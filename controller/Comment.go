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

type CommentController struct {
	commentRepo repo.CommentRepository
}

func NewCommentController(commentRepo repo.CommentRepository) *CommentController {
	return &CommentController{commentRepo}
}

func (p *CommentController) InsertComment(ctx *gin.Context) {
	// PAYLOAD
	payload := param.InsertComment

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

	userID := ctx.MustGet("user_id")
	parseUserID := int32(userID.(int32))

	currentTime := time.Now()
	data := model.Comment{
		Message:   payload.Message,
		PhotoID:   payload.Photo_id,
		UserID:    parseUserID,
		CreatedAt: currentTime.Format("2006-01-02 15:04:05"),
	}

	resp, err := p.commentRepo.InsertComment(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dataResp := param.ResponseInsertComment{
		ID:        resp.ID,
		UserID:    resp.UserID,
		PhotoID:   resp.PhotoID,
		Message:   resp.Message,
		CreatedAt: resp.CreatedAt,
	}

	ctx.JSON(201, gin.H{"status": "success", "payload": dataResp})
}

func (p *CommentController) GetAllComment(ctx *gin.Context) {
	// RESPONSE
	resp, errResp := p.commentRepo.GetAllComment()
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errResp.Error(),
		})
		return
	}

	var response []param.ResponseComment
	if resp != nil {
		for _, val := range *resp {
			data := param.ResponseComment{
				ID:        val.ID,
				Message:   val.Message,
				PhotoID:   val.PhotoID,
				UserID:    val.UserID,
				UpdatedAt: val.UpdatedAt,
				CreatedAt: val.CreatedAt,
			}
			if val.User != nil {
				var respUser = param.ResponseUserComment{}
				respUser.ID = val.User.ID
				respUser.Username = val.User.Username
				respUser.Email = val.User.Email
				data.User = respUser
			}

			if val.Photo != nil {
				var respPhoto = param.ResponsePhotoComment{}
				respPhoto.ID = val.Photo.ID
				respPhoto.Title = val.Photo.Title
				respPhoto.Caption = val.Photo.Caption
				respPhoto.Photo_url = val.Photo.Photo_url
				respPhoto.UserID = val.Photo.UserID
				data.Photo = respPhoto
			}
			response = append(response, data)
		}
	}

	ctx.JSON(200, gin.H{"status": "success", "payload": response})
}

func (u *CommentController) UpdateComment(ctx *gin.Context) {

	// PARAM ID
	id_param := helper.StrToInt(ctx.Param("id"))

	// PAYLOAD
	payload := param.UpdateComment

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

	currentTime := time.Now()
	dataUpdate := model.Comment{
		Message:   payload.Message,
		UpdatedAt: currentTime.Format("2006-01-02 15:04:05"),
	}

	// ACTION UPDATE
	errUpdate := u.commentRepo.UpdateComment(id_param, &dataUpdate)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdate.Error(),
		})
		return
	}

	// RESPONSE
	resp, errResp := u.commentRepo.GetDetailComment(int32(id_param))
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errResp.Error(),
		})
		return
	}

	dataResp := param.ResponseUpdateComment{
		ID:        resp.ID,
		Message:   resp.Message,
		PhotoID:   resp.PhotoID,
		UserID:    resp.UserID,
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
	}

	ctx.JSON(200, gin.H{"status": "success", "payload": dataResp})
}

func (u *CommentController) DeleteComment(ctx *gin.Context) {
	id_param := helper.StrToInt(ctx.Param("id"))
	// ACTION UPDATE
	err := u.commentRepo.DeleteComment(id_param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{"message": "Your Comment has been successfully deleted"})
}
