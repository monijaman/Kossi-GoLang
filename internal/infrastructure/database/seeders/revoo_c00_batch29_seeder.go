package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type RevooC00Batch29 struct {
	BaseSeeder
}

func NewRevooC00Batch29Seeder() *RevooC00Batch29 {
	return &RevooC00Batch29{BaseSeeder: BaseSeeder{name: "Revoo C00 Batch29 Review"}}
}

func (s *RevooC00Batch29) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Revoo C00").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("revoo c00 product not found")
		}
		return fmt.Errorf("error finding product: %w", err)
	}

	var existingReview models.ProductReviewModel
	if err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error; err == nil {
		fmt.Printf("   ℹ️  Review already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Revoo C00 Review Bangladesh 2024 - Versatile Street Performer</h1>
<p class="summary-text">The Revoo C00 delivers versatile street performance at ৳1,39,900. With its 150cc engine, dual-channel ABS, modern design, and 4.7/5.0 rating, it offers excellent balance between performance and practicality for urban riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Revoo C00 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>Engine:</strong> 150cc engine</li>
<li class="highlight-item"><strong>Safety:</strong> Dual-channel ABS</li>
<li class="highlight-item"><strong>Design:</strong> Modern street styling</li>
<li class="highlight-item"><strong>Value:</strong> Excellent price-performance</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Revoo C00 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Versatile Performance:</strong> Street-ready capability</li>
<li class="pro-item"><strong>Safety Features:</strong> Dual-channel ABS</li>
<li class="pro-item"><strong>Fuel Efficiency:</strong> Good economy</li>
<li class="pro-item"><strong>Comfortable Ride:</strong> Urban-friendly ergonomics</li>
<li class="pro-item"><strong>Value Pricing:</strong> Affordable quality</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Revoo C00 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Limited Service:</strong> Fewer centers</li>
<li class="con-item"><strong>Resale Value:</strong> Lower than established brands</li>
<li class="con-item"><strong>Parts Availability:</strong> Can be limited</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Revoo C00 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,39,900 - Versatile street</span></p>
<p class="value-item"><strong>Engine:</strong> <span>150cc engine</span></p>
<p class="value-item"><strong>Safety:</strong> <span>Dual-channel ABS</span></p>
<p class="value-item"><strong>Efficiency:</strong> <span>Good fuel economy</span></p>
<p class="value-item"><strong>Comfort:</strong> <span>Urban ergonomics</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Revoo C00 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Modern street</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - Versatile power</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Urban-friendly</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - Great pricing</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - Solid build</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Revoo C00?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,39,900, the Revoo C00 is an excellent choice for urban riders seeking versatile street performance with modern safety features and good fuel economy. Recommended for daily commuting and city riding.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For urban riders</span></p>
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
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">Revoo C00 রিভিউ বাংলাদেশ 2024 - বহুমুখী স্ট্রিট পারফরমার</h1>
<p class="summary-text">Revoo C00 ৳২,৯৫,০০০ টাকায় বহুমুখী স্ট্রিট পারফরম্যান্স প্রদান করে। এর ১৫০সিসি ইঞ্জিন, ডুয়াল-চ্যানেল ABS, আধুনিক ডিজাইন এবং ৪.৭/৫.০ রেটিং সহ, এটি শহুরে রাইডারদের জন্য পারফরম্যান্স এবং ব্যবহারিকতার মধ্যে চমৎকার ভারসাম্য অফার করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Revoo C00 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>ইঞ্জিন:</strong> ১৫০সিসি ইঞ্জিন</li>
<li class="highlight-item"><strong>নিরাপত্তা:</strong> ডুয়াল-চ্যানেল ABS</li>
<li class="highlight-item"><strong>ডিজাইন:</strong> আধুনিক স্ট্রিট স্টাইলিং</li>
<li class="highlight-item"><strong>মূল্য:</strong> চমৎকার মূল্য-পারফরম্যান্স</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Revoo C00 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>বহুমুখী পারফরম্যান্স:</strong> স্ট্রিট-রেডি ক্ষমতা</li>
<li class="pro-item"><strong>নিরাপত্তা ফিচার:</strong> ডুয়াল-চ্যানেল ABS</li>
<li class="pro-item"><strong>জ্বালানি দক্ষতা:</strong> ভাল অর্থনীতি</li>
<li class="pro-item"><strong>আরামদায়ক রাইড:</strong> শহুরে-বন্ধুত্বপূর্ণ এরগোনমিক্স</li>
<li class="pro-item"><strong>মূল্য মূল্যায়ন:</strong> সাশ্রয়ী মান</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Revoo C00 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>সীমিত সেবা:</strong> কম কেন্দ্র</li>
<li class="con-item"><strong>পুনঃবিক্রয় মূল্য:</strong> প্রতিষ্ঠিত ব্র্যান্ডের চেয়ে কম</li>
<li class="con-item"><strong>যন্ত্রাংশ উপলব্ধতা:</strong> সীমিত হতে পারে</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: Revoo C00 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳২,৯৫,০০০ টাকায়, Revoo C00 আধুনিক নিরাপত্তা ফিচার এবং ভাল জ্বালানি অর্থনীতি সহ বহুমুখী স্ট্রিট পারফরম্যান্স খুঁজছেন শহুরে রাইডারদের জন্য একটি চমৎকার পছন্দ। দৈনিক যাতায়াত এবং শহর রাইডিংয়ের জন্য সুপারিশকৃত।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- শহুরে রাইডারদের জন্য</span></p>
</div>
</section>
</article>`

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating translation: %w", err)
	}

	fmt.Printf("   ✅ Created Revoo C00\n")
	return nil
}

func (s *RevooC00Batch29) GetName() string {
	return "RevooC00Batch29"
}
