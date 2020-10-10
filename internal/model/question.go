package model

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/meloalright/guora/internal/database"
)

// Question struct
type Question struct {
	GORMBase
	Title             string   `json:"title"`
	Desc              string   `json:"desc" gorm:"type:varchar(4000)"`
	Type              int      `json:"type"`
	QuestionProfile   Profile  `json:"questionProfile" gorm:"ForeignKey:QuestionProfileID"`
	QuestionProfileID int      `json:"questionProfileID" gorm:"index"`
	Answers           []Answer `json:"-"`
	AnswersCounts     int      `json:"answersCounts"`
}

// Get func
func (q *Question) Get() (question Question, err error) {

	if err = database.DB.Where(&q).Preload("QuestionProfile").First(&question).Error; err != nil {
		log.Print(err)
	}

	return
}

// Create func
func (q *Question) Create() (ra int64, err error) {

	if err = database.DB.Create(&q).Error; err != nil {
		ra = -1
	} else {
		ra = 1
	}

	return
}

// Update func
func (q *Question) Update() (ra int64, err error) {

	if err = database.DB.Model(&q).Updates(q).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}

	return
}

// Delete func
func (q *Question) Delete() (ra int64, err error) {

	if err = database.DB.Delete(&q).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}

	return
}

// GetList func
func (q *Question) GetList(limit int, offset int) (questions []Question, err error) {

	if err = database.DB.Offset(offset).Limit(limit).Where(&q).Preload("QuestionProfile").Find(&questions).Error; err != nil {
		log.Print(err)
	}

	return
}

// GetOrderList func
func (q *Question) GetOrderList(limit int, offset int, order string) (questions []Question, err error) {

	if err = database.DB.Offset(offset).Limit(limit).Where(&q).Preload("QuestionProfile").Order(order).Find(&questions).Error; err != nil {
		log.Print(err)
	}

	return
}

// GetCounts func
func (q *Question) GetCounts() (counts int, err error) {

	if err = database.DB.Model(&Question{}).Where(&q).Count(&counts).Error; err != nil {
		log.Print(err)
	}

	return
}

// AfterDelete func
func (q *Question) AfterDelete(tx *gorm.DB) (err error) {

	var a Answer
	a.QuestionID = q.ID

	if err = tx.Where(&a).Delete(&a).Error; err != nil {
		log.Print(err)
	}

	return
}
