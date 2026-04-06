package repository

import (
	"context"
	"database/sql"

	"rapdev-graphql/pkg/domain"

	"github.com/lib/pq"
)

type projectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) domain.ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) FindAll(ctx context.Context) ([]domain.Project, error) {
	query := `
		SELECT id, title, subtitle, description, image, tags, "demoUrl", "githubUrl",
		       "createdAt", "updatedAt"
		FROM projects
		ORDER BY "createdAt" DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []domain.Project

	for rows.Next() {
		var p domain.Project
		err := rows.Scan(
			&p.ID, &p.Title, &p.Subtitle, &p.Description, &p.Image,
			pq.Array(&p.Tags), &p.DemoURL, &p.GithubURL, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
