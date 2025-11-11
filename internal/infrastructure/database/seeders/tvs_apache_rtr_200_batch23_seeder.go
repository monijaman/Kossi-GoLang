package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type TVSApacheRTR200Batch23 struct {
	BaseSeeder
}

func NewTVSApacheRTR200Batch23Seeder() *TVSApacheRTR200Batch23 {
	return &TVSApacheRTR200Batch23{BaseSeeder: BaseSeeder{name: "TVS Apache RTR 200 Batch23 Review"}}
}

func (s *TVSApacheRTR200Batch23) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Apache RTR 200").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tvs apache rtr 200 product not found")
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
<h1 class="review-main-title">TVS Apache RTR 200 Review Bangladesh 2024 - India's Thrilling Performance Street Racer</h1>
<p class="summary-text">The TVS Apache RTR 200 is India's thrilling 200cc performance street racer representing accessible sport power. Priced around ৳4,95,000, it delivers 20.2 bhp exhilarating power, aggressive sport styling, sharp handling, good 48+ km/l efficiency, and excellent value for youth racers seeking spirited street performance character.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache RTR 200 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>200cc Fuel-Injected:</strong> Sport performance</li>
<li class="highlight-item"><strong>20.2 bhp Power:</strong> Exhilarating spirit</li>
<li class="highlight-item"><strong>Aggressive Styling:</strong> Street racer</li>
<li class="highlight-item"><strong>48+ km/l Efficiency:</strong> Good practical</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache RTR 200 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Exhilarating 20.2 bhp:</strong> Performance sport</li>
<li class="pro-item"><strong>Aggressive Styling:</strong> Street racer character</li>
<li class="pro-item"><strong>Sharp Handling:</strong> Responsive agile</li>
<li class="pro-item"><strong>TVS Reliability:</strong> Proven quality</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 48+ km/l practical</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache RTR 200 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳4,95,000 upscale</li>
<li class="con-item"><strong>Firm Suspension:</strong> Sport-focused ride</li>
<li class="con-item"><strong>Limited Comfort:</strong> Sporty seating</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Apache RTR 200 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳4,95,000 - Sport premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>200cc fuel-injected sport</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>20.2 bhp exhilarating</span></p>
<p class="value-item"><strong>Torque:</strong> <span>18.3 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>148kg sport-tuned</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>48+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Apache RTR 200 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 20.2 bhp exhilarating</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.7</span> - Sharp agile</li>
<li class="rating-item"><strong>Styling:</strong> <span>4.7</span> - Aggressive racer</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - TVS proven</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳4,95,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Apache RTR 200?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳4,95,000, the TVS Apache RTR 200 is perfect for youth racers seeking India's thrilling street racer, exhilarating 20.2 bhp power, aggressive sport styling, sharp handling, and accessible sport performance character.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For youth street racers</span></p>
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
<h1 class="review-main-title">টিভিএস অ্যাপাচ আরটিআর 200 রিভিউ বাংলাদেশ 2024 - ভারতের রোমাঞ্চকর পারফরম্যান্স রাস্তা রেসার</h1>
<p class="summary-text">টিভিএস অ্যাপাচ আরটিআর 200 ভারতের রোমাঞ্চকর 200cc পারফরম্যান্স রাস্তা রেসার যা অ্যাক্সেসযোগ্য খেলাধুলা শক্তি প্রতিনিধিত্ব করে। ৳4,95,000 টাকায় মূল্যায়িত, এটি 20.2 bhp উত্তেজনাপূর্ণ শক্তি এবং আক্রমণাত্মক খেলাধুলা স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">টিভিএস অ্যাপাচ আরটিআর 200 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>200cc জ্বালানী-ইনজেকশন:</strong> খেলাধুলা পারফরম্যান্স</li>
<li class="highlight-item"><strong>20.2 bhp শক্তি:</strong> উত্তেজনাপূর্ণ মন</li>
<li class="highlight-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> রাস্তা রেসার</li>
<li class="highlight-item"><strong>48+ km/l দক্ষতা:</strong> ভালো ব্যবহারিক</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">টিভিএস অ্যাপাচ আরটিআর 200 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>উত্তেজনাপূর্ণ 20.2 bhp:</strong> পারফরম্যান্স খেলাধুলা</li>
<li class="pro-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> রাস্তা রেসার চরিত্র</li>
<li class="pro-item"><strong>তীক্ষ্ণ হ্যান্ডলিং:</strong> প্রতিক্রিয়াশীল চটপটে</li>
<li class="pro-item"><strong>টিভিএস নির্ভরযোগ্যতা:</strong> প্রমাণিত গুণমান</li>
<li class="pro-item"><strong>ভালো দক্ষতা:</strong> 48+ km/l ব্যবহারিক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">টিভিএস অ্যাপাচ আরটিআর 200 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳4,95,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>দৃঢ় সাসপেনশন:</strong> খেলাধুলা-ফোকাসড রাইড</li>
<li class="con-item"><strong>সীমিত আরাম:</strong> খেলাধুলাপূর্ণ আসন</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: টিভিএস অ্যাপাচ আরটিআর 200 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳4,95,000 টাকায়, টিভিএস অ্যাপাচ আরটিআর 200 যুব রেসারদের জন্য নিখুঁত যারা ভারতের রোমাঞ্চকর রাস্তা রেসার, উত্তেজনাপূর্ণ 20.2 bhp শক্তি এবং আক্রমণাত্মক খেলাধুলা চরিত্র চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- যুব রাস্তা রেসারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created TVS Apache RTR 200\n")
	return nil
}

func (s *TVSApacheRTR200Batch23) GetName() string {
	return "TVSApacheRTR200Batch23"
}
