package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsarAS250Review handles seeding Bajaj Pulsar AS250 product review and translation
type BajajPulsarAS250Review struct {
	BaseSeeder
}

// NewBajajPulsarAS250ReviewSeeder creates a new BajajPulsarAS250Review
func NewBajajPulsarAS250ReviewSeeder() *BajajPulsarAS250Review {
	return &BajajPulsarAS250Review{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar AS250 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsarAS250Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar AS250").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj pulsar as250 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bajaj pulsar as250 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar AS250 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar AS250 Review Bangladesh 2024 - Adventure Tourer</h1>
<p class="summary-text">The Bajaj Pulsar AS250 is a versatile 250cc adventure-sport bike combining on-road sportiness with light touring capability. Priced around ৳3,10,000, it delivers balanced performance, upright seating, and practical functionality for weekend getaways and daily commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar AS250 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">250cc Adventure Engine:</strong> <span class="highlight-value">Dual-purpose performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Upright Seating:</strong> <span class="highlight-value">Comfortable ergonomics</span></li>
<li class="highlight-item"><strong class="highlight-label">Touring Focused:</strong> <span class="highlight-value">Weekend friendly</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">38-42 km/l decent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar AS250 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Versatile Engine:</strong> <span class="pro-description">250cc adventure-sport</span></li>
<li class="pro-item"><strong class="pro-label">Upright Ergonomics:</strong> <span class="pro-description">Comfortable long rides</span></li>
<li class="pro-item"><strong class="pro-label">Touring Capability:</strong> <span class="pro-description">Weekend ready</span></li>
<li class="pro-item"><strong class="pro-label">Good Ground Clearance:</strong> <span class="pro-description">Light offroad capable</span></li>
<li class="pro-item"><strong class="pro-label">Bajaj Reliability:</strong> <span class="pro-description">Proven performance</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar AS250 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Weight:</strong> <span class="con-description">Heavier than street bikes</span></li>
<li class="con-item"><strong class="con-label">Handling:</strong> <span class="con-description">Not as nimble</span></li>
<li class="con-item"><strong class="con-label">Fuel Economy:</strong> <span class="con-description">38-42 km/l average</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar AS250 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,10,000 - Adventure tourer</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳6-7 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">38-42 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar AS250 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Performance:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- 250cc versatile</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Upright touring</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Adventure oriented</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Bajaj proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good versatility</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar AS250?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,10,000, the Bajaj Pulsar AS250 is ideal for riders seeking a versatile adventure-sport bike that combines daily commuting practicality with weekend touring comfort and light offroad capability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For adventure-touring enthusiasts</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.4,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating bajaj pulsar as250 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar AS250 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ Pulsar AS250 রিভিউ বাংলাদেশ 2024 - অ্যাডভেঞ্চার ট্যুরার</h1>
<p class="summary-text">বাজাজ Pulsar AS250 একটি বহুমুখী 250cc অ্যাডভেঞ্চার-ক্রীড়া বাইক যা অন-রোড ক্রীড়া ক্ষমতা সহ হালকা ট্যুরিং ক্ষমতা একত্রিত করে। ৳3,10,000 টাকায় মূল্যায়িত, এটি সুষম কর্মক্ষমতা, সোজা আসন এবং সপ্তাহান্তের গেটঅ্যাওয়ে এবং দৈনন্দিন যাতায়াতের জন্য ব্যবহারিক কার্যকারিতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ Pulsar AS250 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">250cc অ্যাডভেঞ্চার ইঞ্জিন:</strong> <span class="highlight-value">দ্বৈত-উদ্দেশ্য কর্মক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">সোজা আসন:</strong> <span class="highlight-value">আরামদায়ক এরগনমিক্স</span></li>
<li class="highlight-item"><strong class="highlight-label">ট্যুরিং ফোকাসড:</strong> <span class="highlight-value">সপ্তাহান্তে বান্ধব</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">38-42 km/l মানসম্পন্ন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ Pulsar AS250 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">বহুমুখী ইঞ্জিন:</strong> <span class="pro-description">250cc অ্যাডভেঞ্চার-ক্রীড়া</span></li>
<li class="pro-item"><strong class="pro-label">সোজা এরগনমিক্স:</strong> <span class="pro-description">আরামদায়ক দীর্ঘ যাত্রা</span></li>
<li class="pro-item"><strong class="pro-label">ট্যুরিং ক্ষমতা:</strong> <span class="pro-description">সপ্তাহান্ত প্রস্তুত</span></li>
<li class="pro-item"><strong class="pro-label">ভাল গ্রাউন্ড ক্লিয়ারেন্স:</strong> <span class="pro-description">হালকা অফরোড সক্ষম</span></li>
<li class="pro-item"><strong class="pro-label">বাজাজ নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত কর্মক্ষমতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ Pulsar AS250 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">ওজন:</strong> <span class="con-description">স্ট্রিট বাইকের চেয়ে ভারী</span></li>
<li class="con-item"><strong class="con-label">হ্যান্ডলিং:</strong> <span class="con-description">তত চটপটে নয়</span></li>
<li class="con-item"><strong class="con-label">জ্বালানী অর্থনীতি:</strong> <span class="con-description">38-42 km/l গড়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ Pulsar AS250 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,10,000 টাকায়, বাজাজ Pulsar AS250 রাইডারদের জন্য আদর্শ যারা একটি বহুমুখী অ্যাডভেঞ্চার-ক্রীড়া বাইক খুঁজছেন যা দৈনন্দিন যাতায়াত ব্যবহারিকতা সহ সপ্তাহান্তের ট্যুরিং আরাম এবং হালকা অফরোড ক্ষমতা একত্রিত করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- অ্যাডভেঞ্চার-ট্যুরিং উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar AS250 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar AS250\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPulsarAS250Review) GetName() string {
	return "BajajPulsarAS250Review"
}
