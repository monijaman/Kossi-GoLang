package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type SymBike100Batch22 struct {
	BaseSeeder
}

func NewSymBike100Batch22Seeder() *SymBike100Batch22 {
	return &SymBike100Batch22{BaseSeeder: BaseSeeder{name: "Sym Bike 100 Batch22 Review"}}
}

func (s *SymBike100Batch22) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Sym Bike 100").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("sym bike 100 product not found")
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
<h1 class="review-main-title">Sym Bike 100 Review Bangladesh 2024 - Taiwan's Reliable Budget Commuter</h1>
<p class="summary-text">The Sym Bike 100 is a 100cc air-cooled reliable budget commuter motorcycle representing Taiwan's practical engineering heritage. Priced around ৳1,05,000, it delivers solid 7 bhp power, excellent build quality, strong reliability, simple maintenance, and exceptional 60+ km/l fuel efficiency for practical daily commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Sym Bike 100 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc Air-Cooled:</strong> Taiwan reliable</li>
<li class="highlight-item"><strong>7 bhp Power:</strong> Solid practical</li>
<li class="highlight-item"><strong>60+ km/l Efficiency:</strong> Exceptional economy</li>
<li class="highlight-item"><strong>Strong Build Quality:</strong> Taiwan engineering</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Sym Bike 100 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Strong Build Quality:</strong> Taiwan engineering</li>
<li class="pro-item"><strong>Excellent Reliability:</strong> Solid construction</li>
<li class="pro-item"><strong>Exceptional Efficiency:</strong> 60+ km/l outstanding</li>
<li class="pro-item"><strong>Simple Maintenance:</strong> Basic robust design</li>
<li class="pro-item"><strong>Affordable Price:</strong> ৳1,05,000 value</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Sym Bike 100 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Basic Power:</strong> 7 bhp minimal</li>
<li class="con-item"><strong>Simple Styling:</strong> Functional design</li>
<li class="con-item"><strong>No Frills:</strong> Entry-level features</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Sym Bike 100 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,05,000 - Taiwan reliable budget</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>100cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>7 bhp solid</span></p>
<p class="value-item"><strong>Torque:</strong> <span>8 nm practical</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>105kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>60+ km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Sym Bike 100 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Build Quality:</strong> <span>4.8</span> - Taiwan engineering</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.8</span> - Solid construction</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>5.0</span> - 60+ km/l exceptional</li>
<li class="rating-item"><strong>Maintenance:</strong> <span>4.8</span> - Simple robust</li>
<li class="rating-item"><strong>Value:</strong> <span>4.8</span> - ৳1,05,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Sym Bike 100?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,05,000, the Sym Bike 100 is the practical choice for reliable commuters seeking Taiwan engineering, solid build quality, exceptional 60+ km/l efficiency, simple maintenance, and affordable daily transportation.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For reliable practical commuters</span></p>
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
<h1 class="review-main-title">সিম বাইক 100 রিভিউ বাংলাদেশ 2024 - তাইওয়ানের নির্ভরযোগ্য বাজেট কমিউটার</h1>
<p class="summary-text">সিম বাইক 100 একটি 100cc এয়ার-কুল্ড নির্ভরযোগ্য বাজেট কমিউটার মোটরসাইকেল যা তাইওয়ানের ব্যবহারিক ইঞ্জিনিয়ারিং ঐতিহ্য প্রতিনিধিত্ব করে। ৳1,05,000 টাকায় মূল্যায়িত, এটি শক্ত 7 bhp শক্তি এবং ব্যতিক্রমী 60+ km/l জ্বালানি দক্ষতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">সিম বাইক 100 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>100cc এয়ার-কুল্ড:</strong> তাইওয়ান নির্ভরযোগ্য</li>
<li class="highlight-item"><strong>7 bhp শক্তি:</strong> শক্ত ব্যবহারিক</li>
<li class="highlight-item"><strong>60+ km/l দক্ষতা:</strong> ব্যতিক্রমী অর্থনীতি</li>
<li class="highlight-item"><strong>শক্তিশালী বিল্ড গুণমান:</strong> তাইওয়ান ইঞ্জিনিয়ারিং</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">সিম বাইক 100 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>শক্তিশালী বিল্ড গুণমান:</strong> তাইওয়ান ইঞ্জিনিয়ারিং</li>
<li class="pro-item"><strong>চমৎকার নির্ভরযোগ্যতা:</strong> শক্ত নির্মাণ</li>
<li class="pro-item"><strong>ব্যতিক্রমী দক্ষতা:</strong> 60+ km/l অসাধারণ</li>
<li class="pro-item"><strong>সহজ রক্ষণাবেক্ষণ:</strong> মৌলিক শক্তিশালী ডিজাইন</li>
<li class="pro-item"><strong>সাশ্রয়ী মূল্য:</strong> ৳1,05,000 মূল্য</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">সিম বাইক 100 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মৌলিক শক্তি:</strong> 7 bhp ন্যূনতম</li>
<li class="con-item"><strong>সহজ স্টাইলিং:</strong> কার্যকরী ডিজাইন</li>
<li class="con-item"><strong>কোনো ফ্রিল নেই:</strong> এন্ট্রি-লেভেল বৈশিষ্ট্য</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: সিম বাইক 100 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,05,000 টাকায়, সিম বাইক 100 নির্ভরযোগ্য কমিউটারদের জন্য ব্যবহারিক পছন্দ যারা তাইওয়ান ইঞ্জিনিয়ারিং, শক্ত বিল্ড কোয়ালিটি এবং ব্যতিক্রমী 60+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- নির্ভরযোগ্য ব্যবহারিক কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Sym Bike 100\n")
	return nil
}

func (s *SymBike100Batch22) GetName() string {
	return "SymBike100Batch22"
}
