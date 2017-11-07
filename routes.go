package main

import "net/http"

// Route handles a HTTP request by method and path
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Routes is a table of HTTP routes/handlers
type Routes []Route

// routes are our table of routes
var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "TodoIndex",
        "GET",
        "/todos",
        TodoIndex,
    },
    Route{
        "TodoShow",
        "GET",
        "/todos/{todoId}",
        TodoShow,
    },
    Route{
        "TodoCreate",
        "POST",
        "/todos",
        TodoCreate,
    },
    Route{
        "TodoDelete",
        "DELETE",
        "/todos/{todoId}",
        TodoDelete,
    },
}
