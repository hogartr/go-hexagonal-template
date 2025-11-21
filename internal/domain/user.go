package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserID uuid.UUID
type Date time.Time

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

func ParseDate(s string) (Date, error) {
	t, err := time.Parse(time.RFC3339, s)
	return Date(t), err
}

func (d Date) String() string {
	return time.Time(d).Format(time.RFC3339)
}

func (d Date) ToTime() time.Time {
	return time.Time(d)
}

// Entity
type User struct {
	ID        UserID
	Name      string
	Email     string
	DeletedAt *Date
	CreatedAt Date
	UpdatedAt Date
}

// Constructor
func NewUser(id UserID, name, email string, now Date) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Getters
func (u *User) GetID() UserID       { return u.ID }
func (u *User) GetName() string     { return u.Name }
func (u *User) GetEmail() string    { return u.Email }
func (u *User) GetDeletedAt() *Date { return u.DeletedAt }
func (u *User) GetCreatedAt() Date  { return u.CreatedAt }
func (u *User) GetUpdatedAt() Date  { return u.UpdatedAt }

// Mutations
func (u *User) Update(name, email string, now Date) {
	u.Name = name
	u.Email = email
	u.UpdatedAt = now
}

func (u *User) SoftDelete(now Date) {
	u.DeletedAt = &now
	u.UpdatedAt = now
}
