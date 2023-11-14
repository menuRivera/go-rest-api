package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	utils "go-rest-api/utils"
)

func GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	fmt.Println(fmt.Sprintf("GET /todos/%s", id))

	todo, err := utils.FindTodo(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	ctx.IndentedJSON(http.StatusOK, todo)
}
