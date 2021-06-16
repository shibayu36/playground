import React, { useEffect, useReducer } from "react";

import API, { graphqlOperation } from "@aws-amplify/api";
import PubSub from "@aws-amplify/pubsub";

import { createTodo, createTodoList } from "./graphql/mutations";
import { listTodoLists } from "./graphql/queries";
import { onCreateTodo } from "./graphql/subscriptions";

import awsconfig from "./aws-exports";
import "./App.css";

API.configure(awsconfig);
PubSub.configure(awsconfig);

// Action Types
const QUERY = "QUERY";
const SUBSCRIPTION = "SUBSCRIPTION";

const initialState = {
  todoLists: [],
};

const reducer = (state, action) => {
  switch (action.type) {
    case QUERY:
      return { ...state, todoLists: action.todoLists };
    case SUBSCRIPTION:
      return { ...state, todos: [...state.todos, action.todo ]}
    default:
      return state;
  }
}

async function createNewTodo() {
  const todo = { name: "Use AWS AppSync", description: "RealTime and Offline" };
  await API.graphql(graphqlOperation(createTodo, { input: todo }));
}

async function createNewTodoList(num) {
  const todoList = { name: `New TodoList ${num}` };
  await API.graphql(graphqlOperation(createTodoList, { input: todoList }));
}

function App() {
  const [state, dispatch] = useReducer(reducer, initialState);

  useEffect(() => {
    // TodoListとTodo含めて全体を一斉に読み込むようにしたい
    async function getData() {
      const todoListData = await API.graphql(graphqlOperation(listTodoLists));
      dispatch({ type: QUERY, todoLists: todoListData.data.listTodoLists.items });
    }
    getData();

    // const subscription = API.graphql(graphqlOperation(onCreateTodo)).subscribe({
    //   next: (eventData) => {
    //     const todo = eventData.value.data.onCreateTodo;
    //     dispatch({ type: SUBSCRIPTION, todo });
    //   }
    // })

  //   return () => subscription.unsubscribe();
  }, []);

  // TodoListを作れるように
  // TodoList配下にTodoを作れるように
  return (
    <div className="App">
      <h1>Todo List</h1>
      <button onClick={() => { createNewTodoList(state.todoLists.length + 1) }}>Add TodoList</button>
      <div>
        {state.todoLists.length > 0 ?
          state.todoLists.map((todoList) => <p key={todoList.id}>{todoList.name}</p>) :
          <p>Add some todoLists!</p>
        }
      </div>
    </div>
  )
}

export default App;
