package product

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// CategoryResponse represents the response format for categories
type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// BrandResponse represents the response format for brands
type BrandResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ProductResponse represents the response format for products
type ProductResponse struct {
	ID           uint              `json:"id"`
	Name         string            `json:"name"`
	Description  *string           `json:"description,omitempty"`
	Slug         string            `json:"slug"`
	Price        float64           `json:"price"`
	CategoryID   *uint             `json:"category_id,omitempty"`
	CategorySlug *string           `json:"category_slug,omitempty"`
	BrandID      *uint             `json:"brand_id,omitempty"`
	BrandSlug    *string           `json:"brand_slug,omitempty"`
	Category     *CategoryResponse `json:"category,omitempty"`
	Brand        *BrandResponse    `json:"brand,omitempty"`
	ViewsCount   int64             `json:"views_count"`
	CreatedAt    string            `json:"created_at"`
	UpdatedAt    string            `json:"updated_at"`
}

// ProductListResponse represents paginated product list response
type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
	Count    int               `json:"count"`
	Limit    int               `json:"limit"`
	Offset   int               `json:"offset"`
	Total    int64             `json:"total,omitempty"`
}

// convertProductToResponse converts domain entity to response format
func convertProductToResponse(product *entities.Product, categoryRepo repository.CategoryRepository, brandRepo repository.BrandRepository) ProductResponse {
	response := ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Slug:        product.Slug,
		Price:       product.Price,
		CategoryID:  product.CategoryID,
		BrandID:     product.BrandID,
		ViewsCount:  product.ViewsCount,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	// Fetch category information if category ID exists
	if product.CategoryID != nil && categoryRepo != nil {
		category, err := categoryRepo.GetByID(context.Background(), *product.CategoryID)
		if err == nil && category != nil {
			response.CategorySlug = &category.Slug
			response.Category = &CategoryResponse{
				ID:   category.ID,
				Name: category.Name,
				Slug: category.Slug,
			}
		}
	}

	// Fetch brand information if brand ID exists
	if product.BrandID != nil && brandRepo != nil {
		brand, err := brandRepo.GetByID(context.Background(), *product.BrandID)
		if err == nil && brand != nil {
			response.BrandSlug = &brand.Slug
			response.Brand = &BrandResponse{
				ID:   brand.ID,
				Name: brand.Name,
				Slug: brand.Slug,
			}
		}
	}

	return response
}

// convertProductToResponseSimple converts domain entity to response format without fetching related data
func convertProductToResponseSimple(product *entities.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Slug:        product.Slug,
		Price:       product.Price,
		CategoryID:  product.CategoryID,
		BrandID:     product.BrandID,
		ViewsCount:  product.ViewsCount,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// GetProductByIDHandler handles GET /products/{id}
func GetProductByIDHandler(w http.ResponseWriter, r *http.Request, repo repository.ProductRepository, categoryRepo repository.CategoryRepository, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL path - handle both /products/{id} and /products/{id}/
	path := strings.TrimPrefix(r.URL.Path, "/products/")
	path = strings.TrimSuffix(path, "/")

	// Split by slash and get the first part as ID
	pathParts := strings.Split(path, "/")
	if len(pathParts) == 0 || pathParts[0] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Product ID is required"})
		return
	}

	idStr := pathParts[0]
	productID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid product ID format",
			"received": idStr,
		})
		return
	}

	product, err := repo.GetByID(r.Context(), uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Product not found",
			"id":    idStr,
		})
		return
	}

	response := convertProductToResponse(product, categoryRepo, brandRepo)
	json.NewEncoder(w).Encode(response)
}

// GetProductBySlugHandler handles GET /products-by-slug/{slug}
func GetProductBySlugHandler(w http.ResponseWriter, r *http.Request, repo repository.ProductRepository, categoryRepo repository.CategoryRepository, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract slug from URL path
	path := strings.TrimPrefix(r.URL.Path, "/products-by-slug/")
	path = strings.TrimSuffix(path, "/")

	if path == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Product slug is required"})
		return
	}

	product, err := repo.GetBySlug(r.Context(), path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Product not found",
			"slug":  path,
		})
		return
	}

	response := convertProductToResponse(product, categoryRepo, brandRepo)
	json.NewEncoder(w).Encode(response)
}

// CreateProductHandler handles POST /products
func CreateProductHandler(w http.ResponseWriter, r *http.Request, repo repository.ProductRepository) {
	w.Header().Set("Content-Type", "application/json")

	var req entities.Product
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	if req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Product name is required"})
		return
	}

	product, err := repo.Create(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := convertProductToResponseSimple(product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// UpdateProductHandler handles PATCH /products/{id}
func UpdateProductHandler(w http.ResponseWriter, r *http.Request, repo repository.ProductRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/products/")
	path = strings.TrimSuffix(path, "/")

	pathParts := strings.Split(path, "/")
	if len(pathParts) == 0 || pathParts[0] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Product ID is required"})
		return
	}

	idStr := pathParts[0]
	productID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid product ID format",
			"received": idStr,
		})
		return
	}

	var req entities.Product
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	product, err := repo.Update(r.Context(), uint(productID), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := convertProductToResponseSimple(product)
	json.NewEncoder(w).Encode(response)
}

// ListProductsHandler handles GET /products
func ListProductsHandler(w http.ResponseWriter, r *http.Request, repo repository.ProductRepository, categoryRepo repository.CategoryRepository, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	searchQuery := r.URL.Query().Get("search")
	categoryIDStr := r.URL.Query().Get("category_id")
	brandIDStr := r.URL.Query().Get("brand_id")

	// Set defaults
	limit := 10
	offset := 0

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	var products []*entities.Product
	var err error

	// Handle different query types
	if searchQuery != "" {
		products, err = repo.Search(r.Context(), searchQuery, limit, offset)
	} else if categoryIDStr != "" {
		categoryID, parseErr := strconv.ParseUint(categoryIDStr, 10, 32)
		if parseErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid category ID"})
			return
		}
		products, err = repo.GetByCategory(r.Context(), uint(categoryID), limit, offset)
	} else if brandIDStr != "" {
		brandID, parseErr := strconv.ParseUint(brandIDStr, 10, 32)
		if parseErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid brand ID"})
			return
		}
		products, err = repo.GetByBrand(r.Context(), uint(brandID), limit, offset)
	} else {
		products, err = repo.List(r.Context(), limit, offset)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Convert products to response format
	productResponses := make([]ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = convertProductToResponse(product, categoryRepo, brandRepo)
	}

	// Get total count
	total, _ := repo.Count(r.Context())

	response := ProductListResponse{
		Products: productResponses,
		Count:    len(productResponses),
		Limit:    limit,
		Offset:   offset,
		Total:    total,
	}

	json.NewEncoder(w).Encode(response)
}

// GetFilteredProductsHandler handles Laravel-compatible filtered product listing
// GET /products?locale=en&page=1&limit=10&category=&brand=&priceRange=&searchterm=&sortby=
func GetFilteredProductsHandler(w http.ResponseWriter, r *http.Request, repo repository.ProductRepository, categoryRepo repository.CategoryRepository, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters with Laravel compatibility
	locale := r.URL.Query().Get("locale")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	category := r.URL.Query().Get("category")
	brandParam := r.URL.Query().Get("brand")
	priceRange := r.URL.Query().Get("priceRange")
	searchterm := r.URL.Query().Get("searchterm")
	sortby := r.URL.Query().Get("sortby")

	// Set defaults
	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	// Convert brand parameter to array (Laravel sends comma-separated)
	var brands []string
	if brandParam != "" {
		brands = strings.Split(brandParam, ",")
		// Trim whitespace from each brand
		for i, brand := range brands {
			brands[i] = strings.TrimSpace(brand)
		}
	}

	// Create filters struct
	filters := &repository.ProductFilters{
		Page:       page,
		Limit:      limit,
		Locale:     locale,
		SearchTerm: searchterm,
		Category:   category,
		Brand:      brandParam, // Store original comma-separated string
		BrandSlugs: brands,     // Store as array for filtering
		PriceRange: priceRange,
		SortBy:     sortby,
	}

	// Get filtered products
	products, totalCount, err := repo.GetWithFilters(r.Context(), filters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Convert products to response format
	productResponses := make([]ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = convertProductToResponse(product, categoryRepo, brandRepo)
	}

	// Calculate pagination info
	totalPages := (totalCount + int64(limit) - 1) / int64(limit)
	hasNextPage := page < int(totalPages)
	hasPrevPage := page > 1

	// Laravel-compatible response format
	response := map[string]interface{}{
		"data": productResponses,
		"meta": map[string]interface{}{
			"current_page":  page,
			"per_page":      limit,
			"total":         totalCount,
			"last_page":     totalPages,
			"from":          (page-1)*limit + 1,
			"to":            (page-1)*limit + len(productResponses),
			"has_next_page": hasNextPage,
			"has_prev_page": hasPrevPage,
		},
		"filters": map[string]interface{}{
			"locale":      locale,
			"category":    category,
			"brand":       brandParam,
			"price_range": priceRange,
			"search_term": searchterm,
			"sort_by":     sortby,
		},
	}

	json.NewEncoder(w).Encode(response)
}

// GetPopularProductsHandler handles GET /api/popular-products
func GetPopularProductsHandler(w http.ResponseWriter, r *http.Request, repo repository.ProductRepository, categoryRepo repository.CategoryRepository, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	limitStr := r.URL.Query().Get("limit")
	limit := 10

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	products, err := repo.GetPopular(r.Context(), limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	productResponses := make([]ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = convertProductToResponse(product, categoryRepo, brandRepo)
	}

	response := ProductListResponse{
		Products: productResponses,
		Count:    len(productResponses),
		Limit:    limit,
		Offset:   0,
	}

	json.NewEncoder(w).Encode(response)
}

// IncrementProductViewsHandler handles POST /api/products/{id}/increment-views
func IncrementProductViewsHandler(w http.ResponseWriter, r *http.Request, repo repository.ProductRepository, categoryRepo repository.CategoryRepository, brandRepo repository.BrandRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL path - expecting /api/products/{id}/increment-views
	path := strings.TrimPrefix(r.URL.Path, "/api/products/")
	pathParts := strings.Split(path, "/")

	if len(pathParts) < 2 || pathParts[0] == "" || pathParts[1] != "increment-views" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid URL format"})
		return
	}

	idStr := pathParts[0]
	productID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid product ID format",
			"received": idStr,
		})
		return
	}

	err = repo.IncrementViews(r.Context(), uint(productID))
	if err != nil {
		if err.Error() == "product not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Product not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
		return
	}

	// Return updated product
	product, err := repo.GetByID(r.Context(), uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get updated product"})
		return
	}

	response := convertProductToResponse(product, categoryRepo, brandRepo)
	json.NewEncoder(w).Encode(response)
}

// AddProductImageHandler handles POST /api/addproductimage/{productId}
func AddProductImageHandler(w http.ResponseWriter, r *http.Request, imageRepo repository.ImageRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/addproductimage/")
	productIDStr := strings.Trim(path, "/")

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid product ID format",
			"received": productIDStr,
		})
		return
	}

	// Parse multipart form
	err = r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to parse form"})
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "No image file provided"})
		return
	}
	defer file.Close()

	// Create uploads directory if it doesn't exist
	uploadDir := "uploads/products"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create upload directory"})
		return
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d_%d_%s", productID, time.Now().Unix(), handler.Filename)
	filePath := filepath.Join(uploadDir, filename)

	// Create the file
	dst, err := os.Create(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create file"})
		return
	}
	defer dst.Close()

	// Copy file content
	_, err = io.Copy(dst, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save file"})
		return
	}

	// Save image record to database
	image := &entities.Image{
		ImageableType: "product",
		ImageableID:   uint(productID),
		ImagePath:     filePath,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	savedImage, err := imageRepo.Create(r.Context(), image)
	if err != nil {
		// Clean up file if database save fails
		os.Remove(filePath)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save image record"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "Image uploaded successfully",
		"image_id":   savedImage.ID,
		"image_path": savedImage.ImagePath,
	})
}

// GetProductImageHandler handles GET /api/get-product-image/{productId}
func GetProductImageHandler(w http.ResponseWriter, r *http.Request, imageRepo repository.ImageRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/get-product-image/")
	productIDStr := strings.Trim(path, "/")

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid product ID format",
			"received": productIDStr,
		})
		return
	}

	images, err := imageRepo.GetByImageableID(r.Context(), "product", uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get images"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"product_id": productID,
		"images":     images,
	})
}

// AddProductImageAltHandler handles POST /api/products/{product}/image
func AddProductImageAltHandler(w http.ResponseWriter, r *http.Request, imageRepo repository.ImageRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/products/")
	pathParts := strings.Split(path, "/")

	if len(pathParts) < 2 || pathParts[1] != "image" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid URL format"})
		return
	}

	productIDStr := pathParts[0]
	_, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid product ID format",
			"received": productIDStr,
		})
		return
	}

	// Reconstruct the URL for AddProductImageHandler
	r.URL.Path = "/api/addproductimage/" + productIDStr
	AddProductImageHandler(w, r, imageRepo)
}

// GetProductImagesHandler handles GET /api/products/{product}/images
func GetProductImagesHandler(w http.ResponseWriter, r *http.Request, imageRepo repository.ImageRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/products/")
	pathParts := strings.Split(path, "/")

	if len(pathParts) < 2 || pathParts[1] != "images" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid URL format"})
		return
	}

	productIDStr := pathParts[0]
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid product ID format",
			"received": productIDStr,
		})
		return
	}

	images, err := imageRepo.GetByImageableID(r.Context(), "product", uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get images"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"product_id": productID,
		"images":     images,
	})
}

// CreateProductTranslationHandler handles POST /api/product-trans/{id}
func CreateProductTranslationHandler(w http.ResponseWriter, r *http.Request, productRepo repository.ProductRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/product-trans/")
	productIDStr := strings.Trim(path, "/")

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid product ID format",
			"received": productIDStr,
		})
		return
	}

	var request struct {
		Locale                string  `json:"locale"`
		TranslatedName        string  `json:"translated_name"`
		TranslatedDescription *string `json:"translated_description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.Locale == "" || request.TranslatedName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Locale and translated_name are required"})
		return
	}

	translation := &entities.ProductTranslation{
		ProductID:             uint(productID),
		Locale:                request.Locale,
		TranslatedName:        request.TranslatedName,
		TranslatedDescription: request.TranslatedDescription,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}

	savedTranslation, err := productRepo.CreateTranslation(r.Context(), translation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create translation"})
		return
	}

	json.NewEncoder(w).Encode(savedTranslation)
}

// GetPublicReviewsHandler handles GET /public-reviews/{id}
// This endpoint matches the Laravel API: Route::get('/public-reviews/{id}', [ProductReviewController::class, 'getPublicReviews']);
func GetPublicReviewsHandler(w http.ResponseWriter, r *http.Request, productRepo repository.ProductRepository) {
	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/public-reviews/")
	pathParts := strings.Split(path, "/")

	if len(pathParts) == 0 || pathParts[0] == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Product ID is required"}`))
		return
	}

	productIDStr := pathParts[0]
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid product ID"}`))
		return
	}

	// Get locale from query parameter (default to 'en' like Laravel)
	locale := r.URL.Query().Get("locale")
	if locale == "" {
		locale = "en"
	}

	// For now, return a simple response structure matching Laravel
	// TODO: Implement actual review fetching logic when review repository is available
	reviews := []map[string]interface{}{
		{
			"id":                 1,
			"product_id":         uint(productID),
			"rating":             5,
			"reviews":            "Excellent product! Highly recommended.",
			"additional_details": "Great quality and fast delivery",
			"price":              999.99,
			"locale":             locale,
			"created_at":         time.Now().Format("2006-01-02T15:04:05Z07:00"),
			"updated_at":         time.Now().Format("2006-01-02T15:04:05Z07:00"),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"data":    reviews,
		"success": true,
		"message": "Public reviews retrieved successfully",
	}

	json.NewEncoder(w).Encode(response)
}
