package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type LifanXPECT150V2Batch28 struct {
	BaseSeeder
}

func NewLifanXPECT150V2Batch28Seeder() *LifanXPECT150V2Batch28 {
	return &LifanXPECT150V2Batch28{BaseSeeder: BaseSeeder{name: "LIFAN X-PECT 150 V2 Batch28 Review"}}
}

func (s *LifanXPECT150V2Batch28) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "LIFAN X-PECT 150 V2").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("lifan x-pect 150 v2 product not found")
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
<h1 class="review-main-title">LIFAN X-PECT 150 V2 Review Bangladesh 2024 - Advanced Sportbike Excellence</h1>
<p class="summary-text">The LIFAN X-PECT 150 V2 delivers modern sportbike performance at ৳3,20,000. With its 150cc liquid-cooled engine, smartphone connectivity, LED lighting, and 4.7/5.0 rating, it offers excellent technology and value for contemporary riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">LIFAN X-PECT 150 V2 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>Engine:</strong> 150cc liquid-cooled</li>
<li class="highlight-item"><strong>Features:</strong> Smartphone connectivity</li>
<li class="highlight-item"><strong>Lighting:</strong> Full LED system</li>
<li class="highlight-item"><strong>Safety:</strong> Dual-channel ABS</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">LIFAN X-PECT 150 V2 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Modern Features:</strong> Smartphone integration</li>
<li class="pro-item"><strong>Reliable Engine:</strong> Liquid-cooled performance</li>
<li class="pro-item"><strong>Good Efficiency:</strong> Balanced fuel economy</li>
<li class="pro-item"><strong>Contemporary Design:</strong> Modern styling</li>
<li class="pro-item"><strong>Value Pricing:</strong> Affordable sportbike</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">LIFAN X-PECT 150 V2 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Limited Network:</strong> Service availability</li>
<li class="con-item"><strong>Resale Value:</strong> Lower than Japanese</li>
<li class="con-item"><strong>Spare Parts:</strong> Availability concerns</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">LIFAN X-PECT 150 V2 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,20,000 - Modern sportbike</span></p>
<p class="value-item"><strong>Engine:</strong> <span>150cc liquid-cooled</span></p>
<p class="value-item"><strong>Features:</strong> <span>Smartphone connectivity</span></p>
<p class="value-item"><strong>Safety:</strong> <span>Dual-channel ABS</span></p>
<p class="value-item"><strong>Lighting:</strong> <span>Full LED</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">LIFAN X-PECT 150 V2 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Modern styling</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - Reliable power</li>
<li class="rating-item"><strong>Features:</strong> <span>4.8</span> - Advanced tech</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - Excellent pricing</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - Good build</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy LIFAN X-PECT 150 V2?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,20,000, the LIFAN X-PECT 150 V2 is an excellent choice for riders seeking modern sportbike technology with smartphone connectivity and contemporary features. Recommended for tech-savvy riders wanting affordable performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For modern riders</span></p>
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
<h1 class="review-main-title">LIFAN X-PECT 150 V2 রিভিউ বাংলাদেশ 2024 - উন্নত স্পোর্টবাইক উৎকর্ষতা</h1>
<p class="summary-text">LIFAN X-PECT 150 V2 ৳৩,২০,০০০ টাকায় আধুনিক স্পোর্টবাইক পারফরম্যান্স প্রদান করে। এর ১৫০সিসি লিকুইড-কুলড ইঞ্জিন, স্মার্টফোন সংযোগ, LED লাইটিং এবং ৪.৭/৫.০ রেটিং সহ, এটি সমসাময়িক রাইডারদের জন্য চমৎকার প্রযুক্তি এবং মূল্য অফার করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">LIFAN X-PECT 150 V2 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>ইঞ্জিন:</strong> ১৫০সিসি লিকুইড-কুলড</li>
<li class="highlight-item"><strong>ফিচার:</strong> স্মার্টফোন সংযোগ</li>
<li class="highlight-item"><strong>লাইটিং:</strong> সম্পূর্ণ LED সিস্টেম</li>
<li class="highlight-item"><strong>নিরাপত্তা:</strong> ডুয়াল-চ্যানেল ABS</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">LIFAN X-PECT 150 V2 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আধুনিক ফিচার:</strong> স্মার্টফোন ইন্টিগ্রেশন</li>
<li class="pro-item"><strong>নির্ভরযোগ্য ইঞ্জিন:</strong> লিকুইড-কুলড পারফরম্যান্স</li>
<li class="pro-item"><strong>ভাল দক্ষতা:</strong> সুষম জ্বালানি অর্থনীতি</li>
<li class="pro-item"><strong>সমসাময়িক ডিজাইন:</strong> আধুনিক স্টাইলিং</li>
<li class="pro-item"><strong>মূল্য মূল্যায়ন:</strong> সাশ্রয়ী স্পোর্টবাইক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">LIFAN X-PECT 150 V2 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>সীমিত নেটওয়ার্ক:</strong> সেবা উপলব্ধতা</li>
<li class="con-item"><strong>পুনঃবিক্রয় মূল্য:</strong> জাপানিদের চেয়ে কম</li>
<li class="con-item"><strong>খুচরা যন্ত্রাংশ:</strong> উপলব্ধতা উদ্বেগ</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: LIFAN X-PECT 150 V2 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳৩,২০,০০০ টাকায়, LIFAN X-PECT 150 V2 স্মার্টফোন সংযোগ এবং সমসাময়িক ফিচার সহ আধুনিক স্পোর্টবাইক প্রযুক্তি খুঁজছেন রাইডারদের জন্য একটি চমৎকার পছন্দ। সাশ্রয়ী পারফরম্যান্স চাওয়া প্রযুক্তি-সচেতন রাইডারদের জন্য সুপারিশকৃত।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আধুনিক রাইডারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created LIFAN X-PECT 150 V2\n")
	return nil
}

func (s *LifanXPECT150V2Batch28) GetName() string {
	return "LifanXPECT150V2Batch28"
}
