package domain

import (
	"context"
	"time"
)

type Experience struct {
	ID          string
	Company     string
	Position    string
	Type        string
	Logo        *string
	StartDate   time.Time
	EndDate     *time.Time
	Tags        []string
	Description []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ExperienceRepository interface {
	FindAll(ctx context.Context) ([]Experience, error)
}
