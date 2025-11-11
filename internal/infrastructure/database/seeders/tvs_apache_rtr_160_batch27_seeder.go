package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type TVSApacheRTR160Batch27 struct {
	BaseSeeder
}

func NewTVSApacheRTR160Batch27Seeder() *TVSApacheRTR160Batch27 {
	return &TVSApacheRTR160Batch27{BaseSeeder: BaseSeeder{name: "TVS Apache RTR 160 Batch27 Review"}}
}

func (s *TVSApacheRTR160Batch27) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "TVS Apache RTR 160").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("tvs apache rtr 160 product not found")
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
<h1 class="review-main-title">TVS Apache RTR 160 Review Bangladesh 2024 - Performance Warrior</h1>
<p class="summary-text">The TVS Apache RTR 160 is a dynamic performance warrior combining racing DNA and street capability. Priced around ৳3,55,000, it delivers 15.8 bhp power, aggressive styling, sharp handling, advanced features, and thrilling 50+ km/l efficiency for sport riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS Apache RTR 160 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>160cc Sport:</strong> Warrior performance</li>
<li class="highlight-item"><strong>15.8 bhp Power:</strong> Thrilling responsive</li>
<li class="highlight-item"><strong>Racing DNA:</strong> Track inspired</li>
<li class="highlight-item"><strong>50+ km/l Efficiency:</strong> Sport economical</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS Apache RTR 160 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Racing Heritage:</strong> Track DNA proven</li>
<li class="pro-item"><strong>Thrilling Performance:</strong> 15.8 bhp exciting</li>
<li class="pro-item"><strong>Sharp Handling:</strong> Precision agile</li>
<li class="pro-item"><strong>Advanced Features:</strong> Modern tech equipped</li>
<li class="pro-item"><strong>Sport Efficiency:</strong> 50+ km/l balanced</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS Apache RTR 160 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Sport Positioning:</strong> Aggressive riding stance</li>
<li class="con-item"><strong>Limited Comfort:</strong> Performance focused</li>
<li class="con-item"><strong>Service Availability:</strong> Growing network</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">TVS Apache RTR 160 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,55,000 - Sport warrior</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>160cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>15.8 bhp thrilling</span></p>
<p class="value-item"><strong>Torque:</strong> <span>14 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>139kg agile</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50+ km/l sport</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">TVS Apache RTR 160 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Aggressive racing</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.5</span> - Sport focused</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.8</span> - 15.8 bhp thrilling</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.6</span> - 50+ km/l sport</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.6</span> - TVS reliable</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy TVS Apache RTR 160?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,55,000, the TVS Apache RTR 160 is perfect for sport riders seeking racing DNA heritage, thrilling 15.8 bhp performance, sharp handling, and balanced 50+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sport riders</span></p>
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
<h1 class="review-main-title">TVS অ্যাপাচি RTR 160 রিভিউ বাংলাদেশ 2024 - পারফরম্যান্স ওয়ারিয়র</h1>
<p class="summary-text">TVS অ্যাপাচি RTR 160 একটি গতিশীল পারফরম্যান্স ওয়ারিয়র যা রেসিং ডিএনএ এবং রাস্তার সক্ষমতা একত্রিত করে। ৳3,55,000 টাকায় মূল্যায়িত, এটি 15.8 bhp শক্তি এবং আক্রমণাত্মক স্টাইলিং বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">TVS অ্যাপাচি RTR 160 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>160cc স্পোর্ট:</strong> ওয়ারিয়র পারফরম্যান্স</li>
<li class="highlight-item"><strong>15.8 bhp শক্তি:</strong> রোমাঞ্চকর প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>রেসিং ডিএনএ:</strong> ট্র্যাক অনুপ্রাণিত</li>
<li class="highlight-item"><strong>50+ km/l দক্ষতা:</strong> স্পোর্ট সাশ্রয়ী</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">TVS অ্যাপাচি RTR 160 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>রেসিং ঐতিহ্য:</strong> ট্র্যাক ডিএনএ প্রমাণিত</li>
<li class="pro-item"><strong>রোমাঞ্চকর কর্মক্ষমতা:</strong> 15.8 bhp উত্তেজনাপূর্ণ</li>
<li class="pro-item"><strong>তীক্ষ্ণ হ্যান্ডলিং:</strong> নির্ভুল চতুর</li>
<li class="pro-item"><strong>উন্নত বৈশিষ্ট্য:</strong> আধুনিক প্রযুক্তি সজ্জিত</li>
<li class="pro-item"><strong>স্পোর্ট দক্ষতা:</strong> 50+ km/l সুষম</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">TVS অ্যাপাচি RTR 160 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>স্পোর্ট অবস্থান:</strong> আক্রমণাত্মক রাইডিং অবস্থান</li>
<li class="con-item"><strong>সীমিত আরাম:</strong> পারফরম্যান্স দৃষ্টিনিবদ্ধ</li>
<li class="con-item"><strong>সেবা প্রাপ্যতা:</strong> ক্রমবর্ধমান নেটওয়ার্ক</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: TVS অ্যাপাচি RTR 160 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,55,000 টাকায়, TVS অ্যাপাচি RTR 160 স্পোর্ট রাইডারদের জন্য নিখুঁত যারা রেসিং ডিএনএ ঐতিহ্য এবং রোমাঞ্চকর 15.8 bhp কর্মক্ষমতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্পোর্ট রাইডারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created TVS Apache RTR 160\n")
	return nil
}

func (s *TVSApacheRTR160Batch27) GetName() string {
	return "TVSApacheRTR160Batch27"
}
