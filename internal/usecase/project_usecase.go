package usecase

import (
	"context"

	"rapdev-graphql/internal/domain"
)

type ProjectUsecase struct {
	repo domain.ProjectRepository
}

func NewProjectUsecase(repo domain.ProjectRepository) *ProjectUsecase {
	return &ProjectUsecase{repo: repo}
}

func (u *ProjectUsecase) GetAll(ctx context.Context) ([]domain.Project, error) {
	return u.repo.FindAll(ctx)
}
