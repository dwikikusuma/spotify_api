package memberships

import (
	"spotify_api/internal/configs"
	"spotify_api/internal/model/memberhsips"
)

type Repository interface {
	CreateUser(model memberhsips.User) error
	GetUser(username, email string, id uint) (*memberhsips.User, error)
}

type Service struct {
	config     *configs.Config
	repository Repository
}

func NewService(repo Repository, config *configs.Config) *Service {
	return &Service{
		repository: repo,
		config:     config,
	}
}
