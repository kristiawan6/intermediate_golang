package categorymodel

import (
	"blanja_api/src/config"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string
}

func SelectAllCategory() *gorm.DB {
	items := []Category{}
	return config.DB.Find(&items)
}

func SelectCategoryById(id string) *gorm.DB {
	var item Category
	return config.DB.First(&item, "id = ?", id)
}

func PostCategory(item *Category) *gorm.DB {
	return config.DB.Create(&item)
}

func UpdateCategory(id string, newCategory *Category) *gorm.DB {
	var item Category
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newCategory)
}

func DeleteCategory(id string) *gorm.DB {
	var item Category
	return config.DB.Delete(&item, "id = ?", id)
}

func FindData(name string) *gorm.DB {
	items := []Category{}
	name = "%" + name + "%"
	return config.DB.Where("name LIKE ?", name).Find(&items)
}

func FindCond(sort string,limit int, offset int) *gorm.DB {
	item := []Category{}
	return config.DB.Order(sort).Limit(limit).Offset(offset).Find(&item)
}

func CountData() int {
    var item int
    config.DB.Table("categories").Where("deleted_at IS NULL").Count(&item)
    return item
}