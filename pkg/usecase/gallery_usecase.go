package usecase

import (
	"context"

	"rapdev-graphql/pkg/domain"
)

type GalleryUsecase struct {
	repo domain.GalleryRepository
}

func NewGalleryUsecase(repo domain.GalleryRepository) *GalleryUsecase {
	return &GalleryUsecase{repo: repo}
}

func (u *GalleryUsecase) GetAll(ctx context.Context, params domain.PaginationParams) (domain.PaginatedResult[domain.Gallery], error) {
	return u.repo.FindAll(ctx, params)
}
