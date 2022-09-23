package services

import "gin-mongo-api/models"

type UserService interface {
	CreateUser(*models.CreateUserRequest) (*models.UserDBResponse, error)
}