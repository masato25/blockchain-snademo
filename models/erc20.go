package models

import (
	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);unique_index:token_name"`
	Symbol string `gorm:"type:varchar(100);unique_index:token_symbol"`
	Unit   string `gorm:"type:varchar(100)"`
	Amount int
}

func InitToken(qdb *gorm.DB) error {
	d := Token{Name: "medison1", Symbol: "md1", Amount: 1000, Unit: "mg"}
	qdb.Create(&d)
	return qdb.Error
}
