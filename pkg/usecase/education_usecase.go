package usecase

import (
	"context"

	"rapdev-graphql/pkg/domain"
)

type EducationUsecase struct {
	repo domain.EducationRepository
}

func NewEducationUsecase(repo domain.EducationRepository) *EducationUsecase {
	return &EducationUsecase{repo: repo}
}

func (u *EducationUsecase) GetAll(ctx context.Context, params domain.PaginationParams) (domain.PaginatedResult[domain.Education], error) {
	return u.repo.FindAll(ctx, params)
}
