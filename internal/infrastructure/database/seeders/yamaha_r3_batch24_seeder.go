package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type YamahaR3Batch24 struct {
	BaseSeeder
}

func NewYamahaR3Batch24Seeder() *YamahaR3Batch24 {
	return &YamahaR3Batch24{BaseSeeder: BaseSeeder{name: "Yamaha R3 Batch24 Review"}}
}

func (s *YamahaR3Batch24) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Yamaha R3").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("yamaha r3 product not found")
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
<h1 class="review-main-title">Yamaha R3 Review Bangladesh 2024 - Beginner Sports Bike</h1>
<p class="summary-text">The Yamaha R3 is a beginner-friendly 321cc sports bike combining power and accessibility. Priced around ৳7,25,000, it delivers 42 bhp power, sporty design, lightweight frame, precise handling, and impressive 50+ km/l efficiency for riders starting their sport bike journey.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Yamaha R3 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>321cc Sporty:</strong> Beginner accessible</li>
<li class="highlight-item"><strong>42 bhp Power:</strong> Thrilling performance</li>
<li class="highlight-item"><strong>Lightweight Frame:</strong> Precise handling</li>
<li class="highlight-item"><strong>50+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Yamaha R3 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Beginner Friendly:</strong> Accessible sports</li>
<li class="pro-item"><strong>Thrilling Power:</strong> 42 bhp exciting</li>
<li class="pro-item"><strong>Precise Handling:</strong> Responsive frame</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 50+ km/l decent</li>
<li class="pro-item"><strong>Yamaha Quality:</strong> Proven reliability</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Yamaha R3 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Higher Price:</strong> ৳7,25,000 upscale</li>
<li class="con-item"><strong>Limited Parts:</strong> Import availability</li>
<li class="con-item"><strong>Harder Suspension:</strong> Sports oriented</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Yamaha R3 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳7,25,000 - Sports accessible</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>321cc twin cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>42 bhp thrilling</span></p>
<p class="value-item"><strong>Torque:</strong> <span>29.6 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>169kg lightweight</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Yamaha R3 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Sporty modern</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.6</span> - Sports focused</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.8</span> - 42 bhp thrilling</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.6</span> - 50+ km/l good</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - Yamaha proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Yamaha R3?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳7,25,000, the Yamaha R3 is ideal for riders starting their sport bike journey, seeking thrilling 42 bhp power, precise handling, sporty design, and good 50+ km/l efficiency in an accessible package.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For beginner sport bike riders</span></p>
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
<h1 class="review-main-title">ইয়ামাহা R3 রিভিউ বাংলাদেশ 2024 - শুরুর খেলাধুলামূলক বাইক</h1>
<p class="summary-text">ইয়ামাহা R3 একটি শুরুর জন্য বান্ধব 321cc খেলাধুলামূলক বাইক যা শক্তি এবং অ্যাক্সেসযোগ্যতা একত্রিত করে। ৳7,25,000 টাকায় মূল্যায়িত, এটি 42 bhp শক্তি এবং খেলাধুলামূলক ডিজাইন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ইয়ামাহা R3 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>321cc খেলাধুলামূলক:</strong> শুরুর অ্যাক্সেসযোগ্য</li>
<li class="highlight-item"><strong>42 bhp শক্তি:</strong> রোমাঞ্চকর পারফরম্যান্স</li>
<li class="highlight-item"><strong>লাইটওয়েট ফ্রেম:</strong> নির্ভুল হ্যান্ডলিং</li>
<li class="highlight-item"><strong>50+ km/l দক্ষতা:</strong> ভাল অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ইয়ামাহা R3 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>শুরুর জন্য বান্ধব:</strong> অ্যাক্সেসযোগ্য খেলাধুলা</li>
<li class="pro-item"><strong>রোমাঞ্চকর শক্তি:</strong> 42 bhp রোমাঞ্চকর</li>
<li class="pro-item"><strong>নির্ভুল হ্যান্ডলিং:</strong> প্রতিক্রিয়াশীল ফ্রেম</li>
<li class="pro-item"><strong>ভাল দক্ষতা:</strong> 50+ km/l শালীন</li>
<li class="pro-item"><strong>ইয়ামাহা গুণমান:</strong> প্রমাণিত নির্ভরযোগ্যতা</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ইয়ামাহা R3 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>উচ্চতর মূল্য:</strong> ৳7,25,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>সীমিত যন্ত্রাংশ:</strong> আমদানি উপলব্ধতা</li>
<li class="con-item"><strong>কঠোর সাসপেনশন:</strong> খেলাধুলামূলক ভিত্তিক</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ইয়ামাহা R3 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳7,25,000 টাকায়, ইয়ামাহা R3 তাদের খেলাধুলামূলক বাইক যাত্রা শুরু করছেন এমন চালকদের জন্য আদর্শ যারা রোমাঞ্চকর 42 bhp শক্তি এবং নির্ভুল হ্যান্ডলিং চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- শুরুর খেলাধুলামূলক চালকদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Yamaha R3\n")
	return nil
}

func (s *YamahaR3Batch24) GetName() string {
	return "YamahaR3Batch24"
}
