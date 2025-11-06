package seeders

import (
	"kossti/internal/domain/entities"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// CategoryTranslationSeeder handles seeding category translations
type CategoryTranslationSeeder struct {
	BaseSeeder
}

// NewCategoryTranslationSeeder creates a new category translation seeder
func NewCategoryTranslationSeeder() *CategoryTranslationSeeder {
	return &CategoryTranslationSeeder{
		BaseSeeder: BaseSeeder{name: "Category Translations"},
	}
}

// Seed implements the Seeder interface
func (cts *CategoryTranslationSeeder) Seed(db *gorm.DB) error {
	translations := []struct {
		CategoryID uint
		Locale     string
		Name       string
	}{
		{1, "bn", "অ্যাক্সেসরিজ"},
		{2, "bn", "এয়ারোস্পেস"},
		{3, "bn", "এয়ারলাইন"},
		{4, "bn", "অ্যালকোহলিক পানীয়"},
		{5, "bn", "অডিও সরঞ্জাম"},
		{6, "bn", "অটোমোটিভ"},
		{7, "bn", "বেবি পণ্য"},
		{8, "bn", "ব্যাগ এবং লাগেজ"},
		{9, "bn", "বেকারি"},
		{10, "bn", "ব্যাংকিং"},
		{11, "bn", "ব্যাটারি"},
		{12, "bn", "সৌন্দর্য পণ্য"},
		{13, "bn", "বিয়ার"},
		{14, "bn", "পানীয়"},
		{15, "bn", "বই"},
		{16, "bn", "বাল্ব"},
		{17, "bn", "ক্যামেরা"},
		{18, "bn", "গাড়ি"},
		{19, "bn", "ক্যাজুয়াল পরিধান"},
		{20, "bn", "শ্যাম্পেন"},
		{21, "bn", "পোশাক"},
		{21, "bn", "পোশাক"},
		{22, "bn", "ক্লাউড স্টোরেজ"},
		{23, "bn", "কফি"},
		{24, "bn", "কগনাক"},
		{25, "bn", "কম্পিউটার"},
		{26, "bn", "কম্পিউটার হার্ডওয়্যার"},
		{27, "bn", "মিষ্টান্ন"},
		{28, "bn", "সমষ্টিগত"},
		{29, "bn", "পরামর্শ"},
		{30, "bn", "কুকিজ"},
		{31, "bn", "কসমেটিক্স"},
		{32, "bn", "কারুশিল্প সরবরাহ"},
		{33, "bn", "ক্রিস্টাল"},
		{34, "bn", "দুগ্ধজাত পণ্য"},
		{35, "bn", "ডেনিম"},
		{36, "bn", "বৈদ্যুতিক গাড়ি"},
		{37, "bn", "ইলেকট্রনিক্স"},
		{38, "bn", "শক্তি পানীয়"},
		{39, "bn", "বিনোদন"},
		{40, "bn", "চশমা"},
		{41, "bn", "ফ্যাশন"},
		{42, "bn", "ফাস্ট ফুড"},
		{43, "bn", "আর্থিক সেবা"},
		{44, "bn", "ফিটনেস"},
		{45, "bn", "ফিটনেস সরঞ্জাম"},
		{46, "bn", "ফ্লিপ-ফ্লপ"},
		{47, "bn", "খাদ্য ও পানীয়"},
		{48, "bn", "জুতা"},
		{49, "bn", "সুগন্ধি ও সুবাস"},
		{50, "bn", "আসবাবপত্র"},
		{51, "bn", "গেমিং"},
		{52, "bn", "বাগান এবং আউটডোর"},
		{53, "bn", "গ্লাস"},
		{54, "bn", "মুদি"},
		{55, "bn", "চুলের যত্ন"},
		{56, "bn", "স্বাস্থ্য এবং সৌন্দর্য"},
		{57, "bn", "স্বাস্থ্যসেবা"},
		{58, "bn", "গৃহস্থালী যন্ত্রপাতি"},
		{59, "bn", "আতিথেয়তা"},
		{60, "bn", "হোটেল"},
		{61, "bn", "গৃহস্থালী পণ্য"},
		{62, "bn", "হাইব্রিড গাড়ি"},
		{63, "bn", "আইসক্রিম"},
		{64, "bn", "শিল্প সরঞ্জাম"},
		{65, "bn", "বীমা"},
		{66, "bn", "ইন্টারনেট সেবা"},
		{67, "bn", "বিনিয়োগ"},
		{68, "bn", "গহনা"},
		{69, "bn", "রান্নাঘরের যন্ত্রপাতি"},
		{70, "bn", "রান্নাঘরের জিনিসপত্র"},
		{71, "bn", "ল্যাপটপ"},
		{72, "bn", "আলো"},
		{73, "bn", "লজিস্টিক্স"},
		{74, "bn", "বিলাসবহুল গাড়ি"},
		{75, "bn", "বিলাসবহুল পণ্য"},
		{76, "bn", "বিলাসবহুল সেবা"},
		{77, "bn", "মিডিয়া"},
		{78, "bn", "মেলামিন"},
		{79, "bn", "মোবাইল ফোন"},
		{80, "bn", "মোবাইল সেবা"},
		{81, "bn", "মোটরসাইকেল"},
		{82, "bn", "একাধিক শিল্প"},
		{83, "bn", "বাদ্যযন্ত্র"},
		{84, "bn", "বাদ্যযন্ত্র"},
		{85, "bn", "নেটওয়ার্কিং"},
		{86, "bn", "পুষ্টি"},
		{87, "bn", "অফিস সরবরাহ"},
		{88, "bn", "অনলাইন মার্কেটপ্লেস"},
		{89, "bn", "মুখের যত্ন"},
		{90, "bn", "আউটডোর গিয়ার"},
		{91, "bn", "প্যাকেজিং"},
		{92, "bn", "পেমেন্ট সেবা"},
		{93, "bn", "পারফিউম"},
		{94, "bn", "ব্যক্তিগত যত্ন"},
		{95, "bn", "পোষা প্রাণীর সরবরাহ"},
		{96, "bn", "ফার্মাসিউটিক্যালস"},
		{97, "bn", "ফটোগ্রাফি"},
		{98, "bn", "প্রিন্টার"},
		{99, "bn", "প্রিন্টিং"},
		{100, "bn", "রেফ্রিজারেটর"},
		{101, "bn", "রেস্তোরাঁ"},
		{102, "bn", "খুচরা"},
		{103, "bn", "সার্চ ইঞ্জিন"},
		{104, "bn", "সেমিকন্ডাক্টর"},
		{105, "bn", "শিপিং"},
		{106, "bn", "স্কেটবোর্ডিং"},
		{107, "bn", "স্কিনকেয়ার"},
		{108, "bn", "সামাজিক মিডিয়া"},
		{109, "bn", "নরম পানীয়"},
		{110, "bn", "সফটওয়্যার"},
		{111, "bn", "খেলাধুলা"},
		{112, "bn", "ক্রীড়া গাড়ি"},
		{113, "bn", "ক্রীড়া পানীয়"},
		{114, "bn", "খেলার সরঞ্জাম"},
		{115, "bn", "স্টেশনারি"},
		{116, "bn", "মিষ্টান্ন"},
		{117, "bn", "প্রযুক্তি"},
		{118, "bn", "টেলিকমিউনিকেশন"},
		{119, "bn", "থিম পার্ক"},
		{120, "bn", "টায়ার"},
		{121, "bn", "খেলনা"},
		{122, "bn", "ভ্রমণ"},
		{123, "bn", "ট্রাক"},
		{124, "bn", "টিভি"},
		{125, "bn", "ঘড়ি"},
		{126, "bn", "পরিধানযোগ্য"},
		{127, "bn", "হুইস্কি"},
		{128, "bn", "লেখার যন্ত্র"},
	}

	for _, translation := range translations {
		// Create CategoryTranslation entity
		categoryTranslation := &entities.CategoryTranslation{
			CategoryID:     translation.CategoryID,
			Locale:         translation.Locale,
			TranslatedName: translation.Name,
		}

		// Check if translation already exists
		var existingModel models.CategoryTranslationModel
		result := db.Where("category_id = ? AND locale = ?", translation.CategoryID, translation.Locale).First(&existingModel)

		if result.Error == gorm.ErrRecordNotFound {
			// Convert entity to model for database insertion
			model := &models.CategoryTranslationModel{
				CategoryID: categoryTranslation.CategoryID,
				Locale:     categoryTranslation.Locale,
				Name:       categoryTranslation.TranslatedName,
			}

			if err := db.Create(model).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
