package models

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/valueobjects"
)

type User struct {
	ID           string `gorm:"type:VARCHAR(255);primarykey;not null"`
	Email        string `gorm:"type:VARCHAR(255);unique;not null"`
	PasswordHash string `gorm:"type:VARCHAR(255);column:password_hash"`
	Role         string `gorm:"type:VARCHAR(5);not null"`
}

func (u *User) MapToModel(entity entities.UserEntity) {
	u.Email = entity.Email
	u.ID = entity.ID.GetValue()
	u.PasswordHash = entity.PasswordHash.GetValue()
	u.Role = entity.Role.GetValue()
}

func (u *User) MapToEntity() entities.UserEntity {
	var (
		id           valueobjects.EntityID
		passwordHash valueobjects.PasswordHash
		role         valueobjects.UserRole
	)

	id.Assign(u.ID)
	passwordHash.Assign(u.PasswordHash)
	role.Assign(u.Role)

	return entities.UserEntity{
		ID:           id,
		Email:        u.Email,
		PasswordHash: passwordHash,
		Role:         role,
	}
}
