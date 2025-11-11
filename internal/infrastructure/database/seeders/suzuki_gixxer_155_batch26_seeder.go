package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type SuzukiGixxer155Batch26 struct {
	BaseSeeder
}

func NewSuzukiGixxer155Batch26Seeder() *SuzukiGixxer155Batch26 {
	return &SuzukiGixxer155Batch26{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer 155 Batch26 Review"}}
}

func (s *SuzukiGixxer155Batch26) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer 155").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki gixxer 155 product not found")
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
<h1 class="review-main-title">Suzuki Gixxer 155 Review Bangladesh 2024 - Balanced Street Fighter</h1>
<p class="summary-text">The Suzuki Gixxer 155 is a balanced street fighter combining sporty styling and everyday practicality. Priced around ৳3,45,000, it delivers 14.1 bhp power, muscular design, comfortable riding, excellent handling, and impressive 55+ km/l efficiency for urban warriors.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer 155 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>155cc Street:</strong> Fighter balanced</li>
<li class="highlight-item"><strong>14.1 bhp Power:</strong> Sporty capable</li>
<li class="highlight-item"><strong>Muscular Design:</strong> Aggressive contemporary</li>
<li class="highlight-item"><strong>55+ km/l Efficiency:</strong> Impressive economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer 155 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Muscular Styling:</strong> Aggressive street fighter</li>
<li class="pro-item"><strong>Balanced Performance:</strong> 14.1 bhp capable</li>
<li class="pro-item"><strong>Excellent Handling:</strong> Agile responsive</li>
<li class="pro-item"><strong>Comfortable Seating:</strong> Daily ride ready</li>
<li class="pro-item"><strong>Impressive Efficiency:</strong> 55+ km/l economical</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer 155 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Moderate Power:</strong> 14.1 bhp adequate</li>
<li class="con-item"><strong>Limited Features:</strong> Basic amenities</li>
<li class="con-item"><strong>Service Network:</strong> Growing availability</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer 155 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,45,000 - Street fighter</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>155cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>14.1 bhp sporty</span></p>
<p class="value-item"><strong>Torque:</strong> <span>14 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>140kg balanced</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>55+ km/l impressive</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer 155 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Muscular aggressive</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.6</span> - Daily comfortable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 14.1 bhp balanced</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 55+ km/l impressive</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - Suzuki reliable</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer 155?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,45,000, the Suzuki Gixxer 155 is perfect for urban warriors seeking muscular styling, balanced 14.1 bhp performance, excellent handling, and impressive 55+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For urban warriors</span></p>
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
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">সুজুকি গিক্সার 155 রিভিউ বাংলাদেশ 2024 - সুষম স্ট্রিট ফাইটার</h1>
<p class="summary-text">সুজুকি গিক্সার 155 একটি সুষম স্ট্রিট ফাইটার যা স্পোর্টি স্টাইলিং এবং প্রতিদিনের ব্যবহারিকতা একত্রিত করে। ৳3,45,000 টাকায় মূল্যায়িত, এটি 14.1 bhp শক্তি এবং পেশীবহুল ডিজাইন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি গিক্সার 155 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>155cc স্ট্রিট:</strong> ফাইটার সুষম</li>
<li class="highlight-item"><strong>14.1 bhp শক্তি:</strong> স্পোর্টি সক্ষম</li>
<li class="highlight-item"><strong>পেশীবহুল ডিজাইন:</strong> আক্রমণাত্মক সমসাময়িক</li>
<li class="highlight-item"><strong>55+ km/l দক্ষতা:</strong> চিত্তাকর্ষক অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি গিক্সার 155 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>পেশীবহুল স্টাইলিং:</strong> আক্রমণাত্মক স্ট্রিট ফাইটার</li>
<li class="pro-item"><strong>সুষম কর্মক্ষমতা:</strong> 14.1 bhp সক্ষম</li>
<li class="pro-item"><strong>চমৎকার হ্যান্ডলিং:</strong> চতুর প্রতিক্রিয়াশীল</li>
<li class="pro-item"><strong>আরামদায়ক আসন:</strong> দৈনিক যাত্রা প্রস্তুত</li>
<li class="pro-item"><strong>চিত্তাকর্ষক দক্ষতা:</strong> 55+ km/l সাশ্রয়ী</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি গিক্সার 155 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মধ্যম শক্তি:</strong> 14.1 bhp পর্যাপ্ত</li>
<li class="con-item"><strong>সীমিত বৈশিষ্ট্য:</strong> মৌলিক সুবিধাবিদ</li>
<li class="con-item"><strong>সেবা নেটওয়ার্ক:</strong> ক্রমবর্ধমান প্রাপ্যতা</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি গিক্সার 155 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,45,000 টাকায়, সুজুকি গিক্সার 155 শহুরে যোদ্ধাদের জন্য নিখুঁত যারা পেশীবহুল স্টাইলিং এবং সুষম 14.1 bhp কর্মক্ষমতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- শহুরে যোদ্ধাদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Suzuki Gixxer 155\n")
	return nil
}

func (s *SuzukiGixxer155Batch26) GetName() string {
	return "SuzukiGixxer155Batch26"
}
