package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BajajPulsar220FBatch23 struct {
	BaseSeeder
}

func NewBajajPulsar220FBatch23Seeder() *BajajPulsar220FBatch23 {
	return &BajajPulsar220FBatch23{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar 220F Batch23 Review"}}
}

func (s *BajajPulsar220FBatch23) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar 220F").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj pulsar 220f product not found")
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
<h1 class="review-main-title">Bajaj Pulsar 220F Review Bangladesh 2024 - India's Budget Muscle Performance Street Fighter</h1>
<p class="summary-text">The Bajaj Pulsar 220F is India's budget 220cc muscle performance street fighter representing accessible sport power. Priced around ৳4,65,000, it delivers 21.7 bhp exhilarating power, aggressive muscle styling, good handling, practical 45+ km/l efficiency, and excellent value for budget performance enthusiasts seeking muscle sport character.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar 220F Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>220cc Fuel-Injected:</strong> Budget muscle</li>
<li class="highlight-item"><strong>21.7 bhp Power:</strong> Exhilarating spirit</li>
<li class="highlight-item"><strong>Aggressive Styling:</strong> Muscle street fighter</li>
<li class="highlight-item"><strong>45+ km/l Efficiency:</strong> Practical economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar 220F Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Exhilarating 21.7 bhp:</strong> Performance budget</li>
<li class="pro-item"><strong>Aggressive Muscle Styling:</strong> Street fighter</li>
<li class="pro-item"><strong>Good Handling:</strong> Responsive sport</li>
<li class="pro-item"><strong>Bajaj Value:</strong> Affordable power</li>
<li class="pro-item"><strong>Practical Efficiency:</strong> 45+ km/l</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar 220F Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Build Quality:</strong> Budget materials</li>
<li class="con-item"><strong>Limited Comfort:</strong> Sport focused</li>
<li class="con-item"><strong>Vibration:</strong> Higher frequency</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar 220F Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳4,65,000 - Performance budget</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>220cc fuel-injected single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>21.7 bhp exhilarating</span></p>
<p class="value-item"><strong>Torque:</strong> <span>20.4 nm strong</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>155kg sport-tuned</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>45+ km/l practical</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar 220F Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 21.7 bhp exhilarating</li>
<li class="rating-item"><strong>Styling:</strong> <span>4.7</span> - Aggressive muscle</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.6</span> - Good responsive</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳4,65,000 affordable</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.6</span> - Pulsar proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar 220F?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳4,65,000, the Bajaj Pulsar 220F is perfect for budget performance enthusiasts seeking accessible muscle power, exhilarating 21.7 bhp sport character, aggressive styling, good handling, and budget-friendly muscle sport performance.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget muscle performance seekers</span></p>
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
<h1 class="review-main-title">বাজাজ পালসার 220এফ রিভিউ বাংলাদেশ 2024 - ভারতের বাজেট পেশী পারফরম্যান্স রাস্তা ফাইটার</h1>
<p class="summary-text">বাজাজ পালসার 220এফ ভারতের বাজেট 220cc পেশী পারফরম্যান্স রাস্তা ফাইটার যা অ্যাক্সেসযোগ্য খেলাধুলা শক্তি প্রতিনিধিত্ব করে। ৳4,65,000 টাকায় মূল্যায়িত, এটি 21.7 bhp উত্তেজনাপূর্ণ শক্তি এবং আক্রমণাত্মক পেশী স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ পালসার 220এফ মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>220cc জ্বালানী-ইনজেকশন:</strong> বাজেট পেশী</li>
<li class="highlight-item"><strong>21.7 bhp শক্তি:</strong> উত্তেজনাপূর্ণ মন</li>
<li class="highlight-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> পেশী রাস্তা ফাইটার</li>
<li class="highlight-item"><strong>45+ km/l দক্ষতা:</strong> ব্যবহারিক অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ পালসার 220এফ সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>উত্তেজনাপূর্ণ 21.7 bhp:</strong> পারফরম্যান্স বাজেট</li>
<li class="pro-item"><strong>আক্রমণাত্মক পেশী স্টাইলিং:</strong> রাস্তা ফাইটার</li>
<li class="pro-item"><strong>ভালো হ্যান্ডলিং:</strong> প্রতিক্রিয়াশীল খেলাধুলা</li>
<li class="pro-item"><strong>বাজাজ মূল্য:</strong> সাশ্রয়ী শক্তি</li>
<li class="pro-item"><strong>ব্যবহারিক দক্ষতা:</strong> 45+ km/l</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ পালসার 220এফ অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>বিল্ড গুণমান:</strong> বাজেট উপকরণ</li>
<li class="con-item"><strong>সীমিত আরাম:</strong> খেলাধুলা ফোকাসড</li>
<li class="con-item"><strong>কম্পন:</strong> উচ্চ ফ্রিকোয়েন্সি</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ পালসার 220এফ কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳4,65,000 টাকায়, বাজাজ পালসার 220এফ বাজেট পারফরম্যান্স উৎসাহীদের জন্য নিখুঁত যারা অ্যাক্সেসযোগ্য পেশী শক্তি, উত্তেজনাপূর্ণ 21.7 bhp খেলাধুলা চরিত্র এবং আক্রমণাত্মক স্টাইলিং চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট পেশী পারফরম্যান্স খোঁজার জন্য</span></p>
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

	fmt.Printf("   ✅ Created Bajaj Pulsar 220F\n")
	return nil
}

func (s *BajajPulsar220FBatch23) GetName() string {
	return "BajajPulsar220FBatch23"
}
