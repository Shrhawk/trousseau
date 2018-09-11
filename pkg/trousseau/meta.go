package trousseau

import "time"

type Meta struct {
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Description string            `json:"description"`
	Labels      map[string]string `json:"labels"`
}

type Describable interface {
	Meta() Meta
	SetMeta(m Meta) error
}

func NewMetaWithDefaults() *Meta {
	return &Meta{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: "",
		Labels:      map[string]string{},
	}
}
