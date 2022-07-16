package dblayer

import (
	"MyWallet/api/models"
	"MyWallet/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	config, err := util.LoadDBConfig(".")
	if err != nil {
		log.Fatal("error parsing config file")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", config.Host, config.User, config.Password, config.DBname, config.Port, config.SSLmode, config.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//mysql connection bellow:
	//db, err := gorm.Open(dbname, con+"?parseTime=true")
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllFixedCosts() (fc []models.FixedCost, err error) {
	return fc, db.Find(&fc).Error
}
