package repository

import (
	"context"
	"database/sql"

	"rapdev-graphql/pkg/domain"
)

type educationRepository struct {
	db *sql.DB
}

func NewEducationRepository(db *sql.DB) domain.EducationRepository {
	return &educationRepository{db: db}
}

func (r *educationRepository) FindAll(ctx context.Context, params domain.PaginationParams) (domain.PaginatedResult[domain.Education], error) {
	var totalCount int
	countQuery := `SELECT COUNT(*) FROM educations`
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&totalCount); err != nil {
		return domain.PaginatedResult[domain.Education]{}, err
	}

	query := `
		SELECT id, institution, degree, logo, "startDate", "endDate", gpa, achievements,
		       "createdAt", "updatedAt"
		FROM educations
		ORDER BY "startDate" DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, params.Limit, params.Offset())
	if err != nil {
		return domain.PaginatedResult[domain.Education]{}, err
	}
	defer rows.Close()

	var educations []domain.Education

	for rows.Next() {
		var e domain.Education
		err := rows.Scan(
			&e.ID, &e.Institution, &e.Degree, &e.Logo, &e.StartDate, &e.EndDate,
			&e.GPA, (*stringArray)(&e.Achievements), &e.CreatedAt, &e.UpdatedAt,
		)
		if err != nil {
			return domain.PaginatedResult[domain.Education]{}, err
		}
		educations = append(educations, e)
	}

	if err := rows.Err(); err != nil {
		return domain.PaginatedResult[domain.Education]{}, err
	}

	return domain.NewPaginatedResult(educations, totalCount, params), nil
}
