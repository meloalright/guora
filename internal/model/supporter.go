package model

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/meloalright/guora/internal/database"
)

// Supporter struct
type Supporter struct {
	GORMBase
	Answer    Answer  `json:"-" gorm:"ForeignKey:AnswerID"`
	AnswerID  int     `json:"answerID"`
	Profile   Profile `json:"profile" gorm:"ForeignKey:ProfileID"`
	ProfileID int     `json:"profileID"`
}

// Create func
func (su *Supporter) Create() (ra int64, err error) {

	if err = database.DB.Create(&su).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}

	return
}

// Delete func
func (su *Supporter) Delete() (ra int64, err error) {

	if err = database.DB.Delete(&su).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

// AfterCreate func
func (su *Supporter) AfterCreate(tx *gorm.DB) (err error) {

	var a Answer
	a.ID = su.AnswerID

	if err = tx.Model(&a).UpdateColumn("supporters_counts", gorm.Expr("supporters_counts + ?", 1)).Error; err != nil {
		log.Print(err)
	}

	return
}

// AfterDelete func
func (su *Supporter) AfterDelete(tx *gorm.DB) (err error) {

	var a Answer
	a.ID = su.AnswerID

	if err = tx.Model(&a).UpdateColumn("supporters_counts", gorm.Expr("supporters_counts - ?", 1)).Error; err != nil {
		log.Print(err)
	}

	return
}
