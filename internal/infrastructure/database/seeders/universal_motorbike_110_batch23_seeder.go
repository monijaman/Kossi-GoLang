package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type UniversalMotorbike110Batch23 struct {
	BaseSeeder
}

func NewUniversalMotorbike110Batch23Seeder() *UniversalMotorbike110Batch23 {
	return &UniversalMotorbike110Batch23{BaseSeeder: BaseSeeder{name: "Universal Motorbike 110 Batch23 Review"}}
}

func (s *UniversalMotorbike110Batch23) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Universal Motorbike 110").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("universal motorbike 110 product not found")
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
<h1 class="review-main-title">Universal Motorbike 110 Review Bangladesh 2024 - Bangladesh's Affordable Daily Workhorse</h1>
<p class="summary-text">The Universal Motorbike 110 is Bangladesh's affordable 110cc daily workhorse representing practical everyday utility. Priced around ৳1,45,000, it delivers 9.5 bhp practical power, robust simple design, durable construction, excellent 75+ km/l efficiency, and outstanding value for practical daily workers seeking reliable affordable utility.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Universal Motorbike 110 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc Air-Cooled:</strong> Daily workhorse</li>
<li class="highlight-item"><strong>9.5 bhp Power:</strong> Practical steady</li>
<li class="highlight-item"><strong>Robust Design:</strong> Utility focused</li>
<li class="highlight-item"><strong>75+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Universal Motorbike 110 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Excellent Efficiency:</strong> 75+ km/l superior</li>
<li class="pro-item"><strong>Robust Durability:</strong> Daily workhorse</li>
<li class="pro-item"><strong>Simple Maintenance:</strong> Easy upkeep</li>
<li class="pro-item"><strong>Affordability:</strong> ৳1,45,000 value</li>
<li class="pro-item"><strong>Local Support:</strong> Bangladesh made</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Universal Motorbike 110 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Modest Power:</strong> 9.5 bhp basic</li>
<li class="con-item"><strong>Basic Styling:</strong> Utilitarian only</li>
<li class="con-item"><strong>Limited Features:</strong> Minimal comfort</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Universal Motorbike 110 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,45,000 - Practical work</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>110cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>9.5 bhp practical</span></p>
<p class="value-item"><strong>Torque:</strong> <span>9.8 nm steady</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>110kg practical</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>75+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Universal Motorbike 110 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Durability:</strong> <span>4.7</span> - Workhorse built</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 75+ km/l excellent</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.7</span> - Steady proven</li>
<li class="rating-item"><strong>Practicality:</strong> <span>4.7</span> - Utility daily</li>
<li class="rating-item"><strong>Value:</strong> <span>4.8</span> - ৳1,45,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Universal Motorbike 110?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,45,000, the Universal Motorbike 110 is perfect for practical daily workers seeking Bangladesh's reliable workhorse, robust durability, excellent 75+ km/l efficiency, and affordable practical everyday utility.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For practical daily workers</span></p>
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
<h1 class="review-main-title">ইউনিভার্সাল মোটরবাইক 110 রিভিউ বাংলাদেশ 2024 - বাংলাদেশের সাশ্রয়ী দৈনন্দিন কর্মঘোড়া</h1>
<p class="summary-text">ইউনিভার্সাল মোটরবাইক 110 বাংলাদেশের সাশ্রয়ী 110cc দৈনন্দিন কর্মঘোড়া যা ব্যবহারিক দৈনন্দিন উপযোগিতা প্রতিনিধিত্ব করে। ৳1,45,000 টাকায় মূল্যায়িত, এটি 9.5 bhp ব্যবহারিক শক্তি এবং শক্তিশালী সরল ডিজাইন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইউনিভার্সাল মোটরবাইক 110 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc এয়ার-কুল্ড:</strong> দৈনন্দিন কর্মঘোড়া</li>
<li class="highlight-item"><strong>9.5 bhp শক্তি:</strong> ব্যবহারিক স্থিতিশীল</li>
<li class="highlight-item"><strong>শক্তিশালী ডিজাইন:</strong> উপযোগিতা ফোকাসড</li>
<li class="highlight-item"><strong>75+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইউনিভার্সাল মোটরবাইক 110 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 75+ km/l উচ্চতর</li>
<li class="pro-item"><strong>শক্তিশালী স্থায়িত্ব:</strong> দৈনন্দিন কর্মঘোড়া</li>
<li class="pro-item"><strong>সহজ রক্ষণাবেক্ষণ:</strong> সহজ রক্ষণাবেক্ষণ</li>
<li class="pro-item"><strong>সামর্থ্য:</strong> ৳1,45,000 মূল্য</li>
<li class="pro-item"><strong>স্থানীয় সমর্থন:</strong> বাংলাদেশ তৈরি</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইউনিভার্সাল মোটরবাইক 110 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>বিনম্র শক্তি:</strong> 9.5 bhp মৌলিক</li>
<li class="con-item"><strong>মৌলিক স্টাইলিং:</strong> ইউটিলিটারিয়ান শুধুমাত্র</li>
<li class="con-item"><strong>সীমিত বৈশিষ্ট্য:</strong> ন্যূনতম আরাম</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইউনিভার্সাল মোটরবাইক 110 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,45,000 টাকায়, ইউনিভার্সাল মোটরবাইক 110 ব্যবহারিক দৈনন্দিন কর্মীদের জন্য নিখুঁত যারা বাংলাদেশের নির্ভরযোগ্য কর্মঘোড়া, শক্তিশালী স্থায়িত্ব এবং চমৎকার 75+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ব্যবহারিক দৈনন্দিন কর্মীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Universal Motorbike 110\n")
	return nil
}

func (s *UniversalMotorbike110Batch23) GetName() string {
	return "UniversalMotorbike110Batch23"
}
