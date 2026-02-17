package models

import (
	"context"

	"github.com/corbinlazarone/todue/cmd/internals/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseModel struct {
	DB *pgxpool.Pool
}

// Creates a course entry and returns the created UUID
func (c *CourseModel) CreateNewCourseEntry(ctx context.Context, courseId int) (int, error) {
	statement := `INSERT INTO courses (course_name) VALUES ($1) RETURNING id`

	// FIX: todue-server  | ERROR: 2026/02/17 19:56:33
	// handlers_calendar.go:146: failed to encode args[0]: unable to encode 1 into text
	// format for text (OID 25): cannot find encode plan
	var id int
	err := c.DB.QueryRow(ctx, statement, courseId).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// Creates a new assignment for the course matching the provided course id
func (c *CourseModel) CreateNewAssignment(ctx context.Context, courseId int, assignment types.Assignment) error {
	statement :=
		`INSERT INTO assignment (
		course_id,
		name,
		descritpion,
		due_date,
		all_day,
		color,
		start_time,
		end_time,
		reminder
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9
	)`

	_, err := c.DB.Exec(ctx, statement, courseId)

	if err != nil {
		return err
	}

	return nil
}
