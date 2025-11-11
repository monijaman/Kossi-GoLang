package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type PiaggioMP3500ReviewBatch20 struct {
	BaseSeeder
}

func NewPiaggioMP3500ReviewBatch20Seeder() *PiaggioMP3500ReviewBatch20 {
	return &PiaggioMP3500ReviewBatch20{BaseSeeder: BaseSeeder{name: "Piaggio MP3 500 Batch20 Review"}}
}

func (s *PiaggioMP3500ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Piaggio MP3 500").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("piaggio mp3 500 product not found")
		}
		return fmt.Errorf("error finding product: %w", err)
	}

	var existingReview models.ProductReviewModel
	if err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error; err == nil {
		fmt.Printf("   ℹ️  Review already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Piaggio MP3 500 Review Bangladesh 2024 - Three-Wheel Innovation Scooter</h1>
<p class="summary-text">The Piaggio MP3 500 is a 500cc liquid-cooled innovative three-wheel maxi-scooter representing revolutionary commuting design. Priced around ৳14,50,000, it delivers unique three-wheel stability, 42 bhp power, spacious storage, comfortable riding position, and innovative Italian engineering for riders seeking alternative fun commuting experience.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Piaggio MP3 500 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>500cc Liquid-Cooled:</strong> Maxi-scooter</li>
<li class="highlight-item"><strong>Three-Wheel Design:</strong> Unique stability</li>
<li class="highlight-item"><strong>42 bhp:</strong> Smooth practical</li>
<li class="highlight-item"><strong>Spacious Storage:</strong> Generous compartment</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Piaggio MP3 500 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Three-Wheel Stability:</strong> Unique safe design</li>
<li class="pro-item"><strong>Spacious Storage:</strong> Generous compartment</li>
<li class="pro-item"><strong>Comfortable Riding:</strong> Ergonomic position</li>
<li class="pro-item"><strong>Innovation:</strong> Italian design leadership</li>
<li class="pro-item"><strong>Fun Factor:</strong> Alternative commuting</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Piaggio MP3 500 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Fuel Economy:</strong> 22-26 km/l average</li>
<li class="con-item"><strong>Heavy Weight:</strong> 225kg substantial</li>
<li class="con-item"><strong>Premium Price:</strong> ৳14,50,000 specialized</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Piaggio MP3 500 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳14,50,000 - Premium three-wheel</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>500cc liquid-cooled single</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>42 bhp smooth</span></p>
<p class="value-item"><strong>Torque:</strong> <span>44 nm adequate</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>225kg substantial</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>22-26 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Piaggio MP3 500 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Stability:</strong> <span>4.8</span> - Three-wheel unique</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.7</span> - Spacious seating</li>
<li class="rating-item"><strong>Design Innovation:</strong> <span>4.7</span> - Revolutionary</li>
<li class="rating-item"><strong>Storage:</strong> <span>4.8</span> - Generous capacity</li>
<li class="rating-item"><strong>Value:</strong> <span>4.5</span> - ৳14,50,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Piaggio MP3 500?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳14,50,000, the MP3 500 is the innovative choice for riders seeking unique three-wheel stability, spacious storage, comfortable riding, fun commuting experience, and revolutionary Italian design innovation for alternative daily transport.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For innovative commuters</span></p>
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
		return fmt.Errorf("error creating review: %w", err)
	}

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">পিয়াগিও MP3 500 রিভিউ বাংলাদেশ 2024 - তিন-চাকার উদ্ভাবন স্কুটার</h1>
<p class="summary-text">পিয়াগিও MP3 500 একটি 500cc লিকুইড-কুল্ড উদ্ভাবনী তিন-চাকার মক্সি-স্কুটার যা বিপ্লবী কমিউটিং ডিজাইন প্রতিনিধিত্ব করে। ৳14,50,000 টাকায় মূল্যায়িত, এটি অনন্য তিন-চাকার স্থিতিশীলতা এবং আরামদায়ক রাইডিং অবস্থান প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">পিয়াগিও MP3 500 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>500cc লিকুইড-কুল্ড:</strong> মক্সি-স্কুটার</li>
<li class="highlight-item"><strong>তিন-চাকার ডিজাইন:</strong> অনন্য স্থিতিশীলতা</li>
<li class="highlight-item"><strong>42 bhp:</strong> মসৃণ ব্যবহারিক</li>
<li class="highlight-item"><strong>প্রশস্ত স্টোরেজ:</strong> উদার বগি</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">পিয়াগিও MP3 500 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>তিন-চাকার স্থিতিশীলতা:</strong> অনন্য নিরাপদ ডিজাইন</li>
<li class="pro-item"><strong>প্রশস্ত স্টোরেজ:</strong> উদার বগি</li>
<li class="pro-item"><strong>আরামদায়ক রাইডিং:</strong> এরগোনমিক অবস্থান</li>
<li class="pro-item"><strong>উদ্ভাবন:</strong> ইতালিয়ান ডিজাইন নেতৃত্ব</li>
<li class="pro-item"><strong>মজার ফ্যাক্টর:</strong> বিকল্প কমিউটিং</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">পিয়াগিও MP3 500 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>জ্বালানি অর্থনীতি:</strong> 22-26 km/l গড়</li>
<li class="con-item"><strong>ভারী ওজন:</strong> 225kg উল্লেখযোগ্য</li>
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳14,50,000 বিশেষায়িত</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: পিয়াগিও MP3 500 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳14,50,000 টাকায়, MP3 500 উদ্ভাবনী রাইডারদের জন্য উদ্ভাবনী পছন্দ যারা অনন্য তিন-চাকার স্থিতিশীলতা, প্রশস্ত স্টোরেজ এবং বিপ্লবী ইতালিয়ান ডিজাইন খুঁজছেন।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- উদ্ভাবনী কমিউটারদের জন্য</span></p>
</div>
</section>
</article>`

	translation := &models.ProductReviewTranslationModel{
		ProductReviewID: review.ID,
		Locale:          "bn",
		Reviews:         banglaReview,
	}

	if err := db.Create(translation).Error; err != nil {
		return fmt.Errorf("error creating translation: %w", err)
	}

	fmt.Printf("   ✅ Created Piaggio MP3 500\n")
	return nil
}

func (s *PiaggioMP3500ReviewBatch20) GetName() string {
	return "PiaggioMP3500ReviewBatch20"
}
