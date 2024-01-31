package productmodel

import (
	"blanja_api/src/config"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Price       uint
	Stock       int
	Description string
	Condition   string
	Size        string
	UserId      int
	CategoryId  int
}

func SelectAllProduct() *gorm.DB {
	items := []Product{}
	return config.DB.Find(&items)
}

func SelectProductById(id string) *gorm.DB {
	var item Product
	return config.DB.First(&item, "id = ?", id)
}

func PostProduct(item *Product) *gorm.DB {
	return config.DB.Create(&item)
}

func UpdatesProduct(id string, newProduct *Product) *gorm.DB {
	var item Product
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
}

func DeletesProduct(id string) *gorm.DB {
	var item Product
	return config.DB.Delete(&item, "id = ?", id)
}
