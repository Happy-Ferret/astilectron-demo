import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    todos: [],
  },
  mutations: {
    fetchTodos(state) {
      document.addEventListener("astilectron-ready", () => {
        astilectron.sendMessage({ name: "getTodos" }, message => {
          if (message.payload != null) {
            state.todos = message.payload;
          }
        });
      });
    },
    addTodo(state, payload) {
      astilectron.sendMessage(
        { name: "create", payload: payload.title },
        message => {
          state.todos.push({
            title: payload.title,
            done: false,
            id: message.payload
          });
        }
      );
    },
    removeTodo(state, payload) {
      const todoIndex = state.todos.indexOf(payload);
      state.todos.splice(todoIndex, 1);
      astilectron.sendMessage(
        { name: "delete", payload: payload.id }
      );
    },
    updateTodo(state, payload) {
      astilectron.sendMessage(
        { name: "update", payload: { done: payload.done, id: payload.id } },
      );
    }
  },
  actions: {},
  getters: {
    todos: state => state.todos,
  }
})
