package usecase

import (
	"context"

	"rapdev-graphql/internal/domain"
)

type EducationUsecase struct {
	repo domain.EducationRepository
}

func NewEducationUsecase(repo domain.EducationRepository) *EducationUsecase {
	return &EducationUsecase{repo: repo}
}

func (u *EducationUsecase) GetAll(ctx context.Context) ([]domain.Education, error) {
	return u.repo.FindAll(ctx)
}
