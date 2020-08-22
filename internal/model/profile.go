package model

import (
	"log"

	"github.com/meloalright/guora/internal/database"
)

type Profile struct {
	GORMBase
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (p *Profile) Get() (profile Profile, err error) {

	if err = database.SQLITE3DB.Where(&p).First(&profile).Error; err != nil {
		log.Print(err)
	}

	return
}

func (p *Profile) Create() (ra int64, err error) {

	if err = database.SQLITE3DB.Create(&p).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

func (p *Profile) Update() (ra int64, err error) {

	if err = database.SQLITE3DB.Model(&p).Updates(p).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

func (p *Profile) Delete() (ra int64, err error) {

	if err = database.SQLITE3DB.Delete(&p).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

func (p *Profile) GetList(limit int, offset int) (profiles []Profile, err error) {

	if err = database.SQLITE3DB.Offset(offset).Limit(limit).Find(&profiles).Error; err != nil {
		log.Print(err)
	}

	return
}

func (p *Profile) GetCounts() (counts int, err error) {

	if err = database.SQLITE3DB.Model(&Profile{}).Count(&counts).Error; err != nil {
		log.Print(err)
	}

	return
}
