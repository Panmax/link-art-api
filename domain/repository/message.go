package repository

import "link-art-api/domain/model"

func FindMessageFlow(id uint) (*model.MessageFlow, error) {
	messageFlow := &model.MessageFlow{}
	err := model.DB.Unscoped().First(messageFlow, id).Error
	return messageFlow, err
}

func DeleteMessageFlow(id uint) error {
	err := model.DB.Unscoped().Where(
		"id = ?", id).Delete(&model.MessageFlow{}).Error
	return err
}

func FindAllMessageFlow(args ...interface{}) ([]model.MessageFlow, error) {
	var flows []model.MessageFlow
	cond := model.DB
	if len(args) >= 2 {
		cond = cond.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		cond = cond.Where(args[0])
	}
	err := cond.Order("id desc").Find(&flows).Error

	return flows, err
}
