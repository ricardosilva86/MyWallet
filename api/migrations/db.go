package migrations

import (
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error

	dsn := "host=localhost user=gorm password=gorm dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Amsterdam"
	DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//where myhost is port is the port postgres is running on
	//user is your postgres use name
	//password is your postgres password
	if err != nil {

		panic("failed to connect database")
	}

}
