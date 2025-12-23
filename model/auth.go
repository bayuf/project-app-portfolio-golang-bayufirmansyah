package model

import "time"

type User struct {
	Model     Model
	Email     string
	Password  string
	LastLogin time.Time
}
