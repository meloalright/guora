package database

import (
	"log"

	"github.com/meloalright/guora/conf"

	"github.com/jinzhu/gorm"

	// for multi select of db
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB reference
var DB *gorm.DB
var err error

func init() {

	DB, err = gorm.Open(conf.Config().DB.Driver, conf.Config().DB.Addr)

	DB.Callback().Create().Remove("gorm:update_time_stamp")
	DB.Callback().Update().Remove("gorm:update_time_stamp")

	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
}
