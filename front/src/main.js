import Vue from 'vue'
import App from './App.vue'
import router from './router'

import Loading from 'vue-loading-overlay';
import 'vue-loading-overlay/dist/vue-loading.css';
Vue.use(Loading);

import ApiService from "./common/api.service";
ApiService.init()

Vue.config.productionTip = false

require("bulma/css/bulma.min.css")

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
