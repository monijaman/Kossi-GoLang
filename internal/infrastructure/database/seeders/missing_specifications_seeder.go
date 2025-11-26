package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"time"

	"gorm.io/gorm"
)

// MissingSpecificationsSeeder handles seeding missing specifications and their translations
type MissingSpecificationsSeeder struct {
	BaseSeeder
}

// NewMissingSpecificationsSeeder creates a new seeder for missing specifications
func NewMissingSpecificationsSeeder() *MissingSpecificationsSeeder {
	return &MissingSpecificationsSeeder{
		BaseSeeder: BaseSeeder{name: "Missing Specifications"},
	}
}

// Seed implements the Seeder interface
func (ms *MissingSpecificationsSeeder) Seed(db *gorm.DB) error {
	// Missing specification keys and their translations
	specifications := map[string]string{
		"Fuel Efficiency (km/l)":                  "Fuel Efficiency (km/l)",
		"Transmission Type":                       "Transmission Type",
		"Engine Displacement (cc)":                "Engine Displacement (cc)",
		"Cylinders":                               "Cylinders",
		"Max Power (hp)":                          "Max Power (hp)",
		"Max Torque (Nm)":                         "Max Torque (Nm)",
		"City Fuel Economy (km/l)":                "City Fuel Economy (km/l)",
		"Highway Fuel Economy (km/l)":             "Highway Fuel Economy (km/l)",
		"Combined Fuel Economy (km/l)":            "Combined Fuel Economy (km/l)",
		"Fuel Tank Capacity (l)":                  "Fuel Tank Capacity (l)",
		"Length (mm)":                             "Length (mm)",
		"Width (mm)":                              "Width (mm)",
		"Height (mm)":                             "Height (mm)",
		"Wheelbase (mm)":                          "Wheelbase (mm)",
		"AC Type":                                 "AC Type",
		"Power Windows":                           "Power Windows",
		"Seat Material":                           "Seat Material",
		"Driver Airbag":                           "Driver Airbag",
		"Passenger Airbag":                        "Passenger Airbag",
		"Side Airbags":                            "Side Airbags",
		"Curtain Airbags":                         "Curtain Airbags",
		"Camera":                                  "Camera",
		"Child Safety Lock":                       "Child Safety Lock",
		"Stability Control":                       "Stability Control",
		"Service Interval (km)":                   "Service Interval (km)",
		"Spare Parts Availability":                "Spare Parts Availability",
		"Estimated Annual Maintenance Cost (BDT)": "Estimated Annual Maintenance Cost (BDT)",
		"Registration Cost (BDT)":                 "Registration Cost (BDT)",
		"Insurance Cost (BDT)":                    "Insurance Cost (BDT)",
	}

	// Insert missing specifications and translations
	for key, translation := range specifications {
		var existingKey models.SpecificationKeyModel
		result := db.Where("specification_key = ?", key).First(&existingKey)

		if result.Error == gorm.ErrRecordNotFound {
			// Create new specification key
			specKey := models.SpecificationKeyModel{
				SpecificationKey: key,
			}
			if err := db.Create(&specKey).Error; err != nil {
				continue
			}
		}

		// Insert translation
		var existingTranslation models.SpecificationTranslationModel
		result = db.Where("specification_key = ? AND language = ?", key, "en").First(&existingTranslation)

		if result.Error == gorm.ErrRecordNotFound {
			// Retrieve the SpecificationID for the current key
			var specKey models.SpecificationKeyModel
			result := db.Where("specification_key = ?", key).First(&specKey)
			if result.Error != nil {
				continue // Skip if the specification key is not found
			}

			// Create new translation
			translation := models.SpecificationTranslationModel{
				SpecificationID: specKey.ID,  // Use the ID of the specification key
				Locale:          "en",        // Specify the language as English
				Value:           translation, // Use the translation value
			}
			_ = db.Create(&translation)
		}
	}

	// Add Bangla translations for the specifications
	banglaTranslations := map[string]string{
		"Fuel Efficiency (km/l)":                  "জ্বালানি দক্ষতা (কিমি/লি)",
		"Transmission Type":                       "ট্রান্সমিশন প্রকার",
		"Engine Displacement (cc)":                "ইঞ্জিন স্থানচ্যুতি (সিসি)",
		"Cylinders":                               "সিলিন্ডার",
		"Max Power (hp)":                          "সর্বাধিক শক্তি (এইচপি)",
		"Max Torque (Nm)":                         "সর্বাধিক টর্ক (এনএম)",
		"City Fuel Economy (km/l)":                "শহরের জ্বালানি অর্থনীতি (কিমি/লি)",
		"Highway Fuel Economy (km/l)":             "হাইওয়ে জ্বালানি অর্থনীতি (কিমি/লি)",
		"Combined Fuel Economy (km/l)":            "সমন্বিত জ্বালানি অর্থনীতি (কিমি/লি)",
		"Fuel Tank Capacity (l)":                  "জ্বালানি ট্যাঙ্কের ক্ষমতা (লি)",
		"Length (mm)":                             "দৈর্ঘ্য (মিমি)",
		"Width (mm)":                              "প্রস্থ (মিমি)",
		"Height (mm)":                             "উচ্চতা (মিমি)",
		"Wheelbase (mm)":                          "হুইলবেস (মিমি)",
		"AC Type":                                 "এসি প্রকার",
		"Power Windows":                           "পাওয়ার উইন্ডোজ",
		"Seat Material":                           "আসনের উপাদান",
		"Driver Airbag":                           "ড্রাইভার এয়ারব্যাগ",
		"Passenger Airbag":                        "যাত্রী এয়ারব্যাগ",
		"Side Airbags":                            "পার্শ্ব এয়ারব্যাগ",
		"Curtain Airbags":                         "কার্টেন এয়ারব্যাগ",
		"Camera":                                  "ক্যামেরা",
		"Child Safety Lock":                       "শিশু নিরাপত্তা লক",
		"Stability Control":                       "স্থিতিশীলতা নিয়ন্ত্রণ",
		"Service Interval (km)":                   "পরিষেবা ব্যবধান (কিমি)",
		"Spare Parts Availability":                "খুচরা যন্ত্রাংশের প্রাপ্যতা",
		"Estimated Annual Maintenance Cost (BDT)": "প্রাক্কলিত বার্ষিক রক্ষণাবেক্ষণ খরচ (টাকা)",
		"Registration Cost (BDT)":                 "নিবন্ধন খরচ (টাকা)",
		"Insurance Cost (BDT)":                    "বীমা খরচ (টাকা)",
	}

	// Correct all field names for SpecificationTranslationModel
	for key, translation := range specifications {
		var specKey models.SpecificationKeyModel
		result := db.Where("specification_key = ?", key).First(&specKey)
		if result.Error != nil {
			continue // Skip if the specification key is not found
		}

		// Insert English translation
		var existingTranslation models.SpecificationTranslationModel
		result = db.Where("specification_id = ? AND locale = ?", specKey.ID, "en").First(&existingTranslation)

		if result.Error == gorm.ErrRecordNotFound {
			// Create new English translation
			translation := models.SpecificationTranslationModel{
				SpecificationID: specKey.ID,
				Locale:          "en",
				Value:           translation,
			}
			_ = db.Create(&translation)
		}
	}

	// Insert Bangla translations
	for key, translation := range banglaTranslations {
		var specKey models.SpecificationKeyModel
		result := db.Where("specification_key = ?", key).First(&specKey)
		if result.Error != nil {
			continue // Skip if the specification key is not found
		}

		var existingTranslation models.SpecificationTranslationModel
		result = db.Where("specification_id = ? AND locale = ?", specKey.ID, "bn").First(&existingTranslation)

		if result.Error == gorm.ErrRecordNotFound {
			// Create new Bangla translation
			translation := models.SpecificationTranslationModel{
				SpecificationID: specKey.ID,
				Locale:          "bn",
				Value:           translation,
			}
			_ = db.Create(&translation)
		}
	}

	// Updating the Bangla translations map to avoid redeclaration
	banglaTranslations["Estimated Annual Maintenance Cost (BDT)"] = "প্রাক্কলিত বার্ষিক রক্ষণাবেক্ষণ খরচ (টাকা)"
	banglaTranslations["Length (mm)"] = "দৈর্ঘ্য (মিমি)"
	banglaTranslations["Highway Fuel Economy (km/l)"] = "হাইওয়ে জ্বালানি অর্থনীতি (কিমি/লি)"
	banglaTranslations["Service Interval (km)"] = "পরিষেবা ব্যবধান (কিমি)"
	banglaTranslations["Driver Airbag"] = "ড্রাইভার এয়ারব্যাগ"
	banglaTranslations["Power Windows"] = "পাওয়ার উইন্ডোজ"
	banglaTranslations["Combined Fuel Economy (km/l)"] = "সমন্বিত জ্বালানি অর্থনীতি (কিমি/লি)"
	banglaTranslations["Spare Parts Availability"] = "খুচরা যন্ত্রাংশের প্রাপ্যতা"
	banglaTranslations["Cylinders"] = "সিলিন্ডার"
	banglaTranslations["Engine Displacement (cc)"] = "ইঞ্জিন স্থানচ্যুতি (সিসি)"
	banglaTranslations["Registration Cost (BDT)"] = "নিবন্ধন খরচ (টাকা)"
	banglaTranslations["Stability Control"] = "স্থিতিশীলতা নিয়ন্ত্রণ"
	banglaTranslations["Camera"] = "ক্যামেরা"
	banglaTranslations["City Fuel Economy (km/l)"] = "শহরের জ্বালানি অর্থনীতি (কিমি/লি)"
	banglaTranslations["Wheelbase (mm)"] = "হুইলবেস (মিমি)"
	banglaTranslations["Max Torque (Nm)"] = "সর্বাধিক টর্ক (এনএম)"
	banglaTranslations["Fuel Efficiency (km/l)"] = "জ্বালানি দক্ষতা (কিমি/লি)"
	banglaTranslations["Side Airbags"] = "পার্শ্ব এয়ারব্যাগ"
	banglaTranslations["AC Type"] = "এসি প্রকার"
	banglaTranslations["Insurance Cost (BDT)"] = "বীমা খরচ (টাকা)"
	banglaTranslations["Child Safety Lock"] = "শিশু নিরাপত্তা লক"
	banglaTranslations["Passenger Airbag"] = "যাত্রী এয়ারব্যাগ"
	banglaTranslations["Seat Material"] = "আসনের উপাদান"
	banglaTranslations["Fuel Tank Capacity (l)"] = "জ্বালানি ট্যাঙ্কের ক্ষমতা (লি)"
	banglaTranslations["Transmission Type"] = "ট্রান্সমিশন প্রকার"
	banglaTranslations["Curtain Airbags"] = "কার্টেন এয়ারব্যাগ"
	banglaTranslations["Height (mm)"] = "উচ্চতা (মিমি)"
	banglaTranslations["Width (mm)"] = "প্রস্থ (মিমি)"
	banglaTranslations["Max Power (hp)"] = "সর্বাধিক শক্তি (এইচপি)"

	// Insert Bangla translations
	for key, translation := range banglaTranslations {
		var specKey models.SpecificationKeyModel
		result := db.Where("specification_key = ?", key).First(&specKey)
		if result.Error != nil {
			continue // Skip if the specification key is not found
		}

		var existingTranslation models.SpecificationTranslationModel
		result = db.Where("specification_id = ? AND locale = ?", specKey.ID, "bn").First(&existingTranslation)

		if result.Error == gorm.ErrRecordNotFound {
			// Create new Bangla translation
			translation := models.SpecificationTranslationModel{
				SpecificationID: specKey.ID,
				Locale:          "bn",
				Value:           translation,
			}
			_ = db.Create(&translation)
		}
	}

	// Correcting field names in SpecificationKeyTranslationModel
	// Using the correct fields: `SpecificationKeyID` and `SpecificationKey`
	fmt.Println("Starting to seed missing specification key translations...")

	// Ensure all specification keys have translations
	var specificationKeys []models.SpecificationKeyModel
	if err := db.Where("id > ?", 584).Find(&specificationKeys).Error; err != nil {
		fmt.Printf("Error fetching specification keys: %v\n", err)
		return err
	}

	fmt.Printf("Fetched %d specification keys with ID > 584\n", len(specificationKeys))

	for _, specKey := range specificationKeys {
		fmt.Printf("Processing specification key: ID=%d, Key=%s\n", specKey.ID, specKey.SpecificationKey)

		// Check if English translation exists in `specification_key_translations`
		var existingTranslation models.SpecificationKeyTranslationModel
		result := db.Where("specification_key_id = ? AND locale = ?", specKey.ID, "en").First(&existingTranslation)

		if result.Error == gorm.ErrRecordNotFound {
			// Add English translation
			translation := models.SpecificationKeyTranslationModel{
				SpecificationKeyID: specKey.ID,
				Locale:             "en",
				SpecificationKey:   specKey.SpecificationKey, // Use the key as the default translation
			}
			if err := db.Create(&translation).Error; err != nil {
				fmt.Printf("Error creating English translation for ID=%d: %v\n", specKey.ID, err)
				continue
			}
			fmt.Printf("Inserted English translation for ID=%d\n", specKey.ID)
		} else if result.Error != nil {
			fmt.Printf("Error checking English translation for ID=%d: %v\n", specKey.ID, result.Error)
			continue
		}

		// Check if Bangla translation exists in `specification_key_translations`
		result = db.Where("specification_key_id = ? AND locale = ?", specKey.ID, "bn").First(&existingTranslation)

		if result.Error == gorm.ErrRecordNotFound {
			// Add Bangla translation if available
			banglaTranslation, exists := banglaTranslations[specKey.SpecificationKey]
			if exists {
				translation := models.SpecificationKeyTranslationModel{
					SpecificationKeyID: specKey.ID,
					Locale:             "bn",
					SpecificationKey:   banglaTranslation,
				}
				if err := db.Create(&translation).Error; err != nil {
					fmt.Printf("Error creating Bangla translation for ID=%d: %v\n", specKey.ID, err)
					continue
				}
				fmt.Printf("Inserted Bangla translation for ID=%d\n", specKey.ID)
			} else {
				fmt.Printf("No Bangla translation available for key: %s\n", specKey.SpecificationKey)
			}
		} else if result.Error != nil {
			fmt.Printf("Error checking Bangla translation for ID=%d: %v\n", specKey.ID, result.Error)
			continue
		}
	}

	fmt.Println("Finished seeding missing specification key translations.")

	return nil
}

// SpecificationKeyTranslation represents the translation of a specification key
type SpecificationKeyTranslation struct {
	ID               int       `gorm:"primaryKey"`
	SpecificationKey string    `gorm:"column:specification_key"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

// SeedSpecificationKeyTranslations seeds the specification key translations
func SeedSpecificationKeyTranslations(db *gorm.DB) error {
	translations := []SpecificationKeyTranslation{
		{ID: 613, SpecificationKey: "Estimated Annual Maintenance Cost (BDT)", CreatedAt: time.Date(2025, 11, 26, 10, 38, 0, 724641000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 38, 0, 724641000, time.UTC)},
		{ID: 612, SpecificationKey: "Length (mm)", CreatedAt: time.Date(2025, 11, 26, 10, 38, 0, 478137000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 38, 0, 478137000, time.UTC)},
		{ID: 611, SpecificationKey: "Highway Fuel Economy (km/l)", CreatedAt: time.Date(2025, 11, 26, 10, 38, 0, 229385000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 38, 0, 229385000, time.UTC)},
		{ID: 610, SpecificationKey: "Service Interval (km)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 59, 959252000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 59, 959252000, time.UTC)},
		{ID: 609, SpecificationKey: "Driver Airbag", CreatedAt: time.Date(2025, 11, 26, 10, 37, 59, 703289000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 59, 703289000, time.UTC)},
		{ID: 608, SpecificationKey: "Power Windows", CreatedAt: time.Date(2025, 11, 26, 10, 37, 59, 451283000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 59, 451283000, time.UTC)},
		{ID: 607, SpecificationKey: "Combined Fuel Economy (km/l)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 59, 187792000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 59, 187792000, time.UTC)},
		{ID: 606, SpecificationKey: "Spare Parts Availability", CreatedAt: time.Date(2025, 11, 26, 10, 37, 58, 936268000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 58, 936268000, time.UTC)},
		{ID: 605, SpecificationKey: "Cylinders", CreatedAt: time.Date(2025, 11, 26, 10, 37, 58, 683081000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 58, 683081000, time.UTC)},
		{ID: 604, SpecificationKey: "Engine Displacement (cc)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 58, 437565000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 58, 437565000, time.UTC)},
		{ID: 603, SpecificationKey: "Registration Cost (BDT)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 58, 172608000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 58, 172608000, time.UTC)},
		{ID: 602, SpecificationKey: "Stability Control", CreatedAt: time.Date(2025, 11, 26, 10, 37, 57, 914505000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 57, 914505000, time.UTC)},
		{ID: 601, SpecificationKey: "Camera", CreatedAt: time.Date(2025, 11, 26, 10, 37, 57, 658190000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 57, 658190000, time.UTC)},
		{ID: 600, SpecificationKey: "City Fuel Economy (km/l)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 57, 400539000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 57, 400539000, time.UTC)},
		{ID: 599, SpecificationKey: "Wheelbase (mm)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 57, 155499000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 57, 155499000, time.UTC)},
		{ID: 598, SpecificationKey: "Max Torque (Nm)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 56, 901563000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 56, 901563000, time.UTC)},
		{ID: 597, SpecificationKey: "Fuel Efficiency (km/l)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 56, 653796000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 56, 653796000, time.UTC)},
		{ID: 596, SpecificationKey: "Side Airbags", CreatedAt: time.Date(2025, 11, 26, 10, 37, 56, 406226000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 56, 406226000, time.UTC)},
		{ID: 595, SpecificationKey: "AC Type", CreatedAt: time.Date(2025, 11, 26, 10, 37, 56, 155246000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 56, 155246000, time.UTC)},
		{ID: 594, SpecificationKey: "Insurance Cost (BDT)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 55, 909202000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 55, 909202000, time.UTC)},
		{ID: 593, SpecificationKey: "Child Safety Lock", CreatedAt: time.Date(2025, 11, 26, 10, 37, 55, 658743000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 55, 658743000, time.UTC)},
		{ID: 592, SpecificationKey: "Passenger Airbag", CreatedAt: time.Date(2025, 11, 26, 10, 37, 55, 414315000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 55, 414315000, time.UTC)},
		{ID: 591, SpecificationKey: "Seat Material", CreatedAt: time.Date(2025, 11, 26, 10, 37, 55, 159681000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 55, 159681000, time.UTC)},
		{ID: 590, SpecificationKey: "Fuel Tank Capacity (l)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 54, 905848000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 54, 905848000, time.UTC)},
		{ID: 589, SpecificationKey: "Transmission Type", CreatedAt: time.Date(2025, 11, 26, 10, 37, 54, 662669000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 54, 662669000, time.UTC)},
		{ID: 588, SpecificationKey: "Curtain Airbags", CreatedAt: time.Date(2025, 11, 26, 10, 37, 54, 414404000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 54, 414404000, time.UTC)},
		{ID: 587, SpecificationKey: "Height (mm)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 54, 167032000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 54, 167032000, time.UTC)},
		{ID: 586, SpecificationKey: "Width (mm)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 53, 908895000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 53, 908895000, time.UTC)},
		{ID: 585, SpecificationKey: "Max Power (hp)", CreatedAt: time.Date(2025, 11, 26, 10, 37, 53, 612573000, time.UTC), UpdatedAt: time.Date(2025, 11, 26, 10, 37, 53, 612573000, time.UTC)},
	}

	return db.Create(&translations).Error
}
