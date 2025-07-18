package user

import (
	"encoding/json"

	"github.com/marcelofabianov/chronos/internal/platform/msg"
	"github.com/marcelofabianov/chronos/internal/platform/port/hasher"
	"github.com/marcelofabianov/chronos/internal/platform/types"
)

type NewUserInput struct {
	Name        string
	Email       string
	Phone       string
	Password    string
	Role        Role
	Preferences json.RawMessage `json:"preferences,omitempty"`
}

type UserExistsInput struct {
	Email *types.Email
	Phone *types.Phone
}

type FromUserInput struct {
	ID             types.UUID
	Name           string
	Email          types.Email
	Phone          types.Phone
	HashedPassword types.HashedPassword
	Role           Role
	Status         UserLoginStatus
	Preferences    json.RawMessage
	CreatedAt      types.CreatedAt
	UpdatedAt      types.UpdatedAt
	ArchivedAt     types.ArchivedAt
	DeletedAt      types.DeletedAt
	Version        types.Version
}

type User struct {
	ID             types.UUID           `db:"id" json:"id"`
	Name           string               `db:"name" json:"name"`
	Email          types.Email          `db:"email" json:"email"`
	Phone          types.Phone          `db:"phone" json:"phone"`
	HashedPassword types.HashedPassword `db:"hashed_password" json:"-"`
	Role           Role                 `db:"role" json:"role"`
	LoginStatus    UserLoginStatus      `db:"login_status" json:"login_status"`
	Preferences    json.RawMessage      `db:"preferences" json:"preferences,omitempty"`
	CreatedAt      types.CreatedAt      `db:"created_at" json:"created_at"`
	UpdatedAt      types.UpdatedAt      `db:"updated_at" json:"updated_at"`
	ArchivedAt     types.ArchivedAt     `db:"archived_at" json:"archived_at,omitempty"`
	DeletedAt      types.DeletedAt      `db:"deleted_at" json:"-"`
	Version        types.Version        `db:"version" json:"-"`
}

func NewUser(input NewUserInput, h hasher.Hasher) (*User, error) {
	hashedPassword, err := PasswordHash(input.Password, h)
	if err != nil {
		return nil, err
	}

	phone, err := types.NewPhone(input.Phone)
	if err != nil {
		return nil, err
	}

	email, err := types.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	var preferencesValue json.RawMessage
	if input.Preferences != nil {
		if !json.Valid(input.Preferences) {
			return nil, msg.NewValidationError(nil, map[string]any{"field": "Preferences"}, ErrUserPreferencesNotValidJSON)
		}
		if string(input.Preferences) != "null" && len(input.Preferences) > 0 {
			preferencesValue = input.Preferences
		}
	}

	return &User{
		ID:             types.MustNewUUID(),
		Name:           input.Name,
		Email:          email,
		Phone:          phone,
		HashedPassword: hashedPassword,
		Role:           input.Role,
		LoginStatus:    UserLoginStatusPending,
		Preferences:    preferencesValue,
		CreatedAt:      types.NewCreatedAt(),
		UpdatedAt:      types.NewUpdatedAt(),
		ArchivedAt:     types.NewNilArchivedAt(),
		DeletedAt:      types.NewNilDeletedAt(),
		Version:        types.NewVersion(),
	}, nil
}

func FromUser(input FromUserInput) *User {
	return &User{
		ID:             input.ID,
		Name:           input.Name,
		Email:          input.Email,
		Phone:          input.Phone,
		HashedPassword: input.HashedPassword,
		Role:           input.Role,
		LoginStatus:    input.Status,
		Preferences:    input.Preferences,
		CreatedAt:      input.CreatedAt,
		UpdatedAt:      input.UpdatedAt,
		ArchivedAt:     input.ArchivedAt,
		DeletedAt:      input.DeletedAt,
		Version:        input.Version,
	}
}

func PasswordHash(pass string, h hasher.Hasher) (types.HashedPassword, error) {
	password, err := types.NewPassword(pass)
	if err != nil {
		return "", err
	}

	hashedPassword, err := h.Hash(password.String())
	if err != nil {
		return "", err
	}

	ha := types.HashedPassword(hashedPassword)

	return ha, nil
}

// @TODO: add func para validar dados do input antes de retornar a struct

// @TODO: add func para update com um input

// @TODO: add func para verificar se um usuario é igual ao informado pelo input

// @TODO: add func(s) para manipular dados de email, password, status, role, phone...
