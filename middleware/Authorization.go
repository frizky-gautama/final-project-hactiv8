package middleware

import (
	"MyGram/helper"
	"MyGram/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	authorRepo repo.AuthorRepository
}

func NewAuthorController(authorRepo repo.AuthorRepository) *AuthorHandler {
	return &AuthorHandler{authorRepo}
}

func AuthorUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.MustGet("user_id")
		parseUserID := int(userID.(int32))

		param_user_id := helper.StrToInt(ctx.Param("id"))

		if param_user_id != parseUserID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func (p *AuthorHandler) AuthorPhoto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.MustGet("user_id")
		parseUserID := int(userID.(int32))

		param_user_id := helper.StrToInt(ctx.Param("id"))

		check, err := p.authorRepo.CheckPhoto(param_user_id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		if int(check.UserID) != parseUserID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func (p *AuthorHandler) AuthorComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.MustGet("user_id")
		parseUserID := int(userID.(int32))

		param_user_id := helper.StrToInt(ctx.Param("id"))

		check, err := p.authorRepo.CheckComment(param_user_id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		if int(check.UserID) != parseUserID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func (p *AuthorHandler) AuthorSocmed() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.MustGet("user_id")
		parseUserID := int(userID.(int32))

		param_user_id := helper.StrToInt(ctx.Param("id"))

		check, err := p.authorRepo.CheckSocmed(param_user_id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		if int(check.UserID) != parseUserID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}
