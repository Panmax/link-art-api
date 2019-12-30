package representation

import "link-art-api/domain/model"

type MessageRepresentation struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Read      bool   `json:"read"`
	CreatedAt int64  `json:"created_at"`
}

func NewMessageRepresentation(messageFlow *model.MessageFlow) *MessageRepresentation {
	message := &MessageRepresentation{
		ID:        messageFlow.ID,
		Title:     messageFlow.Title,
		Content:   messageFlow.Content,
		Read:      messageFlow.Read,
		CreatedAt: messageFlow.CreatedAt.Unix(),
	}
	return message
}
