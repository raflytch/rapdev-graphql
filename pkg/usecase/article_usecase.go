package usecase

import (
	"context"

	"rapdev-graphql/pkg/domain"
)

type ArticleUsecase struct {
	repo domain.ArticleRepository
}

func NewArticleUsecase(repo domain.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

func (u *ArticleUsecase) GetAll(ctx context.Context) ([]domain.Article, error) {
	return u.repo.FindAll(ctx)
}
