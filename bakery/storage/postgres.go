package storage

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	logger  *logrus.Logger
	db      *sql.DB
	sb      squirrel.StatementBuilderType
	querier Querier
}

func NewRepository(logger *logrus.Logger, pgURL *url.URL) (*Repository, error) {
	config, err := pgx.ParseConfig(pgURL.String())
	if err != nil {
		return nil, fmt.Errorf("parsing postgres URI: %w", err)
	}

	config.Logger = logrusadapter.NewLogger(logger)
	db := stdlib.OpenDB(*config)

	err = validateSchema(db)
	if err != nil {
		return nil, fmt.Errorf("validating schema: %w", err)
	}

	return &Repository{
		logger:  logger,
		db:      db,
		sb:      squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db),
		querier: New(db),
	}, nil
}

func (r Repository) Ping() error {
	return r.db.Ping()
}

func (r Repository) CreateBake(ctx context.Context, name string) (*Bake, error) {
	bake, err := r.querier.CreateBake(ctx, CreateBakeParams{
		Name:       name,
		BakeStatus: StatusEnumPreparing,
	})
	if err != nil {
		return nil, err
	}
	return &bake, nil

}

func (r Repository) UpdateBake(ctx context.Context, id pgtype.UUID, status StatusEnum) error {
	err := r.querier.UpdateBake(ctx, UpdateBakeParams{
		ID:         id,
		BakeStatus: status,
	})

	return err
}

func (r Repository) GetBake(ctx context.Context, id pgtype.UUID) (*Bake, error) {
	bake, err := r.querier.GetBake(ctx, id)
	if err != nil {
		return nil, err
	}
	return &bake, nil
}

func (r Repository) GetAllBake(ctx context.Context) (*[]Bake, error) {
	allBakes, err := r.querier.GetAllBake(ctx)
	if err != nil {
		return nil, err
	}
	return &allBakes, nil
}
