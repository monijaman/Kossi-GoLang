package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HeroSplendorPlusDTSiReview handles seeding Hero Splendor Plus DTS-i product review and translation
type HeroSplendorPlusDTSiReview struct {
	BaseSeeder
}

// NewHeroSplendorPlusDTSiReviewSeeder creates a new HeroSplendorPlusDTSiReview
func NewHeroSplendorPlusDTSiReviewSeeder() *HeroSplendorPlusDTSiReview {
	return &HeroSplendorPlusDTSiReview{BaseSeeder: BaseSeeder{name: "Hero Splendor Plus DTS-i Review"}}
}

// Seed implements the Seeder interface
func (s *HeroSplendorPlusDTSiReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero Splendor Plus DTS-i").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero splendor plus dts-i product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding hero splendor plus dts-i product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Hero Splendor Plus DTS-i already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Hero Splendor Plus DTS-i Review Bangladesh 2024 - Budget Commuter Champion</h1>
<p class="summary-text">The Hero Splendor Plus DTS-i is a dependable 100cc commuter with twin-spark technology, fuel efficiency champion, reliable performance, and budget-friendly pricing. Priced around ৳1,65,000, it's India's best-selling commuter bike.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero Splendor Plus DTS-i Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">100cc Twin-Spark:</strong> <span class="highlight-value">Efficient commuter power</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficiency:</strong> <span class="highlight-value">55-60 km/l best-in-class</span></li>
<li class="highlight-item"><strong class="highlight-label">Low Price:</strong> <span class="highlight-value">Most affordable option</span></li>
<li class="highlight-item"><strong class="highlight-label">Proven Reliability:</strong> <span class="highlight-value">Millions sold worldwide</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero Splendor Plus DTS-i Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Excellent Mileage:</strong> <span class="pro-description">55-60 km/l best fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Low Price:</strong> <span class="pro-description">Most budget-friendly</span></li>
<li class="pro-item"><strong class="pro-label">Proven Reliability:</strong> <span class="pro-description">Bestselling model globally</span></li>
<li class="pro-item"><strong class="pro-label">Twin-Spark Technology:</strong> <span class="pro-description">Better performance</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance:</strong> <span class="pro-description">Simple, robust design</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero Splendor Plus DTS-i Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Basic Design:</strong> <span class="con-description">Simple, no-frills styling</span></li>
<li class="con-item"><strong class="con-label">Limited Features:</strong> <span class="con-description">Minimal modern tech</span></li>
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">100cc limited performance</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero Splendor Plus DTS-i Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,65,000 - Most affordable</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very Low - ৳3 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">55-60 km/l best-in-class</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero Splendor Plus DTS-i Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Proven bestseller</span></li>
<li class="rating-item"><strong class="rating-label">Fuel Efficiency:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- 55-60 km/l best</span></li>
<li class="rating-item"><strong class="rating-label">Price Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Most affordable</span></li>
<li class="rating-item"><strong class="rating-label">Maintenance:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Simple design</span></li>
<li class="rating-item"><strong class="rating-label">Overall Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Best budget champion</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero Splendor Plus DTS-i?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳1,65,000, the Hero Splendor Plus DTS-i is the ultimate budget commuter choice with world-leading fuel efficiency, proven reliability, low cost of ownership, and maximum value for money.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- Best budget commuter choice</span></p>
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
		return fmt.Errorf("error creating hero splendor plus dts-i review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Hero Splendor Plus DTS-i (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হিরো Splendor Plus DTS-i রিভিউ বাংলাদেশ 2024 - বাজেট কমিউটার চ্যাম্পিয়ন</h1>
<p class="summary-text">হিরো Splendor Plus DTS-i একটি নির্ভরযোগ্য 100cc কমিউটার যা টুইন-স্পার্ক প্রযুক্তি, জ্বালানি দক্ষতা চ্যাম্পিয়ন, নির্ভরযোগ্য কর্মক্ষমতা এবং বাজেট-বান্ধব মূল্য সহ আসে। ৳1,65,000 টাকায় মূল্যায়িত, এটি বিশ্বের সবচেয়ে বেশি বিক্রিত কমিউটার বাইক।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো Splendor Plus DTS-i মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">100cc টুইন-স্পার্ক:</strong> <span class="highlight-value">দক্ষ কমিউটার শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষতা:</strong> <span class="highlight-value">55-60 km/l সেরা-ইন-ক্লাস</span></li>
<li class="highlight-item"><strong class="highlight-label">কম মূল্য:</strong> <span class="highlight-value">সবচেয়ে সাশ্রয়ী বিকল্প</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রমাণিত নির্ভরযোগ্যতা:</strong> <span class="highlight-value">বিশ্বব্যাপী লক্ষ লক্ষ বিক্রি</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো Splendor Plus DTS-i সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">চমৎকার মাইলেজ:</strong> <span class="pro-description">55-60 km/l সেরা জ্বালানি দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">কম মূল্য:</strong> <span class="pro-description">সবচেয়ে বাজেট-বান্ধব</span></li>
<li class="pro-item"><strong class="pro-label">প্রমাণিত নির্ভরযোগ্যতা:</strong> <span class="pro-description">বিশ্বব্যাপী বেস্টসেলার মডেল</span></li>
<li class="pro-item"><strong class="pro-label">টুইন-স্পার্ক প্রযুক্তি:</strong> <span class="pro-description">ভাল কর্মক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সহজ, শক্তিশালী ডিজাইন</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো Splendor Plus DTS-i অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মৌলিক ডিজাইন:</strong> <span class="con-description">সহজ, কোন-প্রস্তাবনা স্টাইলিং</span></li>
<li class="con-item"><strong class="con-label">সীমিত বৈশিষ্ট্য:</strong> <span class="con-description">ন্যূনতম আধুনিক প্রযুক্তি</span></li>
<li class="con-item"><strong class="con-label">কম শক্তি:</strong> <span class="con-description">100cc সীমিত কর্মক্ষমতা</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো Splendor Plus DTS-i কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳1,65,000 টাকায়, হিরো Splendor Plus DTS-i চূড়ান্ত বাজেট কমিউটার পছন্দ যা বিশ্ব-নেতৃস্থানীয় জ্বালানি দক্ষতা, প্রমাণিত নির্ভরযোগ্যতা এবং সর্বোচ্চ অর্থের জন্য মূল্য প্রদান করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সেরা বাজেট কমিউটার পছন্দ</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Hero Splendor Plus DTS-i already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Hero Splendor Plus DTS-i\n")

	return nil
}

// GetName returns the seeder name
func (s *HeroSplendorPlusDTSiReview) GetName() string {
	return "HeroSplendorPlusDTSiReview"
}
