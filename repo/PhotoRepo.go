package repo

import (
	"MyGram/model"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	InsertPhoto(photo *model.Photo) (int32, error)
	ResponsePostPhoto(id int32) (model.ResponsePostPhoto, error)
	GetAllPhoto() (*[]model.Photo, error)
	GetUserPhoto(id int32) (model.User, error)
	UpdatePhoto(id int, photo model.Photo) error
	GetDetailPhoto(id int32) (model.Photo, error)
	DeletePhoto(id int) error
}

type connectionPhoto struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *connectionPhoto {
	return &connectionPhoto{db}
}

func (r *connectionPhoto) InsertPhoto(photo *model.Photo) (int32, error) {
	result := r.db.Create(photo)
	return photo.ID, result.Error
}

func (r *connectionPhoto) ResponsePostPhoto(id int32) (model.ResponsePostPhoto, error) {
	var photos model.ResponsePostPhoto
	result := r.db.Where("id = ?", id).Find(&photos).Error
	return photos, result
}

func (r *connectionPhoto) GetAllPhoto() (*[]model.Photo, error) {
	var photos *[]model.Photo
	result := r.db.Model(photos).Preload("User").Find(&photos)
	return photos, result.Error
}

func (r *connectionPhoto) GetUserPhoto(id int32) (model.User, error) {
	var users model.User
	result := r.db.Where("id = ?", id).Find(&users)
	return users, result.Error
}

func (r *connectionPhoto) UpdatePhoto(id int, photo model.Photo) error {
	result := r.db.Where("id = ?", id).Updates(&photo)
	return result.Error
}

func (r *connectionPhoto) GetDetailPhoto(id int32) (model.Photo, error) {
	var photos model.Photo
	result := r.db.Where("id = ?", id).Find(&photos)
	return photos, result.Error
}

func (r *connectionPhoto) DeletePhoto(id int) error {
	var photo model.Photo
	result := r.db.Where("id = ?", id).Delete(&photo).Error
	return result
}
