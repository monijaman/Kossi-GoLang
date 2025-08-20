package formgenerator

import (
	"encoding/json"
	"fmt"
	"kossti/internal/domain/entities"
	"kossti/internal/domain/repository"
	"kossti/internal/usecase/formgenerator"
	"net/http"
	"strconv"
	"strings"
)

// Request/Response structures
type CreateFormGeneratorRequest struct {
	CategoryID      uint   `json:"category_id"`
	SpecificationID []uint `json:"specification_id"`
	Status          string `json:"status,omitempty"`
}

type UpdateFormGeneratorRequest struct {
	CategoryID      *uint   `json:"category_id,omitempty"`
	SpecificationID []uint  `json:"specification_id,omitempty"`
	Status          *string `json:"status,omitempty"`
}

type FormGeneratorResponse struct {
	ID              uint   `json:"id"`
	CategoryID      uint   `json:"category_id"`
	SpecificationID string `json:"specification_id"`
	Status          string `json:"status"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type FormGeneratorDataResponse struct {
	Message string                  `json:"message"`
	Data    []FormGeneratorResponse `json:"data"`
}

type SingleFormGeneratorResponse struct {
	Message string                `json:"message"`
	Data    FormGeneratorResponse `json:"data"`
}

type SpecificationResponse struct {
	ID               uint   `json:"id"`
	SpecificationKey string `json:"specification_key"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type CategorySpecResponse struct {
	Message string                  `json:"message"`
	Data    []SpecificationResponse `json:"data"`
}

// convertFormGeneratorToResponse converts entity to response format
func convertFormGeneratorToResponse(fg *entities.FormGenerator) FormGeneratorResponse {
	return FormGeneratorResponse{
		ID:              fg.ID,
		CategoryID:      fg.CategoryID,
		SpecificationID: fg.SpecificationID,
		Status:          fg.Status,
		CreatedAt:       fg.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:       fg.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// FormGeneratorHandler handles HTTP requests for form generator operations
type FormGeneratorHandler struct {
	repo repository.FormGeneratorRepository
}

// NewFormGeneratorHandler creates a new form generator handler
func NewFormGeneratorHandler(repo repository.FormGeneratorRepository) *FormGeneratorHandler {
	return &FormGeneratorHandler{
		repo: repo,
	}
}

// CreateFormGenerator handles POST /formgenerator
func (h *FormGeneratorHandler) CreateFormGenerator(w http.ResponseWriter, r *http.Request) {
	var req CreateFormGeneratorRequest

	path := strings.TrimPrefix(r.URL.Path, "/formgenerator/")

	idStr := strings.Split(path, "/")[0]

	id, err := strconv.ParseUint(idStr, 10, 32)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if id == 0 {
		http.Error(w, "category_id is required", http.StatusBadRequest)
		return
	}

	if len(req.SpecificationID) == 0 {
		http.Error(w, "specification_id array is required", http.StatusBadRequest)
		return
	}

	// Try to update first (if form generator exists for this category)
	fg, err := formgenerator.UpdateFormGenerator(r.Context(), h.repo, uint(id), req.SpecificationID, &req.Status)

	if err != nil {
		// If update fails (form generator doesn't exist), create a new one
		if strings.Contains(err.Error(), "form generator not found") {
			fg, err = formgenerator.CreateFormGenerator(r.Context(), h.repo, uint(id), req.SpecificationID, req.Status)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Return the created/updated form generator
	response := convertFormGeneratorToResponse(fg)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SingleFormGeneratorResponse{
		Message: "Form generator created/updated successfully!",
		Data:    response,
	})
}

// GetFormGeneratorByID handles GET /formgenerator/{id}
func (h *FormGeneratorHandler) GetFormGeneratorByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/formgenerator/")
	if path == r.URL.Path {
		path = strings.TrimPrefix(r.URL.Path, "/formgenerator/")
	}

	idStr := strings.Split(path, "/")[0]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid form generator ID", http.StatusBadRequest)
		return
	}

	fg, err := formgenerator.GetFormGeneratorByCategoryID(r.Context(), h.repo, uint(id))
	if err != nil {
		// Return empty array if not found, matching Laravel behavior
		response := FormGeneratorDataResponse{
			Message: "Form Generator data retrieved successfully!",
			Data:    []FormGeneratorResponse{},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := FormGeneratorDataResponse{
		Message: "Form Generator data retrieved successfully!",
		Data:    []FormGeneratorResponse{convertFormGeneratorToResponse(fg)},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateFormGenerator handles PUT /formgenerator/{id}
func (h *FormGeneratorHandler) UpdateFormGenerator(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/formgenerator/")

	idStr := strings.Split(path, "/")[0]

	id, err := strconv.ParseUint(idStr, 10, 32)

	var req UpdateFormGeneratorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	fg, err := formgenerator.UpdateFormGenerator(r.Context(), h.repo, uint(id), req.SpecificationID, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := SingleFormGeneratorResponse{
		Message: "Form Generator updated successfully!",
		Data:    convertFormGeneratorToResponse(fg),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetCategorySpec handles GET /category-specs/{categoryId}
func (h *FormGeneratorHandler) GetCategorySpec(w http.ResponseWriter, r *http.Request) {
	// Extract category ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/category-specs/")

	idStr := strings.Split(path, "/")[0]
	categoryID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	specifications, err := formgenerator.GetCategorySpecifications(r.Context(), h.repo, uint(categoryID))
	fmt.Println("-----------------specifications", specifications)

	if err != nil {
		// Return empty array if not found
		response := CategorySpecResponse{
			Message: "Form Generator data retrieved successfully!",
			Data:    []SpecificationResponse{},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Convert specifications to response format
	var specResponses []SpecificationResponse
	for _, spec := range specifications {
		specResponses = append(specResponses, SpecificationResponse{
			ID:               spec.ID,
			SpecificationKey: spec.SpecificationKey,
			CreatedAt:        spec.CreatedAt.Format("2006-01-02T15:04:05.000000Z07:00"),
			UpdatedAt:        spec.UpdatedAt.Format("2006-01-02T15:04:05.000000Z07:00"),
		})
	}

	response := CategorySpecResponse{
		Message: "Form Generator data retrieved successfully!",
		Data:    specResponses,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
