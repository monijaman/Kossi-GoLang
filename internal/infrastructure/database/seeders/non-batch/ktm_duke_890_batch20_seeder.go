package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type KTMDuke890ReviewBatch20 struct {
	BaseSeeder
}

func NewKTMDuke890ReviewBatch20Seeder() *KTMDuke890ReviewBatch20 {
	return &KTMDuke890ReviewBatch20{BaseSeeder: BaseSeeder{name: "KTM Duke 890 Batch20 Review"}}
}

func (s *KTMDuke890ReviewBatch20) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "KTM Duke 890").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("ktm duke 890 product not found")
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
<h1 class="review-main-title">KTM Duke 890 Review Bangladesh 2024 - Aggressive Naked Street Performance Bike</h1>
<p class="summary-text">The KTM Duke 890 is an 889cc liquid-cooled aggressive naked sportbike representing raw extreme street performance. Priced around ৳15,50,000, it delivers 105 bhp raw power, sharp aggressive handling, track-capable performance, advanced technology, and adrenaline-fueled character for aggressive riders.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">KTM Duke 890 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>889cc Liquid-Cooled:</strong> Twin aggressive</li>
<li class="highlight-item"><strong>105 bhp Power:</strong> Extreme street performance</li>
<li class="highlight-item"><strong>Sharp Handling:</strong> Track-capable aggressive</li>
<li class="highlight-item"><strong>Advanced Tech:</strong> Performance systems</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">KTM Duke 890 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>105 bhp Raw:</strong> Extreme 889cc twin power</li>
<li class="pro-item"><strong>Sharp Handling:</strong> Track-capable aggressive control</li>
<li class="pro-item"><strong>Aggressive Design:</strong> Naked streetbike attitude</li>
<li class="pro-item"><strong>Advanced Tech:</strong> Performance electronics</li>
<li class="pro-item"><strong>Track-Capable:</strong> Street-legal performance</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">KTM Duke 890 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Raw Power:</strong> 105 bhp demanding</li>
<li class="con-item"><strong>Fuel Consumption:</strong> 18-20 km/l average</li>
<li class="con-item"><strong>Premium Price:</strong> ৳15,50,000 performance</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">KTM Duke 890 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳15,50,000 - Aggressive naked premium</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>889cc liquid-cooled twin</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>105 bhp raw</span></p>
<p class="value-item"><strong>Torque:</strong> <span>98 nm aggressive</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>189kg nimble</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>18-20 km/l average</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">KTM Duke 890 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Power:</strong> <span>4.9</span> - 105 bhp raw</li>
<li class="rating-item"><strong>Handling:</strong> <span>4.8</span> - Sharp track-ready</li>
<li class="rating-item"><strong>Design:</strong> <span>4.7</span> - Aggressive naked</li>
<li class="rating-item"><strong>Technology:</strong> <span>4.6</span> - Advanced systems</li>
<li class="rating-item"><strong>Value:</strong> <span>4.5</span> - ৳15,50,000</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy KTM Duke 890?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳15,50,000, the Duke 890 is the aggressive choice for thrill-seekers demanding extreme 105 bhp raw power, sharp track-capable handling, aggressive naked design, and street-legal performance for dynamic urban adventure.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For aggressive street riders</span></p>
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
<h1 class="review-main-title">কেটিএম ডিউক 890 রিভিউ বাংলাদেশ 2024 - আক্রমণাত্মক নেকেড রাস্তা পারফরম্যান্স বাইক</h1>
<p class="summary-text">কেটিএম ডিউক 890 একটি 889cc লিকুইড-কুল্ড আক্রমণাত্মক নেকেড স্পোর্টবাইক যা কাঁচা চরম রাস্তা পারফরম্যান্স প্রতিনিধিত্ব করে। ৳15,50,000 টাকায় মূল্যায়িত, এটি 105 bhp কাঁচা শক্তি এবং তীক্ষ্ণ আক্রমণাত্মক হ্যান্ডলিং প্রদান করে।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">কেটিএম ডিউক 890 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>889cc লিকুইড-কুল্ড:</strong> টুইন আক্রমণাত্মক</li>
<li class="highlight-item"><strong>105 bhp শক্তি:</strong> চরম রাস্তা পারফরম্যান্স</li>
<li class="highlight-item"><strong>তীক্ষ্ণ হ্যান্ডলিং:</strong> ট্র্যাক-সক্ষম আক্রমণাত্মক</li>
<li class="highlight-item"><strong>উন্নত প্রযুক্তি:</strong> পারফরম্যান্স সিস্টেম</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">কেটিএম ডিউক 890 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>105 bhp কাঁচা:</strong> চরম 889cc টুইন শক্তি</li>
<li class="pro-item"><strong>তীক্ষ্ণ হ্যান্ডলিং:</strong> ট্র্যাক-সক্ষম আক্রমণাত্মক নিয়ন্ত্রণ</li>
<li class="pro-item"><strong>আক্রমণাত্মক ডিজাইন:</strong> নেকেড রাস্তা মনোভাব</li>
<li class="pro-item"><strong>উন্নত প্রযুক্তি:</strong> পারফরম্যান্স ইলেকট্রনিক্স</li>
<li class="pro-item"><strong>ট্র্যাক-সক্ষম:</strong> রাস্তা-বৈধ কর্মক্ষমতা</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">কেটিএম ডিউক 890 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>কাঁচা শক্তি:</strong> 105 bhp দাবিপূর্ণ</li>
<li class="con-item"><strong>জ্বালানি ব্যবহার:</strong> 18-20 km/l গড়</li>
<li class="con-item"><strong>প্রিমিয়াম মূল্য:</strong> ৳15,50,000 কর্মক্ষমতা</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: কেটিএম ডিউক 890 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳15,50,000 টাকায়, ডিউক 890 রোমাঞ্চ-চাহনাকারীদের জন্য আক্রমণাত্মক পছন্দ যারা চরম 105 bhp কাঁচা শক্তি এবং তীক্ষ্ণ ট্র্যাক-সক্ষম হ্যান্ডলিং চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- আক্রমণাত্মক রাস্তা চালকদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created KTM Duke 890\n")
	return nil
}

func (s *KTMDuke890ReviewBatch20) GetName() string {
	return "KTMDuke890ReviewBatch20"
}
