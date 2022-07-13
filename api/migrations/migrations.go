package migrations

import (
	"MyWallet/api/dblayer"
	"MyWallet/api/models"
)

func Migrate() {
	dblayer.DBORM.AutoMigrate(models.FixedCost{})
}
