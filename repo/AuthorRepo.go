package repo

import (
	"MyGram/model"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	CheckPhoto(id int) (model.Photo, error)
	CheckComment(id int) (model.Comment, error)
	CheckSocmed(id int) (model.SocialMedia, error)
}

type connectionAuthor struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *connectionAuthor {
	return &connectionAuthor{db}
}

func (r *connectionAuthor) CheckPhoto(id int) (model.Photo, error) {
	var photo model.Photo
	result := r.db.Where("id = ?", id).Find(&photo).Error
	return photo, result
}

func (r *connectionAuthor) CheckComment(id int) (model.Comment, error) {
	var comments model.Comment
	result := r.db.Where("id = ?", id).Find(&comments).Error
	return comments, result
}

func (r *connectionAuthor) CheckSocmed(id int) (model.SocialMedia, error) {
	var socmed model.SocialMedia
	result := r.db.Where("id = ?", id).Find(&socmed).Error
	return socmed, result
}
