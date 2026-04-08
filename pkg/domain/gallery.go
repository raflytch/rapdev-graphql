package domain

import (
	"context"
	"time"
)

type Gallery struct {
	ID          string
	Image       string
	ImageFileID string
	Caption     *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GalleryRepository interface {
	FindAll(ctx context.Context, params PaginationParams) (PaginatedResult[Gallery], error)
}
