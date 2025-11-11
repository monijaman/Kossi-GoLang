package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type YamahaFZS150Batch25 struct {
	BaseSeeder
}

func NewYamahaFZS150Batch25Seeder() *YamahaFZS150Batch25 {
	return &YamahaFZS150Batch25{BaseSeeder: BaseSeeder{name: "Yamaha FZS 150 Batch25 Review"}}
}

func (s *YamahaFZS150Batch25) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha FZS 150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha fzs 150 product not found")
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
<h1 class="review-main-title">Yamaha FZS 150 Review Bangladesh 2024 - Premium Street Commuter</h1>
<p class="summary-text">The Yamaha FZS 150 is a premium street commuter combining modern design with reliable performance. Priced around ৳4,25,000, it delivers 16.1 bhp power, modern styling, comfortable ergonomics, excellent 62+ km/l efficiency, and outstanding value for premium commuter seekers.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha FZS 150 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Street:</strong> Premium commuter</li>
<li class="highlight-item"><strong>16.1 bhp Power:</strong> Responsive capable</li>
<li class="highlight-item"><strong>Modern Design:</strong> Contemporary styled</li>
<li class="highlight-item"><strong>62+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha FZS 150 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Modern Design:</strong> Premium styled</li>
<li class="pro-item"><strong>Responsive Power:</strong> 16.1 bhp capable</li>
<li class="pro-item"><strong>Comfortable Ergonomics:</strong> Daily friendly</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 62+ km/l great</li>
<li class="pro-item"><strong>Yamaha Quality:</strong> Proven reliability</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha FZS 150 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳4,25,000 upscale</li>
<li class="con-item"><strong>Limited Service:</strong> Yamaha dealers only</li>
<li class="con-item"><strong>Parts Availability:</strong> Import dependent</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha FZS 150 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳4,25,000 - Premium commuter</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>16.1 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>14 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>146kg modern</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>62+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha FZS 150 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Modern premium</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Ergonomic friendly</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 16.1 bhp capable</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 62+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Yamaha proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha FZS 150?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳4,25,000, the Yamaha FZS 150 is perfect for premium commuter seekers desiring modern design, responsive 16.1 bhp power, comfortable ergonomics, and excellent 62+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium commuter enthusiasts</span></p>
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
<h1 class="review-main-title">ইয়ামাহা FZS 150 রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম স্ট্রিট যাতায়াত</h1>
<p class="summary-text">ইয়ামাহা FZS 150 একটি প্রিমিয়াম স্ট্রিট যাতায়াত যা আধুনিক ডিজাইন এবং নির্ভরযোগ্য পারফরম্যান্স একত্রিত করে। ৳4,25,000 টাকায় মূল্যায়িত, এটি 16.1 bhp শক্তি এবং আধুনিক স্টাইলিং বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা FZS 150 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc স্ট্রিট:</strong> প্রিমিয়াম যাতায়াত</li>
<li class="highlight-item"><strong>16.1 bhp শক্তি:</strong> প্রতিক্রিয়াশীল সক্ষম</li>
<li class="highlight-item"><strong>আধুনিক ডিজাইন:</strong> সমসাময়িক স্টাইল</li>
<li class="highlight-item"><strong>62+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা FZS 150 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আধুনিক ডিজাইন:</strong> প্রিমিয়াম স্টাইল</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল শক্তি:</strong> 16.1 bhp সক্ষম</li>
<li class="pro-item"><strong>আরামদায়ক এরগনমিক্স:</strong> দৈনিক বান্ধব</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 62+ km/l দুর্দান্ত</li>
<li class="pro-item"><strong>ইয়ামাহা গুণমান:</strong> প্রমাণিত নির্ভরযোগ্যতা</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা FZS 150 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳4,25,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>সীমিত সেবা:</strong> ইয়ামাহা ডিলার শুধু</li>
<li class="con-item"><strong>যন্ত্রাংশ উপলব্ধতা:</strong> আমদানি নির্ভর</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা FZS 150 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳4,25,000 টাকায়, ইয়ামাহা FZS 150 প্রিমিয়াম যাতায়াত খোঁজা চালকদের জন্য নিখুঁত যারা আধুনিক ডিজাইন এবং প্রতিক্রিয়াশীল 16.1 bhp শক্তি চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম যাতায়াত উত্সাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Yamaha FZS 150\n")
	return nil
}

func (s *YamahaFZS150Batch25) GetName() string {
	return "YamahaFZS150Batch25"
}
