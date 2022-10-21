package controller

import (
	"MyGram/helper"
	"MyGram/model"
	"MyGram/param"
	"MyGram/repo"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userRepo repo.UserRepository
}

func NewUserController(userRepo repo.UserRepository) *UserController {
	return &UserController{userRepo}
}

func (u *UserController) UpdateUser(ctx *gin.Context) {

	// PARAM ID
	id_param := helper.StrToInt(ctx.Param("id"))

	// PAYLOAD
	payload := param.UpdateUserPayload

	err := ctx.ShouldBindJSON(&payload)
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
	userEmail, errUser := u.userRepo.GetUserByEmail(payload.Email)
	if errUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUser.Error(),
		})
		return
	}

	if userEmail.Email != "" && !strings.EqualFold(userEmail.Email, payload.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUser.Error(),
		})
		return
	}

	// CHECK USERNAME
	userUname, errUser := u.userRepo.GetUserByUsername(payload.Username)
	if errUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUser.Error(),
		})
		return
	}

	if userUname.Username != "" && !strings.EqualFold(userUname.Username, payload.Username) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUser.Error(),
		})
		return
	}

	// Data request for update
	var email string
	var uname string

	if payload.Email == "" {
		email = userEmail.Email
	} else {
		email = payload.Email
	}

	if payload.Username == "" {
		uname = userEmail.Username
	} else {
		uname = payload.Username
	}

	currentTime := time.Now()
	dataUpdate := model.User{
		Email:     email,
		Username:  uname,
		UpdatedAt: currentTime.Format("2006-01-02 15:04:05"),
	}

	// ACTION UPDATE
	errUpdate := u.userRepo.UpdateUser(id_param, dataUpdate)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdate.Error(),
		})
		return
	}

	// RESPONSE
	resp, errResp := u.userRepo.ResponseUpdate(id_param)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errResp.Error(),
		})
		return
	}

	dataResp := map[string]interface{}{
		"id":         resp.ID,
		"email":      resp.Email,
		"username":   resp.Username,
		"age":        resp.Age,
		"updated_at": resp.UpdatedAt,
	}

	ctx.JSON(200, gin.H{"status": "success", "payload": dataResp})
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	// PARAM ID
	id_param := helper.StrToInt(ctx.Param("id"))

	// ACTION UPDATE
	err := u.userRepo.DeleteUser(id_param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{"message": "Your account has been successfully deleted"})
}
