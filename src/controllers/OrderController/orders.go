package ordercontroller

import (
	"blanja_api/src/helper"
	"blanja_api/src/middleware"
	models "blanja_api/src/models/OrderModel"
	"encoding/json"
	"net/http"
)

func Data_orders(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectAllOrder().Value)
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
	} else if r.Method == "POST" {
		var order models.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item := models.Order{
			Address:     order.Address,
			Amount:      order.Amount,
			Deliveryfee: order.Deliveryfee,
			Total:       order.Total,
			UserId:      order.UserId,
			ProductId:   order.ProductId,
			PaymentId:   order.PaymentId,
		}
		models.PostOrder(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Your order has been sucessful",
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

func Data_order(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/order/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectOrderById(id).Value)
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
		var updateOrder models.Order
		err := json.NewDecoder(r.Body).Decode(&updateOrder)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newOrder := models.Order{
			Address:     updateOrder.Address,
			Amount:      updateOrder.Amount,
			Deliveryfee: updateOrder.Deliveryfee,
			Total:       updateOrder.Total,
			ProductId:   updateOrder.ProductId,
			PaymentId:   updateOrder.PaymentId,
			UserId:      updateOrder.UserId,
		}
		models.UpdateOrder(id, &newOrder)
		msg := map[string]string{
			"Message": "Order Updated",
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
		models.DeleteOrder(id)
		msg := map[string]string{
			"Message": "Order Deleted",
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
