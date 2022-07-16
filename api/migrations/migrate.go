package migrations

import "MyWallet/api/models"

func Migrate() error {
	var err error
	err = DBCon.AutoMigrate(models.FixedCost{}, models.VariableSpend{})
	if err != nil {
		return err
	}

	return nil
}
