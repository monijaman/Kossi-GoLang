package specification

import (
	"encoding/json"
	"fmt"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SpecificationResponse represents the response format for specifications
type SpecificationResponse struct {
	ID                 uint   `json:"id"`
	ProductID          uint   `json:"product_id"`
	SpecificationKeyID uint   `json:"specification_key_id"`
	SpecificationKey   string `json:"specification_key"`
	Value              string `json:"value"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

// SpecificationKeyResponse represents the response format for specification keys
type SpecificationKeyResponse struct {
	ID               uint   `json:"id"`
	SpecificationKey string `json:"specification_key"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

// convertSpecificationToResponse converts entity to response format
func convertSpecificationToResponse(spec *entities.Specification) SpecificationResponse {
	return SpecificationResponse{
		ID:                 spec.ID,
		ProductID:          spec.ProductID,
		SpecificationKeyID: spec.SpecificationKeyID,
		SpecificationKey:   spec.SpecificationKey,
		Value:              spec.Value,
		CreatedAt:          spec.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          spec.UpdatedAt.Format(time.RFC3339),
	}
}

// convertSpecificationKeyToResponse converts entity to response format
func convertSpecificationKeyToResponse(key *entities.SpecificationKey) SpecificationKeyResponse {
	return SpecificationKeyResponse{
		ID:               key.ID,
		SpecificationKey: key.SpecificationKey,
		CreatedAt:        key.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        key.UpdatedAt.Format(time.RFC3339),
	}
}

// CreateSpecificationHandler handles POST /specifications
func CreateSpecificationHandler(w http.ResponseWriter, r *http.Request, specRepo repository.SpecificationRepository, keyRepo repository.SpecificationKeyRepository) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		ProductID          uint   `json:"product_id"`
		SpecificationKeyID *uint  `json:"specification_key_id,omitempty"`
		Value              string `json:"value"`
		SpecificationKey   *struct {
			Name         string `json:"name"`
			IsVisible    bool   `json:"is_visible"`
			IsFilterable bool   `json:"is_filterable"`
		} `json:"specification_key,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.ProductID == 0 || request.Value == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "product_id and value are required"})
		return
	}

	// Either specification_key_id or specification_key object must be provided
	if request.SpecificationKeyID == nil && request.SpecificationKey == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Either specification_key_id or specification_key object is required"})
		return
	}

	var specKeyID uint

	// If SpecificationKeyID is provided, use it directly
	if request.SpecificationKeyID != nil {
		specKeyID = *request.SpecificationKeyID
	} else if request.SpecificationKey != nil {
		// Create or find the specification key
		specKeyName := request.SpecificationKey.Name

		// First, try to find existing key by name
		existingKey, err := keyRepo.GetByKey(r.Context(), specKeyName)
		if err == nil && existingKey != nil {
			// Key already exists, use its ID
			specKeyID = existingKey.ID
		} else {
			// Key doesn't exist, create a new one
			newKey := &entities.SpecificationKey{
				SpecificationKey: specKeyName,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			}

			createdKey, err := keyRepo.Create(r.Context(), newKey)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create specification key"})
				return
			}
			specKeyID = createdKey.ID
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Either specification_key_id or specification_key object is required"})
		return
	}

	spec := &entities.Specification{
		ProductID:          request.ProductID,
		SpecificationKeyID: specKeyID,
		Value:              request.Value,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	savedSpec, err := specRepo.Create(r.Context(), spec)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create or update specification"})
		return
	}

	response := convertSpecificationToResponse(savedSpec)
	json.NewEncoder(w).Encode(response)
}

// BulkUpsertSpecificationHandler handles POST /specifications/bulk
func BulkUpsertSpecificationHandler(w http.ResponseWriter, r *http.Request, specRepo repository.SpecificationRepository, keyRepo repository.SpecificationKeyRepository) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		Specifications []struct {
			ID                 *uint  `json:"id,omitempty"`
			ProductID          uint   `json:"product_id"`
			SpecificationKeyID *uint  `json:"specification_key_id,omitempty"`
			Value              string `json:"value"`
			SpecificationKey   *struct {
				Name         string `json:"name"`
				IsVisible    bool   `json:"is_visible"`
				IsFilterable bool   `json:"is_filterable"`
			} `json:"specification_key,omitempty"`
		} `json:"specifications"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if len(request.Specifications) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "specifications array cannot be empty"})
		return
	}

	var specs []*entities.Specification

	// Process each specification in the request
	for i, specReq := range request.Specifications {
		if specReq.ProductID == 0 || specReq.Value == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "product_id and value are required for all specifications",
				"index": strconv.Itoa(i),
			})
			return
		}

		// Either specification_key_id or specification_key object must be provided
		if specReq.SpecificationKeyID == nil && specReq.SpecificationKey == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Either specification_key_id or specification_key object is required",
				"index": strconv.Itoa(i),
			})
			return
		}

		var specKeyID uint

		// If SpecificationKeyID is provided, use it directly
		if specReq.SpecificationKeyID != nil {
			specKeyID = *specReq.SpecificationKeyID
		} else if specReq.SpecificationKey != nil {
			// Create or find the specification key
			specKeyName := specReq.SpecificationKey.Name

			// First, try to find existing key by name
			existingKey, err := keyRepo.GetByKey(r.Context(), specKeyName)
			if err == nil && existingKey != nil {
				// Key already exists, use its ID
				specKeyID = existingKey.ID
			} else {
				// Key doesn't exist, create a new one
				newKey := &entities.SpecificationKey{
					SpecificationKey: specKeyName,
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				}

				createdKey, err := keyRepo.Create(r.Context(), newKey)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(map[string]string{
						"error": "Failed to create specification key",
						"index": strconv.Itoa(i),
					})
					return
				}
				specKeyID = createdKey.ID
			}
		}

		spec := &entities.Specification{
			ProductID:          specReq.ProductID,
			SpecificationKeyID: specKeyID,
			Value:              specReq.Value,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		// Set ID if provided (for updates)
		if specReq.ID != nil {
			spec.ID = *specReq.ID
		}

		specs = append(specs, spec)
	}

	// Perform bulk upsert
	savedSpecs, err := specRepo.BulkUpsert(r.Context(), specs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to bulk upsert specifications"})
		return
	}

	// Convert to response format
	responses := make([]SpecificationResponse, len(savedSpecs))
	for i, spec := range savedSpecs {
		responses[i] = convertSpecificationToResponse(spec)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"specifications": responses,
		"count":          len(responses),
		"message":        "Bulk upsert completed successfully",
	})
}

// UpdateSpecificationHandler handles PUT /specifications/{id}
func UpdateSpecificationHandler(w http.ResponseWriter, r *http.Request, specRepo repository.SpecificationRepository, keyRepo repository.SpecificationKeyRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract specification ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/specifications/")
	specIDStr := strings.Trim(path, "/")

	specID, err := strconv.ParseUint(specIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid specification ID format",
			"received": specIDStr,
		})
		return
	}

	var request struct {
		SpecificationKeyID *uint  `json:"specification_key_id,omitempty"`
		Value              string `json:"value"`
		SpecificationKey   *struct {
			Name         string `json:"name"`
			IsVisible    bool   `json:"is_visible"`
			IsFilterable bool   `json:"is_filterable"`
		} `json:"specification_key,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.Value == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "value is required"})
		return
	}

	// Get existing specification
	existingSpec, err := specRepo.GetByID(r.Context(), uint(specID))
	if err != nil {
		if err.Error() == "specification not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Specification not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get specification"})
		}
		return
	}

	var specKeyID uint = existingSpec.SpecificationKeyID

	// Handle specification key update if provided
	if request.SpecificationKeyID != nil {
		specKeyID = *request.SpecificationKeyID
	} else if request.SpecificationKey != nil {
		// Create or find the specification key
		specKeyName := request.SpecificationKey.Name

		// First, try to find existing key by name
		existingKey, err := keyRepo.GetByKey(r.Context(), specKeyName)
		if err == nil && existingKey != nil {
			// Key already exists, use its ID
			specKeyID = existingKey.ID
		} else {
			// Key doesn't exist, create a new one
			newKey := &entities.SpecificationKey{
				SpecificationKey: specKeyName,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			}

			createdKey, err := keyRepo.Create(r.Context(), newKey)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create specification key"})
				return
			}
			specKeyID = createdKey.ID
		}
	}

	updatedSpec := &entities.Specification{
		ProductID:          existingSpec.ProductID,
		SpecificationKeyID: specKeyID,
		Value:              request.Value,
		UpdatedAt:          time.Now(),
	}

	savedSpec, err := specRepo.Update(r.Context(), uint(specID), updatedSpec)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update specification"})
		return
	}

	response := convertSpecificationToResponse(savedSpec)
	json.NewEncoder(w).Encode(response)
}

// GetSpecificationsByProductHandler handles GET /get-specifications/{id}
func GetSpecificationsByProductHandler(w http.ResponseWriter, r *http.Request, specRepo repository.SpecificationRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/get-specifications/")
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

	specifications, err := specRepo.GetByProductID(r.Context(), uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get specifications"})
		return
	}

	responses := make([]SpecificationResponse, len(specifications))
	for i, spec := range specifications {
		responses[i] = convertSpecificationToResponse(spec)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"product_id":     productID,
		"specifications": responses,
		"count":          len(responses),
	})
}

// GetSpecificationByIDHandler handles GET /specifications/{id}
func GetSpecificationByIDHandler(w http.ResponseWriter, r *http.Request, specRepo repository.SpecificationRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract specification ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/specifications/")
	specIDStr := strings.Trim(path, "/")

	specID, err := strconv.ParseUint(specIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid specification ID format",
			"received": specIDStr,
		})
		return
	}

	specification, err := specRepo.GetByID(r.Context(), uint(specID))
	if err != nil {
		if err.Error() == "specification not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Specification not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get specification"})
		}
		return
	}

	response := convertSpecificationToResponse(specification)
	json.NewEncoder(w).Encode(response)
}

// DeleteSpecificationHandler handles DELETE /specifications/{id}
func DeleteSpecificationHandler(w http.ResponseWriter, r *http.Request, specRepo repository.SpecificationRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract specification ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/specifications/")
	specIDStr := strings.Trim(path, "/")

	specID, err := strconv.ParseUint(specIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid specification ID format",
			"received": specIDStr,
		})
		return
	}

	err = specRepo.Delete(r.Context(), uint(specID))
	if err != nil {
		if err.Error() == "specification not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Specification not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete specification"})
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Specification deleted successfully"})
}

// SearchSpecificationsHandler handles GET /specificationsearch
func SearchSpecificationsHandler(w http.ResponseWriter, r *http.Request, specRepo repository.SpecificationRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	query := r.URL.Query().Get("query")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	// if query == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(map[string]string{"error": "query parameter is required"})
	// 	return
	// }

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

	specifications, err := specRepo.Search(r.Context(), query, limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to search specifications"})
		return
	}

	responses := make([]SpecificationResponse, len(specifications))
	for i, spec := range specifications {
		responses[i] = convertSpecificationToResponse(spec)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"specifications": responses,
		"count":          len(responses),
		"query":          query,
		"limit":          limit,
		"offset":         offset,
	})
}

// CreateSpecificationTranslationHandler handles POST /spec_translation
func CreateSpecificationTranslationHandler(w http.ResponseWriter, r *http.Request, specRepo repository.SpecificationRepository) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		ProductID      uint `json:"productId"`
		Specifications []struct {
			ID              uint   `json:"id"` // This could be either specification_id or specification_key_id
			Locale          string `json:"locale"`
			TranslatedKey   string `json:"translated_key"`
			TranslatedValue string `json:"translated_value"`
		} `json:"specifications"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.ProductID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "productId is required"})
		return
	}

	if len(request.Specifications) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "specifications array cannot be empty"})
		return
	}

	var savedTranslations []*entities.SpecificationTranslation

	// Process each specification translation
	for i, specReq := range request.Specifications {
		if specReq.ID == 0 || specReq.Locale == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("specification id and locale are required for all specifications at index %d", i),
			})
			return
		}

		// The frontend is sending the actual specification_id as id
		// First, get all specifications for this product to find the matching specification
		productSpecs, err := specRepo.GetByProductID(r.Context(), request.ProductID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get product specifications"})
			return
		}

		// Find the specification that matches this specification_id
		var targetSpec *entities.Specification
		for _, spec := range productSpecs {
			if spec.ID == specReq.ID {
				targetSpec = spec
				break
			}
		}

		if targetSpec == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("No specification found for specification_id %d in product %d", specReq.ID, request.ProductID),
			})
			return
		}

		// Check if translation already exists
		existingTranslations, err := specRepo.GetTranslations(r.Context(), targetSpec.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to check existing translations"})
			return
		}

		var existingTranslation *entities.SpecificationTranslation
		for _, trans := range existingTranslations {
			if trans.Locale == specReq.Locale {
				existingTranslation = trans
				break
			}
		}

		// Use translated_value if provided, otherwise skip this translation
		translatedValue := specReq.TranslatedValue
		if translatedValue == "" {
			continue // Skip empty translations
		}

		if existingTranslation != nil {
			// Update existing translation
			existingTranslation.TranslatedValue = translatedValue
			existingTranslation.UpdatedAt = time.Now()

			// Note: We need to implement UpdateTranslation method in repository
			// For now, we'll delete and recreate (not ideal, but works)
			savedTranslation, err := specRepo.CreateTranslation(r.Context(), existingTranslation)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update specification translation"})
				return
			}
			savedTranslations = append(savedTranslations, savedTranslation)
		} else {
			// Create new translation
			translation := &entities.SpecificationTranslation{
				SpecificationID: targetSpec.ID,
				Locale:          specReq.Locale,
				TranslatedValue: translatedValue,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			}

			savedTranslation, err := specRepo.CreateTranslation(r.Context(), translation)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create specification translation"})
				return
			}
			savedTranslations = append(savedTranslations, savedTranslation)
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "Specifications translations saved successfully!",
		"data":       savedTranslations,
		"count":      len(savedTranslations),
		"product_id": request.ProductID,
	})
}

// GetSpecificationTranslationHandler handles GET /spec_translation/{id}?locale=xx
func GetSpecificationTranslationHandler(w http.ResponseWriter, r *http.Request, specRepo repository.SpecificationRepository, productRepo repository.ProductRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/spec_translation/")
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

	// Get locale from query parameter
	locale := r.URL.Query().Get("locale")
	if locale == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "locale parameter is required"})
		return
	}

	// Get all specifications for the product
	specifications, err := specRepo.GetByProductID(r.Context(), uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get product specifications"})
		return
	}

	// Build response in Laravel-compatible format
	var formattedDataset []map[string]interface{}

	for _, spec := range specifications {
		// Get translations for this specification
		translations, err := specRepo.GetTranslations(r.Context(), spec.ID)
		if err != nil {
			continue // Skip on error, don't fail the whole request
		}

		// Find translation for the requested locale
		var translation *entities.SpecificationTranslation
		for _, trans := range translations {
			if trans.Locale == locale {
				translation = trans
				break
			}
		}

		specData := map[string]interface{}{
			"specification_key_id": spec.SpecificationKeyID,
			"translations":         nil,
		}

		if translation != nil {
			specData["translations"] = map[string]interface{}{
				"specification_id": translation.SpecificationID,
				"locale":           translation.Locale,
				"translated_key":   "", // This would come from SpecificationKeyTranslation
				"translated_value": translation.TranslatedValue,
			}
		}

		formattedDataset = append(formattedDataset, specData)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"dataset": formattedDataset,
	})
}

// GetPublicSpecificationHandler handles GET /get-public-spec/{id} - Get specifications for a product (Laravel compatibility)
func GetPublicSpecificationHandler(w http.ResponseWriter, r *http.Request, specRepo repository.SpecificationRepository, keyRepo repository.SpecificationKeyRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/get-public-spec/")
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

	// Get locale from query parameters (default to 'en')
	locale := r.URL.Query().Get("locale")
	if locale == "" {
		locale = "en"
	}

	// Get all specifications for the product
	specifications, err := specRepo.GetByProductID(r.Context(), uint(productID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get specifications"})
		return
	}

	type SpecificationResponse struct {
		SpecificationKeyID uint   `json:"specification_key_id"`
		TranslatedKey      string `json:"translated_key"`
		TranslatedValue    string `json:"translated_value"`
	}

	var dataset []SpecificationResponse
	for _, spec := range specifications {
		translatedKey := spec.SpecificationKey // fallback
		translatedValue := spec.Value          // fallback

		// Get translated key if locale is not 'en'
		if locale != "en" {
			if keyTranslation, err := keyRepo.GetKeyTranslationByLocale(r.Context(), spec.SpecificationKeyID, locale); err == nil && keyTranslation != nil && keyTranslation.TranslatedKey != "" {
				translatedKey = keyTranslation.TranslatedKey
			}

			// Get translated value if locale is not 'en'
			if valueTranslation, err := specRepo.GetTranslationByLocale(r.Context(), spec.ID, locale); err == nil && valueTranslation != nil && valueTranslation.TranslatedValue != "" {
				translatedValue = valueTranslation.TranslatedValue
			}
		}

		dataset = append(dataset, SpecificationResponse{
			SpecificationKeyID: spec.SpecificationKeyID,
			TranslatedKey:      translatedKey,
			TranslatedValue:    translatedValue,
		})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"dataset": dataset,
	})
}

// UpdateSpecificationTranslationValues handles PUT/PATCH requests to update only translated_value
func UpdateSpecificationTranslationValues(w http.ResponseWriter, r *http.Request, specRepo repository.SpecificationRepository) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		ProductID      uint `json:"productId"`
		Specifications []struct {
			ID              uint   `json:"id"` // specification_id
			Locale          string `json:"locale"`
			TranslatedValue string `json:"translated_value"`
		} `json:"specifications"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.ProductID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "productId is required"})
		return
	}

	var savedTranslations []*entities.SpecificationTranslation

	// Process each specification translation
	for i, specReq := range request.Specifications {
		if specReq.ID == 0 || specReq.Locale == "" || specReq.TranslatedValue == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": fmt.Sprintf("specification id, locale and translated_value are required at index %d", i),
			})
			return
		}

		// Create or update translation using upsert approach
		translation := &entities.SpecificationTranslation{
			SpecificationID: specReq.ID,
			Locale:          specReq.Locale,
			TranslatedValue: specReq.TranslatedValue,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		// Use the new UpsertTranslation method which handles create/update automatically
		savedTranslation, err := specRepo.UpsertTranslation(r.Context(), translation)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Failed to upsert specification translation: %v", err)})
			return
		}
		savedTranslations = append(savedTranslations, savedTranslation)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "Specification translation values updated successfully!",
		"data":       savedTranslations,
		"count":      len(savedTranslations),
		"product_id": request.ProductID,
	})
}
