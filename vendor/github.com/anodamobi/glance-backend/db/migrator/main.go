package migrator

//go:generate go-bindata -nometadata -ignore .+\.go$ -pkg migrator -o bindata.go ./...
//go:generate gofmt -w bindata.go

const (
	// MigrationsDir is folder where migrations are store
	MigrationsDir = "migrations"
)

var (
	Migrations *MigrationsLoader
)

func init() {
	Migrations = NewMigrationsLoader()

	Migrations.loadDir(MigrationsDir)
}
