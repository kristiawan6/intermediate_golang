package cartcontroller

import (
	"blanja_api/src/helper"
	"blanja_api/src/middleware"
	models "blanja_api/src/models/CartModel"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

func Data_carts(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == http.MethodGet {
		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")

		page := 1
		limit := 10

		if pageStr != "" {
			page, _ = strconv.Atoi(pageStr)
		}

		if limitStr != "" {
			limit, _ = strconv.Atoi(limitStr)
		}

		items, totalCount, err := models.SelectAllCartPaginated(page, limit)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"status":      "Success",
			"data":        items,
			"currentPage": page,
			"limit":       limit,
			"totalData":   totalCount,
			"totalPage":   int(math.Ceil(float64(totalCount) / float64(limit))),
		}

		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return

	} else if r.Method == "POST" {
		var cart models.Cart
		err := json.NewDecoder(r.Body).Decode(&cart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item := models.Cart{
			Quantity:  cart.Quantity,
			Amount:    cart.Amount,
			ProductId: cart.ProductId,
		}
		models.PostCart(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Cart Created",
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

func Data_cart(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/cart/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectCartById(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		return
	} else if r.Method == "PUT" {
		var updateCart models.Cart
		err := json.NewDecoder(r.Body).Decode(&updateCart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newCart := models.Cart{
			Quantity:  updateCart.Quantity,
			Amount:    updateCart.Amount,
			ProductId: updateCart.ProductId,
		}
		models.UpdateCart(id, &newCart)
		msg := map[string]string{
			"Message": "Cart Updated",
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
	} else if r.Method == "DELETE" {
		models.DeleteCart(id)
		msg := map[string]string{
			"Message": "Cart Deleted",
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
