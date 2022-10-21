package repo

import (
	"MyGram/model"

	"gorm.io/gorm"
)

type SocmedRepository interface {
	InsertSocmed(socmed *model.SocialMedia) (*model.SocialMedia, error)
	GetAllSocmed() (*[]model.SocialMedia, error)
	UpdateSocmed(id int, socmed *model.SocialMedia) error
	GetDetailSocmed(id int32) (model.SocialMedia, error)
	DeleteSocmed(id int) error
}

type connectionSocmed struct {
	db *gorm.DB
}

func NewSocmedRepository(db *gorm.DB) *connectionSocmed {
	return &connectionSocmed{db}
}

func (r *connectionSocmed) InsertSocmed(socmed *model.SocialMedia) (*model.SocialMedia, error) {
	result := r.db.Create(socmed)
	return socmed, result.Error
}

func (r *connectionSocmed) GetAllSocmed() (*[]model.SocialMedia, error) {
	var socmed *[]model.SocialMedia
	result := r.db.Model(socmed).Preload("User").Find(&socmed)
	return socmed, result.Error
}

func (r *connectionSocmed) UpdateSocmed(id int, socmed *model.SocialMedia) error {
	result := r.db.Where("id = ?", id).Updates(socmed)
	return result.Error
}

func (r *connectionSocmed) DeleteSocmed(id int) error {
	var socmed model.SocialMedia
	result := r.db.Where("id = ?", id).Delete(&socmed).Error
	return result
}

func (r *connectionSocmed) GetDetailSocmed(id int32) (model.SocialMedia, error) {
	var socmed model.SocialMedia
	result := r.db.Where("id = ?", id).Find(&socmed)
	return socmed, result.Error
}
