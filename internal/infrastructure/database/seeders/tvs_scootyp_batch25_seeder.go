package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type TVSScootypBatch25 struct {
	BaseSeeder
}

func NewTVSScootypBatch25Seeder() *TVSScootypBatch25 {
	return &TVSScootypBatch25{BaseSeeder: BaseSeeder{name: "TVS Scootyp Batch25 Review"}}
}

func (s *TVSScootypBatch25) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Scootyp").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tvs scootyp product not found")
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
<h1 class="review-main-title">TVS Scootyp Review Bangladesh 2024 - Budget-Friendly Urban Scooter</h1>
<p class="summary-text">The TVS Scootyp is a budget-friendly 125cc urban scooter offering practicality and affordability. Priced around ৳2,85,000, it delivers 9 bhp power, spacious design, practical features, excellent 65+ km/l efficiency, and outstanding value for budget urban riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Scootyp Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Budget:</strong> Urban practical</li>
<li class="highlight-item"><strong>9 bhp Power:</strong> Responsive efficient</li>
<li class="highlight-item"><strong>Spacious Design:</strong> Practical ready</li>
<li class="highlight-item"><strong>65+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Scootyp Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Budget Price:</strong> ৳2,85,000 affordable</li>
<li class="pro-item"><strong>Spacious Design:</strong> Practical urban</li>
<li class="pro-item"><strong>Responsive Power:</strong> 9 bhp efficient</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 65+ km/l great</li>
<li class="pro-item"><strong>TVS Service:</strong> Wide network</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Scootyp Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Limited Power:</strong> 9 bhp modest</li>
<li class="con-item"><strong>Basic Features:</strong> Minimal styling</li>
<li class="con-item"><strong>Budget Build:</strong> Cost optimized</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Scootyp Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳2,85,000 - Budget urban</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc scooter single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>9 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>9.8 nm efficient</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>128kg practical</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>65+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Scootyp Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.6</span> - Practical spacious</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Urban comfortable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.6</span> - 9 bhp efficient</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 65+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - TVS proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Scootyp?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳2,85,000, the TVS Scootyp is perfect for budget urban riders seeking practical spacious design, responsive 9 bhp power, and excellent 65+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget urban riders</span></p>
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
<h1 class="review-main-title">টিভিএস স্কুটিপ রিভিউ বাংলাদেশ 2024 - বাজেট-বান্ধব শহুরে স্কুটার</h1>
<p class="summary-text">টিভিএস স্কুটিপ একটি বাজেট-বান্ধব 125cc শহুরে স্কুটার যা ব্যবহারিকতা এবং সামর্থ্য প্রদান করে। ৳2,85,000 টাকায় মূল্যায়িত, এটি 9 bhp শক্তি এবং প্রশস্ত ডিজাইন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">টিভিএস স্কুটিপ মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc বাজেট:</strong> শহুরে ব্যবহারিক</li>
<li class="highlight-item"><strong>9 bhp শক্তি:</strong> প্রতিক্রিয়াশীল দক্ষ</li>
<li class="highlight-item"><strong>প্রশস্ত ডিজাইন:</strong> ব্যবহারিক প্রস্তুত</li>
<li class="highlight-item"><strong>65+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">টিভিএস স্কুটিপ সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>বাজেট মূল্য:</strong> ৳2,85,000 সাশ্রয়ী</li>
<li class="pro-item"><strong>প্রশস্ত ডিজাইন:</strong> ব্যবহারিক শহুরে</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল শক্তি:</strong> 9 bhp দক্ষ</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 65+ km/l দুর্দান্ত</li>
<li class="pro-item"><strong>টিভিএস সেবা:</strong> বিস্তৃত নেটওয়ার্ক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">টিভিএস স্কুটিপ অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>সীমিত শক্তি:</strong> 9 bhp মামুলি</li>
<li class="con-item"><strong>মৌলিক বৈশিষ্ট্য:</strong> ন্যূনতম স্টাইলিং</li>
<li class="con-item"><strong>বাজেট বিল্ড:</strong> খরচ অপ্টিমাইজড</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: টিভিএস স্কুটিপ কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳2,85,000 টাকায়, টিভিএস স্কুটিপ বাজেট শহুরে চালকদের জন্য নিখুঁত যারা ব্যবহারিক প্রশস্ত ডিজাইন এবং চমৎকার 65+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট শহুরে চালকদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created TVS Scootyp\n")
	return nil
}

func (s *TVSScootypBatch25) GetName() string {
	return "TVSScootypBatch25"
}
