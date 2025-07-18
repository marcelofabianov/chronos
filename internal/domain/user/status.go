package user

type UserLoginStatus string

const (
	UserLoginStatusActive   UserLoginStatus = "active"
	UserLoginStatusInactive UserLoginStatus = "inactive"
	UserLoginStatusBlocked  UserLoginStatus = "blocked"
	UserLoginStatusPending  UserLoginStatus = "pending"
)

// @TODO: add func para tratamento do dados
