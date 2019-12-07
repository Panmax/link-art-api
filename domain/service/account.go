package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"link-art-api/application/command"
	"link-art-api/domain/model"
	"link-art-api/domain/repository"
	"time"
)

func AccountRegister(phone, password string) (*model.Account, error) {
	_, err := repository.FindAccountByPhone(phone)

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
	account, err := repository.FindAccount(id)
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
	_, err := repository.FindApprovalByUser(accountId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("认证审核中，请勿重复提交")
	}

	approval := model.NewApproval(accountId, submitCommand.Type, submitCommand.CompanyName, submitCommand.Photo)
	return model.CreateOne(approval)
}

func ApprovalPass(id uint) error {
	approval, err := repository.FundApproval(id)
	if err != nil {
		return nil
	}

	account, err := repository.FindAccount(approval.AccountId)
	if err != nil {
		return nil
	}

	approval.Pass()
	account.BeArtist()

	tx := model.DB.Begin()
	tx.Save(approval)
	tx.Save(account)
	return tx.Commit().Error
}

func ApprovalReject(id uint) error {
	approval, err := repository.FundApproval(id)
	if err != nil {
		return nil
	}

	account, err := repository.FindAccount(approval.AccountId)
	if err != nil {
		return nil
	}

	approval.Reject()
	account.CancelArtist()

	tx := model.DB.Begin()
	tx.Save(approval)
	tx.Save(account)
	return tx.Commit().Error
}
