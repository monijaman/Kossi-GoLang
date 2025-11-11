package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// KawasakiNinja400ReviewBatch19 handles seeding Kawasaki Ninja 400 product review and translation
type KawasakiNinja400ReviewBatch19 struct {
	BaseSeeder
}

// NewKawasakiNinja400ReviewBatch19Seeder creates a new KawasakiNinja400ReviewBatch19
func NewKawasakiNinja400ReviewBatch19Seeder() *KawasakiNinja400ReviewBatch19 {
	return &KawasakiNinja400ReviewBatch19{BaseSeeder: BaseSeeder{name: "Kawasaki Ninja 400 Batch19 Review"}}
}

// Seed implements the Seeder interface
func (s *KawasakiNinja400ReviewBatch19) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Kawasaki Ninja 400").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("kawasaki ninja 400 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding kawasaki ninja 400 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Kawasaki Ninja 400 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Kawasaki Ninja 400 Review Bangladesh 2024 - Entry-Level Sportbike Accessible Performance</h1>
<p class="summary-text">The Kawasaki Ninja 400 is a 400cc liquid-cooled parallel twin entry-level sportbike combining affordable pricing with exciting sport character. Priced around ৳4,80,000, it delivers user-friendly ergonomics, nimble handling, ABS safety, efficient power delivery, and Kawasaki's sports DNA for newcomers and experienced riders seeking accessible sportbike thrills.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Kawasaki Ninja 400 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">400cc Liquid-Cooled:</strong> <span class="highlight-value">Entry sportbike accessible</span></li>
<li class="highlight-item"><strong class="highlight-label">Nimble Character:</strong> <span class="highlight-value">Lightweight agile</span></li>
<li class="highlight-item"><strong class="highlight-label">ABS Safety:</strong> <span class="highlight-value">Modern safety equipped</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficiency:</strong> <span class="highlight-value">24-28 km/l good</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Kawasaki Ninja 400 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Entry-Friendly:</strong> <span class="pro-description">Accessible sportbike</span></li>
<li class="pro-item"><strong class="pro-label">Nimble Handling:</strong> <span class="pro-description">169kg lightweight</span></li>
<li class="pro-item"><strong class="pro-label">Fun Factor:</strong> <span class="pro-description">Sport character thrills</span></li>
<li class="pro-item"><strong class="pro-label">ABS Equipped:</strong> <span class="pro-description">Modern safety technology</span></li>
<li class="pro-item"><strong class="pro-label">Kawasaki DNA:</strong> <span class="pro-description">Genuine sportbike DNA</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Kawasaki Ninja 400 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Entry Power:</strong> <span class="con-description">35 bhp modest</span></li>
<li class="con-item"><strong class="con-label">Seat Comfort:</strong> <span class="con-description">Sport-focused seating</span></li>
<li class="con-item"><strong class="con-label">Storage Limited:</strong> <span class="con-description">Minimal luggage</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Kawasaki Ninja 400 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳4,80,000 - Entry sportbike</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳12-14 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">24-28 km/l good</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">400cc liquid-cooled parallel twin</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">35 bhp accessible sports</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">169kg lightweight agile</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Kawasaki Ninja 400 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Sportbike Character:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Genuine sport DNA</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Nimble lightweight</span></li>
<li class="rating-item"><strong class="rating-label">Entry-Friendly:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Accessible for all</span></li>
<li class="rating-item"><strong class="rating-label">Safety:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- ABS equipped</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- ৳4,80,000 solid</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Kawasaki Ninja 400?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳4,80,000, the Ninja 400 is perfect for newcomers and experienced riders seeking accessible sportbike thrills, nimble 400cc handling, genuine Kawasaki sport DNA, ABS safety technology, and entry-level performance with fun factor.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For entry-level sportbike seekers</span></p>
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
		return fmt.Errorf("error creating kawasaki ninja 400 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Kawasaki Ninja 400 (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">কাওয়াসাকি নিনজা 400 রিভিউ বাংলাদেশ 2024 - এন্ট্রি-লেভেল স্পোর্টবাইক সহজলভ্য পারফরম্যান্স</h1>
<p class="summary-text">কাওয়াসাকি নিনজা 400 একটি 400cc তরল-শীতলকৃত প্যারালাল টুইন এন্ট্রি-লেভেল স্পোর্টবাইক যা সাশ্রয়ী মূল্যকে উত্তেজনাপূর্ণ স্পোর্ট চরিত্রের সাথে একত্রিত করে। ৳4,80,000 টাকায় মূল্যায়িত, এটি ব্যবহারকারী-বান্ধব এরগোনমিক্স এবং কাওয়াসাকি স্পোর্টস DNA প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">কাওয়াসাকি নিনজা 400 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">400cc তরল-শীতলকৃত:</strong> <span class="highlight-value">এন্ট্রি স্পোর্টবাইক অ্যাক্সেসযোগ্য</span></li>
<li class="highlight-item"><strong class="highlight-label">চতুর চরিত্র:</strong> <span class="highlight-value">হালকা তত্পর</span></li>
<li class="highlight-item"><strong class="highlight-label">ABS নিরাপত্তা:</strong> <span class="highlight-value">আধুনিক নিরাপত্তা সজ্জিত</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষতা:</strong> <span class="highlight-value">24-28 km/l ভালো</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">কাওয়াসাকি নিনজা 400 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">এন্ট্রি-বান্ধব:</strong> <span class="pro-description">সহজলভ্য স্পোর্টবাইক</span></li>
<li class="pro-item"><strong class="pro-label">চতুর হ্যান্ডলিং:</strong> <span class="pro-description">169kg হালকা</span></li>
<li class="pro-item"><strong class="pro-label">মজা ফ্যাক্টর:</strong> <span class="pro-description">স্পোর্ট চরিত্র রোমাঞ্চ</span></li>
<li class="pro-item"><strong class="pro-label">ABS সজ্জিত:</strong> <span class="pro-description">আধুনিক নিরাপত্তা প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">কাওয়াসাকি DNA:</strong> <span class="pro-description">খাঁটি স্পোর্টবাইক DNA</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">কাওয়াসাকি নিনজা 400 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">এন্ট্রি শক্তি:</strong> <span class="con-description">35 bhp বিনম্র</span></li>
<li class="con-item"><strong class="con-label">সিট আরাম:</strong> <span class="con-description">স্পোর্ট-ফোকাসড সিটিং</span></li>
<li class="con-item"><strong class="con-label">স্টোরেজ সীমিত:</strong> <span class="con-description">ন্যূনতম লাগেজ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: কাওয়াসাকি নিনজা 400 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳4,80,000 টাকায়, নিনজা 400 নতুন এবং অভিজ্ঞ রাইডারদের জন্য নিখুঁত যারা সহজলভ্য স্পোর্টবাইক রোমাঞ্চ, চতুর 400cc হ্যান্ডলিং এবং খাঁটি কাওয়াসাকি স্পোর্টস DNA চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- এন্ট্রি-লেভেল স্পোর্টবাইক সন্ধানকারীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Kawasaki Ninja 400 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Kawasaki Ninja 400\n")

	return nil
}

// GetName returns the seeder name
func (s *KawasakiNinja400ReviewBatch19) GetName() string {
	return "KawasakiNinja400ReviewBatch19"
}
