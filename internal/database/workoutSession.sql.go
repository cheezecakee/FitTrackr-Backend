// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: workoutSession.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const completeExerciseSession = `-- name: CompleteExerciseSession :one
UPDATE exercise_session
SET completed = true, is_active = false, ended_at = NOW()
WHERE id = $1 AND is_active = true
RETURNING id
`

func (q *Queries) CompleteExerciseSession(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, completeExerciseSession, id)
	err := row.Scan(&id)
	return id, err
}

const completeWorkoutSession = `-- name: CompleteWorkoutSession :one
UPDATE workout_session
SET completed = true, is_active = false, ended_at = NOW()
WHERE id = $1 AND is_active = true
RETURNING id
`

func (q *Queries) CompleteWorkoutSession(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, completeWorkoutSession, id)
	err := row.Scan(&id)
	return id, err
}

const createExerciseSession = `-- name: CreateExerciseSession :one
WITH active_session AS (
    SELECT id
    FROM workout_session
    WHERE user_id = $1 AND is_active = true AND completed = false
    LIMIT 1
)
INSERT INTO exercise_session (workout_session_id, workout_exercise_id, is_active)
SELECT (SELECT id FROM active_session), $2, true
WHERE EXISTS (SELECT 1 FROM active_session)
RETURNING id
`

type CreateExerciseSessionParams struct {
	UserID            uuid.UUID
	WorkoutExerciseID uuid.UUID
}

func (q *Queries) CreateExerciseSession(ctx context.Context, arg CreateExerciseSessionParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createExerciseSession, arg.UserID, arg.WorkoutExerciseID)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createWorkoutSession = `-- name: CreateWorkoutSession :one
WITH existing_session AS (
    SELECT id
    FROM workout_session
    WHERE user_id = $1 AND is_active = true AND completed = false
    LIMIT 1
)
INSERT INTO workout_session (user_id, workout_id, is_active, expires_at)
SELECT $1, $2, true, NOW() + INTERVAL '6 hours'
WHERE NOT EXISTS (SELECT 1 FROM existing_session)
RETURNING id
`

type CreateWorkoutSessionParams struct {
	UserID    uuid.UUID
	WorkoutID uuid.UUID
}

func (q *Queries) CreateWorkoutSession(ctx context.Context, arg CreateWorkoutSessionParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createWorkoutSession, arg.UserID, arg.WorkoutID)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getActiveExerciseSession = `-- name: GetActiveExerciseSession :one
SELECT id, workout_exercise_id
FROM exercise_session
WHERE workout_session_id = $1 AND is_active = true AND completed = false
LIMIT 1
`

type GetActiveExerciseSessionRow struct {
	ID                uuid.UUID
	WorkoutExerciseID uuid.UUID
}

func (q *Queries) GetActiveExerciseSession(ctx context.Context, workoutSessionID uuid.UUID) (GetActiveExerciseSessionRow, error) {
	row := q.db.QueryRowContext(ctx, getActiveExerciseSession, workoutSessionID)
	var i GetActiveExerciseSessionRow
	err := row.Scan(&i.ID, &i.WorkoutExerciseID)
	return i, err
}

const getActiveWorkoutSession = `-- name: GetActiveWorkoutSession :one
SELECT id
FROM workout_session
WHERE user_id = $1 AND is_active = true AND completed = false
LIMIT 1
`

func (q *Queries) GetActiveWorkoutSession(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getActiveWorkoutSession, userID)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const stopExerciseSession = `-- name: StopExerciseSession :one
UPDATE exercise_session
SET is_active = false, ended_at = NOW()
WHERE id = $1 AND is_active = true
RETURNING id
`

func (q *Queries) StopExerciseSession(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, stopExerciseSession, id)
	err := row.Scan(&id)
	return id, err
}

const stopSession = `-- name: StopSession :one
BEGIN
`

type StopSessionRow struct {
}

func (q *Queries) StopSession(ctx context.Context) (StopSessionRow, error) {
	row := q.db.QueryRowContext(ctx, stopSession)
	var i StopSessionRow
	err := row.Scan()
	return i, err
}

const stopWorkoutSession = `-- name: StopWorkoutSession :one
UPDATE workout_session
SET is_active = false, ended_at = NOW()
WHERE id = $1 AND is_active = true
RETURNING id
`

func (q *Queries) StopWorkoutSession(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, stopWorkoutSession, id)
	err := row.Scan(&id)
	return id, err
}
