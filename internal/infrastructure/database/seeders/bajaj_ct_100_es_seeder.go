package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajCT100ESReview handles seeding Bajaj CT 100 ES product review and translation
type BajajCT100ESReview struct {
	BaseSeeder
}

// NewBajajCT100ESReviewSeeder creates a new BajajCT100ESReview
func NewBajajCT100ESReviewSeeder() *BajajCT100ESReview {
	return &BajajCT100ESReview{BaseSeeder: BaseSeeder{name: "Bajaj CT 100 ES Review"}}
}

// Seed implements the Seeder interface
func (s *BajajCT100ESReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj CT 100 ES").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj CT 100 ES product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj CT 100 ES product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj CT 100 ES already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj CT 100 ES Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj CT 100 ES is an ultra-reliable, highly practical work motorcycle featuring 100cc engine, fuel efficiency up to 65 km/l, robust build quality, and proven durability. Priced at ৳75,000, it represents exceptional value for rural and semi-urban riders who prioritize reliability and low running costs over modern features. Perfect for daily work commuting and occasional long-distance touring on budget.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj CT 100 ES Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Fuel Efficiency:</strong> <span class="highlight-value">60-65 km/l exceptional mileage</span></li>
<li class="highlight-item"><strong class="highlight-label">100cc Engine:</strong> <span class="highlight-value">Simple, proven, reliable technology</span></li>
<li class="highlight-item"><strong class="highlight-label">Robust Build:</strong> <span class="highlight-value">Built for durability and heavy use</span></li>
<li class="highlight-item"><strong class="highlight-label">Work-Ready:</strong> <span class="highlight-value">Practical design for daily work</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj CT 100 ES Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Exceptional Mileage:</strong> <span class="pro-description">60-65 km/l outstanding fuel economy</span></li>
<li class="pro-item"><strong class="pro-label">Ultra Affordable:</strong> <span class="pro-description">Starting at just ৳75,000</span></li>
<li class="pro-item"><strong class="pro-label">Proven Reliability:</strong> <span class="pro-description">Simple design ensures longevity</span></li>
<li class="pro-item"><strong class="pro-label">Low Maintenance:</strong> <span class="pro-description">Affordable spare parts and service</span></li>
<li class="pro-item"><strong class="pro-label">Practical Design:</strong> <span class="pro-description">Perfect for work and daily commuting</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj CT 100 ES Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Outdated Design:</strong> <span class="con-description">Very basic styling, not modern</span></li>
<li class="con-item"><strong class="con-label">Minimal Features:</strong> <span class="con-description">No fancy amenities or technology</span></li>
<li class="con-item"><strong class="con-label">Low Power:</strong> <span class="con-description">100cc engine has limited performance</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj CT 100 ES Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳75,000 - Absolute budget option</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Minimal - ৳2.5 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳2.5 per km (60-65 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj CT 100 ES Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">3.2</span> <span class="rating-note">- Basic but adequate for work</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Exceptional fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Highly proven and durable</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Unbeatable budget value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj CT 100 ES?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.1/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At just ৳75,000, the Bajaj CT 100 ES is an absolute value champion for budget-conscious buyers in rural and semi-urban areas. With exceptional fuel efficiency and proven reliability, it's ideal for work and daily commuting despite its basic design and limited features.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget commuters and work usage</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.1,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Bajaj CT 100 ES review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj CT 100 ES (Rating: 4.1/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ সিটি ১०० ইএস রিভিউ বাংলাদেশ २०२४ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ সিটি १००ईএস একটি অতি নির্ভরযোগ্য, অত্যন্ত ব্যবহারিক কর্ম মোটরসাইকেল যা १००সিসি ইঞ্জিন, ६०-६५ কিমি/লিটার জ্বালানি দক্ষতা এবং প্রমাণিত স্থায়িত্ব সহ আসে। ৳७५,००० টাকায় মূল্যায়িত, এটি গ্রামীণ এবং আধা-শহুরে রাইডারদের জন্য অসাধারণ মূল্য প্রতিনিধিত্ব করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ সিটি १०० ইএস মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষতা:</strong> <span class="highlight-value">६०-६५ किमी/लिटer অসাধারণ মাইলেজ</span></li>
<li class="highlight-item"><strong class="highlight-label">१०० সিসি ইঞ্জিন:</strong> <span class="highlight-value">সহজ, প্রমাণিত, নির্ভরযোগ্য প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">শক্তিশালী বিল্ড:</strong> <span class="highlight-value">স্থায়িত্ব এবং ভারী ব্যবহারের জন্য নির্মিত</span></li>
<li class="highlight-item"><strong class="highlight-label">কর্ম-প্রস্তুত:</strong> <span class="highlight-value">দৈনন্দিন কর্মের জন্য অনুশীলনী ডিজাইন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ সিটি १०० ইএস সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">অসাধারণ মাইলেজ:</strong> <span class="pro-description">६०-६५ किमी/लिটr দুর্দান্ত জ্বালানি অর্থনীতি</span></li>
<li class="pro-item"><strong class="pro-label">অত্যন্ত সাশ্রয়ী:</strong> <span class="pro-description">মাত্র ৳७५,००० থেকে শুরু</span></li>
<li class="pro-item"><strong class="pro-label">প্রমাণিত নির্ভরযোগ্যতা:</strong> <span class="pro-description">সহজ ডিজাইন দীর্ঘায়ু নিশ্চিত করে</span></li>
<li class="pro-item"><strong class="pro-label">কম রক্ষণাবেক্ষণ:</strong> <span class="pro-description">সাশ্রয়ী যন্ত্রাংশ এবং সেবা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ সিটি १००ईएस অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">পুরানো ডিজাইন:</strong> <span class="con-description">অত্যন্ত মৌলিক স্টাইলিং, আধুনিক নয়</span></li>
<li class="con-item"><strong class="con-label">ন্যূনতম বৈশিষ্ট্য:</strong> <span class="con-description">কোন ফ্যান্সি সুবিধা বা প্রযুক্তি নেই</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাজাজ সিটি १०० ইএস किনबेन?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ्रिक রেটিং:</strong> <span class="score-value">४.१/५.०</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">মাত्र ৳७५,००० টাকায় বাজাজ सिटी १००ईएस বাজেট-সচেতন ক্রেতাদের জন্য একটি সম্পূর্ণ মূল্য চ্যাম্পিয়ন। অসাধারণ জ্বালানি দক্ষতা এবং প্রমাণিত নির্ভরযোগ্যতার সাথে, এটি কাজ এবং দৈনন্দিন যাতায়াতের জন্য আদর্শ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট যাতায়াত এবং কর্ম ব্যবহারের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj CT 100 ES already exists\n")
		return nil
	}

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating Bangla translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Bajaj CT 100 ES\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajCT100ESReview) GetName() string {
	return "BajajCT100ESReview"
}
