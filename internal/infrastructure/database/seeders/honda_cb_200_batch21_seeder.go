package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HondaCB200Batch21 struct {
	BaseSeeder
}

func NewHondaCB200Batch21Seeder() *HondaCB200Batch21 {
	return &HondaCB200Batch21{BaseSeeder: BaseSeeder{name: "Honda CB 200 Batch21 Review"}}
}

func (s *HondaCB200Batch21) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CB 200").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda cb 200 product not found")
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
<h1 class="review-main-title">Honda CB 200 Review Bangladesh 2024 - Japanese Reliability Commuter</h1>
<p class="summary-text">The Honda CB 200 is a 200cc air-cooled reliable commuter motorcycle representing Japanese engineering heritage and rock-solid durability. Priced around ৳3,50,000, it delivers 17 bhp smooth power, excellent reliability, low maintenance, comfortable seating, and practical fuel efficiency for daily commuting with confidence.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CB 200 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>200cc Air-Cooled:</strong> Japanese reliable</li>
<li class="highlight-item"><strong>17 bhp Smooth:</strong> Practical power</li>
<li class="highlight-item"><strong>Excellent Reliability:</strong> Japanese heritage</li>
<li class="highlight-item"><strong>40+ km/l Efficiency:</strong> Practical economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CB 200 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Excellent Reliability:</strong> Japanese heritage trusted</li>
<li class="pro-item"><strong>Smooth 200cc:</strong> Practical power delivery</li>
<li class="pro-item"><strong>Low Maintenance:</strong> Simple robust design</li>
<li class="pro-item"><strong>Comfortable Seating:</strong> Daily riding friendly</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 40+ km/l practical</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CB 200 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Moderate Power:</strong> 17 bhp practical only</li>
<li class="con-item"><strong>Basic Styling:</strong> Functional design</li>
<li class="con-item"><strong>Premium Price:</strong> ৳3,50,000 Japanese brand</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CB 200 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,50,000 - Japanese premium reliable</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>200cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>17 bhp smooth</span></p>
<p class="value-item"><strong>Torque:</strong> <span>17 nm practical</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>130kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>40+ km/l practical</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CB 200 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Reliability:</strong> <span>4.9</span> - Japanese heritage</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.5</span> - 17 bhp practical</li>
<li class="rating-item"><strong>Maintenance:</strong> <span>4.8</span> - Simple robust</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Daily riding</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳3,50,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CB 200?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,50,000, the CB 200 is the reliable choice for practical commuters seeking Japanese heritage, excellent reliability, smooth 17 bhp power, low maintenance, and confident daily commuting.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For reliable commuters</span></p>
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
<h1 class="review-main-title">হোন্ডা CB 200 রিভিউ বাংলাদেশ 2024 - জাপানি নির্ভরযোগ্যতা কমিউটার</h1>
<p class="summary-text">হোন্ডা CB 200 একটি 200cc এয়ার-কুল্ড নির্ভরযোগ্য কমিউটার মোটরসাইকেল যা জাপানি ইঞ্জিনিয়ারিং ঐতিহ্য এবং শক্ত স্থায়িত্ব প্রতিনিধিত্ব করে। ৳3,50,000 টাকায় মূল্যায়িত, এটি 17 bhp মসৃণ শক্তি এবং উৎকৃষ্ট নির্ভরযোগ্যতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা CB 200 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>200cc এয়ার-কুল্ড:</strong> জাপানি নির্ভরযোগ্য</li>
<li class="highlight-item"><strong>17 bhp মসৃণ:</strong> ব্যবহারিক শক্তি</li>
<li class="highlight-item"><strong>উৎকৃষ্ট নির্ভরযোগ্যতা:</strong> জাপানি ঐতিহ্য</li>
<li class="highlight-item"><strong>40+ km/l দক্ষতা:</strong> ব্যবহারিক অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা CB 200 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>উৎকৃষ্ট নির্ভরযোগ্যতা:</strong> জাপানি ঐতিহ্য বিশ্বস্ত</li>
<li class="pro-item"><strong>মসৃণ 200cc:</strong> ব্যবহারিক শক্তি সরবরাহ</li>
<li class="pro-item"><strong>কম রক্ষণাবেক্ষণ:</strong> সহজ শক্তিশালী ডিজাইন</li>
<li class="pro-item"><strong>আরামদায়ক আসন:</strong> দৈনিক যাত্রা বান্ধব</li>
<li class="pro-item"><strong>ভালো দক্ষতা:</strong> 40+ km/l ব্যবহারিক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা CB 200 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মধ্যম শক্তি:</strong> 17 bhp মাত্র ব্যবহারিক</li>
<li class="con-item"><strong>মৌলিক স্টাইলিং:</strong> কার্যকরী ডিজাইন</li>
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳3,50,000 জাপানি ব্র্যান্ড</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হোন্ডা CB 200 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,50,000 টাকায়, CB 200 ব্যবহারিক কমিউটারদের জন্য নির্ভরযোগ্য পছন্দ যারা জাপানি ঐতিহ্য, উৎকৃষ্ট নির্ভরযোগ্যতা, মসৃণ 17 bhp শক্তি এবং কম রক্ষণাবেক্ষণ চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- নির্ভরযোগ্য কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Honda CB 200\n")
	return nil
}

func (s *HondaCB200Batch21) GetName() string {
	return "HondaCB200Batch21"
}
