<template>
  <div class="app">
    <div class="container">
      <div class="row header">
        <h1 class="col s6 offset-s3 center-align teal-text">To-Do List!</h1>
      </div>
      <div class="row">
        <form @submit.prevent="submitTodo" class="col s6 offset-s3">
          <div class="input-field">
            <input
              v-model="newTodo"
              id="icon_prefix2"
              class="form-control"
              placeholder="What needs to be done?"
            >
          </div>
          <button class="btn btn-primary col s12">Add</button>
        </form>
      </div>
      <div class="row">
        <ul class="collection col s6 offset-s3">
          <li class="collection-item" v-for="todo in todos" :key="todo.id">
            <p>
              <input
                type="checkbox"
                class="checkbox-round"
                :checked="todo.done"
                @change="todo.done = !todo.done"
              >
              <label for="checkbox"></label>
              <del v-if="todo.done">
                <span>{{todo.title}}</span>
              </del>
              <span v-else>{{todo.title}}</span>
              <span class="delete">
                <a @click.prevent="deleteTodo(todo)">
                  <font-awesome-icon :icon="[ 'fa', 'trash' ]"/>
                </a>
              </span>
            </p>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "app",
  data() {
    return {
      todos: [],
      newTodo: ""
    };
  },
  watch: {
    todos: {
      handler() {
        localStorage.todos = JSON.stringify(this.todos);
      },
      deep: true
    }
  },
  mounted() {
    document.addEventListener("astilectron-ready", () => {
      astilectron.sendMessage({ name: "getTodos" }, message => {
        if (message.payload != null) this.todos = message.payload;
      });
    });
  },
  methods: {
    submitTodo() {
      astilectron.sendMessage(
        { name: "create", payload: this.newTodo },
        message => {
          this.todos.push({
            title: this.newTodo,
            done: false,
            id: message.payload
          });
          this.newTodo = "";
        }
      );
    },
    deleteTodo(todo) {
      const todoIndex = this.todos.indexOf(todo);
      this.todos.splice(todoIndex, 1);
      astilectron.sendMessage(
        { name: "delete", payload: todo.id },
        message => {}
      );
    },
    send(todo) {}
  }
};
</script>

<style>
.header {
  margin-top: 100px;
}
.checkbox-round {
  width: 1.3em;
  height: 1.3em;
  background-color: white;
  border-radius: 50%;
  vertical-align: middle;
  border: 1px solid #ddd;
  -webkit-appearance: none;
  outline: none;
  cursor: pointer;
  float: left;
}
.checkbox-round:checked {
  background-color: green;
}
.collection {
  list-style: none;
}

.delete {
  float: right;
}
</style>