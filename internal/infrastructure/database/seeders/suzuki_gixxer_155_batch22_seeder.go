package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type SuzukiGixxer155Batch22 struct {
	BaseSeeder
}

func NewSuzukiGixxer155Batch22Seeder() *SuzukiGixxer155Batch22 {
	return &SuzukiGixxer155Batch22{BaseSeeder: BaseSeeder{name: "Suzuki Gixxer 155 Batch22 Review"}}
}

func (s *SuzukiGixxer155Batch22) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Suzuki Gixxer 155").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("suzuki gixxer 155 product not found")
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
<h1 class="review-main-title">Suzuki Gixxer 155 Review Bangladesh 2024 - India's Most Aggressive Sporty Street Performer</h1>
<p class="summary-text">The Suzuki Gixxer 155 is India's most aggressive 155cc sporty street performer representing ultimate youth aggression. Priced around ৳3,95,000, it delivers 14.8 bhp spirited power, aggressive styling, excellent handling, good 48+ km/l efficiency, and outstanding value for youth sport seekers wanting authentic street performance daily.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Suzuki Gixxer 155 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>155cc Fuel-Injected:</strong> Sport aggression</li>
<li class="highlight-item"><strong>14.8 bhp Power:</strong> Spirited responsive</li>
<li class="highlight-item"><strong>Aggressive Styling:</strong> Street performer attitude</li>
<li class="highlight-item"><strong>Excellent Handling:</strong> Agile performance</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Suzuki Gixxer 155 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Aggressive Styling:</strong> Authentic street performer</li>
<li class="pro-item"><strong>Spirited 14.8 bhp:</strong> Youth aggression</li>
<li class="pro-item"><strong>Excellent Handling:</strong> Agile responsive</li>
<li class="pro-item"><strong>Suzuki Reputation:</strong> Quality engineering</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 48+ km/l practical</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Suzuki Gixxer 155 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳3,95,000 youth sport</li>
<li class="con-item"><strong>Sport Focused:</strong> Limited comfort</li>
<li class="con-item"><strong>Build Quality:</strong> Aggressive positioning</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Suzuki Gixxer 155 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,95,000 - Sport premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>155cc fuel-injected sport</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>14.8 bhp spirited</span></p>
<p class="value-item"><strong>Torque:</strong> <span>14 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>138kg sport-tuned</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>48+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Suzuki Gixxer 155 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Aggression:</strong> <span>4.8</span> - Street performer</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.7</span> - Agile excellent</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 14.8 bhp</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - Suzuki engineering</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳3,95,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Suzuki Gixxer 155?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,95,000, the Suzuki Gixxer 155 is perfect for youth sport seekers wanting India's most aggressive street performer, spirited 14.8 bhp power, authentic aggressive styling, excellent handling, and street performance daily.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For aggressive youth sport performers</span></p>
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
<h1 class="review-main-title">সুজুকি গিক্সার 155 রিভিউ বাংলাদেশ 2024 - ভারতের সবচেয়ে আক্রমণাত্মক খেলাধুলাপূর্ণ রাস্তা পারফর্মার</h1>
<p class="summary-text">সুজুকি গিক্সার 155 ভারতের সবচেয়ে আক্রমণাত্মক 155cc খেলাধুলাপূর্ণ রাস্তা পারফর্মার যা চূড়ান্ত যুব আক্রমণ প্রতিনিধিত্ব করে। ৳3,95,000 টাকায় মূল্যায়িত, এটি 14.8 bhp প্রাণবন্ত শক্তি এবং আক্রমণাত্মক স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সুজুকি গিক্সার 155 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>155cc জ্বালানী-ইনজেকশন:</strong> স্পোর্ট আক্রমণ</li>
<li class="highlight-item"><strong>14.8 bhp শক্তি:</strong> প্রাণবন্ত প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> রাস্তা পারফর্মার মনোভাব</li>
<li class="highlight-item"><strong>চমৎকার হ্যান্ডলিং:</strong> চটপটে কর্মক্ষমতা</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সুজুকি গিক্সার 155 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> খাঁটি রাস্তা পারফর্মার</li>
<li class="pro-item"><strong>প্রাণবন্ত 14.8 bhp:</strong> যুব আক্রমণ</li>
<li class="pro-item"><strong>চমৎকার হ্যান্ডলিং:</strong> চটপটে প্রতিক্রিয়াশীল</li>
<li class="pro-item"><strong>সুজুকি খ্যাতি:</strong> গুণমান প্রকৌশল</li>
<li class="pro-item"><strong>ভালো দক্ষতা:</strong> 48+ km/l ব্যবহারিক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সুজুকি গিক্সার 155 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳3,95,000 যুব স্পোর্ট</li>
<li class="con-item"><strong>স্পোর্ট ফোকাসড:</strong> সীমিত আরাম</li>
<li class="con-item"><strong>বিল্ড গুণমান:</strong> আক্রমণাত্মক অবস্থান</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সুজুকি গিক্সার 155 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,95,000 টাকায়, সুজুকি গিক্সার 155 যুব স্পোর্ট খোঁজার জন্য নিখুঁত যারা ভারতের সবচেয়ে আক্রমণাত্মক রাস্তা পারফর্মার এবং প্রাণবন্ত 14.8 bhp শক্তি চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আক্রমণাত্মক যুব স্পোর্ট পারফরমারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Suzuki Gixxer 155\n")
	return nil
}

func (s *SuzukiGixxer155Batch22) GetName() string {
	return "SuzukiGixxer155Batch22"
}
