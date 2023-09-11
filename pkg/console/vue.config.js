const path = require("path");

// TODO: get this variable from setting.ts
const devServerPort = 9520;
const mockServerPort = 9528;
const name = "镜像回源配置端";

module.exports = {
  publicPath: "./",
  lintOnSave: process.env.NODE_ENV === "development",
  productionSourceMap: false,
  pages: {
    index: {
      entry: "src/main.ts",
      template: "pages/index.html",
      filename: "index.html",
      chunks: ["chunk-libs", "chunk-app", "index"],
    },
  },
  devServer: {
    port: devServerPort,
    open: true,
    overlay: {
      warnings: false,
      errors: true,
    },
    progress: false,
    proxy: {
      "/": {
        target: process.env.VUE_APP_BASE_HOST,
        changeOrigin: true,
        debug: true,
      },
    },
  },

  pluginOptions: {
    "style-resources-loader": {
      preProcessor: "scss",
      patterns: [
        path.resolve(__dirname, "src/assets/styles/_variables.scss"),
        path.resolve(__dirname, "src/assets/styles/_mixins.scss"),
      ],
    },
  },
  chainWebpack(config) {
    // provide the app's title in webpack's name field, so that
    // it can be accessed in index.html to inject the correct title.
    config.set("name", name);

    // https://webpack.js.org/configuration/devtool/#development
    config.when(process.env.NODE_ENV === "development", (config) =>
      config.devtool("cheap-eval-source-map")
    );

    // remove vue-cli-service's progress output
    // config.plugins.delete("progress");
    // replace with another progress output plugin to solve the this bug:
    // https://github.com/vuejs/vue-cli/issues/4557
    //config
    //  .plugin("simple-progress-webpack-plugin")
    //  .use(require.resolve("simple-progress-webpack-plugin"), [
    //    {
    //      format: "compact",
    //    },
    //  ]);

    config.optimization.splitChunks({
      chunks: "all",
      cacheGroups: {
        default: false,
        libs: {
          name: "chunk-libs",
          test: /[\\/]node_modules[\\/]/,
          priority: 10,
          chunks: "initial", // only package third parties that are initially dependent
        },
        common: {
          name: "chunk-app",
          test: /[\\/]console[\\/]src[\\/]/,
          priority: 5,
        },
      },
    });
  },
};
