package model

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/meloalright/guora/internal/database"
)

type Supporter struct {
	GORMBase
	Answer    Answer  `json:"-" gorm:"ForeignKey:AnswerID"`
	AnswerID  int     `json:"answerID"`
	Profile   Profile `json:"profile" gorm:"ForeignKey:ProfileID"`
	ProfileID int     `json:"profileID"`
}

func (su *Supporter) Create() (ra int64, err error) {

	if err = database.SQLITE3DB.Create(&su).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}

	return
}

func (su *Supporter) Delete() (ra int64, err error) {

	if err = database.SQLITE3DB.Delete(&su).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

func (su *Supporter) AfterCreate(tx *gorm.DB) (err error) {

	var a Answer
	a.ID = su.AnswerID

	if err = tx.Model(&a).UpdateColumn("supporters_counts", gorm.Expr("supporters_counts + ?", 1)).Error; err != nil {
		log.Print(err)
	}

	return
}

func (su *Supporter) AfterDelete(tx *gorm.DB) (err error) {

	var a Answer
	a.ID = su.AnswerID

	if err = tx.Model(&a).UpdateColumn("supporters_counts", gorm.Expr("supporters_counts - ?", 1)).Error; err != nil {
		log.Print(err)
	}

	return
}
