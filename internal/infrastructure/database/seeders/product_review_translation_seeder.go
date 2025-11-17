package seeders

import (
	"encoding/json"
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// ProductReviewTranslationSeeder handles seeding Bangla translations for product reviews
type ProductReviewTranslationSeeder struct {
	BaseSeeder
}

// NewProductReviewTranslationSeeder creates a new ProductReviewTranslationSeeder
func NewProductReviewTranslationSeeder() *ProductReviewTranslationSeeder {
	return &ProductReviewTranslationSeeder{BaseSeeder: BaseSeeder{name: "Product Review Translations"}}
}

// Seed implements the Seeder interface
func (prts *ProductReviewTranslationSeeder) Seed(db *gorm.DB) error {
	translations := []struct {
		ProductName       string
		BanglaReview      string
		AdditionalDetails map[string]interface{}
	}{
		{
			ProductName: "Honda CD 70",
			BanglaReview: `<div class="review-content">
<h3>সুবিধা</h3>
<ul>
<li>অত্যন্ত জ্বালানি সাশ্রয়ী - প্রতি লিটারে ৬০-৭০ কিলোমিটার মাইলেজ</li>
<li>খুবই সাশ্রয়ী এবং রক্ষণাবেক্ষণ সহজ</li>
<li>নির্ভরযোগ্য হোন্ডা ইঞ্জিন সুপ্রমাণিত ট্র্যাক রেকর্ড সহ</li>
<li>হালকা ওজনের এবং নতুনদের জন্য চালানো সহজ</li>
<li>কম রানিং খরচ</li>
<li>ভালো রিসেল ভ্যালু</li>
<li>শহরের ট্রাফিকে দৈনন্দিন যাতায়াতের জন্য নিখুঁত</li>
</ul>

<h3>অসুবিধা</h3>
<ul>
<li>হাইওয়েতে চালানোর জন্য কম শক্তিশালী</li>
<li>ন্যূনতম ফিচার সহ সাধারণ ডিজাইন</li>
<li>লম্বা রাইডারদের জন্য উপযুক্ত নয়</li>
<li>আধুনিক স্টাইলিং এর অভাব</li>
<li>কোনো ডিজিটাল ইন্সট্রুমেন্টেশন নেই</li>
<li>নতুন মডেলের তুলনায় গড় বিল্ড কোয়ালিটি</li>
<li>সীমিত টপ স্পিড</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=BanglaCD70Review",
					"https://www.youtube.com/watch?v=CD70BanglaTest",
				},
				"verdict":     "অতুলনীয় জ্বালানি সাশ্রয়ী সেরা বাজেট কমিউটার বাইক",
				"recommended": true,
			},
		},
		{
			ProductName: "Honda Livo",
			BanglaReview: `<div class="review-content">
<h3>সুবিধা</h3>
<ul>
<li>চমৎকার জ্বালানি দক্ষতা প্রতি লিটারে ৫৫-৬৫ কিলোমিটার</li>
<li>দৈনন্দিন যাতায়াতের জন্য আরামদায়ক রাইডিং পজিশন</li>
<li>আধুনিক এবং আকর্ষণীয় ডিজাইন</li>
<li>হোন্ডার সাধারণ ভালো বিল্ড কোয়ালিটি</li>
<li>মসৃণ এবং পরিশীলিত ইঞ্জিন</li>
<li>শহরে চালানোর জন্য পর্যাপ্ত শক্তি</li>
<li>কম রক্ষণাবেক্ষণ প্রয়োজন</li>
<li>রাইডার এবং পিলিয়ন উভয়ের জন্য আরামদায়ক সিট</li>
</ul>

<h3>অসুবিধা</h3>
<ul>
<li>প্রতিযোগীদের তুলনায় সামান্য কম শক্তিশালী</li>
<li>হাইওয়েতে গড় পারফরম্যান্স</li>
<li>বেসিক ইন্সট্রুমেন্টেশন ক্লাস্টার</li>
<li>সাসপেনশন আরো নরম হতে পারত</li>
<li>সীমিত স্টোরেজ বিকল্প</li>
<li>সিডি ৭০ এর তুলনায় বেশি দাম</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=LivoBanglaReview",
					"https://www.youtube.com/watch?v=HondaLivoBD",
				},
				"verdict":     "আধুনিক লুক সহ ভালো ভারসাম্যপূর্ণ কমিউটার",
				"recommended": true,
			},
		},
		{
			ProductName: "Honda CB125F",
			BanglaReview: `<div class="review-content">
<h3>সুবিধা</h3>
<ul>
<li>চমৎকার পাওয়ার ডেলিভারি সহ প্রাণবন্ত ১২৫সিসি ইঞ্জিন</li>
<li>চমৎকার জ্বালানি দক্ষতা প্রতি লিটারে ৫০-৬০ কিলোমিটার</li>
<li>স্পোর্টি এবং আক্রমণাত্মক স্টাইলিং</li>
<li>হোন্ডার কিংবদন্তী নির্ভরযোগ্যতা এবং বিল্ড কোয়ালিটি</li>
<li>শহর এবং হাইওয়ে উভয়ের জন্য আরামদায়ক</li>
<li>ভালো হ্যান্ডলিং এবং ম্যানুভারেবিলিটি</li>
<li>ডিজিটাল-অ্যানালগ ইন্সট্রুমেন্ট ক্লাস্টার</li>
<li>শক্তিশালী ব্রেকিং পারফরম্যান্স</li>
</ul>

<h3>অসুবিধা</h3>
<ul>
<li>প্রতিযোগীদের তুলনায় প্রিমিয়াম মূল্য</li>
<li>লম্বা রাইডের জন্য সিট আরো আরামদায়ক হতে পারত</li>
<li>বেস ভ্যারিয়েন্টে কোনো এবিএস অপশন নেই</li>
<li>গ্রাউন্ড ক্লিয়ারেন্স আরো ভালো হতে পারত</li>
<li>সীমিত আন্ডার-সিট স্টোরেজ</li>
<li>হেডলাইটের উজ্জ্বলতা উন্নত করা যেত</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=CB125FBangla",
					"https://www.youtube.com/watch?v=HondaCB125Review",
				},
				"verdict":     "দৈনন্দিন রাইডিং এর জন্য চমৎকার অল-রাউন্ডার",
				"recommended": true,
			},
		},
		{
			ProductName: "Honda CB Shine",
			BanglaReview: `<div class="review-content">
<h3>সুবিধা</h3>
<ul>
<li>পরিশীলিত এবং শক্তিশালী ১২৫সিসি ইঞ্জিন</li>
<li>চমৎকার বিল্ড কোয়ালিটি এবং প্রিমিয়াম অনুভূতি</li>
<li>অত্যন্ত আরামদায়ক রাইডিং পজিশন</li>
<li>দুর্দান্ত জ্বালানি দক্ষতা প্রতি লিটারে ৫৫-৬৫ কিলোমিটার</li>
<li>মসৃণ গিয়ার শিফট এবং ক্লাচ অপারেশন</li>
<li>ভারতীয় সড়ক পরিস্থিতির জন্য ভালো সাসপেনশন</li>
<li>নির্ভরযোগ্য হোন্ডা ব্র্যান্ড খ্যাতি</li>
<li>কম রক্ষণাবেক্ষণ খরচ</li>
<li>শক্তিশালী রিসেল ভ্যালু</li>
</ul>

<h3>অসুবিধা</h3>
<ul>
<li>রক্ষণশীল ডিজাইন তরুণ রাইডারদের কাছে আবেদনময় নাও হতে পারে</li>
<li>প্রিমিয়াম মূল্য</li>
<li>এলইডি লাইটের মতো আধুনিক ফিচার নেই (বেস ভ্যারিয়েন্টে)</li>
<li>চালাতে সবচেয়ে রোমাঞ্চকর বাইক নয়</li>
<li>কিছু প্রতিযোগীদের তুলনায় ভারী</li>
<li>বেসিক ইন্সট্রুমেন্ট কনসোল</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=CBShineBangla",
					"https://www.youtube.com/watch?v=HondaShineBD",
				},
				"verdict":     "ব্যতিক্রমী পরিশীলিততা সহ প্রিমিয়াম কমিউটার",
				"recommended": true,
			},
		},
		{
			ProductName: "Hero Splendor Plus",
			BanglaReview: `<div class="review-content">
<h3>সুবিধা</h3>
<ul>
<li>অসাধারণ জ্বালানি দক্ষতা প্রতি লিটারে ৬৫-৭৫ কিলোমিটার</li>
<li>সবচেয়ে সাশ্রয়ী ১০০সিসি মোটরসাইকেল</li>
<li>অত্যন্ত কম রক্ষণাবেক্ষণ খরচ</li>
<li>কয়েক দশকের প্রমাণিত নির্ভরযোগ্যতা</li>
<li>হালকা ওজনের এবং চালানো সহজ</li>
<li>বাংলাদেশ জুড়ে বিস্তৃত সার্ভিস নেটওয়ার্ক</li>
<li>নতুনদের জন্য ভালো</li>
<li>শক্তিশালী রিসেল ভ্যালু</li>
</ul>

<h3>অসুবিধা</h3>
<ul>
<li>অত্যন্ত বেসিক ডিজাইন এবং ফিচার</li>
<li>হাইওয়ে ব্যবহারের জন্য কম শক্তিশালী ইঞ্জিন</li>
<li>গড় বিল্ড কোয়ালিটি</li>
<li>পুরানো স্টাইলিং</li>
<li>লম্বা রাইডে অস্বস্তিকর সিট</li>
<li>ভেজা পরিস্থিতিতে দুর্বল ব্রেকিং পারফরম্যান্স</li>
<li>আধুনিক সুবিধার অভাব</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=SplendorBangla",
					"https://www.youtube.com/watch?v=HeroSplendorBD",
				},
				"verdict":     "সর্বোত্তম ভ্যালু ফর মানি কমিউটার মোটরসাইকেল",
				"recommended": true,
			},
		},
		{
			ProductName: "Yamaha YZF R15 V4",
			BanglaReview: `<div class="review-content">
<h3>সুবিধা</h3>
<ul>
<li>শক্তিশালী এবং রেভ-হ্যাপি ১৫৫সিসি ইঞ্জিন</li>
<li>অত্যাশ্চর্য স্পোর্টি ডিজাইন এয়ারোডাইনামিক ফেয়ারিং সহ</li>
<li>চমৎকার হ্যান্ডলিং এবং কর্নারিং ক্ষমতা</li>
<li>প্রিমিয়াম বিল্ড কোয়ালিটি এবং বিস্তারিত মনোযোগ</li>
<li>কানেক্টিভিটি সহ সম্পূর্ণ ডিজিটাল এলসিডি ইন্সট্রুমেন্ট ক্লাস্টার</li>
<li>উচ্চতর নিরাপত্তার জন্য ডুয়াল-চ্যানেল এবিএস</li>
<li>এলইডি হেডলাইট এবং টেইললাইট</li>
<li>রোমাঞ্চকর রাইডিং অভিজ্ঞতা</li>
<li>শক্তিশালী ব্র্যান্ড ইমেজ</li>
</ul>

<h3>অসুবিধা</h3>
<ul>
<li>প্রতিযোগীদের তুলনায় ব্যয়বহুল</li>
<li>আক্রমণাত্মক রাইডিং পজিশন দৈনন্দিন যাতায়াতের জন্য আরামদায়ক নয়</li>
<li>কম জ্বালানি দক্ষতা (৪০-৪৫ কিমি/লিটার)</li>
<li>উচ্চ রক্ষণাবেক্ষণ খরচ</li>
<li>পিলিয়ন আরামের জন্য উপযুক্ত নয়</li>
<li>ব্যয়বহুল স্পেয়ার পার্টস</li>
<li>সীমিত আন্ডার-সিট স্টোরেজ</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=R15V4Bangla",
					"https://www.youtube.com/watch?v=YamahaR15BD",
				},
				"verdict":     "উৎসাহীদের জন্য সেরা স্পোর্টি মোটরসাইকেল",
				"recommended": true,
			},
		},
		{
			ProductName: "Bajaj Pulsar 150",
			BanglaReview: `<div class="review-content">
<h3>সুবিধা</h3>
<ul>
<li>শক্তিশালী মিড-রেঞ্জ সহ শক্তিশালী ১৫০সিসি ইঞ্জিন</li>
<li>আক্রমণাত্মক এবং পেশীবহুল স্টাইলিং</li>
<li>ভালো ভ্যালু ফর মানি</li>
<li>লম্বা রাইডের জন্য আরামদায়ক রাইডিং পজিশন</li>
<li>ভালো জ্বালানি দক্ষতা প্রতি লিটারে ৪৫-৫০ কিলোমিটার</li>
<li>শক্তিশালী এবং মজবুত বিল্ড</li>
<li>ভালো হ্যান্ডলিং বৈশিষ্ট্য</li>
<li>স্পেয়ার পার্টসের ব্যাপক উপলব্ধতা</li>
</ul>

<h3>অসুবিধা</h3>
<ul>
<li>উচ্চ আরপিএমে কম্পন</li>
<li>জাপানি বাইকের তুলনায় গড় বিল্ড কোয়ালিটি</li>
<li>পুরানো সিঙ্গেল-চ্যানেল এবিএস</li>
<li>ভারী ক্লাচ অপারেশন</li>
<li>বেসিক ইন্সট্রুমেন্ট ক্লাস্টার</li>
<li>সাসপেনশন আরো ভালোভাবে টিউন করা যেত</li>
<li>কম রিসেল ভ্যালু</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=Pulsar150Bangla",
					"https://www.youtube.com/watch?v=BajajPulsarBD",
				},
				"verdict":     "সাশ্রয়ী পারফরম্যান্স মোটরসাইকেল",
				"recommended": true,
			},
		},
		{
			ProductName: "Suzuki Gixxer SF",
			BanglaReview: `<div class="review-content">
<h3>সুবিধা</h3>
<ul>
<li>পরিশীলিত এবং শক্তিশালী ১৫৫সিসি ইঞ্জিন</li>
<li>আকর্ষণীয় ফুল-ফেয়ার্ড স্পোর্টি ডিজাইন</li>
<li>চমৎকার হ্যান্ডলিং এবং স্থিতিশীলতা</li>
<li>ভালো বিল্ড কোয়ালিটি</li>
<li>শহর এবং হাইওয়ে উভয়ের জন্য আরামদায়ক</li>
<li>ভালো দক্ষতার জন্য ফুয়েল ইনজেকশন (৪৮-৫২ কিমি/লি)</li>
<li>নিরাপত্তার জন্য সিঙ্গেল-চ্যানেল এবিএস</li>
<li>সম্পূর্ণ ডিজিটাল ইন্সট্রুমেন্ট কনসোল</li>
</ul>

<h3>অসুবিধা</h3>
<ul>
<li>প্রিমিয়াম মূল্য</li>
<li>হোন্ডা/হিরোর তুলনায় সীমিত সার্ভিস নেটওয়ার্ক</li>
<li>হেডলাইটের উজ্জ্বলতা আরো ভালো হতে পারত</li>
<li>লম্বা রাইডারদের জন্য সামান্য সংকীর্ণ</li>
<li>গড় পিলিয়ন আরাম</li>
<li>ব্যয়বহুল স্পেয়ার পার্টস</li>
<li>কম গ্রাউন্ড ক্লিয়ারেন্স</li>
</ul>
</div>`,
			AdditionalDetails: map[string]interface{}{
				"youtube_reviews": []string{
					"https://www.youtube.com/watch?v=GixxerSFBangla",
					"https://www.youtube.com/watch?v=SuzukiGixxerBD",
				},
				"verdict":     "ভালো ভারসাম্যপূর্ণ স্পোর্ট-কমিউটার মোটরসাইকেল",
				"recommended": true,
			},
		},
	}

	for _, transData := range translations {
		// Find the product by name
		var product models.ProductModel
		if err := db.Where("name = ?", transData.ProductName).First(&product).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Printf("   ⚠️  Product '%s' not found, skipping translation\n", transData.ProductName)
				continue
			}
			return fmt.Errorf("error finding product '%s': %w", transData.ProductName, err)
		}

		// Find the review for this product
		var review models.ProductReviewModel
		if err := db.Where("product_id = ?", product.ID).First(&review).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Printf("   ⚠️  Review for '%s' not found, skipping translation\n", transData.ProductName)
				continue
			}
			return fmt.Errorf("error finding review for '%s': %w", transData.ProductName, err)
		}

		// Check if translation already exists
		var existing models.ProductReviewTranslationModel
		err := db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existing).Error
		if err == nil {
			fmt.Printf("   ℹ️  Bangla translation for '%s' already exists, skipping\n", transData.ProductName)
			continue
		}

		// Marshal additional details to JSON
		additionalDetailsJSON, err := json.Marshal(transData.AdditionalDetails)
		if err != nil {
			return fmt.Errorf("error marshaling additional details for '%s': %w", transData.ProductName, err)
		}

		// Create new translation
		translation := &models.ProductReviewTranslationModel{
			ProductReviewID:   review.ID,
			Locale:            "bn",
			Rating:            fmt.Sprintf("%.2f", review.Rating), // Convert float32 rating to string
			Reviews:           transData.BanglaReview,
			AdditionalDetails: additionalDetailsJSON,
		}

		if err := db.Create(translation).Error; err != nil {
			return fmt.Errorf("error creating translation for '%s': %w", transData.ProductName, err)
		}

		fmt.Printf("   ✅ Created Bangla translation for '%s'\n", transData.ProductName)
	}

	return nil
}
