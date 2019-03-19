import Vue from 'vue'
import App from './App.vue'
import store from './store'

Vue.config.productionTip = false

new Vue({
  store,
<<<<<<< HEAD
  render: h => h(App)
=======
  data: {
    currentRoute: window.location.pathname
  },
  mounted() {
    document.addEventListener('astilectron-ready', () => {
      this.listen();
      this.send();
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
>>>>>>> 35eae05... WIP
}).$mount('#app')
