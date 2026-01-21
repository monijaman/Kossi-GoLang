package seeders

import (
	"encoding/json"
	"fmt"
	"kossti/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// TVReviewTranslationSeeder handles seeding Bangla translations for TV product reviews
type TVReviewTranslationSeeder struct {
	BaseSeeder
}

// NewTVReviewTranslationSeeder creates a new TVReviewTranslationSeeder
func NewTVReviewTranslationSeeder() *TVReviewTranslationSeeder {
	return &TVReviewTranslationSeeder{BaseSeeder: BaseSeeder{name: "TV Review Translations"}}
}

// Seed implements the Seeder interface
func (tvrs *TVReviewTranslationSeeder) Seed(db *gorm.DB) error {
	translations := []struct {
		ProductName       string
		BanglaReview      string
		AdditionalDetails map[string]interface{}
	}{
		{
			ProductName:       "Samsung 55\" 4K Smart TV",
			BanglaReview:      "স্যামসাং 55\" 4K টিভি অসাধারণ ছবির গুণমান প্রদান করে যার সাথে প্রাণবন্ত রং এবং গভীর কালো রয়েছে। QLED প্যানেল গেমিং এবং চলচ্চিত্রের জন্য চমৎকার। স্মার্ট টিভি OS স্বজ্ঞাত এবং প্রতিক্রিয়াশীল। বিনোদন প্রেমীদের জন্য অত্যন্ত সুপারিশকৃত।",
			AdditionalDetails: map[string]interface{}{"verdict": "চমৎকার গেমিং ফিচার সহ প্রিমিয়াম 4K QLED ডিসপ্লে", "recommended": true},
		},
		{
			ProductName:       "Samsung 65\" 8K Smart TV",
			BanglaReview:      "65\" 8K রেজোলিউশন অসাধারণ স্পষ্টতা এবং বিবরণ সহ অবিশ্বাস্য। 120Hz রিফ্রেশ রেট গেমিংকে মসৃণ এবং প্রতিক্রিয়াশীল করে তোলে। চমৎকার উজ্জ্বলতা এবং বৈপরীত্য। প্রিমিয়াম মানের নির্মাণ। প্রিমিয়াম দেখার অভিজ্ঞতার জন্য বিনিয়োগের যোগ্য।",
			AdditionalDetails: map[string]interface{}{"verdict": "ভবিষ্যতের জন্য প্রস্তুত অসাধারণ 8K ডিসপ্লে", "recommended": true},
		},
		{
			ProductName:       "LG 55\" OLED TV",
			BanglaReview:      "LG OLED প্রযুক্তি নিখুঁত কালো এবং অসীম বৈপরীত্য প্রদান করে। 120Hz রিফ্রেশ রেট খেলাধুলা এবং গেমিংয়ের জন্য দুর্দান্ত। webOS মসৃণ এবং বৈশিষ্ট্য সমৃদ্ধ। প্রিমিয়াম বিল্ড কোয়ালিটি। সামান্য ব্যয়বহুল কিন্তু চমৎকার কর্মক্ষমতা।",
			AdditionalDetails: map[string]interface{}{"verdict": "নিখুঁত কালো স্তর সহ স্বর্ণমান OLED", "recommended": true},
		},
		{
			ProductName:       "LG 43\" HD Smart TV",
			BanglaReview:      "LG 43\" HD টিভি মূল্যের জন্য শালীন ছবির গুণমান প্রদান করে। webOS নেভিগেশনকে সহজ করে তোলে। দৈনন্দিন দেখার জন্য যথেষ্ট ভাল। কম রিফ্রেশ রেটের কারণে গেমিংয়ের জন্য আদর্শ নয়, তবে নৈমিত্তিক দেখার জন্য ভাল।",
			AdditionalDetails: map[string]interface{}{"verdict": "বাজেট সচেতন ক্রেতাদের জন্য ভাল মূল্য", "recommended": true},
		},
		{
			ProductName:       "Sony Bravia 55\" 4K TV",
			BanglaReview:      "Sony Bravia চমৎকার রঙের নির্ভুলতা এবং প্রাকৃতিক ছবির টোন প্রদান করে। Android TV ভাল অ্যাপ সমর্থন প্রদান করে। মজবুত বিল্ড কোয়ালিটি। চলচ্চিত্র উত্সাহীদের জন্য ভাল পছন্দ। প্রতিযোগীদের তুলনায় সীমিত স্মার্ট বৈশিষ্ট্য।",
			AdditionalDetails: map[string]interface{}{"verdict": "চলচ্চিত্র অনুরাগীদের জন্য চমৎকার রঙ নির্ভুলতা", "recommended": true},
		},
		{
			ProductName:       "Xiaomi 43\" Full HD Smart TV",
			BanglaReview:      "Xiaomi 43\" অর্থের জন্য চমৎকার মূল্য প্রদান করে। উজ্জ্বল প্রদর্শন ভাল রঙ পুনরুৎপাদনসহ। Android TV মসৃণভাবে কাজ করে। দৈনন্দিন ব্যবহারের জন্য ভাল। অডিও আরও ভাল হতে পারে। বাজেট ক্রেতাদের জন্য চমৎকার পছন্দ।",
			AdditionalDetails: map[string]interface{}{"verdict": "দুর্দান্ত Android ইকোসিস্টেম সহ সেরা বাজেট স্মার্ট টিভি", "recommended": true},
		},
		{
			ProductName:       "Xiaomi 55\" 4K Smart TV",
			BanglaReview:      "Xiaomi 55\" 4K টিভি তীক্ষ্ণ বিবরণ সহ দুর্দান্ত 4K সামগ্রী সরবরাহ করে। ভাল উজ্জ্বলতা এবং রঙের নির্ভুলতা। বিল্ড কোয়ালিটি শক্তিশালী। গেমিং কর্মক্ষমতা গ্রহণযোগ্য। অর্থের জন্য দুর্দান্ত মূল্য।",
			AdditionalDetails: map[string]interface{}{"verdict": "বাজেট মূল্যে চমৎকার 4K কর্মক্ষমতা", "recommended": true},
		},
		{
			ProductName:       "Walton 32\" HD Smart TV",
			BanglaReview:      "Walton 32\" HD টিভি স্থানীয় ব্র্যান্ডের জন্য একটি শক্তিশালী পছন্দ। HD সামগ্রীর জন্য ভাল ছবির গুণমান। আফটার সেলস সার্ভিস নির্ভরযোগ্য। নির্মাণ মজবুত। গেমিংয়ের জন্য রিফ্রেশ রেট আরও ভাল হতে পারে।",
			AdditionalDetails: map[string]interface{}{"verdict": "স্থানীয় সেবা সহ বিশ্বস্ত বাজেট বিকল্প", "recommended": true},
		},
		{
			ProductName:       "Walton 43\" Full HD Smart TV",
			BanglaReview:      "Walton 43\" ফুল এইচডি টিভি ভাল ছবির গুণমান এবং মসৃণ কর্মক্ষমতা প্রদান করে। স্থানীয় ব্র্যান্ড মানে সহজ সেবা এবং যন্ত্রাংশ। রঙ পুনরুৎপাদন শালীন। স্থানীয় বাজারের জন্য ভাল বিকল্প।",
			AdditionalDetails: map[string]interface{}{"verdict": "স্থানীয় সমর্থন সহ ভাল ফুল এইচডি কর্মক্ষমতা", "recommended": true},
		},
		{
			ProductName:       "Hisense 50\" 4K Smart TV",
			BanglaReview:      "Hisense 50\" 4K টিভি যুক্তিসঙ্গত মূল্যে চিত্তাকর্ষক 4K সামগ্রীর গুণমান প্রদান করে। ভাল রঙের নির্ভুলতা এবং উজ্জ্বলতা। Android TV প্ল্যাটফর্ম নির্ভরযোগ্য। গ্রহণযোগ্য বিল্ড কোয়ালিটি। ভাল বাজেট 4K বিকল্প।",
			AdditionalDetails: map[string]interface{}{"verdict": "সাশ্রয়ী 4K পারফরম্যান্স এবং নির্ভরযোগ্য সেবা", "recommended": true},
		},
		{
			ProductName:       "TCL 43\" Full HD Smart TV",
			BanglaReview:      "TCL 43\" ফুল এইচডি টিভি মৌলিক ফুল এইচডি সামগ্রী ভালভাবে সরবরাহ করে। সাশ্রয়ী মূল্য পয়েন্ট। ছবির গুণমান নৈমিত্তিক দেখার জন্য গ্রহণযোগ্য। স্মার্ট বৈশিষ্ট্য সীমিত। ভাল এন্ট্রি-লেভেল টিভি।",
			AdditionalDetails: map[string]interface{}{"verdict": "সর্বনিম্ন মূল্যে প্রবেশ-স্তর স্মার্ট টিভি", "recommended": true},
		},
		{
			ProductName:       "Haier 32\" HD Smart TV",
			BanglaReview:      "Haier 32\" এইচডি টিভি ছোট রুম এবং শোবার ঘরের জন্য উপযুক্ত। কমপ্যাক্ট সাইজ শালীন ছবির গুণমানসহ। Android TV ভাল কাজ করে। বিল্ড কোয়ালিটি গড়। বাজেট সচেতন ক্রেতাদের জন্য ভাল।",
			AdditionalDetails: map[string]interface{}{"verdict": "ছোট স্থানের জন্য কমপ্যাক্ট এবং সাশ্রয়ী", "recommended": true},
		},
		{
			ProductName:       "Panasonic 43\" Full HD Smart TV",
			BanglaReview:      "Panasonic 43\" ফুল এইচডি টিভি ভাল ছবির গুণমান সহ নির্ভরযোগ্য কর্মক্ষমতা প্রদান করে। জাপানি বিল্ড কোয়ালিটি স্পষ্ট। Android TV ইন্টিগ্রেশন মসৃণভাবে কাজ করে। ভাল অডিও আউটপুট। স্থায়িত্বের জন্য বিশ্বস্ত ব্র্যান্ড।",
			AdditionalDetails: map[string]interface{}{"verdict": "জাপানি প্রযুক্তি এবং দীর্ঘস্থায়িত্ব", "recommended": true},
		},
		{
			ProductName:       "Singer 32\" HD Smart TV",
			BanglaReview:      "Singer 32\" এইচডি টিভি সাশ্রয়ী এবং মৌলিক দেখার জন্য ভাল কাজ করে। ছবির গুণমান এইচডি সামগ্রীর জন্য গ্রহণযোগ্য। বিল্ড আরও শক্তিশালী হতে পারে। মৌলিক ব্যবহারের জন্য ভাল বাজেট বিকল্প।",
			AdditionalDetails: map[string]interface{}{"verdict": "মৌলিক ব্যবহারের জন্য সাশ্রয়ী বিকল্প", "recommended": true},
		},
		{
			ProductName:       "Vision 32\" Full HD Smart TV",
			BanglaReview:      "Vision 32\" ফুল এইচডি টিভি ফুল এইচডি স্পষ্টতা সহ ভাল মূল্য প্রদান করে। বিল্ড কোয়ালিটি শালীন। স্থানীয় ব্র্যান্ড যুক্তিসঙ্গত সমর্থনসহ। পারিবারিক দেখার জন্য ভাল।",
			AdditionalDetails: map[string]interface{}{"verdict": "পারিবারিক ব্যবহারের জন্য স্থানীয় ব্র্যান্ড", "recommended": true},
		},
		{
			ProductName:       "Minister 24\" HD Smart TV",
			BanglaReview:      "Minister 24\" এইচডি টিভি ছোট স্থান এবং ব্যক্তিগত ব্যবহারের জন্য চমৎকার। সাশ্রয়ী মূল্য। ছবির গুণমান মৌলিক কিন্তু গ্রহণযোগ্য। সীমিত বৈশিষ্ট্য কিন্তু কার্যকরী।",
			AdditionalDetails: map[string]interface{}{"verdict": "অতি-কমপ্যাক্ট এবং সর্বনিম্ন মূল্য", "recommended": true},
		},
		{
			ProductName:       "Rangs 32\" Full HD Smart TV",
			BanglaReview:      "Rangs 32\" ফুল এইচডি টিভি ভাল ফুল এইচডি দেখার অভিজ্ঞতা প্রদান করে। বিল্ড কোয়ালিটি শক্তিশালী। আফটার সেলস সার্ভিস স্থানীয়ভাবে উপলব্ধ। মূল্যের জন্য ভাল মূল্য।",
			AdditionalDetails: map[string]interface{}{"verdict": "শালীন স্থানীয় মধ্য-পরিসর বিকল্প", "recommended": true},
		},
		{
			ProductName:       "Rangs 43\" Full HD Smart TV",
			BanglaReview:      "Rangs 43\" ফুল এইচডি টিভি ভাল ছবির গুণমান সহ বড় স্ক্রিনের জন্য চমৎকার পছন্দ। বিল্ড মজবুত। সার্ভিস সমর্থন সহ স্থানীয় ব্র্যান্ড। দুর্দান্ত মূল্য বিকল্প।",
			AdditionalDetails: map[string]interface{}{"verdict": "মধ্য-থেকে-বড় স্ক্রিনের জন্য দুর্দান্ত মূল্য", "recommended": true},
		},
		{
			ProductName:       "MyOne 32\" Full HD Smart TV",
			BanglaReview:      "MyOne 32\" ফুল এইচডি টিভি গ্রহণযোগ্য ছবির গুণমানের সাথে বাজেট-বান্ধব। নৈমিত্তিক দেখার জন্য ভাল। বিল্ড কোয়ালিটি গড়। শালীন স্থানীয় ব্র্যান্ড বিকল্প।",
			AdditionalDetails: map[string]interface{}{"verdict": "বাজেট-বান্ধব স্থানীয় স্মার্ট টিভি", "recommended": true},
		},
		{
			ProductName:       "Toshiba 43\" Full HD Smart TV",
			BanglaReview:      "Toshiba 43\" ফুল এইচডি টিভি ভাল ছবির গুণমান সহ শক্তিশালী কর্মক্ষমতা সরবরাহ করে। আন্তর্জাতিক ব্র্যান্ড নির্ভরযোগ্যতা স্পষ্ট। বিল্ড মজবুত। গুণমান খোঁজেন এমনদের জন্য ভাল পছন্দ।",
			AdditionalDetails: map[string]interface{}{"verdict": "আন্তর্জাতিক ব্র্যান্ডের নির্ভরযোগ্যতা এবং মান", "recommended": true},
		},
		{
			ProductName:       "Toshiba 50\" 4K Smart TV",
			BanglaReview:      "Toshiba 50\" 4K টিভি চমৎকার রঙ রেন্ডারিং সহ চিত্তাকর্ষক 4K ডিসপ্লে প্রদান করে। আন্তর্জাতিক মান বজায় রাখা হয়েছে। বিল্ড প্রিমিয়াম। হোম থিয়েটার সেটআপের জন্য চমৎকার পছন্দ।",
			AdditionalDetails: map[string]interface{}{"verdict": "হোম থিয়েটার জন্য চমৎকার আন্তর্জাতিক 4K টিভি", "recommended": true},
		},
		{
			ProductName:       "Sharp 43\" Full HD Smart TV",
			BanglaReview:      "Sharp 43\" ফুল এইচডি টিভি ভাল রঙের নির্ভুলতা সহ উজ্জ্বল প্রদর্শন প্রদান করে। বিল্ড কোয়ালিটি শক্তিশালী। Android TV প্ল্যাটফর্ম ভাল কাজ করে। উজ্জ্বল রুমের জন্য ভাল পছন্দ।",
			AdditionalDetails: map[string]interface{}{"verdict": "উজ্জ্বল রুমের জন্য চমৎকার ডিসপ্লে", "recommended": true},
		},
		{
			ProductName:       "Konka 32\" Full HD Smart TV",
			BanglaReview:      "Konka 32\" ফুল এইচডি টিভি শালীন গুণমান সহ একটি ভাল চীনা ব্র্যান্ড বিকল্প। বিল্ড গ্রহণযোগ্য। মৌলিক দেখার চাহিদার জন্য ভাল। যুক্তিসঙ্গত মূল্য পয়েন্ট।",
			AdditionalDetails: map[string]interface{}{"verdict": "মৌলিক ব্যবহারের জন্য চীনা ব্র্যান্ড বিকল্প", "recommended": true},
		},
		{
			ProductName:       "ECO+ 24\" HD Smart TV",
			BanglaReview:      "ECO+ 24\" এইচডি টিভি ডর্ম বা অফিসের মতো খুব ছোট স্থানের জন্য আদর্শ। অতি-সাশ্রয়ী মূল্য। ছবির গুণমান মৌলিক কিন্তু কার্যকরী। সীমিত স্মার্ট বৈশিষ্ট্য।",
			AdditionalDetails: map[string]interface{}{"verdict": "অত্যন্ত কমপ্যাক্ট এবং সাশ্রয়ী বাজেট টিভি", "recommended": true},
		},
		{
			ProductName:       "Haiko 28\" Full HD Smart TV",
			BanglaReview:      "Haiko 28\" ফুল এইচডি টিভি কমপ্যাক্ট সাইজে ফুল এইচডি স্পষ্টতা প্রদান করে। ছোট শোবার ঘর বা ব্যক্তিগত ব্যবহারের জন্য ভাল। বিল্ড কোয়ালিটি শালীন। অর্থের জন্য মূল্য ভাল।",
			AdditionalDetails: map[string]interface{}{"verdict": "ছোট ব্যক্তিগত ব্যবহারের জন্য আদর্শ", "recommended": true},
		},
	}

	// Get TV category
	var tvCategory models.CategoryModel
	if err := db.Where("id = ?", 124).First(&tvCategory).Error; err != nil {
		return nil
	}

	for _, transData := range translations {
		// Find the product by name
		var product models.ProductModel
		if err := db.Where("name = ?", transData.ProductName).First(&product).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Printf("   ⚠️  পণ্য '%s' পাওয়া যায়নি, অনুবাদ এড়িয়ে যাচ্ছি\n", transData.ProductName)
				continue
			}
			return fmt.Errorf("error finding product '%s': %w", transData.ProductName, err)
		}

		// Find the review for this product
		var review models.ProductReviewModel
		if err := db.Where("product_id = ?", product.ID).First(&review).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Printf("   ⚠️  '%s' এর জন্য রিভিউ পাওয়া যায়নি, অনুবাদ এড়িয়ে যাচ্ছি\n", transData.ProductName)
				continue
			}
			return fmt.Errorf("error finding review for '%s': %w", transData.ProductName, err)
		}

		// Check if translation already exists
		var existing models.ProductReviewTranslationModel
		err := db.Where("product_review_id = ? AND locale = ?", review.ID, "bn").First(&existing).Error
		if err == nil {
			fmt.Printf("   ℹ️  '%s' এর জন্য বাংলা অনুবাদ ইতিমধ্যে বিদ্যমান, এড়িয়ে যাচ্ছি\n", transData.ProductName)
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
			Rating:            review.Rating,
			Reviews:           transData.BanglaReview,
			AdditionalDetails: additionalDetailsJSON,
		}

		if err := db.Create(translation).Error; err != nil {
			return fmt.Errorf("error creating translation for '%s': %w", transData.ProductName, err)
		}

		fmt.Printf("   ✅ '%s' এর জন্য বাংলা অনুবাদ তৈরি করা হয়েছে\n", transData.ProductName)
	}

	return nil
}
