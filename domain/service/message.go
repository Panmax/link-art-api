package service

import (
	"link-art-api/application/representation"
	"time"
)

func ListMessage(accountId uint) ([]*representation.MessageRepresentation, error) {
	results := make([]*representation.MessageRepresentation, 0)

	m := representation.MessageRepresentation{
		ID:        1,
		Title:     "标题标题",
		Content:   "这是正文",
		Read:      false,
		CreatedAt: time.Now().Unix(),
	}
	results = append(results, &m) // TODO

	return results, nil
}

func GetMessage(id uint) (*representation.MessageRepresentation, error) {
	// TODO
	m := representation.MessageRepresentation{
		ID:        1,
		Title:     "标题标题",
		Content:   "这是正文",
		Read:      false,
		CreatedAt: time.Now().Unix(),
	}
	return &m, nil
}

func DeleteMessage(id uint) (bool, error) {
	return true, nil
}

func ReadMessage(id uint) (bool, error) {
	return true, nil
}

func CheckNewMessage(accountId uint) (bool, error) {
	return false, nil
}
