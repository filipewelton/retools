package entities

import "backend/internal/domain/valueobjects"

type UserEntity struct {
	ID           valueobjects.EntityID
	Email        string
	PasswordHash valueobjects.PasswordHash
	Role         valueobjects.UserRole
}

type MappedUserEntity struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserEntityCreation struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (u *UserEntity) Create(params UserEntityCreation) {
	var (
		id           valueobjects.EntityID
		passwordHash valueobjects.PasswordHash
		role         valueobjects.UserRole
	)

	id.Generate()
	passwordHash.Encrypt(params.Password)
	role.Assign(params.Role)

	u.ID = id
	u.Email = params.Email
	u.PasswordHash = passwordHash
	u.Role = role
}

func (u *UserEntity) Map() MappedUserEntity {
	return MappedUserEntity{
		ID:    u.ID.GetValue(),
		Email: u.Email,
		Role:  u.Role.GetValue(),
	}
}
