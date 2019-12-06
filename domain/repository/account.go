package repository

import "link-art-api/domain/model"

func FindAccountByPhone(phone string) (*model.Account, error) {
	account := &model.Account{}
	err := model.DB.Where("phone = ?", phone).First(&account).Error
	return account, err
}

func FindAccount(id uint) (*model.Account, error) {
	account := &model.Account{}
	err := model.DB.Unscoped().First(account, id).Error
	return account, err
}

func FindApprovalByUser(accountId uint) (*model.Approval, error) {
	approval := &model.Approval{}
	err := model.DB.Where("account_id = ?", accountId).Find(&approval).Error
	return approval, err
}
