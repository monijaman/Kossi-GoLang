package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// SuzukiGSXS1000ReviewBatch16 handles seeding Suzuki GSX-S1000 product review and translation
type SuzukiGSXS1000ReviewBatch16 struct {
	BaseSeeder
}

// NewSuzukiGSXS1000ReviewBatch16Seeder creates a new SuzukiGSXS1000ReviewBatch16
func NewSuzukiGSXS1000ReviewBatch16Seeder() *SuzukiGSXS1000ReviewBatch16 {
	return &SuzukiGSXS1000ReviewBatch16{BaseSeeder: BaseSeeder{name: "Suzuki GSX-S1000 Batch16 Review"}}
}

// Seed implements the Seeder interface
func (s *SuzukiGSXS1000ReviewBatch16) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki GSX-S1000").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki gsx-s1000 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding suzuki gsx-s1000 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Suzuki GSX-S1000 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Suzuki GSX-S1000 Review Bangladesh 2024 - Litre-Class Performance Beast</h1>
<p class="summary-text">The Suzuki GSX-S1000 is a premium 1000cc naked superbike delivering extraordinary power, advanced technology, razor-sharp handling, and aggressive styling. Priced around ৳12,50,000, it's the ultimate street machine for experienced riders seeking extreme performance and cutting-edge engineering.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki GSX-S1000 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">1000cc Beast Engine:</strong> <span class="highlight-value">Extreme superbike power</span></li>
<li class="highlight-item"><strong class="highlight-label">Advanced Technology:</strong> <span class="highlight-value">Cutting-edge features</span></li>
<li class="highlight-item"><strong class="highlight-label">Sharp Handling:</strong> <span class="highlight-value">Track-tuned dynamics</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Styling:</strong> <span class="highlight-value">Premium performance looks</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki GSX-S1000 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Extreme Power:</strong> <span class="pro-description">1000cc thrilling</span></li>
<li class="pro-item"><strong class="pro-label">Advanced Tech:</strong> <span class="pro-description">ABS, traction control</span></li>
<li class="pro-item"><strong class="pro-label">Razor Handling:</strong> <span class="pro-description">Track-tuned sharp</span></li>
<li class="pro-item"><strong class="pro-label">Premium Build:</strong> <span class="pro-description">Excellent quality</span></li>
<li class="pro-item"><strong class="pro-label">Suzuki Heritage:</strong> <span class="pro-description">Racing pedigree</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki GSX-S1000 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Extreme Price:</strong> <span class="con-description">৳12,50,000 very expensive</span></li>
<li class="con-item"><strong class="con-label">Maintenance Cost:</strong> <span class="con-description">Very high</span></li>
<li class="con-item"><strong class="con-label">Insurance:</strong> <span class="con-description">Expensive coverage</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki GSX-S1000 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳12,50,000 - Premium superbike</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very high - ৳15-20 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">18-25 km/l low</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki GSX-S1000 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Power:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- 1000cc extreme</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Track sharp</span></li>
<li class="rating-item"><strong class="rating-label">Technology:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Cutting-edge</span></li>
<li class="rating-item"><strong class="rating-label">Build Quality:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Premium excellent</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Expensive premium</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki GSX-S1000?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳12,50,000, the Suzuki GSX-S1000 is for extreme performance enthusiasts seeking raw 1000cc power, track-tuned handling, cutting-edge technology, and the ultimate expression of litre-class motorcycle engineering.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For extreme performance seekers</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.8,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating suzuki gsx-s1000 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Suzuki GSX-S1000 (Rating: 4.8/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি GSX-S1000 রিভিউ বাংলাদেশ 2024 - লিটার-ক্লাস পারফরম্যান্স বিস্ট</h1>
<p class="summary-text">সুজুকি GSX-S1000 একটি প্রিমিয়াম 1000cc ন্যাকেড সুপারবাইক যা অসাধারণ শক্তি, উন্নত প্রযুক্তি, ধারালো হ্যান্ডলিং এবং আক্রমণাত্মক স্টাইলিং প্রদান করে। ৳12,50,000 টাকায় মূল্যায়িত, এটি অভিজ্ঞ রাইডারদের জন্য চূড়ান্ত স্ট্রিট মেশিন যারা চরম কর্মক্ষমতা এবং অত্যাধুনিক প্রকৌশল খুঁজছেন।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি GSX-S1000 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">1000cc বিস্ট ইঞ্জিন:</strong> <span class="highlight-value">চরম সুপারবাইক শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">উন্নত প্রযুক্তি:</strong> <span class="highlight-value">অত্যাধুনিক বৈশিষ্ট্য</span></li>
<li class="highlight-item"><strong class="highlight-label">ধারালো হ্যান্ডলিং:</strong> <span class="highlight-value">ট্র্যাক-টিউনড গতিশীলতা</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="highlight-value">প্রিমিয়াম কর্মক্ষমতা চেহারা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি GSX-S1000 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চরম শক্তি:</strong> <span class="pro-description">1000cc রোমাঞ্চকর</span></li>
<li class="pro-item"><strong class="pro-label">উন্নত প্রযুক্তি:</strong> <span class="pro-description">ABS, ট্র্যাকশন নিয়ন্ত্রণ</span></li>
<li class="pro-item"><strong class="pro-label">ধারালো হ্যান্ডলিং:</strong> <span class="pro-description">ট্র্যাক-টিউনড তীক্ষ্ণ</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম বিল্ড:</strong> <span class="pro-description">চমৎকার গুণমান</span></li>
<li class="pro-item"><strong class="pro-label">সুজুকি ঐতিহ্য:</strong> <span class="pro-description">রেসিং বংশপরম্পরা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি GSX-S1000 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">চরম মূল্য:</strong> <span class="con-description">৳12,50,000 অত্যন্ত ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ খরচ:</strong> <span class="con-description">অত্যন্ত উচ্চ</span></li>
<li class="con-item"><strong class="con-label">বীমা:</strong> <span class="con-description">ব্যয়বহুল কভারেজ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি GSX-S1000 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳12,50,000 টাকায়, সুজুকি GSX-S1000 চরম কর্মক্ষমতা উত্সাহীদের জন্য যারা কাঁচা 1000cc শক্তি, ট্র্যাক-টিউনড হ্যান্ডলিং, অত্যাধুনিক প্রযুক্তি এবং লিটার-ক্লাস মোটরসাইকেল প্রকৌশলের চূড়ান্ত অভিব্যক্তি খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- চরম কর্মক্ষমতা সন্ধানকারীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Suzuki GSX-S1000 already exists\n")
		return nil
	}

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating bangla translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Suzuki GSX-S1000\n")

	return nil
}

// GetName returns the seeder name
func (s *SuzukiGSXS1000ReviewBatch16) GetName() string {
	return "SuzukiGSXS1000ReviewBatch16"
}
