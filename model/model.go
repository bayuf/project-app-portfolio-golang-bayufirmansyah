package model

import "time"

type Model struct {
	ID         int
	Created_At time.Time
	Updated_At time.Time
	Deleted_AT *time.Time
}
