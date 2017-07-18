package main

import "time"

// Todo represents a to do task.
type Todo struct {
    Id        int       `json:"id"`
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

// Table of existing Todo 
type Todos []Todo
