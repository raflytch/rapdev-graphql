package domain

import "time"

type User struct {
	ID          string
	Name        string
	Email       string
	Role        string
	Image       *string
	ImageFileID *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
