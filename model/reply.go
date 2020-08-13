package model

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/meloalright/guora/database"
)

type Reply struct {
	GORMBase
	Content            string  `json:"content"`
	Type               int     `json:"type"`
	Comment            Comment `json:"-" gorm:"ForeignKey:CommentID"`
	CommentID          int     `json:"commentID"`
	ReplyFromProfile   Profile `json:"replyFromProfile" gorm:"ForeignKey:ReplyFromProfileID"`
	ReplyFromProfileID int     `json:"replyFromProfileID"`
	ReplyToProfile     Profile `json:"replyToProfile" gorm:"ForeignKey:ReplyToProfileID"`
	ReplyToProfileID   int     `json:"replyToProfileID"`
}

func (r *Reply) Get() (reply Reply, err error) {

	if err = database.SQLITE3DB.Where(&r).Preload("ReplyFromProfile").Preload("ReplyToProfile").First(&reply).Error; err != nil {
		log.Print(err)
	}

	return
}

func (r *Reply) Create() (ra int64, err error) {

	if err = database.SQLITE3DB.Create(&r).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}

	return
}

func (r *Reply) Update() (ra int64, err error) {

	if err = database.SQLITE3DB.Model(&r).Updates(r).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

func (r *Reply) Delete() (ra int64, err error) {

	if err = database.SQLITE3DB.Where(&r).First(&r).Delete(&r).Error; err != nil {
		ra = -1
		log.Print(err)
	} else {
		ra = 1
	}
	return
}

func (r *Reply) GetList(limit int, offset int) (replies []Reply, err error) {

	if err = database.SQLITE3DB.Offset(offset).Limit(limit).Preload("ReplyFromProfile").Preload("ReplyToProfile").Find(&replies, r).Error; err != nil {
		log.Print(err)
	}

	return
}

func (r *Reply) GetCounts() (counts int, err error) {

	if err = database.SQLITE3DB.Model(&Reply{}).Where(&r).Count(&counts).Error; err != nil {
		log.Print(err)
	}

	return
}

func (r *Reply) AfterCreate(tx *gorm.DB) (err error) {

	var co Comment
	co.ID = r.CommentID

	if err = tx.Model(&co).UpdateColumn("replies_counts", gorm.Expr("replies_counts + ?", 1)).Error; err != nil {
		log.Print(err)
	}

	return
}

func (r *Reply) AfterDelete(tx *gorm.DB) (err error) {

	var co Comment
	co.ID = r.CommentID

	if err = tx.Model(&co).UpdateColumn("replies_counts", gorm.Expr("replies_counts - ?", 1)).Error; err != nil {
		log.Print(err)
	}

	return
}
