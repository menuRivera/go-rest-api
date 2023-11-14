package handlers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	utils "go-rest-api/utils"
)

func GetTodos(ctx *gin.Context) {
	fmt.Println("GET todos...")
	fmt.Println(utils.Todos)
	ctx.IndentedJSON(http.StatusOK, utils.Todos)
}
