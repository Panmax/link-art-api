package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Account struct {
	Model

	Name         string     `gorm:"size:16;not null"`
	Phone        string     `gorm:"size:16;unique_index;not null"`
	Gender       uint8      `gorm:"not null"`
	Introduce    string     `gorm:"size:512;not null"`
	PasswordHash string     `gorm:"column:password;not null"`
	Birth        *time.Time `gorm:"type:date;default null"`
}

func CreateAccount(phone, password string) (*Account, error) {
	account := &Account{
		Phone:        phone,
		PasswordHash: passwordHash(password),
	}
	if err := db.Create(account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

func passwordHash(password string) string {
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(passwordHash)
}

func (a *Account) CheckPassword(password string) bool {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(a.PasswordHash)
	if bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword) == nil {
		return true
	}

	return false
}

func FindAccountByPhone(phone string) (account Account, err error) {
	err = db.Where("phone = ?", phone).First(&account).Error
	return
}

func FindAccount(id uint) (*Account, error) {
	account := &Account{}
	err := db.Unscoped().First(account, id).Error
	return account, err
}
