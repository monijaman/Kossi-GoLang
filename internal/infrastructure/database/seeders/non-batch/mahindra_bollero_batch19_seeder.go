package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// MahindraBolleroReviewBatch19 handles seeding Mahindra Bollero product review and translation
type MahindraBolleroReviewBatch19 struct {
	BaseSeeder
}

// NewMahindraBolleroReviewBatch19Seeder creates a new MahindraBolleroReviewBatch19
func NewMahindraBolleroReviewBatch19Seeder() *MahindraBolleroReviewBatch19 {
	return &MahindraBolleroReviewBatch19{BaseSeeder: BaseSeeder{name: "Mahindra Bollero Batch19 Review"}}
}

// Seed implements the Seeder interface
func (s *MahindraBolleroReviewBatch19) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Mahindra Bollero").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("mahindra bollero product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding mahindra bollero product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Mahindra Bollero already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Mahindra Bollero Review Bangladesh 2024 - Budget Three-Wheeler Practical Transport</h1>
<p class="summary-text">The Mahindra Bollero is a 110cc air-cooled budget three-wheeler optimized for commercial and personal transport. Priced around ৳2,45,000, it combines affordability, generous cargo capacity, simple mechanics, and reliable Indian engineering for small business operators and rural transport seekers valuing practicality and cost-efficiency.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Mahindra Bollero Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">110cc Engine:</strong> <span class="highlight-value">Simple reliable</span></li>
<li class="highlight-item"><strong class="highlight-label">Cargo Capacity:</strong> <span class="highlight-value">300kg generous</span></li>
<li class="highlight-item"><strong class="highlight-label">Budget Price:</strong> <span class="highlight-value">৳2,45,000 accessible</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Economy:</strong> <span class="highlight-value">40-45 km/l excellent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Mahindra Bollero Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Affordable Price:</strong> <span class="pro-description">৳2,45,000 budget</span></li>
<li class="pro-item"><strong class="pro-label">Cargo Space:</strong> <span class="pro-description">300kg capacity</span></li>
<li class="pro-item"><strong class="pro-label">Simple Mechanics:</strong> <span class="pro-description">Easy maintenance</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Engine:</strong> <span class="pro-description">110cc proven</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Efficient:</strong> <span class="pro-description">40-45 km/l excellent</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Mahindra Bollero Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Three-Wheeler Class:</strong> <span class="con-description">Limited comfort</span></li>
<li class="con-item"><strong class="con-label">Speed Limitation:</strong> <span class="con-description">60 km/h modest</span></li>
<li class="con-item"><strong class="con-label">Weather Protection:</strong> <span class="con-description">Basic canopy</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Mahindra Bollero Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,45,000 - Budget transport</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Ultra-low - ৳5-7 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">40-45 km/l excellent</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">110cc air-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">5 bhp simple utility</span></p>
<p class="value-item"><strong class="value-label">Cargo Capacity:</strong> <span class="value-amount">300kg generous transport</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Mahindra Bollero Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Affordability:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- ৳2,45,000 good</span></li>
<li class="rating-item"><strong class="rating-label">Cargo Utility:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- 300kg capacity</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Economy:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- 40-45 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Simple proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Transport champion</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Mahindra Bollero?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,45,000, the Bollero is perfect for small business operators and rural transport seekers valuing generous 300kg cargo capacity, exceptional 40-45 km/l fuel efficiency, simple mechanics, and reliable Indian engineering for practical, cost-efficient daily transport.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For commercial transport</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.7,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating mahindra bollero review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Mahindra Bollero (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">মহিন্দ্রা বলেরো রিভিউ বাংলাদেশ 2024 - বাজেট তিন-চাকা ব্যবহারিক পরিবহন</h1>
<p class="summary-text">মহিন্দ্রা বলেরো একটি 110cc এয়ার-কুল্ড বাজেট তিন-চাকা বাণিজ্যিক এবং ব্যক্তিগত পরিবহনের জন্য অপ্টিমাইজ করা। ৳2,45,000 টাকায় মূল্যায়িত, এটি সাশ্রয়যোগ্যতা এবং উদার কার্গো ক্ষমতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">মহিন্দ্রা বলেরো মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">110cc ইঞ্জিন:</strong> <span class="highlight-value">সরল নির্ভরযোগ্য</span></li>
<li class="highlight-item"><strong class="highlight-label">কার্গো ক্ষমতা:</strong> <span class="highlight-value">300kg উদার</span></li>
<li class="highlight-item"><strong class="highlight-label">বাজেট মূল্য:</strong> <span class="highlight-value">৳2,45,000 অ্যাক্সেসযোগ্য</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি অর্থনীতি:</strong> <span class="highlight-value">40-45 km/l চমৎকার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">মহিন্দ্রা বলেরো সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">৳2,45,000 বাজেট</span></li>
<li class="pro-item"><strong class="pro-label">কার্গো স্থান:</strong> <span class="pro-description">300kg ক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">সরল মেকানিক্স:</strong> <span class="pro-description">সহজ রক্ষণাবেক্ষণ</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="pro-description">110cc প্রমাণিত</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি দক্ষ:</strong> <span class="pro-description">40-45 km/l চমৎকার</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">মহিন্দ্রা বলেরো অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">তিন-চাকা শ্রেণী:</strong> <span class="con-description">সীমিত আরাম</span></li>
<li class="con-item"><strong class="con-label">গতি সীমাবদ্ধতা:</strong> <span class="con-description">60 km/h বিনম্র</span></li>
<li class="con-item"><strong class="con-label">আবহাওয়া সুরক্ষা:</strong> <span class="con-description">মৌলিক ক্যানোপি</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: মহিন্দ্রা বলেরো কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,45,000 টাকায়, বলেরো ছোট ব্যবসা পরিচালক এবং গ্রামীণ পরিবহন সন্ধানকারীদের জন্য নিখুঁত যারা উদার 300kg কার্গো ক্ষমতা এবং ব্যবহারিক বাণিজ্যিক সমাধান চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাণিজ্যিক পরিবহনের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Mahindra Bollero already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Mahindra Bollero\n")

	return nil
}

// GetName returns the seeder name
func (s *MahindraBolleroReviewBatch19) GetName() string {
	return "MahindraBolleroReviewBatch19"
}
