package seeders

import (
	"kossti/internal/infrastructure/database/models"
	"log"

	"gorm.io/gorm"
)

type MobileSpecificationsExistingKeysSeeder struct{}

func NewMobileSpecificationsExistingKeysSeeder() *MobileSpecificationsExistingKeysSeeder {
	return &MobileSpecificationsExistingKeysSeeder{}
}

func (s *MobileSpecificationsExistingKeysSeeder) GetName() string {
	return "MobileSpecificationsExistingKeysSeeder"
}

func (s *MobileSpecificationsExistingKeysSeeder) Seed(db *gorm.DB) error {
	log.Println("Starting Mobile Specifications Existing Keys Seeder...")

	// Get all mobile products (category_id = 79)
	var products []models.ProductModel

	if err := db.Where("category_id = ?", 79).Order("id").Find(&products).Error; err != nil {
		return err
	}

	log.Printf("Found %d mobile products", len(products))

	// Get default mobile specs structure
	specs := DefaultMobileSpecs()
	totalSpecs := 0

	for _, product := range products {
		log.Printf("Processing product: %s (ID: %d)", product.Name, product.ID)

		for keyName, value := range specs {
			// Find or create the specification key
			sk, err := CreateOrFindSpecificationKey(db, keyName)
			if err != nil {
				log.Printf("Error finding/creating key '%s': %v", keyName, err)
				continue
			}

			// Check if specification already exists
			var existing models.SpecificationModel
			if err := db.Where("product_id = ? AND specification_key_id = ?", product.ID, sk.ID).First(&existing).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// Create new specification with empty value
					spec := &models.SpecificationModel{
						ProductID:          product.ID,
						SpecificationKeyID: sk.ID,
						Value:              value,
						Status:             1,
					}

					if err := db.Create(spec).Error; err != nil {
						log.Printf("Error creating specification for product %d, key %s: %v", product.ID, keyName, err)
						continue
					}
					totalSpecs++
				} else {
					log.Printf("Error checking existing specification: %v", err)
					continue
				}
			}
		}
	}

	log.Printf("Successfully seeded %d specifications for %d products",
		totalSpecs, len(products))

	return nil
}
