package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type HeroHFDeluxeBatch26 struct {
	BaseSeeder
}

func NewHeroHFDeluxeBatch26Seeder() *HeroHFDeluxeBatch26 {
	return &HeroHFDeluxeBatch26{BaseSeeder: BaseSeeder{name: "Hero HF Deluxe Batch26 Review"}}
}

func (s *HeroHFDeluxeBatch26) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Hero HF Deluxe").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("hero hf deluxe product not found")
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
<h1 class="review-main-title">Hero HF Deluxe Review Bangladesh 2024 - Budget Champion</h1>
<p class="summary-text">The Hero HF Deluxe is Bangladesh's ultimate budget commuter combining affordability and reliability. Priced around ৳1,55,000, it delivers 8.2 bhp power, simple maintenance, comfortable seating, exceptional 80+ km/l efficiency, and outstanding value for budget-conscious riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Hero HF Deluxe Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc Budget:</strong> Champion affordability</li>
<li class="highlight-item"><strong>8.2 bhp Power:</strong> Reliable efficient</li>
<li class="highlight-item"><strong>Simple Maintenance:</strong> Cost-effective service</li>
<li class="highlight-item"><strong>80+ km/l Efficiency:</strong> Exceptional economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Hero HF Deluxe Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Ultra Affordable:</strong> Budget champion pricing</li>
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 80+ km/l outstanding</li>
<li class="pro-item"><strong>Simple Design:</strong> Easy maintenance</li>
<li class="pro-item"><strong>Reliable Performance:</strong> 8.2 bhp dependable</li>
<li class="pro-item"><strong>Wide Service:</strong> Extensive network</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Hero HF Deluxe Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Basic Features:</strong> Minimal amenities</li>
<li class="con-item"><strong>Limited Power:</strong> 8.2 bhp modest</li>
<li class="con-item"><strong>Traditional Design:</strong> Basic styling</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Hero HF Deluxe Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,55,000 - Budget champion</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>100cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>8.2 bhp reliable</span></p>
<p class="value-item"><strong>Torque:</strong> <span>8.3 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>112kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>80+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Hero HF Deluxe Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.5</span> - Traditional basic</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.6</span> - Simple comfortable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.5</span> - 8.2 bhp reliable</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.9</span> - 80+ km/l exceptional</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - Hero proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Hero HF Deluxe?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,55,000, the Hero HF Deluxe is perfect for budget-conscious commuters seeking ultra-affordable pricing, exceptional 80+ km/l efficiency, simple maintenance, and reliable 8.2 bhp performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget commuters</span></p>
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
<h1 class="review-main-title">হিরো HF ডিলাক্স রিভিউ বাংলাদেশ 2024 - বাজেট চ্যাম্পিয়ন</h1>
<p class="summary-text">হিরো HF ডিলাক্স হল বাংলাদেশের চূড়ান্ত বাজেট কমিউটার যা সামর্থ্য এবং নির্ভরযোগ্যতা একত্রিত করে। ৳1,55,000 টাকায় মূল্যায়িত, এটি 8.2 bhp শক্তি এবং সহজ রক্ষণাবেক্ষণ বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হিরো HF ডিলাক্স মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc বাজেট:</strong> চ্যাম্পিয়ন সামর্থ্য</li>
<li class="highlight-item"><strong>8.2 bhp শক্তি:</strong> নির্ভরযোগ্য দক্ষ</li>
<li class="highlight-item"><strong>সহজ রক্ষণাবেক্ষণ:</strong> সাশ্রয়ী সেবা</li>
<li class="highlight-item"><strong>80+ km/l দক্ষতা:</strong> ব্যতিক্রমী অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হিরো HF ডিলাক্স সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>অতি সাশ্রয়ী:</strong> বাজেট চ্যাম্পিয়ন মূল্য</li>
<li class="pro-item"><strong>ব্যতিক্রমী দক্ষতা:</strong> 80+ km/l অসাধারণ</li>
<li class="pro-item"><strong>সহজ ডিজাইন:</strong> সহজ রক্ষণাবেক্ষণ</li>
<li class="pro-item"><strong>নির্ভরযোগ্য কর্মক্ষমতা:</strong> 8.2 bhp নির্ভরযোগ্য</li>
<li class="pro-item"><strong>বিস্তৃত সেবা:</strong> বিস্তৃত নেটওয়ার্ক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হিরো HF ডিলাক্স অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মৌলিক বৈশিষ্ট্য:</strong> ন্যূনতম সুবিধাবিদ</li>
<li class="con-item"><strong>সীমিত শক্তি:</strong> 8.2 bhp মধ্যম</li>
<li class="con-item"><strong>ঐতিহ্যবাহী ডিজাইন:</strong> মৌলিক স্টাইলিং</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হিরো HF ডিলাক্স কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,55,000 টাকায়, হিরো HF ডিলাক্স বাজেট-সচেতন যাতায়াতকারীদের জন্য নিখুঁত যারা অতি সাশ্রয়ী মূল্য এবং ব্যতিক্রমী 80+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট যাতায়াতকারীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Hero HF Deluxe\n")
	return nil
}

func (s *HeroHFDeluxeBatch26) GetName() string {
	return "HeroHFDeluxeBatch26"
}
