package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajPulsarF150Review handles seeding Bajaj Pulsar F150 product review and translation
type BajajPulsarF150Review struct {
	BaseSeeder
}

// NewBajajPulsarF150ReviewSeeder creates a new BajajPulsarF150Review
func NewBajajPulsarF150ReviewSeeder() *BajajPulsarF150Review {
	return &BajajPulsarF150Review{BaseSeeder: BaseSeeder{name: "Bajaj Pulsar F150 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajPulsarF150Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Pulsar F150").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj pulsar f150 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bajaj pulsar f150 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Pulsar F150 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Pulsar F150 Review Bangladesh 2024 - Faired Performance Bike</h1>
<p class="summary-text">The Bajaj Pulsar F150 is a faired 150cc street-sport bike combining aerodynamic design, sporty performance, and everyday practicality. Priced around ৳2,45,000, it delivers balanced power with wind protection and aggressive styling for daily urban commuting.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Pulsar F150 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc Faired Engine:</strong> <span class="highlight-value">Street-sport performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Aerodynamic Design:</strong> <span class="highlight-value">Wind protection fairing</span></li>
<li class="highlight-item"><strong class="highlight-label">Sporty Styling:</strong> <span class="highlight-value">Aggressive looks</span></li>
<li class="highlight-item"><strong class="highlight-label">Good Mileage:</strong> <span class="highlight-value">48-52 km/l decent</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Pulsar F150 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Faired Design:</strong> <span class="pro-description">Wind protection aerodynamic</span></li>
<li class="pro-item"><strong class="pro-label">Sporty Performance:</strong> <span class="pro-description">150cc responsive</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Modern faired looks</span></li>
<li class="pro-item"><strong class="pro-label">Practical Bike:</strong> <span class="pro-description">Urban commute ready</span></li>
<li class="pro-item"><strong class="pro-label">Bajaj Reliability:</strong> <span class="pro-description">Proven performance</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Pulsar F150 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Sport-focused firm</span></li>
<li class="con-item"><strong class="con-label">Handling:</strong> <span class="con-description">Slightly heavier</span></li>
<li class="con-item"><strong class="con-label">Power:</strong> <span class="con-description">150cc moderate</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Pulsar F150 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,45,000 - Faired street bike</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳4-5 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">48-52 km/l decent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Pulsar F150 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Power:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- 150cc street</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Faired sporty</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Commute suitable</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Bajaj proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good balance</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Pulsar F150?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,45,000, the Bajaj Pulsar F150 is perfect for urban riders seeking a faired sport bike that balances aggressive styling, wind protection, responsive performance, and practical daily commuting capability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For faired sport commuters</span></p>
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
		return fmt.Errorf("error creating bajaj pulsar f150 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Pulsar F150 (Rating: 4.4/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ Pulsar F150 রিভিউ বাংলাদেশ 2024 - ফেয়ারড পারফরম্যান্স বাইক</h1>
<p class="summary-text">বাজাজ Pulsar F150 একটি ফেয়ারড 150cc স্ট্রিট-ক্রীড়া বাইক যা বায়ুগতিশীল ডিজাইন, খেলাধুলা কর্মক্ষমতা এবং দৈনন্দিন ব্যবহারিকতা একত্রিত করে। ৳2,45,000 টাকায় মূল্যায়িত, এটি দৈনিক শহুরে যাতায়াতের জন্য বায়ু সুরক্ষা এবং আক্রমণাত্মক স্টাইলিং সহ সুষম শক্তি প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ Pulsar F150 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">150cc ফেয়ারড ইঞ্জিন:</strong> <span class="highlight-value">স্ট্রিট-ক্রীড়া কর্মক্ষমতা</span></li>
<li class="highlight-item"><strong class="highlight-label">বায়ুগতিশীল ডিজাইন:</strong> <span class="highlight-value">বায়ু সুরক্ষা ফেয়ারিং</span></li>
<li class="highlight-item"><strong class="highlight-label">খেলাধুলা স্টাইলিং:</strong> <span class="highlight-value">আক্রমণাত্মক চেহারা</span></li>
<li class="highlight-item"><strong class="highlight-label">ভাল মাইলেজ:</strong> <span class="highlight-value">48-52 km/l মানসম্পন্ন</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ Pulsar F150 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ফেয়ারড ডিজাইন:</strong> <span class="pro-description">বায়ু সুরক্ষা বায়ুগতিশীল</span></li>
<li class="pro-item"><strong class="pro-label">খেলাধুলা কর্মক্ষমতা:</strong> <span class="pro-description">150cc প্রতিক্রিয়াশীল</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">আধুনিক ফেয়ারড চেহারা</span></li>
<li class="pro-item"><strong class="pro-label">ব্যবহারিক বাইক:</strong> <span class="pro-description">শহুরে যাত্রা প্রস্তুত</span></li>
<li class="pro-item"><strong class="pro-label">বাজাজ নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত কর্মক্ষমতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ Pulsar F150 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">আসনের আরাম:</strong> <span class="con-description">ক্রীড়া-ফোকাসড দৃঢ়</span></li>
<li class="con-item"><strong class="con-label">হ্যান্ডলিং:</strong> <span class="con-description">সামান্য ভারী</span></li>
<li class="con-item"><strong class="con-label">শক্তি:</strong> <span class="con-description">150cc মধ্যম</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ Pulsar F150 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.4/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳2,45,000 টাকায়, বাজাজ Pulsar F150 শহুরে রাইডারদের জন্য নিখুঁত যারা একটি ফেয়ারড ক্রীড়া বাইক খুঁজছেন যা আক্রমণাত্মক স্টাইলিং, বায়ু সুরক্ষা, প্রতিক্রিয়াশীল কর্মক্ষমতা এবং ব্যবহারিক দৈনিক যাতায়াত ক্ষমতা একত্রিত করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ফেয়ারড ক্রীড়া কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Pulsar F150 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Pulsar F150\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajPulsarF150Review) GetName() string {
	return "BajajPulsarF150Review"
}
