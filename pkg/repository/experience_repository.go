package repository

import (
	"context"
	"database/sql"

	"rapdev-graphql/pkg/domain"

	"github.com/lib/pq"
)

type experienceRepository struct {
	db *sql.DB
}

func NewExperienceRepository(db *sql.DB) domain.ExperienceRepository {
	return &experienceRepository{db: db}
}

func (r *experienceRepository) FindAll(ctx context.Context) ([]domain.Experience, error) {
	query := `
		SELECT id, company, position, "type", logo, "startDate", "endDate", tags, description,
		       "createdAt", "updatedAt"
		FROM experiences
		ORDER BY "startDate" DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiences []domain.Experience

	for rows.Next() {
		var e domain.Experience
		err := rows.Scan(
			&e.ID, &e.Company, &e.Position, &e.Type, &e.Logo, &e.StartDate, &e.EndDate,
			pq.Array(&e.Tags), pq.Array(&e.Description), &e.CreatedAt, &e.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		experiences = append(experiences, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return experiences, nil
}
