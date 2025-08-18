package specification

import (
	"encoding/json"
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

	keys, err := keyRepo.GetAll(r.Context(), limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get specification keys"})
		return
	}

	responses := make([]SpecificationKeyResponse, len(keys))
	for i, key := range keys {
		responses[i] = convertSpecificationKeyToResponse(key)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"specification_keys": responses,
		"count":              len(responses),
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

	translation := &entities.SpecificationKeyTranslation{
		SpecificationKeyID: request.SpecificationKeyID,
		Locale:             request.Locale,
		TranslatedKey:      request.TranslatedKey,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	savedTranslation, err := keyRepo.CreateKeyTranslation(r.Context(), translation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create specification key translation"})
		return
	}

	json.NewEncoder(w).Encode(savedTranslation)
}

// GetAllSpecificationKeyTranslationsHandler handles GET /speckey-translation
func GetAllSpecificationKeyTranslationsHandler(w http.ResponseWriter, r *http.Request, keyRepo repository.SpecificationKeyRepository) {
	w.Header().Set("Content-Type", "application/json")

	translations, err := keyRepo.GetAllKeyTranslations(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get specification key translations"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"translations": translations,
		"count":        len(translations),
	})
}
