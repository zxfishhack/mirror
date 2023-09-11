<template>
  <el-container class="container">
    <el-header>
      <el-container>
        <el-menu :router="true" mode="horizontal">
          <div v-for="router in routers" :key="router.path">
            <el-menu-item v-if="!router.children" :index="router.path">
              <span slot="title">{{ router.name }}</span>
            </el-menu-item>
            <el-submenu v-else :index="router.path">
              <span slot="title">{{ router.name }}</span>
              <el-menu-item
                v-for="subRouter in router.children"
                :index="subRouter.path"
                :key="router.path + '-' + subRouter.path"
              >
                <span slot="title">{{ subRouter.name }}</span>
              </el-menu-item>
            </el-submenu>
          </div>
        </el-menu>
      </el-container>
    </el-header>
    <el-container>
      <el-main style="padding: 0px; position: relative">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script lang="ts">
import { Vue, Watch, Prop, Component } from "vue-property-decorator";
import settings from "@/settings";
import { displayRouters } from "@/router/routers";

@Component({})
export default class AppLayout extends Vue {
  private routers: any[] = [];
  private appName: string = settings.title;

  created() {
    this.routers = displayRouters.map(this.routerProcess(""));
  }

  mounted() {
    console.log(this.$children);
  }

  routerProcess(prefix: string) {
    return (value: any) => {
      let p = prefix + value.path;
      let np = p;
      if (np[np.length - 1] !== "/") {
        np += "/";
      }
      return {
        path: p,
        name: value.meta?.display || value.name,
        children: !value.children
          ? undefined
          : value.children.map(this.routerProcess(np)),
      };
    };
  }
}
</script>

<style lang="scss" scoped>
.container {
  height: 100%;
}
</style>
