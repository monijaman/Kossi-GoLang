package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type YamahaFZ25Batch27 struct {
	BaseSeeder
}

func NewYamahaFZ25Batch27Seeder() *YamahaFZ25Batch27 {
	return &YamahaFZ25Batch27{BaseSeeder: BaseSeeder{name: "Yamaha FZ 25 Batch27 Review"}}
}

func (s *YamahaFZ25Batch27) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha FZ 25").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha fz 25 product not found")
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
<h1 class="review-main-title">Yamaha FZ 25 Review Bangladesh 2024 - Naked Street King</h1>
<p class="summary-text">The Yamaha FZ 25 is a powerful naked street king combining aggressive styling and urban dominance. Priced around ৳4,35,000, it delivers 20.9 bhp power, muscular design, comfortable riding, excellent handling, and impressive 45+ km/l efficiency for street warriors.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha FZ 25 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>250cc Naked:</strong> Street king powerful</li>
<li class="highlight-item"><strong>20.9 bhp Power:</strong> Strong responsive</li>
<li class="highlight-item"><strong>Muscular Design:</strong> Aggressive naked</li>
<li class="highlight-item"><strong>45+ km/l Efficiency:</strong> Impressive balanced</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha FZ 25 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Powerful Performance:</strong> 20.9 bhp strong</li>
<li class="pro-item"><strong>Muscular Styling:</strong> Aggressive street fighter</li>
<li class="pro-item"><strong>Excellent Handling:</strong> Precision agile</li>
<li class="pro-item"><strong>Comfortable Seating:</strong> Daily ride capable</li>
<li class="pro-item"><strong>Yamaha Quality:</strong> Japanese reliability</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha FZ 25 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Higher Price:</strong> Premium positioning</li>
<li class="con-item"><strong>Limited Features:</strong> Basic technology</li>
<li class="con-item"><strong>Moderate Efficiency:</strong> 45+ km/l sport</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha FZ 25 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳4,35,000 - Naked street</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>250cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>20.9 bhp powerful</span></p>
<p class="value-item"><strong>Torque:</strong> <span>20 nm strong</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>148kg balanced</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>45+ km/l impressive</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha FZ 25 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Muscular aggressive</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.6</span> - Daily capable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 20.9 bhp strong</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.6</span> - 45+ km/l balanced</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.8</span> - Yamaha premium</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha FZ 25?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳4,35,000, the Yamaha FZ 25 is perfect for street warriors seeking powerful 20.9 bhp performance, muscular naked styling, excellent handling, and impressive 45+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For street warriors</span></p>
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
<h1 class="review-main-title">ইয়ামাহা FZ 25 রিভিউ বাংলাদেশ 2024 - নগ্ন স্ট্রিট কিং</h1>
<p class="summary-text">ইয়ামাহা FZ 25 একটি শক্তিশালী নগ্ন স্ট্রিট কিং যা আক্রমণাত্মক স্টাইলিং এবং শহুরে আধিপত্য একত্রিত করে। ৳4,35,000 টাকায় মূল্যায়িত, এটি 20.9 bhp শক্তি এবং পেশীবহুল ডিজাইন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা FZ 25 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>250cc নগ্ন:</strong> স্ট্রিট কিং শক্তিশালী</li>
<li class="highlight-item"><strong>20.9 bhp শক্তি:</strong> শক্তিশালী প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>পেশীবহুল ডিজাইন:</strong> আক্রমণাত্মক নগ্ন</li>
<li class="highlight-item"><strong>45+ km/l দক্ষতা:</strong> চিত্তাকর্ষক সুষম</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা FZ 25 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>শক্তিশালী কর্মক্ষমতা:</strong> 20.9 bhp শক্তিশালী</li>
<li class="pro-item"><strong>পেশীবহুল স্টাইলিং:</strong> আক্রমণাত্মক স্ট্রিট ফাইটার</li>
<li class="pro-item"><strong>চমৎকার হ্যান্ডলিং:</strong> নির্ভুল চতুর</li>
<li class="pro-item"><strong>আরামদায়ক আসন:</strong> দৈনিক যাত্রা সক্ষম</li>
<li class="pro-item"><strong>ইয়ামাহা গুণমান:</strong> জাপানী নির্ভরযোগ্যতা</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা FZ 25 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>উচ্চতর মূল্য:</strong> প্রিমিয়াম অবস্থান</li>
<li class="con-item"><strong>সীমিত বৈশিষ্ট্য:</strong> মৌলিক প্রযুক্তি</li>
<li class="con-item"><strong>মধ্যম দক্ষতা:</strong> 45+ km/l স্পোর্ট</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা FZ 25 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳4,35,000 টাকায়, ইয়ামাহা FZ 25 স্ট্রিট যোদ্ধাদের জন্য নিখুঁত যারা শক্তিশালী 20.9 bhp কর্মক্ষমতা এবং চিত্তাকর্ষক 45+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- স্ট্রিট যোদ্ধাদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Yamaha FZ 25\n")
	return nil
}

func (s *YamahaFZ25Batch27) GetName() string {
	return "YamahaFZ25Batch27"
}
