package categorycontroller

import (
	"blanja_api/src/helper"
	"blanja_api/src/middleware"
	models "blanja_api/src/models/CategoryModel"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func Data_categories(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == http.MethodGet {
		var page, limit int

		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")

		if pageStr != "" {
			page, _ = strconv.Atoi(pageStr)
		}

		if limitStr != "" {
			limit, _ = strconv.Atoi(limitStr)
		}

		offset := (page - 1) * limit

		sort := r.URL.Query().Get("sort")
		if sort == "" {
			sort = "ASC"
		}
		sortBy := r.URL.Query().Get("sortBy")
		if sortBy == "" {
			sortBy = "name"
		}
		sort = sortBy + " " + strings.ToLower(sort)
		response := models.FindCond(sort, limit, offset)
		totalData := models.CountData()
		totalPage := math.Ceil(float64(totalData) / float64(limit))

		result := map[string]interface{}{
			"status":      "Success",
			"data":        response.Value,
			"currentPage": page,
			"limit":       limit,
			"totalData":   totalData,
			"totalPage":   totalPage,
		}

		res, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		return
	} else if r.Method == "POST" {
		var category models.Category
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item := models.Category{
			Name: category.Name,
		}
		models.PostCategory(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Category Created",
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

func Data_category(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/category/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectCategoryById(id).Value)
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
		var updateCategory models.Category
		err := json.NewDecoder(r.Body).Decode(&updateCategory)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newCategory := models.Category{
			Name: updateCategory.Name,
		}
		models.UpdateCategory(id, &newCategory)
		msg := map[string]string{
			"Message": "Category Updated",
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
		models.DeleteCategory(id)
		msg := map[string]string{
			"Message": "Category Deleted",
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

func Search_category(w http.ResponseWriter, r *http.Request) {
	keyWord := r.URL.Query().Get("search")
	res, err := json.Marshal(models.FindData(keyWord).Value)
	if err != nil {
		http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
