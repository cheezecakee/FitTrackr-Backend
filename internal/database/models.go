// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Exercise struct {
	ID        uuid.UUID
	Name      string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type RefreshToken struct {
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	IsRevoked bool
	ExpiresAt time.Time
	RevokedAt sql.NullTime
}

type User struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	PasswordHash string
	Email        string
	Age          sql.NullInt32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsPremium    bool
}

type Workout struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Name        string
	Description sql.NullString
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

type WorkoutExercise struct {
	ID         uuid.UUID
	WorkoutID  uuid.UUID
	ExerciseID uuid.UUID
	Sets       int32
	RepsMin    int32
	RepsMax    int32
	Weight     float64
	Interval   int32
	Rest       int32
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
}
