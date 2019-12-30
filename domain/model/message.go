package model

type Message struct {
	Model

	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}

type MessageFlow struct {
	Model

	AccountId uint `gorm:"index:idx_account;not null"`

	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	Read    bool   `gorm:"not null"`
}

func (m *MessageFlow) MarkRead() {
	m.Read = true
}
