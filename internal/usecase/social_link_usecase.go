package usecase

import (
	"context"

	"rapdev-graphql/internal/domain"
)

type SocialLinkUsecase struct {
	repo domain.SocialLinkRepository
}

func NewSocialLinkUsecase(repo domain.SocialLinkRepository) *SocialLinkUsecase {
	return &SocialLinkUsecase{repo: repo}
}

func (u *SocialLinkUsecase) GetAll(ctx context.Context) ([]domain.SocialLink, error) {
	return u.repo.FindAll(ctx)
}
