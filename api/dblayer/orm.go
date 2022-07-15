package dblayer

import (
	"MyWallet/api/models"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Amsterdam"
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
