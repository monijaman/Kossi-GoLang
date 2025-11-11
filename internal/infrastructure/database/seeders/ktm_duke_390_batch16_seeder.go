package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// KTMDuke390Review handles seeding KTM Duke 390 product review and translation
type KTMDuke390Review struct {
	BaseSeeder
}

// NewKTMDuke390ReviewSeeder creates a new KTMDuke390Review
func NewKTMDuke390ReviewSeeder() *KTMDuke390Review {
	return &KTMDuke390Review{BaseSeeder: BaseSeeder{name: "KTM Duke 390 Review"}}
}

// Seed implements the Seeder interface
func (s *KTMDuke390Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "KTM Duke 390").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("ktm duke 390 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding ktm duke 390 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for KTM Duke 390 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">KTM Duke 390 Review Bangladesh 2024 - Premium Street Beast</h1>
<p class="summary-text">The KTM Duke 390 is a premium 390cc street fighter with powerful engine, advanced technology, sharp handling, and aggressive design. Priced around ৳4,50,000, it's the ultimate street performance bike for riders seeking thrilling acceleration and cutting-edge features.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">KTM Duke 390 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">390cc Powerful Engine:</strong> <span class="highlight-value">Premium performance beast</span></li>
<li class="highlight-item"><strong class="highlight-label">Advanced Technology:</strong> <span class="highlight-value">Cutting-edge features</span></li>
<li class="highlight-item"><strong class="highlight-label">Sharp Handling:</strong> <span class="highlight-value">Sport-oriented dynamics</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Design:</strong> <span class="highlight-value">Premium street styling</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">KTM Duke 390 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Powerful Engine:</strong> <span class="pro-description">390cc thrilling</span></li>
<li class="pro-item"><strong class="pro-label">Advanced Tech:</strong> <span class="pro-description">ABS, slipper clutch</span></li>
<li class="pro-item"><strong class="pro-label">Sharp Handling:</strong> <span class="pro-description">Sport-oriented agile</span></li>
<li class="pro-item"><strong class="pro-label">Premium Design:</strong> <span class="pro-description">Aggressive styling</span></li>
<li class="pro-item"><strong class="pro-label">KTM Performance:</strong> <span class="pro-description">Racing pedigree</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">KTM Duke 390 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">₳4,50,000 expensive</span></li>
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Sport-focused firm</span></li>
<li class="con-item"><strong class="con-label">Fuel Efficiency:</strong> <span class="con-description">35-40 km/l average</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">KTM Duke 390 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,50,000 - Premium street</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳8-10 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">35-40 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">KTM Duke 390 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine Power:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- 390cc powerful</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Sport sharp</span></li>
<li class="rating-item"><strong class="rating-label">Technology:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Advanced ABS</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Aggressive premium</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Premium price</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy KTM Duke 390?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳4,50,000, the KTM Duke 390 is the choice for performance enthusiasts seeking premium engineering, thrilling acceleration, sport-oriented handling, advanced technology, and authentic street-fighting character.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance enthusiasts</span></p>
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
		return fmt.Errorf("error creating ktm duke 390 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for KTM Duke 390 (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">KTM Duke 390 রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম স্ট্রিট বিস্ট</h1>
<p class="summary-text">KTM Duke 390 একটি প্রিমিয়াম 390cc স্ট্রিট ফাইটার যা শক্তিশালী ইঞ্জিন, উন্নত প্রযুক্তি, তীক্ষ্ণ হ্যান্ডলিং এবং আক্রমণাত্মক ডিজাইন সহ আসে। ৳4,50,000 টাকায় মূল্যায়িত, এটি রোমাঞ্চকর ত্বরণ এবং অত্যাধুনিক বৈশিষ্ট্য খুঁজছেন রাইডারদের জন্য চূড়ান্ত স্ট্রিট কর্মক্ষমতা বাইক।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">KTM Duke 390 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">390cc শক্তিশালী ইঞ্জিন:</strong> <span class="highlight-value">প্রিমিয়াম কর্মক্ষমতা শক্তিশালী</span></li>
<li class="highlight-item"><strong class="highlight-label">উন্নত প্রযুক্তি:</strong> <span class="highlight-value">অত্যাধুনিক বৈশিষ্ট্য</span></li>
<li class="highlight-item"><strong class="highlight-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="highlight-value">ক্রীড়া-ভিত্তিক গতিশীলতা</span></li>
<li class="highlight-item"><strong class="highlight-label">আক্রমণাত্মক ডিজাইন:</strong> <span class="highlight-value">প্রিমিয়াম স্ট্রিট স্টাইলিং</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">KTM Duke 390 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী ইঞ্জিন:</strong> <span class="pro-description">390cc রোমাঞ্চকর</span></li>
<li class="pro-item"><strong class="pro-label">উন্নত প্রযুক্তি:</strong> <span class="pro-description">ABS, স্লিপার ক্লাচ</span></li>
<li class="pro-item"><strong class="pro-label">তীক্ষ্ণ হ্যান্ডলিং:</strong> <span class="pro-description">ক্রীড়া-ভিত্তিক চটপটে</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম ডিজাইন:</strong> <span class="pro-description">আক্রমণাত্মক স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">KTM কর্মক্ষমতা:</strong> <span class="pro-description">রেসিং বংশপরম্পরা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">KTM Duke 390 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳4,50,000 ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">আসনের আরাম:</strong> <span class="con-description">ক্রীড়া-ফোকাসড দৃঢ়</span></li>
<li class="con-item"><strong class="con-label">জ্বালানী দক্ষতা:</strong> <span class="con-description">35-40 km/l গড়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: KTM Duke 390 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳4,50,000 টাকায়, KTM Duke 390 কর্মক্ষমতা উত্সাহীদের পছন্দ যারা প্রিমিয়াম প্রকৌশল, রোমাঞ্চকর ত্বরণ, ক্রীড়া-ভিত্তিক হ্যান্ডলিং, উন্নত প্রযুক্তি এবং খাঁটি স্ট্রিট-ফাইটিং চরিত্র খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- কর্মক্ষমতা উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for KTM Duke 390 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for KTM Duke 390\n")

	return nil
}

// GetName returns the seeder name
func (s *KTMDuke390Review) GetName() string {
	return "KTMDuke390Review"
}
