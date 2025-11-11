package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type YamahaR15V4Batch26 struct {
	BaseSeeder
}

func NewYamahaR15V4Batch26Seeder() *YamahaR15V4Batch26 {
	return &YamahaR15V4Batch26{BaseSeeder: BaseSeeder{name: "Yamaha R15 V4 Batch26 Review"}}
}

func (s *YamahaR15V4Batch26) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha R15 V4").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha r15 v4 product not found")
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
<h1 class="review-main-title">Yamaha R15 V4 Review Bangladesh 2024 - Premium Sport Icon</h1>
<p class="summary-text">The Yamaha R15 V4 is Bangladesh's ultimate premium sport motorcycle combining Japanese engineering excellence with thrilling performance. Priced around ৳4,85,000, it delivers 18.4 bhp power, advanced technology, refined handling, aggressive styling, and unforgettable riding dynamics.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha R15 V4 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>155cc Sport:</strong> Premium icon flagship</li>
<li class="highlight-item"><strong>18.4 bhp Power:</strong> Thrilling responsive</li>
<li class="highlight-item"><strong>Advanced Tech:</strong> ABS enabled modern</li>
<li class="highlight-item"><strong>Aggressive Styling:</strong> Premium contemporary</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha R15 V4 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Premium Engineering:</strong> Yamaha excellence</li>
<li class="pro-item"><strong>Thrilling Performance:</strong> 18.4 bhp exciting</li>
<li class="pro-item"><strong>Advanced ABS:</strong> Safety technology</li>
<li class="pro-item"><strong>Refined Handling:</strong> Precision control</li>
<li class="pro-item"><strong>Aggressive Styling:</strong> Modern premium</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha R15 V4 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Higher Cost:</strong> Premium pricing</li>
<li class="con-item"><strong>Sport Focus:</strong> Not family commuter</li>
<li class="con-item"><strong>Fuel Consumption:</strong> Sport bike economy</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha R15 V4 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳4,85,000 - Premium sport</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>155cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>18.4 bhp thrilling</span></p>
<p class="value-item"><strong>Torque:</strong> <span>15 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>142kg agile</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>45+ km/l sport</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha R15 V4 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.9</span> - Aggressive premium</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.6</span> - Sport positioned</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.9</span> - 18.4 bhp thrilling</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.5</span> - 45+ km/l sport</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.8</span> - Yamaha premium</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha R15 V4?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳4,85,000, the Yamaha R15 V4 is perfect for enthusiasts seeking premium Japanese engineering, thrilling 18.4 bhp performance, advanced ABS technology, and unforgettable sport riding dynamics.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sport enthusiasts</span></p>
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
<h1 class="review-main-title">ইয়ামাহা R15 V4 রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম স্পোর্ট আইকন</h1>
<p class="summary-text">ইয়ামাহা R15 V4 হল বাংলাদেশের চূড়ান্ত প্রিমিয়াম স্পোর্ট মোটরসাইকেল যা জাপানী প্রকৌশল উৎকর্ষতা এবং রোমাঞ্চকর কর্মক্ষমতা একত্রিত করে। ৳4,85,000 টাকায় মূল্যায়িত, এটি 18.4 bhp শক্তি এবং আধুনিক প্রযুক্তি বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা R15 V4 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>155cc স্পোর্ট:</strong> প্রিমিয়াম আইকন ফ্ল্যাগশিপ</li>
<li class="highlight-item"><strong>18.4 bhp শক্তি:</strong> রোমাঞ্চকর প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>উন্নত প্রযুক্তি:</strong> ABS সক্ষম আধুনিক</li>
<li class="highlight-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> প্রিমিয়াম সমসাময়িক</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা R15 V4 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রিমিয়াম প্রকৌশল:</strong> ইয়ামাহা উৎকর্ষতা</li>
<li class="pro-item"><strong>রোমাঞ্চকর কর্মক্ষমতা:</strong> 18.4 bhp উত্তেজনাপূর্ণ</li>
<li class="pro-item"><strong>উন্নত ABS:</strong> নিরাপত্তা প্রযুক্তি</li>
<li class="pro-item"><strong>পরিমার্জিত হ্যান্ডলিং:</strong> নির্ভুল নিয়ন্ত্রণ</li>
<li class="pro-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> আধুনিক প্রিমিয়াম</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা R15 V4 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>উচ্চতর খরচ:</strong> প্রিমিয়াম মূল্য নির্ধারণ</li>
<li class="con-item"><strong>স্পোর্ট ফোকাস:</strong> পারিবারিক যাতায়াত নয়</li>
<li class="con-item"><strong>জ্বালানি খরচ:</strong> স্পোর্ট বাইক অর্থনীতি</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা R15 V4 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳4,85,000 টাকায়, ইয়ামাহা R15 V4 উৎসাহীদের জন্য নিখুঁত যারা প্রিমিয়াম জাপানী প্রকৌশল এবং রোমাঞ্চকর 18.4 bhp কর্মক্ষমতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্পোর্ট উৎসাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Yamaha R15 V4\n")
	return nil
}

func (s *YamahaR15V4Batch26) GetName() string {
	return "YamahaR15V4Batch26"
}
