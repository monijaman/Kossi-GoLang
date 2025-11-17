package migrations

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// UpdateAirtelReviewsWithComparison updates product reviews with HTML comparison content
func UpdateAirtelReviewsWithComparison(db *gorm.DB) error {
	fmt.Println("\n========== UPDATING AIRTEL REVIEWS WITH HTML COMPARISON ==========")

	// English review HTML content
	englishReview := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Airtel Review</title>
</head>
<body>
    <article class="airtel-review">
        <header>
            <h1>Airtel: Growing Network with Regional Options</h1>
            <p class="rating"><strong>Overall Rating: 3.05/5 ⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>Executive Summary</h2>
            <p>Airtel is Bangladesh's fourth mobile operator with 8+ million subscribers, offering growing network infrastructure. Ranked #4 in the market with a 3.05/5 overall rating, Airtel is best suited for users in areas with Airtel coverage seeking budget-friendly options.</p>
        </section>

        <section class="network-quality">
            <h2>Network Quality &amp; Coverage</h2>
            <p><strong>Rating: 3.0/5</strong></p>
            <ul>
                <li>Urban Coverage: 90% with growing infrastructure</li>
                <li>Rural Coverage: 35% with limited expansion</li>
                <li>Average Speed: 10 Mbps download, 5 Mbps peak hours</li>
                <li>Network Uptime: 96.5% reliability</li>
                <li>Latency: 45ms average</li>
            </ul>
        </section>

        <section class="pros">
            <h2>Pros (6 Major Advantages)</h2>
            <ul>
                <li> <strong>Budget-friendly rates</strong> - Very competitive pricing</li>
                <li> <strong>Growing network</strong> - Continuous infrastructure expansion</li>
                <li> <strong>International presence</strong> - Part of global Airtel Group</li>
                <li> <strong>Simple plans</strong> - Easy to understand packages</li>
                <li> <strong>Good for urban areas</strong> - Better urban coverage</li>
                <li> <strong>Promotional offers</strong> - Regular special deals</li>
            </ul>
        </section>

        <section class="cons">
            <h2>Cons (7 Main Limitations)</h2>
            <ul>
                <li> <strong>Limited rural coverage</strong> - Only 35% in villages</li>
                <li> <strong>Slower speeds</strong> - Slower data speeds than competitors</li>
                <li> <strong>Network congestion</strong> - Frequent slowdowns during peak hours</li>
                <li> <strong>Limited customer support</strong> - Fewer support channels</li>
                <li> <strong>Fewer digital services</strong> - Limited banking and content options</li>
                <li> <strong>Smaller subscriber base</strong> - Limited ecosystem</li>
                <li> <strong>High latency</strong> - Not suitable for gaming/streaming</li>
            </ul>
        </section>

        <section class="verdict">
            <h2>Verdict</h2>
            <p><strong>AVERAGE - Budget Option Only</strong></p>
            <p>Airtel is suitable only for budget-conscious users in areas with coverage. Not recommended for those prioritizing network quality or comprehensive services.</p>
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
    <title>এয়ারটেল রিভিউ</title>
</head>
<body>
    <article class="airtel-review-bn">
        <header>
            <h1>এয়ারটেল: ক্রমবর্ধমান নেটওয়ার্ক এবং আঞ্চলিক বিকল্প</h1>
            <p class="rating"><strong>সামগ্রিক রেটিং: ৩.০৫/৫ ⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>সংক্ষিপ্ত পর্যালোচনা</h2>
            <p>এয়ারটেল বাংলাদেশের চতুর্থ মোবাইল অপারেটর, যা ৮+ মিলিয়ন গ্রাহক সেবা প্রদান করে। ক্রমবর্ধমান নেটওয়ার্ক অবকাঠামো অফার করে; বাজারে #৪ নম্বরে এবং সামগ্রিক রেটিং ৩.০৫/৫। এয়ারটেল কভারেজ সহ অঞ্চলে বাজেট-সচেতন ব্যবহারকারীদের জন্য একটি উপযুক্ত বিকল্প।</p>
        </section>

        <section class="network-quality">
            <h2>নেটওয়ার্ক গুণমান ও কভারেজ</h2>
            <p><strong>রেটিং: ৩.০/৫</strong></p>
            <ul>
                <li>শহরের কভারেজ: ৯০% ক্রমবর্ধমান অবকাঠামো সহ</li>
                <li>গ্রামীণ কভারেজ: ৩৫% সীমিত সম্প্রসারণ</li>
                <li>গড় গতি: ১০ এমবিপিএস ডাউনলোড, ৫ এমবিপিএস পিক আওয়ার্স</li>
                <li>নেটওয়ার্ক আপটাইম: ৯৬.৫% নির্ভরযোগ্যতা</li>
                <li>লেটেন্সি: ৪৫ এমএস গড়</li>
            </ul>
        </section>

        <section class="pros">
            <h2>সুবিধাসমূহ (৬টি প্রধান সুবিধা)</h2>
            <ul>
                <li><strong>বাজেট-ফ্রেন্ডলি রেট</strong> - প্রতিযোগিতামূলক মূল্য</li>
                <li><strong>ক্রমবর্ধমান নেটওয়ার্ক</strong> - অবকাঠামো সম্প্রসারণ চলমান</li>
                <li><strong>আন্তর্জাতিক উপস্থিতি</strong> - গ্লোবাল এয়ারটেল গ্রুপের অংশ</li>
                <li><strong>সরল পরিকল্পনা</strong> - বোঝা সহজ প্যাকেজ</li>
                <li><strong>শহুরে ক্ষেত্রে ভালো</strong> - উন্নত শহুরে কভারেজ</li>
                <li><strong>প্রচারণা ও অফার</strong> - নিয়মিত বিশেষ ডিল</li>
            </ul>
        </section>

        <section class="cons">
            <h2>অসুবিধাসমূহ (৭টি প্রধান সীমাবদ্ধতা)</h2>
            <ul>
                <li><strong>সীমিত গ্রামীণ কভারেজ</strong> - গ্রামে মাত্র ৩৫%</li>
                <li><strong>ধীর গতি</strong> - প্রতিযোগীদের তুলনায় ধীর ডেটা গতি</li>
                <li><strong>নেটওয়ার্ক ভিড়</strong> - পিক আওয়ারে মন্থরতা</li>
                <li><strong>সীমিত গ্রাহক সহায়তা</strong> - কম চ্যানেল</li>
                <li><strong>কম ডিজিটাল সেবা</strong> - সীমিত ব্যাংকিং ও কন্টেন্ট অপশন</li>
                <li><strong>ছোট গ্রাহক ভিত্তি</strong> - সীমিত ইকোসিস্টেম</li>
                <li><strong>উচ্চ লেটেন্সি</strong> - গেমিং/স্ট্রিমিং-এ অনুপযুক্ত</li>
            </ul>
        </section>

        <section class="verdict">
            <h2>চূড়ান্ত মূল্যায়ন</h2>
            <p><strong>গড় - শুধুমাত্র বাজেট বিকল্প</strong></p>
            <p>এয়ারটেল কেবল সেই ব্যবহারকারীদের জন্য উপযুক্ত যারা কভারেজ এলাকায় রয়েছেন এবং বাজেট-মুখী অপারেটর চান; নেটওয়ার্ক মান অগ্রাধিকার হলে অন্য অপশন বিবেচনা করুন।</p>
        </section>
    </article>
</body>
</html>`

	// Additional details JSON
	additionalDetails := datatypes.JSONMap{
		"market_rank":             4,
		"subscribers_millions":    8,
		"overall_rating":          3.05,
		"network_quality_rating":  3.0,
		"coverage_rating":         3.0,
		"data_speed_rating":       2.9,
		"pricing_rating":          3.8,
		"package_variety_rating":  3.0,
		"customer_support_rating": 2.8,
		"value_for_money_rating":  3.2,
		"avg_data_speed_mbps":     10,
		"peak_hour_speed_mbps":    5,
		"network_uptime_percent":  96.5,
		"latency_ms":              45,
		"urban_coverage_percent":  90.0,
		"rural_coverage_percent":  35.0,
		"hotline_number":          "121",
		"website_url":             "https://www.airtel.com.bd",
		"verdict":                 "Average - Budget Option Only",
		"recommendation_score":    5,
	}

	detailsJSON, err := json.Marshal(additionalDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Update English review
	result := db.Model(&map[string]interface{}{}).
		Table("product_reviews").
		Where("id = ? AND product_id = (SELECT id FROM products WHERE slug = ?)", 3, "airtel").
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
		"market_rank":          4,
		"subscribers_millions": 8,
		"overall_rating":       3.05,
		"verdict":              "গড়বেনে - শুধুমাত্র বাজেট বিকল্প",
		"recommendation_score": 5,
	}

	bengaliJSON, err := json.Marshal(bengaliDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal Bengali JSON: %w", err)
	}

	result = db.Model(&map[string]interface{}{}).
		Table("product_review_translations").
		Where("product_review_id = ? AND locale = ?", 3, "bn").
		Updates(map[string]interface{}{
			"reviews":            bengaliReview,
			"additional_details": string(bengaliJSON),
			"updated_at":         gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if result.Error != nil {
		log.Printf("Error updating Bangla translation: %v", result.Error)
		return result.Error
	}

	fmt.Println("✅ Airtel reviews updated successfully!")

	return nil
}
