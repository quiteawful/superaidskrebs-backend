package packer

import "github.com/gobuffalo/packr/v2"

//PostgresBox holds the psql migrations
var PostgresBox *packr.Box

//PackAll packs the migrations into the go executable
func PackAll() {
	PostgresBox = packr.New("psqlmigrations", "../../scripts/sql/migrations/postgres")
}
