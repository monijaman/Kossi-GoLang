package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// HondaCB500FReviewBatch19 handles seeding Honda CB500F product review and translation
type HondaCB500FReviewBatch19 struct {
	BaseSeeder
}

// NewHondaCB500FReviewBatch19Seeder creates a new HondaCB500FReviewBatch19
func NewHondaCB500FReviewBatch19Seeder() *HondaCB500FReviewBatch19 {
	return &HondaCB500FReviewBatch19{BaseSeeder: BaseSeeder{name: "Honda CB500F Batch19 Review"}}
}

// Seed implements the Seeder interface
func (s *HondaCB500FReviewBatch19) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found, please run user seeder first: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "Honda CB500F").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("honda cb500f product not found, please run product seeder first")
		}
		return fmt.Errorf("error finding honda cb500f product: %w", err)
	}

	var existingReview models.ProductReviewModel
	err := db.Where("product_id = ? AND user_id = ?", product.ID, adminUser.ID).First(&existingReview).Error
	if err == nil {
		fmt.Printf("   ℹ️  Review for Honda CB500F already exists, skipping\n")
		return nil
	}

	englishReview := `<article class="product-review-wrapper">
<header class="review-summary card">
<h1 class="review-main-title">Honda CB500F Review Bangladesh 2024 - Mid-Weight Sport Naked Versatile Master</h1>
<p class="summary-text">The Honda CB500F is a 500cc liquid-cooled parallel twin versatile sport-tourer combining practical ergonomics with exciting performance. Priced around ৳7,45,000, it delivers balanced power delivery, excellent fuel efficiency, upright comfortable positioning, comprehensive ABS, and Honda reliability for all-rounder riders seeking daily commute to weekend adventure capability.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">Honda CB500F Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">500cc Liquid-Cooled:</strong> <span class="highlight-value">Mid-weight balanced power</span></li>
<li class="highlight-item"><strong class="highlight-label">Sport-Tourer DNA:</strong> <span class="highlight-value">Versatile all-rounder</span></li>
<li class="highlight-item"><strong class="highlight-label">Comfort Ergonomics:</strong> <span class="highlight-value">Upright positioning</span></li>
<li class="highlight-item"><strong class="highlight-label">Fuel Efficiency:</strong> <span class="highlight-value">22-25 km/l good</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">Honda CB500F Pros & Advantages</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">Balanced Performance:</strong> <span class="pro-description">500cc versatile</span></li>
<li class="pro-item"><strong class="pro-label">Comfort Focus:</strong> <span class="pro-description">Upright riding</span></li>
<li class="pro-item"><strong class="pro-label">ABS Equipped:</strong> <span class="pro-description">Advanced safety</span></li>
<li class="pro-item"><strong class="pro-label">Practical Design:</strong> <span class="pro-description">All-rounder capable</span></li>
<li class="pro-item"><strong class="pro-label">Honda Reliability:</strong> <span class="pro-description">Proven dependability</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">Honda CB500F Cons & Disadvantages</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">Moderate Power:</strong> <span class="con-description">47 bhp accessible</span></li>
<li class="con-item"><strong class="con-label">Premium Price:</strong> <span class="con-description">৳7,45,000 investment</span></li>
<li class="con-item"><strong class="con-label">Fuel Tank:</strong> <span class="con-description">189mm ground clearance moderate</span></li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">Honda CB500F Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong class="value-label">Price Range:</strong> <span class="value-amount">৳7,45,000 - Mid-weight sport</span></p>
<p class="value-item"><strong class="value-label">Running Cost:</strong> <span class="value-amount">Low - ৳11-13 per km</span></p>
<p class="value-item"><strong class="value-label">Fuel Efficiency:</strong> <span class="value-amount">22-25 km/l good</span></p>
<p class="value-item"><strong class="value-label">Engine Type:</strong> <span class="value-amount">500cc liquid-cooled parallel twin</span></p>
<p class="value-item"><strong class="value-label">Power Output:</strong> <span class="value-amount">47 bhp balanced smooth</span></p>
<p class="value-item"><strong class="value-label">Kerb Weight:</strong> <span class="value-amount">189kg lightweight</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">Honda CB500F Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong class="rating-label">Versatility:</strong> <span class="rating-score">4.8</span> <span class="rating-note">- All-rounder champion</span></li>
<li class="rating-item"><strong class="rating-label">Comfort:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- Upright excellent</span></li>
<li class="rating-item"><strong class="rating-label">Performance:</strong> <span class="rating-score">4.6</span> <span class="rating-note">- 500cc balanced</span></li>
<li class="rating-item"><strong class="rating-label">Safety:</strong> <span class="rating-score">4.7</span> <span class="rating-note">- ABS equipped</span></li>
<li class="rating-item"><strong class="rating-label">Value:</strong> <span class="rating-score">4.5</span> <span class="rating-note">- ৳7,45,000 solid</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy Honda CB500F?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">Overall Rating:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">At ৳7,45,000, the CB500F is ideal for all-rounder riders seeking balanced 500cc performance, comfortable upright ergonomics, excellent versatility from daily commute to weekend adventures, comprehensive ABS safety, and Honda's proven reliability.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For versatile all-rounder riders</span></p>
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
		return fmt.Errorf("error creating honda cb500f review: %w", err)
	}

	fmt.Printf("   ✅ Created English review for Honda CB500F (Rating: 4.7/5.0)\n")

	banglaReview := `<article class="product-review-wrapper bangla-content">
<header class="review-summary card">
<h1 class="review-main-title">হোন্ডা CB500F রিভিউ বাংলাদেশ 2024 - মিড-ওয়েট স্পোর্ট নেকেড বহুমুখী মাস্টার</h1>
<p class="summary-text">হোন্ডা CB500F একটি 500cc তরল-শীতলকৃত প্যারালাল টুইন বহুমুখী স্পোর্ট-ট্যুরার যা ব্যবহারিক এরগোনমিক্সকে উত্তেজনাপূর্ণ পারফরম্যান্সের সাথে একত্রিত করে। ৳7,45,000 টাকায় মূল্যায়িত, এটি ভারসাম্যপূর্ণ শক্তি প্রদান এবং চমৎকার জ্বালানি দক্ষতা প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">হোন্ডা CB500F মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong class="highlight-label">500cc তরল-শীতলকৃত:</strong> <span class="highlight-value">মিড-ওয়েট ভারসাম্যপূর্ণ শক্তি</span></li>
<li class="highlight-item"><strong class="highlight-label">স্পোর্ট-ট্যুরার DNA:</strong> <span class="highlight-value">বহুমুখী সর্ব-উদ্দেশ্য</span></li>
<li class="highlight-item"><strong class="highlight-label">আরাম এরগোনমিক্স:</strong> <span class="highlight-value">সোজা পজিশনিং</span></li>
<li class="highlight-item"><strong class="highlight-label">জ্বালানি দক্ষতা:</strong> <span class="highlight-value">22-25 km/l ভালো</span></li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">হোন্ডা CB500F সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong class="pro-label">ভারসাম্যপূর্ণ পারফরম্যান্স:</strong> <span class="pro-description">500cc বহুমুখী</span></li>
<li class="pro-item"><strong class="pro-label">আরাম ফোকাস:</strong> <span class="pro-description">সোজা রাইডিং</span></li>
<li class="pro-item"><strong class="pro-label">ABS সজ্জিত:</strong> <span class="pro-description">উন্নত নিরাপত্তা</span></li>
<li class="pro-item"><strong class="pro-label">ব্যবহারিক ডিজাইন:</strong> <span class="pro-description">সর্ব-উদ্দেশ্য সক্ষম</span></li>
<li class="pro-item"><strong class="pro-label">হোন্ডা নির্ভরযোগ্যতা:</strong> <span class="pro-description">প্রমাণিত নির্ভরযোগ্যতা</span></li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">হোন্ডা CB500F অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong class="con-label">মধ্যম শক্তি:</strong> <span class="con-description">47 bhp অ্যাক্সেসযোগ্য</span></li>
<li class="con-item"><strong class="con-label">প্রিমিয়াম মূল্য:</strong> <span class="con-description">৳7,45,000 বিনিয়োগ</span></li>
<li class="con-item"><strong class="con-label">জ্বালানি ট্যাংক:</strong> <span class="con-description">মধ্যম ক্ষমতা</span></li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: হোন্ডা CB500F কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong class="score-label">সামগ্রিক রেটিং:</strong> <span class="score-value">4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p class="verdict-description">৳7,45,000 টাকায়, CB500F সর্ব-উদ্দেশ্য রাইডারদের জন্য আদর্শ যারা ভারসাম্যপূর্ণ 500cc পারফরম্যান্স, আরামদায়ক সোজা এরগোনমিক্স এবং দৈনিক যাতায়াত থেকে সপ্তাহান্তের অ্যাডভেঞ্চার ক্ষমতা চান।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong class="recommendation-label">সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- বহুমুখী সর্ব-উদ্দেশ্য রাইডারদের জন্য</span></p>
</div>
</section>
</article>`

	var existingTranslation models.ProductReviewTranslationModel
	err = db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existingTranslation).Error
	if err == nil {
		fmt.Printf("   ℹ️  Bangla translation for Honda CB500F already exists\n")
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

	fmt.Printf("   ✅ Created Bangla translation for Honda CB500F\n")

	return nil
}

// GetName returns the seeder name
func (s *HondaCB500FReviewBatch19) GetName() string {
	return "HondaCB500FReviewBatch19"
}
