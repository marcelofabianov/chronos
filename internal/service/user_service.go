package service

import (
	"context"

	"github.com/marcelofabianov/chronos/internal/domain/user"
	"github.com/marcelofabianov/chronos/internal/platform/port/hasher"
)

type UserService struct {
	repo   user.UserRepository
	hasher hasher.Hasher
}

func NewUserService(repo user.UserRepository, hasher hasher.Hasher) *UserService {
	return &UserService{repo: repo, hasher: hasher}
}

func (u *UserService) Create(ctx context.Context, input user.NewUserInput) (*user.User, error) {
	newUser, err := user.NewUser(input, u.hasher)
	if err != nil {
		return nil, err
	}

	// @TODO: validar se o usuario ja existe para impedir a criacao

	rInput := user.CreateUserRepoInput{User: newUser}

	err = u.repo.CreateUser(ctx, rInput)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *UserService) Exists(ctx context.Context, user user.User) (bool, error) {
	//...

	// @TODO: utilizar o user.Equals(input user.UserExistsInput)

	return true, nil
}

// @TODO: adicionar demais func...
