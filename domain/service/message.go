package service

import (
	"link-art-api/application/representation"
	"link-art-api/domain/model"
	"link-art-api/domain/repository"
)

func ListMessage(accountId uint) ([]*representation.MessageRepresentation, error) {
	flows, err := repository.FindAllMessageFlow("account_id = ?", accountId)
	if err != nil {
		return nil, err
	}

	results := make([]*representation.MessageRepresentation, 0)
	for _, flow := range flows {
		message, err := GetMessage(flow.ID)
		if err != nil {
			return nil, err
		}

		results = append(results, message)
	}

	return results, nil
}

func GetMessage(id uint) (*representation.MessageRepresentation, error) {
	flow, err := repository.FindMessageFlow(id)
	if err != nil {
		return nil, err
	}
	return representation.NewMessageRepresentation(flow), nil
}

func DeleteMessage(id uint) error {
	return repository.DeleteMessageFlow(id)
}

func ReadMessage(id uint) error {
	flow, err := repository.FindMessageFlow(id)
	if err != nil {
		return err
	}
	flow.MarkRead()

	return model.SaveOne(flow)
}

func CheckNewMessage(accountId uint) (bool, error) {
	flows, err := repository.FindAllMessageFlow("account_id = ?", accountId)
	if err != nil {
		return false, err
	}
	if len(flows) > 0 {
		return !flows[0].Read, nil
	}

	return false, nil
}
