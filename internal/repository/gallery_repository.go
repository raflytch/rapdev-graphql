package repository

import (
	"context"
	"database/sql"

	"rapdev-graphql/internal/domain"
)

type galleryRepository struct {
	db *sql.DB
}

func NewGalleryRepository(db *sql.DB) domain.GalleryRepository {
	return &galleryRepository{db: db}
}

func (r *galleryRepository) FindAll(ctx context.Context) ([]domain.Gallery, error) {
	query := `
		SELECT id, image, "imageFileId", caption, "createdAt", "updatedAt"
		FROM galleries
		ORDER BY "createdAt" DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var galleries []domain.Gallery

	for rows.Next() {
		var g domain.Gallery
		err := rows.Scan(&g.ID, &g.Image, &g.ImageFileID, &g.Caption, &g.CreatedAt, &g.UpdatedAt)
		if err != nil {
			return nil, err
		}
		galleries = append(galleries, g)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return galleries, nil
}
