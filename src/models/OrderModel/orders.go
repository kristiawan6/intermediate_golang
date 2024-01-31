package ordermodel

import (
	"blanja_api/src/config"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Address     string
	Amount      uint
	Deliveryfee uint
	Total       uint
	UserId      int
	ProductId   int
	PaymentId   int
}

func SelectAllOrder() *gorm.DB {
	items := []Order{}
	return config.DB.Find(&items)
}

func SelectOrderById(id string) *gorm.DB {
	var item Order
	return config.DB.First(&item, "id = ?", id)
}

func PostOrder(item *Order) *gorm.DB {
	return config.DB.Create(&item)
}

func UpdateOrder(id string, newOrder *Order) *gorm.DB {
	var item Order
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newOrder)
}

func DeleteOrder(id string) *gorm.DB {
	var item Order
	return config.DB.Delete(&item, "id = ?", id)
}
