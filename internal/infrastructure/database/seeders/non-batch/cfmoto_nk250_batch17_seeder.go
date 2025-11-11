package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// CFMotoNK250ReviewBatch17 handles seeding CFMOTO NK250 product review and translation
type CFMotoNK250ReviewBatch17 struct {
	BaseSeeder
}

// NewCFMotoNK250ReviewBatch17Seeder creates a new CFMotoNK250ReviewBatch17
func NewCFMotoNK250ReviewBatch17Seeder() *CFMotoNK250ReviewBatch17 {
	return &CFMotoNK250ReviewBatch17{BaseSeeder: BaseSeeder{name: "CFMOTO NK250 Batch17 Review"}}
}

// Seed implements the Seeder interface
func (s *CFMotoNK250ReviewBatch17) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "CFMOTO NK250").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("cfmoto nk250 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding cfmoto nk250 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for CFMOTO NK250 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">CFMOTO NK250 Review Bangladesh 2024 - Chinese Sport Aggressor</h1>
<p class="summary-text">The CFMOTO NK250 is a 250cc aggressive naked bike combining sporty styling, responsive performance, and budget-friendly pricing. Priced around ৳3,25,000, it delivers raw street energy, modern features, decent handling, and excellent value for performance-hungry budget riders seeking true street spirit.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">CFMOTO NK250 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">250cc Naked:</strong> <span class="highlight-value">Street aggression</span></li>
<li class="highlight-item"><strong class="highlight-label">Responsive Engine:</strong> <span class="highlight-value">Direct throttle</span></li>
<li class="highlight-item"><strong class="highlight-label">Modern Styling:</strong> <span class="highlight-value">Aggressive lines</span></li>
<li class="highlight-item"><strong class="highlight-label">Budget Value:</strong> <span class="highlight-value">Great price</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">CFMOTO NK250 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Street attitude</span></li>
<li class="pro-item"><strong class="pro-label">Responsive Power:</strong> <span class="pro-description">Direct performance</span></li>
<li class="pro-item"><strong class="pro-label">Great Handling:</strong> <span class="pro-description">Agile dynamics</span></li>
<li class="pro-item"><strong class="pro-label">Budget Price:</strong> <span class="pro-description">Excellent value</span></li>
<li class="pro-item"><strong class="pro-label">Modern Features:</strong> <span class="pro-description">ABS equipped</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">CFMOTO NK250 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Limited Service:</strong> <span class="con-description">Fewer dealers</span></li>
<li class="con-item"><strong class="con-label">Moderate Power:</strong> <span class="con-description">Not high performance</span></li>
<li class="con-item"><strong class="con-label">Build Quality:</strong> <span class="con-description">Budget level</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">CFMOTO NK250 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳3,25,000 - Budget sport</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳6-8 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">32-36 km/l excellent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">CFMOTO NK250 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Direct response</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Agile dynamics</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Modern aggressive</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Excellent budget</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Modern equipped</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy CFMOTO NK250?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳3,25,000, the NK250 is perfect for budget-conscious street enthusiasts wanting authentic 250cc naked sport with aggressive styling, responsive performance, and excellent value without premium pricing.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For budget sport enthusiasts</span></p>
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
		return fmt.Errorf("error creating cfmoto nk250 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for CFMOTO NK250 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">CFMOTO NK250 রিভিউ বাংলাদেশ 2024 - চাইনিজ স্পোর্ট এগ্রেসর</h1>
<p class="summary-text">CFMOTO NK250 একটি 250cc আক্রমণাত্মক নেকেড বাইক যা খেলাধুলামূলক স্টাইলিং, প্রতিক্রিয়াশীল কর্মক্ষমতা এবং বাজেট-বান্ধব মূল্য একত্রিত করে। ৳3,25,000 টাকায় মূল্যায়িত, এটি কাঁচা রাস্তার শক্তি, আধুনিক বৈশিষ্ট্য এবং চমৎকার মূল্য প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">CFMOTO NK250 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">250cc নেকেড:</strong> <span class="highlight-value">রাস্তা আক্রমণ</span></li>
<li class="highlight-item"><strong class="highlight-label">প্রতিক্রিয়াশীল ইঞ্জিন:</strong> <span class="highlight-value">সরাসরি থ্রটল</span></li>
<li class="highlight-item"><strong class="highlight-label">আধুনিক স্টাইলিং:</strong> <span class="highlight-value">আক্রমণাত্মক লাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">বাজেট মূল্য:</strong> <span class="highlight-value">দুর্দান্ত দাম</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">CFMOTO NK250 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">রাস্তা মনোভাব</span></li>
<li class="pro-item"><strong class="pro-label">প্রতিক্রিয়াশীল শক্তি:</strong> <span class="pro-description">সরাসরি কর্মক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">দুর্দান্ত হ্যান্ডলিং:</strong> <span class="pro-description">চটপটে গতিশীলতা</span></li>
<li class="pro-item"><strong class="pro-label">বাজেট দাম:</strong> <span class="pro-description">চমৎকার মূল্য</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক বৈশিষ্ট্য:</strong> <span class="pro-description">ABS সজ্জিত</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">CFMOTO NK250 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">সীমিত সেবা:</strong> <span class="con-description">কম ডিলার</span></li>
<li class="con-item"><strong class="con-label">মধ্যম শক্তি:</strong> <span class="con-description">উচ্চ কর্মক্ষমতা নয়</span></li>
<li class="con-item"><strong class="con-label">বিল্ড গুণমান:</strong> <span class="con-description">বাজেট স্তর</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: CFMOTO NK250 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳3,25,000 টাকায়, NK250 বাজেট-সচেতন রাস্তা উত্সাহীদের জন্য নিখুঁত যারা খাঁটি 250cc নেকেড স্পোর্ট চান আক্রমণাত্মক স্টাইলিং এবং চমৎকার মূল্য সহ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বাজেট খেলাধুলা উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for CFMOTO NK250 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for CFMOTO NK250\n")

	return nil
}

// GetName returns the seeder name
func (s *CFMotoNK250ReviewBatch17) GetName() string {
	return "CFMotoNK250ReviewBatch17"
}
