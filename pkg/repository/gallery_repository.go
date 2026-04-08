package repository

import (
	"context"
	"database/sql"

	"rapdev-graphql/pkg/domain"
)

type galleryRepository struct {
	db *sql.DB
}

func NewGalleryRepository(db *sql.DB) domain.GalleryRepository {
	return &galleryRepository{db: db}
}

func (r *galleryRepository) FindAll(ctx context.Context, params domain.PaginationParams) (domain.PaginatedResult[domain.Gallery], error) {
	var totalCount int
	countQuery := `SELECT COUNT(*) FROM galleries`
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&totalCount); err != nil {
		return domain.PaginatedResult[domain.Gallery]{}, err
	}

	query := `
		SELECT id, image, "imageFileId", caption, "createdAt", "updatedAt"
		FROM galleries
		ORDER BY "createdAt" DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, params.Limit, params.Offset())
	if err != nil {
		return domain.PaginatedResult[domain.Gallery]{}, err
	}
	defer rows.Close()

	var galleries []domain.Gallery

	for rows.Next() {
		var g domain.Gallery
		err := rows.Scan(&g.ID, &g.Image, &g.ImageFileID, &g.Caption, &g.CreatedAt, &g.UpdatedAt)
		if err != nil {
			return domain.PaginatedResult[domain.Gallery]{}, err
		}
		galleries = append(galleries, g)
	}

	if err := rows.Err(); err != nil {
		return domain.PaginatedResult[domain.Gallery]{}, err
	}

	return domain.NewPaginatedResult(galleries, totalCount, params), nil
}
