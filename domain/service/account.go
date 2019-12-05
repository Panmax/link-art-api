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

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("手机号码已注册，可直接登录")
	}

	account := model.NewAccount(phone, password)
	if err := model.CreateOne(account); err != nil {
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

func SubmitApproval(accountId uint, submitCommand *command.SubmitApprovalCommand) error {
	_, err := model.FindApprovalByUser(accountId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("认证审核中，请勿重复提交")
	}

	approval := model.NewApproval(accountId, submitCommand.Type, submitCommand.CompanyName, submitCommand.Photo)
	return model.CreateOne(approval)
}
