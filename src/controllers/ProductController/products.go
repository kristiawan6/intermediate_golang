package productcontroller

import (
    "blanja_api/src/helper"
    "blanja_api/src/middleware"
    models "blanja_api/src/models/ProductModel"
    "encoding/json"
    "fmt"
    "math"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"

    "github.com/dgrijalva/jwt-go"
  
)

func DataProducts(w http.ResponseWriter, r *http.Request) {
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
            "data":        response,
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

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(res)
        return

    } else {
        http.Error(w, "Method tidak Diizinkan", http.StatusMethodNotAllowed)
    }
}

func AddProducts(w http.ResponseWriter, r *http.Request) {
    helper.EnableCors(w)
    if r.Method == "POST" {
        tokenString := middleware.ExtractToken(r)

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Check the signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            secretKey := []byte(os.Getenv("SECRETKEY"))
            return secretKey, nil
        })

        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            if role, ok := claims["role"].(string); ok && role == "Seller" {
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
                    http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
                    return
                }

                if _, err := w.Write(res); err != nil {
                    http.Error(w, "Failed to write response", http.StatusInternalServerError)
                    return
                }
            } else {
                http.Error(w, "User does not have the necessary role for creating a product", http.StatusForbidden)
                return
            }

        } else {
            http.Error(w, "User does not have the necessary role for creating a product", http.StatusForbidden)
            return
        }
    }
}

func DataProduct(w http.ResponseWriter, r *http.Request) {
    middleware.GetCleanedInput(r)
    helper.EnableCors(w)
    id := r.URL.Path[len("/product/"):]

    if r.Method == "GET" {
        res, err := json.Marshal(models.SelectProductById(id))
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
            http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
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
            http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
            return
        }
        if _, err := w.Write(res); err != nil {
            http.Error(w, "Failed to write response", http.StatusInternalServerError)
            return
        }
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
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
        http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
        return
    }
    w.Write(res)
}

func SearchProduct(w http.ResponseWriter, r *http.Request) {
    keyWord := r.URL.Query().Get("search")
    res, err := json.Marshal(models.FindData(keyWord))
    if err != nil {
        http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
        return
    }
    w.Write(res)
}
