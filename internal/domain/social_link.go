package domain

import (
	"context"
	"time"
)

type SocialLink struct {
	ID        string
	Title     string
	URL       string
	Order     int
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SocialLinkRepository interface {
	FindAll(ctx context.Context) ([]SocialLink, error)
}
