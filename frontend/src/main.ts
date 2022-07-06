

import "bootstrap/dist/css/bootstrap.css"; 
//import { BootstrapVue } from "bootstrap-vue";
import { createApp  } from 'vue'
import App from './App.vue'

import router from "./router";
//import { Model } from "vue-api-query";

//createApp.config.productionTip = false;


createApp(App).use(router).mount('#app')
import "bootstrap/dist/js/bootstrap.js"; 