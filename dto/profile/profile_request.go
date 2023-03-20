package profilesdto

type CreateProfileRequest struct {
	Image    string `json:"image"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	Phone    string `json:"phone"`
	UserID   int    `json:"user_id"`
}

type UpdateProfileRequest struct {
	Image    string `json:"image"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	Phone    string `json:"phone"`
}
