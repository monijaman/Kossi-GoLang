package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCRF450RallyReviewBatch18 handles seeding Honda CRF450 Rally product review and translation
type HondaCRF450RallyReviewBatch18 struct {
	BaseSeeder
}

// NewHondaCRF450RallyReviewBatch18Seeder creates a new HondaCRF450RallyReviewBatch18
func NewHondaCRF450RallyReviewBatch18Seeder() *HondaCRF450RallyReviewBatch18 {
	return &HondaCRF450RallyReviewBatch18{BaseSeeder: BaseSeeder{name: "Honda CRF450 Rally Batch18 Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCRF450RallyReviewBatch18) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CRF450 Rally").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda crf450 rally product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding honda crf450 rally product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CRF450 Rally already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CRF450 Rally Review Bangladesh 2024 - Premium Rally Adventure Weapon</h1>
<p class="summary-text">The Honda CRF450 Rally is a 450cc mid-weight rally-adventure bike combining rugged durability, advanced rally racing DNA, and versatile terrain capability. Priced around ৳15,80,000, it delivers aggressive off-road performance, liquid-cooled reliability, sophisticated electronics, and the perfect balance between on-road manners and extreme off-road prowess for rally adventurers and serious explorers.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CRF450 Rally Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">450cc Rally Engine:</strong> <span class="highlight-value">Rally-tuned performance</span></li>
<li class="highlight-item"><strong class="highlight-label">Liquid Cooling:</strong> <span class="highlight-value">Extreme conditions reliability</span></li>
<li class="highlight-item"><strong class="highlight-label">Rally Heritage:</strong> <span class="highlight-value">Paris-Dakar proven technology</span></li>
<li class="highlight-item"><strong class="highlight-label">Versatile Platform:</strong> <span class="highlight-value">On/off-road balance</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CRF450 Rally Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Rally DNA:</strong> <span class="pro-description">Proven motorsport pedigree</span></li>
<li class="pro-item"><strong class="pro-label">Off-Road Mastery:</strong> <span class="pro-description">Sand/mud/rock capable</span></li>
<li class="pro-item"><strong class="pro-label">Electronic Aids:</strong> <span class="pro-description">TCS and ABS systems</span></li>
<li class="pro-item"><strong class="pro-label">Fuel Tank:</strong> <span class="pro-description">20L extra-large range</span></li>
<li class="pro-item"><strong class="pro-label">Build Quality:</strong> <span class="pro-description">Honda durability legend</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CRF450 Rally Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">High Price:</strong> <span class="con-description">৳15,80,000 premium segment</span></li>
<li class="con-item"><strong class="con-label">Heavy Weight:</strong> <span class="con-description">190kg substantial</span></li>
<li class="con-item"><strong class="con-label">Fuel Consumption:</strong> <span class="con-description">22-26 km/l moderate</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CRF450 Rally Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳15,80,000 - Premium adventure rally</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Moderate-high - ৳12-16 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">22-26 km/l moderate</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">450cc liquid-cooled single</span></li>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">40 bhp with rally torque</span></li>
<p class="value-item"><strong class="value-label">Wheel Travel:</strong> <span class="value-amount">250mm suspension adventure-ready</span></li>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CRF450 Rally Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Off-Road:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- Rally-class terrain master</span></li>
<li class="rating-item"><strong class="rating-label">Engine:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Reliable liquid-cooled</span></li>
<li class="rating-item"><strong class="rating-label">Handling:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- Balanced agile</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.4</span> <span class="rating-note">- Rally focused</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.3</span> <span class="rating-note">- Premium investment</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CRF450 Rally?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳15,80,000, the CRF450 Rally is for serious rally adventurers and off-road explorers wanting proven Paris-Dakar technology, extreme terrain mastery, and the confidence of Honda's legendary reliability for expeditions across all conditions.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For rally adventure enthusiasts</span></p>
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
		return fmt.Errorf("error creating honda crf450 rally review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CRF450 Rally (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা CRF450 র্যালি রিভিউ বাংলাদেশ 2024 - প্রিমিয়াম র্যালি অ্যাডভেঞ্চার ওয়েপন</h1>
<p class="summary-text">হোন্ডা CRF450 র্যালি একটি 450cc মধ্য-ওজন র্যালি-অ্যাডভেঞ্চার বাইক যা মজবুত স্থায়িত্ব, উন্নত র্যালি রেসিং DNA এবং বহুমুখী ভূখণ্ড ক্ষমতা একত্রিত করে। ৳15,80,000 টাকায় মূল্যায়িত, এটি আক্রমণাত্মক অফ-রোড কর্মক্ষমতা, তরল-শীতলকৃত নির্ভরযোগ্যতা এবং পরিশীলিত ইলেকট্রনিক্স প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা CRF450 র্যালি মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">450cc র্যালি ইঞ্জিন:</strong> <span class="highlight-value">র্যালি-টিউনড পারফরম্যান্স</span></li>
<li class="highlight-item"><strong class="highlight-label">তরল শীতলন:</strong> <span class="highlight-value">চরম অবস্থা নির্ভরযোগ্যতা</span></li>
<li class="highlight-item"><strong class="highlight-label">র্যালি ঐতিহ্য:</strong> <span class="highlight-value">পারিস-ডাকার প্রমাণিত প্রযুক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">বহুমুখী প্ল্যাটফর্ম:</strong> <span class="highlight-value">অন/অফ-রোড ভারসাম্য</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা CRF450 র্যালি সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">র্যালি DNA:</strong> <span class="pro-description">প্রমাণিত মোটরস্পোর্ট বংশানুক্রম</span></li>
<li class="pro-item"><strong class="pro-label">অফ-রোড মাস্টারি:</strong> <span class="pro-description">বালি/কাদা/শিলা সক্ষম</span></li>
<li class="pro-item"><strong class="pro-label">ইলেকট্রনিক সহায়তা:</strong> <span class="pro-description">TCS এবং ABS সিস্টেম</span></li>
<li class="pro-item"><strong class="pro-label">জ্বালানি ট্যাঙ্ক:</strong> <span class="pro-description">20L অতিরিক্ত-বৃহৎ পরিসর</span></li>
<li class="pro-item"><strong class="pro-label">বিল্ড গুণমান:</strong> <span class="pro-description">হোন্ডা স্থায়িত্ব কিংবদন্তি</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা CRF450 র্যালি অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">উচ্চ মূল্য:</strong> <span class="con-description">৳15,80,000 প্রিমিয়াম বিভাগ</span></li>
<li class="con-item"><strong class="con-label">ভারী ওজন:</strong> <span class="con-description">190kg উল্লেখযোগ্য</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি খরচ:</strong> <span class="con-description">22-26 km/l মধ্যমা</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হোন্ডা CRF450 র্যালি কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳15,80,000 টাকায়, CRF450 র্যালি গুরুতর র্যালি অ্যাডভেঞ্চারারদের জন্য চরম ভূখণ্ড দক্ষতা এবং নির্ভরযোগ্যতার জন্য একটি প্রিমিয়াম বিনিয়োগ।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- র্যালি অ্যাডভেঞ্চার উত্সাহীদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CRF450 Rally already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda CRF450 Rally\n")

	return nil
}

// GetName returns the seeder name
func (s *HondaCRF450RallyReviewBatch18) GetName() string {
	return "HondaCRF450RallyReviewBatch18"
}
