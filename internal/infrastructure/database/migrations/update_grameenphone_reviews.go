package migrations

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// UpdateGrameenphoneReviewsWithComparison updates product reviews with HTML comparison content
func UpdateGrameenphoneReviewsWithComparison(db *gorm.DB) error {
	fmt.Println("\n========== UPDATING GRAMEENPHONE REVIEWS WITH HTML COMPARISON ==========")

	// English review HTML content
	englishReview := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Grameenphone Review</title>
</head>
<body>
    <article class="grameenphone-review">
        <header>
            <h1>Grameenphone: Premium Quality Network in Bangladesh</h1>
            <p class="rating"><strong>Overall Rating: 4.65/5 ⭐⭐⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>Executive Summary</h2>
            <p>Grameenphone is Bangladesh's largest mobile operator with 50+ million subscribers, offering the highest quality network coverage with premium services. Ranked #1 in the market with a 4.65/5 overall rating, Grameenphone is the ideal choice for users seeking the best network quality and most comprehensive service offerings, albeit at a premium price point.</p>
        </section>

        <section class="network-quality">
            <h2>Network Quality &amp; Coverage</h2>
            <p><strong>Rating: 4.65/5</strong></p>
            <ul>
                <li>Urban Coverage: 99.5% with exceptional 4G availability</li>
                <li>Rural Coverage: 75% with continuous expansion</li>
                <li>Average Speed: 18 Mbps download, 10 Mbps peak hours</li>
                <li>Network Uptime: 99.2% reliability</li>
                <li>Latency: 20ms average (excellent for gaming and streaming)</li>
            </ul>
        </section>

        <section class="market-comparison">
            <h2>Market Comparison</h2>
            <table border="1" cellpadding="10" cellspacing="0" style="width: 100%; border-collapse: collapse;">
                <thead>
                    <tr style="background-color: #f5f5f5;">
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">Operator</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Rank</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Rating</th>
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">Best For</th>
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">Position vs Grameenphone</th>
                    </tr>
                </thead>
                <tbody>
                    <tr style="background-color: #fff3cd;">
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Grameenphone</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;"><strong>#1</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;"><strong>4.65/5 ⭐</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Best Quality</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>CURRENT - Premium network leader</strong></td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Banglalink</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#2</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">4.1/5 ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Best Value</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Good budget alternative, 14% cheaper</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Robi</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#3</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">4.05/5 ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Good Alternative</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Lower quality, better pricing</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Airtel</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#4</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">3.05/5 ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Regional Option</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Limited coverage and services</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="pricing">
            <h2>Pricing Comparison</h2>
            <p><strong>Rating: 3.8/5 (Premium Pricing)</strong></p>
            <table border="1" cellpadding="10" cellspacing="0" style="width: 100%; border-collapse: collapse;">
                <thead>
                    <tr style="background-color: #f5f5f5;">
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">Service</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Grameenphone</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Banglalink</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Robi</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>100MB Daily</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">80 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">50 BDT ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">60 BDT</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>1GB Weekly</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">180 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">120 BDT ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">140 BDT</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="pros">
            <h2>Pros (10 Major Advantages)</h2>
            <ul>
                <li> <strong>Best network quality</strong> - Fastest speeds and most reliable coverage</li>
                <li> <strong>Widest coverage</strong> - 99.5% urban and 75% rural coverage</li>
                <li> <strong>Excellent customer service</strong> - 24/7 support across multiple channels</li>
                <li> <strong>Advanced mobile banking</strong> - bKash integration and digital services</li>
                <li> <strong>Generous data plans</strong> - Generous FUP and rollover data</li>
                <li> <strong>Premium roaming</strong> - Better international roaming rates compared to others</li>
                <li> <strong>Entertainment content</strong> - Music, video, and gaming bundles</li>
                <li> <strong>Business solutions</strong> - Dedicated corporate packages</li>
                <li> <strong>Loyalty rewards</strong> - GP points and exclusive member benefits</li>
                <li> <strong>Latest technology</strong> - 5G readiness and modern infrastructure</li>
            </ul>
        </section>

        <section class="cons">
            <h2>Cons (5 Main Limitations)</h2>
            <ul>
                <li> <strong>Premium pricing</strong> - 14-25% more expensive than competitors</li>
                <li> <strong>Complex pricing structure</strong> - Multiple plans can be confusing</li>
                <li> <strong>Congestion in peak hours</strong> - Network slowdowns during busy times</li>
                <li> <strong>High activation costs</strong> - SIM cards and setup fees are premium</li>
                <li> <strong>Limited budget options</strong> - Few entry-level plans available</li>
            </ul>
        </section>

        <section class="verdict">
            <h2>Verdict</h2>
            <p><strong>HIGHLY RECOMMENDED - Best Quality Network</strong></p>
            <p>Grameenphone is Bangladesh's premium mobile operator offering the best network quality, fastest speeds, and most comprehensive services. While pricing is higher than competitors, the superior network quality and service excellence make it worth the investment for users who prioritize reliability and performance.</p>
        </section>
    </article>
</body>
</html>`

	// Bangla (Bengali) review
	bengaliReview := `<!DOCTYPE html>
<html lang="bn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>গ্রামীণফোন রিভিউ</title>
</head>
<body>
    <article class="grameenphone-review-bn">
        <header>
            <h1>গ্রামীণফোন: বাংলাদেশের সেরা মানের নেটওয়ার্ক</h1>
            <p class="rating"><strong>সামগ্রিক রেটিং: ৪.৬৫/৫ ⭐⭐⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>সংক্ষিপ্ত পর্যালোচনা</h2>
            <p>গ্রামীণফোন বাংলাদেশের বৃহত্তম মোবাইল অপারেটর ৫০+ মিলিয়ন গ্রাহক সেবা প্রদান করে প্রিমিয়াম মানের নেটওয়ার্ক কভারেজ সহ। বাজারে ১নম্বর স্থানে ৪.৬৫/৫ রেটিং সহ, গ্রামীণফোন সর্বোত্তম নেটওয়ার্ক মান এবং ব্যাপক সেবা খোঁজেন এমন ব্যবহারকারীদের জন্য আদর্শ।</p>
        </section>

        <section class="network-quality">
            <h2>নেটওয়ার্ক গুণমান এবং কভারেজ</h2>
            <p><strong>রেটিং: ৪.৬৫/৫</strong></p>
            <ul>
                <li>শহরের কভারেজ: ৯৯.৫% ব্যতিক্রমী ৪জি উপলব্ধতা</li>
                <li>গ্রামীণ কভারেজ: ৭৫% ক্রমাগত সম্প্রসারণ সহ</li>
                <li>গড় গতি: ১৮ এমবিপিএস ডাউনলোড, ১০ এমবিপিএস পিক আওয়ার্স</li>
                <li>নেটওয়ার্ক আপটাইম: ৯৯.২% নির্ভরযোগ্যতা</li>
                <li>লেটেন্সি: ২০ এমএস গড় (গেমিং এবং স্ট্রিমিংয়ের জন্য চমৎকার)</li>
            </ul>
        </section>

        <section class="pros">
            <h2>সুবিধাসমূহ (১০টি প্রধান সুবিধা)</h2>
            <ul>
                <li> <strong>সেরা নেটওয়ার্ক মান</strong> - সবচেয়ে দ্রুত গতি এবং নির্ভরযোগ্য কভারেজ</li>
                <li> <strong>সবচেয়ে বিস্তৃত কভারেজ</strong> - ৯৯.৫% শহুরে এবং ৭৫% গ্রামীণ</li>
                <li> <strong>চমৎকার গ্রাহক সেবা</strong> - ২৪/৭ সহায়তা একাধিক চ্যানেল জুড়ে</li>
                <li> <strong>উন্নত মোবাইল ব্যাংকিং</strong> - বিকাশ একীকরণ এবং ডিজিটাল সেবা</li>
                <li> <strong>উদার ডেটা প্ল্যান</strong> - উদার এফইউপি এবং রোলওভার ডেটা</li>
                <li> <strong>প্রিমিয়াম রোমিং</strong> - আন্তর্জাতিক রোমিং এ ভাল হার</li>
                <li> <strong>বিনোদন সামগ্রী</strong> - সঙ্গীত, ভিডিও এবং গেমিং বান্ডেল</li>
                <li> <strong>ব্যবসায়িক সমাধান</strong> - কর্পোরেট প্যাকেজ উপলব্ধ</li>
                <li> <strong>লয়্যালটি পুরস্কার</strong> - জিপি পয়েন্ট এবং এক্সক্লুসিভ সুবিধা</li>
                <li> <strong>অত্যাধুনিক প্রযুক্তি</strong> - ৫জি প্রস্তুতি এবং আধুনিক অবকাঠামো</li>
            </ul>
        </section>

        <section class="cons">
            <h2>অসুবিধাসমূহ (৫টি প্রধান সীমাবদ্ধতা)</h2>
            <ul>
                <li> <strong>প্রিমিয়াম মূল্য</strong> - প্রতিযোগীদের চেয়ে ১৪-২৫% বেশি ব্যয়বহুল</li>
                <li> <strong>জটিল মূল্য নির্ধারণ</strong> - একাধিক পরিকল্পনা বিভ্রান্তিকর হতে পারে</li>
                <li> <strong>পিক আওয়ার ভিড়</strong> - ব্যস্ত সময়ে নেটওয়ার্ক মন্থরতা</li>
                <li> <strong>উচ্চ সক্রিয়করণ খরচ</strong> - সিম কার্ড এবং সেটআপ ফি প্রিমিয়াম</li>
                <li> <strong>সীমিত বাজেট বিকল্প</strong> - কয়েকটি প্রবেশ-স্তরের পরিকল্পনা উপলব্ধ</li>
            </ul>
        </section>

        <section class="verdict">
            <h2>চূড়ান্ত মূল্যায়ন</h2>
            <p><strong>অত্যন্ত সুপারিশকৃত - সেরা মানের নেটওয়ার্ক</strong></p>
            <p>গ্রামীণফোন বাংলাদেশের প্রিমিয়াম মোবাইল অপারেটর যা সেরা নেটওয়ার্ক মান, দ্রুততম গতি এবং সবচেয়ে ব্যাপক সেবা প্রদান করে।</p>
        </section>
    </article>
</body>
</html>`

	// Additional details JSON
	additionalDetails := datatypes.JSONMap{
		"market_rank":             1,
		"subscribers_millions":    50,
		"overall_rating":          4.65,
		"network_quality_rating":  4.65,
		"coverage_rating":         4.65,
		"data_speed_rating":       4.7,
		"pricing_rating":          3.8,
		"package_variety_rating":  4.5,
		"customer_support_rating": 4.6,
		"value_for_money_rating":  4.2,
		"avg_data_speed_mbps":     18,
		"peak_hour_speed_mbps":    10,
		"network_uptime_percent":  99.2,
		"latency_ms":              20,
		"urban_coverage_percent":  99.5,
		"rural_coverage_percent":  75.0,
		"hotline_number":          "121",
		"website_url":             "https://www.grameenphone.com",
		"verdict":                 "Highly Recommended - Best Quality Network",
		"recommendation_score":    9,
	}

	detailsJSON, err := json.Marshal(additionalDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Update English review
	result := db.Model(&map[string]interface{}{}).
		Table("product_reviews").
		Where("id = ? AND product_id = (SELECT id FROM products WHERE slug = ?)", 1, "grameenphone").
		Updates(map[string]interface{}{
			"reviews":            englishReview,
			"additional_details": string(detailsJSON),
			"updated_at":         gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if result.Error != nil {
		log.Printf("Error updating English review: %v", result.Error)
		return result.Error
	}

	log.Printf("Updated English review: %d rows affected", result.RowsAffected)

	// Update Bangla translation
	bengaliDetails := datatypes.JSONMap{
		"language":             "Bengali",
		"market_rank":          1,
		"subscribers_millions": 50,
		"overall_rating":       4.65,
		"verdict":              "অত্যন্ত সুপারিশকৃত - সেরা মানের নেটওয়ার্ক",
		"recommendation_score": 9,
	}

	bengaliJSON, err := json.Marshal(bengaliDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal Bengali JSON: %w", err)
	}

	result = db.Model(&map[string]interface{}{}).
		Table("product_review_translations").
		Where("product_review_id = ? AND locale = ?", 1, "bn").
		Updates(map[string]interface{}{
			"reviews":            bengaliReview,
			"additional_details": string(bengaliJSON),
			"updated_at":         gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if result.Error != nil {
		log.Printf("Error updating Bangla translation: %v", result.Error)
		return result.Error
	}

	log.Printf("Updated Bangla translation: %d rows affected", result.RowsAffected)

	fmt.Println("✅ Grameenphone reviews updated successfully!")

	return nil
}
