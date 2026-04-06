package domain

import (
	"context"
	"time"
)

type Article struct {
	ID        string
	Title     string
	Content   string
	Path      string
	ViewCount int
	Likes     int
	AuthorID  string
	Published bool
	Author    *User
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ArticleRepository interface {
	FindAll(ctx context.Context) ([]Article, error)
}
