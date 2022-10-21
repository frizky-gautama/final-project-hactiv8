package middleware

import (
	"MyGram/helper"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "need signin",
		})
		return
	}

	bearer := strings.Split(auth, "Bearer ")

	if len(bearer) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "need signin",
		})
		return
	}

	tokStr := bearer[1]

	tok, err := helper.ValidateToken(tokStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("%+v\n", tok)
	ctx.Set("user_id", tok.UserID)
	ctx.Next()
}
