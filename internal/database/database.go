package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Connect stabilish connection with database
func Connect(ctx context.Context, URL string) (*pgxpool.Pool, error) {
	pool, erro := pgxpool.Connect(
		ctx,
		URL,
	)
	if erro != nil {
		return &pgxpool.Pool{}, erro
	}
	return pool, nil
}
