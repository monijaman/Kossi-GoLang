package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// KawasakiNinja650ReviewBatch16 handles seeding Kawasaki Ninja 650 product review and translation
type KawasakiNinja650ReviewBatch16 struct {
	BaseSeeder
}

// NewKawasakiNinja650ReviewBatch16Seeder creates a new KawasakiNinja650ReviewBatch16
func NewKawasakiNinja650ReviewBatch16Seeder() *KawasakiNinja650ReviewBatch16 {
	return &KawasakiNinja650ReviewBatch16{BaseSeeder: BaseSeeder{name: "Kawasaki Ninja 650 Batch16 Review"}}
}

// Seed implements the Seeder interface
func (s *KawasakiNinja650ReviewBatch16) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Kawasaki Ninja 650").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("kawasaki ninja 650 product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding kawasaki ninja 650 product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Kawasaki Ninja 650 already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Kawasaki Ninja 650 Review Bangladesh 2024 - Mid-Weight Sports Thrills</h1>
<p class="summary-text">The Kawasaki Ninja 650 is a mid-weight 650cc sport bike combining thrilling acceleration, agile handling, modern styling, and practical ergonomics. Priced around ৳7,80,000, it's the perfect entry-level supersport for riders seeking real sport performance with everyday usability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Kawasaki Ninja 650 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">650cc Sport Engine:</strong> <span class="highlight-value">Mid-weight thrilling</span></li>
<li class="highlight-item"><strong class="highlight-label">Agile Handling:</strong> <span class="highlight-value">Sport-tuned responsive</span></li>
<li class="highlight-item"><strong class="highlight-label">Modern Design:</strong> <span class="highlight-value">Contemporary styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Practical Ergonomics:</strong> <span class="highlight-value">Usable for daily riding</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Kawasaki Ninja 650 Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Sport Performance:</strong> <span class="pro-description">650cc thrilling</span></li>
<li class="pro-item"><strong class="pro-label">Agile Handling:</strong> <span class="pro-description">Sport-tuned sharp</span></li>
<li class="pro-item"><strong class="pro-label">Practical Seating:</strong> <span class="pro-description">Usable comfortable</span></li>
<li class="pro-item"><strong class="pro-label">Kawasaki Quality:</strong> <span class="pro-description">Proven reliability</span></li>
<li class="pro-item"><strong class="pro-label">Good Mileage:</strong> <span class="pro-description">35-40 km/l decent</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Kawasaki Ninja 650 Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">₳7,80,000 expensive</span></li>
<li class="con-item"><strong class="con-label">Insurance Cost:</strong> <span class="con-description">Higher than average</span></li>
<li class="con-item"><strong class="con-label">Maintenance:</strong> <span class="con-description">Premium parts</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Kawasaki Ninja 650 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳7,80,000 - Mid-weight sport</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate-high - ৳10-12 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">35-40 km/l decent</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Kawasaki Ninja 650 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Sport Performance:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- 650cc thrilling</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Sport sharp</span></li>
<li class="rating-item"><strong class="rating-label">Design:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Modern sleek</span></li>
<li class="rating-item"><strong class="rating-label">Reliability:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Kawasaki proven</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Mid-weight premium</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Kawasaki Ninja 650?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳7,80,000, the Kawasaki Ninja 650 is ideal for sport riders seeking a mid-weight bike that delivers real performance thrills, agile handling, modern design, practical everyday usability, and Kawasaki's trusted quality.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For sport enthusiasts</span></p>
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
		return fmt.Errorf("error creating kawasaki ninja 650 review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Kawasaki Ninja 650 (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">কাওয়াসাকি নিনজা 650 রিভিউ বাংলাদেশ 2024 - মিড-ওয়েট স্পোর্টস রোমাঞ্চ</h1>
<p class="summary-text">কাওয়াসাকি নিনজা 650 একটি মিড-ওয়েট 650cc খেলাধুলা বাইক যা রোমাঞ্চকর ত্বরণ, চটপটে হ্যান্ডলিং, আধুনিক স্টাইলিং এবং ব্যবহারিক এরগনমিক্স একত্রিত করে। ৳7,80,000 টাকায় মূল্যায়িত, এটি রাইডারদের জন্য নিখুঁত এন্ট্রি-লেভেল সুপারস্পোর্ট যারা দৈনন্দিন ব্যবহারযোগ্যতা সহ বাস্তব খেলাধুলা কর্মক্ষমতা খুঁজছেন।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">কাওয়াসাকি নিনজা 650 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">650cc খেলাধুলা ইঞ্জিন:</strong> <span class="highlight-value">মিড-ওয়েট রোমাঞ্চকর</span></li>
<li class="highlight-item"><strong class="highlight-label">চটপটে হ্যান্ডলিং:</strong> <span class="highlight-value">খেলাধুলা-টিউনড প্রতিক্রিয়াশীল</span></li>
<li class="highlight-item"><strong class="highlight-label">আধুনিক ডিজাইন:</strong> <span class="highlight-value">সমসাময়িক স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্যবহারিক এরগনমিক্স:</strong> <span class="highlight-value">দৈনিক যাত্রার জন্য ব্যবহারযোগ্য</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">কাওয়াসাকি নিনজা 650 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">খেলাধুলা কর্মক্ষমতা:</strong> <span class="pro-description">650cc রোমাঞ্চকর</span></li>
<li class="pro-item"><strong class="pro-label">চটপটে হ্যান্ডলিং:</strong> <span class="pro-description">খেলাধুলা-টিউনড তীক্ষ্ণ</span></li>
<li class="pro-item"><strong class="pro-label">ব্যবহারিক আসন:</strong> <span class="pro-description">ব্যবহারযোগ্য আরামদায়ক</span></li>
<li class="pro-item"><strong class="pro-label">কাওয়াসাকি গুণমান:</strong> <span class="pro-description">প্রমাণিত নির্ভরযোগ্যতা</span></li>
<li class="pro-item"><strong class="pro-label">ভাল মাইলেজ:</strong> <span class="pro-description">35-40 km/l মানসম্পন্ন</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">কাওয়াসাকি নিনজা 650 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳7,80,000 ব্যয়বহুল</span></li>
<li class="con-item"><strong class="con-label">বীমা খরচ:</strong> <span class="con-description">গড়ের চেয়ে বেশি</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ:</strong> <span class="con-description">প্রিমিয়াম যন্ত্রাংশ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: কাওয়াসাকি নিনজা 650 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳7,80,000 টাকায়, কাওয়াসাকি নিনজা 650 খেলাধুলা রাইডারদের জন্য আদর্শ যারা একটি মিড-ওয়েট বাইক খুঁজছেন যা বাস্তব কর্মক্ষমতা রোমাঞ্চ, চটপটে হ্যান্ডলিং, আধুনিক ডিজাইন, ব্যবহারিক দৈনন্দিন ব্যবহারযোগ্যতা এবং কাওয়াসাকির বিশ্বস্ত গুণমান প্রদান করে।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- খেলাধুলা উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Kawasaki Ninja 650 already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Kawasaki Ninja 650\n")

	return nil
}

// GetName returns the seeder name
func (s *KawasakiNinja650ReviewBatch16) GetName() string {
	return "KawasakiNinja650ReviewBatch16"
}
