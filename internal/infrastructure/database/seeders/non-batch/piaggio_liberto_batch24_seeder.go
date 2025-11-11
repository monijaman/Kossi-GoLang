package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type PiaggioLibertoBatch24 struct {
	BaseSeeder
}

func NewPiaggioLibertoBatch24Seeder() *PiaggioLibertoBatch24 {
	return &PiaggioLibertoBatch24{BaseSeeder: BaseSeeder{name: "Piaggio Liberto Batch24 Review"}}
}

func (s *PiaggioLibertoBatch24) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Piaggio Liberto").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("piaggio liberto product not found")
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
<h1 class="review-main-title">Piaggio Liberto Review Bangladesh 2024 - Premium Italian Scooter</h1>
<p class="summary-text">The Piaggio Liberto is a premium 150cc scooter combining Italian heritage with modern features. Priced around ৳6,85,000, it delivers 11.8 bhp power, premium design, excellent comfort, impressive 57+ km/l efficiency, and outstanding value for riders seeking Italian premium character.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Piaggio Liberto Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Premium:</strong> Italian heritage</li>
<li class="highlight-item"><strong>11.8 bhp Power:</strong> Responsive capable</li>
<li class="highlight-item"><strong>Premium Design:</strong> Modern Italian</li>
<li class="highlight-item"><strong>57+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Piaggio Liberto Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Premium Design:</strong> Italian styled</li>
<li class="pro-item"><strong>Responsive Power:</strong> 11.8 bhp capable</li>
<li class="pro-item"><strong>Excellent Comfort:</strong> Modern features</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 57+ km/l great</li>
<li class="pro-item"><strong>Piaggio Heritage:</strong> Italian quality</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Piaggio Liberto Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳6,85,000 upscale</li>
<li class="con-item"><strong>Limited Service:</strong> Piaggio dealers</li>
<li class="con-item"><strong>Import Dependent:</strong> Parts availability</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Piaggio Liberto Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳6,85,000 - Premium Italian</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc scooter single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>11.8 bhp capable</span></p>
<p class="value-item"><strong>Torque:</strong> <span>12 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>142kg premium</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>57+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Piaggio Liberto Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Premium Italian</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Modern featured</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 11.8 bhp capable</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 57+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Piaggio heritage</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Piaggio Liberto?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳6,85,000, the Piaggio Liberto is perfect for riders seeking premium Italian character, 11.8 bhp responsive power, excellent 57+ km/l efficiency, and modern featured Italian heritage scooter.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For Italian premium scooter seekers</span></p>
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
<h1 class="review-main-title">পিয়াজিও লিবার্টো রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম ইতালীয় স্কুটার</h1>
<p class="summary-text">পিয়াজিও লিবার্টো একটি প্রিমিয়াম 150cc স্কুটার যা ইতালীয় ঐতিহ্য এবং আধুনিক বৈশিষ্ট্য একত্রিত করে। ৳6,85,000 টাকায় মূল্যায়িত, এটি 11.8 bhp শক্তি এবং প্রিমিয়াম ডিজাইন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">পিয়াজিও লিবার্টো মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc প্রিমিয়াম:</strong> ইতালীয় ঐতিহ্য</li>
<li class="highlight-item"><strong>11.8 bhp শক্তি:</strong> প্রতিক্রিয়াশীল সক্ষম</li>
<li class="highlight-item"><strong>প্রিমিয়াম ডিজাইন:</strong> আধুনিক ইতালীয়</li>
<li class="highlight-item"><strong>57+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">পিয়াজিও লিবার্টো সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রিমিয়াম ডিজাইন:</strong> ইতালীয় স্টাইল</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল শক্তি:</strong> 11.8 bhp সক্ষম</li>
<li class="pro-item"><strong>চমৎকার আরাম:</strong> আধুনিক বৈশিষ্ট্য</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 57+ km/l দুর্দান্ত</li>
<li class="pro-item"><strong>পিয়াজিও ঐতিহ্য:</strong> ইতালীয় গুণমান</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">পিয়াজিও লিবার্টো অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳6,85,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>সীমিত সেবা:</strong> পিয়াজিও ডিলার</li>
<li class="con-item"><strong>আমদানি নির্ভর:</strong> যন্ত্রাংশ উপলব্ধতা</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: পিয়াজিও লিবার্টো কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳6,85,000 টাকায়, পিয়াজিও লিবার্টো চালকদের জন্য নিখুঁত যারা প্রিমিয়াম ইতালীয় চরিত্র এবং আধুনিক বৈশিষ্ট্য চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ইতালীয় প্রিমিয়াম স্কুটার খোঁজার জন্য</span></p>
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

	fmt.Printf("   ✅ Created Piaggio Liberto\n")
	return nil
}

func (s *PiaggioLibertoBatch24) GetName() string {
	return "PiaggioLibertoBatch24"
}
