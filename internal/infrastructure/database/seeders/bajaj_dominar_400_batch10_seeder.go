package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajDominar400Review handles seeding Bajaj Dominar 400 product review and translation
type BajajDominar400Review struct {
	BaseSeeder
}

// NewBajajDominar400ReviewSeeder creates a new BajajDominar400Review
func NewBajajDominar400ReviewSeeder() *BajajDominar400Review {
	return &BajajDominar400Review{BaseSeeder: BaseSeeder{name: "Bajaj Dominar 400 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajDominar400Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Dominar 400").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Bajaj Dominar 400 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding Bajaj Dominar 400 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Dominar 400 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Dominar 400 Review Bangladesh 2024 - Complete Guide</h1>
<p class="summary-text">The Bajaj Dominar 400 is a performance-oriented, premium naked bike featuring 400cc liquid-cooled engine, dual ABS, upside-down forks, and aggressive styling. Priced around ৳4,20,000, it delivers genuine sports performance with highway capability and practical daily usability. Perfect for enthusiast riders seeking powerful middleweight performance without premium pricing.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Dominar 400 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">400cc Liquid-Cooled:</strong> <span class="highlight-value">High-performance engine</span></li>
<li class="highlight-item"><strong class="highlight-label">Dual ABS:</strong> <span class="highlight-value">Front and rear ABS safety</span></li>
<li class="highlight-item"><strong class="highlight-label">Upside-Down Forks:</strong> <span class="highlight-value">Premium inverted suspension</span></li>
<li class="highlight-item"><strong class="highlight-label">Aggressive Design:</strong> <span class="highlight-value">Muscular naked bike styling</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Dominar 400 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Strong Performance:</strong> <span class="pro-description">Powerful 400cc engine with impressive acceleration</span></li>
<li class="pro-item"><strong class="pro-label">Dual ABS:</strong> <span class="pro-description">Front and rear ABS for confident braking</span></li>
<li class="pro-item"><strong class="pro-label">Premium Handling:</strong> <span class="pro-description">Upside-down forks and excellent suspension</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Muscular design stands out</span></li>
<li class="pro-item"><strong class="pro-label">Highway Capable:</strong> <span class="pro-description">Comfortable for long-distance touring</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Dominar 400 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Pricing:</strong> <span class="con-description">Expensive compared to smaller bikes</span></li>
<li class="con-item"><strong class="con-label">Higher Running Cost:</strong> <span class="con-description">Fuel consumption and maintenance higher</span></li>
<li class="con-item"><strong class="con-label">Not for Beginners:</strong> <span class="con-description">Powerful bike needs experienced riders</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Dominar 400 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,20,000 - Premium performance value</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳8-10 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Cost Per Month:</strong> <span class="value-amount">৳8-10 per km (25-30 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Dominar 400 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Excellent 400cc power</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Premium suspension and control</span></li>
<li class="rating-item"><strong class="rating-label">Safety:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Dual ABS on both wheels</span></li>
<li class="rating-item"><strong class="rating-label">Features:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Premium components</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Good value for performance</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Dominar 400?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳4,20,000, the Bajaj Dominar 400 offers genuine performance with premium features and safety technology. Excellent for experienced riders seeking powerful 400cc performance without premium brand pricing.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance-oriented riders seeking highway capability</span></p>
</div>
</section>
</article>`

	review := &models.ProductReviewModel{
		ProductID: product.ID,
		UserID:    adminUser.ID,
		Rating:    4.6,
		Reviews:   &englishReview,
		Priority:  1,
		Status:    true,
	}

	if err := db.Create(review).Error; err != nil {
		return fmt.Errorf("error creating Bajaj Dominar 400 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Dominar 400 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাजाج डोमिनार ४०० रिव्यू बांग्लादेश २०२४ - सम्पूर्ण गाइड</h1>
<p class="summary-text">বাজाज ডোমिনার ४०० एक কর्मক्षमता-उन्मुख, প्রিमিয়াম裸 बाइक যা ४००cc liquid-cooled ইঞ్জिन, dual ABS, upside-down forks এবং আক্रમণাত্মक স্টাইলिং সহ আসে। ৳४,२०,००० টाকায় মূल्यায়িত, এটি হাইওয়ে সক্ষমতা এবং ব्যবহارিक দৈनिक usability সহ খাঁটি খেলা পারফরমেন्स প्रदान करता है।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বाजाज डोमिनार ४०० मुख्य विशेषताएं</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">४००cc सतत-ठंडा:</strong> <span class="highlight-value">उच्च-कार्यक्षमता ইঞ్జिन</span></li>
<li class="highlight-item"><strong class="highlight-label">Dual ABS:</strong> <span class="highlight-value">সামনে এবং পিছনে ABS নিরাপত্তা</span></li>
<li class="highlight-item"><strong class="highlight-label">उल्टे forks:</strong> <span class="highlight-value">প্রিমিয়াম আপস ডাউন সাসপেশন</span></li>
<li class="highlight-item"><strong class="highlight-label">आक्रामक ডিজाইন:</strong> <span class="highlight-value">পেশী naked বাইক স্টাইলিং</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">बाजाज डोमिनार ४०० फायदे</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক्তिশালী पारফরמेंस:</strong> <span class="pro-description">চিত্তাকর্ষক ত্বরণ সহ শক্তিশালী ४००cc ইঞ્જिन</span></li>
<li class="pro-item"><strong class="pro-label">Dual ABS:</strong> <span class="pro-description">আত्मবिश्वाস ব्रेकिং এর জন्য সামнে এবং পিছনে ABS</span></li>
<li class="pro-item"><strong class="pro-label">প्रिමيय়াম हैंड़लिং:</strong> <span class="pro-description">उल्टे forks और उत्कृष्ट सरस्पेंशन</span></li>
<li class="pro-item"><strong class="pro-label">আক্রামক স্टाইल:</strong> <span class="pro-description">পেশী ডিজাइন দাঁড়িয়ে থাকে</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">बाजाज डोमिनार ४०० अपूर्णताएं</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমिয়اম मূल्य:</strong> <span class="con-description">ছোট বাইকের তুলনায় ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">उच्चतर চलমান खरच:</strong> <span class="con-description">জ্বালানি এবং রক্ষণাবেক্ষণ খরচ বেশি</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">चূড़ांत निर्णय: बाजाज डोमिनार ४०० किनें?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ्रिक রেটিং:</strong> <span class="score-value">४.६/५.०</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳४,२०,००० টাकায়, বाजाज डोमिनार ४००প্রিমিয়াম বৈশिষ্ট्য এবং নिरাপত्তा প्रযুक्ति সহ খাঁটি পারফরमेंস প्রदान করে। অভिজ্ञ রাইডারদের জন्য চমৎকার।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">सुपारिश:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পারফরমেंस-ओरिएंtéd রাইডারদের জন्য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Dominar 400 already exists\n")
		return nil
	}

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating Bangla translation: %w", err)
	}

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Dominar 400\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajDominar400Review) GetName() string {
	return "BajajDominar400Review"
}
