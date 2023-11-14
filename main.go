package main

import (
	"fmt"
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Do laundry", Completed: true},
	{ID: "3", Item: "Program", Completed: false},
}

func getTodos(ctx *gin.Context) {
	fmt.Println("GET todos...")
	fmt.Println(todos)
	ctx.IndentedJSON(http.StatusOK, todos)	
}

func getTodoById(id string) (*todo, error) {
	for _, t := range todos {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, errors.New("Todo not found")
}

func getTodo(ctx *gin.Context) {
	id := ctx.Param("id")	

	fmt.Println(fmt.Sprintf("GET /todos/%s", id))

	todo, err := getTodoById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	ctx.IndentedJSON(http.StatusOK, todo)
}

func addTodo(ctx *gin.Context) {
	fmt.Println("POST todo...")
	var newTodo todo 

	if err := ctx.BindJSON(&newTodo); err != nil {
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusServiceUnavailable, err.Error())
		return
	}
	todos = append(todos, newTodo)

	ctx.IndentedJSON(http.StatusCreated, todos)
}

func main() {
	// var router = gin.Default() <- equivalent	
	router := gin.Default() 

	// routes
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)

	// server init
	router.Run("localhost:9090")
}
