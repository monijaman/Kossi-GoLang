package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type BMWRnineT900ReviewBatch20 struct {
	BaseSeeder
}

func NewBMWRnineT900ReviewBatch20Seeder() *BMWRnineT900ReviewBatch20 {
	return &BMWRnineT900ReviewBatch20{BaseSeeder: BaseSeeder{name: "BMW RnineT 900 Batch20 Review"}}
}

func (s *BMWRnineT900ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "BMW RnineT 900").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bmw rninet 900 product not found")
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
<h1 class="review-main-title">BMW RnineT 900 Review Bangladesh 2024 - Modern Retro-Classic German Heritage</h1>
<p class="summary-text">The BMW RnineT 900 is a 900cc liquid-cooled modern retro-classic representing iconic horizontal-twin German engineering heritage. Priced around ৳21,00,000, it delivers 110 bhp smooth power, retro-modern styling, superior handling, premium German build quality, and classic soul fused with modern technology.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">BMW RnineT 900 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>900cc Horizontal-Twin:</strong> Iconic German engine</li>
<li class="highlight-item"><strong>110 bhp Smooth:</strong> Twin power delivery</li>
<li class="highlight-item"><strong>Retro-Modern Design:</strong> Classic soul fusion</li>
<li class="highlight-item"><strong>Superior Quality:</strong> Premium German build</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">BMW RnineT 900 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Iconic Horizontal-Twin:</strong> Legendary 900cc engine</li>
<li class="pro-item"><strong>Retro-Modern Design:</strong> Classic beauty fusion</li>
<li class="pro-item"><strong>Superior Handling:</strong> Precision German engineering</li>
<li class="pro-item"><strong>Premium Build Quality:</strong> Exceptional craftsmanship</li>
<li class="pro-item"><strong>Modern Technology:</strong> Classic heritage meets innovation</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">BMW RnineT 900 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳21,00,000 German luxury</li>
<li class="con-item"><strong>Fuel Economy:</strong> 16-19 km/l average</li>
<li class="con-item"><strong>Moderate Weight:</strong> 221kg premium build</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">BMW RnineT 900 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳21,00,000 - Premium retro classic</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>900cc liquid-cooled horizontal-twin</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>110 bhp smooth</span></p>
<p class="value-item"><strong>Torque:</strong> <span>92 nm relentless</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>221kg premium</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>16-19 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">BMW RnineT 900 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.9</span> - Iconic retro-modern</li>
<li class="rating-item"><strong>Engine:</strong> <span>4.9</span> - Legendary twin</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.8</span> - Superior German</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.9</span> - Premium craftsmanship</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳21,00,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy BMW RnineT 900?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳21,00,000, the RnineT 900 is the premium choice for riders seeking modern retro-classic perfection, iconic horizontal-twin engine heritage, superior German engineering, and timeless design that bridges classic soul with contemporary technology.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium classic enthusiasts</span></p>
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
<h1 class="review-main-title">BMW RnineT 900 রিভিউ বাংলাদেশ 2024 - আধুনিক রেট্রো-ক্লাসিক জার্মান ঐতিহ্য</h1>
<p class="summary-text">BMW RnineT 900 একটি 900cc লিকুইড-কুল্ড আধুনিক রেট্রো-ক্লাসিক যা আইকনিক হরিজন্টাল-টুইন জার্মান ইঞ্জিনিয়ারিং ঐতিহ্য প্রতিনিধিত্ব করে। ৳21,00,000 টাকায় মূল্যায়িত, এটি 110 bhp মসৃণ শক্তি এবং প্রিমিয়াম জার্মান বিল্ড কোয়ালিটি প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">BMW RnineT 900 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>900cc হরিজন্টাল-টুইন:</strong> আইকনিক জার্মান ইঞ্জিন</li>
<li class="highlight-item"><strong>110 bhp মসৃণ:</strong> টুইন শক্তি সরবরাহ</li>
<li class="highlight-item"><strong>রেট্রো-আধুনিক ডিজাইন:</strong> ক্লাসিক আত্মা সংমিশ্রণ</li>
<li class="highlight-item"><strong>উচ্চতর গুণমান:</strong> প্রিমিয়াম জার্মান বিল্ড</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">BMW RnineT 900 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আইকনিক হরিজন্টাল-টুইন:</strong> কিংবদন্তী 900cc ইঞ্জিন</li>
<li class="pro-item"><strong>রেট্রো-আধুনিক ডিজাইন:</strong> ক্লাসিক সৌন্দর্য সংমিশ্রণ</li>
<li class="pro-item"><strong>উচ্চতর হ্যান্ডলিং:</strong> নির্ভুল জার্মান ইঞ্জিনিয়ারিং</li>
<li class="pro-item"><strong>প্রিমিয়াম বিল্ড কোয়ালিটি:</strong> ব্যতিক্রমী কারুশিল্প</li>
<li class="pro-item"><strong>আধুনিক প্রযুক্তি:</strong> ক্লাসিক ঐতিহ্য আধুনিকতা মিলন</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">BMW RnineT 900 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳21,00,000 জার্মান বিলাসিতা</li>
<li class="con-item"><strong>জ্বালানি অর্থনীতি:</strong> 16-19 km/l গড়</li>
<li class="con-item"><strong>মধ্যম ওজন:</strong> 221kg প্রিমিয়াম বিল্ড</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: BMW RnineT 900 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳21,00,000 টাকায়, RnineT 900 প্রিমিয়াম পছন্দ যারা আধুনিক রেট্রো-ক্লাসিক নিখুঁততা, আইকনিক হরিজন্টাল-টুইন ইঞ্জিন ঐতিহ্য এবং শ্রেষ্ঠ জার্মান ইঞ্জিনিয়ারিং চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম ক্লাসিক উত্সাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created BMW RnineT 900\n")
	return nil
}

func (s *BMWRnineT900ReviewBatch20) GetName() string {
	return "BMWRnineT900ReviewBatch20"
}
