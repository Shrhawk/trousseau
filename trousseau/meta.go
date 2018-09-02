package trousseau

import "time"

type Meta struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewMeta(createdAt, updatedAt time.Time) *Meta {
	return &Meta{createdAt, updatedAt}
}
