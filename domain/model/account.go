package model

import "golang.org/x/crypto/bcrypt"

type Account struct {
	Model

	Phone        string `gorm:"size:32;unique_index;not null"`
	Bio          string `gorm:"size:1024;not null"`
	PasswordHash string `gorm:"column:password;not null"`
}

func CreateAccount(phone, password string) error {
	account := Account{
		Phone:        phone,
		PasswordHash: passwordHash(password),
	}
	if err := db.Create(&account).Error; err != nil {
		return err
	}

	return nil
}

func FindAccountByPhone(phone string) (account Account, err error) {
	err = db.Where("phone = ?", phone).First(&account).Error
	return
}

func passwordHash(password string) string {
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(passwordHash)
}
