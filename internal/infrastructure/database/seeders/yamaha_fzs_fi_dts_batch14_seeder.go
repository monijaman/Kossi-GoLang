package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// YamahaFZSFiDTSReview handles seeding Yamaha FZS Fi DTS product review and translation
type YamahaFZSFiDTSReview struct {
	BaseSeeder
}

// NewYamahaFZSFiDTSReviewSeeder creates a new YamahaFZSFiDTSReview
func NewYamahaFZSFiDTSReviewSeeder() *YamahaFZSFiDTSReview {
	return &YamahaFZSFiDTSReview{BaseSeeder: BaseSeeder{name: "Yamaha FZS Fi DTS Review"}}
}

// Seed implements the Seeder interface
func (s *YamahaFZSFiDTSReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha FZS Fi DTS").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha fzs fi dts product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding yamaha fzs fi dts product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Yamaha FZS Fi DTS already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Yamaha FZS Fi DTS Review Bangladesh 2024 - Sporty Street Commuter</h1>
<p class="summary-text">The Yamaha FZS Fi DTS is a dynamic 150cc street bike with fuel injection, twin-spark technology, aggressive styling, and excellent performance. Priced around ৳2,65,000, it delivers sporty street character with reliable Yamaha technology.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha FZS Fi DTS Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc Fuel Injection:</strong> <span class="highlight-value">Modern efficient technology</span></li>
<li class="highlight-item"><strong class="highlight-label">Twin-Spark Performance:</strong> <span class="highlight-value">Enhanced power delivery</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Styling:</strong> <span class="highlight-value">Street fighter design</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">45-50 km/l efficiency</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha FZS Fi DTS Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Fuel Injection:</strong> <span class="pro-description">Modern FI technology</span></li>
<li class="pro-item"><strong class="pro-label">Twin-Spark Engine:</strong> <span class="pro-description">Enhanced performance</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Styling:</strong> <span class="pro-description">Aggressive street fighter</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Brand:</strong> <span class="pro-description">Yamaha trusted quality</span></li>
<li class="pro-item"><strong class="pro-label">Good Efficiency:</strong> <span class="pro-description">45-50 km/l mileage</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha FZS Fi DTS Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Premium maintenance</span></li>
<li class="con-item"><strong class="con-label">Front Suspension:</strong> <span class="con-description">Basic telescopic forks</span></li>
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Short duration comfort</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha FZS Fi DTS Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,65,000 - Sporty street bike</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">45-50 km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha FZS Fi DTS Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Performance:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- 150cc FI twin-spark</span></li>
<li class="rating-item"><strong class="rating-label">Styling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Aggressive street fighter</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good street bike</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Yamaha quality</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good features worth</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha FZS Fi DTS?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,65,000, the Yamaha FZS Fi DTS is perfect for street riders seeking a sporty commuter with modern fuel injection, twin-spark performance, aggressive styling, and reliable Yamaha engineering.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sporty street commuters</span></p>
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
		return fmt.Errorf("error creating yamaha fzs fi dts review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Yamaha FZS Fi DTS (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ইয়ামাহা FZS Fi DTS রিভিউ বাংলাদেশ 2024 - স্পোর্টি স্ট্রিট কমিউটার</h1>
<p class="summary-text">ইয়ামাহা FZS Fi DTS একটি গতিশীল 150cc স্ট্রিট বাইক যা জ্বালানি ইঞ্জেকশন, টুইন-স্পার্ক প্রযুক্তি, আক্রমণাত্মক স্টাইলিং এবং চমৎকার কর্মক্ষমতা সহ আসে। ৳2,65,000 টাকায় মূল্যায়িত, এটি নির্ভরযোগ্য ইয়ামাহা প্রযুক্তি সহ স্পোর্টি স্ট্রিট চরিত্র প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা FZS Fi DTS মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc জ্বালানি ইঞ্জেকশন:</strong> <span class="highlight-value">আধুনিক দক্ষ প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">টুইন-স্পার্ক কর্মক্ষমতা:</strong> <span class="highlight-value">উন্নত শক্তি সরবরাহ</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="highlight-value">স্ট্রিট ফাইটার ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">45-50 km/l দক্ষতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা FZS Fi DTS সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">জ্বালানি ইঞ্জেকশন:</strong> <span class="pro-description">আধুনিক FI প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">টুইন-স্পার্ক ইঞ্জিন:</strong> <span class="pro-description">উন্নত কর্মক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্টি স্টাইলিং:</strong> <span class="pro-description">আক্রমণাত্মক স্ট্রিট ফাইটার</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ব্র্যান্ড:</strong> <span class="pro-description">ইয়ামাহা বিশ্বাসযোগ্য গুণমান</span></li>
<li class="pro-item"><strong class="pro-label">ভাল দক্ষতা:</strong> <span class="pro-description">45-50 km/l মাইলেজ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা FZS Fi DTS অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">প্রিমিয়াম রক্ষণাবেক্ষণ</span></li>
<li class="con-item"><strong class="con-label">সামনের সাসপেনশন:</strong> <span class="con-description">মৌলিক টেলিস্কোপিক ফর্ক</span></li>
<li class="con-item"><strong class="con-label">আসনের আরাম:</strong> <span class="con-description">সংক্ষিপ্ত সময়ের জন্য আরাম</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা FZS Fi DTS কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,65,000 টাকায়, ইয়ামাহা FZS Fi DTS স্ট্রিট রাইডারদের জন্য নিখুঁত যারা আধুনিক জ্বালানি ইঞ্জেকশন, টুইন-স্পার্ক কর্মক্ষমতা, আক্রমণাত্মক স্টাইলিং এবং নির্ভরযোগ্য ইয়ামাহা প্রযুক্তি সহ স্পোর্টি কমিউটার খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্পোর্টি স্ট্রিট যাত্রীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Yamaha FZS Fi DTS already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Yamaha FZS Fi DTS\n")

	return nil
}

// GetName returns the seeder name
func (s *YamahaFZSFiDTSReview) GetName() string {
	return "YamahaFZSFiDTSReview"
}
