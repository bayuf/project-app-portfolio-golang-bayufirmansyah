package model

type Project struct {
	Model       Model
	UserId      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IconURL     string `json:"icon_url"`
	LinkURL     string `json:"link_url"`
}
