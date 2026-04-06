package repository

import (
	"context"
	"database/sql"

	"rapdev-graphql/internal/domain"

	"github.com/lib/pq"
)

type educationRepository struct {
	db *sql.DB
}

func NewEducationRepository(db *sql.DB) domain.EducationRepository {
	return &educationRepository{db: db}
}

func (r *educationRepository) FindAll(ctx context.Context) ([]domain.Education, error) {
	query := `
		SELECT id, institution, degree, logo, "startDate", "endDate", gpa, achievements,
		       "createdAt", "updatedAt"
		FROM educations
		ORDER BY "startDate" DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var educations []domain.Education

	for rows.Next() {
		var e domain.Education
		err := rows.Scan(
			&e.ID, &e.Institution, &e.Degree, &e.Logo, &e.StartDate, &e.EndDate,
			&e.GPA, pq.Array(&e.Achievements), &e.CreatedAt, &e.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		educations = append(educations, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return educations, nil
}
