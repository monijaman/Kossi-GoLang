package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type KTMDuke390Batch26 struct {
	BaseSeeder
}

func NewKTMDuke390Batch26Seeder() *KTMDuke390Batch26 {
	return &KTMDuke390Batch26{BaseSeeder: BaseSeeder{name: "KTM Duke 390 Batch26 Review"}}
}

func (s *KTMDuke390Batch26) Seed(db *gorm.DB) error {
	var adminUser models.UserModel
	if err := db.First(&adminUser, 1).Error; err != nil {
		return fmt.Errorf("admin user not found: %w", err)
	}

	var product models.ProductModel
	if err := db.Where("name = ?", "KTM Duke 390").First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("ktm duke 390 product not found")
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
<h1 class="review-main-title">KTM Duke 390 Review Bangladesh 2024 - Performance Beast Unleashed</h1>
<p class="summary-text">The KTM Duke 390 is an extreme performance motorcyclist's dream combining raw power and aggressive engineering. Priced around ৳5,95,000, it delivers 44 bhp power, lightweight design, refined handling, advanced features, and unmatched street riding thrills.</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">KTM Duke 390 Key Highlights</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>390cc Engine:</strong> Performance extreme</li>
<li class="highlight-item"><strong>44 bhp Power:</strong> Incredible responsive</li>
<li class="highlight-item"><strong>Lightweight Design:</strong> Agile powerful</li>
<li class="highlight-item"><strong>Advanced Features:</strong> Tech enabled modern</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">KTM Duke 390 Pros</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>Extreme Power:</strong> 44 bhp incredible</li>
<li class="pro-item"><strong>Lightweight Construction:</strong> Agile performance</li>
<li class="pro-item"><strong>Excellent Handling:</strong> Precision control</li>
<li class="pro-item"><strong>Advanced Technology:</strong> Modern features</li>
<li class="pro-item"><strong>Street Supremacy:</strong> Ultimate thrills</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">KTM Duke 390 Cons</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>Extreme Power:</strong> Requires skilled riding</li>
<li class="con-item"><strong>High Cost:</strong> Performance premium</li>
<li class="con-item"><strong>Limited Comfort:</strong> Sport focused positioning</li>
</ul>
</section>

<section class="value-analysis-section section card">
<h2 class="section-title value-title">KTM Duke 390 Price & Value Analysis</h2>
<div class="value-details">
<p class="value-item"><strong>Price Range:</strong> <span>৳5,95,000 - Performance beast</span></p>
<p class="value-item"><strong>Engine Type:</strong> <span>390cc single cylinder</span></p>
<p class="value-item"><strong>Power Output:</strong> <span>44 bhp incredible</span></p>
<p class="value-item"><strong>Torque:</strong> <span>35 nm powerful</span></p>
<p class="value-item"><strong>Kerb Weight:</strong> <span>149kg agile</span></p>
<p class="value-item"><strong>Fuel Efficiency:</strong> <span>40+ km/l sport</span></p>
</div>
</section>

<section class="performance-rating-section section">
<h2 class="section-title rating-title">KTM Duke 390 Performance Rating</h2>
<ul class="rating-list">
<li class="rating-item"><strong>Design:</strong> <span>4.8</span> - Aggressive extreme</li>
<li class="rating-item"><strong>Comfort:</strong> <span>4.5</span> - Sport positioned</li>
<li class="rating-item"><strong>Performance:</strong> <span>4.9</span> - 44 bhp incredible</li>
<li class="rating-item"><strong>Efficiency:</strong> <span>4.5</span> - 40+ km/l sport</li>
<li class="rating-item"><strong>Quality:</strong> <span>4.7</span> - KTM proven</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">Final Verdict: Should You Buy KTM Duke 390?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>Overall Rating:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>At ৳5,95,000, the KTM Duke 390 is perfect for extreme riders seeking unmatched 44 bhp performance, lightweight agility, advanced technology, and unforgettable street supremacy.</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>Recommended:</strong> <span class="badge badge-success">YES</span> <span class="recommendation-for">- For performance enthusiasts</span></p>
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
<h1 class="review-main-title">KTM ডিউক 390 রিভিউ বাংলাদেশ 2024 - পারফরম্যান্স বিস্ট আনলিশড</h1>
<p class="summary-text">KTM ডিউক 390 একটি চরম কর্মক্ষমতা মোটরসাইক্লিস্টের স্বপ্ন যা কাঁচা শক্তি এবং আক্রমণাত্মক প্রকৌশল একত্রিত করে। ৳5,95,000 টাকায় মূল্যায়িত, এটি 44 bhp শক্তি এবং হালকা ডিজাইন বৈশিষ্ট্যযুক্ত।</p>
</header>

<section class="key-highlights section">
<h2 class="section-title highlight-title">KTM ডিউক 390 মূল বৈশিষ্ট্য</h2>
<ul class="highlights-list">
<li class="highlight-item"><strong>390cc ইঞ্জিন:</strong> পারফরম্যান্স চরম</li>
<li class="highlight-item"><strong>44 bhp শক্তি:</strong> অবিশ্বাস্য প্রতিক্রিয়াশীল</li>
<li class="highlight-item"><strong>হালকা ডিজাইন:</strong> চতুর শক্তিশালী</li>
<li class="highlight-item"><strong>উন্নত বৈশিষ্ট্য:</strong> প্রযুক্তি সক্ষম আধুনিক</li>
</ul>
</section>

<section class="pros-section section">
<h2 class="section-title pros-title">KTM ডিউক 390 সুবিধা</h2>
<ul class="pros-list advantages-list">
<li class="pro-item"><strong>চরম শক্তি:</strong> 44 bhp অবিশ্বাস্য</li>
<li class="pro-item"><strong>হালকা নির্মাণ:</strong> চতুর পারফরম্যান্স</li>
<li class="pro-item"><strong>চমৎকার হ্যান্ডলিং:</strong> নির্ভুল নিয়ন্ত্রণ</li>
<li class="pro-item"><strong>উন্নত প্রযুক্তি:</strong> আধুনিক বৈশিষ্ট্য</li>
<li class="pro-item"><strong>স্ট্রিট শ্রেষ্ঠত্ব:</strong> চূড়ান্ত রোমাঞ্চ</li>
</ul>
</section>

<section class="cons-section section">
<h2 class="section-title cons-title">KTM ডিউক 390 অসুবিধা</h2>
<ul class="cons-list disadvantages-list">
<li class="con-item"><strong>চরম শক্তি:</strong> দক্ষ রাইডিং প্রয়োজন</li>
<li class="con-item"><strong>উচ্চ খরচ:</strong> পারফরম্যান্স প্রিমিয়াম</li>
<li class="con-item"><strong>সীমিত আরাম:</strong> স্পোর্ট দৃষ্টিভঙ্গি অবস্থান</li>
</ul>
</section>

<section class="final-verdict-section section card highlight">
<h2 class="section-title verdict-title">চূড়ান্ত সিদ্ধান্ত: KTM ডিউক 390 কিনবেন?</h2>
<div class="overall-rating">
<p class="overall-score"><strong>সামগ্রিক রেটিং:</strong> <span>4.7/5.0</span></p>
</div>
<div class="verdict-text">
<p>৳5,95,000 টাকায়, KTM ডিউক 390 চরম আরোহীদের জন্য নিখুঁত যারা অতুলনীয় 44 bhp কর্মক্ষমতা এবং চূড়ান্ত স্ট্রিট শ্রেষ্ঠত্ব চায়।</p>
</div>
<div class="recommendation">
<p class="recommendation-badge"><strong>সুপারিশ:</strong> <span class="badge badge-success">হ্যাঁ</span> <span class="recommendation-for">- পারফরম্যান্স উৎসাহীদের জন্য</span></p>
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

	fmt.Printf("   ✅ Created KTM Duke 390\n")
	return nil
}

func (s *KTMDuke390Batch26) GetName() string {
	return "KTMDuke390Batch26"
}
