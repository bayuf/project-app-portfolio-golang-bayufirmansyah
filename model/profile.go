package model

type Profile struct {
	Model      Model
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Headline   string `json:"headline"`
	About      string `json:"about"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Experience int    `json:"experience"`
}
