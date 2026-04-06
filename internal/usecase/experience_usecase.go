package usecase

import (
	"context"

	"rapdev-graphql/internal/domain"
)

type ExperienceUsecase struct {
	repo domain.ExperienceRepository
}

func NewExperienceUsecase(repo domain.ExperienceRepository) *ExperienceUsecase {
	return &ExperienceUsecase{repo: repo}
}

func (u *ExperienceUsecase) GetAll(ctx context.Context) ([]domain.Experience, error) {
	return u.repo.FindAll(ctx)
}
