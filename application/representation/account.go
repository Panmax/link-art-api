package representation

type AccountProfileRepresentation struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Avatar    *string `json:"avatar"`
	Phone     string  `json:"phone"`
	Introduce string  `json:"introduce"`
	Gender    uint8   `json:"gender"`
	Birth     *int64  `json:"birth"`
	Follow    int     `json:"follow_count"`
	Fans      int     `json:"fans_count"`
	Points    int     `json:"points"`
	Artist    bool    `json:"points"`
}

type ArtistRepresentation struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Avatar *string `json:"avatar"`
}
