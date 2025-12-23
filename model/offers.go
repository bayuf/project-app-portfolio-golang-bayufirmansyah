package model

type Offers struct {
	Model       Model
	UserId      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IconURL     string `json:"icon_url"`
}
