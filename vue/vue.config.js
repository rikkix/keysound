module.exports = {
  devServer: {
    proxy: {
      '/api/new': {
        target: 'http://linux-device:8088',
        secure: false
      },
      '/api/get': {
        target: 'http://linux-device:8088',
        secure: false
      }
    }
  },
  lintOnSave: false,
  chainWebpack: config => {
    config
      .plugin('html')
      .tap(args => {
        args[0].title = "Key Sound";
        return args;
      })
  }
};
