package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type SuzukiGixxer250Batch24 struct {
	BaseSeeder
}

func NewSuzukiGixxer250Batch24Seeder() *SuzukiGixxer250Batch24 {
	return &SuzukiGixxer250Batch24{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer 250 Batch24 Review"}}
}

func (s *SuzukiGixxer250Batch24) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer 250").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki gixxer 250 product not found")
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
<h1 class="review-main-title">Suzuki Gixxer 250 Review Bangladesh 2024 - Premium Street Bike</h1>
<p class="summary-text">The Suzuki Gixxer 250 is a premium 250cc street bike combining style and performance. Priced around ৳5,45,000, it delivers 26.5 bhp power, modern styling, comfortable riding, impressive 52+ km/l efficiency, and outstanding value for riders seeking premium street bike character.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer 250 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>250cc Street:</strong> Premium styled</li>
<li class="highlight-item"><strong>26.5 bhp Power:</strong> Thrilling capable</li>
<li class="highlight-item"><strong>Modern Design:</strong> Urban ready</li>
<li class="highlight-item"><strong>52+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer 250 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Premium Styling:</strong> Modern design</li>
<li class="pro-item"><strong>Thrilling Power:</strong> 26.5 bhp capable</li>
<li class="pro-item"><strong>Comfortable Riding:</strong> Urban friendly</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 52+ km/l great</li>
<li class="pro-item"><strong>Suzuki Reliability:</strong> Proven quality</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer 250 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳5,45,000 upscale</li>
<li class="con-item"><strong>Limited Service:</strong> Suzuki dealers only</li>
<li class="con-item"><strong>Parts Availability:</strong> Import dependent</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer 250 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳5,45,000 - Premium street</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>250cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>26.5 bhp capable</span></p>
<p class="value-item"><strong>Torque:</strong> <span>22.2 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>155kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>52+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer 250 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Premium modern</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Urban comfortable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 26.5 bhp capable</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 52+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Suzuki proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer 250?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳5,45,000, the Suzuki Gixxer 250 is perfect for riders seeking premium street bike character, 26.5 bhp thrilling power, modern styling, and excellent 52+ km/l efficiency in an urban-ready package.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium street bike seekers</span></p>
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
<h1 class="review-main-title">সুজুকি গিক্সার 250 রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম স্ট্রিট বাইক</h1>
<p class="summary-text">সুজুকি গিক্সার 250 একটি প্রিমিয়াম 250cc স্ট্রিট বাইক যা শৈলী এবং পারফরম্যান্স একত্রিত করে। ৳5,45,000 টাকায় মূল্যায়িত, এটি 26.5 bhp শক্তি এবং আধুনিক স্টাইলিং বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি গিক্সার 250 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>250cc স্ট্রিট:</strong> প্রিমিয়াম স্টাইল</li>
<li class="highlight-item"><strong>26.5 bhp শক্তি:</strong> রোমাঞ্চকর সক্ষম</li>
<li class="highlight-item"><strong>আধুনিক ডিজাইন:</strong> শহুরে প্রস্তুত</li>
<li class="highlight-item"><strong>52+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি গিক্সার 250 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রিমিয়াম স্টাইলিং:</strong> আধুনিক ডিজাইন</li>
<li class="pro-item"><strong>রোমাঞ্চকর শক্তি:</strong> 26.5 bhp সক্ষম</li>
<li class="pro-item"><strong>আরামদায়ক রাইডিং:</strong> শহুরে বান্ধব</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 52+ km/l দুর্দান্ত</li>
<li class="pro-item"><strong>সুজুকি নির্ভরযোগ্যতা:</strong> প্রমাণিত গুণমান</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি গিক্সার 250 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳5,45,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>সীমিত সেবা:</strong> সুজুকি ডিলার শুধু</li>
<li class="con-item"><strong>যন্ত্রাংশ উপলব্ধতা:</strong> আমদানি নির্ভর</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি গিক্সার 250 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳5,45,000 টাকায়, সুজুকি গিক্সার 250 চালকদের জন্য নিখুঁত যারা প্রিমিয়াম স্ট্রিট বাইক চরিত্র এবং রোমাঞ্চকর 26.5 bhp শক্তি চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম স্ট্রিট বাইক খোঁজার জন্য</span></p>
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

	fmt.Printf("   ✅ Created Suzuki Gixxer 250\n")
	return nil
}

func (s *SuzukiGixxer250Batch24) GetName() string {
	return "SuzukiGixxer250Batch24"
}
