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

type SocmedController struct {
	socmedRepo repo.SocmedRepository
}

func NewSocmedController(socmedRepo repo.SocmedRepository) *SocmedController {
	return &SocmedController{socmedRepo}
}

func (p *SocmedController) InsertSocmed(ctx *gin.Context) {
	// PAYLOAD
	payload := param.InsertSocmed

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
	data := model.SocialMedia{
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		UserID:         parseUserID,
		CreatedAt:      currentTime.Format("2006-01-02 15:04:05"),
	}

	resp, err := p.socmedRepo.InsertSocmed(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dataResp := param.ResponseInsertSocmed{
		ID:             resp.ID,
		Name:           resp.Name,
		UserID:         resp.UserID,
		SocialMediaUrl: resp.SocialMediaUrl,
		CreatedAt:      resp.CreatedAt,
	}

	ctx.JSON(201, gin.H{"status": "success", "payload": dataResp})
}

func (p *SocmedController) GetAllSocmed(ctx *gin.Context) {
	// RESPONSE
	resp, errResp := p.socmedRepo.GetAllSocmed()
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errResp.Error(),
		})
		return
	}

	var response []param.ResponseSocmed
	if resp != nil {
		for _, val := range *resp {
			data := param.ResponseSocmed{
				ID:             val.ID,
				Name:           val.Name,
				SocialMediaUrl: val.SocialMediaUrl,
				UserID:         val.UserID,
				CreatedAt:      val.CreatedAt,
				// Username:      val.User,
			}
			if val.User != nil {
				var respSocmed = param.ResponseUserSocmed{}
				respSocmed.ID = val.User.ID
				respSocmed.Username = val.User.Username
				respSocmed.Email = val.User.Email
				data.User = respSocmed
			}
			response = append(response, data)
		}
	}

	ctx.JSON(200, gin.H{"status": "success", "payload": response})
}

func (u *SocmedController) UpdateSocmed(ctx *gin.Context) {

	// PARAM ID
	id_param := helper.StrToInt(ctx.Param("id"))

	// PAYLOAD
	payload := param.UpdateSocmed

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
	dataUpdate := model.SocialMedia{
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		UpdatedAt:      currentTime.Format("2006-01-02 15:04:05"),
	}

	// ACTION UPDATE
	errUpdate := u.socmedRepo.UpdateSocmed(id_param, &dataUpdate)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdate.Error(),
		})
		return
	}

	// RESPONSE
	resp, errResp := u.socmedRepo.GetDetailSocmed(int32(id_param))
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errResp.Error(),
		})
		return
	}

	dataResp := param.ResponseUpdateSocmed{
		ID:             resp.ID,
		Name:           resp.Name,
		SocialMediaUrl: resp.SocialMediaUrl,
		UserID:         resp.UserID,
		UpdatedAt:      resp.UpdatedAt,
	}

	ctx.JSON(200, gin.H{"status": "success", "payload": dataResp})
}

func (u *SocmedController) DeleteSocmed(ctx *gin.Context) {
	id_param := helper.StrToInt(ctx.Param("id"))
	// ACTION UPDATE
	err := u.socmedRepo.DeleteSocmed(id_param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{"message": "Your Social Media has been successfully deleted"})
}
