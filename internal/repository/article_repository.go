package repository

import (
	"context"
	"database/sql"

	"rapdev-graphql/internal/domain"
)

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) domain.ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) FindAll(ctx context.Context) ([]domain.Article, error) {
	query := `
		SELECT
			a.id, a.title, a.content, a.path, a."viewCount", a.likes,
			a."authorId", a.published, a."createdAt", a."updatedAt",
			u.id, u.name, u.email, u.role, u.image, u."imageFileId",
			u."createdAt", u."updatedAt"
		FROM articles a
		LEFT JOIN users u ON a."authorId" = u.id
		ORDER BY a."createdAt" DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []domain.Article

	for rows.Next() {
		var a domain.Article
		var u domain.User
		var userID sql.NullString

		err := rows.Scan(
			&a.ID, &a.Title, &a.Content, &a.Path, &a.ViewCount, &a.Likes,
			&a.AuthorID, &a.Published, &a.CreatedAt, &a.UpdatedAt,
			&userID, &u.Name, &u.Email, &u.Role, &u.Image, &u.ImageFileID,
			&u.CreatedAt, &u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if userID.Valid {
			u.ID = userID.String
			a.Author = &u
		}

		articles = append(articles, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}
