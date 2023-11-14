package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	utils "go-rest-api/utils"
)

func ToggleTodoStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := utils.FindTodo(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	todo.Completed = !todo.Completed

	ctx.IndentedJSON(http.StatusOK, todo)
}
