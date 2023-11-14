package handlers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"go-rest-api/utils"
)

func AddTodo(ctx *gin.Context) {
	fmt.Println("POST todo...")
	var newTodo utils.Todo

	if err := ctx.BindJSON(&newTodo); err != nil {
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusServiceUnavailable, err.Error())
		return
	}
	utils.Todos = append(utils.Todos, newTodo)

	ctx.IndentedJSON(http.StatusCreated, utils.Todos)
}
