package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type VespaSprintBatch24 struct {
	BaseSeeder
}

func NewVespaSprintBatch24Seeder() *VespaSprintBatch24 {
	return &VespaSprintBatch24{BaseSeeder: BaseSeeder{name: "Vespa Sprint Batch24 Review"}}
}

func (s *VespaSprintBatch24) Seed(db *gorm.DB) error {
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
<h1 class="review-main-title">Vespa Sprint Review Bangladesh 2024 - Iconic Italian Scooter</h1>
<p class="summary-text">The Vespa Sprint is the iconic Italian scooter offering timeless design and authentic Vespa character. Priced around ৳9,85,000, it delivers 12 bhp power, iconic round design, vintage charm, excellent 58+ km/l efficiency, and outstanding prestige for riders seeking legendary scooter character.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Vespa Sprint Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>Iconic Design:</strong> Legendary Italian</li>
<li class="highlight-item"><strong>12 bhp Power:</strong> Responsive capable</li>
<li class="highlight-item"><strong>Vintage Charm:</strong> Timeless appeal</li>
<li class="highlight-item"><strong>58+ km/l Efficiency:</strong> Excellent economy</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Vespa Sprint Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Iconic Design:</strong> Legendary status</li>
<li class="pro-item"><strong>Vintage Character:</strong> Timeless appeal</li>
<li class="pro-item"><strong>Responsive Power:</strong> 12 bhp capable</li>
<li class="pro-item"><strong>Excellent Efficiency:</strong> 58+ km/l superior</li>
<li class="pro-item"><strong>Prestige Factor:</strong> Iconic heritage</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Vespa Sprint Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Premium Price:</strong> ৳9,85,000 luxury</li>
<li class="con-item"><strong>Limited Parts:</strong> Import dependent</li>
<li class="con-item"><strong>Service Availability:</strong> Specialist required</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Vespa Sprint Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳9,85,000 - Premium iconic</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>150cc scooter single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>12 bhp responsive</span></p>
<p class="value-item"><strong>Torque:</strong> <span>11.6 nm capable</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>144kg iconic</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>58+ km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Vespa Sprint Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.9</span> - Iconic legendary</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Vintage comfortable</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.7</span> - 12 bhp capable</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.7</span> - 58+ km/l excellent</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.8</span> - Italian heritage</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Vespa Sprint?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳9,85,000, the Vespa Sprint is ideal for riders seeking iconic Italian character, legendary status, vintage charm, excellent 58+ km/l efficiency, and prestige driving daily.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For iconic scooter enthusiasts</span></p>
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
<h1 class="review-main-title">ভেসপা স্প্রিন্ট রিভিউ বাংলাদেশ 2024 - আইকনিক ইতালীয় স্কুটার</h1>
<p class="summary-text">ভেসপা স্প্রিন্ট আইকনিক ইতালীয় স্কুটার যা চিরন্তন ডিজাইন এবং খাঁটি ভেসপা চরিত্র প্রদান করে। ৳9,85,000 টাকায় মূল্যায়িত, এটি 12 bhp শক্তি, আইকনিক গোলাকার ডিজাইন এবং ভিন্টেজ আকর্ষণ প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ভেসপা স্প্রিন্ট মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>আইকনিক ডিজাইন:</strong> কিংবদন্তি ইতালীয়</li>
<li class="highlight-item"><strong>12 bhp শক্তি:</strong> প্রতিক্রিয়াশীল সক্ষম</li>
<li class="highlight-item"><strong>ভিন্টেজ আকর্ষণ:</strong> চিরন্তন আবেদন</li>
<li class="highlight-item"><strong>58+ km/l দক্ষতা:</strong> চমৎকার অর্থনীতি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ভেসপা স্প্রিন্ট সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>আইকনিক ডিজাইন:</strong> কিংবদন্তি মর্যাদা</li>
<li class="pro-item"><strong>ভিন্টেজ চরিত্র:</strong> চিরন্তন আবেদন</li>
<li class="pro-item"><strong>প্রতিক্রিয়াশীল শক্তি:</strong> 12 bhp সক্ষম</li>
<li class="pro-item"><strong>চমৎকার দক্ষতা:</strong> 58+ km/l উচ্চতর</li>
<li class="pro-item"><strong>মর্যাদা ফ্যাক্টর:</strong> আইকনিক ঐতিহ্য</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ভেসপা স্প্রিন্ট অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳9,85,000 বিলাসবহুল</li>
<li class="con-item"><strong>সীমিত যন্ত্রাংশ:</strong> আমদানি নির্ভর</li>
<li class="con-item"><strong>সেবা উপলব্ধতা:</strong> বিশেষজ্ঞ প্রয়োজন</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ভেসপা স্প্রিন্ট কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳9,85,000 টাকায়, ভেসপা স্প্রিন্ট চালকদের জন্য আদর্শ যারা আইকনিক ইতালীয় চরিত্র, কিংবদন্তি মর্যাদা এবং ভিন্টেজ আকর্ষণ চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আইকনিক স্কুটার উত্সাহীদের জন্য</span></p>
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

func (s *VespaSprintBatch24) GetName() string {
	return "VespaSprintBatch24"
}
