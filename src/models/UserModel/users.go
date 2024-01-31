package usermodel

import (
	"blanja_api/src/config"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Email string
	Phonenumber string
	Storename string
	Password string
	Role string
}

func SelectAllUser() *gorm.DB {
	items := []User{}
	return config.DB.Find(&items)
}

func SelectUserById(id string) *gorm.DB {
	var item User
	return config.DB.First(&item, "id = ?", id)
}

func PostUser(item *User) *gorm.DB {
	return config.DB.Create(&item)
}

func UpdateCustomer(id string, newCustomer *User) *gorm.DB {
	var item User
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newCustomer)
}

func UpdateSeller(id string, newSeller *User) *gorm.DB {
	var item User
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newSeller)
}

func DeleteUser(id string) *gorm.DB {
	var item User
	return config.DB.Delete(&item, "id = ?", id)
}

func FindEmail(input *User) []User {
	items := []User{}
	config.DB.Raw("SELECT * FROM users WHERE email = ?", input.Email).Scan(&items)
	return items
}
