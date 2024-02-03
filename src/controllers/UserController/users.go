package usercontroller

import (
	"blanja_api/src/helper"
	"blanja_api/src/middleware"
	models "blanja_api/src/models/UserModel"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func SellerRegister(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "POST" {
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		item := models.User{
			Name:        input.Name,
			Email:       input.Email,
			Phonenumber: input.Phonenumber,
			Storename:   input.Storename,
			Password:    Password,
			Role:        "Seller",
		}
		models.PostUser(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Seller Account Created",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func CustomerRegister(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "POST" {
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		item := models.User{
			Name:        input.Name,
			Email:       input.Email,
			Phonenumber: "-",
			Storename:   "-",
			Password:    Password,
			Role:        "Customer",
		}
		models.PostUser(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Customer Account Created",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)

	if r.Method == "POST" {
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}

		users := models.FindEmail(&input)
		if len(users) == 0 {
			fmt.Fprintf(w, "Email not Found")
			return
		}

		user := users[0]
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			fmt.Fprintf(w, "Password not Found")
			return
		}

		jwtKey := os.Getenv("SECRETKEY")
		token, err := helper.GenerateToken(jwtKey, input.Email, input.Role)
		if err != nil {
			http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
			return
		}

		item := map[string]string{
			"Message": "HI, " + user.Name + " as a " + user.Role,
			"Email":   input.Email,
			"Role":   user.Role,
			"Token":   token,
		}

		result, _ := json.Marshal(item)
		if _, err := w.Write(result); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
		return
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_users(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectAllUser().Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		return
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_user(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/user/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectUserById(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		return
	} else if r.Method == "DELETE" {
		models.DeleteUser(id)
		msg := map[string]string{
			"Message": "User Deleted",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Update_seller(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/update-seller/"):]

	if r.Method == "PUT" {
		var updateSeller models.User
		err := json.NewDecoder(r.Body).Decode(&updateSeller)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(updateSeller.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		newSeller := models.User{
			Name:        updateSeller.Name,
			Email:       updateSeller.Email,
			Phonenumber: updateSeller.Phonenumber,
			Storename:   updateSeller.Storename,
			Password:    Password,
		}
		models.UpdateSeller(id, &newSeller)
		msg := map[string]string{
			"Message": "Seller Updated",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Update_customer(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/update-customer/"):]

	if r.Method == "PUT" {
		var updateCustomer models.User
		err := json.NewDecoder(r.Body).Decode(&updateCustomer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(updateCustomer.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		newCustomer := models.User{
			Name:     updateCustomer.Name,
			Email:    updateCustomer.Email,
			Password: Password,
		}
		models.UpdateCustomer(id, &newCustomer)
		msg := map[string]string{
			"Message": "Customer Updated",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := middleware.ExtractToken(r)

	if refreshToken == "" {
		http.Error(w, "Refresh token tidak tersedia", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRETKEY")), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	email := claims["email"].(string)
	role := claims["role"].(string)

	newToken, err := helper.GenerateToken(os.Getenv("SECRETKEY"), email, role)
	if err != nil {
		http.Error(w, "Failed to generate new token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"token": newToken,
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
