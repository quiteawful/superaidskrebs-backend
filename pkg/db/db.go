package db

import (
	"errors"
	"log"

	"github.com/gobuffalo/packr/v2"
	"github.com/jmoiron/sqlx"
	"github.com/quiteawful/superaidskrebs-backend/internal/packer"
	"github.com/quiteawful/superaidskrebs-backend/pkg/config"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	pgDriverName = "pgx"
)

func Initialisation(dbc config.DatabaseConf) (*sqlx.DB, error) {
	var err error
	connector := dbc.Driver
	if dbc.Driver == "postgres" {
		connector = pgDriverName
	}
	db, err := sqlx.Connect(connector, dbc.Connection)
	if err != nil {
		return nil, errors.New("Error connecting to database:" + err.Error())
	}

	var migbox *packr.Box

	switch dbc.Driver {
	case "postgres":
		migbox = packer.PostgresBox
	}

	migrations := &migrate.PackrMigrationSource{
		Box: migbox,
		Dir: "./",
	}
	migrate.SetTable("migrations")

	n, err := migrate.Exec(db.DB, dbc.Driver, migrations, migrate.Up)
	if err != nil {
		return nil, errors.New("Error applying migrations:" + err.Error())
	}
	log.Println("Applied ", n, "Migrations")
	return db, nil
}
