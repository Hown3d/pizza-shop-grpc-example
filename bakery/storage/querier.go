// Code generated by sqlc. DO NOT EDIT.

package storage

import (
	"context"

	"github.com/jackc/pgtype"
)

type Querier interface {
	CreateBake(ctx context.Context, arg CreateBakeParams) (Bake, error)
	DeleteBake(ctx context.Context, id pgtype.UUID) error
	GetAllBake(ctx context.Context) ([]Bake, error)
	GetBake(ctx context.Context, id pgtype.UUID) (Bake, error)
	UpdateBake(ctx context.Context, arg UpdateBakeParams) error
}

var _ Querier = (*Queries)(nil)
