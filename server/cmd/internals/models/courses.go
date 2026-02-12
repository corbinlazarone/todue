package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseModel struct {
	DB *pgxpool.Pool
}

func (c *CourseModel) CreateNewCourseEntry(ctx context.Context) error {
	return nil
}
