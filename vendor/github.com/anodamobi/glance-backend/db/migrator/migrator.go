package migrator

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/rubenv/sql-migrate"
)

// MigrateDir represents a direction in which to perform schema migrations.
type MigrateDir string

const (
	// MigrateUp causes migrations to be run in the "up" direction.
	MigrateUp MigrateDir = "up"
	// MigrateDown causes migrations to be run in the "down" direction.
	MigrateDown MigrateDir = "down"
)

type MigrationsLoader struct {
	source *migrate.AssetMigrationSource
}

func NewMigrationsLoader() *MigrationsLoader {
	return &MigrationsLoader{}
}

func (l *MigrationsLoader) loadDir(dir string) {
	l.source = &migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      dir,
	}
}

// Migrate performs schema migration.  Migrations can occur in one of three
// ways:
//
// - up: migrations are performed from the currently installed version upwards.
// If count is 0, all unapplied migrations will be run.
//
// - down: migrations are performed from the current version downard. If count
// is 0, all applied migrations will be run in a downard direction.
//
// - redo: migrations are first ran downard `count` times, and then are ran
// upward back to the current version at the start of the process. If count is
// 0, a count of 1 will be assumed.
func (l *MigrationsLoader) Migrate(dbClient *sql.DB, dir MigrateDir) (int, error) {
	switch dir {
	case MigrateUp:
		return migrate.ExecMax(dbClient, "postgres", l.source, migrate.Up, 0)

	case MigrateDown:
		return migrate.ExecMax(dbClient, "postgres", l.source, migrate.Down, 0)

	default:
		return 0, errors.New("Invalid migration direction")
	}
}
