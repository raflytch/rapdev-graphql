package repository

import (
	"context"
	"database/sql"

	"rapdev-graphql/pkg/domain"
)

type experienceRepository struct {
	db *sql.DB
}

func NewExperienceRepository(db *sql.DB) domain.ExperienceRepository {
	return &experienceRepository{db: db}
}

func (r *experienceRepository) FindAll(ctx context.Context, params domain.PaginationParams) (domain.PaginatedResult[domain.Experience], error) {
	var totalCount int
	countQuery := `SELECT COUNT(*) FROM experiences`
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&totalCount); err != nil {
		return domain.PaginatedResult[domain.Experience]{}, err
	}

	query := `
		SELECT id, company, position, "type", logo, "startDate", "endDate", tags, description,
		       "createdAt", "updatedAt"
		FROM experiences
		ORDER BY "startDate" DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, params.Limit, params.Offset())
	if err != nil {
		return domain.PaginatedResult[domain.Experience]{}, err
	}
	defer rows.Close()

	var experiences []domain.Experience

	for rows.Next() {
		var e domain.Experience
		err := rows.Scan(
			&e.ID, &e.Company, &e.Position, &e.Type, &e.Logo, &e.StartDate, &e.EndDate,
			(*stringArray)(&e.Tags), (*stringArray)(&e.Description), &e.CreatedAt, &e.UpdatedAt,
		)
		if err != nil {
			return domain.PaginatedResult[domain.Experience]{}, err
		}
		experiences = append(experiences, e)
	}

	if err := rows.Err(); err != nil {
		return domain.PaginatedResult[domain.Experience]{}, err
	}

	return domain.NewPaginatedResult(experiences, totalCount, params), nil
}
