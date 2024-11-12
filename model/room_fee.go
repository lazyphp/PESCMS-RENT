package model

import (
	"math"

	"gorm.io/gorm"
)

type Room_fee struct {
	Room_fee_id     int     `gorm:"primary_key;column:fee_id" json:"fee_id"`
	Room_id         int     `gorm:"column:room_id" json:"room_id"`
	Fee_name        string  `gorm:"column:fee_name" json:"fee_name"`
	Fee_value       float64 `gorm:"column:fee_value" json:"fee_value"`
	Fee_price       float64 `gorm:"column:fee_price" json:"fee_price"`
	Fee_used        float64 `gorm:"column:fee_used" json:"fee_used"`
	Fee_date        string  `gorm:"column:fee_date" json:"fee_date"`
	Fee_create_time int     `gorm:"column:fee_create_time" json:"fee_create_time"`
}

func (rf *Room_fee) AfterFind(tx *gorm.DB) (err error) {
	rf.Fee_price = math.Round(rf.Fee_price*100) / 100
	return
}
