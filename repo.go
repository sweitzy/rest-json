package main

import "fmt"

// currentId is index of most recently created Todo
var currentId int

// todos are our table of todos
var todos Todos

/* This is for debug only.
// Give us some seed data
func init() {
    RepoCreateTodo(Todo{Name: "Write presentation"})
    RepoCreateTodo(Todo{Name: "Host meetup"})
}
*/

// RepoFindTodo searches for a Todo by index
func RepoFindTodo(id int) Todo {
    for _, t := range todos {
        if t.Id == id {
            return t
        }
    }
    // return empty Todo if not found
    return Todo{}
}

// RepoCreate creates a new todo in our todo table
func RepoCreateTodo(t Todo) Todo {
    currentId += 1
    t.Id = currentId
    todos = append(todos, t)
    return t
}

// RepoDestroy destroys an existing todo  
func RepoDestroyTodo(id int) error {
    for i, t := range todos {
        if t.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
