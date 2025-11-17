package migrations

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// UpdateRobiReviewsWithComparison updates product reviews with HTML comparison content
func UpdateRobiReviewsWithComparison(db *gorm.DB) error {
	fmt.Println("\n========== UPDATING ROBI REVIEWS WITH HTML COMPARISON ==========")

	// English review HTML content
	englishReview := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Robi Review</title>
</head>
<body>
    <article class="robi-review review-section">
        <header>
            <h1>Robi: Good Balance Between Quality and Price</h1>
            <p class="rating"><strong>Overall Rating: 4.05/5 ⭐⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>Executive Summary</h2>
            <p>Robi is Bangladesh's third-largest mobile operator with 30+ million subscribers, offering a balanced mix of network quality and affordable pricing. Ranked #3 in the market with a 4.05/5 overall rating, Robi provides a solid middle-ground option for budget-conscious users who still value decent network quality.</p>
        </section>

        <section class="network-quality">
            <h2>Network Quality &amp; Coverage</h2>
            <p><strong>Rating: 4.0/5</strong></p>
            <ul>
                <li>Urban Coverage: 98% with good 4G availability</li>
                <li>Rural Coverage: 62% with steady expansion</li>
                <li>Average Speed: 12 Mbps download, 7 Mbps peak hours</li>
                <li>Network Uptime: 97.8% reliability</li>
                <li>Latency: 30ms average</li>
            </ul>
        </section>

        <section class="pros">
            <h2>Pros (9 Major Advantages)</h2>
            <ul>
                <li> <strong>Good balance</strong> - Quality and affordability combined well</li>
                <li> <strong>Competitive pricing</strong> - Similar to Banglalink, slightly cheaper data</li>
                <li> <strong>Decent coverage</strong> - 98% urban and 62% rural coverage</li>
                <li> <strong>User-friendly services</strong> - Simple app and intuitive interface</li>
                <li> <strong>Business packages</strong> - Good corporate solutions available</li>
                <li> <strong>Regular promotions</strong> - Frequent special offers</li>
                <li> <strong>Multiple payment options</strong> - Various digital payment methods</li>
                <li> <strong>Reasonable customer support</strong> - Responsive service team</li>
                <li> <strong>Flexible validity</strong> - Good validity periods on packages</li>
            </ul>
        </section>

        <section class="cons">
            <h2>Cons (6 Main Limitations)</h2>
            <ul>
                <li> <strong>Peak hour slowdowns</strong> - Noticeable speed drops during busy times</li>
                <li> <strong>Limited rural expansion</strong> - Coverage slower than competitors</li>
                <li> <strong>Fewer entertainment options</strong> - Limited streaming bundles</li>
                <li> <strong>Average speeds</strong> - Not as fast as Grameenphone</li>
                <li> <strong>Limited 5G readiness</strong> - Slower technology adoption</li>
                <li> <strong>Less brand recognition</strong> - Smaller loyalty program</li>
            </ul>
        </section>

        <section class="verdict">
            <h2>Verdict</h2>
            <p><strong>RECOMMENDED - Good Alternative</strong></p>
            <p>Robi is a solid choice for users seeking good network quality at reasonable prices. While not as premium as Grameenphone or as budget-friendly as some competitors, Robi offers excellent value for money with its balanced approach.</p>
        </section>
    </article>
</body>
</html>`

	// Bangla review
	bengaliReview := `<!DOCTYPE html>
<html lang="bn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>রোবি রিভিউ</title>
</head>
<body>
    <article class="robi-review-bn review-section">
        <header>
            <h1>রোবি: মান এবং মূল্যের মধ্যে ভাল ভারসাম্য</h1>
            <p class="rating"><strong>সামগ্রিক রেটিং: ৪.०५/৫ ⭐⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>সংক্ষিপ্ত পর্যালোচনা</h2>
            <p>রোবি বাংলাদেশের তৃতীয় বৃহত্তম মোবাইল অপারেটর ৩০+ মিলিয়ন গ্রাহক সেবা প্রদান করে নেটওয়ার্ক মান এবং সাশ্রয়ী মূল্যের মধ্যে ভাল ভারসাম্য প্রদান করে। বাজারে ৩নম্বর স্থানে ৪.०५/५ রেটিং সহ, রোবি বাজেট-সচেতন ব্যবহারকারীদের জন্য একটি নিরাপদ মধ্য-পথ বিকল্প যারা শালীন নেটওয়ার্ক মান মূল্য দেয়।</p>
        </section>

        <section class="network-quality">
            <h2>নেটওয়ার্ক গুণমান এবং কভারেজ</h2>
            <p><strong>রেটিং: ৪.०५/५</strong></p>
            <ul>
                <li>শহরের কভারেজ: ৯৮% ভাল ৪জী উপলব্ধতা সহ</li>
                <li>গ্রামীণ কভারেজ: ৬२% স্থির সম্প্রসারণ সহ</li>
                <li>গড় গতি: १२ এমবিপিএস ডাউনলোড, ७ এমবিপিএস পিক আওয়ার্স</li>
                <li>নেটওয়ার্ক আপটাইম: ९७.८% নির্ভরযোগ্যতা</li>
                <li>লেটেন্সি: ३० এমএস গড়</li>
            </ul>
        </section>

        <section class="market-comparison">
            <h2>বাজার তুলনা</h2>
            <table border="1" cellpadding="10" cellspacing="0" style="width: 100%; border-collapse: collapse;">
                <thead>
                    <tr style="background-color: #f5f5f5;">
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">অপারেটর</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">র‍্যাঙ্ক</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">রেটিং</th>
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">সেরা জন্য</th>
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">রোবির তুলনায় অবস্থান</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>গ্রামীণফোন</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#১</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">४.६५/५ ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">সেরা মান</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">প्रिमिय়াম অপশন, ३०% বেশি দামী</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>ব্যাংলালিংক</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#२</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">४.१/५ ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">সেরা মূল্য</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">সামান্য ভাল রেটিং, সমান মূল্য</td>
                    </tr>
                    <tr style="background-color: #fff3cd;">
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>রোবি</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;"><strong>#३</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;"><strong>४.०५/५ ⭐</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>ভাল ভারসাম্য</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>বর্তমান - মধ্য-স্তরের প্রতিযোগী</strong></td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="pricing">
            <h2>মূল্য নির্ধারণ তুলনা</h2>
            <p><strong>রেটিং: ४.२/५ (প্রতিযোগিতামূলক মূল্য)</strong></p>
            <table border="1" cellpadding="10" cellspacing="0" style="width: 100%; border-collapse: collapse;">
                <thead>
                    <tr style="background-color: #f5f5f5;">
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">সেবা</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">গ্রামীণফোন</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">ব্যাংলালিংক</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">রোবি</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>१००एमबि दैनिक</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">৮০ টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">৫০ টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">৫৫ টাকা ⭐</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>१जीबी साप्ताहिक</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">१८० টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">१२० টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">१३० টাকা ⭐</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="pros">
            <h2>সুবিধাসমূহ (९ প্রধান সুবিধা)</h2>
            <ul>
                <li><strong>ভাল ভারসাম্য</strong> - মান এবং সাশ্রয়যোগ্যতা ভালভাবে সংমিশ্রিত</li>
                <li><strong>প্রতিযোগিতামূলক মূল্য</strong> - বাংলালিংকের অনুরূপ, সামান্য সস্তা ডেটা</li>
                <li><strong>শালীন কভারেজ</strong> - ९८% শহুরে এবং ६२% গ্রামীণ কভারেজ</li>
                <li><strong>ব্যবহারকারী-বান্ধব সেবা</strong> - সহজ অ্যাপ এবং স্বজ্ঞাত ইন্টারফেস</li>
                <li><strong>ব্যবসায়িক প্যাকেজ</strong> - ভাল কর্পোরেট সমাধান উপলব্ধ</li>
                <li><strong>নিয়মিত প্রচার</strong> - ঘন ঘন বিশেষ অফার</li>
                <li><strong>একাধিক পেমেন্ট বিকল্প</strong> - বিভিন্ন ডিজিটাল পেমেন্ট পদ্ধতি</li>
                <li><strong>যুक्तिसঙ্गत গ্রাহক সেবা</strong> - প্রতিক্রিয়াশীল সেবা দল</li>
                <li><strong>নমনীয় বৈধতা</strong> - প্যাকেজগুলিতে ভাল বৈধতা সময়কাল</li>
            </ul>
        </section>

        <section class="cons">
            <h2>অসুবিধাসমূহ (६ প্রধান সীমাবদ্ধতা)</h2>
            <ul>
                <li><strong>পিক আওয়ার মন্থরতা</strong> - ব্যস্ত সময়ে লক্ষণীয় গতি হ্রাস</li>
                <li><strong>সীমিত গ্রামীণ সম্প্রসারণ</strong> - প্রতিযোগীদের তুলনায় কভারেজ ধীর</li>
                <li><strong>কম বিনোদন বিকল্প</strong> - সীমিত স্ট্রিমিং বান্ডেল</li>
                <li><strong>গড় গতি</strong> - গ্রামীণফোনের মতো দ্রুত নয়</li>
                <li><strong>সীমিত ५जी প্রস্তুতি</strong> - ধীর প্রযুক্তি গ্রহণ</li>
                <li><strong>কম ব্র্যান্ড স্বীকৃতি</strong> - ছোট আনুগত্য প্রোগ্রাম</li>
            </ul>
        </section>

        <section class="verdict">
            <h2>চূড়ান্ত মূল্যায়ন</h2>
            <p><strong>সুপারিশকৃত - ভাল বিকল্প</strong></p>
            <p>রোবি যুক्तिसঙ००त মূল্যে ভাল নেটওয়ার्क মান খোঁজেন এমন ব্যবহারকারীদের জন্য একটি নিরাপদ পছন্দ। গ্রামীণফোনের মতো প্রিমিয়াম নয় বা কিছু প্রতিযোগীদের মতো বাজেট-বান্ধব না হলেও, রোবি তার ভারসাম्यপূর্ণ পদ্ধতির সাথে অসাধারণ মূল्य প্রদান করে।</p>
        </section>
    </article>
</body>
</html>`

	// Additional details JSON
	additionalDetails := datatypes.JSONMap{
		"market_rank":             3,
		"subscribers_millions":    30,
		"overall_rating":          4.05,
		"network_quality_rating":  4.0,
		"coverage_rating":         4.0,
		"data_speed_rating":       3.9,
		"pricing_rating":          4.2,
		"package_variety_rating":  4.0,
		"customer_support_rating": 4.0,
		"value_for_money_rating":  4.15,
		"avg_data_speed_mbps":     12,
		"peak_hour_speed_mbps":    7,
		"network_uptime_percent":  97.8,
		"latency_ms":              30,
		"urban_coverage_percent":  98.0,
		"rural_coverage_percent":  62.0,
		"hotline_number":          "123",
		"website_url":             "https://www.robi.com.bd",
		"verdict":                 "Recommended - Good Alternative",
		"recommendation_score":    7,
	}

	detailsJSON, err := json.Marshal(additionalDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Update English review
	result := db.Model(&map[string]interface{}{}).
		Table("product_reviews").
		Where("id = ? AND product_id = (SELECT id FROM products WHERE slug = ?)", 2, "robi").
		Updates(map[string]interface{}{
			"reviews":            englishReview,
			"additional_details": string(detailsJSON),
			"updated_at":         gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if result.Error != nil {
		log.Printf("Error updating English review: %v", result.Error)
		return result.Error
	}

	// Update Bangla translation
	bengaliDetails := datatypes.JSONMap{
		"language":             "Bengali",
		"market_rank":          3,
		"subscribers_millions": 30,
		"overall_rating":       4.05,
		"verdict":              "সুপারিশকৃত - ভাল বিকল্প",
		"recommendation_score": 7,
	}

	bengaliJSON, err := json.Marshal(bengaliDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal Bengali JSON: %w", err)
	}

	result = db.Model(&map[string]interface{}{}).
		Table("product_review_translations").
		Where("product_review_id = ? AND locale = ?", 2, "bn").
		Updates(map[string]interface{}{
			"reviews":            bengaliReview,
			"additional_details": string(bengaliJSON),
			"updated_at":         gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if result.Error != nil {
		log.Printf("Error updating Bangla translation: %v", result.Error)
		return result.Error
	}

	fmt.Println("✅ Robi reviews updated successfully!")

	return nil
}
