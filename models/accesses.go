package models

import "time"

type Accesses struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement"`
	Ip          string    `gorm:"column:ip;type:string;size:255"`
	OriginID    uint      `gorm:"column:origin_id;type:int"`
	ChangedDate time.Time `gorm:"column:changed_date;type:time"`
}

func (Accesses) TableName() string {
	return "accesses"
}
