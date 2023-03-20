package profilesdto

type ProfileResponse struct {
	ID       int    `json:"id"`
	Image    string `json:"image"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	Phone    string `json:"phone"`
	UserID   int    `json:"user_id"`
}
