package mysqlconn

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	DBName      string
	TablePrefix string
}

func GetDB(setting Database) *gorm.DB {
	newDb, err := gorm.Open(setting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.User,
		setting.Password,
		setting.Host,
		setting.DBName))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	newDb.SingularTable(true)
	newDb.DB().SetMaxIdleConns(10)
	newDb.DB().SetMaxOpenConns(100)
	return newDb
}

func CloseDB(db *gorm.DB) {
	defer db.Close()
}
