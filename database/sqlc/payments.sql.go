// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: payments.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments (paymented_time)
VALUES ($1)
RETURNING payment_id, paymented_time
`

type CreatePaymentRow struct {
	PaymentID     uuid.UUID
	PaymentedTime time.Time
}

func (q *Queries) CreatePayment(ctx context.Context, paymentedTime time.Time) (CreatePaymentRow, error) {
	row := q.db.QueryRowContext(ctx, createPayment, paymentedTime)
	var i CreatePaymentRow
	err := row.Scan(&i.PaymentID, &i.PaymentedTime)
	return i, err
}

const deletePayment = `-- name: DeletePayment :exec
DELETE FROM payments
WHERE payment_id = $1
`

// Delete a payment
func (q *Queries) DeletePayment(ctx context.Context, paymentID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePayment, paymentID)
	return err
}

const getPayment = `-- name: GetPayment :one
SELECT payment_id, paymented_time
FROM payments
WHERE payment_id = $1
`

type GetPaymentRow struct {
	PaymentID     uuid.UUID
	PaymentedTime time.Time
}

func (q *Queries) GetPayment(ctx context.Context, paymentID uuid.UUID) (GetPaymentRow, error) {
	row := q.db.QueryRowContext(ctx, getPayment, paymentID)
	var i GetPaymentRow
	err := row.Scan(&i.PaymentID, &i.PaymentedTime)
	return i, err
}

const getPaymentsByOrderID = `-- name: GetPaymentsByOrderID :many
SELECT payment_id, paymented_time
FROM payments
WHERE paymented_time >= $1
`

type GetPaymentsByOrderIDRow struct {
	PaymentID     uuid.UUID
	PaymentedTime time.Time
}

// Get all payments for a given order
func (q *Queries) GetPaymentsByOrderID(ctx context.Context, paymentedTime time.Time) ([]GetPaymentsByOrderIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentsByOrderID, paymentedTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPaymentsByOrderIDRow
	for rows.Next() {
		var i GetPaymentsByOrderIDRow
		if err := rows.Scan(&i.PaymentID, &i.PaymentedTime); err != nil {
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

const updatePayment = `-- name: UpdatePayment :exec
UPDATE payments
SET paymented_time = $1
WHERE payment_id = $2
`

type UpdatePaymentParams struct {
	PaymentedTime time.Time
	PaymentID     uuid.UUID
}

func (q *Queries) UpdatePayment(ctx context.Context, arg UpdatePaymentParams) error {
	_, err := q.db.ExecContext(ctx, updatePayment, arg.PaymentedTime, arg.PaymentID)
	return err
}
