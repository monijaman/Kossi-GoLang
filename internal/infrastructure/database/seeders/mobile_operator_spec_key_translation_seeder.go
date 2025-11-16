package seeders

import (
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// MobileOperatorSpecificationKeyTranslationSeeder handles seeding mobile operator specification key translations
type MobileOperatorSpecificationKeyTranslationSeeder struct {
	BaseSeeder
}

// NewMobileOperatorSpecificationKeyTranslationSeeder creates a new mobile operator specification key translation seeder
func NewMobileOperatorSpecificationKeyTranslationSeeder() *MobileOperatorSpecificationKeyTranslationSeeder {
	return &MobileOperatorSpecificationKeyTranslationSeeder{
		BaseSeeder: BaseSeeder{name: "Mobile Operator Specification Key Translations"},
	}
}

// Seed implements the Seeder interface
func (moskts *MobileOperatorSpecificationKeyTranslationSeeder) Seed(db *gorm.DB) error {
	// Specification key translations in Bengali and English
	translationMap := map[string]map[string]string{
		"Operator Name":                            {"bn": "অপারেটর নাম", "en": "Operator Name"},
		"Established Year":                         {"bn": "প্রতিষ্ঠার বছর", "en": "Established Year"},
		"Headquarters Location":                    {"bn": "সদর দপ্তরের অবস্থান", "en": "Headquarters Location"},
		"Network Type":                             {"bn": "নেটওয়ার্ক প্রকার", "en": "Network Type"},
		"Coverage Area":                            {"bn": "কভারেজ এলাকা", "en": "Coverage Area"},
		"Total Subscribers":                        {"bn": "মোট গ্রাহক সংখ্যা", "en": "Total Subscribers"},
		"Total Base Stations":                      {"bn": "মোট বেস স্টেশন", "en": "Total Base Stations"},
		"2G Service Available":                     {"bn": "২জি সেবা উপলব্ধ", "en": "2G Service Available"},
		"3G Service Available":                     {"bn": "৩জি সেবা উপলব্ধ", "en": "3G Service Available"},
		"4G/LTE Service Available":                 {"bn": "৪জি/এলটিই সেবা উপলব্ধ", "en": "4G/LTE Service Available"},
		"5G Service Available":                     {"bn": "৫জি সেবা উপলব্ধ", "en": "5G Service Available"},
		"Voice Call Service":                       {"bn": "ভয়েস কল সেবা", "en": "Voice Call Service"},
		"SMS Service":                              {"bn": "এসএমএস সেবা", "en": "SMS Service"},
		"Data Service":                             {"bn": "ডেটা সেবা", "en": "Data Service"},
		"MMS Service":                              {"bn": "এমএমএস সেবা", "en": "MMS Service"},
		"Internet Browsing":                        {"bn": "ইন্টারনেট ব্রাউজিং", "en": "Internet Browsing"},
		"Video Call Support":                       {"bn": "ভিডিও কল সমর্থন", "en": "Video Call Support"},
		"Roaming Service":                          {"bn": "রোমিং সেবা", "en": "Roaming Service"},
		"International Roaming":                    {"bn": "আন্তর্জাতিক রোমিং", "en": "International Roaming"},
		"Mobile Money Service":                     {"bn": "মোবাইল মানি সেবা", "en": "Mobile Money Service"},
		"Mobile Money Partner":                     {"bn": "মোবাইল মানি অংশীদার", "en": "Mobile Money Partner"},
		"USSD Support":                             {"bn": "ইউএসএসডি সমর্থন", "en": "USSD Support"},
		"Prepaid Plans Available":                  {"bn": "প্রি-পেইড প্ল্যান উপলব্ধ", "en": "Prepaid Plans Available"},
		"Postpaid Plans Available":                 {"bn": "পোস্ট-পেইড প্ল্যান উপলব্ধ", "en": "Postpaid Plans Available"},
		"Customer Care Phone":                      {"bn": "গ্রাহক সেবা ফোন", "en": "Customer Care Phone"},
		"Customer Care Email":                      {"bn": "গ্রাহক সেবা ইমেইল", "en": "Customer Care Email"},
		"Official Website":                         {"bn": "অফিসিয়াল ওয়েবসাইট", "en": "Official Website"},
		"Mobile App Name":                          {"bn": "মোবাইল অ্যাপ নাম", "en": "Mobile App Name"},
		"Android App Available":                    {"bn": "অ্যান্ড্রয়েড অ্যাপ উপলব্ধ", "en": "Android App Available"},
		"iOS App Available":                        {"bn": "আইওএস অ্যাপ উপলব্ধ", "en": "iOS App Available"},
		"Head Office Address":                      {"bn": "প্রধান অফিস ঠিকানা", "en": "Head Office Address"},
		"Helpline Hours":                           {"bn": "হেল্পলাইন সময়", "en": "Helpline Hours"},
		"Complaint Resolution Time":                {"bn": "অভিযোগ সমাধানের সময়", "en": "Complaint Resolution Time"},
		"Network Speed (Average)":                  {"bn": "নেটওয়ার্ক গতি (গড়)", "en": "Network Speed (Average)"},
		"Customer Satisfaction Rating":             {"bn": "গ্রাহক সন্তুষ্টির রেটিং", "en": "Customer Satisfaction Rating"},
		"CEO / MD Name":                            {"bn": "সিইও / এমডি নাম", "en": "CEO / MD Name"},
		"Chairman Name":                            {"bn": "চেয়ারম্যান নাম", "en": "Chairman Name"},
		"Number of Employees":                      {"bn": "কর্মচারী সংখ্যা", "en": "Number of Employees"},
		"Network Infrastructure Partners":          {"bn": "নেটওয়ার্ক অবকাঠামো অংশীদার", "en": "Network Infrastructure Partners"},
		"Technology Provider":                      {"bn": "প্রযুক্তি প্রদানকারী", "en": "Technology Provider"},
		"ISO Certifications":                       {"bn": "আইএসও সার্টিফিকেশন", "en": "ISO Certifications"},
		"Corporate Social Responsibility Programs": {"bn": "কর্পোরেট সামাজিক দায়বদ্ধতা কর্মসূচি", "en": "Corporate Social Responsibility Programs"},
		"Service Availability SLA":                 {"bn": "সেবা উপলব্ধতা এসএলএ", "en": "Service Availability SLA"},
		"Call Drop Rate":                           {"bn": "কল ড্রপ হার", "en": "Call Drop Rate"},
		"Voice Quality Rating":                     {"bn": "ভয়েস গুণমান রেটিং", "en": "Voice Quality Rating"},
		"Data Speed Rating":                        {"bn": "ডেটা গতি রেটিং", "en": "Data Speed Rating"},
		"Network Reliability":                      {"bn": "নেটওয়ার্ক নির্ভরযোগ্যতা", "en": "Network Reliability"},
		"Emergency Services Support":               {"bn": "জরুরি সেবা সমর্থন", "en": "Emergency Services Support"},
		"Disaster Recovery Support":                {"bn": "দুর্যোগ পুনরুদ্ধার সমর্থন", "en": "Disaster Recovery Support"},
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
