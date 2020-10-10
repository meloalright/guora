package model

import (
	"log"

	"github.com/meloalright/guora/internal/database"
)

// Profile struct
type Profile struct {
	GORMBase
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// Get func
func (p *Profile) Get() (profile Profile, err error) {

	if err = database.DB.Where(&p).First(&profile).Error; err != nil {
		log.Print(err)
	}

	return
}

// Create func
func (p *Profile) Create() (ra int64, err error) {

	if err = database.DB.Create(&p).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

// Update func
func (p *Profile) Update() (ra int64, err error) {

	if err = database.DB.Model(&p).Updates(p).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

// Delete func
func (p *Profile) Delete() (ra int64, err error) {

	if err = database.DB.Delete(&p).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

// GetList func
func (p *Profile) GetList(limit int, offset int) (profiles []Profile, err error) {

	if err = database.DB.Offset(offset).Limit(limit).Find(&profiles).Error; err != nil {
		log.Print(err)
	}

	return
}

// GetCounts func
func (p *Profile) GetCounts() (counts int, err error) {

	if err = database.DB.Model(&Profile{}).Count(&counts).Error; err != nil {
		log.Print(err)
	}

	return
}
