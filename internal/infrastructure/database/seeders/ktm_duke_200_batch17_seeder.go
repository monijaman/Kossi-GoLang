package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// KTMDuke200ReviewBatch17 handles seeding KTM Duke 200 product review and translation
type KTMDuke200ReviewBatch17 struct {
	BaseSeeder
}

// NewKTMDuke200ReviewBatch17Seeder creates a new KTMDuke200ReviewBatch17
func NewKTMDuke200ReviewBatch17Seeder() *KTMDuke200ReviewBatch17 {
	return &KTMDuke200ReviewBatch17{BaseSeeder: BaseSeeder{name: "KTM Duke 200 Batch17 Review"}}
}

// Seed implements the Seeder interface
func (s *KTMDuke200ReviewBatch17) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "KTM Duke 200").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("ktm duke 200 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding ktm duke 200 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for KTM Duke 200 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">KTM Duke 200 Review Bangladesh 2024 - Entry-Level Performance Beast</h1>
<p class="summary-text">The KTM Duke 200 is a 200cc accessible performance street bike combining aggressive styling with fun dynamics and excellent value. Priced around ৳2,85,000, it delivers sharp handling, engaging engine, modern features, and the perfect introduction to performance biking for young enthusiasts seeking thrills.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">KTM Duke 200 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc Beast:</strong> <span class="highlight-value">Performance entry</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Design:</strong> <span class="highlight-value">Street attitude</span></li>
<li class="highlight-item"><strong class="highlight-label">Sharp Handling:</strong> <span class="highlight-value">Engaging dynamics</span></li>
<li class="highlight-item"><strong class="highlight-label">Fun Factor:</strong> <span class="highlight-value">Thrilling ride</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">KTM Duke 200 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Performance look</span></li>
<li class="pro-item"><strong class="pro-label">Responsive Engine:</strong> <span class="pro-description">Direct thrills</span></li>
<li class="pro-item"><strong class="pro-label">Sharp Handling:</strong> <span class="pro-description">Nimble feel</span></li>
<li class="pro-item"><strong class="pro-label">Great Price:</strong> <span class="pro-description">Value packed</span></li>
<li class="pro-item"><strong class="pro-label">Modern Features:</strong> <span class="pro-description">ABS equipped</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">KTM Duke 200 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Build Quality:</strong> <span class="con-description">Entry-level finish</span></li>
<li class="con-item"><strong class="con-label">Fuel Tank:</strong> <span class="con-description">Regular refills</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Performance bikes costly</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">KTM Duke 200 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,85,000 - Entry performance</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳8-10 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">30-35 km/l decent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">KTM Duke 200 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Responsive thrills</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Sharp agile</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Aggressive style</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Modern equipped</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Good budget</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy KTM Duke 200?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,85,000, the Duke 200 is perfect for young performance seekers wanting authentic 200cc street thrills with aggressive styling, sharp handling, and excellent value without premium pricing.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For young performance enthusiasts</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.6,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating ktm duke 200 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for KTM Duke 200 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">KTM ডিউক 200 রিভিউ বাংলাদেশ 2024 - এন্ট্রি-লেভেল পারফরম্যান্স বিস্ট</h1>
<p class="summary-text">KTM ডিউক 200 একটি 200cc অ্যাক্সেসযোগ্য পারফরম্যান্স স্ট্রিট বাইক যা আক্রমণাত্মক স্টাইলিং এবং মজাদার গতিশীলতাকে একত্রিত করে। ৳2,85,000 টাকায় মূল্যায়িত, এটি ধারালো হ্যান্ডলিং, নিযুক্ত ইঞ্জিন এবং পারফরম্যান্স বাইকিংয়ের নিখুঁত পরিচয় প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">KTM ডিউক 200 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc বিস্ট:</strong> <span class="highlight-value">পারফরম্যান্স এন্ট্রি</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক ডিজাইন:</strong> <span class="highlight-value">রাস্তা মনোভাব</span></li>
<li class="highlight-item"><strong class="highlight-label">ধারালো হ্যান্ডলিং:</strong> <span class="highlight-value">নিযুক্ত গতিশীলতা</span></li>
<li class="highlight-item"><strong class="highlight-label">মজা ফ্যাক্টর:</strong> <span class="highlight-value">রোমাঞ্চকর যাত্রা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">KTM ডিউক 200 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">পারফরম্যান্স লুক</span></li>
<li class="pro-item"><strong class="pro-label">প্রতিক্রিয়াশীল ইঞ্জিন:</strong> <span class="pro-description">সরাসরি thrills</span></li>
<li class="pro-item"><strong class="pro-label">ধারালো হ্যান্ডলিং:</strong> <span class="pro-description">নমনীয় অনুভূতি</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত দাম:</strong> <span class="pro-description">মূল্য প্যাকড</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক বৈশিষ্ট্য:</strong> <span class="pro-description">ABS সজ্জিত</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">KTM ডিউক 200 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">বিল্ড গুণমান:</strong> <span class="con-description">এন্ট্রি-লেভেল ফিনিশ</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি ট্যাঙ্ক:</strong> <span class="con-description">নিয়মিত রিফিল</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">পারফরম্যান্স বাইক ব্যয়বহুল</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: KTM ডিউক 200 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,85,000 টাকায়, ডিউক 200 তরুণ পারফরম্যান্স চাহনাকারীদের জন্য নিখুঁত যারা খাঁটি 200cc রাস্তার thrills, আক্রমণাত্মক স্টাইলিং এবং চমৎকার মূল্য চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- তরুণ পারফরম্যান্স উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for KTM Duke 200 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for KTM Duke 200\n")

	return nil
}

// GetName returns the seeder name
func (s *KTMDuke200ReviewBatch17) GetName() string {
	return "KTMDuke200ReviewBatch17"
}
