package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// PiaggioJetstreamReviewBatch19 handles seeding Piaggio Jetstream product review and translation
type PiaggioJetstreamReviewBatch19 struct {
	BaseSeeder
}

// NewPiaggioJetstreamReviewBatch19Seeder creates a new PiaggioJetstreamReviewBatch19
func NewPiaggioJetstreamReviewBatch19Seeder() *PiaggioJetstreamReviewBatch19 {
	return &PiaggioJetstreamReviewBatch19{BaseSeeder: BaseSeeder{name: "Piaggio Jetstream Batch19 Review"}}
}

// Seed implements the Seeder interface
func (s *PiaggioJetstreamReviewBatch19) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Piaggio Jetstream").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("piaggio jetstream product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding piaggio jetstream product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Piaggio Jetstream already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Piaggio Jetstream Review Bangladesh 2024 - Premium Maxi-Scooter Italian Design</h1>
<p class="summary-text">The Piaggio Jetstream is a 300cc liquid-cooled premium maxi-scooter combining Italian design with modern features. Priced around ৳5,20,000, it delivers sophisticated styling, comfortable seating, practical storage, advanced technology, and the Piaggio heritage for urban sophisticates seeking premium scooter comfort with Italian flair.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Piaggio Jetstream Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">300cc Liquid-Cooled:</strong> <span class="highlight-value">Premium maxi-scooter</span></li>
<li class="highlight-item"><strong class="highlight-label">Italian Design:</strong> <span class="highlight-value">Sophisticated styling</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfort Focus:</strong> <span class="highlight-value">Premium seating</span></li>
<li class="highlight-item"><strong class="highlight-label">Storage Generous:</strong> <span class="highlight-value">Practical capacity</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Piaggio Jetstream Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Italian Elegance:</strong> <span class="pro-description">Sophisticated styling</span></li>
<li class="pro-item"><strong class="pro-label">Premium Comfort:</strong> <span class="pro-description">Wide seat luxury</span></li>
<li class="pro-item"><strong class="pro-label">Practical Storage:</strong> <span class="pro-description">Generous capacity</span></li>
<li class="pro-item"><strong class="pro-label">Modern Tech:</strong> <span class="pro-description">Feature-rich equipped</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Efficiency:</strong> <span class="pro-description">25-28 km/l good</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Piaggio Jetstream Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">৳5,20,000 investment</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Premium brand maintenance</span></li>
<li class="con-item"><strong class="con-label">Weight Heavy:</strong> <span class="con-description">210kg substantial</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Piaggio Jetstream Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳5,20,000 - Premium scooter</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate - ৳13-15 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">25-28 km/l good</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">300cc liquid-cooled single</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">26 bhp refined smooth</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">210kg premium substantial</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Piaggio Jetstream Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Design Premium:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Italian sophisticated</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Premium luxury</span></li>
<li class="rating-item"><strong class="rating-label">Practicality:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Storage generous</span></li>
<li class="rating-item"><strong class="rating-label">Technology:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Feature-rich</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- ৳5,20,000 premium</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Piaggio Jetstream?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳5,20,000, the Jetstream is perfect for urban sophisticates seeking premium Italian design, exceptional comfort, generous storage, modern technology, and the Piaggio heritage for discerning scooter enthusiasts valuing elegance and practicality.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For premium scooter seekers</span></p>
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
		return fmt.Errorf("error creating piaggio jetstream review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Piaggio Jetstream (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">পিয়াজিও জেটস্ট্রিম রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম ম্যাক্সি-স্কুটার ইতালিয়ান ডিজাইন</h1>
<p class="summary-text">পিয়াজিও জেটস্ট্রিম একটি 300cc তরল-শীতলকৃত প্রিমিয়াম ম্যাক্সি-স্কুটার যা ইতালিয়ান ডিজাইনকে আধুনিক বৈশিষ্ট্যের সাথে একত্রিত করে। ৳5,20,000 টাকায় মূল্যায়িত, এটি পরিশীলিত স্টাইলিং এবং আরামদায়ক সিটিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">পিয়াজিও জেটস্ট্রিম মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">300cc তরল-শীতলকৃত:</strong> <span class="highlight-value">প্রিমিয়াম ম্যাক্সি-স্কুটার</span></li>
<li class="highlight-item"><strong class="highlight-label">ইতালিয়ান ডিজাইন:</strong> <span class="highlight-value">পরিশীলিত স্টাইলিং</span></li>
<li class="highlight-item"><strong class="highlight-label">আরাম ফোকাস:</strong> <span class="highlight-value">প্রিমিয়াম সিটিং</span></li>
<li class="highlight-item"><strong class="highlight-label">স্টোরেজ উদার:</strong> <span class="highlight-value">ব্যবহারিক ক্ষমতা</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">পিয়াজিও জেটস্ট্রিম সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ইতালিয়ান মার্জিত:</strong> <span class="pro-description">পরিশীলিত স্টাইলিং</span></li>
<li class="pro-item"><strong class="pro-label">প্রিমিয়াম আরাম:</strong> <span class="pro-description">প্রশস্ত আসন বিলাসিতা</span></li>
<li class="pro-item"><strong class="pro-label">ব্যবহারিক স্টোরেজ:</strong> <span class="pro-description">উদার ক্ষমতা</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক প্রযুক্তি:</strong> <span class="pro-description">বৈশিষ্ট্য-সমৃদ্ধ সজ্জিত</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি দক্ষতা:</strong> <span class="pro-description">25-28 km/l ভালো</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">পিয়াজিও জেটস্ট্রিম অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳5,20,000 বিনিয়োগ</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">প্রিমিয়াম ব্র্যান্ড রক্ষণাবেক্ষণ</span></li>
<li class="con-item"><strong class="con-label">ওজন ভারী:</strong> <span class="con-description">210kg উল্লেখযোগ্য</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: পিয়াজিও জেটস্ট্রিম কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳5,20,000 টাকায়, জেটস্ট্রিম শহুরে পরিশীলিতদের জন্য নিখুঁত যারা প্রিমিয়াম ইতালিয়ান ডিজাইন, ব্যতিক্রমী আরাম এবং পিয়াজিও হেরিটেজ চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- প্রিমিয়াম স্কুটার সন্ধানকারীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Piaggio Jetstream already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Piaggio Jetstream\n")

	return nil
}

// GetName returns the seeder name
func (s *PiaggioJetstreamReviewBatch19) GetName() string {
	return "PiaggioJetstreamReviewBatch19"
}
