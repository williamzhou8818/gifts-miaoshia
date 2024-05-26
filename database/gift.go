package database

import (
	"github.com/williamzhou8818/gifts-miaoshia/util"
	"gorm.io/gorm"
)

type Gift struct {
	Id      int    `gorm:"column:id;primaryKey"`
	Name    string `gorm:"column:name"`
	Price   int    `gorm:"column:price"`
	Picture string `gorm:"column:picture"`
	Count   int    `gorm:"column:count"`
}

func (Gift) TableName() string {
	return "inventory"
}

var (
	_all_gitf_field = util.GetGormFields(Gift{})
)

func GetAllGiftsV1() []*Gift {
	db := GetGiftDBConnection()
	var gifts []*Gift
	err := db.Select(_all_gitf_field).Find(&gifts).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			util.LogRus.Errorf("read table %s failed: %s", Gift{}.TableName(), err)
		}
	}
	return gifts
}
