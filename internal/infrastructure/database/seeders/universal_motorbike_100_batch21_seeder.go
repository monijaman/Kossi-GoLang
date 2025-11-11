package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type UniversalMotorBike100Batch21 struct {
	BaseSeeder
}

func NewUniversalMotorBike100Batch21Seeder() *UniversalMotorBike100Batch21 {
	return &UniversalMotorBike100Batch21{BaseSeeder: BaseSeeder{name: "Universal MotorBike 100 Batch21 Review"}}
}

func (s *UniversalMotorBike100Batch21) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Universal MotorBike 100").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("universal motorbike 100 product not found")
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
<h1 class="review-main-title">Universal MotorBike 100 Review Bangladesh 2024 - No-Frills Budget Commuter</h1>
<p class="summary-text">The Universal MotorBike 100 is a 100cc air-cooled ultra-budget commuter motorcycle representing no-frills practical transportation. Priced around ৳95,000, it delivers 7 bhp basic power, minimal maintenance, excellent 60+ km/l fuel economy, simple reliability, and ultimate value for budget-conscious commuters.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Universal MotorBike 100 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc Air-Cooled:</strong> Ultra-budget</li>
<li class="highlight-item"><strong>7 bhp Power:</strong> Basic practical</li>
<li class="highlight-item"><strong>60+ km/l Efficiency:</strong> Exceptional economy</li>
<li class="highlight-item"><strong>Minimal Maintenance:</strong> Simple robust</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Universal MotorBike 100 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 60+ km/l extraordinary</li>
<li class="pro-item"><strong>Ultra-Budget Price:</strong> ৳95,000 entry-level</li>
<li class="pro-item"><strong>Simple Reliability:</strong> No-frills robust</li>
<li class="pro-item"><strong>Minimal Maintenance:</strong> Basic parts easy</li>
<li class="pro-item"><strong>Maximum Value:</strong> Budget transportation focused</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Universal MotorBike 100 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Very Basic Power:</strong> 7 bhp minimal</li>
<li class="con-item"><strong>No Comfort:</strong> Bare essentials only</li>
<li class="con-item"><strong>Limited Features:</strong> Stripped-down design</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Universal MotorBike 100 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳95,000 - Ultra-budget entry-level</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>100cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>7 bhp basic</span></p>
<p class="value-item"><strong>Torque:</strong> <span>8 nm practical</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>100kg minimal</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>60+ km/l extraordinary</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Universal MotorBike 100 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Efficiency:</strong> <span>5.0</span> - 60+ km/l extraordinary</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.7</span> - No-frills simple</li>
<li class="rating-item"><strong>Value:</strong> <span>5.0</span> - ৳95,000 unbeatable</li>
<li class="rating-item"><strong>Maintenance:</strong> <span>4.8</span> - Minimal basic</li>
<li class="rating-item"><strong>Practicality:</strong> <span>4.5</span> - Budget transport</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Universal MotorBike 100?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳95,000, the MotorBike 100 is the ultimate budget commuter for cost-conscious riders seeking exceptional 60+ km/l efficiency, simple reliability, minimal maintenance, and maximum value for basic daily commuting needs.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For ultra-budget commuters</span></p>
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
<h1 class="review-main-title">ইউনিভার্সাল মোটরবাইক 100 রিভিউ বাংলাদেশ 2024 - কোনো-ফ্রিল বাজেট কমিউটার</h1>
<p class="summary-text">ইউনিভার্সাল মোটরবাইক 100 একটি 100cc এয়ার-কুল্ড অতি-বাজেট কমিউটার মোটরসাইকেল যা কোনো-ফ্রিল ব্যবহারিক পরিবহন প্রতিনিধিত্ব করে। ৳95,000 টাকায় মূল্যায়িত, এটি 7 bhp মৌলিক শক্তি এবং ব্যতিক্রমী 60+ km/l জ্বালানি অর্থনীতি প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইউনিভার্সাল মোটরবাইক 100 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc এয়ার-কুল্ড:</strong> অতি-বাজেট</li>
<li class="highlight-item"><strong>7 bhp শক্তি:</strong> মৌলিক ব্যবহারিক</li>
<li class="highlight-item"><strong>60+ km/l দক্ষতা:</strong> ব্যতিক্রমী অর্থনীতি</li>
<li class="highlight-item"><strong>ন্যূনতম রক্ষণাবেক্ষণ:</strong> সহজ শক্তিশালী</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইউনিভার্সাল মোটরবাইক 100 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>ব্যতিক্রমী দক্ষতা:</strong> 60+ km/l অসাধারণ</li>
<li class="pro-item"><strong>অতি-বাজেট মূল্য:</strong> ৳95,000 এন্ট্রি-লেভেল</li>
<li class="pro-item"><strong>সহজ নির্ভরযোগ্যতা:</strong> কোনো-ফ্রিল শক্তিশালী</li>
<li class="pro-item"><strong>ন্যূনতম রক্ষণাবেক্ষণ:</strong> মৌলিক যন্ত্রাংশ সহজ</li>
<li class="pro-item"><strong>সর্বোচ্চ মূল্য:</strong> বাজেট পরিবহন ফোকাসড</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইউনিভার্সাল মোটরবাইক 100 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>অতি মৌলিক শক্তি:</strong> 7 bhp ন্যূনতম</li>
<li class="con-item"><strong>কোনো আরাম নেই:</strong> শুধুমাত্র প্রয়োজনীয় বিষয়</li>
<li class="con-item"><strong>সীমিত বৈশিষ্ট্য:</strong> ছাঁটানো-ডাউন ডিজাইন</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইউনিভার্সাল মোটরবাইক 100 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳95,000 টাকায়, মোটরবাইক 100 খরচ-সচেতন চালকদের জন্য চূড়ান্ত বাজেট কমিউটার যারা ব্যতিক্রমী 60+ km/l দক্ষতা, সহজ নির্ভরযোগ্যতা এবং সর্বোচ্চ মূল্য চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অতি-বাজেট কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Universal MotorBike 100\n")
	return nil
}

func (s *UniversalMotorBike100Batch21) GetName() string {
	return "UniversalMotorBike100Batch21"
}
