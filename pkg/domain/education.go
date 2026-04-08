package domain

import (
	"context"
	"time"
)

type Education struct {
	ID           string
	Institution  string
	Degree       string
	Logo         *string
	StartDate    time.Time
	EndDate      *time.Time
	GPA          *string
	Achievements []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type EducationRepository interface {
	FindAll(ctx context.Context, params PaginationParams) (PaginatedResult[Education], error)
}
