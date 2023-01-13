package services

import "userapi.com/models"

type UserService interface {
	CreateUser(*models.User, error)
	GetUser(*string) (*models.User, error)
	Getall() ([]*models.User, error)
	UpdateUser(*models.User, error)
	DeleteUser(*string) error
}
