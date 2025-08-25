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

// CreateOrUpdateSpecificationKeyHandler handles POST /speckey
func CreateOrUpdateSpecificationKeyHandler(w http.ResponseWriter, r *http.Request, keyRepo repository.SpecificationKeyRepository) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		ID               *uint  `json:"id,omitempty"`
		SpecificationKey string `json:"specification_key"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.SpecificationKey == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "specification_key is required"})
		return
	}

	if request.ID != nil && *request.ID > 0 {
		// Update existing key
		key := &entities.SpecificationKey{
			SpecificationKey: request.SpecificationKey,
			UpdatedAt:        time.Now(),
		}

		updatedKey, err := keyRepo.Update(r.Context(), *request.ID, key)
		if err != nil {
			if err.Error() == "specification key not found" {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"error": "Specification key not found"})
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update specification key"})
			}
			return
		}

		response := convertSpecificationKeyToResponse(updatedKey)
		json.NewEncoder(w).Encode(response)
	} else {
		// Create new key
		key := &entities.SpecificationKey{
			SpecificationKey: request.SpecificationKey,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}

		savedKey, err := keyRepo.Create(r.Context(), key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create specification key"})
			return
		}

		response := convertSpecificationKeyToResponse(savedKey)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

// GetAllSpecificationKeysHandler handles GET /speckey
func GetAllSpecificationKeysHandler(w http.ResponseWriter, r *http.Request, keyRepo repository.SpecificationKeyRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	perPageStr := r.URL.Query().Get("per_page")
	searchTerm := r.URL.Query().Get("search")

	// Set defaults
	limit := 10
	offset := 0

	// Handle per_page parameter (alternative to limit)
	if perPageStr != "" {
		if pp, err := strconv.Atoi(perPageStr); err == nil && pp > 0 {
			limit = pp
		}
	}

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

	var keys []*entities.SpecificationKey
	var totalCount int64
	var err error

	// If search term is provided, use search functionality
	if searchTerm != "" {
		fmt.Print("Searching for " + searchTerm)
		keys, err = keyRepo.Search(r.Context(), searchTerm, limit, offset)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to search specification keys"})
			return
		}

		// Get total count for search results
		totalCount, err = keyRepo.GetSearchCount(r.Context(), searchTerm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get search count"})
			return
		}
	} else {
		// No search term, get all keys
		keys, err = keyRepo.GetAll(r.Context(), limit, offset)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get specification keys"})
			return
		}

		// Get total count of all specification keys
		totalCount, err = keyRepo.GetCount(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get total count"})
			return
		}
	}

	responses := make([]SpecificationKeyResponse, len(keys))
	for i, key := range keys {
		responses[i] = convertSpecificationKeyToResponse(key)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"specification_keys": responses,
		"total":              totalCount,     // Total count of all keys
		"returned":           len(responses), // Number of keys returned in this page
		"limit":              limit,
		"offset":             offset,
	})
}

// GetSpecificationKeyByIDHandler handles GET /speckey/{id}
func GetSpecificationKeyByIDHandler(w http.ResponseWriter, r *http.Request, keyRepo repository.SpecificationKeyRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract key ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/speckey/")
	keyIDStr := strings.Trim(path, "/")

	keyID, err := strconv.ParseUint(keyIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid specification key ID format",
			"received": keyIDStr,
		})
		return
	}

	key, err := keyRepo.GetByID(r.Context(), uint(keyID))
	if err != nil {
		if err.Error() == "specification key not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Specification key not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get specification key"})
		}
		return
	}

	response := convertSpecificationKeyToResponse(key)
	json.NewEncoder(w).Encode(response)
}

// DeleteSpecificationKeyHandler handles POST /specremove/{id}
func DeleteSpecificationKeyHandler(w http.ResponseWriter, r *http.Request, keyRepo repository.SpecificationKeyRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Extract key ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/specremove/")
	keyIDStr := strings.Trim(path, "/")

	keyID, err := strconv.ParseUint(keyIDStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Invalid specification key ID format",
			"received": keyIDStr,
		})
		return
	}

	err = keyRepo.Delete(r.Context(), uint(keyID))
	if err != nil {
		if err.Error() == "specification key not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Specification key not found"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete specification key"})
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Specification key deleted successfully",
		"id":      keyIDStr,
	})
}

// CreateSpecificationKeyTranslationHandler handles POST /speckey-translation
// This handler implements upsert logic: update if exists, create if not
func CreateSpecificationKeyTranslationHandler(w http.ResponseWriter, r *http.Request, keyRepo repository.SpecificationKeyRepository) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		SpecificationKeyID uint   `json:"specification_key_id"`
		Locale             string `json:"locale"`
		TranslatedKey      string `json:"translated_key"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	if request.SpecificationKeyID == 0 || request.Locale == "" || request.TranslatedKey == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "specification_key_id, locale, and translated_key are required"})
		return
	}

	// Check if translation already exists
	existingTranslation, err := keyRepo.GetKeyTranslationByLocale(r.Context(), request.SpecificationKeyID, request.Locale)

	var savedTranslation *entities.SpecificationKeyTranslation

	if err != nil && err.Error() == "key translation not found" {
		// Translation doesn't exist, create new one
		translation := &entities.SpecificationKeyTranslation{
			SpecificationKeyID: request.SpecificationKeyID,
			Locale:             request.Locale,
			TranslatedKey:      request.TranslatedKey,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		savedTranslation, err = keyRepo.CreateKeyTranslation(r.Context(), translation)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create specification key translation"})
			return
		}
	} else if err != nil {
		// Some other error occurred
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to check existing translation"})
		return
	} else {
		// Translation exists, update it
		existingTranslation.TranslatedKey = request.TranslatedKey
		existingTranslation.UpdatedAt = time.Now()

		savedTranslation, err = keyRepo.UpdateKeyTranslation(r.Context(), existingTranslation.ID, existingTranslation)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update specification key translation"})
			return
		}
	}

	json.NewEncoder(w).Encode(savedTranslation)
}

// GetAllSpecificationKeyTranslationsHandler handles GET /speckey-translation
// Supports query parameters: key_id and locale for filtering
func GetAllSpecificationKeyTranslationsHandler(w http.ResponseWriter, r *http.Request, keyRepo repository.SpecificationKeyRepository) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameters
	keyIDStr := r.URL.Query().Get("key_id")
	locale := r.URL.Query().Get("locale")

	// If both key_id and locale are provided, get specific translation
	if keyIDStr != "" && locale != "" {
		keyID, err := strconv.ParseUint(keyIDStr, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error":    "Invalid key_id format",
				"received": keyIDStr,
			})
			return
		}

		translation, err := keyRepo.GetKeyTranslationByLocale(r.Context(), uint(keyID), locale)
		if err != nil {
			if err.Error() == "key translation not found" {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"error": "Translation not found"})
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get translation"})
			}
			return
		}

		json.NewEncoder(w).Encode(translation)
		return
	}

	// If only key_id is provided, get all translations for that key
	if keyIDStr != "" {
		keyID, err := strconv.ParseUint(keyIDStr, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error":    "Invalid key_id format",
				"received": keyIDStr,
			})
			return
		}

		translations, err := keyRepo.GetKeyTranslations(r.Context(), uint(keyID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get translations for key"})
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"translations": translations,
			"total":        len(translations),
		})
		return
	}

	// If no parameters provided, get all translations
	translations, err := keyRepo.GetAllKeyTranslations(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get specification key translations"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"translations": translations,
		"total":        len(translations),
	})
}
