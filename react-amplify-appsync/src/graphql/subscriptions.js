/* eslint-disable */
// this is an auto generated file. This will be overwritten

export const onCreateTodoList = /* GraphQL */ `
  subscription OnCreateTodoList {
    onCreateTodoList {
      id
      name
      items {
        nextToken
      }
      createdAt
      updatedAt
    }
  }
`;
export const onUpdateTodoList = /* GraphQL */ `
  subscription OnUpdateTodoList {
    onUpdateTodoList {
      id
      name
      items {
        nextToken
      }
      createdAt
      updatedAt
    }
  }
`;
export const onDeleteTodoList = /* GraphQL */ `
  subscription OnDeleteTodoList {
    onDeleteTodoList {
      id
      name
      items {
        nextToken
      }
      createdAt
      updatedAt
    }
  }
`;
export const onCreateTodo = /* GraphQL */ `
  subscription OnCreateTodo {
    onCreateTodo {
      id
      todoListId
      name
      description
      createdAt
      updatedAt
    }
  }
`;
export const onUpdateTodo = /* GraphQL */ `
  subscription OnUpdateTodo {
    onUpdateTodo {
      id
      todoListId
      name
      description
      createdAt
      updatedAt
    }
  }
`;
export const onDeleteTodo = /* GraphQL */ `
  subscription OnDeleteTodo {
    onDeleteTodo {
      id
      todoListId
      name
      description
      createdAt
      updatedAt
    }
  }
`;
