package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type CFMotoNK600ReviewBatch20 struct {
	BaseSeeder
}

func NewCFMotoNK600ReviewBatch20Seeder() *CFMotoNK600ReviewBatch20 {
	return &CFMotoNK600ReviewBatch20{BaseSeeder: BaseSeeder{name: "CFMoto NK600 Batch20 Review"}}
}

func (s *CFMotoNK600ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "CFMoto NK600").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("cfmoto nk600 product not found")
		}
		return fmt.Errorf("error finding product: %w", err)
	}

	var existingReview models.ProductReviewModel
	if err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error; err == nil {
		fmt.Printf("   ℹ️  Review already exists, skipping\n")
		return nil
	}

	englishReview := `<article><header><h1>CFMoto NK600 Review - Affordable Naked Roadster</h1><p>The CFMoto NK600 is a 600cc naked roadster representing value-for-money Chinese engineering. Priced around ৳8,50,000, it delivers 45 bhp power, modern styling, practical ergonomics, good handling, and exceptional value for riders seeking affordable performance and daily practicality.</p></header><section><h2>Key Highlights</h2><ul><li>600cc Liquid-Cooled Engine</li><li>Modern Naked Design</li><li>45 bhp Practical Power</li><li>Value-For-Money ৳8,50,000</li></ul></section><section><h2>Performance Rating</h2><ul><li>Value: 4.8/5 - Exceptional price</li><li>Engine: 4.6/5 - Smooth 45 bhp</li><li>Handling: 4.6/5 - Responsive nimble</li><li>Design: 4.5/5 - Modern practical</li><li>Reliability: 4.5/5 - Proven build</li></ul></section><section><h2>Verdict</h2><p>At ৳8,50,000, NK600 offers exceptional value for practical riders wanting affordable 600cc performance with modern styling and reliable handling.</p></section></article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.6,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article><header><h1>সিএফমোটো NK600 রিভিউ - সাশ্রয়ী নেকেড রোডস্টার</h1><p>সিএফমোটো NK600 একটি 600cc নেকেড রোডস্টার যা মূল্য-সাশ্রয়ী চাইনিজ ইঞ্জিনিয়ারিং প্রতিনিধিত্ব করে। ৳8,50,000 টাকায় মূল্যায়িত, এটি ব্যবহারিক কর্মক্ষমতা এবং আধুনিক স্টাইলিং প্রদান করে।</p></header><section><h2>চূড়ান্ত সিদ্ধান্ত</h2><p>৳8,50,000 টাকায়, NK600 ব্যবহারিক রাইডারদের জন্য ব্যতিক্রমী মূল্য প্রদান করে যারা সাশ্রয়ী 600cc পারফরম্যান্স খুঁজছেন।</p><p>সুপারিশ: হ্যাঁ - মূল্য-সচেতন কর্মক্ষমতা চাহনাকারীদের জন্য</p></section></article>`

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating translation: %w", err)
	}

	fmt.Printf("   ✅ Created CFMoto NK600\n")
	return nil
}

func (s *CFMotoNK600ReviewBatch20) GetName() string {
	return "CFMotoNK600ReviewBatch20"
}
