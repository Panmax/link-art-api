package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"link-art-api/domain/model"
)

func GetLoginToken(phone, password string) (string, error) {
	return "token", nil
}

func AccountRegister(phone, password string) (string, error) {
	_, err := model.FindAccountByPhone(phone)

	// https://www.flysnow.org/2019/09/06/go1.13-error-wrapping.html
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("手机号码已注册，可直接登录")
	}
	err = model.CreateAccount(phone, password)
	if err != nil {
		return "", err
	}
	return "token", nil
}
