package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Account struct {
	Model

	Name         string  `gorm:"size:16;not null"`
	Avatar       *string `gorm:"size:512"`
	Phone        string  `gorm:"size:16;unique_index;not null"`
	Gender       uint8   `gorm:"not null"`
	Introduce    string  `gorm:"size:512;not null"`
	PasswordHash string  `gorm:"column:password;not null"`
	Artist       bool    `gorm:"artist;not null"`

	//https://reading.developerlearning.cn/discuss/2019-06-19-gorm-mysql-timestamp/
	Birth *time.Time `gorm:"type:date"`
}

func NewAccount(phone, password string) *Account {
	account := &Account{
		Phone:        phone,
		PasswordHash: passwordHash(password),
	}
	return account
}

type Approval struct {
	Model

	AccountId   uint    `gorm:"not null;unique_index"`
	Type        uint8   `gorm:"not null"`
	CompanyName *string `gorm:"size:64"`
	Photo       string  `gorm:"size:512;not null"`
	Status      uint8   `gorm:"not null"`
}

func NewApproval(accountId uint, _type uint8, companyName *string, photo string) *Approval {
	return &Approval{
		AccountId:   accountId,
		Type:        _type,
		CompanyName: companyName,
		Photo:       photo,
		Status:      0,
	}
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

func (a *Account) UpdateAvatar(url *string) {
	a.Avatar = url
}

func FindAccountByPhone(phone string) (*Account, error) {
	account := &Account{}
	err := db.Where("phone = ?", phone).First(&account).Error
	return account, err
}

func FindAccount(id uint) (*Account, error) {
	account := &Account{}
	err := db.Unscoped().First(account, id).Error
	return account, err
}

func FindApprovalByUser(accountId uint) (*Approval, error) {
	approval := &Approval{}
	err := db.Where("account_id = ?", accountId).Find(&approval).Error
	return approval, err
}
