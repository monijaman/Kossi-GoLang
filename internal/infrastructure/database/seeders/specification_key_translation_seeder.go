package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SpecificationKeyTranslationSeeder handles seeding specification key translations
type SpecificationKeyTranslationSeeder struct {
	BaseSeeder
}

// NewSpecificationKeyTranslationSeeder creates a new specification key translation seeder
func NewSpecificationKeyTranslationSeeder() *SpecificationKeyTranslationSeeder {
	return &SpecificationKeyTranslationSeeder{
		BaseSeeder: BaseSeeder{name: "Specification Key Translations"},
	}
}

// Seed implements the Seeder interface
func (skts *SpecificationKeyTranslationSeeder) Seed(db *gorm.DB) error {
	translations := map[string]string{
		"Technology":                          "প্রযুক্তি",
		"2G Bands":                            "২জি ব্যান্ড",
		"3G Bands":                            "৩জি ব্যান্ড",
		"4G Bands":                            "৪জি ব্যান্ড",
		"Network Speed":                       "নেটওয়ার্ক স্পিড",
		"Announcement Date":                   "ঘোষণা তারিখ",
		"Device Status":                       "ডিভাইসের অবস্থা",
		"Dimensions":                          "আয়তন",
		"Weight":                              "ওজন",
		"Build Material":                      "নির্মাণের উপকরণ",
		"SIM Card Type":                       "সিম কার্ডের ধরন",
		"Display Type":                        "ডিসপ্লে টাইপ",
		"Screen Size":                         "স্ক্রীনের আকার",
		"Resolution":                          "রেজুলেশন",
		"Operating System":                    "অপারেটিং সিস্টেম",
		"Chipset":                             "চিপসেট",
		"CPU Type":                            "সিপিইউ টাইপ",
		"GPU Type":                            "জিপিইউ টাইপ",
		"Card Slot Type":                      "কার্ড স্লট টাইপ",
		"Internal Memory Capacity":            "আভ্যন্তরীণ মেমোরির ধারণক্ষমতা",
		"Dual SIM":                            "ডুয়াল সিম",
		"Special Features":                    "বিশেষ বৈশিষ্ট্য",
		"Main Camera Video Resolution":        "মেইন ক্যামেরার ভিডিও রেজুলেশন",
		"Front Camera Resolution":             "ফ্রন্ট ক্যামেরার রেজুলেশন",
		"Front Camera Video Resolution":       "ফ্রন্ট ক্যামেরার ভিডিও রেজুলেশন",
		"Loudspeaker Quality":                 "লাউডস্পিকার গুণমান",
		"3.5mm Audio Jack":                    "৩.৫মিমি অডিও জ্যাক",
		"WiFi":                                "ওয়াইফাই",
		"Bluetooth Version":                   "ব্লুটুথ সংস্করণ",
		"Positioning System":                  "অবস্থান নির্ধারণের ব্যবস্থা",
		"NFC Support":                         "এনএফসি সমর্থন",
		"Radio Support":                       "রেডিও সমর্থন",
		"USB Type":                            "ইউএসবি টাইপ",
		"Sensors":                             "সেন্সর",
		"Battery Capacity":                    "ব্যাটারির ধারণক্ষমতা",
		"Charging Speed":                      "চার্জিং গতি",
		"Available Colors":                    "উপলব্ধ রং",
		"Model Variants":                      "মডেল ভ্যারিয়েন্ট",
		"SAR (Specific Absorption Rate) EU":   "এসএআর (বিশেষ শোষণ হার) ইউরোপীয় ইউনিয়ন",
		"Price":                               "মূল্য",
		"Triple Camera Setup":                 "ট্রিপল ক্যামেরা সেটআপ",
		"SAR (Specific Absorption Rate)":      "এসএআর (বিশেষ শোষণ হার)",
		"5G Bands":                            "৫জি ব্যান্ড",
		"Performance Metrics":                 "কার্যক্ষমতার পরিমাপক",
		"Display Characteristics":             "ডিসপ্লের বৈশিষ্ট্য",
		"New Battery Capacity":                "নতুন ব্যাটারির ধারণক্ষমতা",
		"Screen Protection":                   "স্ক্রীন সুরক্ষা",
		"Camera Features":                     "ক্যামেরার বৈশিষ্ট্য",
		"Old Battery Capacity":                "পুরাতন ব্যাটারির ধারণক্ষমতা",
		"Quad Camera Setup":                   "কোয়াড ক্যামেরা সেটআপ",
		"Audio Quality":                       "অডিও গুণমান",
		"GPRS (General Packet Radio Service)": "জিপিআরএস (সাধারণ প্যাকেট রেডিও পরিষেবা)",
		"EDGE (Enhanced Data Rates for GSM Evolution)": "এডিজি (জিএসএম বিবর্তনের জন্য উন্নত ডেটা হার)",
		"Talk Time Duration":                           "কথা বলার সময়কাল",
		"Music Playback Duration":                      "সঙ্গীত প্লেব্যাক সময়কাল",
		"Standby Time":                                 "স্ট্যান্ডবাই সময়",
		"Infrared Port":                                "ইনফ্রারেড পোর্ট",
		"Device Type":                                  "ডিভাইসের ধরন",
		"Phonebook Capacity":                           "ফোনবুক ধারণক্ষমতা",
		"Call Records":                                 "কল রেকর্ড",
		"Messaging Features":                           "মেসেজিং বৈশিষ্ট্য",
		"Preinstalled Games":                           "পূর্ব-স্থাপিত গেম",
		"Java Support":                                 "জাভা সমর্থন",
		"Penta Camera Setup":                           "পেন্টা ক্যামেরা সেটআপ",
		"Dual or Triple Camera Setup":                  "ডুয়াল বা ট্রিপল ক্যামেরা সেটআপ",
		"Web Browser":                                  "ওয়েব ব্রাউজার",
		"Alert Types":                                  "সতর্কতা প্রকার",
		"Physical Keyboard":                            "শারীরিক কীবোর্ড",
		"Device Size":                                  "ডিভাইসের আকার",
		"Clock Feature":                                "ঘড়ির বৈশিষ্ট্য",
		"Alarm Feature":                                "অ্যালার্ম বৈশিষ্ট্য",
		"No Support":                                   "কোনো সমর্থন নেই",
		"Supported Languages":                          "সমর্থিত ভাষাসমূহ",
	}

	for englishKey, bengaliTranslation := range translations {
		// Find the specification key by its English name
		var specKeyModel models.SpecificationKeyModel
		result := db.Where("specification_key = ?", englishKey).First(&specKeyModel)

		if result.Error == nil {
			// Check if translation already exists
			var existingTranslation models.SpecificationKeyTranslationModel
			translationResult := db.Where("specification_key_id = ? AND locale = ?", specKeyModel.ID, "bn").First(&existingTranslation)

			if translationResult.Error == gorm.ErrRecordNotFound {
				// Create new translation
				translation := &models.SpecificationKeyTranslationModel{
					SpecificationKeyID: specKeyModel.ID,
					Locale:             "bn",
					SpecificationKey:   bengaliTranslation,
				}

				if err := db.Create(translation).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}
