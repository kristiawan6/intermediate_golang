package routes

import (
	cartcontroller "blanja_api/src/controllers/CartController"
	categorycontroller "blanja_api/src/controllers/CategoryController"
	ordercontroller "blanja_api/src/controllers/OrderController"
	paymentcontroller "blanja_api/src/controllers/PaymentController"
	productcontroller "blanja_api/src/controllers/ProductController"
	usercontroller "blanja_api/src/controllers/UserController"
	"blanja_api/src/middleware"
	"fmt"
	"net/http"

	"github.com/goddtriffin/helmet"
)

func Router() {

	helmet := helmet.Default()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Nothing Here, Try Another Page")
	})

	//Routes User
	//login and register
	http.Handle("/login", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.Login))))
	http.Handle("/register-seller", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.SellerRegister))))
	http.Handle("/register-customer", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.CustomerRegister))))

	/* --------------------------------------------- XSS MIDDLEWARE ROUTES --------------------------------------------*/

	//Routes User

	// http.Handle("/update-seller/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.Update_seller))))
	// http.Handle("/update-customer/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.Update_customer))))

	// http.Handle("/users", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.Data_users))))
	// http.Handle("/user/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.Data_user))))

	// //Route Refresh Token
	// http.Handle("/refresh-token", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontroller.RefreshToken))))

	// //Route Upload
	// http.Handle("/upload", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(productcontroller.HandleUpload))))
	// http.Handle("/upload", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(productcontroller.HandleUpload))))

	// //Routes Product
	http.Handle("/products", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(productcontroller.DataProducts))))
	http.Handle("/product/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(productcontroller.DataProduct))))

	// //Routes Category
	http.Handle("/categories", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(categorycontroller.Data_categories))))
	http.Handle("/category/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(categorycontroller.Data_category))))

	// //Routes Cart
	// http.Handle("/carts", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(cartcontroller.Data_carts))))
	// http.Handle("/cart/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(cartcontroller.Data_cart))))

	// //Routes Order
	// http.Handle("/orders", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(ordercontroller.Data_orders))))
	// http.Handle("/order/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(ordercontroller.Data_order))))

	// //Routes Payment
	// http.Handle("/payments", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(paymentcontroller.Data_payments))))
	// http.Handle("/payment/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(paymentcontroller.Data_payment))))

	/* --------------------------------------------- JWT MIDDLEWARE ROUTES --------------------------------------------*/

	//Routes User
	http.Handle("/update-seller/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(usercontroller.Update_seller))))
	http.Handle("/update-customer/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(usercontroller.Update_customer))))
	
	http.Handle("/users", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(usercontroller.Data_users))))
	http.Handle("/user/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(usercontroller.Data_user))))

	//Route Refresh Token
	http.Handle("/refresh-token", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(usercontroller.RefreshToken))))

	//Route Upload
	http.Handle("/upload", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(productcontroller.HandleUpload))))

	//Routes Product
	// http.Handle("/products",(http.HandlerFunc(usercontroller.RefreshToken)))
	// http.Handle("/add_product", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(productcontroller.Add_products))))
	// http.Handle("/add_product", helmet.Secure(middleware.JwtMiddleware("Seller")(http.HandlerFunc(productcontroller.Add_products))))
	// http.Handle("/product/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(productcontroller.Data_product))))
	// http.Handle("/search-product", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(productcontroller.Search_product))))

	//Routes Category
	// http.Handle("/categories", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(categorycontroller.Data_categories))))
	// http.Handle("/category/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(categorycontroller.Data_category))))
	// http.Handle("/search-category", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(categorycontroller.Search_category))))

	//Routes Cart
	http.Handle("/carts", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(cartcontroller.Data_carts))))
	http.Handle("/cart/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(cartcontroller.Data_cart))))
	http.Handle("/search-cart", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(cartcontroller.Search_cart))))

	//Routes Order
	http.Handle("/orders", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(ordercontroller.Data_orders))))
	http.Handle("/order/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(ordercontroller.Data_order))))

	//Routes Payment
	http.Handle("/payments", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(paymentcontroller.Data_payments))))
	http.Handle("/payment/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(paymentcontroller.Data_payment))))

}
