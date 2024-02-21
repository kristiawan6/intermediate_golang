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

func FindData(keyword string) *gorm.DB {
	items := []Product{}
	keyword = "%" + keyword + "%"
	return config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(price AS TEXT) LIKE ? OR CAST(stock AS TEXT) LIKE ? OR description LIKE ? OR condition LIKE ? OR size LIKE ? ", keyword, keyword, keyword, keyword, keyword, keyword, keyword).Find(&items)
}

func FindCond(sort string, limit int, offset int) *gorm.DB {
	item := []Product{}
	return config.DB.Order(sort).Limit(limit).Offset(offset).Find(&item)
}

func CountData() int {
	var item int
	config.DB.Table("products").Where("deleted_at IS NULL").Count(&item)
	return item
}
