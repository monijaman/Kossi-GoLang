package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type PiaggioLibertoBatch22 struct {
	BaseSeeder
}

func NewPiaggioLibertoBatch22Seeder() *PiaggioLibertoBatch22 {
	return &PiaggioLibertoBatch22{BaseSeeder: BaseSeeder{name: "Piaggio Liberto Batch22 Review"}}
}

func (s *PiaggioLibertoBatch22) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Piaggio Liberto").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("piaggio liberto product not found")
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
<h1 class="review-main-title">Piaggio Liberto Review Bangladesh 2024 - Italy's Premium Classic Commuter Scooter</h1>
<p class="summary-text">The Piaggio Liberto is Italy's premium 125cc classic commuter scooter representing timeless European character. Priced around ৳5,50,000, it delivers 10 bhp steady power, classic Italian styling, comfortable riding, excellent 55+ km/l efficiency, and outstanding value for premium commuters seeking European classic character daily.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Piaggio Liberto Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc Air-Cooled:</strong> Classic European</li>
<li class="highlight-item"><strong>10 bhp Power:</strong> Steady reliable</li>
<li class="highlight-item"><strong>Classic Styling:</strong> Italian timeless</li>
<li class="highlight-item"><strong>55+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Piaggio Liberto Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Classic Italian Design:</strong> Timeless character</li>
<li class="pro-item"><strong>Comfortable Riding:</strong> Premium positioning</li>
<li class="pro-item"><strong>Reliable Power:</strong> Steady 10 bhp</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 55+ km/l practical</li>
<li class="pro-item"><strong>Premium Build:</strong> Italian engineering</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Piaggio Liberto Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳5,50,000 upscale</li>
<li class="con-item"><strong>Modern Performance:</strong> Not sporty</li>
<li class="con-item"><strong>Service Network:</strong> Limited dealers</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Piaggio Liberto Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳5,50,000 - Premium classic</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>125cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>10 bhp steady</span></p>
<p class="value-item"><strong>Torque:</strong> <span>10.2 nm reliable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>128kg scooter</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>55+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Piaggio Liberto Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Classic timeless</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.8</span> - Premium positioning</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 55+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.8</span> - Italian engineering</li>
<li class="rating-item"><strong>Value:</strong> <span>4.6</span> - ৳5,50,000 premium</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Piaggio Liberto?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳5,50,000, the Piaggio Liberto is perfect for premium commuters seeking Italy's classic timeless design, comfortable riding position, excellent 55+ km/l efficiency, and European character for daily sophisticated commuting.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium classic commuters</span></p>
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
<h1 class="review-main-title">পিয়াজিও লিবার্টো রিভিউ বাংলাদেশ 2024 - ইতালির প্রিমিয়াম ক্লাসিক কমিউটার স্কুটার</h1>
<p class="summary-text">পিয়াজিও লিবার্টো ইতালির প্রিমিয়াম 125cc ক্লাসিক কমিউটার স্কুটার যা নিরবধি ইউরোপীয় চরিত্র প্রতিনিধিত্ব করে। ৳5,50,000 টাকায় মূল্যায়িত, এটি 10 bhp স্থিতিশীল শক্তি এবং ক্লাসিক ইতালীয় স্টাইলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">পিয়াজিও লিবার্টো মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>125cc এয়ার-কুল্ড:</strong> ক্লাসিক ইউরোপীয়</li>
<li class="highlight-item"><strong>10 bhp শক্তি:</strong> স্থিতিশীল নির্ভরযোগ্য</li>
<li class="highlight-item"><strong>ক্লাসিক স্টাইলিং:</strong> ইতালীয় নিরবধি</li>
<li class="highlight-item"><strong>55+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">পিয়াজিও লিবার্টো সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>ক্লাসিক ইতালীয় ডিজাইন:</strong> নিরবধি চরিত্র</li>
<li class="pro-item"><strong>আরামদায়ক চালনা:</strong> প্রিমিয়াম অবস্থান</li>
<li class="pro-item"><strong>নির্ভরযোগ্য শক্তি:</strong> স্থিতিশীল 10 bhp</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 55+ km/l ব্যবহারিক</li>
<li class="pro-item"><strong>প্রিমিয়াম বিল্ড:</strong> ইতালীয় প্রকৌশল</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">পিয়াজিও লিবার্টো অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳5,50,000 উচ্চ সম্পদ</li>
<li class="con-item"><strong>আধুনিক কর্মক্ষমতা:</strong> খেলাধুলাপূর্ণ নয়</li>
<li class="con-item"><strong>সেবা নেটওয়ার্ক:</strong> সীমিত ডিলার</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: পিয়াজিও লিবার্টো কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳5,50,000 টাকায়, পিয়াজিও লিবার্টো প্রিমিয়াম কমিউটারদের জন্য নিখুঁত যারা ইতালির ক্লাসিক নিরবধি ডিজাইন, আরামদায়ক চালনা এবং চমৎকার 55+ km/l দক্ষতা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম ক্লাসিক কমিউটারদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Piaggio Liberto\n")
	return nil
}

func (s *PiaggioLibertoBatch22) GetName() string {
	return "PiaggioLibertoBatch22"
}
