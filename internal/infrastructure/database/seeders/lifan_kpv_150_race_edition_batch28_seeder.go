package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type LifanKPV150RaceEditionBatch28 struct {
	BaseSeeder
}

func NewLifanKPV150RaceEditionBatch28Seeder() *LifanKPV150RaceEditionBatch28 {
	return &LifanKPV150RaceEditionBatch28{BaseSeeder: BaseSeeder{name: "Lifan KPV 150 Race Edition Batch28 Review"}}
}

func (s *LifanKPV150RaceEditionBatch28) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Lifan KPV 150 Race Edition").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("lifan kpv 150 race edition product not found")
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
<h1 class="review-main-title">Lifan KPV 150 Race Edition Review Bangladesh 2024 - Track-Ready Performance Beast</h1>
<p class="summary-text">The Lifan KPV 150 Race Edition stands as the pinnacle of Chinese racing technology, delivering professional-grade performance at ৳3,80,000. This track-focused machine combines aggressive styling with genuine racing credentials, 150cc liquid-cooled engine, and comprehensive racing features package.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Lifan KPV 150 Race Edition Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Racing Engine:</strong> Liquid-cooled performance</li>
<li class="highlight-item"><strong>Racing Suspension:</strong> Adjustable setup</li>
<li class="highlight-item"><strong>Digital Dashboard:</strong> Shift light equipped</li>
<li class="highlight-item"><strong>Track-Ready:</strong> Professional-grade</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Lifan KPV 150 Race Edition Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Genuine Racing Pedigree:</strong> Track-ready setup</li>
<li class="pro-item"><strong>Impressive Power:</strong> Strong acceleration</li>
<li class="pro-item"><strong>Professional Suspension:</strong> Adjustable quality</li>
<li class="pro-item"><strong>Competitive Pricing:</strong> Performance value</li>
<li class="pro-item"><strong>Racing Features:</strong> Comprehensive package</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Lifan KPV 150 Race Edition Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Limited Service Network:</strong> Availability concerns</li>
<li class="con-item"><strong>Higher Maintenance:</strong> Racing-spec costs</li>
<li class="con-item"><strong>Aggressive Position:</strong> Comfort compromise</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Lifan KPV 150 Race Edition Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,80,000 - Racing performance</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc liquid-cooled</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>Racing-tuned</span></p>
<p class="value-item"><strong>Suspension:</strong> <span>Adjustable front/rear</span></p>
<p class="value-item"><strong>Brakes:</strong> <span>Radial mount calipers</span></p>
<p class="value-item"><strong>Features:</strong> <span>Full racing package</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Lifan KPV 150 Race Edition Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.9</span> - Aggressive racing</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.8</span> - Track-ready power</li>
<li class="rating-item"><strong>Features:</strong> <span>4.8</span> - Comprehensive racing</li>
<li class="rating-item"><strong>Build Quality:</strong> <span>4.7</span> - Racing-spec</li>
<li class="rating-item"><strong>Value:</strong> <span>4.8</span> - Performance value</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Lifan KPV 150 Race Edition?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,80,000, the Lifan KPV 150 Race Edition is a genuine performance machine that delivers track-ready capability at an accessible price point. Highly recommended for riders who prioritize performance above all else.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance enthusiasts</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.8,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">Lifan KPV 150 Race Edition রিভিউ বাংলাদেশ 2024 - ট্র্যাক-রেডি পারফরম্যান্স বিস্ট</h1>
<p class="summary-text">Lifan KPV 150 Race Edition চাইনিজ রেসিং প্রযুক্তির শিখর হিসেবে দাঁড়িয়ে আছে, ৳৩,৮০,০০০ টাকায় পেশাদার-গ্রেড পারফরম্যান্স প্রদান করে। এই ট্র্যাক-ফোকাসড মেশিন আক্রমণাত্মক স্টাইলিং এর সাথে প্রকৃত রেসিং শংসাপত্র একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Lifan KPV 150 Race Edition মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc রেসিং ইঞ্জিন:</strong> লিকুইড-কুলড পারফরম্যান্স</li>
<li class="highlight-item"><strong>রেসিং সাসপেনশন:</strong> সামঞ্জস্যযোগ্য সেটআপ</li>
<li class="highlight-item"><strong>ডিজিটাল ড্যাশবোর্ড:</strong> শিফট লাইট সজ্জিত</li>
<li class="highlight-item"><strong>ট্র্যাক-রেডি:</strong> পেশাদার-গ্রেড</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Lifan KPV 150 Race Edition সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রকৃত রেসিং বংশধারা:</strong> ট্র্যাক-রেডি সেটআপ</li>
<li class="pro-item"><strong>চিত্তাকর্ষক শক্তি:</strong> শক্তিশালী ত্বরণ</li>
<li class="pro-item"><strong>পেশাদার সাসপেনশন:</strong> সামঞ্জস্যযোগ্য মান</li>
<li class="pro-item"><strong>প্রতিযোগিতামূলক মূল্য:</strong> পারফরম্যান্স মূল্য</li>
<li class="pro-item"><strong>রেসিং ফিচার:</strong> ব্যাপক প্যাকেজ</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Lifan KPV 150 Race Edition অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>সীমিত সেবা নেটওয়ার্ক:</strong> উপলব্ধতা উদ্বেগ</li>
<li class="con-item"><strong>বেশি রক্ষণাবেক্ষণ:</strong> রেসিং-স্পেক খরচ</li>
<li class="con-item"><strong>আক্রমণাত্মক পজিশন:</strong> আরাম আপস</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: Lifan KPV 150 Race Edition কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳৩,৮০,০০০ টাকায়, Lifan KPV 150 Race Edition একটি প্রকৃত পারফরম্যান্স মেশিন যা একটি সহজলভ্য মূল্যে ট্র্যাক-রেডি ক্ষমতা প্রদান করে। পারফরম্যান্স উৎসাহীদের জন্য অত্যন্ত সুপারিশকৃত।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পারফরম্যান্স উৎসাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Lifan KPV 150 Race Edition\n")
	return nil
}

func (s *LifanKPV150RaceEditionBatch28) GetName() string {
	return "LifanKPV150RaceEditionBatch28"
}
