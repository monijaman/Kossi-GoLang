package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVSpecificationKeyTranslationSeeder handles seeding TV specification key translations
type TVSpecificationKeyTranslationSeeder struct {
	BaseSeeder
}

// NewTVSpecificationKeyTranslationSeeder creates a new TV specification key translation seeder
func NewTVSpecificationKeyTranslationSeeder() *TVSpecificationKeyTranslationSeeder {
	return &TVSpecificationKeyTranslationSeeder{
		BaseSeeder: BaseSeeder{name: "TV Specification Key Translations"},
	}
}

// Seed implements the Seeder interface
func (tvskts *TVSpecificationKeyTranslationSeeder) Seed(db *gorm.DB) error {
	// Specification key translations in Bengali and English
	translationMap := map[string]map[string]string{
		"Brand":                             {"bn": "ব্র্যান্ড", "en": "Brand"},
		"Model Name":                        {"bn": "মডেল নাম", "en": "Model Name"},
		"Screen Size":                       {"bn": "স্ক্রীন সাইজ", "en": "Screen Size"},
		"Resolution":                        {"bn": "রেজোলিউশন", "en": "Resolution"},
		"Panel Type":                        {"bn": "প্যানেল প্রকার", "en": "Panel Type"},
		"Refresh Rate":                      {"bn": "রিফ্রেশ রেট", "en": "Refresh Rate"},
		"Brightness (Nits)":                 {"bn": "উজ্জ্বলতা (নিটস)", "en": "Brightness (Nits)"},
		"Contrast Ratio":                    {"bn": "কন্ট্রাস্ট অনুপাত", "en": "Contrast Ratio"},
		"Color Accuracy":                    {"bn": "রঙের নির্ভুলতা", "en": "Color Accuracy"},
		"Viewing Angle":                     {"bn": "দৃশ্য কোণ", "en": "Viewing Angle"},
		"Response Time":                     {"bn": "প্রতিক্রিয়া সময়", "en": "Response Time"},
		"HDR Support":                       {"bn": "এইচডিআর সমর্থন", "en": "HDR Support"},
		"Dolby Vision":                      {"bn": "ডলবি ভিশন", "en": "Dolby Vision"},
		"Local Dimming":                     {"bn": "স্থানীয় ম্লান করা", "en": "Local Dimming"},
		"Backlight Type":                    {"bn": "ব্যাকলাইট প্রকার", "en": "Backlight Type"},
		"Smart TV OS":                       {"bn": "স্মার্ট টিভি অপারেটিং সিস্টেম", "en": "Smart TV OS"},
		"Built-in Apps":                     {"bn": "অন্তর্নির্মিত অ্যাপ্লিকেশন", "en": "Built-in Apps"},
		"WiFi Connectivity":                 {"bn": "ওয়াইফাই সংযোগ", "en": "WiFi Connectivity"},
		"Bluetooth Connectivity":            {"bn": "ব্লুটুথ সংযোগ", "en": "Bluetooth Connectivity"},
		"HDMI Ports":                        {"bn": "এইচডিএমআই পোর্ট", "en": "HDMI Ports"},
		"USB Ports":                         {"bn": "ইউএসবি পোর্ট", "en": "USB Ports"},
		"Audio Output (Watts)":              {"bn": "অডিও আউটপুট (ওয়াট)", "en": "Audio Output (Watts)"},
		"Speaker Configuration":             {"bn": "স্পিকার কনফিগারেশন", "en": "Speaker Configuration"},
		"Dolby Atmos Support":               {"bn": "ডলবি অ্যাটমস সমর্থন", "en": "Dolby Atmos Support"},
		"Built-in Tuner":                    {"bn": "অন্তর্নির্মিত টিউনার", "en": "Built-in Tuner"},
		"Supported Video Formats":           {"bn": "সমর্থিত ভিডিও ফরম্যাট", "en": "Supported Video Formats"},
		"Supported Audio Formats":           {"bn": "সমর্থিত অডিও ফরম্যাট", "en": "Supported Audio Formats"},
		"Frame Rate Support":                {"bn": "ফ্রেম রেট সমর্থন", "en": "Frame Rate Support"},
		"ALLM Support":                      {"bn": "এএলএলএম সমর্থন", "en": "ALLM Support"},
		"VRR Support":                       {"bn": "ভিআরআর সমর্থন", "en": "VRR Support"},
		"Game Mode":                         {"bn": "গেম মোড", "en": "Game Mode"},
		"Motion Smoothing":                  {"bn": "গতি মসৃণকরণ", "en": "Motion Smoothing"},
		"Picture Quality Modes":             {"bn": "ছবির গুণমান মোড", "en": "Picture Quality Modes"},
		"Energy Consumption (Watts)":        {"bn": "শক্তি খরচ (ওয়াট)", "en": "Energy Consumption (Watts)"},
		"Power Consumption Standby (Watts)": {"bn": "স্ট্যান্ডবাই শক্তি খরচ (ওয়াট)", "en": "Power Consumption Standby (Watts)"},
		"Dimensions (W x H x D)":            {"bn": "মাত্রা (প্রস্থ x উচ্চতা x গভীরতা)", "en": "Dimensions (W x H x D)"},
		"Weight Without Stand (kg)":         {"bn": "স্ট্যান্ড ছাড়া ওজন (কেজি)", "en": "Weight Without Stand (kg)"},
		"Weight With Stand (kg)":            {"bn": "স্ট্যান্ড সহ ওজন (কেজি)", "en": "Weight With Stand (kg)"},
		"VESA Mount Pattern":                {"bn": "ভিইএসএ মাউন্ট প্যাটার্ন", "en": "VESA Mount Pattern"},
		"Warranty Period (Years)":           {"bn": "ওয়ারেন্টি সময়কাল (বছর)", "en": "Warranty Period (Years)"},
		"Warranty Coverage":                 {"bn": "ওয়ারেন্টি কভারেজ", "en": "Warranty Coverage"},
		"Lifespan (Hours)":                  {"bn": "জীবনকাল (ঘন্টা)", "en": "Lifespan (Hours)"},
		"Thickness":                         {"bn": "পুরুত্ব", "en": "Thickness"},
		"Bezel Width":                       {"bn": "বেজেল প্রস্থ", "en": "Bezel Width"},
		"Remote Type":                       {"bn": "রিমোট প্রকার", "en": "Remote Type"},
		"Voice Control Support":             {"bn": "ভয়েস নিয়ন্ত্রণ সমর্থন", "en": "Voice Control Support"},
		"Gaming Features":                   {"bn": "গেমিং বৈশিষ্ট্য", "en": "Gaming Features"},
		"Sports Mode":                       {"bn": "খেলাধুলা মোড", "en": "Sports Mode"},
		"Movie Mode":                        {"bn": "মুভি মোড", "en": "Movie Mode"},
		"Eye Comfort":                       {"bn": "চোখের আরাম", "en": "Eye Comfort"},
		"Flicker-Free":                      {"bn": "ফ্লিকার মুক্ত", "en": "Flicker-Free"},
		"Blue Light Filter":                 {"bn": "নীল আলো ফিল্টার", "en": "Blue Light Filter"},
		"Eco Mode":                          {"bn": "ইকো মোড", "en": "Eco Mode"},
		"Auto Brightness":                   {"bn": "স্বয়ংক্রিয় উজ্জ্বলতা", "en": "Auto Brightness"},
		"Motion Interpolation":              {"bn": "গতি ইন্টারপোলেশন", "en": "Motion Interpolation"},
		"Upscaling Technology":              {"bn": "আপস্কেলিং প্রযুক্তি", "en": "Upscaling Technology"},
		"Color Enhancement":                 {"bn": "রঙ উন্নতি", "en": "Color Enhancement"},
		"Price (BDT)":                       {"bn": "মূল্য (টাকা)", "en": "Price (BDT)"},
		"Availability in Bangladesh":        {"bn": "বাংলাদেশে উপলব্ধতা", "en": "Availability in Bangladesh"},
		"After-Sales Service":               {"bn": "বিক্রয়োত্তর সেবা", "en": "After-Sales Service"},
		"Return / Exchange Policy":          {"bn": "রিটার্ন / এক্সচেঞ্জ নীতি", "en": "Return / Exchange Policy"},
	}

	// Create translations for each specification key
	for specKeyName, translations := range translationMap {
		// Find specification key
		var specKey models.SpecificationKeyModel
		if err := db.Where("specification_key = ?", specKeyName).First(&specKey).Error; err != nil {
			continue // Skip if spec key not found
		}

		// Create translations for each locale
		for locale, translatedName := range translations {
			// Check if translation already exists
			var existingTranslation models.SpecificationKeyTranslationModel
			result := db.Where("specification_key_id = ? AND locale = ?", specKey.ID, locale).First(&existingTranslation)

			if result.Error == gorm.ErrRecordNotFound {
				// Create new translation
				translation := models.SpecificationKeyTranslationModel{
					SpecificationKeyID: specKey.ID,
					Locale:             locale,
					SpecificationKey:   translatedName,
				}

				if err := db.Create(&translation).Error; err != nil {
					continue
				}
			}
		}
	}

	return nil
}
