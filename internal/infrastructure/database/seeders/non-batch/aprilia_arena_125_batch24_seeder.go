package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type ApriliaArena125Batch24 struct {
	BaseSeeder
}

func NewApriliaArena125Batch24Seeder() *ApriliaArena125Batch24 {
	return &ApriliaArena125Batch24{BaseSeeder: BaseSeeder{name: "Aprilia Arena 125 Batch24 Review"}}
}

func (s *ApriliaArena125Batch24) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Aprilia Arena 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("aprilia arena 125 product not found")
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
<h1 class="review-main-title">Aprilia Arena 125 Review Bangladesh 2024 - Italy's Modern Urban Scooter</h1>
<p class="summary-text">The Aprilia Arena 125 is Italy's modern 125cc urban scooter representing contemporary European design. Priced around ৳5,75,000, it delivers 10 bhp responsive power, modern minimalist styling, spacious comfortable seating, excellent 58+ km/l efficiency, and outstanding value for urban riders seeking modern European character daily.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Aprilia Arena 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Modern Scooter:</strong> European design</li>
<li class="highlight-item"><strong>10 bhp Power:</strong> Responsive nimble</li>
<li class="highlight-item"><strong>Contemporary Styling:</strong> Urban aesthetic</li>
<li class="highlight-item"><strong>58+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Aprilia Arena 125 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Modern Design:</strong> Contemporary European</li>
<li class="pro-item"><strong>Spacious Seating:</strong> Comfortable urban</li>
<li class="pro-item"><strong>Responsive Power:</strong> 10 bhp nimble</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 58+ km/l superior</li>
<li class="pro-item"><strong>Aprilia Quality:</strong> Italian engineering</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Aprilia Arena 125 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳5,75,000 upscale</li>
<li class="con-item"><strong>Limited Service:</strong> Aprilia dealers</li>
<li class="con-item"><strong>Urban Focused:</strong> Limited off-road</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Aprilia Arena 125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳5,75,000 - Premium European</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc scooter single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>10 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>10.8 nm nimble</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>135kg modern</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>58+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Aprilia Arena 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Modern contemporary</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Spacious urban</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 10 bhp responsive</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 58+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Italian engineering</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Aprilia Arena 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳5,75,000, the Aprilia Arena 125 is perfect for urban riders seeking Italy's modern scooter, contemporary design, spacious comfort, excellent 58+ km/l efficiency, and premium European character.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For modern urban scooter seekers</span></p>
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
<h1 class="review-main-title">অ্যাপ্রিলিয়া অ্যারেনা 125 রিভিউ বাংলাদেশ 2024 - ইতালির আধুনিক শহুরে স্কুটার</h1>
<p class="summary-text">অ্যাপ্রিলিয়া অ্যারেনা 125 ইতালির আধুনিক 125cc শহুরে স্কুটার যা সমসাময়িক ইউরোপীয় ডিজাইন প্রতিনিধিত্ব করে। ৳5,75,000 টাকায় মূল্যায়িত, এটি 10 bhp প্রতিক্রিয়াশীল শক্তি এবং আধুনিক ন্যূনতমবাদী স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">অ্যাপ্রিলিয়া অ্যারেনা 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc আধুনিক স্কুটার:</strong> ইউরোপীয় ডিজাইন</li>
<li class="highlight-item"><strong>10 bhp শক্তি:</strong> প্রতিক্রিয়াশীল চটপটে</li>
<li class="highlight-item"><strong>সমসাময়িক স্টাইলিং:</strong> শহুরে নান্দনিকতা</li>
<li class="highlight-item"><strong>58+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">অ্যাপ্রিলিয়া অ্যারেনা 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আধুনিক ডিজাইন:</strong> সমসাময়িক ইউরোপীয়</li>
<li class="pro-item"><strong>প্রশস্ত আসন:</strong> আরামদায়ক শহুরে</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল শক্তি:</strong> 10 bhp চটপটে</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 58+ km/l উচ্চতর</li>
<li class="pro-item"><strong>অ্যাপ্রিলিয়া গুণমান:</strong> ইতালীয় প্রকৌশল</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">অ্যাপ্রিলিয়া অ্যারেনা 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳5,75,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>সীমিত সেবা:</strong> অ্যাপ্রিলিয়া ডিলার</li>
<li class="con-item"><strong>শহুরে ফোকাসড:</strong> সীমিত অফ-রোড</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: অ্যাপ্রিলিয়া অ্যারেনা 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳5,75,000 টাকায়, অ্যাপ্রিলিয়া অ্যারেনা 125 শহুরে চালকদের জন্য নিখুঁত যারা ইতালির আধুনিক স্কুটার, সমসাময়িক ডিজাইন, চমৎকার 58+ km/l দক্ষতা এবং প্রিমিয়াম ইউরোপীয় চরিত্র চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আধুনিক শহুরে স্কুটার খোঁজার জন্য</span></p>
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

	fmt.Printf("   ✅ Created Aprilia Arena 125\n")
	return nil
}

func (s *ApriliaArena125Batch24) GetName() string {
	return "ApriliaArena125Batch24"
}
