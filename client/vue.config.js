module.exports = {
  lintOnSave: process.env.NODE_ENV !== "production",
  devServer: {
    proxy: {
      "^/api": {
        target: "localhost:3000",
        pathRewrite: {
          "^/api/": "/"
        }
      }
    },
    overlay: {
      warnings: true,
      errors: true
    }
  }
};
