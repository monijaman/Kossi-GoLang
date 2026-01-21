package seeders

import (
	"encoding/json"
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// BanglalinkReviewTranslationSeeder handles seeding Bangla translations for Banglalink reviews
type BanglalinkReviewTranslationSeeder struct {
	BaseSeeder
}

// NewBanglalinkReviewTranslationSeeder creates a new BanglalinkReviewTranslationSeeder
func NewBanglalinkReviewTranslationSeeder() *BanglalinkReviewTranslationSeeder {
	return &BanglalinkReviewTranslationSeeder{BaseSeeder: BaseSeeder{name: "Banglalink Review Translations"}}
}

// Seed implements the Seeder interface
func (brts *BanglalinkReviewTranslationSeeder) Seed(db *gorm.DB) error {
	// Find Banglalink review
	var banglalinkReview models.ProductReviewModel
	var banglalinkProduct models.ProductModel

	// First find the Banglalink product
	if err := db.Where("slug = ?", "banglalink").First(&banglalinkProduct).Error; err != nil {
		return fmt.Errorf("Banglalink product not found: %w", err)
	}

	// Then find the review for that product
	if err := db.Where("product_id = ?", banglalinkProduct.ID).First(&banglalinkReview).Error; err != nil {
		return fmt.Errorf("Banglalink review not found, please run Banglalink review seeder first: %w", err)
	}

	// Check if translation already exists
	var existingTranslation models.ProductReviewTranslationModel
	result := db.Where("product_review_id = ? AND locale = ?", banglalinkReview.ID, "bn").First(&existingTranslation)

	if result.Error == gorm.ErrRecordNotFound {
		// Create Bangla translation
		banglalinkReviewBN := `<div class="review-content">
<h2>বাংলালিংক মোবাইল অপারেটর বাংলাদেশ - সম্পূর্ণ পর্যালোচনা ২০২৪</h2>

<h3>সংক্ষিপ্ত বিবরণ</h3>
<p>বাংলালিংক বাংলাদেশের শীর্ষস্থানীয় মোবাইল নেটওয়ার্ক অপারেটরগুলির মধ্যে একটি, যা ভয়েস, ডেটা এবং ডিজিটাল পরিষেবা প্রদান করে। একটি গ্রাহক-কেন্দ্রিক টেলিকম প্রদানকারী হিসাবে, বাংলালিংক আধুনিক মোবাইল ব্যবহারকারীদের ক্রমবর্ধমান চাহিদা পূরণের জন্য উল্লেখযোগ্যভাবে বিকশিত হয়েছে।</p>

<h3>নেটওয়ার্ক মান ও কভারেজ</h3>
<ul>
<li><strong>৪জি এলটিই কভারেজ:</strong> ঢাকা, চট্টগ্রাম, সিলেট এবং খুলনা সহ প্রধান শহরগুলিতে চমৎকার ৪জি কভারেজ এবং গ্রামীণ এলাকায় ক্রমাগত সম্প্রসারণ</li>
<li><strong>নেটওয়ার্ক নির্ভরযোগ্যতা:</strong> সাধারণত স্থিতিশীল নেটওয়ার্ক কর্মক্ষমতা যা প্রতিযোগিতামূলক ডেটা গতিতে শহুরে এলাকায় গড়ে ১০-১৫ এমবিপিএস</li>
<li><strong>ভয়েস মান:</strong> কভারেজ এলাকায় পরিষ্কার কথোপকথন এবং সর্বনিম্ন ড্রপ-অফ হার</li>
<li><strong>২জি/৩জি ব্যাকআপ:</strong> লিগেসি ডিভাইস এবং দূরবর্তী এলাকার জন্য নির্ভরযোগ্য ফলব্যাক কভারেজ</li>
</ul>

<h3>সুবিধা - বাংলালিংক কেন বেছে নিবেন</h3>
<ul>
<li>সাশ্রয়ী কল রেট - স্থানীয় এবং আন্তর্জাতিক কলের জন্য প্রতিযোগিতামূলক মূল্য</li>
<li>ডেটা-বান্ধব বান্ডেল - দৈনন্দিন ব্রাউজিং এবং স্ট্রিমিংয়ের জন্য আকর্ষণীয় ইন্টারনেট প্যাকেজ</li>
<li>সোশ্যাল মিডিয়া প্যাকেজ - ফেসবুক, হোয়াটসঅ্যাপ এবং অন্যান্য প্ল্যাটফর্মের জন্য ডেডিকেটেড বান্ডেল</li>
<li>নিয়মিত প্রচার - নিয়মিত বিশেষ অফার এবং বোনাস ডেটা ডিল</li>
<li>ভালো গ্রাহক সেবা - হেল্পলাইন এবং অ্যাপ সহ একাধিক সহায়তা চ্যানেল</li>
<li>বাংলালিংক অ্যাপ - ব্যালেন্স চেকিং এবং প্যাকেজ ম্যানেজমেন্টের জন্য ব্যবহারকারী-বান্ধব মোবাইল অ্যাপ্লিকেশন</li>
<li>ডিজিটাল পেমেন্ট অপশন - মোবাইল ব্যাংকিং এবং ডিজিটাল ওয়ালেটের সমর্থন</li>
<li>নমনীয় পরিকল্পনা - বিভিন্ন বাজেটের জন্য বিভিন্ন প্রি-পেইড এবং পোস্ট-পেইড অপশন</li>
<li>রোমিং প্যাকেজ - প্রতিযোগিতামূলক আন্তর্জাতিক রোমিং হার</li>
<li>জরুরী ব্যালেন্স বৈশিষ্ট্য - জরুরী কল জন্য ব্যালেন্স ধার করার বিকল্প</li>
</ul>

<h3>অসুবিধা - উন্নতির ক্ষেত্র</h3>
<ul>
<li>শহরের কিছু এলাকায় পিক আওয়ারে নেটওয়ার্ক ভিড়ের সমস্যা</li>
<li>নির্দিষ্ট পরিকল্পনায় মাঝেমধ্যে ডেটা গতি কমে যাওয়া</li>
<li>পিক সময়ে গ্রাহক সেবা প্রতীক্ষা সময় দীর্ঘ হতে পারে</li>
<li>রোমিং রেট প্রতিযোগীদের তুলনায় এখনও বেশি</li>
<li>বড় প্রতিযোগীদের তুলনায় দূরবর্তী এলাকায় সীমিত কভারেজ</li>
<li>সিম অ্যাক্টিভেশন প্রক্রিয়া সহজ করা যেত</li>
<li>কিছু প্যাকেজের কঠোর বৈধতার সময়সীমা রয়েছে</li>
</ul>

<h3>সক্রিয় ডেটা প্যাকেজ ও বান্ডেল</h3>
<ul>
<li><strong>দৈনিক ডেটা প্যাক:</strong> সাশ্রয়ী মূল্যে ১০০এমবি-৫০০এমবি দৈনিক বান্ডেল</li>
<li><strong>সাপ্তাহিক ডেটা প্যাক:</strong> নিয়মিত ব্যবহারকারীদের জন্য ১জিবি-৩জিবি সাপ্তাহিক বান্ডেল</li>
<li><strong>মাসিক আনলিমিটেড ডেটা:</strong> ভারী ব্যবহারকারীদের জন্য সীমাহীন ডেটা প্ল্যান</li>
<li><strong>সোশ্যাল মিডিয়া বান্ডেল:</strong> ফেসবুক, হোয়াটসঅ্যাপ, ইউটিউব বিশেষ প্যাকেজ</li>
<li><strong>রাত্রিকালীন ডেটা প্যাকেজ:</strong> দেরী রাতের সময় ছাড়ে পাওয়া ডেটা</li>
<li><strong>ওটিটি প্যাকেজ:</strong> ইউটিউব, নেটফ্লিক্সের মতো স্ট্রিমিং সেবার জন্য ছাড়ে ডেটা</li>
</ul>

<h3>সক্রিয় ভয়েস ও এসএমএস প্যাকেজ</h3>
<ul>
<li><strong>স্থানীয় কল বান্ডেল:</strong> প্রতি-মিনিট হার অথবা দৈনিক/সাপ্তাহিক কল প্যাকেজ</li>
<li><strong>এসএমএস বান্ডেল:</strong> সাশ্রয়ী স্থানীয় এবং আন্তর্জাতিক এসএমএস প্যাকেজ</li>
<li><strong>রোমিং প্যাকেজ:</strong> বাংলাদেশের মধ্যে এবং বাইরে ভ্রমণের জন্য বিশেষ অফার</li>
<li><strong>আন্তর্জাতিক কল বান্ডেল:</strong> জনপ্রিয় গন্তব্যে ছাড়প্রাপ্ত হার</li>
<li><strong>এফএনএফ (বন্ধু এবং পরিবার):</strong> ঘন ঘন ডাকা নম্বরে হ্রাসকৃত হার</li>
</ul>

<h3>বাংলালিংক মানি ও আর্থিক সেবা</h3>
<ul>
<li>ডিজিটাল লেনদেনের জন্য মোবাইল ওয়ালেট</li>
<li>অর্থ স্থানান্তর সেবা</li>
<li>ইউটিলিটি বিল পেমেন্ট</li>
<li>দ্রুত বণিক পেমেন্ট সিস্টেম</li>
</ul>

<h3>গ্রাহক সেবার মান</h3>
<ul>
<li>২৪/৭ গ্রাহক হেল্পলাইন সাপোর্ট</li>
<li>দ্রুত ব্যালেন্স এবং প্যাকেজ ম্যানেজমেন্টের জন্য ইউএসএসডি কোড</li>
<li>স্ব-সেবা বৈশিষ্ট্যের জন্য মোবাইল অ্যাপ</li>
<li>ফেসবুক এবং টুইটারের মাধ্যমে সোশ্যাল মিডিয়া সহায়তা</li>
<li>প্রধান শহরগুলি জুড়ে ফিজিক্যাল সার্ভিস সেন্টার</li>
</ul>

<h3>মূল্য নির্ধারণ ও মূল্যের জন্য অর্থ</h3>
<p>বাংলালিংক সমস্ত বিভাগে প্রতিযোগিতামূলক মূল্য নির্ধারণ প্রদান করে। তাদের প্রি-পেইড প্যাকেজগুলি মূল্য-ভিত্তিক, যা বাজেট-সচেতন ব্যবহারকারীদের কাছে অ্যাক্সেসযোগ্য করে তোলে। তবে পোস্ট-পেইড পরিকল্পনাগুলি প্রতিযোগীদের তুলনায় তুলনামূলকভাবে বেশি। সামগ্রিকভাবে, নেটওয়ার্ক গুণমান এবং সেবা অফারগুলি বিবেচনা করে ভালো মূল্য মানি।</p>

<h3>চূড়ান্ত রায়</h3>
<p>বাংলালিংক বাংলাদেশে মোবাইল সেবার জন্য একটি নির্ভরযোগ্য পছন্দ, বিশেষত যারা সাশ্রয়ী ডেটা এবং কল প্যাকেজ খুঁজছেন তাদের জন্য। তাদের ঘন ঘন প্রচার এবং ব্যবহারকারী-বান্ধব অ্যাপ তাদের প্রতিযোগিতামূলক করে তোলে। দৈনন্দিন যাত্রীদের জন্য এবং নমনীয়, চুক্তিহীন সেবা প্রয়োজন এমন ব্যবহারকারীদের জন্য আদর্শ।</p>

<h3>সুপারিশ</h3>
<p><strong>সুপারিশকৃত:</strong> বাজেট-সচেতন ব্যবহারকারী, দৈনন্দিন ডেটা ভোক্তা, বাংলাদেশের মধ্যে ঘন ঘন ভ্রমণকারী, ছোট ব্যবসায়ী, শিক্ষার্থীরা</p>
<p><strong>আদর্শ নয়:</strong> প্রিমিয়াম গ্রাহক সেবা প্রয়োজন এমন ব্যবহারকারী, সীমিত ৪জি কভারেজ এলাকায় থাকা, বা সম্পূর্ণ সস্তা হার খোঁজা</p>

<h3>সামগ্রিক রেটিং বিভাজন</h3>
<ul>
<li>নেটওয়ার্ক মান: ৪/৫</li>
<li>কভারেজ এলাকা: ৪/৫</li>
<li>কল মান: ৪/৫</li>
<li>ডেটা গতি: ৪/৫</li>
<li>প্যাকেজ মূল্য: ৪.৫/৫</li>
<li>গ্রাহক সেবা: ৩.৫/৫</li>
<li>মূল্য নির্ধারণ: ৪/৫</li>
<li>সামগ্রিক: ৪.১/৫</li>
</ul>

<h3>চূড়ান্ত মন্তব্য</h3>
<p>বাংলালিংক বাংলাদেশে মোবাইল অপারেটরগুলির মধ্যে একটি সেরা পছন্দ হিসাবে থাকে। বাংলালিংক অ্যাপের মাধ্যমে সাশ্রয়ী প্যাকেজ, নিয়মিত প্রচার এবং ডিজিটাল উদ্ভাবনের উপর তাদের ফোকাস তাদের আলাদা করে তোলে। গ্রাহক সেবা এবং কভারেজ সম্প্রসারণে উন্নতির জায়গা থাকলেও, গ্রাহক সন্তুষ্টির প্রতি তাদের প্রতিশ্রুতি স্পষ্ট। আপনি হালকা ব্যবহারকারী হোন বা ভারী ডেটা ভোক্তা, বাংলালিংকের প্রতিযোগিতামূলক মূল্যে অফার করার কিছু না কিছু আছে।</p>
</div>`

		additionalDetails := map[string]interface{}{
			"youtube_reviews": []string{
				"https://www.youtube.com/results?search_query=বাংলালিংক+রিভিউ+বাংলাদেশ",
				"https://www.youtube.com/results?search_query=বাংলালিংক+প্যাকেজ+তুলনা",
			},
			"verdict":     "আকর্ষণীয় ডেটা প্যাকেজ সহ সেরা বাজেট-বান্ধব মোবাইল অপারেটর",
			"recommended": true,
			"rating_split": map[string]float32{
				"network_quality":  4.0,
				"coverage":         4.0,
				"call_quality":     4.0,
				"data_speed":       4.0,
				"package_value":    4.5,
				"customer_service": 3.5,
				"pricing":          4.0,
			},
			"seo_keywords": []string{
				"বাংলালিংক মোবাইল অপারেটর বাংলাদেশ",
				"বাংলালিংক প্যাকেজ ২০২৪",
				"বাংলালিংক ডেটা প্ল্যান",
				"বাংলালিংক রিভিউ",
				"সেরা মোবাইল অপারেটর বাংলাদেশ",
				"বাংলালিংক ৪জি কভারেজ",
				"বাংলালিংক কল রেট",
				"বাংলালিংক ইন্টারনেট প্যাকেজ",
			},
			"pros": []string{
				"সাশ্রয়ী কল রেট",
				"আকর্ষণীয় ডেটা বান্ডেল",
				"নিয়মিত প্রচার",
				"ভালো ৪জি কভারেজ",
				"মোবাইল অ্যাপ উপলব্ধ",
				"ডিজিটাল পেমেন্ট সমর্থন",
				"জরুরী ব্যালেন্স বৈশিষ্ট্য",
				"নমনীয় প্ল্যান",
			},
			"cons": []string{
				"পিক আওয়ার ভিড়",
				"দূরবর্তী এলাকায় সীমিত কভারেজ",
				"দীর্ঘ গ্রাহক সেবা প্রতীক্ষা সময়",
				"উচ্চতর রোমিং হার",
				"মাঝেমধ্যে গতি কমে যাওয়া",
			},
		}

		additionalDetailsJSON, err := json.Marshal(additionalDetails)
		if err != nil {
			return fmt.Errorf("failed to marshal additional details: %w", err)
		}

		translation := models.ProductReviewTranslationModel{
			ProductReviewID:   banglalinkReview.ID,
			Locale:            "bn",
			Rating:            fmt.Sprintf("%.2f", banglalinkReview.Rating), // Convert float32 rating to string
			Reviews:           banglalinkReviewBN,
			AdditionalDetails: additionalDetailsJSON,
		}

		if err := db.Create(&translation).Error; err != nil {
			return fmt.Errorf("failed to create Banglalink review translation: %w", err)
		}

		fmt.Printf("✓ Created Banglalink review translation (ID: %d, Locale: bn)\n", translation.ID)
	}

	return nil
}
