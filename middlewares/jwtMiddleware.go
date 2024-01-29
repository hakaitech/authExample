package middlewares

import (
	"authExample/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Validate(ctx *gin.Context) {
	header := ctx.GetHeader("authorization")
	token := strings.Split(header, "Bearer ")[1]
	// check jwt token here and is ok pass
	if _, err := auth.ValidateToken(token); err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		ctx.Next()
	}
}
