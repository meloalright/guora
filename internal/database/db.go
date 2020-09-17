package database

import (
	"github.com/meloalright/guora/conf"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

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
