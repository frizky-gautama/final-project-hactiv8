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

type PhotoController struct {
	photoRepo repo.PhotoRepository
}

func NewPhotoController(photoRepo repo.PhotoRepository) *PhotoController {
	return &PhotoController{photoRepo}
}

func (p *PhotoController) InsertPhoto(ctx *gin.Context) {

	// PAYLOAD
	payload := param.InsertPhoto

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
	data := model.Photo{
		Title:     payload.Title,
		Caption:   payload.Caption,
		Photo_url: payload.Photo_url,
		UserID:    parseUserID,
		CreatedAt: currentTime.Format("2006-01-02 15:04:05"),
	}

	id, err := p.photoRepo.InsertPhoto(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// RESPONSE
	resp, errResp := p.photoRepo.ResponsePostPhoto(id)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errResp.Error(),
		})
		return
	}

	dataResp := map[string]interface{}{
		"id":        resp.ID,
		"title":     resp.Title,
		"caption":   resp.Caption,
		"photo_url": resp.Photo_url,
		"user_id":   resp.UserID,
	}

	ctx.JSON(201, gin.H{"status": "success", "payload": dataResp})

}

func (p *PhotoController) GetAllFoto(ctx *gin.Context) {
	// RESPONSE
	resp, errResp := p.photoRepo.GetAllPhoto()
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errResp.Error(),
		})
		return
	}

	var response []param.ResponsePhoto
	if resp != nil {
		for _, val := range *resp {
			data := param.ResponsePhoto{
				ID:        val.ID,
				Title:     val.Title,
				Caption:   val.Caption,
				Photo_url: val.Photo_url,
				UserID:    val.UserID,
				CreatedAt: val.CreatedAt,
				// Username:      val.User,
			}
			if val.User != nil {
				var respPhoto = param.ResponseUserPhoto{}
				respPhoto.Username = val.User.Username
				respPhoto.Email = val.User.Email
				data.User = respPhoto
			}
			response = append(response, data)
		}
	}

	ctx.JSON(200, gin.H{"status": "success", "payload": response})
}

func (u *PhotoController) UpdatePhoto(ctx *gin.Context) {

	// PARAM ID
	id_param := helper.StrToInt(ctx.Param("id"))

	// PAYLOAD
	payload := param.UpdatePhoto

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
	dataUpdate := model.Photo{
		Title:     payload.Title,
		Caption:   payload.Caption,
		Photo_url: payload.Photo_url,
		UpdatedAt: currentTime.Format("2006-01-02 15:04:05"),
	}

	// ACTION UPDATE
	errUpdate := u.photoRepo.UpdatePhoto(id_param, dataUpdate)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdate.Error(),
		})
		return
	}

	// RESPONSE
	resp, errResp := u.photoRepo.GetDetailPhoto(int32(id_param))
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errResp.Error(),
		})
		return
	}

	dataResp := param.ResponseUpdatePhoto{
		ID:        resp.ID,
		Title:     resp.Title,
		Caption:   resp.Caption,
		Photo_url: resp.Photo_url,
		UserID:    resp.UserID,
		UpdatedAt: resp.UpdatedAt,
	}

	ctx.JSON(200, gin.H{"status": "success", "payload": dataResp})
}

func (u *PhotoController) DeletePhoto(ctx *gin.Context) {
	id_param := helper.StrToInt(ctx.Param("id"))
	// ACTION UPDATE
	err := u.photoRepo.DeletePhoto(id_param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{"message": "Your Photo has been successfully deleted"})
}
