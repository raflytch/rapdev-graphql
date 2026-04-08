package usecase

import (
	"context"

	"rapdev-graphql/pkg/domain"
)

type SocialLinkUsecase struct {
	repo domain.SocialLinkRepository
}

func NewSocialLinkUsecase(repo domain.SocialLinkRepository) *SocialLinkUsecase {
	return &SocialLinkUsecase{repo: repo}
}

func (u *SocialLinkUsecase) GetAll(ctx context.Context, params domain.PaginationParams) (domain.PaginatedResult[domain.SocialLink], error) {
	return u.repo.FindAll(ctx, params)
}
