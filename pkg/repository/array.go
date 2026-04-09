package repository

import (
	"fmt"
	"strings"
)

// stringArray implements sql.Scanner for PostgreSQL text[] columns.
// This replaces pq.Array() to avoid the lib/pq driver dependency.
type stringArray []string

func (a *stringArray) Scan(src interface{}) error {
	if src == nil {
		*a = nil
		return nil
	}

	var s string
	switch v := src.(type) {
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("stringArray.Scan: unsupported type %T", src)
	}

	parsed, err := parsePostgresArray(s)
	if err != nil {
		return err
	}
	*a = parsed
	return nil
}

func parsePostgresArray(s string) ([]string, error) {
	s = strings.TrimSpace(s)
	if s == "{}" {
		return []string{}, nil
	}
	if len(s) < 2 || s[0] != '{' || s[len(s)-1] != '}' {
		return nil, fmt.Errorf("stringArray: invalid PostgreSQL array format: %q", s)
	}
	s = s[1 : len(s)-1]

	var result []string
	var current strings.Builder
	inQuote := false
	escaped := false

	for i := 0; i < len(s); i++ {
		c := s[i]
		if escaped {
			current.WriteByte(c)
			escaped = false
			continue
		}
		if c == '\\' {
			escaped = true
			continue
		}
		if c == '"' {
			inQuote = !inQuote
			continue
		}
		if c == ',' && !inQuote {
			result = append(result, current.String())
			current.Reset()
			continue
		}
		current.WriteByte(c)
	}
	result = append(result, current.String())

	return result, nil
}
