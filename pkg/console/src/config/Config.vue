<template>
  <div style="background: white; height: 100%; padding: 10px">
    <div style="display: flex; margin-top: 10px; margin-bottom: 10px">
      <el-button style="flex: 1" type="primary" @click="activeRules"
        >生效所有配置</el-button
      >
      <el-button style="flex: 1" type="primary" @click="popAdd">新增</el-button>
    </div>
    <el-table :data="rules" style="width: 100%">
      <el-table-column width="50">
        <template slot-scope="scope">
          <div style="display: flex">
            <i
              :style="{ color: scope.row.active ? '#2d7f37' : '' }"
              style="margin: auto; font-size: 20px"
              class="el-icon-s-opportunity"
            />
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="prefix" label="文件名">
        <template slot-scope="scope">
          <p>前缀：{{ scope.row.prefix }}</p>
          <p>后缀：{{ scope.row.postfix }}</p>
        </template>
      </el-table-column>
      <el-table-column prop="upstream" label="回源配置">
        <template slot-scope="scope">
          <p>回源地址：{{ scope.row.upstream }}</p>
          <p>前缀替换：{{ scope.row.replacePrefixWith }}</p>
        </template>
      </el-table-column>
      <el-table-column prop="checkMD5" label="校验MD5" width="100">
        <template slot-scope="scope">
          {{ scope.row.checkMD5 ? "是" : "否" }}
        </template>
      </el-table-column>
      <el-table-column width="200" label="操作">
        <template slot-scope="scope">
          <div style="display: flex; gap: 10px">
            <el-link type="primary" @click="popEdit(scope.row)">编辑</el-link>
            <el-popconfirm
              title="确定要删除吗？"
              @confirm="deleteRule(scope.row.id)"
            >
              <el-link slot="reference" type="danger">删除</el-link>
            </el-popconfirm>
          </div>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
      :title="this.editing ? '编辑...' : '新增'"
      :visible.sync="show"
      :before-close="handleClose"
    >
      <el-form>
        <el-form-item label="文件名前缀">
          <el-input v-model="rule.prefix" />
        </el-form-item>
        <el-form-item label="文件名后缀">
          <el-input v-model="rule.postfix" />
        </el-form-item>
        <el-form-item label="回源地址">
          <el-input v-model="rule.upstream" />
        </el-form-item>
        <el-form-item label="前缀替换">
          <el-input v-model="rule.replacePrefixWith" />
        </el-form-item>
        <el-form-item label="校验MD5">
          <el-switch v-model="rule.checkMD5" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="show = false">取 消</el-button>
        <el-button type="primary" @click="confirmAddOrEdit">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import { Component, Mixins, Watch } from "vue-property-decorator";
import Vue from "vue";
import Service from "@/service";

const emptyRule: Rule = {
  active: false,
  checkMD5: false,
  id: "0",
  postfix: "",
  prefix: "",
  replacePrefixWith: "",
  upstream: "",
};

@Component
export default class ConfigView extends Vue {
  rules: Rule[] = [];
  rule: Rule = { ...emptyRule };
  editing = false;
  show = false;

  handleClose(done: Function) {
    this.$confirm("确认关闭？")
      .then(() => {
        done();
      })
      .catch((_) => {});
  }
  popAdd() {
    this.editing = false;
    this.rule = { ...emptyRule };
    this.show = true;
  }

  popEdit(rule: Rule) {
    this.editing = true;
    this.rule = rule;
    this.show = true;
  }

  async confirmAddOrEdit() {
    if (this.editing) {
      await Service.editRule(this.rule);
    } else {
      await Service.addRule(this.rule);
    }
    await this.loadData();
    this.show = false;
  }

  async activeRules() {
    await Service.activeRules();
    await this.loadData();
  }

  async loadData() {
    this.rules = await Service.getRules();
  }

  async deleteRule(id: string) {
    await Service.deleteRule(id);
    await this.loadData();
  }

  mounted() {
    this.loadData().then(() => {
      console.log(this.rules);
    });
  }
}
</script>

<style scoped lang="scss">
::v-deep .el-form-item__content {
  display: flex;
}
.max-size {
  flex: 1;
}
</style>
