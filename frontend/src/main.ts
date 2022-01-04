import Vue from "vue";
import App from "@/App.vue";
import "@/registerServiceWorker";
import router from "@/router/index";
import "@/assets/css/styles.scss";
import '@/plugins/bootstrap'
import '@/plugins/cookie'
import '@/plugins/toast'

Vue.prototype.$mapFields = Vue;
Vue.config.productionTip = false;

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");
