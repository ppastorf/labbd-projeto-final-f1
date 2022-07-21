
import Vue from 'vue'
import App from './App.vue'
import router from './router'

Vue.config.productionTip = false

// Vue.mixin({
//   methods: {
//     download: function(content, filename, contentType) {
//       if(!contentType) contentType = 'application/octet-stream';
//       var a = document.createElement('a');
//       var blob = new Blob([content], {'type':contentType});
//       a.href = window.URL.createObjectURL(blob);
//       a.download = filename;
//       a.click();
//     },
//     parseCookie: function(str) {
//       str
//       .split(';')
//       .map(v => v.split('='))
//       .reduce((acc, v) => {
//         acc[decodeURIComponent(v[0].trim())] = decodeURIComponent(v[1].trim());
//         return acc;
//       }, {});
//     }
//   },
// })

new Vue({
  render: h => h(App),
  router
}).$mount('#app')