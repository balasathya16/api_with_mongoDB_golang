package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"userapi.com/models"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	return nil
}
