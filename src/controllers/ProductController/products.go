package productcontroller

import (
	"blanja_api/src/helper"
	"blanja_api/src/middleware"
	models "blanja_api/src/models/ProductModel"
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"
)

func Data_products(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectAllProduct().Value)
		if err != nil {
			http.Error(w, "Gagal konversi ke JSON", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		return
	} else {
		http.Error(w, "Method tidak Diizinkan", http.StatusMethodNotAllowed)
	}
}

func Add_products(w http.ResponseWriter, r *http.Request) {
	helper.EnableCors(w)
	if r.Method == "POST" {
		var product models.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		item := models.Product{
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Description: product.Description,
			Condition:   product.Condition,
			Size:        product.Size,
			UserId:      product.UserId,
			CategoryId:  product.CategoryId,
		}
		models.PostProduct(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Product Created",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal konversi ke JSON", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method tidak Diizinkan", http.StatusMethodNotAllowed)
	}

}

func Data_product(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/product/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectProductById(id).Value)
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
	} else if r.Method == "PUT" {
		var updateProduct models.Product
		err := json.NewDecoder(r.Body).Decode(&updateProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newProduct := models.Product{
			Name:        updateProduct.Name,
			Price:       updateProduct.Price,
			Stock:       updateProduct.Stock,
			Description: updateProduct.Description,
			Condition:   updateProduct.Condition,
			Size:        updateProduct.Size,
			CategoryId:  updateProduct.CategoryId,
		}
		models.UpdatesProduct(id, &newProduct)
		msg := map[string]string{
			"Message": "Product Updated",
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
		models.DeletesProduct(id)
		msg := map[string]string{
			"Message": "Product Deleted",
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

func HandleUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	const (
		AllowedExtensions = ".jpg,.jpeg,.pdf,.png"
		MaxFileSize       = 2 << 20 // 2 MB
	)

	file, handler, err := r.FormFile("File")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := filepath.Ext(handler.Filename)
	ext = strings.ToLower(ext)
	allowedExts := strings.Split(AllowedExtensions, ",")
	validExtension := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			validExtension = true
			break
		}
	}
	if !validExtension {
		http.Error(w, "Invalid file extension", http.StatusBadRequest)
		return
	}

	fileSize := handler.Size
	if fileSize > MaxFileSize {
		http.Error(w, "File size exceeds the allowed limit", http.StatusBadRequest)
		return
	}
	helper.Uplaod(w, file, handler)

	msg := map[string]string{
		"Message": "File uploaded successfully",
	}
	res, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
