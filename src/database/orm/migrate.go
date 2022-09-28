package orm

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"gorm.io/gorm"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "database migration",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "enable migration")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "disable migration")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := New()

	if err != nil {
		return err
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Histories{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&models.User{}, &models.Vehicle{}, &models.Histories{})
			},
		},
	})

	if migUp {
		if err := m.Migrate(); err != nil {
			return err
		}
		log.Println("migration successfully")
		// return nil
	}

	if migDown {
		if err := m.RollbackLast(); err != nil {
			return err
		}
		log.Println("rollback successfully")
		// return nil
	}

	log.Println("init schema database successfully")
	return nil
}
