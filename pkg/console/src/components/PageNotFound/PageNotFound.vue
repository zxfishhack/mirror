<template>
  <div style="margin: 30px" class="page-not-found">
    <div style="margin: auto; width: 50%">
      <div class="boxIcon hover" style="text-align: center">
        <i class="icon-unlink iconBig iconRounded"></i>
        <h1>页面无法找到</h1>
        <p>您所请求的页面可能被移除或者暂时无法访问。</p>
        <router-link
          :to="{ name: 'Home' }"
          class="btn btn-lg btn-primary"
          style="margin-top: 15px"
          >回到首页</router-link
        >
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component } from "vue-property-decorator";
import Vue from "vue";
import "@/assets/font-icons/font-icons.scss";

@Component
export default class PageNotFound extends Vue {}
</script>

<style lang="scss" scoped>
.btn:active,
.btn.active {
  -webkit-box-shadow: none;
  box-shadow: none;
}
.btn {
  background: #fff;
  color: lighten($foreground, 6.5%);
  border: 1px solid rgba(0, 0, 0, 0.15);
}
.btn:hover {
  background: $highlight;
  border-color: $highlight;
  color: $background;
}
.btn-primary {
  background: $default;
  color: $background !important;
}
</style>
