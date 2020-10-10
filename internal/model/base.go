package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// GORMBase struct
type GORMBase struct {
	ID        int   `json:"id" gorm:"AUTO_INCREMENT"`
	CreatedAt int64 `json:"createAt"`
	UpdatedAt int64 `json:"updateAt"`
}

// BeforeCreate func
func (m *GORMBase) BeforeCreate(scope *gorm.Scope) error {
	if m.UpdatedAt == 0 {
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}

	scope.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}

// BeforeUpdate func
func (m *GORMBase) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
