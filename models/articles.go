package models

import "time"

type Articles struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement"`
	UserID      uint      `gorm:"column:user_id;type:int"`
	Title       string    `gorm:"column:first_name;type:string;size:255"`
	ContentText string    `gorm:"column:last_name;type:string"`
	ChangedDate time.Time `gorm:"column:changed_date;type:time"`
}

func (Articles) TableName() string {
	return "articles"
}
