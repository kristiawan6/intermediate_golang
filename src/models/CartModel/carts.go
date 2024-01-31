package cartmodel

import (
	"blanja_api/src/config"

	"github.com/jinzhu/gorm"
)

type Cart struct {
	gorm.Model
	Quantity int
	Amount uint
	ProductId int
}

func SelectAllCart() *gorm.DB {
	items := []Cart{}
	return config.DB.Find(&items)
}

func SelectCartById(id string) *gorm.DB {
	var item Cart
	return config.DB.First(&item, "id = ?", id)
}

func PostCart(item *Cart) *gorm.DB {
	return config.DB.Create(&item)
}

func UpdateCart(id string, newCart *Cart) *gorm.DB {
	var item Cart
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newCart)
}

func DeleteCart(id string) *gorm.DB {
	var item Cart
	return config.DB.Delete(&item, "id = ?", id)
}
