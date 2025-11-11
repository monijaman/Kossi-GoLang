package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type KawasakiNinja650Batch25 struct {
	BaseSeeder
}

func NewKawasakiNinja650Batch25Seeder() *KawasakiNinja650Batch25 {
	return &KawasakiNinja650Batch25{BaseSeeder: BaseSeeder{name: "Kawasaki Ninja 650 Batch25 Review"}}
}

func (s *KawasakiNinja650Batch25) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Kawasaki Ninja 650").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("kawasaki ninja 650 product not found")
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
<h1 class="review-main-title">Kawasaki Ninja 650 Review Bangladesh 2024 - Mid-Range Sport Excellence</h1>
<p class="summary-text">The Kawasaki Ninja 650 is a mid-range 650cc sport bike delivering impressive performance. Priced around ৳10,25,000, it features 68 bhp power, aggressive styling, responsive handling, excellent 45+ km/l efficiency, and outstanding appeal for mid-range sport seekers.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Kawasaki Ninja 650 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>650cc Sport:</strong> Mid-range powerful</li>
<li class="highlight-item"><strong>68 bhp Power:</strong> Thrilling capable</li>
<li class="highlight-item"><strong>Aggressive Design:</strong> Sport styled</li>
<li class="highlight-item"><strong>45+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Kawasaki Ninja 650 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Impressive Power:</strong> 68 bhp thrilling</li>
<li class="pro-item"><strong>Aggressive Design:</strong> Sport styled</li>
<li class="pro-item"><strong>Responsive Handling:</strong> Precise control</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 45+ km/l decent</li>
<li class="pro-item"><strong>Kawasaki Heritage:</strong> Performance proven</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Kawasaki Ninja 650 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳10,25,000 upscale</li>
<li class="con-item"><strong>Limited Service:</strong> Specialized dealers</li>
<li class="con-item"><strong>Harder Suspension:</strong> Sport configured</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Kawasaki Ninja 650 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳10,25,000 - Premium mid-range</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>650cc parallel twin</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>68 bhp thrilling</span></p>
<p class="value-item"><strong>Torque:</strong> <span>63.9 nm powerful</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>187kg sport</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>45+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Kawasaki Ninja 650 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.9</span> - Aggressive sport</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.6</span> - Sport configured</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.9</span> - 68 bhp thrilling</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.5</span> - 45+ km/l good</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.8</span> - Kawasaki proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Kawasaki Ninja 650?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳10,25,000, the Kawasaki Ninja 650 is ideal for mid-range sport enthusiasts seeking thrilling 68 bhp power, aggressive styling, precise handling, and good 45+ km/l efficiency.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For mid-range sport seekers</span></p>
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
<h1 class="review-main-title">কাওয়াসাকি নিনজা 650 রিভিউ বাংলাদেশ 2024 - মধ্য-পরিসর খেলাধুলা শ্রেষ্ঠত্ব</h1>
<p class="summary-text">কাওয়াসাকি নিনজা 650 একটি মধ্য-পরিসীমা 650cc খেলাধুলা বাইক যা চিত্তাকর্ষক পারফরম্যান্স প্রদান করে। ৳10,25,000 টাকায় মূল্যায়িত, এটি 68 bhp শক্তি এবং আক্রমণাত্মক স্টাইলিং বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">কাওয়াসাকি নিনজা 650 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>650cc খেলা:</strong> মধ্য-পরিসর শক্তিশালী</li>
<li class="highlight-item"><strong>68 bhp শক্তি:</strong> রোমাঞ্চকর সক্ষম</li>
<li class="highlight-item"><strong>আক্রমণাত্মক ডিজাইন:</strong> খেলাধুলা স্টাইল</li>
<li class="highlight-item"><strong>45+ km/l দক্ষতা:</strong> ভাল অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">কাওয়াসাকি নিনজা 650 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>চিত্তাকর্ষক শক্তি:</strong> 68 bhp রোমাঞ্চকর</li>
<li class="pro-item"><strong>আক্রমণাত্মক ডিজাইন:</strong> খেলাধুলা স্টাইল</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল হ্যান্ডলিং:</strong> নির্ভুল নিয়ন্ত্রণ</li>
<li class="pro-item"><strong>ভাল দক্ষতা:</strong> 45+ km/l শালীন</li>
<li class="pro-item"><strong>কাওয়াসাকি ঐতিহ্য:</strong> পারফরম্যান্স প্রমাণিত</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">কাওয়াসাকি নিনজা 650 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳10,25,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>সীমিত সেবা:</strong> বিশেষায়িত ডিলার</li>
<li class="con-item"><strong>কঠোর সাসপেনশন:</strong> খেলাধুলামূলক কনফিগার</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: কাওয়াসাকি নিনজা 650 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳10,25,000 টাকায়, কাওয়াসাকি নিনজা 650 মধ্য-পরিসীমা খেলাধুলা উত্সাহীদের জন্য আদর্শ যারা রোমাঞ্চকর 68 bhp শক্তি এবং নির্ভুল হ্যান্ডলিং চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- মধ্য-পরিসর খেলাধুলা খোঁজার জন্য</span></p>
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

	fmt.Printf("   ✅ Created Kawasaki Ninja 650\n")
	return nil
}

func (s *KawasakiNinja650Batch25) GetName() string {
	return "KawasakiNinja650Batch25"
}
