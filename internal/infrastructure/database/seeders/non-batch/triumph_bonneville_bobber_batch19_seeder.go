package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TriumphBonnevilleBobberReviewBatch19 handles seeding Triumph Bonneville Bobber product review and translation
type TriumphBonnevilleBobberReviewBatch19 struct {
	BaseSeeder
}

// NewTriumphBonnevilleBobberReviewBatch19Seeder creates a new TriumphBonnevilleBobberReviewBatch19
func NewTriumphBonnevilleBobberReviewBatch19Seeder() *TriumphBonnevilleBobberReviewBatch19 {
	return &TriumphBonnevilleBobberReviewBatch19{BaseSeeder: BaseSeeder{name: "Triumph Bonneville Bobber Batch19 Review"}}
}

// Seed implements the Seeder interface
func (s *TriumphBonnevilleBobberReviewBatch19) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Triumph Bonneville Bobber").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("triumph bonneville bobber product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding triumph bonneville bobber product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Triumph Bonneville Bobber already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Triumph Bonneville Bobber Review Bangladesh 2024 - British Modern Bobber Masterpiece</h1>
<p class="summary-text">The Triumph Bonneville Bobber is a 1200cc liquid-cooled parallel twin modern bobber combining classic styling with contemporary performance. Priced around ৳16,75,000, it delivers iconic Triumph character, minimalist design philosophy, comfortable mid-set controls, sophisticated ABS/traction control, and the legendary British engineering for purists seeking authentic bobber soul with modern reliability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Triumph Bonneville Bobber Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">1200cc Liquid-Cooled:</strong> <span class="highlight-value">Classic parallel twin</span></li>
<li class="highlight-item"><strong class="highlight-label">Modern Bobber Design:</strong> <span class="highlight-value">Minimalist authentic</span></li>
<li class="highlight-item"><strong class="highlight-label">British Heritage:</strong> <span class="highlight-value">Iconic Triumph DNA</span></li>
<li class="highlight-item"><strong class="highlight-label">Contemporary Tech:</strong> <span class="highlight-value">ABS/traction control</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Triumph Bonneville Bobber Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Bobber Soul:</strong> <span class="pro-description">Authentic minimalist</span></li>
<li class="pro-item"><strong class="pro-label">1200cc Power:</strong> <span class="pro-description">96 bhp parallel twin</span></li>
<li class="pro-item"><strong class="pro-label">British Craftsmanship:</strong> <span class="pro-description">Triumph heritage</span></li>
<li class="pro-item"><strong class="pro-label">Smooth Torque:</strong> <span class="pro-description">110 nm cruiser character</span></li>
<li class="pro-item"><strong class="pro-label">Modern Features:</strong> <span class="pro-description">ABS/TC equipped</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Triumph Bonneville Bobber Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">৳16,75,000 investment</span></li>
<li class="con-item"><strong class="con-label">Fuel Economy:</strong> <span class="con-description">18-20 km/l low</span></li>
<li class="con-item"><strong class="con-label">Service Cost:</strong> <span class="con-description">Premium brand maintenance</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Triumph Bonneville Bobber Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳16,75,000 - Premium bobber</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳15-18 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">18-20 km/l low</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">1200cc liquid-cooled parallel twin</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">96 bhp cruiser smooth</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">252kg minimalist design</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Triumph Bonneville Bobber Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Bobber Character:</strong> <span class="rating-score">4.9</span> <span class="rating-note">- Authentic minimalist</span></li>
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- 1200cc smooth</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- Cruiser friendly</span></li>
<li class="rating-item"><strong class="rating-label">British Quality:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Heritage crafted</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- ৳16,75,000 premium</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Triumph Bonneville Bobber?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳16,75,000, the Bonneville Bobber is perfect for bobber purists seeking authentic minimalist design, 1200cc parallel twin character, iconic Triumph heritage, and modern safety technology for soul-searching cruiser experiences.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For bobber purists</span></p>
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
		return fmt.Errorf("error creating triumph bonneville bobber review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Triumph Bonneville Bobber (Rating: 4.6/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">ট্রায়াম্ফ বনেভিল ববার রিভিউ বাংলাদেশ 2024 - ব্রিটিশ মডার্ন ববার মাস্টারপিস</h1>
<p class="summary-text">ট্রায়াম্ফ বনেভিল ববার একটি 1200cc তরল-শীতলকৃত প্যারালাল টুইন মডার্ন ববার যা ক্লাসিক স্টাইলিংকে সমসাময়িক পারফরম্যান্সের সাথে একত্রিত করে। ৳16,75,000 টাকায় মূল্যায়িত, এটি আইকনিক ট্রায়াম্ফ চরিত্র এবং ব্রিটিশ ইঞ্জিনিয়ারিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">ট্রায়াম্ফ বনেভিল ববার মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">1200cc তরল-শীতলকৃত:</strong> <span class="highlight-value">ক্লাসিক প্যারালাল টুইন</span></li>
<li class="highlight-item"><strong class="highlight-label">মডার্ন ববার ডিজাইন:</strong> <span class="highlight-value">মিনিমালিস্ট খাঁটি</span></li>
<li class="highlight-item"><strong class="highlight-label">ব্রিটিশ হেরিটেজ:</strong> <span class="highlight-value">আইকনিক ট্রায়াম্ফ DNA</span></li>
<li class="highlight-item"><strong class="highlight-label">সমসাময়িক প্রযুক্তি:</strong> <span class="highlight-value">ABS/ট্র্যাকশন কন্ট্রোল</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">ট্রায়াম্ফ বনেভিল ববার সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ববার আত্মা:</strong> <span class="pro-description">খাঁটি মিনিমালিস্ট</span></li>
<li class="pro-item"><strong class="pro-label">1200cc শক্তি:</strong> <span class="pro-description">96 bhp প্যারালাল টুইন</span></li>
<li class="pro-item"><strong class="pro-label">ব্রিটিশ কারুশিল্প:</strong> <span class="pro-description">ট্রায়াম্ফ হেরিটেজ</span></li>
<li class="pro-item"><strong class="pro-label">মসৃণ টর্ক:</strong> <span class="pro-description">110 nm ক্রুজার চরিত্র</span></li>
<li class="pro-item"><strong class="pro-label">আধুনিক বৈশিষ্ট্য:</strong> <span class="pro-description">ABS/TC সজ্জিত</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">ট্রায়াম্ফ বনেভিল ববার অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳16,75,000 বিনিয়োগ</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি অর্থনীতি:</strong> <span class="con-description">18-20 km/l কম</span></li>
<li class="con-item"><strong class="con-label">সেবা খরচ:</strong> <span class="con-description">প্রিমিয়াম ব্র্যান্ড রক্ষণাবেক্ষণ</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: ট্রায়াম্ফ বনেভিল ববার কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.6/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳16,75,000 টাকায়, বনেভিল ববার ববার পিউরিস্টদের জন্য নিখুঁত যারা খাঁটি মিনিমালিস্ট ডিজাইন এবং আইকনিক ট্রায়াম্ফ হেরিটেজ চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- ববার পিউরিস্টদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Triumph Bonneville Bobber already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Triumph Bonneville Bobber\n")

	return nil
}

// GetName returns the seeder name
func (s *TriumphBonnevilleBobberReviewBatch19) GetName() string {
	return "TriumphBonnevilleBobberReviewBatch19"
}
