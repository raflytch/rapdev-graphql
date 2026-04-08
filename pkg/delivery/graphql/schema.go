package graphqldelivery

import (
	"context"

	"rapdev-graphql/pkg/domain"

	"github.com/graphql-go/graphql"
)

func NewSchema(r *Resolver) (graphql.Schema, error) {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"articles":    buildArticlesField(r),
			"educations":  buildEducationsField(r),
			"experiences": buildExperiencesField(r),
			"projects":    buildProjectsField(r),
			"galleries":   buildGalleriesField(r),
			"socialLinks": buildSocialLinksField(r),
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}

func extractPaginationParams(p graphql.ResolveParams) domain.PaginationParams {
	page := 1
	limit := 10
	if v, ok := p.Args["page"].(int); ok {
		page = v
	}
	if v, ok := p.Args["limit"].(int); ok {
		limit = v
	}
	return domain.NewPaginationParams(page, limit)
}

func toPaginatedResponse(data []map[string]interface{}, totalCount, page, limit, totalPages int) map[string]interface{} {
	return map[string]interface{}{
		"data": data,
		"pageInfo": map[string]interface{}{
			"totalCount": totalCount,
			"page":       page,
			"limit":      limit,
			"totalPages": totalPages,
		},
	}
}

func buildArticlesField(r *Resolver) *graphql.Field {
	return &graphql.Field{
		Type: PaginatedArticleType,
		Args: PaginationArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			params := extractPaginationParams(p)
			result, err := r.ArticleUsecase.GetAll(resolveContext(p), params)
			if err != nil {
				return nil, err
			}
			return toPaginatedResponse(toArticleMaps(result.Data), result.TotalCount, result.Page, result.Limit, result.TotalPages), nil
		},
	}
}

func buildEducationsField(r *Resolver) *graphql.Field {
	return &graphql.Field{
		Type: PaginatedEducationType,
		Args: PaginationArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			params := extractPaginationParams(p)
			result, err := r.EducationUsecase.GetAll(resolveContext(p), params)
			if err != nil {
				return nil, err
			}
			return toPaginatedResponse(toEducationMaps(result.Data), result.TotalCount, result.Page, result.Limit, result.TotalPages), nil
		},
	}
}

func buildExperiencesField(r *Resolver) *graphql.Field {
	return &graphql.Field{
		Type: PaginatedExperienceType,
		Args: PaginationArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			params := extractPaginationParams(p)
			result, err := r.ExperienceUsecase.GetAll(resolveContext(p), params)
			if err != nil {
				return nil, err
			}
			return toPaginatedResponse(toExperienceMaps(result.Data), result.TotalCount, result.Page, result.Limit, result.TotalPages), nil
		},
	}
}

func buildProjectsField(r *Resolver) *graphql.Field {
	return &graphql.Field{
		Type: PaginatedProjectType,
		Args: PaginationArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			params := extractPaginationParams(p)
			result, err := r.ProjectUsecase.GetAll(resolveContext(p), params)
			if err != nil {
				return nil, err
			}
			return toPaginatedResponse(toProjectMaps(result.Data), result.TotalCount, result.Page, result.Limit, result.TotalPages), nil
		},
	}
}

func buildGalleriesField(r *Resolver) *graphql.Field {
	return &graphql.Field{
		Type: PaginatedGalleryType,
		Args: PaginationArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			params := extractPaginationParams(p)
			result, err := r.GalleryUsecase.GetAll(resolveContext(p), params)
			if err != nil {
				return nil, err
			}
			return toPaginatedResponse(toGalleryMaps(result.Data), result.TotalCount, result.Page, result.Limit, result.TotalPages), nil
		},
	}
}

func buildSocialLinksField(r *Resolver) *graphql.Field {
	return &graphql.Field{
		Type: PaginatedSocialLinkType,
		Args: PaginationArgs,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			params := extractPaginationParams(p)
			result, err := r.SocialLinkUsecase.GetAll(resolveContext(p), params)
			if err != nil {
				return nil, err
			}
			return toPaginatedResponse(toSocialLinkMaps(result.Data), result.TotalCount, result.Page, result.Limit, result.TotalPages), nil
		},
	}
}

func resolveContext(p graphql.ResolveParams) context.Context {
	if p.Context != nil {
		return p.Context
	}
	return context.Background()
}

func toArticleMap(a domain.Article) map[string]interface{} {
	m := map[string]interface{}{
		"id":        a.ID,
		"title":     a.Title,
		"content":   a.Content,
		"path":      a.Path,
		"viewCount": a.ViewCount,
		"likes":     a.Likes,
		"authorId":  a.AuthorID,
		"published": a.Published,
		"createdAt": a.CreatedAt,
		"updatedAt": a.UpdatedAt,
	}
	if a.Author != nil {
		m["author"] = toUserMap(*a.Author)
	}
	return m
}

func toArticleMaps(articles []domain.Article) []map[string]interface{} {
	result := make([]map[string]interface{}, len(articles))
	for i, a := range articles {
		result[i] = toArticleMap(a)
	}
	return result
}

func toUserMap(u domain.User) map[string]interface{} {
	return map[string]interface{}{
		"id":          u.ID,
		"name":        u.Name,
		"email":       u.Email,
		"role":        u.Role,
		"image":       u.Image,
		"imageFileId": u.ImageFileID,
		"createdAt":   u.CreatedAt,
		"updatedAt":   u.UpdatedAt,
	}
}

func toEducationMap(e domain.Education) map[string]interface{} {
	return map[string]interface{}{
		"id":           e.ID,
		"institution":  e.Institution,
		"degree":       e.Degree,
		"logo":         e.Logo,
		"startDate":    e.StartDate,
		"endDate":      e.EndDate,
		"gpa":          e.GPA,
		"achievements": e.Achievements,
		"createdAt":    e.CreatedAt,
		"updatedAt":    e.UpdatedAt,
	}
}

func toEducationMaps(educations []domain.Education) []map[string]interface{} {
	result := make([]map[string]interface{}, len(educations))
	for i, e := range educations {
		result[i] = toEducationMap(e)
	}
	return result
}

func toExperienceMap(e domain.Experience) map[string]interface{} {
	return map[string]interface{}{
		"id":          e.ID,
		"company":     e.Company,
		"position":    e.Position,
		"type":        e.Type,
		"logo":        e.Logo,
		"startDate":   e.StartDate,
		"endDate":     e.EndDate,
		"tags":        e.Tags,
		"description": e.Description,
		"createdAt":   e.CreatedAt,
		"updatedAt":   e.UpdatedAt,
	}
}

func toExperienceMaps(experiences []domain.Experience) []map[string]interface{} {
	result := make([]map[string]interface{}, len(experiences))
	for i, e := range experiences {
		result[i] = toExperienceMap(e)
	}
	return result
}

func toProjectMap(p domain.Project) map[string]interface{} {
	return map[string]interface{}{
		"id":          p.ID,
		"title":       p.Title,
		"subtitle":    p.Subtitle,
		"description": p.Description,
		"image":       p.Image,
		"tags":        p.Tags,
		"demoUrl":     p.DemoURL,
		"githubUrl":   p.GithubURL,
		"createdAt":   p.CreatedAt,
		"updatedAt":   p.UpdatedAt,
	}
}

func toProjectMaps(projects []domain.Project) []map[string]interface{} {
	result := make([]map[string]interface{}, len(projects))
	for i, p := range projects {
		result[i] = toProjectMap(p)
	}
	return result
}

func toGalleryMap(g domain.Gallery) map[string]interface{} {
	return map[string]interface{}{
		"id":          g.ID,
		"image":       g.Image,
		"imageFileId": g.ImageFileID,
		"caption":     g.Caption,
		"createdAt":   g.CreatedAt,
		"updatedAt":   g.UpdatedAt,
	}
}

func toGalleryMaps(galleries []domain.Gallery) []map[string]interface{} {
	result := make([]map[string]interface{}, len(galleries))
	for i, g := range galleries {
		result[i] = toGalleryMap(g)
	}
	return result
}

func toSocialLinkMap(s domain.SocialLink) map[string]interface{} {
	return map[string]interface{}{
		"id":        s.ID,
		"title":     s.Title,
		"url":       s.URL,
		"order":     s.Order,
		"isActive":  s.IsActive,
		"createdAt": s.CreatedAt,
		"updatedAt": s.UpdatedAt,
	}
}

func toSocialLinkMaps(links []domain.SocialLink) []map[string]interface{} {
	result := make([]map[string]interface{}, len(links))
	for i, s := range links {
		result[i] = toSocialLinkMap(s)
	}
	return result
}
