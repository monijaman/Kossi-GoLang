package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BajajAvenger220Batch24 struct {
	BaseSeeder
}

func NewBajajAvenger220Batch24Seeder() *BajajAvenger220Batch24 {
	return &BajajAvenger220Batch24{BaseSeeder: BaseSeeder{name: "Bajaj Avenger 220 Batch24 Review"}}
}

func (s *BajajAvenger220Batch24) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Avenger 220").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj avenger 220 product not found")
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
<h1 class="review-main-title">Bajaj Avenger 220 Review Bangladesh 2024 - Cruiser Muscle Power</h1>
<p class="summary-text">The Bajaj Avenger 220 is a cruiser-styled 220cc motorcycle delivering impressive muscle and power. Priced around ৳4,15,000, it features 19.3 bhp power, cruiser design language, comfortable seating, excellent 58+ km/l efficiency, and outstanding appeal for riders seeking cruiser character.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Avenger 220 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>220cc Cruiser:</strong> Muscle styled</li>
<li class="highlight-item"><strong>19.3 bhp Power:</strong> Responsive thrilling</li>
<li class="highlight-item"><strong>Cruiser Design:</strong> Comfortable poised</li>
<li class="highlight-item"><strong>58+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Avenger 220 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Cruiser Design:</strong> Muscle styled appeal</li>
<li class="pro-item"><strong>Responsive Power:</strong> 19.3 bhp capable</li>
<li class="pro-item"><strong>Comfortable Seating:</strong> Relaxed riding</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 58+ km/l great</li>
<li class="pro-item"><strong>Bajaj Reliability:</strong> Proven brand</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Avenger 220 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Heavier Weight:</strong> 186kg substantial</li>
<li class="con-item"><strong>Limited Agility:</strong> Cruiser positioned</li>
<li class="con-item"><strong>Basic Features:</strong> Cruiser focused</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Avenger 220 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳4,15,000 - Cruiser accessible</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>220cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>19.3 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>17.5 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>186kg cruiser</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>58+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Avenger 220 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Cruiser muscle</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Very comfortable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 19.3 bhp responsive</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 58+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Bajaj proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Avenger 220?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳4,15,000, the Bajaj Avenger 220 is perfect for riders seeking cruiser character, 19.3 bhp responsive power, comfortable relaxed riding, and excellent 58+ km/l efficiency in a muscle-styled package.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For cruiser enthusiasts</span></p>
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
<h1 class="review-main-title">বাজাজ অ্যাভেঞ্জার 220 রিভিউ বাংলাদেশ 2024 - ক্রুজার মাসল শক্তি</h1>
<p class="summary-text">বাজাজ অ্যাভেঞ্জার 220 একটি ক্রুজার-স্টাইল 220cc মোটরসাইকেল যা চিত্তাকর্ষক মাসল এবং শক্তি প্রদান করে। ৳4,15,000 টাকায় মূল্যায়িত, এটি 19.3 bhp শক্তি এবং ক্রুজার ডিজাইন ভাষা বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ অ্যাভেঞ্জার 220 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>220cc ক্রুজার:</strong> মাসল স্টাইল</li>
<li class="highlight-item"><strong>19.3 bhp শক্তি:</strong> প্রতিক্রিয়াশীল রোমাঞ্চকর</li>
<li class="highlight-item"><strong>ক্রুজার ডিজাইন:</strong> আরামদায়ক অবস্থিত</li>
<li class="highlight-item"><strong>58+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ অ্যাভেঞ্জার 220 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>ক্রুজার ডিজাইন:</strong> মাসল স্টাইল আবেদন</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল শক্তি:</strong> 19.3 bhp সক্ষম</li>
<li class="pro-item"><strong>আরামদায়ক আসন:</strong> শিথিল রাইডিং</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 58+ km/l দুর্দান্ত</li>
<li class="pro-item"><strong>বাজাজ নির্ভরযোগ্যতা:</strong> প্রমাণিত ব্র্যান্ড</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ অ্যাভেঞ্জার 220 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>ভারী ওজন:</strong> 186kg উল্লেখযোগ্য</li>
<li class="con-item"><strong>সীমিত চটপটে:</strong> ক্রুজার অবস্থান</li>
<li class="con-item"><strong>মৌলিক বৈশিষ্ট্য:</strong> ক্রুজার ফোকাসড</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ অ্যাভেঞ্জার 220 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳4,15,000 টাকায়, বাজাজ অ্যাভেঞ্জার 220 চালকদের জন্য নিখুঁত যারা ক্রুজার চরিত্র এবং প্রতিক্রিয়াশীল 19.3 bhp শক্তি চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ক্রুজার উত্সাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Bajaj Avenger 220\n")
	return nil
}

func (s *BajajAvenger220Batch24) GetName() string {
	return "BajajAvenger220Batch24"
}
