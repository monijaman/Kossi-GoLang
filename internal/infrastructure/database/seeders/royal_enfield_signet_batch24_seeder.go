package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type RoyalEnfieldSignetBatch24 struct {
	BaseSeeder
}

func NewRoyalEnfieldSignetBatch24Seeder() *RoyalEnfieldSignetBatch24 {
	return &RoyalEnfieldSignetBatch24{BaseSeeder: BaseSeeder{name: "Royal Enfield Signet Batch24 Review"}}
}

func (s *RoyalEnfieldSignetBatch24) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Royal Enfield Signet").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("royal enfield signet product not found")
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
<h1 class="review-main-title">Royal Enfield Signet Review Bangladesh 2024 - Classic Heritage Motorcycle</h1>
<p class="summary-text">The Royal Enfield Signet is a classic 350cc heritage motorcycle combining simplicity and character. Priced around ৳6,25,000, it delivers 20 bhp power, classic design, excellent build, impressive 60+ km/l efficiency, and outstanding appeal for riders seeking authentic motorcycle heritage.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Royal Enfield Signet Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>350cc Classic:</strong> Heritage design</li>
<li class="highlight-item"><strong>20 bhp Power:</strong> Capable responsive</li>
<li class="highlight-item"><strong>Classic Character:</strong> Timeless appeal</li>
<li class="highlight-item"><strong>60+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Royal Enfield Signet Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Classic Design:</strong> Heritage character</li>
<li class="pro-item"><strong>Capable Power:</strong> 20 bhp responsive</li>
<li class="pro-item"><strong>Build Quality:</strong> Excellent construction</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 60+ km/l great</li>
<li class="pro-item"><strong>Heritage Appeal:</strong> Authentic character</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Royal Enfield Signet Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Higher Price:</strong> ৳6,25,000 upscale</li>
<li class="con-item"><strong>Heavier Weight:</strong> 202kg substantial</li>
<li class="con-item"><strong>Limited Modern Features:</strong> Classic approach</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Royal Enfield Signet Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳6,25,000 - Heritage classic</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>350cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>20 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>28 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>202kg classic</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>60+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Royal Enfield Signet Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Classic heritage</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Upright capable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 20 bhp smooth</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 60+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Heritage proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Royal Enfield Signet?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳6,25,000, the Royal Enfield Signet is perfect for riders seeking authentic classic heritage, 20 bhp capable power, excellent 60+ km/l efficiency, and timeless motorcycle character.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For classic heritage seekers</span></p>
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
<h1 class="review-main-title">রয়াল এনফিল্ড সিগনেট রিভিউ বাংলাদেশ 2024 - ক্লাসিক ঐতিহ্য মোটরসাইকেল</h1>
<p class="summary-text">রয়াল এনফিল্ড সিগনেট একটি ক্লাসিক 350cc ঐতিহ্য মোটরসাইকেল যা সরলতা এবং চরিত্র একত্রিত করে। ৳6,25,000 টাকায় মূল্যায়িত, এটি 20 bhp শক্তি এবং ক্লাসিক ডিজাইন প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">রয়াল এনফিল্ড সিগনেট মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>350cc ক্লাসিক:</strong> ঐতিহ্য ডিজাইন</li>
<li class="highlight-item"><strong>20 bhp শক্তি:</strong> সক্ষম প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>ক্লাসিক চরিত্র:</strong> চিরন্তন আবেদন</li>
<li class="highlight-item"><strong>60+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">রয়াল এনফিল্ড সিগনেট সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>ক্লাসিক ডিজাইন:</strong> ঐতিহ্য চরিত্র</li>
<li class="pro-item"><strong>সক্ষম শক্তি:</strong> 20 bhp প্রতিক্রিয়াশীল</li>
<li class="pro-item"><strong>বিল্ড গুণমান:</strong> চমৎকার নির্মাণ</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 60+ km/l দুর্দান্ত</li>
<li class="pro-item"><strong>ঐতিহ্য আবেদন:</strong> খাঁটি চরিত্র</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">রয়াল এনফিল্ড সিগনেট অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>উচ্চতর মূল্য:</strong> ৳6,25,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>ভারী ওজন:</strong> 202kg উল্লেখযোগ্য</li>
<li class="con-item"><strong>সীমিত আধুনিক বৈশিষ্ট্য:</strong> ক্লাসিক পদ্ধতি</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: রয়াল এনফিল্ড সিগনেট কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳6,25,000 টাকায়, রয়াল এনফিল্ড সিগনেট চালকদের জন্য নিখুঁত যারা খাঁটি ক্লাসিক ঐতিহ্য, 20 bhp সক্ষম শক্তি এবং চমৎকার 60+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ক্লাসিক ঐতিহ্য খোঁজার জন্য</span></p>
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

	fmt.Printf("   ✅ Created Royal Enfield Signet\n")
	return nil
}

func (s *RoyalEnfieldSignetBatch24) GetName() string {
	return "RoyalEnfieldSignetBatch24"
}
