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
  lintOnSave: false
};
