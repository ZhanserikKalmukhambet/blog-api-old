package service

import (
	"context"
	"github.com/zhanserikKalmukhambet/blog-api/internal/entity"
	"github.com/zhanserikKalmukhambet/blog-api/pkg/util"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = m.Repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}
