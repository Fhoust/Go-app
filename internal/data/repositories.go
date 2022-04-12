package repositories

import (
	"context"
	"github.com/zerepl/go-app/internal/model"
)

type UserRepository interface {
	GetUser(context.Context, int64) (*model.User, error)
	GetAllUsers(context.Context) ([]model.User, error)
	CreateNewUser(context.Context, model.User) (int64, error)
	UpdateUser(context.Context, model.User) error
	DeleteUser(context.Context, int64) error
}
