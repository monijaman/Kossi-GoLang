package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TriumphStreetTwinReview handles seeding Triumph Street Twin product review and translation
type TriumphStreetTwinReview struct {
	BaseSeeder
}

// NewTriumphStreetTwinReviewSeeder creates a new TriumphStreetTwinReview
func NewTriumphStreetTwinReviewSeeder() *TriumphStreetTwinReview {
	return &TriumphStreetTwinReview{BaseSeeder: BaseSeeder{name: "Triumph Street Twin Review"}}
}

// Seed implements the Seeder interface
func (s *TriumphStreetTwinReview) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Triumph Street Twin").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("triumph street twin product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding triumph street twin product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Triumph Street Twin already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Triumph Street Twin Review Bangladesh 2024 - British Modern Classic</h1>
<p class="summary-text">The Triumph Street Twin is a 900cc modern classic motorcycle combining retro aesthetics with contemporary engineering, delivering authentic British character. Priced around ৳8,50,000, it offers refined simplicity, accessible performance, premium build quality, and commanding presence for riders seeking modern classic sophistication.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Triumph Street Twin Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">900cc Twin:</strong> <span class="highlight-value">British torque</span></li>
<li class="highlight-item"><strong class="highlight-label">Modern Classic:</strong> <span class="highlight-value">Retro styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Premium Build:</strong> <span class="highlight-value">Quality engineering</span></li>
<li class="highlight-item"><strong class="highlight-label">Accessible Power:</strong> <span class="highlight-value">Manageable thrills</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Triumph Street Twin Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">British Heritage:</strong> <span class="pro-description">Iconic character</span></li>
<li class="pro-item"><strong class="pro-label">Retro Design:</strong> <span class="pro-description">Timeless appeal</span></li>
<li class="pro-item"><strong class="pro-label">Twin Engine:</strong> <span class="pro-description">Smooth power</span></li>
<li class="pro-item"><strong class="pro-label">Premium Quality:</strong> <span class="pro-description">Solid build</span></li>
<li class="pro-item"><strong class="pro-label">Handling:</strong> <span class="pro-description">Balanced dynamics</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Triumph Street Twin Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">৳8,50,000 expensive</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">High maintenance</span></li>
<li class="con-item"><strong class="con-label">Speed Modest:</strong> <span class="con-description">Not sportbike fast</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Triumph Street Twin Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳8,50,000 - Premium classic</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳12-16 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">20-23 km/l moderate</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Triumph Street Twin Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Twin character</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Balanced feel</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Upright ergonomics</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Modern classic</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Premium positioning</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Triumph Street Twin?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳8,50,000, the Street Twin is perfect for riders appreciating British heritage, modern classic design, refined simplicity, and understated quality over flashy performance, delivering a sophisticated riding experience.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For British motorcycle enthusiasts</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.5,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating triumph street twin review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Triumph Street Twin (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ট্রায়াম্ফ স্ট্রিট টুইন রিভিউ বাংলাদেশ 2024 - ব্রিটিশ মডার্ন ক্লাসিক</h1>
<p class="summary-text">ট্রায়াম্ফ স্ট্রিট টুইন একটি 900cc আধুনিক ক্লাসিক মোটরসাইকেল যা রেট্রো নান্দনিকতাকে সমসাময়িক প্রকৌশলের সাথে একত্রিত করে, খাঁটি ব্রিটিশ চরিত্র প্রদান করে। ৳8,50,000 টাকায় মূল্যায়িত, এটি পরিমার্জিত সরলতা, অ্যাক্সেসযোগ্য কর্মক্ষমতা এবং প্রিমিয়াম বিল্ড গুণমান অফার করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ট্রায়াম্ফ স্ট্রিট টুইন মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">900cc টুইন:</strong> <span class="highlight-value">ব্রিটিশ টর্ক</span></li>
<li class="highlight-item"><strong class="highlight-label">মডার্ন ক্লাসিক:</strong> <span class="highlight-value">রেট্রো স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রিমিয়াম বিল্ড:</strong> <span class="highlight-value">গুণমান প্রকৌশল</span></li>
<li class="highlight-item"><strong class="highlight-label">অ্যাক্সেসযোগ্য শক্তি:</strong> <span class="highlight-value">পরিচালনাযোগ্য thrills</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ট্রায়াম্ফ স্ট্রিট টুইন সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ব্রিটিশ ঐতিহ্য:</strong> <span class="pro-description">আইকনিক চরিত্র</span></li>
<li class="pro-item"><strong class="pro-label">রেট্রো ডিজাইন:</strong> <span class="pro-description">কালজয়ী আবেদন</span></li>
<li class="pro-item"><strong class="pro-label">টুইন ইঞ্জিন:</strong> <span class="pro-description">মসৃণ শক্তি</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম গুণমান:</strong> <span class="pro-description">শক্ত নির্মাণ</span></li>
<li class="pro-item"><strong class="pro-label">হ্যান্ডলিং:</strong> <span class="pro-description">সুষম গতিশীলতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ট্রায়াম্ফ স্ট্রিট টুইন অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳8,50,000 ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">উচ্চ রক্ষণাবেক্ষণ</span></li>
<li class="con-item"><strong class="con-label">গতি মডেস্ট:</strong> <span class="con-description">স্পোর্টবাইক দ্রুত নয়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ট্রায়াম্ফ স্ট্রিট টুইন কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳8,50,000 টাকায়, স্ট্রিট টুইন চালকদের জন্য নিখুঁত যারা ব্রিটিশ ঐতিহ্য, আধুনিক ক্লাসিক ডিজাইন এবং পরিমার্জিত সরলতা মূল্যবান করেন চমৎকার অভিজ্ঞতার জন্য।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ব্রিটিশ মোটরসাইকেল উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Triumph Street Twin already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Triumph Street Twin\n")

	return nil
}

// GetName returns the seeder name
func (s *TriumphStreetTwinReview) GetName() string {
	return "TriumphStreetTwinReview"
}
