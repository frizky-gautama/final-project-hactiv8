package repo

import (
	"MyGram/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	InsertComment(comment *model.Comment) (*model.Comment, error)
	GetAllComment() (*[]model.Comment, error)
	UpdateComment(id int, photo *model.Comment) error
	DeleteComment(id int) error
	GetDetailComment(id int32) (model.Comment, error)
}

type connectionComment struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *connectionComment {
	return &connectionComment{db}
}

func (r *connectionComment) InsertComment(comment *model.Comment) (*model.Comment, error) {
	result := r.db.Create(comment)
	return comment, result.Error
}

func (r *connectionComment) GetAllComment() (*[]model.Comment, error) {
	var comments *[]model.Comment
	result := r.db.Model(comments).Preload("User").Preload("Photo").Find(&comments)
	return comments, result.Error
}

func (r *connectionComment) UpdateComment(id int, photo *model.Comment) error {
	result := r.db.Where("id = ?", id).Updates(photo)
	return result.Error
}

func (r *connectionComment) DeleteComment(id int) error {
	var comment model.Comment
	result := r.db.Where("id = ?", id).Delete(&comment).Error
	return result
}

func (r *connectionComment) GetDetailComment(id int32) (model.Comment, error) {
	var comment model.Comment
	result := r.db.Where("id = ?", id).Find(&comment)
	return comment, result.Error
}
