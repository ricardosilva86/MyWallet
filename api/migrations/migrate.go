package migrations

import "MyWallet/api/models"

func Migrate() error {
	var err error
	err = DBCon.AutoMigrate(models.FixedCost{})
	if err != nil {
		return err
	}

	return nil
}
