package migrations

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

// TranslateEnglishSpecifications translates remaining English values in specification_translations table to Bengali
// This handles entries added after ID 9808 that may have English values
func TranslateEnglishSpecifications(db *gorm.DB) error {
	fmt.Println("\n========== TRANSLATING ENGLISH SPECIFICATIONS TO BENGALI ==========")

	// English to Bengali translation map for common specification values
	translationMap := map[string]string{
		// Common English words/phrases that need Bengali translation
		"Yes":           "হ্যাঁ",
		"No":            "না",
		"Available":     "উপলব্ধ",
		"Not Available": "উপলব্ধ নয়",
		"Supported":     "সমর্থিত",
		"Not Supported": "সমর্থিত নয়",
		"Active":        "সক্রিয়",
		"Inactive":      "নিষ্ক্রিয়",
		"Enabled":       "সক্ষম",
		"Disabled":      "অক্ষম",
		"True":          "সত্য",
		"False":         "মিথ্যা",
		"None":          "কোন",
		"Optional":      "ঐচ্ছিক",
		"Required":      "প্রয়োজনীয়",
		"Standard":      "মান",
		"Premium":       "প্রিমিয়াম",
		"Basic":         "মৌলিক",
		"Advanced":      "উন্নত",
		"Free":          "বিনামূল্যে",
		"Paid":          "পেইড",
		"Monthly":       "মাসিক",
		"Yearly":        "বার্ষিক",
		"Daily":         "দৈনিক",
		"Weekly":        "সাপ্তাহিক",
		"Hourly":        "প্রতি ঘন্টায়",
		"24/7":          "২৪/৭",
		"Limited":       "সীমিত",
		"Unlimited":     "সীমাহীন",
		"Full":          "সম্পূর্ণ",
		"Partial":       "আংশিক",
		"Complete":      "সম্পূর্ণ",
		"Incomplete":    "অসম্পূর্ণ",
	}

	fmt.Println("\n📊 Step 1: Finding English values in specification_translations (ID > 9808)...")

	// Query to find English values that need translation
	type SpecTranslation struct {
		ID     uint
		Value  string
		Locale string
	}

	var englishSpecs []SpecTranslation
	result := db.Raw(`
		SELECT id, value, locale 
		FROM specification_translations 
		WHERE id > 9808 
		AND locale = 'bn'
		ORDER BY id DESC
	`).Scan(&englishSpecs)

	if result.Error != nil {
		return fmt.Errorf("error querying specifications: %w", result.Error)
	}

	if len(englishSpecs) == 0 {
		fmt.Println("✓ No specification translations found after ID 9808")
		return nil
	}

	fmt.Printf("Found %d entries after ID 9808\n", len(englishSpecs))

	// Find English values that need translation
	fmt.Println("\n🔍 Step 2: Identifying English values to translate...")
	var needsTranslation []SpecTranslation
	for _, spec := range englishSpecs {
		// Check if value contains any English words
		isEnglish := false
		lowerValue := strings.ToLower(spec.Value)

		// Check against common English words
		englishWords := []string{"yes", "no", "available", "not available", "supported", "not supported",
			"active", "inactive", "enabled", "disabled", "true", "false", "optional", "required",
			"standard", "premium", "basic", "advanced", "free", "paid", "monthly", "yearly",
			"daily", "weekly", "hourly", "24/7", "limited", "unlimited", "full", "partial",
			"complete", "incomplete"}

		for _, word := range englishWords {
			if strings.Contains(lowerValue, word) {
				isEnglish = true
				break
			}
		}

		if isEnglish {
			needsTranslation = append(needsTranslation, spec)
		}
	}

	if len(needsTranslation) == 0 {
		fmt.Println("✓ No English values found that need translation")
		return nil
	}

	fmt.Printf("Found %d entries with English values to translate:\n", len(needsTranslation))
	for i, spec := range needsTranslation {
		if i < 20 {
			fmt.Printf("  ID %d: '%s'\n", spec.ID, spec.Value)
		}
	}
	if len(needsTranslation) > 20 {
		fmt.Printf("  ... and %d more\n", len(needsTranslation)-20)
	}

	// Step 3: Translate the values
	fmt.Println("\n🔄 Step 3: Translating values...")
	updatedCount := 0

	for _, spec := range needsTranslation {
		translatedValue := spec.Value

		// Try exact match first (case-insensitive)
		for englishKey, bengaliValue := range translationMap {
			if strings.EqualFold(spec.Value, englishKey) {
				translatedValue = bengaliValue
				break
			}
		}

		// Update the record if translated
		if translatedValue != spec.Value {
			updateResult := db.Exec(
				"UPDATE specification_translations SET value = ? WHERE id = ?",
				translatedValue, spec.ID,
			)
			if updateResult.Error != nil {
				log.Printf("Error updating ID %d: %v", spec.ID, updateResult.Error)
			} else if updateResult.RowsAffected > 0 {
				updatedCount++
				fmt.Printf("✓ ID %d: '%s' → '%s'\n", spec.ID, spec.Value, translatedValue)
			}
		}
	}

	fmt.Printf("\n✅ Translated %d values to Bengali\n", updatedCount)
	fmt.Println("=================================================================\n")

	return nil
}
