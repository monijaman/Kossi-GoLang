package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HondaActiva6Batch23 struct {
	BaseSeeder
}

func NewHondaActiva6Batch23Seeder() *HondaActiva6Batch23 {
	return &HondaActiva6Batch23{BaseSeeder: BaseSeeder{name: "Honda Activa 6 Batch23 Review"}}
}

func (s *HondaActiva6Batch23) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda Activa 6").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda activa 6 product not found")
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
<h1 class="review-main-title">Honda Activa 6 Review Bangladesh 2024 - India's Best-Selling Automatic Scooter Icon</h1>
<p class="summary-text">The Honda Activa 6 is India's best-selling 110cc automatic scooter icon representing ultimate commuter reliability. Priced around ৳3,85,000, it delivers 8 bhp steady power, automatic transmission, comfortable ergonomics, exceptional 60+ km/l efficiency, and outstanding value for daily practical commuters.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda Activa 6 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc Automatic:</strong> Best-selling icon</li>
<li class="highlight-item"><strong>8 bhp Power:</strong> Steady reliable</li>
<li class="highlight-item"><strong>Automatic Transmission:</strong> Hassle-free</li>
<li class="highlight-item"><strong>60+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda Activa 6 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Best-Selling Reputation:</strong> Proven workhorse</li>
<li class="pro-item"><strong>Automatic Transmission:</strong> Zero clutch hassle</li>
<li class="pro-item"><strong>Comfortable Ergonomics:</strong> Daily commuting</li>
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 60+ km/l superior</li>
<li class="pro-item"><strong>Honda Reliability:</strong> Global trust</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda Activa 6 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Low Power:</strong> 8 bhp basic</li>
<li class="con-item"><strong>Modest Styling:</strong> Conservative design</li>
<li class="con-item"><strong>Premium Price:</strong> ৳3,85,000</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda Activa 6 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,85,000 - Premium practical</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>110cc automatic scooter</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>8 bhp steady</span></p>
<p class="value-item"><strong>Torque:</strong> <span>9 nm reliable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>112kg automatic</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>60+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda Activa 6 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Reliability:</strong> <span>4.9</span> - Best-selling proven</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Ergonomic daily</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 60+ km/l excellent</li>
<li class="rating-item"><strong>Transmission:</strong> <span>4.9</span> - Automatic seamless</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳3,85,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda Activa 6?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,85,000, the Honda Activa 6 is perfect for practical daily commuters seeking India's best-selling automatic scooter, seamless transmission, exceptional 60+ km/l efficiency, and reliable workhorse performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For daily practical commuters</span></p>
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
<h1 class="review-main-title">হোন্ডা অ্যাক্টিভা 6 রিভিউ বাংলাদেশ 2024 - ভারতের সেরা বিক্রয় অটোমেটিক স্কুটার আইকন</h1>
<p class="summary-text">হোন্ডা অ্যাক্টিভা 6 ভারতের সেরা বিক্রয় 110cc অটোমেটিক স্কুটার আইকন যা চূড়ান্ত কমিউটার নির্ভরযোগ্যতা প্রতিনিধিত্ব করে। ৳3,85,000 টাকায় মূল্যায়িত, এটি 8 bhp স্থিতিশীল শক্তি এবং অটোমেটিক ট্রান্সমিশন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা অ্যাক্টিভা 6 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc অটোমেটিক:</strong> সেরা বিক্রয় আইকন</li>
<li class="highlight-item"><strong>8 bhp শক্তি:</strong> স্থিতিশীল নির্ভরযোগ্য</li>
<li class="highlight-item"><strong>অটোমেটিক ট্রান্সমিশন:</strong> ঝামেলা-মুক্ত</li>
<li class="highlight-item"><strong>60+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা অ্যাক্টিভা 6 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>সেরা বিক্রয় খ্যাতি:</strong> প্রমাণিত কর্মঘোড়া</li>
<li class="pro-item"><strong>অটোমেটিক ট্রান্সমিশন:</strong> শূন্য ক্লাচ ঝামেলা</li>
<li class="pro-item"><strong>আরামদায়ক এর্গোনমিক্স:</strong> দৈনন্দিন যাতায়াত</li>
<li class="pro-item"><strong>ব্যতিক্রমী দক্ষতা:</strong> 60+ km/l উচ্চতর</li>
<li class="pro-item"><strong>হোন্ডা নির্ভরযোগ্যতা:</strong> বিশ্বব্যাপী বিশ্বাস</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা অ্যাক্টিভা 6 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>কম শক্তি:</strong> 8 bhp মৌলিক</li>
<li class="con-item"><strong>মনোরম স্টাইলিং:</strong> রক্ষণশীল ডিজাইন</li>
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳3,85,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হোন্ডা অ্যাক্টিভা 6 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,85,000 টাকায়, হোন্ডা অ্যাক্টিভা 6 ব্যবহারিক দৈনন্দিন কমিউটারদের জন্য নিখুঁত যারা ভারতের সেরা বিক্রয় অটোমেটিক স্কুটার, নিরবচ্ছিন্ন ট্রান্সমিশন এবং চমৎকার 60+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- দৈনন্দিন ব্যবহারিক কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Honda Activa 6\n")
	return nil
}

func (s *HondaActiva6Batch23) GetName() string {
	return "HondaActiva6Batch23"
}
