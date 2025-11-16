package migrations

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

// ConvertSpecificationsAfter9833ToBengali converts English values in specification_translations
// to Bengali for entries after ID 9833, excluding URLs
func ConvertSpecificationsAfter9833ToBengali(db *gorm.DB) error {
	fmt.Println("\n========== CONVERTING SPECIFICATIONS AFTER ID 9833 TO BENGALI ==========")

	// Comprehensive English to Bengali translation map
	translationMap := map[string]string{
		// Yes/No and Boolean values
		"Yes":   "হ্যাঁ",
		"No":    "না",
		"TRUE":  "সত্য",
		"FALSE": "মিথ্যা",
		"True":  "সত্য",
		"False": "মিথ্যা",

		// Availability
		"Available":     "উপলব্ধ",
		"Not Available": "উপলব্ধ নয়",
		"Unavailable":   "অনুপলব্ধ",

		// Support status
		"Supported":     "সমর্থিত",
		"Not Supported": "সমর্থিত নয়",
		"Unsupported":   "অসমর্থিত",

		// Active/Inactive
		"Active":   "সক্রিয়",
		"Inactive": "নিষ্ক্রিয়",

		// Enabled/Disabled
		"Enabled":  "সক্ষম",
		"Disabled": "অক্ষম",

		// Optional/Required
		"Optional": "ঐচ্ছিক",
		"Required": "প্রয়োজনীয়",

		// Common spec values
		"Standard":     "মান",
		"Premium":      "প্রিমিয়াম",
		"Basic":        "মৌলিক",
		"Advanced":     "উন্নত",
		"Professional": "পেশাদার",

		// Pricing tiers
		"Free": "বিনামূল্যে",
		"Paid": "পেইড",

		// Time periods
		"Monthly": "মাসিক",
		"Yearly":  "বার্ষিক",
		"Annual":  "বার্ষিক",
		"Daily":   "দৈনিক",
		"Weekly":  "সাপ্তাহিক",
		"Hourly":  "প্রতি ঘন্টায়",
		"24/7":    "२४/७",

		// Quantity/Availability
		"Limited":   "সীমিত",
		"Unlimited": "সীমাহীন",

		// Completeness
		"Full":       "সম্পূর্ণ",
		"Partial":    "আংশিক",
		"Complete":   "সম্পূর্ণ",
		"Incomplete": "অসম্পূর্ণ",

		// Connectivity/Status
		"Connected":    "সংযুক্ত",
		"Disconnected": "বিচ্ছিন্ন",
		"Online":       "অনলাইন",
		"Offline":      "অফলাইন",

		// Common features
		"Built-in": "অন্তর্নির্মিত",
		"External": "বাহ্যিক",
		"Internal": "অভ্যন্তরীণ",

		// More options
		"None":     "কোন",
		"All":      "সব",
		"Single":   "একক",
		"Multiple": "একাধিক",
		"Both":     "উভয়",
		"Either":   "যেকোনো একটি",

		// Speed/Performance
		"Fast":   "দ্রুত",
		"Slow":   "ধীর",
		"High":   "উচ্চ",
		"Low":    "নিম্ন",
		"Medium": "মধ্যম",

		// Quality
		"Good":      "ভালো",
		"Bad":       "খারাপ",
		"Excellent": "চমৎকার",
		"Poor":      "খাঁটি",

		// Size
		"Large": "বড়",
		"Small": "ছোট",

		// Color status
		"Red":    "লাল",
		"Green":  "সবুজ",
		"Blue":   "নীল",
		"Yellow": "হলুদ",
		"Black":  "কালো",
		"White":  "সাদা",

		// Generic
		"Default": "ডিফল্ট",
		"Custom":  "কাস্টম",
		"Other":   "অন্যান্য",

		// Telecom & Network
		"2G":    "२जी",
		"3G":    "३जी",
		"4G":    "४जी",
		"LTE":   "एलटीई",
		"5G":    "५जी",
		"WiFi":  "वाईफाई",
		"Wi-Fi": "वाईफाई",

		// Locations (Bengali names)
		"Dhaka":      "ঢাকা",
		"Bangladesh": "বাংলাদেশ",
		"Chittagong": "চট্টগ্রাম",
		"Sylhet":     "সিলেট",
		"Rajshahi":   "রাজশাহী",
		"Khulna":     "খুলনা",
		"Barisal":    "বরিশাল",
		"Rangpur":    "রংপুর",
		"Mymensingh": "ময়মনসিংহ",

		// Brand names
		"Grameenphone": "গ্রামীণফোন",
		"Vodafone":     "ভোডাফোন",
		"Robi":         "রবি",
		"Banglalink":   "বাংলালিংক",
		"Teletalk":     "টেলিটক",
		"Airtel":       "এয়ারটেল",
	}

	fmt.Println("\n📊 Step 1: Finding specifications after ID 9833...")

	type SpecTranslation struct {
		ID     uint
		Value  string
		Locale string
	}

	var specs []SpecTranslation
	result := db.Raw(`
		SELECT id, value, locale 
		FROM specification_translations 
		WHERE id > 9833 
		AND locale = 'bn'
		ORDER BY id ASC
	`).Scan(&specs)

	if result.Error != nil {
		return fmt.Errorf("error querying specifications: %w", result.Error)
	}

	fmt.Printf("Found %d entries after ID 9833\n", len(specs))

	// Helper function to check if string is a URL
	isURL := func(s string) bool {
		s = strings.ToLower(strings.TrimSpace(s))
		return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") || strings.HasPrefix(s, "www.")
	}

	// Helper function to convert numbers to Bengali numerals
	convertNumbersToBengali := func(s string) string {
		// Map English digits to Bengali digits (Bengali numerals)
		englishDigits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		bengaliDigits := []string{"০", "১", "२", "३", "४", "५", "६", "७", "८", "९"}

		result := s
		for i := 0; i < 10; i++ {
			result = strings.ReplaceAll(result, englishDigits[i], bengaliDigits[i])
		}
		return result
	}

	// Helper function to check if already in Bengali
	hasBengaliChars := func(s string) bool {
		return strings.ContainsAny(s, "অআইঈউঊঋএঐওঔকখগঘঙচছজঝঞটঠডঢণতথদধনপফবভমযরলশষসহড়ঢ়য়ৎংঃঁ")
	}

	// Process translations
	fmt.Println("\n🔄 Step 2: Converting English values to Bengali...")
	updatedCount := 0
	numberConvertedCount := 0
	skippedURLCount := 0
	skippedBengaliCount := 0
	skippedUnknownCount := 0

	for _, spec := range specs {
		// Skip URLs
		if isURL(spec.Value) {
			skippedURLCount++
			continue
		}

		translatedValue := spec.Value
		found := false

		// Try exact match first (case-insensitive)
		for englishKey, bengaliValue := range translationMap {
			if strings.EqualFold(strings.TrimSpace(spec.Value), englishKey) {
				translatedValue = bengaliValue
				found = true
				break
			}
		}

		// If found and different from original, update
		if found && translatedValue != spec.Value {
			updateResult := db.Exec(
				"UPDATE specification_translations SET value = ? WHERE id = ?",
				translatedValue, spec.ID,
			)
			if updateResult.Error != nil {
				log.Printf("Error updating ID %d: %v", spec.ID, updateResult.Error)
			} else if updateResult.RowsAffected > 0 {
				updatedCount++
				fmt.Printf("  ✓ ID %d: '%s' → '%s'\n", spec.ID, spec.Value, translatedValue)
			}
		} else if !found {
			// Check if contains English digits and convert them to Bengali
			if regexp.MustCompile(`\d`).MatchString(spec.Value) && !hasBengaliChars(spec.Value) {
				bengaliVersion := convertNumbersToBengali(spec.Value)
				if bengaliVersion != spec.Value {
					updateResult := db.Exec(
						"UPDATE specification_translations SET value = ? WHERE id = ?",
						bengaliVersion, spec.ID,
					)
					if updateResult.Error != nil {
						log.Printf("Error updating ID %d: %v", spec.ID, updateResult.Error)
					} else if updateResult.RowsAffected > 0 {
						numberConvertedCount++
						fmt.Printf("  ✓ ID %d (Numbers): '%s' → '%s'\n", spec.ID, spec.Value, bengaliVersion)
					}
				}
			} else if hasBengaliChars(spec.Value) {
				skippedBengaliCount++
			} else {
				skippedUnknownCount++
			}
		}
	}

	fmt.Printf("\n📝 Conversion Summary:\n")
	fmt.Printf("  ✓ Direct Translation: %d entries\n", updatedCount)
	fmt.Printf("  ✓ Number Conversion: %d entries\n", numberConvertedCount)
	fmt.Printf("  ⊘ Skipped (URLs): %d entries\n", skippedURLCount)
	fmt.Printf("  ⊘ Skipped (Already Bengali): %d entries\n", skippedBengaliCount)
	fmt.Printf("  ⊘ Skipped (No translation found): %d entries\n", skippedUnknownCount)
	fmt.Printf("  → Total processed: %d entries\n", len(specs))

	fmt.Println("\n✅ Conversion completed successfully!")
	return nil
}
