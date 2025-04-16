package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User представляет пользователя в базе данных
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

// BeforeCreate хеширует пароль перед созданием пользователя
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Хешируем пароль перед сохранением в базу
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
