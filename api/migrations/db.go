package migrations

import (
	"MyWallet/util"
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error

	config, err := util.LoadDBConfig(".")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", config.Host, config.User, config.Password, config.DBname, config.Port, config.SSLmode, config.TimeZone)
	DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//where myhost is port is the port postgres is running on
	//user is your postgres use name
	//password is your postgres password
	if err != nil {

		panic("failed to connect database")
	}

}
