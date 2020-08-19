package database

import (
	"github.com/meloalright/guora/conf"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var SQLITE3DB *gorm.DB
var err error

func init() {

	SQLITE3DB, err = gorm.Open("sqlite3", conf.Config().Sql.Addr)

	SQLITE3DB.Callback().Create().Remove("gorm:update_time_stamp")
	SQLITE3DB.Callback().Update().Remove("gorm:update_time_stamp")

	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
}
