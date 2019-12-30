package representation

type MessageRepresentation struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Read      bool   `json:"read"`
	CreatedAt int64  `json:"created_at"`
}
