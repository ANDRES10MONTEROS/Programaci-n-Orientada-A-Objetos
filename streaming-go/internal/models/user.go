package models

import (
	"errors"
	"time"
)

type User struct {
	id        string
	name      string
	email     string
	subscribed bool
	createdAt time.Time
}

func NewUser(id, name, email string) (*User, error) {
	if id == "" || name == "" || email == "" {
		return nil, errors.New("id, name and email are required")
	}
	return &User{
		id:        id,
		name:      name,
		email:     email,
		subscribed: false,
		createdAt: time.Now(),
	}, nil
}

func (u *User) GetID() string { return u.id }
func (u *User) GetName() string { return u.name }
func (u *User) SetName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	u.name = name
	return nil
}
func (u *User) GetEmail() string { return u.email }
func (u *User) SetEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	u.email = email
	return nil
}
func (u *User) IsSubscribed() bool { return u.subscribed }
func (u *User) Subscribe() { u.subscribed = true }
func (u *User) Unsubscribe() { u.subscribed = false }
func (u *User) GetCreatedAt() time.Time { return u.createdAt }
