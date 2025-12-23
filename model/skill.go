package model

type Skill struct {
	Model  Model
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Level  string `json:"level"`
	ImgURL string `json:"img_url"`
}
