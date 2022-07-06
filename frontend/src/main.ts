

import "bootstrap/dist/css/bootstrap.css"; 
//import { BootstrapVue } from "bootstrap-vue";
import { createApp  } from 'vue'
import App from './App.vue'
import VueCookies from 'vue3-cookies'

import router from "./router";
//import { Model } from "vue-api-query";

//createApp.config.productionTip = false;


createApp(App).use(router).use(VueCookies).mount('#app')
import "bootstrap/dist/js/bootstrap.js"; 