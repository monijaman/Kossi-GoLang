package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVSApacheRTR200Review handles seeding TVS Apache RTR 200 product review and translation
type TVSApacheRTR200Review struct {
	BaseSeeder
}

// NewTVSApacheRTR200ReviewSeeder creates a new TVSApacheRTR200Review
func NewTVSApacheRTR200ReviewSeeder() *TVSApacheRTR200Review {
	return &TVSApacheRTR200Review{BaseSeeder: BaseSeeder{name: "TVS Apache RTR 200 Review"}}
}

// Seed implements the Seeder interface
func (s *TVSApacheRTR200Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Apache RTR 200").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tvs apache rtr 200 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding tvs apache rtr 200 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for TVS Apache RTR 200 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">TVS Apache RTR 200 Review Bangladesh 2024 - Performance Racer</h1>
<p class="summary-text">The TVS Apache RTR 200 is a performance-oriented sport bike featuring 200cc engine, aggressive design, excellent handling, and ABS safety. Priced around ৳2,20,000, it delivers genuine racing performance with Indian reliability for enthusiast riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache RTR 200 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc Engine:</strong> <span class="highlight-value">Performance race engine</span></li>
<li class="highlight-item"><strong class="highlight-label">ABS Safety:</strong> <span class="highlight-value">Front and rear ABS protection</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Design:</strong> <span class="highlight-value">Sporty racing styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Excellent Handling:</strong> <span class="highlight-value">Responsive chassis tuning</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache RTR 200 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Strong Performance:</strong> <span class="pro-description">Powerful 200cc with impressive acceleration</span></li>
<li class="pro-item"><strong class="pro-label">Dual ABS:</strong> <span class="pro-description">Front and rear ABS for safety</span></li>
<li class="pro-item"><strong class="pro-label">Excellent Handling:</strong> <span class="pro-description">Responsive and agile performance</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Design:</strong> <span class="pro-description">Racing-inspired styling</span></li>
<li class="pro-item"><strong class="pro-label">Good Value:</strong> <span class="pro-description">Excellent sport bike features</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache RTR 200 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Sport Focused:</strong> <span class="con-description">Not ideal for casual commuting</span></li>
<li class="con-item"><strong class="con-label">Fuel Consumption:</strong> <span class="con-description">35-40 km/l moderate efficiency</span></li>
<li class="con-item"><strong class="con-label">Maintenance Cost:</strong> <span class="con-description">Sport bike service expensive</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Apache RTR 200 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,20,000 - Premium sport performance</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳5-6 per km (35-40 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Apache RTR 200 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Strong 200cc power</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Excellent chassis control</span></li>
<li class="rating-item"><strong class="rating-label">Safety:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Dual ABS equipped</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Aggressive racing styling</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good sport value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Apache RTR 200?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,20,000, the TVS Apache RTR 200 is excellent for enthusiast riders seeking genuine performance, dual ABS safety, and racing-inspired handling with strong Indian engineering and affordability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance enthusiasts</span></p>
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
		return fmt.Errorf("error creating tvs apache rtr 200 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for TVS Apache RTR 200 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">টিভিএস Apache RTR 200 রিভিউ বাংলাদেশ 2024 - পারফরমেন্স রেসার</h1>
<p class="summary-text">টিভিএস Apache RTR 200 একটি পারফরমেন্স-উন্মুখ স্পোর্ট বাইক যা 200cc ইঞ্জিন, আক্রমণাত্মক ডিজাইন, চমৎকার হ্যান্ডলিং এবং ABS নিরাপত্তা বৈশিষ্ট্যযুক্ত। ৳2,20,000 টাকায় মূল্যায়িত, এটি উত্সাহী রাইডারদের জন্য খাঁটি রেসিং পারফরমেন্স সহ ভারতীয় নির্ভরযোগ্যতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">টিভিএস Apache RTR 200 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">200cc ইঞ্জিন:</strong> <span class="highlight-value">পারফরমেন্স রেস ইঞ्जिन</span></li>
<li class="highlight-item"><strong class="highlight-label">ABS নিরাপত্তা:</strong> <span class="highlight-value">সামনে এবং পিছনে ABS সুরক্ষা</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক ডিজাইন:</strong> <span class="highlight-value">খেলাধুলাপূর্ণ রেসিং স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">চমৎকার হ্যান্ডলিং:</strong> <span class="highlight-value">প্রতিক্রিয়াশীল চ্যাসিস টিউনিং</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">টিভিএস Apache RTR 200 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী পারফরমেন্স:</strong> <span class="pro-description">চিত্তাকর্ষক ত্বরণ সহ শক্তিশালী 200cc</span></li>
<li class="pro-item"><strong class="pro-label">Dual ABS:</strong> <span class="pro-description">নিরাপত্তার জন্য সামনে এবং পিছনে ABS</span></li>
<li class="pro-item"><strong class="pro-label">চমৎকার হ্যান্ডলিং:</strong> <span class="pro-description">প্রতিক্রিয়াশীল এবং চটপটে পারফরমেন্স</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক ডিজাইন:</strong> <span class="pro-description">রেসিং-অনুপ্রাণিত স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মূল্য:</strong> <span class="pro-description">চমৎকার স্পোর্ট বাইক বৈশিষ্ট্য</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">টিভিএস Apache RTR 200 অসুবिधा</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">খেলা ফোকাসড:</strong> <span class="con-description">নৈমিত্তিক যাতায়াতের জন্য আদর্শ নয়</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি খরচ:</strong> <span class="con-description">35-40 km/l মধ্যম দক্ষতা</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ খরচ:</strong> <span class="con-description">স্পোর্ট বাইক সেবা ব্যয়বহুল</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: টিভিএস Apache RTR 200 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,20,000 টাকায়, টিভিএস Apache RTR 200 খাঁটি পারফরমেন্স, ডুয়াল ABS নিরাপত্তা এবং রেসিং-অনুপ্রাণিত হ্যান্ডলিং সহ শক্তিশালী ভারতীয় প্রকৌশল এবং সামর্থ্য সহ উত্সাহী রাইডারদের জন্য চমৎকার।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারिश:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পারফরমেন্স উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for TVS Apache RTR 200 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for TVS Apache RTR 200\n")

	return nil
}

// GetName returns the seeder name
func (s *TVSApacheRTR200Review) GetName() string {
	return "TVSApacheRTR200Review"
}
