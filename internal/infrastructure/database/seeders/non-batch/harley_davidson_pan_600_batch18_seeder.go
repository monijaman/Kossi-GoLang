package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HarleyDavidsonPan600ReviewBatch18 handles seeding Harley-Davidson Pan 600 product review and translation
type HarleyDavidsonPan600ReviewBatch18 struct {
	BaseSeeder
}

// NewHarleyDavidsonPan600ReviewBatch18Seeder creates a new HarleyDavidsonPan600ReviewBatch18
func NewHarleyDavidsonPan600ReviewBatch18Seeder() *HarleyDavidsonPan600ReviewBatch18 {
	return &HarleyDavidsonPan600ReviewBatch18{BaseSeeder: BaseSeeder{name: "Harley-Davidson Pan 600 Batch18 Review"}}
}

// Seed implements the Seeder interface
func (s *HarleyDavidsonPan600ReviewBatch18) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Harley-Davidson Pan 600").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("harley-davidson pan 600 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding harley-davidson pan 600 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Harley-Davidson Pan 600 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Harley-Davidson Pan 600 Review Bangladesh 2024 - Entry-Level American Cruiser Heritage</h1>
<p class="summary-text">The Harley-Davidson Pan 600 is an 600cc air-cooled v-twin entry cruiser bringing legendary H-D heritage to accessible price points. Priced around ৳11,50,000, it delivers classic American cruiser DNA, authentic styling, comfortable seating, and the iconic rumbling powerplant for beginners entering the cruiser lifestyle.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Harley-Davidson Pan 600 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">600cc V-Twin:</strong> <span class="highlight-value">Entry-level cruiser power</span></li>
<li class="highlight-item"><strong class="highlight-label">American Heritage:</strong> <span class="highlight-value">Iconic H-D styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Air-Cooled Engine:</strong> <span class="highlight-value">Classic rumble character</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfort Focus:</strong> <span class="highlight-value">Low seat cruising</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Harley-Davidson Pan 600 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Authentic Style:</strong> <span class="pro-description">Genuine American cruiser</span></li>
<li class="pro-item"><strong class="pro-label">Heritage Appeal:</strong> <span class="pro-description">120+ year legacy</span></li>
<li class="pro-item"><strong class="pro-label">Accessible Price:</strong> <span class="pro-description">৳11,50,000 entry point</span></li>
<li class="pro-item"><strong class="pro-label">Comfortable Seat:</strong> <span class="pro-description">Relaxed riding position</span></li>
<li class="pro-item"><strong class="pro-label">Iconic Sound:</strong> <span class="pro-description">Classic V-twin rumble</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Harley-Davidson Pan 600 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Fuel Efficiency:</strong> <span class="con-description">17-20 km/l moderate</span></li>
<li class="con-item"><strong class="con-label">Handling Limitations:</strong> <span class="con-description">Cruiser ergonomics</span></li>
<li class="con-item"><strong class="con-label">Maintenance Cost:</strong> <span class="con-description">American brand service</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Harley-Davidson Pan 600 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳11,50,000 - Entry cruiser</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳12-15 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">17-20 km/l moderate</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">600cc air-cooled V-twin</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">46 bhp low-end torque</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">210kg comfortable cruiser</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Harley-Davidson Pan 600 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Cruiser Character:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Authentic American</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Relaxed seating</span></li>
<li class="rating-item"><strong class="rating-label">Heritage Appeal:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Iconic styling</span></li>
<li class="rating-item"><strong class="rating-label">Sound:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Signature rumble</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Entry cruiser access</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Harley-Davidson Pan 600?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳11,50,000, the Pan 600 is perfect for cruiser enthusiasts seeking authentic Harley-Davidson heritage, classic American styling, and iconic V-twin character in an entry-level price point with genuine cruiser comfort and lifestyle appeal.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For cruiser lifestyle seekers</span></p>
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
		return fmt.Errorf("error creating harley-davidson pan 600 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Harley-Davidson Pan 600 (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হার্লে-ডেভিডসন প্যান 600 রিভিউ বাংলাদেশ 2024 - এন্ট্রি-লেভেল আমেরিকান ক্রুজার হেরিটেজ</h1>
<p class="summary-text">হার্লে-ডেভিডসন প্যান 600 একটি 600cc এয়ার-কুল্ড ভি-টুইন এন্ট্রি ক্রুজার যা কিংবদন্তি এইচ-ডি হেরিটেজকে সাশ্রয়ী মূল্যে নিয়ে আসে। ৳11,50,000 টাকায় মূল্যায়িত, এটি ক্লাসিক আমেরিকান ক্রুজার DNA, খাঁটি স্টাইলিং এবং আইকনিক রাম্বলিং পাওয়ার প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হার্লে-ডেভিডসন প্যান 600 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">600cc ভি-টুইন:</strong> <span class="highlight-value">এন্ট্রি-লেভেল ক্রুজার শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">আমেরিকান হেরিটেজ:</strong> <span class="highlight-value">আইকনিক এইচ-ডি স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">এয়ার-কুল্ড ইঞ্জিন:</strong> <span class="highlight-value">ক্লাসিক রাম্বল চরিত্র</span></li>
<li class="highlight-item"><strong class="highlight-label">আরাম ফোকাস:</strong> <span class="highlight-value">লো সিট ক্রুজিং</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হার্লে-ডেভিডসন প্যান 600 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">খাঁটি স্টাইল:</strong> <span class="pro-description">জেনুইন আমেরিকান ক্রুজার</span></li>
<li class="pro-item"><strong class="pro-label">হেরিটেজ আবেদন:</strong> <span class="pro-description">120+ বছর লিগেসি</span></li>
<li class="pro-item"><strong class="pro-label">সাশ্রয়ী মূল্য:</strong> <span class="pro-description">৳11,50,000 এন্ট্রি পয়েন্ট</span></li>
<li class="pro-item"><strong class="pro-label">আরাম আসন:</strong> <span class="pro-description">শান্ত রাইডিং পজিশন</span></li>
<li class="pro-item"><strong class="pro-label">আইকনিক সাউন্ড:</strong> <span class="pro-description">ক্লাসিক ভি-টুইন রাম্বল</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হার্লে-ডেভিডসন প্যান 600 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">জ্বালানি দক্ষতা:</strong> <span class="con-description">17-20 km/l মধ্যম</span></li>
<li class="con-item"><strong class="con-label">হ্যান্ডলিং সীমাবদ্ধতা:</strong> <span class="con-description">ক্রুজার এরগোনমিক্স</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ খরচ:</strong> <span class="con-description">আমেরিকান ব্র্যান্ড সেবা</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হার্লে-ডেভিডসন প্যান 600 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳11,50,000 টাকায়, প্যান 600 ক্রুজার উত্সাহীদের জন্য নিখুঁত যারা খাঁটি হার্লে-ডেভিডসন হেরিটেজ এবং আইকনিক ভি-টুইন চরিত্র চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ক্রুজার লাইফস্টাইল সন্ধানকারীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Harley-Davidson Pan 600 already exists\n")
		return nil
	}

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating bangla translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Harley-Davidson Pan 600\n")

	return nil
}

// GetName returns the seeder name
func (s *HarleyDavidsonPan600ReviewBatch18) GetName() string {
	return "HarleyDavidsonPan600ReviewBatch18"
}
