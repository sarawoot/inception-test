package model

import "time"

type PaymentTransaction struct {
	ID         uint64    `gorm:"primaryKey,type:autoIncrement"`
	ChargeID   string    `gorm:"type:varchar(255)"`
	SourceID   string    `gorm:"type:varchar(255)"`
	Amount     int64     ``
	RawRespone string    `gorm:"type:text"`
	Status     string    `gorm:"type:varchar(20)"`
	CreatedAt  time.Time ``
	UpdatedAt  time.Time ``
}
