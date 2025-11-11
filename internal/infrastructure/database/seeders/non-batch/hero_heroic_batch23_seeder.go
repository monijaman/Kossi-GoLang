package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HeroHeroicBatch23 struct {
	BaseSeeder
}

func NewHeroHeroicBatch23Seeder() *HeroHeroicBatch23 {
	return &HeroHeroicBatch23{BaseSeeder: BaseSeeder{name: "Hero Heroic Batch23 Review"}}
}

func (s *HeroHeroicBatch23) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero Heroic").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero heroic product not found")
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
<h1 class="review-main-title">Hero Heroic Review Bangladesh 2024 - India's Practical Work-Duty Commuter</h1>
<p class="summary-text">The Hero Heroic is India's practical 97cc work-duty commuter representing essential workhorse functionality. Priced around ৳1,15,000, it delivers 8 bhp steady power, robust simplified design, durable construction, exceptional 80+ km/l efficiency, and supreme value for hard-working practical commuters seeking reliable duty vehicle.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero Heroic Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>97cc Air-Cooled:</strong> Work-duty</li>
<li class="highlight-item"><strong>8 bhp Power:</strong> Steady practical</li>
<li class="highlight-item"><strong>Robust Design:</strong> Workhorse character</li>
<li class="highlight-item"><strong>80+ km/l Efficiency:</strong> Exceptional economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero Heroic Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 80+ km/l ultimate</li>
<li class="pro-item"><strong>Robust Durability:</strong> Workhorse built</li>
<li class="pro-item"><strong>Simple Maintenance:</strong> Easy service</li>
<li class="pro-item"><strong>Affordability:</strong> ৳1,15,000 value</li>
<li class="pro-item"><strong>Hero Reliability:</strong> Proven duty</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero Heroic Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Low Power:</strong> 8 bhp basic</li>
<li class="con-item"><strong>Basic Styling:</strong> Utilitarian design</li>
<li class="con-item"><strong>Minimal Features:</strong> Practical only</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero Heroic Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,15,000 - Work value</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>97cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>8 bhp steady</span></p>
<p class="value-item"><strong>Torque:</strong> <span>8 nm practical</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>98kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>80+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero Heroic Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Durability:</strong> <span>4.8</span> - Workhorse built</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.9</span> - 80+ km/l exceptional</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.8</span> - Hero proven</li>
<li class="rating-item"><strong>Practicality:</strong> <span>4.7</span> - Work-duty focused</li>
<li class="rating-item"><strong>Value:</strong> <span>4.8</span> - ৳1,15,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero Heroic?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,15,000, the Hero Heroic is perfect for hard-working practical commuters seeking India's reliable work-duty vehicle, robust durability, exceptional 80+ km/l efficiency, and supreme affordable practical value.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For practical work-duty commuters</span></p>
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
<h1 class="review-main-title">হিরো হিরোইক রিভিউ বাংলাদেশ 2024 - ভারতের ব্যবহারিক কর্ম-দায়িত্ব কমিউটার</h1>
<p class="summary-text">হিরো হিরোইক ভারতের ব্যবহারিক 97cc কর্ম-দায়িত্ব কমিউটার যা অপরিহার্য কর্মঘোড়া কার্যকারিতা প্রতিনিধিত্ব করে। ৳1,15,000 টাকায় মূল্যায়িত, এটি 8 bhp স্থিতিশীল শক্তি এবং শক্তিশালী সরলীকৃত ডিজাইন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো হিরোইক মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>97cc এয়ার-কুল্ড:</strong> কর্ম-দায়িত্ব</li>
<li class="highlight-item"><strong>8 bhp শক্তি:</strong> স্থিতিশীল ব্যবহারিক</li>
<li class="highlight-item"><strong>শক্তিশালী ডিজাইন:</strong> কর্মঘোড়া চরিত্র</li>
<li class="highlight-item"><strong>80+ km/l দক্ষতা:</strong> ব্যতিক্রমী অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো হিরোইক সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>ব্যতিক্রমী দক্ষতা:</strong> 80+ km/l চূড়ান্ত</li>
<li class="pro-item"><strong>শক্তিশালী স্থায়িত্ব:</strong> কর্মঘোড়া নির্মিত</li>
<li class="pro-item"><strong>সহজ রক্ষণাবেক্ষণ:</strong> সহজ সেবা</li>
<li class="pro-item"><strong>সামর্থ্য:</strong> ৳1,15,000 মূল্য</li>
<li class="pro-item"><strong>হিরো নির্ভরযোগ্যতা:</strong> প্রমাণিত দায়িত্ব</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো হিরোইক অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>কম শক্তি:</strong> 8 bhp মৌলিক</li>
<li class="con-item"><strong>মৌলিক স্টাইলিং:</strong> ইউটিলিটারিয়ান ডিজাইন</li>
<li class="con-item"><strong>ন্যূনতম বৈশিষ্ট্য:</strong> ব্যবহারিক শুধুমাত্র</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো হিরোইক কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,15,000 টাকায়, হিরো হিরোইক কঠোর পরিশ্রমী ব্যবহারিক কমিউটারদের জন্য নিখুঁত যারা ভারতের নির্ভরযোগ্য কর্ম-দায়িত্ব যানবাহন, শক্তিশালী স্থায়িত্ব এবং ব্যতিক্রমী 80+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ব্যবহারিক কর্ম-দায়িত্ব কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Hero Heroic\n")
	return nil
}

func (s *HeroHeroicBatch23) GetName() string {
	return "HeroHeroicBatch23"
}
