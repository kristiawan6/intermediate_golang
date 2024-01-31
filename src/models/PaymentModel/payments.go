package paymentmodel

import (
	"blanja_api/src/config"

	"github.com/jinzhu/gorm"
)

type Payment struct {
	gorm.Model
	Method string
	OrderId int
}

func SelectAllPayment() *gorm.DB {
	items := []Payment{}
	return config.DB.Find(&items)
}

func SelectPaymentById(id string) *gorm.DB {
	var item Payment
	return config.DB.First(&item, "id = ?", id)
}

func PostPayment(item *Payment) *gorm.DB {
	return config.DB.Create(&item)
}

func UpdatePayment(id string, newPayment *Payment) *gorm.DB {
	var item Payment
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newPayment)
}

func DeletePayment(id string) *gorm.DB {
	var item Payment
	return config.DB.Delete(&item, "id = ?", id)
}
