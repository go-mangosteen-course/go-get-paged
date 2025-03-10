// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: items.sql

package queries

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const countItems = `-- name: CountItems :one
SELECT count(*) FROM items
`

func (q *Queries) CountItems(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countItems)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createItem = `-- name: CreateItem :one
INSERT INTO items (
  user_id,
  amount,
  kind,
  happened_at,
  tag_ids
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING id, user_id, amount, tag_ids, kind, happened_at, created_at, updated_at
`

type CreateItemParams struct {
	UserID     int32     `json:"user_id"`
	Amount     int32     `json:"amount"`
	Kind       Kind      `json:"kind"`
	HappenedAt time.Time `json:"happened_at"`
	TagIds     []int32   `json:"tag_ids"`
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.UserID,
		arg.Amount,
		arg.Kind,
		arg.HappenedAt,
		pq.Array(arg.TagIds),
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		pq.Array(&i.TagIds),
		&i.Kind,
		&i.HappenedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAllItems = `-- name: DeleteAllItems :exec
DELETE FROM items
`

func (q *Queries) DeleteAllItems(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllItems)
	return err
}

const listItems = `-- name: ListItems :many
SELECT id, user_id, amount, tag_ids, kind, happened_at, created_at, updated_at from items
ORDER BY happened_at DESC
OFFSET $1
LIMIT $2
`

type ListItemsParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListItems(ctx context.Context, arg ListItemsParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItems, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			pq.Array(&i.TagIds),
			&i.Kind,
			&i.HappenedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
