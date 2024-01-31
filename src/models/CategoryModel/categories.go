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
