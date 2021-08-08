package storage

import (
	"database/sql"
	"embed"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgtype"
)

//go:embed migrations/*.sql
var fs embed.FS

// version defines the current migration version. This ensures the app
// is always compatible with the version of the database.
const version = 1

// Migrate migrates the Postgres schema to the current version.
func validateSchema(db *sql.DB) error {
	sourceInstance, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}
	targetInstance, err := postgres.WithInstance(db, new(postgres.Config))
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", sourceInstance, "postgres", targetInstance)
	if err != nil {
		return err
	}
	err = m.Migrate(version) // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return sourceInstance.Close()
}

func IdPostgresToProto(id pgtype.UUID) (proto_id string, err error) {
	err = id.AssignTo(&proto_id)
	if err != nil {
		return "", err
	}
	return
}

func IdProtoToPostgres(id string) (postgres_id pgtype.UUID, err error) {
	err = postgres_id.Set(id)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return
}
