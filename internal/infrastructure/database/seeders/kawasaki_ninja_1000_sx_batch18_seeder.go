package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// KawasakiNinja1000SXReviewBatch18 handles seeding Kawasaki Ninja 1000 SX product review and translation
type KawasakiNinja1000SXReviewBatch18 struct {
	BaseSeeder
}

// NewKawasakiNinja1000SXReviewBatch18Seeder creates a new KawasakiNinja1000SXReviewBatch18
func NewKawasakiNinja1000SXReviewBatch18Seeder() *KawasakiNinja1000SXReviewBatch18 {
	return &KawasakiNinja1000SXReviewBatch18{BaseSeeder: BaseSeeder{name: "Kawasaki Ninja 1000 SX Batch18 Review"}}
}

// Seed implements the Seeder interface
func (s *KawasakiNinja1000SXReviewBatch18) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Kawasaki Ninja 1000 SX").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("kawasaki ninja 1000 sx product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding kawasaki ninja 1000 sx product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Kawasaki Ninja 1000 SX already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Kawasaki Ninja 1000 SX Review Bangladesh 2024 - Middleweight Sport-Tourer Precision</h1>
<p class="summary-text">The Kawasaki Ninja 1000 SX is a 1000cc super-sport tourer combining litre-class performance with practical touring capability and sophisticated electronics. Priced around ৳18,20,000, it delivers parallel-twin thrust, advanced rider aids, long-distance comfort, and the perfect balance between sportbike aggression and touring practicality for performance touring enthusiasts.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Kawasaki Ninja 1000 SX Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">1000cc Parallel Twin:</strong> <span class="highlight-value">Litre-class performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Sport-Touring DNA:</strong> <span class="highlight-value">Track and tour capable</span></li>
<li class="highlight-item"><strong class="highlight-label">Advanced Electronics:</strong> <span class="highlight-value">Rider assists equipped</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfort Focus:</strong> <span class="highlight-value">Long-distance capable</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Kawasaki Ninja 1000 SX Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Litre Performance:</strong> <span class="pro-description">1000cc acceleration</span></li>
<li class="pro-item"><strong class="pro-label">Sport-Touring Blend:</strong> <span class="pro-description">Track and tour ready</span></li>
<li class="pro-item"><strong class="pro-label">Cornering ABS:</strong> <span class="pro-description">Advanced safety tech</span></li>
<li class="pro-item"><strong class="pro-label">Comfort Seat:</strong> <span class="pro-description">Long-distance luxury</span></li>
<li class="pro-item"><strong class="pro-label">Storage:</strong> <span class="pro-description">Practical luggage mounting</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Kawasaki Ninja 1000 SX Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Price:</strong> <span class="con-description">৳18,20,000 premium segment</span></li>
<li class="con-item"><strong class="con-label">Fuel Consumption:</strong> <span class="con-description">19-23 km/l low</span></li>
<li class="con-item"><strong class="con-label">Maintenance Cost:</strong> <span class="con-description">Sport-tourer expensive</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Kawasaki Ninja 1000 SX Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳18,20,000 - Premium sport-tourer</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">High - ৳16-20 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">19-23 km/l low</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">1000cc liquid-cooled parallel twin</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">143 bhp with touring torque</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">238kg equipped tourer</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Kawasaki Ninja 1000 SX Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Litre-class acceleration</span></li>
<li class="rating-item"><strong class="rating-label">Touring Comfort:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Long-distance capable</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Precise responsive</span></li>
<li class="rating-item"><strong class="rating-label">Technology:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Advanced systems</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Premium investment</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Kawasaki Ninja 1000 SX?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳18,20,000, the Ninja 1000 SX is ideal for performance touring enthusiasts wanting litre-class acceleration with sophisticated sport-touring capability, advanced technology, and practical long-distance comfort for global adventures.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance touring enthusiasts</span></p>
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
		return fmt.Errorf("error creating kawasaki ninja 1000 sx review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Kawasaki Ninja 1000 SX (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">কাওয়াসাকি নিনজা 1000 SX রিভিউ বাংলাদেশ 2024 - মিডলওয়েট স্পোর্ট-ট্যুরার নির্ভুলতা</h1>
<p class="summary-text">কাওয়াসাকি নিনজা 1000 SX একটি 1000cc সুপার-স্পোর্ট ট্যুরার যা লিটার-ক্লাস পারফরম্যান্সকে ব্যবহারিক ট্যুরিং ক্ষমতার সাথে একত্রিত করে। ৳18,20,000 টাকায় মূল্যায়িত, এটি সমান্তরাল-টুইন থ্রাস্ট, উন্নত রাইডার সহায়তা এবং দীর্ঘ-দূরত্ব আরাম প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">কাওয়াসাকি নিনজা 1000 SX মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">1000cc সমান্তরাল টুইন:</strong> <span class="highlight-value">লিটার-ক্লাস পারফরম্যান্স</span></li>
<li class="highlight-item"><strong class="highlight-label">স্পোর্ট-ট্যুরিং DNA:</strong> <span class="highlight-value">ট্র্যাক এবং ট্যুর সক্ষম</span></li>
<li class="highlight-item"><strong class="highlight-label">উন্নত ইলেকট্রনিক্স:</strong> <span class="highlight-value">রাইডার সহায়তা সজ্জিত</span></li>
<li class="highlight-item"><strong class="highlight-label">আরাম ফোকাস:</strong> <span class="highlight-value">দীর্ঘ-দূরত্ব সক্ষম</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">কাওয়াসাকি নিনজা 1000 SX সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">লিটার পারফরম্যান্স:</strong> <span class="pro-description">1000cc ত্বরণ</span></li>
<li class="pro-item"><strong class="pro-label">স্পোর্ট-ট্যুরিং ব্লেন্ড:</strong> <span class="pro-description">ট্র্যাক এবং ট্যুর প্রস্তুত</span></li>
<li class="pro-item"><strong class="pro-label">কর্নারিং ABS:</strong> <span class="pro-description">উন্নত নিরাপত্তা প্রযুক্তি</span></li>
<li class="pro-item"><strong class="pro-label">আরাম আসন:</strong> <span class="pro-description">দীর্ঘ-দূরত্ব বিলাসিতা</span></li>
<li class="pro-item"><strong class="pro-label">স্টোরেজ:</strong> <span class="pro-description">ব্যবহারিক লাগেজ মাউন্টিং</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">কাওয়াসাকি নিনজা 1000 SX অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">৳18,20,000 প্রিমিয়াম বিভাগ</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি খরচ:</strong> <span class="con-description">19-23 km/l কম</span></li>
<li class="con-item"><strong class="con-label">রক্ষণাবেক্ষণ খরচ:</strong> <span class="con-description">স্পোর্ট-ট্যুরার ব্যয়বহুল</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: কাওয়াসাকি নিনজা 1000 SX কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳18,20,000 টাকায়, নিনজা 1000 SX পারফরম্যান্স ট্যুরিং উত্সাহীদের জন্য আদর্শ যারা লিটার-ক্লাস ত্বরণ এবং পরিশীলিত স্পোর্ট-ট্যুরিং ক্ষমতা চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পারফরম্যান্স ট্যুরিং উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Kawasaki Ninja 1000 SX already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Kawasaki Ninja 1000 SX\n")

	return nil
}

// GetName returns the seeder name
func (s *KawasakiNinja1000SXReviewBatch18) GetName() string {
	return "KawasakiNinja1000SXReviewBatch18"
}
