package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type KTMDuke125Batch23 struct {
	BaseSeeder
}

func NewKTMDuke125Batch23Seeder() *KTMDuke125Batch23 {
	return &KTMDuke125Batch23{BaseSeeder: BaseSeeder{name: "KTM Duke 125 Batch23 Review"}}
}

func (s *KTMDuke125Batch23) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "KTM Duke 125").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("ktm duke 125 product not found")
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
<h1 class="review-main-title">KTM Duke 125 Review Bangladesh 2024 - Austria's Youthful Performance Character</h1>
<p class="summary-text">The KTM Duke 125 is Austria's youthful 125cc performance character representing accessible sport fun. Priced around ৳4,15,000, it delivers 14.5 bhp exhilarating power, sharp aggressive styling, responsive handling, good 50+ km/l efficiency, and excellent value for youth riders seeking authentic sport fun character.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">KTM Duke 125 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Air-Cooled:</strong> Youth performance</li>
<li class="highlight-item"><strong>14.5 bhp Power:</strong> Exhilarating fun</li>
<li class="highlight-item"><strong>Sharp Styling:</strong> Aggressive character</li>
<li class="highlight-item"><strong>50+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">KTM Duke 125 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Exhilarating 14.5 bhp:</strong> Youth fun thrill</li>
<li class="pro-item"><strong>Sharp Aggressive Styling:</strong> Street character</li>
<li class="pro-item"><strong>Responsive Handling:</strong> Fun agile</li>
<li class="pro-item"><strong>KTM Quality:</strong> European engineering</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 50+ km/l practical</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">KTM Duke 125 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳4,15,000 upscale</li>
<li class="con-item"><strong>Limited Service:</strong> KTM dealers</li>
<li class="con-item"><strong>Aggressive Posture:</strong> Sport-only focus</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">KTM Duke 125 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳4,15,000 - Sport premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>14.5 bhp exhilarating</span></p>
<p class="value-item"><strong>Torque:</strong> <span>12 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>148kg sport-tuned</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">KTM Duke 125 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 14.5 bhp exhilarating</li>
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Sharp aggressive</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.7</span> - Responsive fun</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - KTM European</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳4,15,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy KTM Duke 125?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳4,15,000, the KTM Duke 125 is perfect for youth riders seeking Austria's accessible sport fun, exhilarating 14.5 bhp power, sharp aggressive styling, responsive handling, and authentic sport character.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For youth sport fun seekers</span></p>
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
<h1 class="review-main-title">কেটিএম ডিউক 125 রিভিউ বাংলাদেশ 2024 - অস্ট্রিয়ার যুববান পারফরম্যান্স চরিত্র</h1>
<p class="summary-text">কেটিএম ডিউক 125 অস্ট্রিয়ার যুববান 125cc পারফরম্যান্স চরিত্র যা অ্যাক্সেসযোগ্য খেলাধুলা মজা প্রতিনিধিত্ব করে। ৳4,15,000 টাকায় মূল্যায়িত, এটি 14.5 bhp উত্তেজনাপূর্ণ শক্তি এবং তীক্ষ্ণ আক্রমণাত্মক স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">কেটিএম ডিউক 125 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc এয়ার-কুল্ড:</strong> যুব পারফরম্যান্স</li>
<li class="highlight-item"><strong>14.5 bhp শক্তি:</strong> উত্তেজনাপূর্ণ মজা</li>
<li class="highlight-item"><strong>তীক্ষ্ণ স্টাইলিং:</strong> আক্রমণাত্মক চরিত্র</li>
<li class="highlight-item"><strong>50+ km/l দক্ষতা:</strong> ভালো অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">কেটিএম ডিউক 125 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>উত্তেজনাপূর্ণ 14.5 bhp:</strong> যুব মজা রোমাঞ্চ</li>
<li class="pro-item"><strong>তীক্ষ্ণ আক্রমণাত্মক স্টাইলিং:</strong> রাস্তা চরিত্র</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল হ্যান্ডলিং:</strong> মজা চটপটে</li>
<li class="pro-item"><strong>কেটিএম গুণমান:</strong> ইউরোপীয় প্রকৌশল</li>
<li class="pro-item"><strong>ভালো দক্ষতা:</strong> 50+ km/l ব্যবহারিক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">কেটিএম ডিউক 125 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳4,15,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>সীমিত সেবা:</strong> কেটিএম ডিলার</li>
<li class="con-item"><strong>আক্রমণাত্মক ভঙ্গি:</strong> খেলাধুলা-শুধুমাত্র ফোকাস</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: কেটিএম ডিউক 125 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳4,15,000 টাকায়, কেটিএম ডিউক 125 যুব চালকদের জন্য নিখুঁত যারা অস্ট্রিয়ার অ্যাক্সেসযোগ্য খেলাধুলা মজা, উত্তেজনাপূর্ণ 14.5 bhp শক্তি এবং খাঁটি খেলাধুলা চরিত্র চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- যুব খেলাধুলা মজা খোঁজার জন্য</span></p>
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

	fmt.Printf("   ✅ Created KTM Duke 125\n")
	return nil
}

func (s *KTMDuke125Batch23) GetName() string {
	return "KTMDuke125Batch23"
}
