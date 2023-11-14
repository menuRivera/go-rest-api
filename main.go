package main

import (
	"go-rest-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// var router = gin.Default() <- equivalent
	router := gin.Default()

	// routes
	router.GET("/todos", handlers.GetTodos)
	router.GET("/todos/:id", handlers.GetTodo)
	router.PATCH("/todos/:id", handlers.ToggleTodoStatus)
	router.POST("/todos", handlers.AddTodo)

	// server init
	router.Run("localhost:9090")
}
