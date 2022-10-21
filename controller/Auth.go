package controller

import (
	"MyGram/helper"
	"MyGram/model"
	"MyGram/param"
	"MyGram/repo"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	userRepo repo.UserRepository
}

func NewAuthController(userRepo repo.UserRepository) *AuthController {
	return &AuthController{userRepo}
}

func (u *AuthController) Register(ctx *gin.Context) {
	req := param.RegisReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("%s has a value of %s which does not satisfy %s", e.Field(), e.Value(), e.Tag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errorMessages, "data": nil})
		return
	}

	// CHECK EMAIL
	userEmail, errUser := u.userRepo.GetUserByEmail(req.Email)
	if errUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUser.Error(),
		})
		return
	}

	if userEmail.Email != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email Sudah Terdaftar, Validasi Unique",
		})
		return
	}

	// CHECK USERNAME
	userUname, errUser := u.userRepo.GetUserByUsername(req.Username)
	if errUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUser.Error(),
		})
		return
	}

	if userUname.Username != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Username Sudah Terdaftar, Validasi Unique",
		})
		return
	}

	// HASH PASSWORD
	hash, errHash := helper.GeneratePassword([]byte(req.Password))
	if errHash != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errHash.Error(),
		})
		return
	}

	// INSERT TO DATABASE
	currentTime := time.Now()
	data := model.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hash,
		Age:       req.Age,
		CreatedAt: currentTime.Format("2006-01-02 15:04:05"),
	}

	dataRegis, errRegis := u.userRepo.RegisterUser(&data)
	if errRegis != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": errRegis,
		})
		return
	}

	resp := map[string]interface{}{
		"age":      dataRegis.Age,
		"email":    dataRegis.Email,
		"id":       dataRegis.ID,
		"username": dataRegis.Username,
	}

	ctx.JSON(201, gin.H{"status": "success", "payload": resp})
}

func (u *AuthController) Login(ctx *gin.Context) {
	req := param.LoginReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := e.Error()
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errorMessages, "data": nil})
		return
	}

	user, err := u.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "no data",
		})
		return
	}

	if user.Email == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Email is not registered",
		})
		return
	}

	err = helper.ValidatePassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	payload := helper.Token{
		Email:  user.Email,
		UserID: user.ID,
	}

	token, err := helper.GenerateToken(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
