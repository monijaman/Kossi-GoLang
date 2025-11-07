package seeders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// ProductSeeder handles seeding products (motorcycles) for Bangladesh
type ProductSeeder struct {
	BaseSeeder
}

// NewProductSeeder creates a new ProductSeeder
func NewProductSeeder() *ProductSeeder {
	return &ProductSeeder{BaseSeeder: BaseSeeder{name: "Products (Motorcycles)"}}
}

// Seed implements the Seeder interface
func (ps *ProductSeeder) Seed(db *gorm.DB) error {
	// Ensure category exists
	_, err := CreateOrFindCategory(db, "Motorcycles", "motorcycles")
	if err != nil {
		return err
	}

	// Load dataset from JSON file for maintainability
	jsonPath := filepath.Join("init-db", "seeders", "motorbikes.json")
	raw, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("failed to read motorbikes dataset %s: %w", jsonPath, err)
	}

	var bikes []struct {
		NameEN string   `json:"name_en"`
		NameBN string   `json:"name_bn"`
		Brand  string   `json:"brand"`
		Model  string   `json:"model"`
		Price  float64  `json:"price"`
		Images []string `json:"images"`
	}

	if err := json.Unmarshal(raw, &bikes); err != nil {
		return fmt.Errorf("failed to parse motorbikes dataset: %w", err)
	}

	// Create or skip existing
	for _, b := range bikes {
		// Ensure brand exists (slug generated inside helper)
		brandEntity, err := CreateOrFindBrand(db, b.Brand, GenerateSlug(b.Brand))
		if err != nil {
			return err
		}

		// Find category id
		var catModel models.CategoryModel
		if err := db.Where("slug = ?", "motorcycles").First(&catModel).Error; err != nil {
			return err
		}

		// Compose product entity
		prod := &entities.Product{
			Name:       b.NameEN,
			Slug:       GenerateSlug(b.NameEN),
			Price:      b.Price,
			CategoryID: &catModel.ID,
			BrandID:    &brandEntity.ID,
			ViewsCount: 0,
			Status:     true,
			Priority:   1,
		}

		var existing models.ProductModel
		if err := db.Where("slug = ?", prod.Slug).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// create
				prodModel := &models.ProductModel{}
				prodModel.FromEntity(prod)
				// set model string field
				m := b.Model
				prodModel.Model = &m
				// ensure price pointer
				p := prod.Price
				prodModel.Price = &p

				if err := db.Create(prodModel).Error; err != nil {
					return err
				}

				// add images
				for _, img := range b.Images {
					image := &models.ImageModel{
						ImageableType: "Product",
						ImageableID:   prodModel.ID,
						ImagePath:     img,
						Status:        1,
					}
					if err := db.Create(image).Error; err != nil {
						return err
					}
				}

				// add Bangla translation
				priceStr := fmt.Sprintf("%.0f", b.Price)
				pt := &models.ProductTranslationModel{
					ProductID:      prodModel.ID,
					Locale:         "bn",
					TranslatedName: b.NameBN,
					Price:          &priceStr,
				}
				if err := db.Create(pt).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
