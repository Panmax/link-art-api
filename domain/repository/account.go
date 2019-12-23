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

func FindAllAccount(args ...interface{}) ([]model.Account, error) {
	var accounts []model.Account

	cond := model.DB
	if len(args) >= 2 {
		cond = cond.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		cond = cond.Where(args[0])
	}

	err := cond.Find(&accounts).Error
	return accounts, err
}

func FindApprovalByAccount(accountId uint) (*model.Approval, error) {
	approval := &model.Approval{}
	err := model.DB.Where("account_id = ?", accountId).Find(&approval).Error
	return approval, err
}

func FindApproval(id uint) (*model.Approval, error) {
	approval := &model.Approval{}
	err := model.DB.Unscoped().First(approval, id).Error
	return approval, err
}

func FindAllFollowFlow(args ...interface{}) ([]model.FollowFlow, error) {
	var flows []model.FollowFlow
	cond := model.DB
	if len(args) >= 2 {
		cond = cond.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		cond = cond.Where(args[0])
	}
	err := cond.Find(&flows).Error

	return flows, err
}

func DeleteFollowFlow(accountId, followerId uint) error {
	err := model.DB.Unscoped().Where(
		"account_id = ? AND follower_id = ?", accountId, followerId).Delete(&model.FollowFlow{}).Error
	return err
}
