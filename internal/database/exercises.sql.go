// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: exercises.sql

package database

import (
	"context"
)

const insertDummyExercises = `-- name: InsertDummyExercises :exec
INSERT INTO exercises (name) VALUES
('Push-up'),
('Squat'),
('Deadlift'),
('Bench Press'),
('Pull-up')
`

func (q *Queries) InsertDummyExercises(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, insertDummyExercises)
	return err
}
