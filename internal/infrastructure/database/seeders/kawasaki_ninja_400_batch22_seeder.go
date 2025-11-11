package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type KawasakiNinja400Batch22 struct {
	BaseSeeder
}

func NewKawasakiNinja400Batch22Seeder() *KawasakiNinja400Batch22 {
	return &KawasakiNinja400Batch22{BaseSeeder: BaseSeeder{name: "Kawasaki Ninja 400 Batch22 Review"}}
}

func (s *KawasakiNinja400Batch22) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Kawasaki Ninja 400").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("kawasaki ninja 400 product not found")
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
<h1 class="review-main-title">Kawasaki Ninja 400 Review Bangladesh 2024 - Japan's Premium Entry Sport Fighter</h1>
<p class="summary-text">The Kawasaki Ninja 400 is Japan's premium 400cc entry sport fighter representing ultimate beginner performance. Priced around ৳7,50,000, it delivers 45 bhp exhilarating power, aggressive sport styling, excellent handling, good 50+ km/l efficiency, and outstanding value for entry sport enthusiasts seeking genuine sport character daily.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Kawasaki Ninja 400 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>400cc Parallel Twin:</strong> Entry sport power</li>
<li class="highlight-item"><strong>45 bhp Power:</strong> Exhilarating responsive</li>
<li class="highlight-item"><strong>Aggressive Styling:</strong> Genuine sport fighter</li>
<li class="highlight-item"><strong>50+ km/l Efficiency:</strong> Good economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Kawasaki Ninja 400 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Exhilarating 45 bhp:</strong> Entry sport power</li>
<li class="pro-item"><strong>Aggressive Sport Styling:</strong> Genuine fighter</li>
<li class="pro-item"><strong>Excellent Handling:</strong> Responsive sport</li>
<li class="pro-item"><strong>Kawasaki Quality:</strong> Premium engineering</li>
<li class="pro-item"><strong>Good Efficiency:</strong> 50+ km/l practical</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Kawasaki Ninja 400 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳7,50,000 upscale</li>
<li class="con-item"><strong>Limited Service:</strong> Few Kawasaki dealers</li>
<li class="con-item"><strong>Sport Focused:</strong> Limited comfort</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Kawasaki Ninja 400 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳7,50,000 - Premium sport</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>400cc parallel twin</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>45 bhp exhilarating</span></p>
<p class="value-item"><strong>Torque:</strong> <span>38 nm responsive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>168kg sport-tuned</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>50+ km/l good</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Kawasaki Ninja 400 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Performance:</strong> <span>4.8</span> - 45 bhp exhilarating</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.8</span> - Excellent responsive</li>
<li class="rating-item"><strong>Style:</strong> <span>4.7</span> - Genuine sport fighter</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.8</span> - Kawasaki premium</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳7,50,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Kawasaki Ninja 400?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳7,50,000, the Kawasaki Ninja 400 is perfect for entry sport enthusiasts seeking Japan's premium sport fighter, exhilarating 45 bhp power, aggressive authentic styling, excellent handling, and genuine sport character.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For entry sport enthusiasts</span></p>
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
<h1 class="review-main-title">কাওয়াসাকি নিনজা 400 রিভিউ বাংলাদেশ 2024 - জাপানের প্রিমিয়াম প্রবেশ স্পোর্ট ফাইটার</h1>
<p class="summary-text">কাওয়াসাকি নিনজা 400 জাপানের প্রিমিয়াম 400cc প্রবেশ স্পোর্ট ফাইটার যা চূড়ান্ত নতুন কর্মক্ষমতা প্রতিনিধিত্ব করে। ৳7,50,000 টাকায় মূল্যায়িত, এটি 45 bhp উত্তেজনাপূর্ণ শক্তি এবং আক্রমণাত্মক খেলাধুলা স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">কাওয়াসাকি নিনজা 400 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>400cc সমান্তরাল টুইন:</strong> প্রবেশ খেলাধুলা শক্তি</li>
<li class="highlight-item"><strong>45 bhp শক্তি:</strong> উত্তেজনাপূর্ণ প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>আক্রমণাত্মক স্টাইলিং:</strong> খাঁটি খেলাধুলা ফাইটার</li>
<li class="highlight-item"><strong>50+ km/l দক্ষতা:</strong> ভালো অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">কাওয়াসাকি নিনজা 400 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>উত্তেজনাপূর্ণ 45 bhp:</strong> প্রবেশ খেলাধুলা শক্তি</li>
<li class="pro-item"><strong>আক্রমণাত্মক খেলাধুলা স্টাইলিং:</strong> খাঁটি ফাইটার</li>
<li class="pro-item"><strong>চমৎকার হ্যান্ডলিং:</strong> প্রতিক্রিয়াশীল খেলাধুলা</li>
<li class="pro-item"><strong>কাওয়াসাকি গুণমান:</strong> প্রিমিয়াম প্রকৌশল</li>
<li class="pro-item"><strong>ভালো দক্ষতা:</strong> 50+ km/l ব্যবহারিক</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">কাওয়াসাকি নিনজা 400 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳7,50,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>সীমিত সেবা:</strong> কয়েক কাওয়াসাকি ডিলার</li>
<li class="con-item"><strong>খেলাধুলা ফোকাসড:</strong> সীমিত আরাম</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: কাওয়াসাকি নিনজা 400 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳7,50,000 টাকায়, কাওয়াসাকি নিনজা 400 প্রবেশ খেলাধুলা উৎসাহীদের জন্য নিখুঁত যারা জাপানের প্রিমিয়াম খেলাধুলা ফাইটার, উত্তেজনাপূর্ণ 45 bhp শক্তি এবং খাঁটি খেলাধুলা চরিত্র চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রবেশ খেলাধুলা উৎসাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Kawasaki Ninja 400\n")
	return nil
}

func (s *KawasakiNinja400Batch22) GetName() string {
	return "KawasakiNinja400Batch22"
}
