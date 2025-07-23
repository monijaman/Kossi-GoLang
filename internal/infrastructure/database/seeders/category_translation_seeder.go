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
		{1, "bn", "এয়ারলাইন"},
		{2, "bn", "অটোমোটিভ"},
		{3, "bn", "বেবি পণ্য"},
		{4, "bn", "ব্যাগ এবং লাগেজ"},
		{5, "bn", "বই"},
		{6, "bn", "ক্যামেরা"},
		{7, "bn", "পোশাক"},
		{8, "bn", "কসমেটিক্স"},
		{9, "bn", "কারুশিল্প সরবরাহ"},
		{10, "bn", "ইলেকট্রনিক্স"},
		{11, "bn", "ফিটনেস সরঞ্জাম"},
		{12, "bn", "জুতা"},
		{13, "bn", "আসবাবপত্র"},
		{14, "bn", "গেমিং"},
		{15, "bn", "বাগান এবং আউটডোর"},
		{16, "bn", "মুদি"},
		{17, "bn", "স্বাস্থ্য এবং সৌন্দর্য"},
		{18, "bn", "গৃহস্থালী যন্ত্রপাতি"},
		{19, "bn", "গহনা"},
		{20, "bn", "রান্নাঘরের জিনিসপত্র"},
		{21, "bn", "ল্যাপটপ"},
		{22, "bn", "মোবাইল ফোন"},
		{23, "bn", "বাদ্যযন্ত্র"},
		{24, "bn", "অফিস সরবরাহ"},
		{25, "bn", "পারফিউম"},
		{26, "bn", "ব্যক্তিগত যত্ন"},
		{27, "bn", "পোষা প্রাণীর সরবরাহ"},
		{28, "bn", "রেফ্রিজারেটর"},
		{29, "bn", "খেলাধুলা"},
		{30, "bn", "স্টেশনারি"},
		{31, "bn", "টিভি"},
		{32, "bn", "খেলনা"},
		{33, "bn", "ঘড়ি"},
		{34, "bn", "খেলার সরঞ্জাম"},
		{35, "bn", "বিলাসবহুল পণ্য"},
		{36, "bn", "অ্যাক্সেসরিজ"},
		{37, "bn", "পানীয়"},
		{38, "bn", "অ্যালকোহলিক পানীয়"},
		{39, "bn", "বিয়ার"},
		{40, "bn", "ফটোগ্রাফি"},
		{41, "bn", "ফ্যাশন"},
		{42, "bn", "সুগন্ধি ও সুবাস"},
		{43, "bn", "গাড়ি"},
		{44, "bn", "নেটওয়ার্কিং"},
		{45, "bn", "টেলিকমিউনিকেশন"},
		{46, "bn", "নরম পানীয়"},
		{47, "bn", "মুখের যত্ন"},
		{48, "bn", "লজিস্টিক্স"},
		{49, "bn", "শিপিং"},
		{50, "bn", "বিনোদন"},
		{51, "bn", "মিডিয়া"},
		{52, "bn", "থিম পার্ক"},
		{53, "bn", "সৌন্দর্য পণ্য"},
		{54, "bn", "ক্লাউড স্টোরেজ"},
		{55, "bn", "অনলাইন মার্কেটপ্লেস"},
		{56, "bn", "ভ্রমণ"},
		{57, "bn", "সামাজিক মিডিয়া"},
		{58, "bn", "ট্রাক"},
		{59, "bn", "ক্রীড়া পানীয়"},
		{60, "bn", "সার্চ ইঞ্জিন"},
		{61, "bn", "সফটওয়্যার"},
		{62, "bn", "খুচরা"},
		{63, "bn", "কম্পিউটার"},
		{64, "bn", "প্রিন্টার"},
		{65, "bn", "মোটরসাইকেল"},
		{66, "bn", "পরামর্শ"},
		{67, "bn", "কম্পিউটার হার্ডওয়্যার"},
		{68, "bn", "সেমিকন্ডাক্টর"},
		{69, "bn", "বিলাসবহুল গাড়ি"},
		{70, "bn", "স্বাস্থ্যসেবা"},
		{71, "bn", "ফার্মাসিউটিক্যালস"},
		{72, "bn", "ব্যাংকিং"},
		{73, "bn", "বিনিয়োগ"},
		{74, "bn", "খাদ্য ও পানীয়"},
		{75, "bn", "ফাস্ট ফুড"},
		{76, "bn", "রেস্তোরাঁ"},
		{77, "bn", "ডেনিম"},
		{78, "bn", "আর্থিক সেবা"},
		{79, "bn", "পেমেন্ট সেবা"},
		{80, "bn", "শিল্প সরঞ্জাম"},
		{81, "bn", "কনফেকশনারি"},
		{82, "bn", "পুষ্টি"},
		{83, "bn", "স্কিনকেয়ার"},
		{84, "bn", "আলো"},
		{85, "bn", "ক্রীড়া গাড়ি"},
		{86, "bn", "গৃহস্থালী পণ্য"},
		{87, "bn", "শক্তি পানীয়"},
		{88, "bn", "কফি"},
		{89, "bn", "মোবাইল সেবা"},
		{90, "bn", "বৈদ্যুতিক গাড়ি"},
		{91, "bn", "প্রযুক্তি"},
		{92, "bn", "আউটডোর গিয়ার"},
		{93, "bn", "হাইব্রিড গাড়ি"},
		{94, "bn", "ইন্টারনেট সেবা"},
		{95, "bn", "রান্নাঘরের যন্ত্রপাতি"},
		{96, "bn", "বীমা"},
		{97, "bn", "আইসক্রিম"},
		{98, "bn", "মিষ্টান্ন"},
		{99, "bn", "ক্যাজুয়াল পরিধান"},
		{100, "bn", "দুগ্ধজাত পণ্য"},
		{101, "bn", "ব্যাটারি"},
		{102, "bn", "ফিটনেস"},
		{103, "bn", "পরিধানযোগ্য"},
		{104, "bn", "ফ্লিপ-ফ্লপ"},
		{105, "bn", "লেখার যন্ত্র"},
		{106, "bn", "শ্যাম্পেন"},
		{107, "bn", "কুকিজ"},
		{108, "bn", "চশমা"},
		{109, "bn", "কগনাক"},
		{110, "bn", "হোটেল"},
		{111, "bn", "আতিথেয়তা"},
		{112, "bn", "বিলাসবহুল সেবা"},
		{113, "bn", "ক্রিস্টাল"},
		{114, "bn", "স্কেটবোর্ডিং"},
		{115, "bn", "চুলের যত্ন"},
		{116, "bn", "বাদ্যযন্ত্র"},
		{117, "bn", "একাধিক শিল্প"},
		{118, "bn", "সমষ্টিগত"},
		{119, "bn", "এয়ারোস্পেস"},
		{120, "bn", "অডিও সরঞ্জাম"},
		{121, "bn", "টায়ার"},
		{122, "bn", "হুইস্কি"},
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
