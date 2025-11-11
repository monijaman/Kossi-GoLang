package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type MaybachXTX110ReviewBatch20 struct {
	BaseSeeder
}

func NewMaybachXTX110ReviewBatch20Seeder() *MaybachXTX110ReviewBatch20 {
	return &MaybachXTX110ReviewBatch20{BaseSeeder: BaseSeeder{name: "Maybach XTX 110 Batch20 Review"}}
}

func (s *MaybachXTX110ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Maybach XTX 110").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("maybach xtx 110 product not found")
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
<h1 class="review-main-title">Maybach XTX 110 Review Bangladesh 2024 - Sporty Youth Commuter</h1>
<p class="summary-text">The Maybach XTX 110 is a 110cc air-cooled sporty youth commuter representing practical modern design for young riders. Priced around ৳1,50,000, it delivers sporty styling, 8.5 bhp power, practical fuel efficiency, youthful appeal, and value-for-money for young commuters seeking stylish daily transport.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Maybach XTX 110 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc Air-Cooled:</strong> Youth commuter</li>
<li class="highlight-item"><strong>Sporty Design:</strong> Stylish modern</li>
<li class="highlight-item"><strong>50+ km/l:</strong> Practical efficiency</li>
<li class="highlight-item"><strong>Budget Value:</strong> ৳1,50,000</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Maybach XTX 110 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Sporty Design:</strong> Youth appeal styling</li>
<li class="pro-item"><strong>Budget Price:</strong> ৳1,50,000 affordable</li>
<li class="pro-item"><strong>Fuel Economy:</strong> 50+ km/l practical</li>
<li class="pro-item"><strong>Easy Maintenance:</strong> Simple service</li>
<li class="pro-item"><strong>Reliable Build:</strong> Proven quality</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Maybach XTX 110 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Basic Features:</strong> Minimal equipment</li>
<li class="con-item"><strong>Low Power:</strong> 8.5 bhp modest</li>
<li class="con-item"><strong>Limited Storage:</strong> Minimal carrying</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Maybach XTX 110 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳1,50,000 - Budget youth segment</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>110cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>8.5 bhp practical</span></p>
<p class="value-item"><strong>Torque:</strong> <span>9 nm adequate</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>115kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50-55 km/l exceptional</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Maybach XTX 110 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.6</span> - Sporty youth appeal</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 50+ km/l</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳1,50,000</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.5</span> - Practical seating</li>
<li class="rating-item"><strong>Reliability:</strong> <span>4.5</span> - Proven quality</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Maybach XTX 110?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳1,50,000, the XTX 110 is an excellent youth commuter choice offering sporty styling, practical 50+ km/l efficiency, budget value, reliable build quality, and stylish daily transport for young riders seeking affordable practical commuting.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For young budget commuters</span></p>
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
<h1 class="review-main-title">মায়বাক XTX 110 রিভিউ বাংলাদেশ 2024 - ক্রীড়াবহুল যুব কমিউটার</h1>
<p class="summary-text">মায়বাক XTX 110 একটি 110cc এয়ার-কুল্ড ক্রীড়াবহুল যুব কমিউটার যা তরুণ রাইডারদের জন্য ব্যবহারিক আধুনিক ডিজাইন প্রতিনিধিত্ব করে। ৳1,50,000 টাকায় মূল্যায়িত, এটি ক্রীড়াবহুল স্টাইলিং এবং ব্যবহারিক জ্বালানি দক্ষতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">মায়বাক XTX 110 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>110cc এয়ার-কুল্ড:</strong> যুব কমিউটার</li>
<li class="highlight-item"><strong>ক্রীড়াবহুল ডিজাইন:</strong> আধুনিক স্টাইলিশ</li>
<li class="highlight-item"><strong>50+ km/l:</strong> ব্যবহারিক দক্ষতা</li>
<li class="highlight-item"><strong>বাজেট মূল্য:</strong> ৳1,50,000</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">মায়বাক XTX 110 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>ক্রীড়াবহুল ডিজাইন:</strong> যুব আবেদন স্টাইলিং</li>
<li class="pro-item"><strong>বাজেট মূল্য:</strong> ৳1,50,000 সাশ্রয়ী</li>
<li class="pro-item"><strong>জ্বালানি অর্থনীতি:</strong> 50+ km/l ব্যবহারিক</li>
<li class="pro-item"><strong>সহজ রক্ষণাবেক্ষণ:</strong> সরল সেবা</li>
<li class="pro-item"><strong>নির্ভরযোগ্য বিল্ড:</strong> প্রমাণিত গুণমান</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">মায়বাক XTX 110 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>মৌলিক বৈশিষ্ট্য:</strong> ন্যূনতম সরঞ্জাম</li>
<li class="con-item"><strong>নিম্ন শক্তি:</strong> 8.5 bhp বিনম্র</li>
<li class="con-item"><strong>সীমিত স্টোরেজ:</strong> ন্যূনতম বহন</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: মায়বাক XTX 110 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳1,50,000 টাকায়, XTX 110 যুব কমিউটারদের জন্য চমৎকার পছন্দ যা ক্রীড়াবহুল স্টাইলিং, ব্যবহারিক দক্ষতা এবং সাশ্রয়ী মূল্যে স্টাইলিশ দৈনন্দিন পরিবহন খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- তরুণ বাজেট কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Maybach XTX 110\n")
	return nil
}

func (s *MaybachXTX110ReviewBatch20) GetName() string {
	return "MaybachXTX110ReviewBatch20"
}
