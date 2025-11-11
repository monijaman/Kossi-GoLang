package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BajajDualDo125Batch24 struct {
	BaseSeeder
}

func NewBajajDualDo125Batch24Seeder() *BajajDualDo125Batch24 {
	return &BajajDualDo125Batch24{BaseSeeder: BaseSeeder{name: "Bajaj Dual Do 125 Batch24 Review"}}
}

func (s *BajajDualDo125Batch24) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Dual Do 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj dual do 125 product not found")
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
<h1 class="review-main-title">Bajaj Dual Do 125 Review Bangladesh 2024 - Unique Adventure Scooter</h1>
<p class="summary-text">The Bajaj Dual Do 125 is a unique adventure scooter combining urban and light off-road capability. Priced around ৳4,85,000, it features 8.6 bhp power, adventure styling, unique dual design, excellent 62+ km/l efficiency, and outstanding versatility for urban-plus riders seeking weekend adventure capability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Dual Do 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>Adventure Scooter:</strong> Unique design</li>
<li class="highlight-item"><strong>8.6 bhp Power:</strong> Adventure capable</li>
<li class="highlight-item"><strong>Urban + Off-Road:</strong> Dual capability</li>
<li class="highlight-item"><strong>62+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Dual Do 125 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Unique Design:</strong> Adventure oriented</li>
<li class="pro-item"><strong>Light Off-Road:</strong> Dual capability</li>
<li class="pro-item"><strong>Urban Practical:</strong> Daily commute ready</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 62+ km/l superior</li>
<li class="pro-item"><strong>Bajaj Reliability:</strong> Proven service network</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Dual Do 125 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Limited Power:</strong> 8.6 bhp modest</li>
<li class="con-item"><strong>Adventure Light:</strong> Not full capability</li>
<li class="con-item"><strong>Niche Category:</strong> Unique positioning</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Dual Do 125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳4,85,000 - Adventure oriented</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc scooter single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>8.6 bhp adventure</span></p>
<p class="value-item"><strong>Torque:</strong> <span>10 nm adventure</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>130kg adventure</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>62+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Dual Do 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Unique adventure</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.6</span> - Urban focused</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 8.6 bhp dual</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 62+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - Bajaj proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Dual Do 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳4,85,000, the Bajaj Dual Do 125 is ideal for urban riders seeking unique adventure capability, light off-road versatility, excellent 62+ km/l efficiency, and weekend adventure capability with urban practicality.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For adventure-minded urban riders</span></p>
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
<h1 class="review-main-title">বাজাজ ডুয়াল ডু 125 রিভিউ বাংলাদেশ 2024 - অনন্য অ্যাডভেঞ্চার স্কুটার</h1>
<p class="summary-text">বাজাজ ডুয়াল ডু 125 একটি অনন্য অ্যাডভেঞ্চার স্কুটার যা শহুরে এবং হালকা অফ-রোড সক্ষমতা একত্রিত করে। ৳4,85,000 টাকায় মূল্যায়িত, এটি 8.6 bhp শক্তি, অ্যাডভেঞ্চার স্টাইলিং এবং অনন্য ডুয়াল ডিজাইন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ ডুয়াল ডু 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>অ্যাডভেঞ্চার স্কুটার:</strong> অনন্য ডিজাইন</li>
<li class="highlight-item"><strong>8.6 bhp শক্তি:</strong> অ্যাডভেঞ্চার সক্ষম</li>
<li class="highlight-item"><strong>শহুরে + অফ-রোড:</strong> ডুয়াল সক্ষমতা</li>
<li class="highlight-item"><strong>62+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ ডুয়াল ডু 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>অনন্য ডিজাইন:</strong> অ্যাডভেঞ্চার ভিত্তিক</li>
<li class="pro-item"><strong>হালকা অফ-রোড:</strong> ডুয়াল সক্ষমতা</li>
<li class="pro-item"><strong>শহুরে ব্যবহারিক:</strong> দৈনন্দিন যাতায়াত প্রস্তুত</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 62+ km/l উচ্চতর</li>
<li class="pro-item"><strong>বাজাজ নির্ভরযোগ্যতা:</strong> প্রমাণিত সেবা নেটওয়ার্ক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ ডুয়াল ডু 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>সীমিত শক্তি:</strong> 8.6 bhp মামুলি</li>
<li class="con-item"><strong>হালকা অ্যাডভেঞ্চার:</strong> সম্পূর্ণ সক্ষমতা নয়</li>
<li class="con-item"><strong>নিশ বিভাগ:</strong> অনন্য অবস্থান</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ ডুয়াল ডু 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳4,85,000 টাকায়, বাজাজ ডুয়াল ডু 125 শহুরে চালকদের জন্য আদর্শ যারা অনন্য অ্যাডভেঞ্চার সক্ষমতা এবং সপ্তাহান্তে অ্যাডভেঞ্চার ক্ষমতা খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অ্যাডভেঞ্চার মনের শহুরে চালকদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Bajaj Dual Do 125\n")
	return nil
}

func (s *BajajDualDo125Batch24) GetName() string {
	return "BajajDualDo125Batch24"
}
