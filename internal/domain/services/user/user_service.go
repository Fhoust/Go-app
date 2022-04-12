package user

import (
	"context"
	repositories "github.com/zerepl/go-app/internal/data"
	api "github.com/zerepl/go-app/internal/domain/services"
	"github.com/zerepl/go-app/internal/model"
	"log"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userService repositories.UserRepository) api.UserService {
	return UserService{userRepository: userService}
}

func (u UserService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	user, err := u.userRepository.GetUser(ctx, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (u UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	users, err := u.userRepository.GetAllUsers(ctx)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func (u UserService) CreateNewUser(ctx context.Context, user model.User) (int64, error) {
	id, err := u.userRepository.CreateNewUser(ctx, user)

	if err != nil {
		log.Println(err)
		return -1, err
	}

	return id, nil
}

func (u UserService) UpdateUser(ctx context.Context, user model.User) error {
	err := u.userRepository.UpdateUser(ctx, user)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u UserService) DeleteUser(ctx context.Context, id int64) error {
	err := u.userRepository.DeleteUser(ctx, id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
