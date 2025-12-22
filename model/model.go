package model

import "time"

type Model struct {
	ID         int        `json:"id"`
	Created_At time.Time  `json:"created_at"`
	Updated_At time.Time  `json:"updates_at"`
	Deleted_At *time.Time `json:"deleted_at,omitempty"`
}
