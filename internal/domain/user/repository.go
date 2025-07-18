package user

import (
	"context"
)

// --- UserExistsRepository ---
type UserExistsRepoInput struct {
	UserExistsInput
}

type UserExistsRepository interface {
	UserExists(ctx context.Context, input UserExistsRepoInput) (bool, error)
}

// --- CreateUserRepository ---

type CreateUserRepoInput struct {
	User *User
}

type CreateUserRepository interface {
	UserExistsRepository
	CreateUser(ctx context.Context, input CreateUserRepoInput) error
}

// --- UserRepository ---
type UserRepository interface {
	UserExistsRepository
	CreateUserRepository
}
