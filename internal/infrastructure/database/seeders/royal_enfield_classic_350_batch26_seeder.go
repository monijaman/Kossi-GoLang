package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type RoyalEnfieldClassic350Batch26 struct {
	BaseSeeder
}

func NewRoyalEnfieldClassic350Batch26Seeder() *RoyalEnfieldClassic350Batch26 {
	return &RoyalEnfieldClassic350Batch26{BaseSeeder: BaseSeeder{name: "Royal Enfield Classic 350 Batch26 Review"}}
}

func (s *RoyalEnfieldClassic350Batch26) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Royal Enfield Classic 350").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("royal enfield classic 350 product not found")
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
<h1 class="review-main-title">Royal Enfield Classic 350 Review Bangladesh 2024 - Timeless Cruiser Icon</h1>
<p class="summary-text">The Royal Enfield Classic 350 is a timeless cruiser icon combining vintage charm and modern reliability. Priced around ৳4,65,000, it delivers 20.2 bhp power, retro styling, comfortable cruising, distinctive thump, and unforgettable heritage riding experience.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Royal Enfield Classic 350 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>350cc Cruiser:</strong> Iconic timeless</li>
<li class="highlight-item"><strong>20.2 bhp Power:</strong> Torquey capable</li>
<li class="highlight-item"><strong>Retro Styling:</strong> Vintage charming</li>
<li class="highlight-item"><strong>Distinctive Thump:</strong> Heritage soundtrack</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Royal Enfield Classic 350 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Timeless Design:</strong> Retro iconic heritage</li>
<li class="pro-item"><strong>Distinctive Character:</strong> Unique thump soul</li>
<li class="pro-item"><strong>Comfortable Cruising:</strong> Relaxed riding</li>
<li class="pro-item"><strong>Strong Torque:</strong> 20.2 bhp capable</li>
<li class="pro-item"><strong>Brand Heritage:</strong> Royal Enfield legacy</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Royal Enfield Classic 350 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Heavy Weight:</strong> 195kg substantial</li>
<li class="con-item"><strong>Fuel Consumption:</strong> 35+ km/l moderate</li>
<li class="con-item"><strong>Higher Price:</strong> Premium cruiser</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Royal Enfield Classic 350 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳4,65,000 - Cruiser icon</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>350cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>20.2 bhp torquey</span></p>
<p class="value-item"><strong>Torque:</strong> <span>27 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>195kg substantial</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>35+ km/l moderate</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Royal Enfield Classic 350 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.9</span> - Timeless retro</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Cruiser relaxed</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 20.2 bhp torquey</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.5</span> - 35+ km/l moderate</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Royal heritage</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Royal Enfield Classic 350?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳4,65,000, the Royal Enfield Classic 350 is perfect for heritage enthusiasts seeking timeless retro styling, distinctive thump character, comfortable cruising, and unforgettable Royal Enfield legacy.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For heritage enthusiasts</span></p>
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
<h1 class="review-main-title">রয়্যাল এনফিল্ড ক্ল্যাসিক 350 রিভিউ বাংলাদেশ 2024 - চিরন্তন ক্রুজার আইকন</h1>
<p class="summary-text">রয়্যাল এনফিল্ড ক্ল্যাসিক 350 একটি চিরন্তন ক্রুজার আইকন যা ভিনটেজ মোহ এবং আধুনিক নির্ভরযোগ্যতা একত্রিত করে। ৳4,65,000 টাকায় মূল্যায়িত, এটি 20.2 bhp শক্তি এবং রেট্রো স্টাইলিং বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">রয়্যাল এনফিল্ড ক্ল্যাসিক 350 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>350cc ক্রুজার:</strong> আইকনিক চিরন্তন</li>
<li class="highlight-item"><strong>20.2 bhp শক্তি:</strong> টর্কি সক্ষম</li>
<li class="highlight-item"><strong>রেট্রো স্টাইলিং:</strong> ভিনটেজ মোহনীয়</li>
<li class="highlight-item"><strong>স্বতন্ত্র থাম্প:</strong> ঐতিহ্য সাউন্ডট্র্যাক</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">রয়্যাল এনফিল্ড ক্ল্যাসিক 350 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>চিরন্তন ডিজাইন:</strong> রেট্রো আইকনিক ঐতিহ্য</li>
<li class="pro-item"><strong>স্বতন্ত্র চরিত্র:</strong> অনন্য থাম্প আত্মা</li>
<li class="pro-item"><strong>আরামদায়ক ক্রুজিং:</strong> শিথিল রাইডিং</li>
<li class="pro-item"><strong>শক্তিশালী টর্ক:</strong> 20.2 bhp সক্ষম</li>
<li class="pro-item"><strong>ব্র্যান্ড ঐতিহ্য:</strong> রয়্যাল এনফিল্ড উত্তরাধিকার</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">রয়্যাল এনফিল্ড ক্ল্যাসিক 350 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>ভারী ওজন:</strong> 195kg উল্লেখযোগ্য</li>
<li class="con-item"><strong>জ্বালানি খরচ:</strong> 35+ km/l মধ্যম</li>
<li class="con-item"><strong>উচ্চতর মূল্য:</strong> প্রিমিয়াম ক্রুজার</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: রয়্যাল এনফিল্ড ক্ল্যাসিক 350 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳4,65,000 টাকায়, রয়্যাল এনফিল্ড ক্ল্যাসিক 350 ঐতিহ্য উৎসাহীদের জন্য নিখুঁত যারা চিরন্তন রেট্রো স্টাইলিং এবং স্বতন্ত্র থাম্প চরিত্র চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ঐতিহ্য উৎসাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Royal Enfield Classic 350\n")
	return nil
}

func (s *RoyalEnfieldClassic350Batch26) GetName() string {
	return "RoyalEnfieldClassic350Batch26"
}
