package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type YamahaYZF150Batch22 struct {
	BaseSeeder
}

func NewYamahaYZF150Batch22Seeder() *YamahaYZF150Batch22 {
	return &YamahaYZF150Batch22{BaseSeeder: BaseSeeder{name: "Yamaha YZF 150 Batch22 Review"}}
}

func (s *YamahaYZF150Batch22) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha YZF 150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha yzf 150 product not found")
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
<h1 class="review-main-title">Yamaha YZF 150 Review Bangladesh 2024 - Thailand's Aggressive Sport Street Fighter</h1>
<p class="summary-text">The Yamaha YZF 150 is a Thailand-built 150cc aggressive sport street fighter representing ultimate youth performance. Priced around ৳3,75,000, it delivers 15 bhp spirited power, responsive handling, aggressive styling, good 48+ km/l efficiency, and excellent value for youth sport enthusiasts seeking street-ready aggressive performance.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha YZF 150 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Fuel-Injected:</strong> Sport performance</li>
<li class="highlight-item"><strong>15 bhp Power:</strong> Spirited aggressive</li>
<li class="highlight-item"><strong>Responsive Handling:</strong> Street-ready agile</li>
<li class="highlight-item"><strong>Aggressive Styling:</strong> Street fighter attitude</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha YZF 150 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Spirited 15 bhp:</strong> Youth performance</li>
<li class="pro-item"><strong>Responsive Handling:</strong> Street-ready agile</li>
<li class="pro-item"><strong>Aggressive Styling:</strong> Authentic sport character</li>
<li class="pro-item"><strong>Yamaha Reputation:</strong> Quality engineering</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 48+ km/l practical</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha YZF 150 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳3,75,000 youth</li>
<li class="con-item"><strong>Firm Suspension:</strong> Sport-focused comfort</li>
<li class="con-item"><strong>Sporty Seating:</strong> Limited capacity</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha YZF 150 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳3,75,000 - Youth sport premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc fuel-injected sport</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>15 bhp spirited</span></p>
<p class="value-item"><strong>Torque:</strong> <span>14 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>142kg sport-tuned</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>48+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha YZF 150 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 15 bhp spirited</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.7</span> - Responsive agile</li>
<li class="rating-item"><strong>Style:</strong> <span>4.8</span> - Aggressive authentic</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Yamaha engineering</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳3,75,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha YZF 150?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳3,75,000, the Yamaha YZF 150 is perfect for youth sport enthusiasts seeking Thailand's aggressive street fighter, spirited 15 bhp performance, responsive handling, authentic aggressive styling, and street-ready sport character daily.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For youth sport enthusiasts</span></p>
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
<h1 class="review-main-title">ইয়ামাহা ওয়াইজেড এফ 150 রিভিউ বাংলাদেশ 2024 - থাইল্যান্ডের আক্রমণাত্মক স্পোর্ট রাস্তা ফাইটার</h1>
<p class="summary-text">ইয়ামাহা ওয়াইজেড এফ 150 থাইল্যান্ড-নির্মিত 150cc আক্রমণাত্মক স্পোর্ট রাস্তা ফাইটার যা চূড়ান্ত যুব কর্মক্ষমতা প্রতিনিধিত্ব করে। ৳3,75,000 টাকায় মূল্যায়িত, এটি 15 bhp প্রাণবন্ত শক্তি এবং প্রতিক্রিয়াশীল হ্যান্ডলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা ওয়াইজেড এফ 150 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc জ্বালানী-ইনজেকশন:</strong> স্পোর্ট কর্মক্ষমতা</li>
<li class="highlight-item"><strong>15 bhp শক্তি:</strong> প্রাণবন্ত আক্রমণাত্মক</li>
<li class="highlight-item"><strong>প্রতিক্রিয়াশীল হ্যান্ডলিং:</strong> রাস্তা-প্রস্তুত চটপটে</li>
<li class="highlight-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> রাস্তা ফাইটার মনোভাব</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা ওয়াইজেড এফ 150 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>প্রাণবন্ত 15 bhp:</strong> যুব কর্মক্ষমতা</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল হ্যান্ডলিং:</strong> রাস্তা-প্রস্তুত চটপটে</li>
<li class="pro-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> খাঁটি স্পোর্ট চরিত্র</li>
<li class="pro-item"><strong>ইয়ামাহা খ্যাতি:</strong> গুণমান প্রকৌশল</li>
<li class="pro-item"><strong>ভালো দক্ষতা:</strong> 48+ km/l ব্যবহারিক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা ওয়াইজেড এফ 150 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳3,75,000 যুব</li>
<li class="con-item"><strong>দৃঢ় সাসপেনশন:</strong> স্পোর্ট-ফোকাসড আরাম</li>
<li class="con-item"><strong>খেলাধুলাপূর্ণ আসন:</strong> সীমিত ক্ষমতা</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা ওয়াইজেড এফ 150 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳3,75,000 টাকায়, ইয়ামাহা ওয়াইজেড এফ 150 যুব স্পোর্ট উৎসাহীদের জন্য নিখুঁত যারা থাইল্যান্ডের আক্রমণাত্মক রাস্তা ফাইটার এবং প্রাণবন্ত 15 bhp কর্মক্ষমতা খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- যুব স্পোর্ট উৎসাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Yamaha YZF 150\n")
	return nil
}

func (s *YamahaYZF150Batch22) GetName() string {
	return "YamahaYZF150Batch22"
}
