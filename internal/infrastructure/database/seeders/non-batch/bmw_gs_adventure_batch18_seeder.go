package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BMWGSAdventureReviewBatch18 handles seeding BMW GS Adventure product review and translation
type BMWGSAdventureReviewBatch18 struct {
	BaseSeeder
}

// NewBMWGSAdventureReviewBatch18Seeder creates a new BMWGSAdventureReviewBatch18
func NewBMWGSAdventureReviewBatch18Seeder() *BMWGSAdventureReviewBatch18 {
	return &BMWGSAdventureReviewBatch18{BaseSeeder: BaseSeeder{name: "BMW GS Adventure Batch18 Review"}}
}

// Seed implements the Seeder interface
func (s *BMWGSAdventureReviewBatch18) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "BMW GS Adventure").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bmw gs adventure product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bmw gs adventure product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for BMW GS Adventure already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">BMW GS Adventure Review Bangladesh 2024 - Premium Adventure Touring Legend</h1>
<p class="summary-text">The BMW GS Adventure is a 1200cc premium adventure touring bike delivering German engineering excellence, legendary versatility, and unmatched capability for world-class expeditions. Priced around ৳28,50,000, it combines advanced technology, rugged durability, sophisticated suspension, and legendary reliability making it the ultimate adventure platform for serious explorers and long-distance touring enthusiasts.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">BMW GS Adventure Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">1200cc Boxer Twin:</strong> <span class="highlight-value">Legend torque delivery</span></li>
<li class="highlight-item"><strong class="highlight-label">Adventure Heritage:</strong> <span class="highlight-value">40+ year proven platform</span></li>
<li class="highlight-item"><strong class="highlight-label">Off-Road Capability:</strong> <span class="highlight-value">Extreme terrain ready</span></li>
<li class="highlight-item"><strong class="highlight-label">German Engineering:</strong> <span class="highlight-value">Precision reliability</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">BMW GS Adventure Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Legendary Platform:</strong> <span class="pro-description">40+ year heritage</span></li>
<li class="pro-item"><strong class="pro-label">Off-Road Beast:</strong> <span class="pro-description">Extreme terrain ready</span></li>
<li class="pro-item"><strong class="pro-label">Boxer Engine:</strong> <span class="pro-description">Torquey character</span></li>
<li class="pro-item"><strong class="pro-label">Comfort Touring:</strong> <span class="pro-description">Long-distance capable</span></li>
<li class="pro-item"><strong class="pro-label">Tech Packed:</strong> <span class="pro-description">Advanced systems</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">BMW GS Adventure Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Ultra Premium Price:</strong> <span class="con-description">৳28,50,000 extreme</span></li>
<li class="con-item"><strong class="con-label">High Maintenance:</strong> <span class="con-description">Very expensive upkeep</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">234kg challenging</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">BMW GS Adventure Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳28,50,000 - Ultra-premium adventure</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Very high - ৳20-28 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">18-22 km/l low</span></p>
<p class="value-item"><strong class="value-label">Engine Displacement:</strong> <span class="value-amount">1200cc air-cooled boxer</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">100+ bhp legendary torque</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">BMW GS Adventure Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Legendary boxer power</span></li>
<li class="rating-item"><strong class="rating-label">Off-Road Capability:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Extreme terrain master</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Long-distance luxury</span></li>
<li class="rating-item"><strong class="rating-label">Technology:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Advanced systems</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Ultra-premium cost</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy BMW GS Adventure?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳28,50,000, the GS Adventure is for serious explorers and touring enthusiasts seeking the ultimate adventure platform with legendary reliability, extreme versatility, German precision engineering, and proven world-class capability for expeditions across all terrain.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For serious adventure touring enthusiasts</span></p>
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
		return fmt.Errorf("error creating bmw gs adventure review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for BMW GS Adventure (Rating: 4.8/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">BMW GS অ্যাডভেঞ্চার রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম অ্যাডভেঞ্চার ট্যুরিং লেজেন্ড</h1>
<p class="summary-text">BMW GS অ্যাডভেঞ্চার একটি 1200cc প্রিমিয়াম অ্যাডভেঞ্চার ট্যুরিং বাইক যা জার্মান প্রকৌশল শ্রেষ্ঠত্ব, কিংবদন্তি বহুমুখিতা এবং অতুলনীয় ক্ষমতা প্রদান করে। ৳28,50,000 টাকায় মূল্যায়িত, এটি উন্নত প্রযুক্তি, মজবুত স্থায়িত্ব এবং কিংবদন্তি নির্ভরযোগ্যতা একত্রিত করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">BMW GS অ্যাডভেঞ্চার মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">1200cc বক্সার টুইন:</strong> <span class="highlight-value">কিংবদন্তি টর্ক ডেলিভারি</span></li>
<li class="highlight-item"><strong class="highlight-label">অ্যাডভেঞ্চার ঐতিহ্য:</strong> <span class="highlight-value">40+ বছর প্রমাণিত প্ল্যাটফর্ম</span></li>
<li class="highlight-item"><strong class="highlight-label">অফ-রোড ক্ষমতা:</strong> <span class="highlight-value">চরম ভূখণ্ড প্রস্তুত</span></li>
<li class="highlight-item"><strong class="highlight-label">জার্মান প্রকৌশল:</strong> <span class="highlight-value">নির্ভুলতা নির্ভরযোগ্যতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">BMW GS অ্যাডভেঞ্চার সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">কিংবদন্তি প্ল্যাটফর্ম:</strong> <span class="pro-description">40+ বছর ঐতিহ্য</span></li>
<li class="pro-item"><strong class="pro-label">অফ-রোড বিস্ট:</strong> <span class="pro-description">চরম ভূখণ্ড প্রস্তুত</span></li>
<li class="pro-item"><strong class="pro-label">বক্সার ইঞ্জিন:</strong> <span class="pro-description">টর্কি চরিত্র</span></li>
<li class="pro-item"><strong class="pro-label">আরাম ট্যুরিং:</strong> <span class="pro-description">দীর্ঘ-দূরত্ব সক্ষম</span></li>
<li class="pro-item"><strong class="pro-label">প্রযুক্তি প্যাকড:</strong> <span class="pro-description">উন্নত সিস্টেম</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">BMW GS অ্যাডভেঞ্চার অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">অতি প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳28,50,000 চরম</span></li>
<li class="con-item"><strong class="con-label">উচ্চ রক্ষণাবেক্ষণ:</strong> <span class="con-description">খুবই ব্যয়বহুল রক্ষণাবেক্ষণ</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">234kg চ্যালেঞ্জিং</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: BMW GS অ্যাডভেঞ্চার কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.8/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳28,50,000 টাকায়, GS অ্যাডভেঞ্চার গুরুতর অনুসন্ধানকারীদের জন্য চূড়ান্ত প্ল্যাটফর্ম যারা কিংবদন্তি নির্ভরযোগ্যতা এবং চরম বহুমুখিতা খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- গুরুতর অ্যাডভেঞ্চার ট্যুরিং উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for BMW GS Adventure already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for BMW GS Adventure\n")

	return nil
}

// GetName returns the seeder name
func (s *BMWGSAdventureReviewBatch18) GetName() string {
	return "BMWGSAdventureReviewBatch18"
}
