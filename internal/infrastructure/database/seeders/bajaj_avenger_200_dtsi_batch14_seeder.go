package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajAvenger200DTSiReview handles seeding Bajaj Avenger 200 DTS-i product review and translation
type BajajAvenger200DTSiReview struct {
	BaseSeeder
}

// NewBajajAvenger200DTSiReviewSeeder creates a new BajajAvenger200DTSiReview
func NewBajajAvenger200DTSiReviewSeeder() *BajajAvenger200DTSiReview {
	return &BajajAvenger200DTSiReview{BaseSeeder: BaseSeeder{name: "Bajaj Avenger 200 DTS-i Review"}}
}

// Seed implements the Seeder interface
func (s *BajajAvenger200DTSiReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Avenger 200 DTS-i").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj avenger 200 dts-i product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bajaj avenger 200 dts-i product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Avenger 200 DTS-i already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Avenger 200 DTS-i Review Bangladesh 2024 - Cruiser Performance Bike</h1>
<p class="summary-text">The Bajaj Avenger 200 DTS-i is a power-packed 200cc cruiser with twin-spark technology, muscle-bike styling, commanding road presence, and excellent torque. Priced around ৳2,95,000, it delivers thrilling performance with impressive acceleration.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Avenger 200 DTS-i Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc Twin-Spark:</strong> <span class="highlight-value">Powerful performance engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Cruiser Design:</strong> <span class="highlight-value">Muscle-bike styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Torque:</strong> <span class="highlight-value">Strong acceleration</span></li>
<li class="highlight-item"><strong class="highlight-label">Road Presence:</strong> <span class="highlight-value">Commanding appearance</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Avenger 200 DTS-i Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Powerful Engine:</strong> <span class="pro-description">200cc twin-spark delivery</span></li>
<li class="pro-item"><strong class="pro-label">Strong Torque:</strong> <span class="pro-description">Impressive acceleration</span></li>
<li class="pro-item"><strong class="pro-label">Muscle Design:</strong> <span class="pro-description">Aggressive cruiser styling</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Performance:</strong> <span class="pro-description">Bajaj tested technology</span></li>
<li class="pro-item"><strong class="pro-label">Affordable Power:</strong> <span class="pro-description">Good value for performance</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Avenger 200 DTS-i Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Fuel Consumption:</strong> <span class="con-description">Higher consumption rate</span></li>
<li class="con-item"><strong class="con-label">Seating Comfort:</strong> <span class="con-description">Long rides less comfortable</span></li>
<li class="con-item"><strong class="con-label">Vibration:</strong> <span class="con-description">Noticeable at high revs</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Avenger 200 DTS-i Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,95,000 - Cruiser performance</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">40-45 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Avenger 200 DTS-i Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Power:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- 200cc twin-spark</span></li>
<li class="rating-item"><strong class="rating-label">Acceleration:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Strong torque delivery</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Muscle cruiser</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Bajaj proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Good performance value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Avenger 200 DTS-i?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,95,000, the Bajaj Avenger 200 DTS-i is perfect for riders seeking a powerful cruiser with muscle styling, strong torque, and thrilling performance for exciting urban riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For cruiser performance seekers</span></p>
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
		return fmt.Errorf("error creating bajaj avenger 200 dts-i review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Avenger 200 DTS-i (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ Avenger 200 DTS-i রিভিউ বাংলাদেশ 2024 - ক্রুজার পারফরম্যান্স বাইক</h1>
<p class="summary-text">বাজাজ Avenger 200 DTS-i একটি শক্তিশালী 200cc ক্রুজার যা টুইন-স্পার্ক প্রযুক্তি, মাসল-বাইক স্টাইলিং, কমান্ডিং রোড উপস্থিতি এবং চমৎকার টর্ক সহ আসে। ৳2,95,000 টাকায় মূল্যায়িত, এটি প্রভাবশালী ত্বরণ সহ রোমাঞ্চকর কর্মক্ষমতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ Avenger 200 DTS-i মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc টুইন-স্পার্ক:</strong> <span class="highlight-value">শক্তিশালী কর্মক্ষমতা ইঞ্জিন</span></li>
<li class="highlight-item"><strong class="highlight-label">ক্রুজার ডিজাইন:</strong> <span class="highlight-value">মাসল-বাইক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার টর্ক:</strong> <span class="highlight-value">শক্তিশালী ত্বরণ</span></li>
<li class="highlight-item"><strong class="highlight-label">রোড উপস্থিতি:</strong> <span class="highlight-value">কমান্ডিং চেহারা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ Avenger 200 DTS-i সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী ইঞ্জিন:</strong> <span class="pro-description">200cc টুইন-স্পার্ক সরবরাহ</span></li>
<li class="pro-item"><strong class="pro-label">শক্তিশালী টর্ক:</strong> <span class="pro-description">প্রভাবশালী ত্বরণ</span></li>
<li class="pro-item"><strong class="pro-label">মাসল ডিজাইন:</strong> <span class="pro-description">আক্রমণাত্মক ক্রুজার স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য কর্মক্ষমতা:</strong> <span class="pro-description">বাজাজ পরীক্ষিত প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী শক্তি:</strong> <span class="pro-description">কর্মক্ষমতার জন্য ভাল মূল্য</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ Avenger 200 DTS-i অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">জ্বালানি খরচ:</strong> <span class="con-description">উচ্চতর খরচ হার</span></li>
<li class="con-item"><strong class="con-label">আসনের আরাম:</strong> <span class="con-description">দীর্ঘ যাত্রা কম আরামদায়ক</span></li>
<li class="con-item"><strong class="con-label">কম্পন:</strong> <span class="con-description">উচ্চ বিপ্লব লক্ষণীয়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ Avenger 200 DTS-i কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,95,000 টাকায়, বাজাজ Avenger 200 DTS-i শক্তিশালী ক্রুজার, মাসল স্টাইলিং এবং রোমাঞ্চকর শহুরে যাত্রার জন্য শক্তিশালী টর্ক খুঁজছেন এমন রাইডারদের জন্য নিখুঁত।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ক্রুজার কর্মক্ষমতা খোঁজারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Avenger 200 DTS-i already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Avenger 200 DTS-i\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajAvenger200DTSiReview) GetName() string {
	return "BajajAvenger200DTSiReview"
}
