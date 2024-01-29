package apis

import (
	"authExample/auth"

	"github.com/gin-gonic/gin"
)

func HelloWord(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "helloWorld"})
}

func GenToken(ctx *gin.Context) {
	token, _ := auth.GenerateToken()
	ctx.JSON(200, gin.H{"Token": token})
}
