package productmodel

import (
	"blanja_api/src/config"

	"gorm.io/gorm"
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
	CategoryId  uint
	Category    Category `gorm:"foreignKey:CategoryId"`
}

type Category struct {
	gorm.Model
	Name string
}

func SelectAllProduct() []*Product {
	items := []*Product{}
	config.DB.Find(&items)
	return items
}

func SelectProductById(id string) *Product {
	var item Product
	if err := config.DB.Preload("Category").First(&item, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return nil
	}
	return &item
}

func PostProduct(item *Product) *Product {
	config.DB.Create(&item)
	return item
}

func UpdatesProduct(id string, newProduct *Product) *Product {
	var item Product
	config.DB.Model(&item).Where("id = ?", id).Updates(newProduct)
	return &item
}

func DeletesProduct(id string) {
	var item Product
	config.DB.Delete(&item, "id = ?", id)
}

func FindData(keyword string) []*Product {
	items := []*Product{}
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(price AS TEXT) LIKE ? OR CAST(stock AS TEXT) LIKE ? OR description LIKE ? OR condition LIKE ? OR size LIKE ? ", keyword, keyword, keyword, keyword, keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Product {
	item := []*Product{}
	config.DB.Order(sort).Limit(limit).Offset(offset).Preload("Category").Find(&item)
	return item
}

func CountData() int64 {
	var item int64
	config.DB.Table("products").Where("deleted_at IS NULL").Count(&item)
	return item
}
