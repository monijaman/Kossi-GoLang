package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsar150SingleDiscABSReview handles seeding Bajaj Pulsar 150 Single Disc ABS product review and translation
type BajajPulsar150SingleDiscABSReview struct {
	BaseSeeder
}

// NewBajajPulsar150SingleDiscABSReviewSeeder creates a new BajajPulsar150SingleDiscABSReview
func NewBajajPulsar150SingleDiscABSReviewSeeder() *BajajPulsar150SingleDiscABSReview {
	return &BajajPulsar150SingleDiscABSReview{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar 150 Single Disc ABS Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsar150SingleDiscABSReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar 150 Single Disc ABS").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Pulsar 150 Single Disc ABS product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Pulsar 150 Single Disc ABS product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar 150 Single Disc ABS already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar 150 Single Disc ABS Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Pulsar 150 Single Disc ABS is a practical commuter motorcycle featuring front ABS braking system for enhanced safety, 150cc DTS-i fuel-injected engine, comfortable ergonomics, and excellent fuel efficiency. Priced at ৳1,52,000, it offers the perfect balance of safety, performance, and affordability making it ideal for daily commuters seeking ABS protection without breaking the bank.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar 150 Single Disc ABS Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">Front ABS:</strong> <span class="highlight-value">Anti-lock braking system for safer stops</span></li>
<li class="highlight-item"><strong class="highlight-label">150cc Engine:</strong> <span class="highlight-value">DTS-i with digital fuel injection</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Ride:</strong> <span class="highlight-value">Well-designed seat and ergonomics</span></li>
<li class="highlight-item"><strong class="highlight-label">Great Mileage:</strong> <span class="highlight-value">40-45 km/l fuel efficiency</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar 150 Single Disc ABS Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ABS Safety:</strong> <span class="pro-description">Front ABS prevents wheel lock during emergency braking</span></li>
<li class="pro-item"><strong class="pro-label">Affordable ABS:</strong> <span class="pro-description">Most affordable ABS commuter bike at ৳1,52,000</span></li>
<li class="pro-item"><strong class="pro-label">Great Mileage:</strong> <span class="pro-description">40-45 km/l ideal for daily commuting</span></li>
<li class="pro-item"><strong class="pro-label">Reliable Engine:</strong> <span class="pro-description">Proven DTS-i technology with good performance</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seat:</strong> <span class="pro-description">Well-padded for long daily commutes</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar 150 Single Disc ABS Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Single Disc Only:</strong> <span class="con-description">Rear drum brake not as effective</span></li>
<li class="con-item"><strong class="con-label">No Rear ABS:</strong> <span class="con-description">ABS only on front wheel</span></li>
<li class="con-item"><strong class="con-label">Basic Design:</strong> <span class="con-description">Less aggressive styling compared to sportier variants</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar 150 Single Disc ABS Price & Running Cost</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳1,52,000 - Best value ABS commuter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very Low - ৳4-5 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳4-5 per km (40-45 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar 150 Single Disc ABS Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.1</span> <span class="rating-note">- Good commuter performance</span></li>
<li class="rating-item"><strong class="rating-label">Mileage:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Excellent fuel efficiency</span></li>
<li class="rating-item"><strong class="rating-label">Safety:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Good front ABS safety</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.2</span> <span class="rating-note">- Comfortable commuter seat</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Excellent value for money</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar 150 Single Disc ABS?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.3/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At just ৳1,52,000, the Bajaj Pulsar 150 Single Disc ABS provides excellent value for daily commuting with the safety benefit of front ABS braking. Perfect for budget-conscious riders who don't want to compromise on safety features.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For affordable ABS safety and excellent mileage</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.3,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Bajaj Pulsar 150 Single Disc ABS review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar 150 Single Disc ABS (Rating: 4.3/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ পালসার ১৫০ সিঙ্গেল ডিস্ক এবিএস রিভিউ বাংলাদেশ ২০২৪ - সম্পূর্ণ গাইড</h1>
<p class="summary-text">বাজাজ পালসার ১৫০ সিঙ্গেল ডিস্ক এবিএস একটি ব্যবহারিক কমিউটার মোটরসাইকেল যা সামনের এবিএস ব্রেকিং সিস্টেম, ১৫০সিসি ডিটিএস-আই জ্বালানি-ইনজেক্টেড ইঞ্জিন এবং চমৎকার জ্বালানি দক্ষতা সহ আসে। ৳১,৫২,০০০ টাকায় মূল্যায়িত, এটি নিরাপত্তা, পারফরমেন্স এবং সামর্থ্যের নিখুঁত ভারসাম্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার ১৫০ সিঙ্গেল ডিস্ক এবিএস মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">সামনের এবিএস:</strong> <span class="highlight-value">নিরাপদ ব্রেকিংয়ের জন্য অ্যান্টি-লক সিস্টেম</span></li>
<li class="highlight-item"><strong class="highlight-label">১৫০সিসি ইঞ্জিন:</strong> <span class="highlight-value">ডিজিটাল জ্বালানি ইনজেকশন সহ ডিটিএস-আই</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক রাইড:</strong> <span class="highlight-value">ভাল-ডিজাইন করা সিট এবং এরগোনমিক্স</span></li>
<li class="highlight-item"><strong class="highlight-label">দুর্দান্ত মাইলেজ:</strong> <span class="highlight-value">৪০-৪৫ কিমি/লিটার জ্বালানি দক্ষতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার ১৫০ সিঙ্গেল ডিস্ক এবিএস সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">এবিএস নিরাপত্তা:</strong> <span class="pro-description">জরুরি ব্রেকিংয়ে চাকার লক প্রতিরোধ করে</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী এবিএস:</strong> <span class="pro-description">সবচেয়ে সাশ্রয়ী এবিএস কমিউটার বাইক</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত মাইলেজ:</strong> <span class="pro-description">দৈনন্দিন যাতায়াতের জন্য আদর্শ</span></li>
<li class="pro-item"><strong class="pro-label">নির্ভরযোগ্য ইঞ্জিন:</strong> <span class="pro-description">প্রমাণিত ডিটিএস-আই প্রযুক্তি ভাল পারফরমেন্স সহ</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার ১৫০ সিঙ্গেল ডিস্ক এবিএস অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">শুধুমাত্র সিঙ্গেল ডিস্ক:</strong> <span class="con-description">পিছনের ড্রাম ব্রেক তত কার্যকর নয়</span></li>
<li class="con-item"><strong class="con-label">পিছনে এবিএস নেই:</strong> <span class="con-description">শুধুমাত্র সামনের চাকায় এবিএস</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত রায়: বাজাজ পালসার ১৫০ সিঙ্গেল ডিস্ক এবিএস কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.3/৫.০</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">মাত্র ৳১,৫২,০০০ টাকায় বাজাজ পালসার ১৫০ সিঙ্গেল ডিস্ক এবিএস দৈনন্দিন যাতায়াতের জন্য চমৎকার মূল্য প্রদান করে। বাজেট-সচেতন রাইডারদের জন্য নিরাপত্তা সুবিধা চাওয়াদের জন্য আদর্শ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশকৃত:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- সাশ্রয়ী এবিএস নিরাপত্তা এবং উত্তম মাইলেজের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar 150 Single Disc ABS already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar 150 Single Disc ABS\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPulsar150SingleDiscABSReview) GetName() string {
	return "BajajPulsar150SingleDiscABSReview"
}
