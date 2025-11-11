package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type KTMDuke200Batch24 struct {
	BaseSeeder
}

func NewKTMDuke200Batch24Seeder() *KTMDuke200Batch24 {
	return &KTMDuke200Batch24{BaseSeeder: BaseSeeder{name: "KTM Duke 200 Batch24 Review"}}
}

func (s *KTMDuke200Batch24) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "KTM Duke 200").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("ktm duke 200 product not found")
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
<h1 class="review-main-title">KTM Duke 200 Review Bangladesh 2024 - Sporty Performance Naked</h1>
<p class="summary-text">The KTM Duke 200 is a sporty 200cc naked bike delivering impressive performance. Priced around ৳5,95,000, it features 24.4 bhp power, aggressive design, lightweight frame, responsive handling, and excellent 50+ km/l efficiency for riders seeking thrilling naked bike performance.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">KTM Duke 200 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>200cc Sporty:</strong> Performance focused</li>
<li class="highlight-item"><strong>24.4 bhp Power:</strong> Thrilling performance</li>
<li class="highlight-item"><strong>Aggressive Design:</strong> Modern styled</li>
<li class="highlight-item"><strong>50+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">KTM Duke 200 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Sporty Performance:</strong> 24.4 bhp thrilling</li>
<li class="pro-item"><strong>Aggressive Design:</strong> Modern styling</li>
<li class="pro-item"><strong>Lightweight Frame:</strong> Responsive handling</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 50+ km/l decent</li>
<li class="pro-item"><strong>KTM Quality:</strong> Performance engineering</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">KTM Duke 200 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳5,95,000 upscale</li>
<li class="con-item"><strong>Limited Service:</strong> KTM dealers only</li>
<li class="con-item"><strong>Harder Suspension:</strong> Sporty oriented</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">KTM Duke 200 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳5,95,000 - Performance oriented</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>200cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>24.4 bhp thrilling</span></p>
<p class="value-item"><strong>Torque:</strong> <span>19.5 nm performance</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>148kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">KTM Duke 200 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Aggressive modern</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.5</span> - Sporty configured</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.8</span> - 24.4 bhp thrilling</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.6</span> - 50+ km/l good</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - KTM proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy KTM Duke 200?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳5,95,000, the KTM Duke 200 is ideal for performance enthusiasts seeking thrilling 24.4 bhp power, aggressive styling, responsive handling, and good 50+ km/l efficiency in a sporty naked package.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance naked enthusiasts</span></p>
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
<h1 class="review-main-title">কেটিএম ডিউক 200 রিভিউ বাংলাদেশ 2024 - খেলাধুলামূলক পারফরম্যান্স নেকেড</h1>
<p class="summary-text">কেটিএম ডিউক 200 একটি খেলাধুলামূলক 200cc নেকেড বাইক যা চিত্তাকর্ষক পারফরম্যান্স প্রদান করে। ৳5,95,000 টাকায় মূল্যায়িত, এটি 24.4 bhp শক্তি, আক্রমণাত্মক ডিজাইন এবং লাইটওয়েট ফ্রেম বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">কেটিএম ডিউক 200 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>200cc খেলাধুলামূলক:</strong> পারফরম্যান্স ফোকাসড</li>
<li class="highlight-item"><strong>24.4 bhp শক্তি:</strong> রোমাঞ্চকর পারফরম্যান্স</li>
<li class="highlight-item"><strong>আক্রমণাত্মক ডিজাইন:</strong> আধুনিক স্টাইল</li>
<li class="highlight-item"><strong>50+ km/l দক্ষতা:</strong> ভাল অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">কেটিএম ডিউক 200 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>খেলাধুলামূলক পারফরম্যান্স:</strong> 24.4 bhp রোমাঞ্চকর</li>
<li class="pro-item"><strong>আক্রমণাত্মক ডিজাইন:</strong> আধুনিক স্টাইলিং</li>
<li class="pro-item"><strong>লাইটওয়েট ফ্রেম:</strong> প্রতিক্রিয়াশীল হ্যান্ডলিং</li>
<li class="pro-item"><strong>ভাল দক্ষতা:</strong> 50+ km/l শালীন</li>
<li class="pro-item"><strong>কেটিএম গুণমান:</strong> পারফরম্যান্স প্রকৌশল</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">কেটিএম ডিউক 200 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳5,95,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>সীমিত সেবা:</strong> কেটিএম ডিলার শুধুমাত্র</li>
<li class="con-item"><strong>কঠোর সাসপেনশন:</strong> খেলাধুলামূলক ভিত্তিক</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: কেটিএম ডিউক 200 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳5,95,000 টাকায়, কেটিএম ডিউক 200 পারফরম্যান্স উত্সাহীদের জন্য আদর্শ যারা রোমাঞ্চকর 24.4 bhp শক্তি, আক্রমণাত্মক স্টাইলিং এবং প্রতিক্রিয়াশীল হ্যান্ডলিং চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পারফরম্যান্স নেকেড উত্সাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created KTM Duke 200\n")
	return nil
}

func (s *KTMDuke200Batch24) GetName() string {
	return "KTMDuke200Batch24"
}
