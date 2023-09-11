r
<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script lang="ts">
import { Component } from "vue-property-decorator";
import Vue from "vue";

@Component
export default class App extends Vue {}
</script>
