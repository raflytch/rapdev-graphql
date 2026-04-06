package graphqldelivery

import (
	"rapdev-graphql/internal/usecase"
)

type Resolver struct {
	ArticleUsecase    *usecase.ArticleUsecase
	EducationUsecase  *usecase.EducationUsecase
	ExperienceUsecase *usecase.ExperienceUsecase
	ProjectUsecase    *usecase.ProjectUsecase
	GalleryUsecase    *usecase.GalleryUsecase
	SocialLinkUsecase *usecase.SocialLinkUsecase
}

func NewResolver(
	articleUC *usecase.ArticleUsecase,
	educationUC *usecase.EducationUsecase,
	experienceUC *usecase.ExperienceUsecase,
	projectUC *usecase.ProjectUsecase,
	galleryUC *usecase.GalleryUsecase,
	socialLinkUC *usecase.SocialLinkUsecase,
) *Resolver {
	return &Resolver{
		ArticleUsecase:    articleUC,
		EducationUsecase:  educationUC,
		ExperienceUsecase: experienceUC,
		ProjectUsecase:    projectUC,
		GalleryUsecase:    galleryUC,
		SocialLinkUsecase: socialLinkUC,
	}
}
