package main

import (
	"database/sql"
	"fmt"
	"github.com/anodamobi/glance-backend/db/migrator"
	dbx "github.com/go-ozzo/ozzo-dbx"

	"github.com/anodamobi/glance-backend"
	"github.com/anodamobi/glance-backend/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type Migrator func(*sql.DB, migrator.MigrateDir) (int, error)

func MigrateDB(direction string, dbClient *dbx.DB, migratorFn Migrator) (int, error) {
	applied, err := migratorFn(dbClient.DB(), migrator.MigrateDir(direction))

	return applied, errors.Wrap(err, "failed to apply migrations")
}

func main() {
	apiConfig := config.New()
	log := apiConfig.Log()

	rootCmd := &cobra.Command{}

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "run command",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := MigrateDB("up", apiConfig.DB().DBX(), migrator.Migrations.Migrate)

			if err != nil {
				log.WithError(err).Error("migration failed")
				return
			}

			api := app.New(apiConfig)
			if err := api.Start(); err != nil {
				panic(errors.Wrap(err, "failed to start api"))
			}
		},
	}

	rootCmd.AddCommand(runCmd)
	if err := rootCmd.Execute(); err != nil {
		log.WithField("cobra", "read").Error(fmt.Sprintf("failed to read command %s", err.Error()))
		return
	}
}
