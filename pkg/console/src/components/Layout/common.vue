<template>
  <div class="app-container">
    <div class="app-box">
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class AppLayout extends Vue {}
</script>

<style lang="scss" scoped>
.app-container {
  position: absolute;
  top: 2px;
  right: 2px;
  bottom: 2px;
  left: 2px;
  overflow: auto;
}
</style>
