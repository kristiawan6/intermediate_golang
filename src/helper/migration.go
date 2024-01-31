package helper

import (
	"blanja_api/src/config"
	cartmodel "blanja_api/src/models/CartModel"
	categorymodel "blanja_api/src/models/CategoryModel"
	ordermodel "blanja_api/src/models/OrderModel"
	paymentmodel "blanja_api/src/models/PaymentModel"
	productmodel "blanja_api/src/models/ProductModel"
	usermodel "blanja_api/src/models/UserModel"
)

func Migration() {
	config.DB.AutoMigrate(&productmodel.Product{})
	config.DB.AutoMigrate(&categorymodel.Category{})
	config.DB.AutoMigrate(&cartmodel.Cart{})
	config.DB.AutoMigrate(&ordermodel.Order{})
	config.DB.AutoMigrate(&paymentmodel.Payment{})
	config.DB.AutoMigrate(&usermodel.User{})
}
