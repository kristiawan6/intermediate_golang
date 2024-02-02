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

func SelectAllCartPaginated(page, limit int) ([]Cart, int, error) {
	var items []Cart
	var totalCount int

	config.DB.Model(&Cart{}).Count(&totalCount)

	offset := (page - 1) * limit

	err := config.DB.Offset(offset).Limit(limit).Find(&items).Error
	if err != nil {
		return nil, 0, err
	}

	return items, totalCount, nil
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

func CountData() int {
	var item int
	config.DB.Table("carts").Count(&item)
	return item
}
