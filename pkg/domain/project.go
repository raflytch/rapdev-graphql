package domain

import (
	"context"
	"time"
)

type Project struct {
	ID          string
	Title       string
	Subtitle    *string
	Description string
	Image       *string
	Tags        []string
	DemoURL     *string
	GithubURL   *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProjectRepository interface {
	FindAll(ctx context.Context) ([]Project, error)
}
