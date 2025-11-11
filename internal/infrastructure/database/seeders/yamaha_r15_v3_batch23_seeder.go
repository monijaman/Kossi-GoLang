package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type YamahaR15V3Batch23 struct {
	BaseSeeder
}

func NewYamahaR15V3Batch23Seeder() *YamahaR15V3Batch23 {
	return &YamahaR15V3Batch23{BaseSeeder: BaseSeeder{name: "Yamaha R15 V3 Batch23 Review"}}
}

func (s *YamahaR15V3Batch23) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha R15 V3").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha r15 v3 product not found")
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
<h1 class="review-main-title">Yamaha R15 V3 Review Bangladesh 2024 - India's Ultimate Sports Bike Thrill</h1>
<p class="summary-text">The Yamaha R15 V3 is India's ultimate 155cc sports bike thrill representing genuine performance engineering. Priced around ৳5,35,000, it delivers 16.8 bhp exhilarating power, extreme race styling, sharp handling, good 50+ km/l efficiency, and outstanding value for youth racers seeking extreme sport performance character daily.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha R15 V3 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>155cc Fuel-Injected:</strong> Sports performance</li>
<li class="highlight-item"><strong>16.8 bhp Power:</strong> Exhilarating thrill</li>
<li class="highlight-item"><strong>Extreme Race Design:</strong> Sport character</li>
<li class="highlight-item"><strong>50+ km/l Efficiency:</strong> Practical economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha R15 V3 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Exhilarating 16.8 bhp:</strong> Young racer thrill</li>
<li class="pro-item"><strong>Extreme Race Styling:</strong> Authentic sport DNA</li>
<li class="pro-item"><strong>Sharp Handling:</strong> Track-ready agile</li>
<li class="pro-item"><strong>Yamaha Reputation:</strong> Quality engineering</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 50+ km/l practical</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha R15 V3 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳5,35,000 upscale</li>
<li class="con-item"><strong>Limited Comfort:</strong> Race-focused</li>
<li class="con-item"><strong>Aggressive Character:</strong> Sport-only appeal</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha R15 V3 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳5,35,000 - Sport premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>155cc fuel-injected sport</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>16.8 bhp exhilarating</span></p>
<p class="value-item"><strong>Torque:</strong> <span>15 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>142kg sport-race</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50+ km/l practical</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha R15 V3 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.8</span> - 16.8 bhp exhilarating</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.8</span> - Sharp track-ready</li>
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Extreme race DNA</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.8</span> - Yamaha proven</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳5,35,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha R15 V3?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳5,35,000, the Yamaha R15 V3 is perfect for youth racers seeking India's ultimate sports bike, exhilarating 16.8 bhp thrill, extreme race styling, sharp track-ready handling, and genuine performance sport character.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For youth sport racers</span></p>
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
<h1 class="review-main-title">ইয়ামাহা আর15 ভি3 রিভিউ বাংলাদেশ 2024 - ভারতের চূড়ান্ত খেলাধুলা বাইক রোমাঞ্চ</h1>
<p class="summary-text">ইয়ামাহা আর15 ভি3 ভারতের চূড়ান্ত 155cc খেলাধুলা বাইক রোমাঞ্চ যা খাঁটি পারফরম্যান্স প্রকৌশল প্রতিনিধিত্ব করে। ৳5,35,000 টাকায় মূল্যায়িত, এটি 16.8 bhp উত্তেজনাপূর্ণ শক্তি এবং চরম রেস স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা আর15 ভি3 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>155cc জ্বালানী-ইনজেকশন:</strong> খেলাধুলা পারফরম্যান্স</li>
<li class="highlight-item"><strong>16.8 bhp শক্তি:</strong> উত্তেজনাপূর্ণ রোমাঞ্চ</li>
<li class="highlight-item"><strong>চরম রেস ডিজাইন:</strong> খেলাধুলা চরিত্র</li>
<li class="highlight-item"><strong>50+ km/l দক্ষতা:</strong> ব্যবহারিক অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা আর15 ভি3 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>উত্তেজনাপূর্ণ 16.8 bhp:</strong> যুব রেসার রোমাঞ্চ</li>
<li class="pro-item"><strong>চরম রেস স্টাইলিং:</strong> খাঁটি খেলাধুলা ডিএনএ</li>
<li class="pro-item"><strong>তীক্ষ্ণ হ্যান্ডলিং:</strong> ট্র্যাক-প্রস্তুত চটপটে</li>
<li class="pro-item"><strong>ইয়ামাহা খ্যাতি:</strong> গুণমান প্রকৌশল</li>
<li class="pro-item"><strong>ভালো দক্ষতা:</strong> 50+ km/l ব্যবহারিক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা আর15 ভি3 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳5,35,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>সীমিত আরাম:</strong> রেস-ফোকাসড</li>
<li class="con-item"><strong>আক্রমণাত্মক চরিত্র:</strong> খেলাধুলা-শুধুমাত্র আবেদন</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা আর15 ভি3 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳5,35,000 টাকায়, ইয়ামাহা আর15 ভি3 যুব রেসারদের জন্য নিখুঁত যারা ভারতের চূড়ান্ত খেলাধুলা বাইক, উত্তেজনাপূর্ণ 16.8 bhp রোমাঞ্চ এবং চরম রেস চরিত্র চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- যুব খেলাধুলা রেসারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Yamaha R15 V3\n")
	return nil
}

func (s *YamahaR15V3Batch23) GetName() string {
	return "YamahaR15V3Batch23"
}
