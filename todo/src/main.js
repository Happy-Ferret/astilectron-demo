import Vue from 'vue'
import App from './App.vue'
import store from './store'

import AboutVue from './views/About.vue'

Vue.config.productionTip = false

const routes = {
  '/': App,
  '/about': AboutVue
}

new Vue({
  store,
  data: {
    currentRoute: window.location.pathname
  },
  mounted() {
    document.addEventListener('astilectron-ready', () => {
      this.listen();
    })
  },
  computed: {
    ViewComponent() {
      return routes[this.currentRoute] || App
    }
  },
  methods: {
    listen() {
      astilectron.onMessage(function (message) {
        switch (message.name) {
          case "about":
            window.open("/about");
            break;
        }
      });
    }
  },
  render(h) { return h(this.ViewComponent) }
}).$mount('#app')
