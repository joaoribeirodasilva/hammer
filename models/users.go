package models

import "time"

type Users struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement"`
	FirstName   string    `gorm:"column:first_name;type:string;size:255"`
	LastName    string    `gorm:"column:last_name;type:string;size:255"`
	Email       string    `gorm:"column:email;type:string;size:255"`
	Password    string    `gorm:"column:password;type:string;size:255"`
	ChangedDate time.Time `gorm:"->;column:changed_date;type:time"`
}

func (Users) TableName() string {
	return "users"
}
