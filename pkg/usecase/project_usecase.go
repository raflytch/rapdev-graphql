package usecase

import (
	"context"

	"rapdev-graphql/pkg/domain"
)

type ProjectUsecase struct {
	repo domain.ProjectRepository
}

func NewProjectUsecase(repo domain.ProjectRepository) *ProjectUsecase {
	return &ProjectUsecase{repo: repo}
}

func (u *ProjectUsecase) GetAll(ctx context.Context, params domain.PaginationParams) (domain.PaginatedResult[domain.Project], error) {
	return u.repo.FindAll(ctx, params)
}
