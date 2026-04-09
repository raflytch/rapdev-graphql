package repository

import (
	"context"
	"database/sql"

	"rapdev-graphql/pkg/domain"
)

type projectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) domain.ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) FindAll(ctx context.Context, params domain.PaginationParams) (domain.PaginatedResult[domain.Project], error) {
	var totalCount int
	countQuery := `SELECT COUNT(*) FROM projects`
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&totalCount); err != nil {
		return domain.PaginatedResult[domain.Project]{}, err
	}

	query := `
		SELECT id, title, subtitle, description, image, tags, "demoUrl", "githubUrl",
		       "createdAt", "updatedAt"
		FROM projects
		ORDER BY "createdAt" DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, params.Limit, params.Offset())
	if err != nil {
		return domain.PaginatedResult[domain.Project]{}, err
	}
	defer rows.Close()

	var projects []domain.Project

	for rows.Next() {
		var p domain.Project
		err := rows.Scan(
			&p.ID, &p.Title, &p.Subtitle, &p.Description, &p.Image,
			(*stringArray)(&p.Tags), &p.DemoURL, &p.GithubURL, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return domain.PaginatedResult[domain.Project]{}, err
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return domain.PaginatedResult[domain.Project]{}, err
	}

	return domain.NewPaginatedResult(projects, totalCount, params), nil
}
