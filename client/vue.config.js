module.exports = {
  lintOnSave: process.env.NODE_ENV !== "production",
  devServer: {
    // proxy: {
    //   "^/api": {
    //     target: "http://lolachost:3000",
    //     pathRewrite: {
    //       "^/api/": "/"
    //     }
    //   }

    // },
    proxy: {
      "/api/*": {
        target: "http://localhost:3000",
        changeOrigin: true
      }
    }
    // overlay: {
    //   warnings: true,
    //   errors: true
    // }
  }
};
