package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type VespaSprintReviewBatch20 struct {
	BaseSeeder
}

func NewVespaSprintReviewBatch20Seeder() *VespaSprintReviewBatch20 {
	return &VespaSprintReviewBatch20{BaseSeeder: BaseSeeder{name: "Vespa Sprint Batch20 Review"}}
}

func (s *VespaSprintReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Vespa Sprint").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("vespa sprint product not found")
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
<h1 class="review-main-title">Vespa Sprint Review Bangladesh 2024 - Italian Retro Design Legend</h1>
<p class="summary-text">The Vespa Sprint is a 150cc air-cooled Italian style icon representing timeless scooter heritage and retro design perfection. Priced around ৳5,00,000, it delivers iconic design, comfortable seating, excellent fuel efficiency, authentic Italian heritage, and vintage soul for style enthusiasts seeking iconic retro elegance.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Vespa Sprint Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc Air-Cooled:</strong> Icon scooter</li>
<li class="highlight-item"><strong>Iconic Italian Design:</strong> Timeless heritage</li>
<li class="highlight-item"><strong>40 km/l Efficiency:</strong> Excellent economy</li>
<li class="highlight-item"><strong>Retro Heritage:</strong> ৳5,00,000 legend</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Vespa Sprint Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Iconic Design:</strong> Legendary Italian style</li>
<li class="pro-item"><strong>Heritage Character:</strong> Timeless Vespa soul</li>
<li class="pro-item"><strong>Comfortable Seating:</strong> Spacious classic design</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 40 km/l economy</li>
<li class="pro-item"><strong>Retro Charm:</strong> Vintage authentic appeal</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Vespa Sprint Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Limited Power:</strong> 11 bhp modest</li>
<li class="con-item"><strong>Moderate Weight:</strong> 150kg substantial</li>
<li class="con-item"><strong>Premium Price:</strong> ৳5,00,000 retro heritage</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Vespa Sprint Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳5,00,000 - Premium retro icon</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc air-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>11 bhp modest</span></p>
<p class="value-item"><strong>Torque:</strong> <span>12 nm character</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>150kg heritage</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>40 km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Vespa Sprint Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.9</span> - Iconic legendary</li>
<li class="rating-item"><strong>Heritage:</strong> <span>4.9</span> - Vespa soul</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Spacious seating</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.8</span> - 40 km/l excellent</li>
<li class="rating-item"><strong>Value:</strong> <span>4.7</span> - ৳5,00,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Vespa Sprint?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳5,00,000, the Vespa Sprint is the iconic choice for style enthusiasts seeking authentic Italian design heritage, timeless retro elegance, comfortable classic seating, excellent 40 km/l efficiency, and vintage Vespa soul for lifestyle-focused commuting.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For retro style enthusiasts</span></p>
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
<h1 class="review-main-title">ভেসপা স্প্রিন্ট রিভিউ বাংলাদেশ 2024 - ইতালিয়ান রেট্রো ডিজাইন কিংবদন্তী</h1>
<p class="summary-text">ভেসপা স্প্রিন্ট একটি 150cc এয়ার-কুল্ড ইতালিয়ান স্টাইল আইকন যা কালজয়ী স্কুটার ঐতিহ্য এবং রেট্রো ডিজাইন নিখুঁততা প্রতিনিধিত্ব করে। ৳5,00,000 টাকায় মূল্যায়িত, এটি আইকনিক ডিজাইন এবং খাঁটি ইতালিয়ান ঐতিহ্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ভেসপা স্প্রিন্ট মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>150cc এয়ার-কুল্ড:</strong> আইকন স্কুটার</li>
<li class="highlight-item"><strong>আইকনিক ইতালিয়ান ডিজাইন:</strong> কালজয়ী ঐতিহ্য</li>
<li class="highlight-item"><strong>40 km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
<li class="highlight-item"><strong>রেট্রো ঐতিহ্য:</strong> ৳5,00,000 কিংবদন্তী</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ভেসপা স্প্রিন্ট সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আইকনিক ডিজাইন:</strong> কিংবদন্তী ইতালিয়ান স্টাইল</li>
<li class="pro-item"><strong>ঐতিহ্য চরিত্র:</strong> কালজয়ী ভেসপা আত্মা</li>
<li class="pro-item"><strong>আরামদায়ক আসন:</strong> প্রশস্ত ক্লাসিক ডিজাইন</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 40 km/l অর্থনীতি</li>
<li class="pro-item"><strong>রেট্রো আকর্ষণ:</strong> ভিনটেজ খাঁটি আবেদন</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ভেসপা স্প্রিন্ট অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>সীমিত শক্তি:</strong> 11 bhp বিনম্র</li>
<li class="con-item"><strong>মধ্যম ওজন:</strong> 150kg উল্লেখযোগ্য</li>
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳5,00,000 রেট্রো ঐতিহ্য</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ভেসপা স্প্রিন্ট কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳5,00,000 টাকায়, ভেসপা স্প্রিন্ট স্টাইল উত্সাহীদের জন্য আইকনিক পছন্দ যারা খাঁটি ইতালিয়ান ডিজাইন ঐতিহ্য, কালজয়ী রেট্রো মার্জিত এবং ভিনটেজ ভেসপা আত্মা চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- রেট্রো স্টাইল উত্সাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created Vespa Sprint\n")
	return nil
}

func (s *VespaSprintReviewBatch20) GetName() string {
	return "VespaSprintReviewBatch20"
}
