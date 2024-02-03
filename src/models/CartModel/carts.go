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

func FindData(id string) *gorm.DB {
	items := []Cart{}
	id = "%" + id + "%"
	return config.DB.Where("id LIKE ?", id).Find(&items)
}

func FindCond(sort string,limit int, offset int) *gorm.DB {
	item := []Cart{}
	return config.DB.Order(sort).Limit(limit).Offset(offset).Find(&item)
}

func CountData() int {
    var item int
    config.DB.Table("carts").Where("deleted_at IS NULL").Count(&item)
    return item
}