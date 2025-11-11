package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// MahindraClusterReview handles seeding Mahindra Cluster 110 product review and translation
type MahindraClusterReview struct {
	BaseSeeder
}

// NewMahindraClusterReviewSeeder creates a new MahindraClusterReview
func NewMahindraClusterReviewSeeder() *MahindraClusterReview {
	return &MahindraClusterReview{BaseSeeder: BaseSeeder{name: "Mahindra Cluster 110 Review"}}
}

// Seed implements the Seeder interface
func (s *MahindraClusterReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Mahindra Cluster 110").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("mahindra cluster 110 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding mahindra cluster 110 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Mahindra Cluster 110 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Mahindra Cluster 110 Review Bangladesh 2024 - Digital Commuter</h1>
<p class="summary-text">The Mahindra Cluster 110 is an innovative 110cc commuter bike featuring advanced digital cluster, efficient engine, stylish design, and modern technology. Priced at ৳1,72,000, it combines practicality with contemporary features for tech-savvy riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Mahindra Cluster 110 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Digital Cluster:</strong> <span class="highlight-value">Advanced instrumentation</span></li>
<li class="highlight-item"><strong class="highlight-label">110cc Engine:</strong> <span class="highlight-value">Smooth efficient commuting</span></li>
<li class="highlight-item"><strong class="highlight-label">Modern Design:</strong> <span class="highlight-value">Sleek contemporary styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Great Mileage:</strong> <span class="highlight-value">55-60 km/l excellent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Mahindra Cluster 110 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Digital Cluster:</strong> <span class="pro-description">Modern instrumentation</span></li>
<li class="pro-item"><strong class="pro-label">Efficient Engine:</strong> <span class="pro-description">Smooth 110cc performance</span></li>
<li class="pro-item"><strong class="pro-label">Stylish Design:</strong> <span class="pro-description">Contemporary looks</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Mileage:</strong> <span class="pro-description">55-60 km/l efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Tech Features:</strong> <span class="pro-description">Modern commuter bike</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Mahindra Cluster 110 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Brand Awareness:</strong> <span class="con-description">Less known in market</span></li>
<li class="con-item"><strong class="con-label">Service Network:</strong> <span class="con-description">Limited compared to others</span></li>
<li class="con-item"><strong class="con-label">Parts Availability:</strong> <span class="con-description">Harder to find spares</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Mahindra Cluster 110 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,72,000 - Digital commuter bike</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very low - ৳3-4 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">55-60 km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Mahindra Cluster 110 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Smoothness:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Smooth 110cc</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- 55-60 km/l</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Modern contemporary</span></li>
<li class="rating-item"><strong class="rating-label">Technology:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Digital cluster</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Good features</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Mahindra Cluster 110?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,72,000, the Mahindra Cluster 110 is ideal for tech-savvy riders seeking digital instrumentation, modern design, excellent fuel efficiency, and reliable commuting in a package that stands out with contemporary features.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For modern tech-forward riders</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.5,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating mahindra cluster 110 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Mahindra Cluster 110 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">মাহিন্দ্রা Cluster 110 রিভিউ বাংলাদেশ 2024 - ডিজিটাল কমিউটার</h1>
<p class="summary-text">মাহিন্দ্রা Cluster 110 একটি উদ্ভাবনী 110cc কমিউটার বাইক যা উন্নত ডিজিটাল ক্লাস্টার, দক্ষ ইঞ্জিন, স্টাইলিশ ডিজাইন এবং আধুনিক প্রযুক্তি বৈশিষ্ট্যযুক্ত। ৳1,72,000 টাকায় মূল্যায়িত, এটি প্রযুক্তি-সচেতন রাইডারদের জন্য ব্যবহারিকতার সাথে সমসাময়িক বৈশিষ্ট্য সমন্বয় করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">মাহিন্দ্রা Cluster 110 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">ডিজিটাল ক্লাস্টার:</strong> <span class="highlight-value">উন্নত যন্ত্র মাপকাঠি</span></li>
<li class="highlight-item"><strong class="highlight-label">110cc ইঞ্জিন:</strong> <span class="highlight-value">মসৃণ দক্ষ যাতায়াত</span></li>
<li class="highlight-item"><strong class="highlight-label">আধুনিক ডিজাইন:</strong> <span class="highlight-value">চিকন সমসাময়িক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">দুর্দান্ত মাইলেজ:</strong> <span class="highlight-value">55-60 km/l চমৎকার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">মাহিন্দ্রা Cluster 110 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ডিজিটাল ক্লাস্টার:</strong> <span class="pro-description">আধুনিক যন্ত্র মাপকাঠি</span></li>
<li class="pro-item"><strong class="pro-label">দক্ষ ইঞ্জিন:</strong> <span class="pro-description">মসৃণ 110cc কর্মক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">স্টাইলিশ ডিজাইন:</strong> <span class="pro-description">সমসাময়িক চেহারা</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার মাইলেজ:</strong> <span class="pro-description">55-60 km/l দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">প্রযুক্তি বৈশিষ্ট্য:</strong> <span class="pro-description">আধুনিক কমিউটার বাইক</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">মাহিন্দ্রা Cluster 110 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">ব্র্যান্ড সচেতনতা:</strong> <span class="con-description">বাজারে কম পরিচিত</span></li>
<li class="con-item"><strong class="con-label">সেবা নেটওয়ার্ক:</strong> <span class="con-description">অন্যদের তুলনায় সীমিত</span></li>
<li class="con-item"><strong class="con-label">যন্ত্রাংশ উপলব্ধতা:</strong> <span class="con-description">খুঁজে পাওয়া কঠিন</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: মাহিন্দ্রা Cluster 110 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,72,000 টাকায়, মাহিন্দ্রা Cluster 110 প্রযুক্তি-সচেতন রাইডারদের জন্য আদর্শ যারা ডিজিটাল যন্ত্র মাপকাঠি, আধুনিক ডিজাইন, চমৎকার জ্বালানী দক্ষতা এবং নির্ভরযোগ্য যাতায়াত খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আধুনিক প্রযুক্তি-সচেতন রাইডারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Mahindra Cluster 110 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Mahindra Cluster 110\n")

	return nil
}

// GetName returns the seeder name
func (s *MahindraClusterReview) GetName() string {
	return "MahindraClusterReview"
}
