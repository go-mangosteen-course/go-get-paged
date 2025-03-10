// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: validation_codes.sql

package queries

import (
	"context"
)

const countValidationCodes = `-- name: CountValidationCodes :one
SELECT count(*) FROM validation_codes WHERE email = $1
`

func (q *Queries) CountValidationCodes(ctx context.Context, email string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countValidationCodes, email)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createValidationCode = `-- name: CreateValidationCode :one
INSERT INTO validation_codes (
  email, code
) VALUES (
  $1, $2
)
RETURNING id, code, email, used_at, created_at, updated_at
`

type CreateValidationCodeParams struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (q *Queries) CreateValidationCode(ctx context.Context, arg CreateValidationCodeParams) (ValidationCode, error) {
	row := q.db.QueryRowContext(ctx, createValidationCode, arg.Email, arg.Code)
	var i ValidationCode
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Email,
		&i.UsedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findValidationCode = `-- name: FindValidationCode :one
SELECT id, code, email, used_at, created_at, updated_at FROM validation_codes
WHERE
  email = $1
  AND
  code = $2
  AND
  used_at is null
ORDER BY created_at desc
LIMIT 1
`

type FindValidationCodeParams struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (q *Queries) FindValidationCode(ctx context.Context, arg FindValidationCodeParams) (ValidationCode, error) {
	row := q.db.QueryRowContext(ctx, findValidationCode, arg.Email, arg.Code)
	var i ValidationCode
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Email,
		&i.UsedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
