package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"link-art-api/application/command"
	"link-art-api/domain/model"
	"time"
)

func AccountRegister(phone, password string) (*model.Account, error) {
	_, err := model.FindAccountByPhone(phone)

	// https://www.flysnow.org/2019/09/06/go1.13-error-wrapping.html
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("手机号码已注册，可直接登录")
	}
	account, err := model.CreateAccount(phone, password)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func UpdateProfile(id uint, updateCommand *command.UpdateProfileCommand) (bool, error) {
	account, err := model.FindAccount(id)
	if err != nil {
		return false, err
	}
	account.Name = updateCommand.Name
	account.Gender = updateCommand.Gender
	account.Introduce = updateCommand.Introduce
	if updateCommand.Birth != nil {
		birth := time.Unix(*updateCommand.Birth, 0)
		account.Birth = &birth
	}
	err = model.SaveOne(account)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ListAccountFollow(id uint) []map[string]string {
	return make([]map[string]string, 0)
}

func ListAccountFans(id uint) []map[string]string {
	return make([]map[string]string, 0)
}
