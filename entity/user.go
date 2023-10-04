package entity

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultCost = 10
)

type User struct {
	ID        int64
	Username  string
	Password  string
	FirstName string
	LastName  string
}

func (u *User) SetPassword(password string) {
	u.Password = encryptPassword(password)
}

func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func encryptPassword(password string) string {
	encrypted, _ := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	return string(encrypted)
}
