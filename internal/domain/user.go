package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserID uuid.UUID

// Helpers
func NewUserId() UserID {
	id, _ := uuid.NewV7()
	return UserID(id)
}

func ParseUserId(s string) (UserID, error) {
	id, err := uuid.Parse(s)
	return UserID(id), err
}

func (id UserID) String() string {
	return uuid.UUID(id).String()
}

// Entity
type User struct {
	ID        UserID
	Name      string
	Email     string
	DeletedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Constructor
func NewUser(id UserID, name, email string, now time.Time) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Getters
func (u *User) GetID() UserID            { return u.ID }
func (u *User) GetName() string          { return u.Name }
func (u *User) GetEmail() string         { return u.Email }
func (u *User) GetDeletedAt() *time.Time { return u.DeletedAt }
func (u *User) GetCreatedAt() time.Time  { return u.CreatedAt }
func (u *User) GetUpdatedAt() time.Time  { return u.UpdatedAt }

// Mutations
func (u *User) Update(name, email string, now time.Time) {
	u.Name = name
	u.Email = email
	u.UpdatedAt = now
}

func (u *User) SoftDelete(now time.Time) {
	u.DeletedAt = &now
	u.UpdatedAt = now
}
