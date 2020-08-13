package model

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/meloalright/guora/database"
)

type Answer struct {
	GORMBase
	Content          string      `json:"content"`
	Type             int         `json:"type"`
	Question         Question    `json:"question" gorm:"ForeignKey:QuestionID"`
	QuestionID       int         `json:"questionID"`
	AnswerProfile    Profile     `json:"answerProfile" gorm:"ForeignKey:AnswerProfileID"`
	AnswerProfileID  int         `json:"answerProfileID"`
	Comments         []Comment   `json:"-"`
	CommentsCounts   int         `json:"commentsCounts"`
	Supporters       []Supporter `json:"-"`
	SupportersCounts int         `json:"supportersCounts"`
	Supported        bool        `json:"supported" gorm:"-"`
}

func (a *Answer) Get() (answer Answer, err error) {

	if err = database.SQLITE3DB.Where(&a).Preload("Question").Preload("AnswerProfile").First(&answer).Error; err != nil {
		log.Print(err)
	}

	return
}

func (a *Answer) Create() (ra int64, err error) {

	if err = database.SQLITE3DB.Create(&a).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}

	return
}

func (a *Answer) Update() (ra int64, err error) {

	if err = database.SQLITE3DB.Model(&a).Updates(a).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

func (a *Answer) Delete() (ra int64, err error) {

	if err = database.SQLITE3DB.Where(&a).First(&a).Delete(&a).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

func (a *Answer) GetList(limit int, offset int) (answers []Answer, err error) {

	if err = database.SQLITE3DB.Offset(offset).Limit(limit).Preload("Question").Preload("AnswerProfile").Find(&answers, a).Error; err != nil {
		log.Print(err)
	}

	return
}

func (a *Answer) GetOrderList(limit int, offset int, order string) (answers []Answer, err error) {

	if err = database.SQLITE3DB.Offset(offset).Limit(limit).Preload("Question").Preload("AnswerProfile").Order(order).Find(&answers, a).Error; err != nil {
		log.Print(err)
	}

	return
}

func (a *Answer) GetCounts() (counts int, err error) {

	if err = database.SQLITE3DB.Model(&Answer{}).Where(&a).Count(&counts).Error; err != nil {
		log.Print(err)
	}
	return

}

func (a *Answer) AfterCreate(tx *gorm.DB) (err error) {

	var q Question
	q.ID = a.QuestionID

	if err = tx.Model(&q).UpdateColumn("answers_counts", gorm.Expr("answers_counts + ?", 1)).Error; err != nil {
		log.Print(err)
	}

	return
}

func (a *Answer) AfterDelete(tx *gorm.DB) (err error) {

	var q Question
	q.ID = a.QuestionID

	if err = tx.Model(&q).UpdateColumn("answers_counts", gorm.Expr("answers_counts - ?", 1)).Error; err != nil {
		log.Print(err)
	}

	return
}
