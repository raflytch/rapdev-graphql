package domain

type PaginationParams struct {
	Page  int
	Limit int
}

func NewPaginationParams(page, limit int) PaginationParams {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	return PaginationParams{Page: page, Limit: limit}
}

func (p PaginationParams) Offset() int {
	return (p.Page - 1) * p.Limit
}

type PaginatedResult[T any] struct {
	Data       []T
	TotalCount int
	Page       int
	Limit      int
	TotalPages int
}

func NewPaginatedResult[T any](data []T, totalCount int, params PaginationParams) PaginatedResult[T] {
	totalPages := totalCount / params.Limit
	if totalCount%params.Limit != 0 {
		totalPages++
	}
	return PaginatedResult[T]{
		Data:       data,
		TotalCount: totalCount,
		Page:       params.Page,
		Limit:      params.Limit,
		TotalPages: totalPages,
	}
}
