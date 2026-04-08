package repository

import (
	"context"
	"database/sql"

	"rapdev-graphql/pkg/domain"
)

type socialLinkRepository struct {
	db *sql.DB
}

func NewSocialLinkRepository(db *sql.DB) domain.SocialLinkRepository {
	return &socialLinkRepository{db: db}
}

func (r *socialLinkRepository) FindAll(ctx context.Context, params domain.PaginationParams) (domain.PaginatedResult[domain.SocialLink], error) {
	var totalCount int
	countQuery := `SELECT COUNT(*) FROM social_links WHERE "isActive" = true`
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&totalCount); err != nil {
		return domain.PaginatedResult[domain.SocialLink]{}, err
	}

	query := `
		SELECT id, title, url, "order", "isActive", "createdAt", "updatedAt"
		FROM social_links
		WHERE "isActive" = true
		ORDER BY "order" ASC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, params.Limit, params.Offset())
	if err != nil {
		return domain.PaginatedResult[domain.SocialLink]{}, err
	}
	defer rows.Close()

	var links []domain.SocialLink

	for rows.Next() {
		var s domain.SocialLink
		err := rows.Scan(&s.ID, &s.Title, &s.URL, &s.Order, &s.IsActive, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return domain.PaginatedResult[domain.SocialLink]{}, err
		}
		links = append(links, s)
	}

	if err := rows.Err(); err != nil {
		return domain.PaginatedResult[domain.SocialLink]{}, err
	}

	return domain.NewPaginatedResult(links, totalCount, params), nil
}
