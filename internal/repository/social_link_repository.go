package repository

import (
	"context"
	"database/sql"

	"rapdev-graphql/internal/domain"
)

type socialLinkRepository struct {
	db *sql.DB
}

func NewSocialLinkRepository(db *sql.DB) domain.SocialLinkRepository {
	return &socialLinkRepository{db: db}
}

func (r *socialLinkRepository) FindAll(ctx context.Context) ([]domain.SocialLink, error) {
	query := `
		SELECT id, title, url, "order", "isActive", "createdAt", "updatedAt"
		FROM social_links
		WHERE "isActive" = true
		ORDER BY "order" ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []domain.SocialLink

	for rows.Next() {
		var s domain.SocialLink
		err := rows.Scan(&s.ID, &s.Title, &s.URL, &s.Order, &s.IsActive, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, err
		}
		links = append(links, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}
