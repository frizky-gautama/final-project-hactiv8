package repo

import (
	"MyGram/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	RegisterUser(user *model.User) (*model.User, error)
	ResponseRegister(email string) (model.RegResponse, error)
	UpdateUser(id int, user model.User) error
	ResponseUpdate(id int) (model.UpdateResponse, error)
	DeleteUser(id int) error
}

type connectionUser struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *connectionUser {
	return &connectionUser{db}
}

func (r *connectionUser) RegisterUser(user *model.User) (*model.User, error) {
	result := r.db.Create(user)
	return user, result.Error
}

func (r *connectionUser) GetUserByEmail(email string) (model.User, error) {
	var users model.User
	result := r.db.Where("email = ?", email).Find(&users).Error
	return users, result
}

func (r *connectionUser) GetUserByUsername(username string) (model.User, error) {
	var users model.User
	result := r.db.Where("username = ?", username).Find(&users).Error
	return users, result
}

func (r *connectionUser) ResponseRegister(email string) (model.RegResponse, error) {
	var users model.RegResponse
	result := r.db.Where("username = ?", email).Find(&users).Error
	return users, result
}

func (r *connectionUser) UpdateUser(id int, user model.User) error {
	result := r.db.Where("id = ?", id).Updates(&user)
	return result.Error
}

func (r *connectionUser) ResponseUpdate(id int) (model.UpdateResponse, error) {
	var users model.UpdateResponse
	result := r.db.Where("id = ?", id).Find(&users).Error
	return users, result
}

func (r *connectionUser) DeleteUser(id int) error {
	var users model.User
	result := r.db.Where("id = ?", id).Delete(&users).Error
	return result
}
