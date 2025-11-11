package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BajajAvenger220Review handles seeding Bajaj Avenger 220 product review and translation
type BajajAvenger220Review struct {
	BaseSeeder
}

// NewBajajAvenger220ReviewSeeder creates a new BajajAvenger220Review
func NewBajajAvenger220ReviewSeeder() *BajajAvenger220Review {
	return &BajajAvenger220Review{BaseSeeder: BaseSeeder{name: "Bajaj Avenger 220 Review"}}
}

// Seed implements the Seeder interface
func (s *BajajAvenger220Review) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Bajaj Avenger 220").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("bajaj avenger 220 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding bajaj avenger 220 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Bajaj Avenger 220 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Bajaj Avenger 220 Review Bangladesh 2024 - Cruiser Performance</h1>
<p class="summary-text">The Bajaj Avenger 220 is a muscular cruiser motorcycle featuring 220cc air-cooled engine, modern styling, comfortable cruiser seating, and excellent highway capability. Priced around ৳2,80,000, it offers genuine cruiser experience with powerful torque delivery and impressive mileage for long-distance riding.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Bajaj Avenger 220 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">220cc Air-Cooled:</strong> <span class="highlight-value">Powerful torque delivery</span></li>
<li class="highlight-item"><strong class="highlight-label">Cruiser Design:</strong> <span class="highlight-value">Muscular and aggressive styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfortable Seating:</strong> <span class="highlight-value">Low seat height and spacious design</span></li>
<li class="highlight-item"><strong class="highlight-label">Highway Capable:</strong> <span class="highlight-value">Excellent for long-distance touring</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Bajaj Avenger 220 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Strong Torque:</strong> <span class="pro-description">Powerful 220cc engine with excellent acceleration</span></li>
<li class="pro-item"><strong class="pro-label">Cruiser Comfort:</strong> <span class="pro-description">Relaxed riding position and spacious seating</span></li>
<li class="pro-item"><strong class="pro-label">Aggressive Styling:</strong> <span class="pro-description">Muscular cruiser design stands out</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">32-36 km/l fuel efficiency</span></li>
<li class="pro-item"><strong class="pro-label">Highway Ready:</strong> <span class="pro-description">Perfect for long-distance riding</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Bajaj Avenger 220 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Moderate Pricing:</strong> <span class="con-description">More expensive than smaller models</span></li>
<li class="con-item"><strong class="con-label">Heavy Bike:</strong> <span class="con-description">Heavier than sport bikes of same displacement</span></li>
<li class="con-item"><strong class="con-label">City Traffic:</strong> <span class="con-description">Not ideal for congested city riding</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Bajaj Avenger 220 Price in Bangladesh & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳2,80,000 - Premium cruiser value</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳5-6 per km</span></p>
<p class="value-item"><strong class="value-label">Monthly Fuel Cost:</strong> <span class="value-amount">৳5-6 per km (32-36 km/l)</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Bajaj Avenger 220 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Power:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Strong 220cc torque</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Excellent cruiser seating</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Aggressive styling</span></li>
<li class="rating-item"><strong class="rating-label">Highway Performance:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Excellent touring capability</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Good cruiser value</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Bajaj Avenger 220?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳2,80,000, the Bajaj Avenger 220 is excellent for riders seeking genuine cruiser experience with powerful performance, comfortable seating, and highway capability for long-distance adventures.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For cruiser enthusiasts and highway riders</span></p>
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
		return fmt.Errorf("error creating bajaj avenger 220 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Bajaj Avenger 220 (Rating: 4.5/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">বাজাজ অ্যাভেঞ্জার ২২০ রিভিউ বাংলাদেশ 2024 - ক্রুজার পারফরমেন্স</h1>
<p class="summary-text">বাজাজ অ্যাভেঞ্জার ২२० একটি পেশীবহুল ক্রুজার মোটরসাইকেল যা 220cc এয়ার-কুল্ড ইঞ্জিন, আধুনিক স্টাইলিং, আরামদায়ক ক্রুজার সিটিং এবং চমৎকার হাইওয়ে সামর্থ্য বৈশিষ্ট্যযুক্ত। ৳2,80,000 টাকায় মূল্যায়িত, এটি শক্তিশালী টর্ক ডেলিভারি এবং দীর্ঘ দূরত্বের যাত্রার জন্য চিত্তাকর্ষক মাইলেজ সহ খাঁটি ক্রুজার অভিজ্ঞতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">বাজাজ অ্যাভেঞ্জার ২२० মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">220cc এয়ার-কুল্ড:</strong> <span class="highlight-value">শক্তিশালী টর্ক ডেলিভারি</span></li>
<li class="highlight-item"><strong class="highlight-label">ক্রুজার ডিজাইন:</strong> <span class="highlight-value">পেশীবহুল এবং আক্রমণাত্মক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">আরামদায়ক সিটিং:</strong> <span class="highlight-value">কম সিট হাইট এবং প্রশস্ত ডিজাইন</span></li>
<li class="highlight-item"><strong class="highlight-label">হাইওয়ে সক্ষম:</strong> <span class="highlight-value">দীর্ঘ দূরত্বের ট্যুরিংয়ের জন্য চমৎকার</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">বাজাজ অ্যাভেঞ্জার ২२० সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">শক্তিশালী টর্ক:</strong> <span class="pro-description">চমৎকার ত্বরণ সহ শক্তিশালী 220cc ইঞ্জিন</span></li>
<li class="pro-item"><strong class="pro-label">ক্রুজার আরাম:</strong> <span class="pro-description">শিথিল রাইডিং অবস্থান এবং প্রশস্ত সিটিং</span></li>
<li class="pro-item"><strong class="pro-label">আক্রমণাত্মক স্টাইলিং:</strong> <span class="pro-description">পেশীবহুল ক্রুজার ডিজাইন দাঁড়িয়ে থাকে</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">32-36 km/l জ্বালানি দক্ষতা</span></li>
<li class="pro-item"><strong class="pro-label">হাইওয়ে প্রস্তুত:</strong> <span class="pro-description">দীর্ঘ দূরত্বের যাত্রার জন্য নিখুঁত</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">বাজাজ অ্যাভেঞ্জার २२० অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মধ্যম মূল্য:</strong> <span class="con-description">ছোট মডেলের চেয়ে বেশি ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">ভারী বাইক:</strong> <span class="con-description">একই স্থানচ্যুতির স্পোর্ট বাইকের চেয়ে ভারী</span></li>
<li class="con-item"><strong class="con-label">শহর ট্রাফিক:</strong> <span class="con-description">ভিড়ের শহর যাত্রার জন্য আদর্শ নয়</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: বাজাজ অ্যাভেঞ্জার २२२ কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.5/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳२,८०,००० টাকায়, বাজাজ অ্যাভেঞ্জার २२२ যারা শক্তিশালী পারফরমেন্স, আরামদায়ক সিটিং এবং দীর্ঘ দূরত্বের অ্যাডভেঞ্চারের জন্য হাইওয়ে সামর্থ্য সহ খাঁটি ক্রুজার অভিজ্ঞতা খুঁজছেন তাদের জন্য চমৎকার।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ক্রুজার উত্সাহী এবং হাইওয়ে রাইডারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Bajaj Avenger 220 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Bajaj Avenger 220\n")

	return nil
}

// GetName returns the seeder name
func (s *BajajAvenger220Review) GetName() string {
	return "BajajAvenger220Review"
}
