import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './assets/tailwind.css'
// import store from './store'
import FlashMessage from '@smartweb/vue-flash-message'

Vue.config.productionTip = false;
Vue.use(FlashMessage, {
  strategy: 'multiple'
});

new Vue({
  router,
  // store,
  render: h => h(App)
}).$mount('#app');
