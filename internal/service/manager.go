package service

import (
	"github.com/zhanserikKalmukhambet/blog-api/internal/config"
	"github.com/zhanserikKalmukhambet/blog-api/internal/repository"
)

type Manager struct {
	Repository repository.Repository
	Config     *config.Config
}

func New(repository repository.Repository, config *config.Config) *Manager {
	return &Manager{
		Repository: repository,
		Config:     config,
	}
}
