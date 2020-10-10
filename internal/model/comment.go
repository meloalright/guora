package model

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/meloalright/guora/internal/database"
)

// Comment struct
type Comment struct {
	GORMBase
	Content          string  `json:"content"`
	Type             int     `json:"type"`
	Answer           Answer  `json:"-" gorm:"ForeignKey:AnswerID"`
	AnswerID         int     `json:"answerID"`
	CommentProfile   Profile `json:"commentProfile" gorm:"ForeignKey:CommentProfileID"`
	CommentProfileID int     `json:"commentProfileID"`
	Replies          []Reply `json:"-"`
	RepliesCounts    int     `json:"repliesCounts"`
}

// Get func
func (co *Comment) Get() (comment Comment, err error) {

	if err = database.DB.Where(&co).Preload("CommentProfile").First(&comment).Error; err != nil {
		log.Print(err)
	}

	return
}

// Create func
func (co *Comment) Create() (ra int64, err error) {

	if err = database.DB.Create(&co).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}

	return
}

// Update func
func (co *Comment) Update() (ra int64, err error) {

	if err = database.DB.Model(&co).Updates(co).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

// Delete func
func (co *Comment) Delete() (ra int64, err error) {

	if err = database.DB.Where(&co).First(&co).Delete(&co).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

// GetList func
func (co *Comment) GetList(limit int, offset int) (comments []Comment, err error) {

	if err = database.DB.Offset(offset).Limit(limit).Preload("CommentProfile").Find(&comments, co).Error; err != nil {
		log.Print(err)
	}

	return
}

// GetCounts func
func (co *Comment) GetCounts() (counts int, err error) {

	if err = database.DB.Model(&Comment{}).Where(&co).Count(&counts).Error; err != nil {
		log.Print(err)
	}

	return
}

// AfterCreate func
func (co *Comment) AfterCreate(tx *gorm.DB) (err error) {

	var a Answer
	a.ID = co.AnswerID

	if err = tx.Model(&a).UpdateColumn("comments_counts", gorm.Expr("comments_counts + ?", 1)).Error; err != nil {
		log.Print(err)
	}

	return
}

// AfterDelete func
func (co *Comment) AfterDelete(tx *gorm.DB) (err error) {

	var a Answer
	a.ID = co.AnswerID

	if err = tx.Model(&a).UpdateColumn("comments_counts", gorm.Expr("comments_counts - ?", 1)).Error; err != nil {
		log.Print(err)
	}

	return
}
