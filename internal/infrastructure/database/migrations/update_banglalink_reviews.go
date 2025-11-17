package migrations

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// UpdateBanglalinkReviewsWithComparison updates product reviews with HTML comparison content
func UpdateBanglalinkReviewsWithComparison(db *gorm.DB) error {
	fmt.Println("\n========== UPDATING BANGLALINK REVIEWS WITH HTML COMPARISON ==========")

	// English review HTML content
	englishReview := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Banglalink Review</title>
</head>
<body>
    <article class="banglalink-review">
        <header>
            <h1>Banglalink: Best Value Network in Bangladesh</h1>
            <p class="rating"><strong>Overall Rating: 4.1/5 ⭐⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>Executive Summary</h2>
            <p>Banglalink is Bangladesh's second-largest mobile operator with 35+ million subscribers, offering the best value for money with competitive pricing, solid 4G coverage, and diverse service offerings. Ranked #2 in the market with a 4.1/5 overall rating, Banglalink is the ideal choice for budget-conscious users, students, and professionals seeking reliable connectivity without premium pricing.</p>
        </section>

        <section class="network-quality">
            <h2>Network Quality &amp; Coverage</h2>
            <p><strong>Rating: 4.0/5</strong></p>
            <ul>
                <li>Urban Coverage: 99% with consistent 4G availability</li>
                <li>Rural Coverage: 60% with ongoing expansion</li>
                <li>Average Speed: 14 Mbps download, 8 Mbps peak hours</li>
                <li>Network Uptime: 97.5% reliability</li>
                <li>Latency: 25ms average (good for gaming)</li>
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
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">Position vs Banglalink</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Grameenphone</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#1</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">4.65/5 ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Best Quality</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Premium network, 14% more expensive</td>
                    </tr>
                    <tr style="background-color: #fff3cd;">
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Banglalink</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#2</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;"><strong>4.1/5 ⭐</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Best Value</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>CURRENT - Best price-to-quality ratio</strong></td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Robi</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#3</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">4.05/5 ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Good Alternative</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Similar quality, slightly cheaper data</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Airtel</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#4</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">3.05/5 ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Regional Option</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Growing network, limited coverage</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>Teletalk</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">#5</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">2.85/5 ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Budget Only</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">Cheapest but outdated infrastructure</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="pricing">
            <h2>Pricing Comparison</h2>
            <p><strong>Rating: 4.5/5 (Most Competitive)</strong></p>
            <table border="1" cellpadding="10" cellspacing="0" style="width: 100%; border-collapse: collapse;">
                <thead>
                    <tr style="background-color: #f5f5f5;">
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">Service</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Banglalink</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Grameenphone</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Robi</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Airtel</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">Teletalk</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>100MB Daily</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">50 BDT ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">80 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">60 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">50 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">40 BDT</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>1GB Weekly</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">120 BDT ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">180 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">140 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">130 BDT</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">100 BDT</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="pros">
            <h2>Pros (11 Major Advantages)</h2>
            <ul>
                <li>✅ <strong>Affordable call rates</strong> - Competitive pricing for local and international calls</li>
                <li>✅ <strong>Attractive data bundles</strong> - Regular promotions with bonus offers</li>
                <li>✅ <strong>Good 4G coverage</strong> - 99% urban and 60% rural coverage</li>
                <li>✅ <strong>User-friendly app</strong> - Intuitive account management interface</li>
                <li>✅ <strong>Digital payment support</strong> - Integrated with multiple payment platforms</li>
                <li>✅ <strong>Emergency balance</strong> - Feature for urgent calls when balance low</li>
                <li>✅ <strong>Flexible plans</strong> - Prepaid and postpaid options available</li>
                <li>✅ <strong>Frequent promotions</strong> - Regular special offers and bonuses</li>
                <li>✅ <strong>Social media bundles</strong> - Unlimited access to popular platforms</li>
                <li>✅ <strong>Money services</strong> - Comprehensive digital payment ecosystem</li>
                <li>✅ <strong>Multiple support channels</strong> - Phone, app, USSD, and social media</li>
            </ul>
        </section>

        <section class="cons">
            <h2>Cons (7 Main Limitations)</h2>
            <ul>
                <li>⚠️ <strong>Peak hour congestion</strong> - Network slowdowns during 6-10 PM</li>
                <li>⚠️ <strong>Limited rural coverage</strong> - Only 60% coverage in villages</li>
                <li>⚠️ <strong>Long customer wait times</strong> - Support queues during peak periods</li>
                <li>⚠️ <strong>Higher roaming rates</strong> - More expensive than competitors internationally</li>
                <li>⚠️ <strong>Speed throttling</strong> - Occasional limitations on some packages</li>
                <li>⚠️ <strong>Complex SIM activation</strong> - KYC process could be simplified</li>
                <li>⚠️ <strong>Restrictive validity</strong> - Some packages have strict expiry periods</li>
            </ul>
        </section>

        <section class="verdict">
            <h2>Verdict</h2>
            <p><strong>✅ HIGHLY RECOMMENDED - Best Value for Money</strong></p>
            <p>Banglalink offers the best value-to-quality ratio in Bangladesh's mobile market. As the #2 operator with 35+ million subscribers, it provides reliable 4G coverage in urban areas, competitive pricing that undercuts Grameenphone by 14%, and a growing digital services ecosystem.</p>
        </section>
    </article>
</body>
</html>`

	// Bangla (Bengali) review - PURE BENGALI SCRIPT
	bengaliReview := `<!DOCTYPE html>
<html lang="bn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>বাংলালিংক রিভিউ</title>
</head>
<body>
    <article class="banglalink-review-bn">
        <header>
            <h1>বাংলালিংক: বাংলাদেশের সেরা মূল্যের নেটওয়ার্ক</h1>
            <p class="rating"><strong>সামগ্রিক রেটিং: ৪.১/৫ ⭐⭐⭐⭐</strong></p>
        </header>

        <section class="executive-summary">
            <h2>সংক্ষিপ্ত পর্যালোচনা</h2>
            <p>বাংলালিংক বাংলাদেশের দ্বিতীয় বৃহত্তম মোবাইল অপারেটর যা ৩৫+ মিলিয়ন গ্রাহকদের সেবা প্রদান করে। প্রতিযোগিতামূলক মূল্য, শক্তিশালী ৪জি কভারেজ এবং বৈচিত্র্যময় সেবার মাধ্যমে সেরা মূল্যের মূল্য প্রদান করে। বাজারে ২নম্বর স্থানে ৪.১/৫ রেটিং সহ, বাংলালিংক বাজেট-সচেতন ব্যবহারকারী, শিক্ষার্থী এবং নির্ভরযোগ্য সংযোগ খোঁজেন এমন পেশাদারদের জন্য আদর্শ।</p>
        </section>

        <section class="network-quality">
            <h2>নেটওয়ার্ক গুণমান এবং কভারেজ</h2>
            <p><strong>রেটিং: ৪.০/৫</strong></p>
            <ul>
                <li>শহরের কভারেজ: ৯৯% ধারাবাহিক ৪জি উপলব্ধতা</li>
                <li>গ্রামীণ কভারেজ: ৬০% এবং ক্রমাগত সম্প্রসারণ</li>
                <li>গড় গতি: ১৪ এমবিপিএস ডাউনলোড, ৮ এমবিপিএস পিক আওয়ার্স</li>
                <li>নেটওয়ার্ক আপটাইম: ৯৭.৫% নির্ভরযোগ্যতা</li>
                <li>লেটেন্সি: ২৫ এমএস গড় (গেমিংয়ের জন্য ভাল)</li>
            </ul>
        </section>

        <section class="market-comparison">
            <h2>বাজার তুলনা</h2>
            <table border="1" cellpadding="10" cellspacing="0" style="width: 100%; border-collapse: collapse;">
                <thead>
                    <tr style="background-color: #f5f5f5;">
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">অপারেটর</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">র্যাঙ্ক</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">রেটিং</th>
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">সেরা</th>
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">বাংলালিংকের বিপরীতে</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>গ্রামীণফোন</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">১ম</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">৪.৬৫/৫ ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">সেরা মান</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">প্রিমিয়াম নেটওয়ার্ক, ১৪% বেশি দামী</td>
                    </tr>
                    <tr style="background-color: #fff3cd;">
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>বাংলালিংক</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">২য়</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;"><strong>৪.১/৫ ⭐</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>সেরা মূল্য</strong></td>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>বর্তমান - সেরা মূল্য-গুণমান অনুপাত</strong></td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>রোবি</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">৩য়</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">৪.০৫/৫ ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">ভাল বিকল্প</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">অনুরূপ মান, সামান্য সস্তা ডেটা</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>এয়ারটেল</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">৪র্থ</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">৩.০৫/৫ ⭐</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">আঞ্চলিক বিকল্প</td>
                        <td style="padding: 10px; border: 1px solid #ddd;">ক্রমবর্ধমান নেটওয়ার্ক, সীমিত কভারেজ</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="pricing">
            <h2>মূল্য তুলনা</h2>
            <p><strong>রেটিং: ৪.৫/৫ (সবচেয়ে প্রতিযোগিতামূলক)</strong></p>
            <table border="1" cellpadding="10" cellspacing="0" style="width: 100%; border-collapse: collapse;">
                <thead>
                    <tr style="background-color: #f5f5f5;">
                        <th style="text-align: left; padding: 10px; border: 1px solid #ddd;">সেবা</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">বাংলালিংক</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">গ্রামীণফোন</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">রোবি</th>
                        <th style="text-align: center; padding: 10px; border: 1px solid #ddd;">এয়ারটেল</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>১০০ এমবি দৈনিক</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">৫০ টাকা ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">৮০ টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">৬০ টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">৫০ টাকা</td>
                    </tr>
                    <tr>
                        <td style="padding: 10px; border: 1px solid #ddd;"><strong>১ গিগাবাইট সাপ্তাহিক</strong></td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd; background-color: #d4edda;">১২০ টাকা ⭐</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">১৮০ টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">১৪০ টাকা</td>
                        <td style="text-align: center; padding: 10px; border: 1px solid #ddd;">১৩০ টাকা</td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="pros">
            <h2>সুবিধাসমূহ (১১টি প্রধান সুবিধা)</h2>
            <ul>
                <li>✅ <strong>সাশ্রয়ী কল রেট</strong> - স্থানীয় এবং আন্তর্জাতিক কলের প্রতিযোগিতামূলক মূল্য</li>
                <li>✅ <strong>আকর্ষণীয় ডেটা প্যাকেজ</strong> - নিয়মিত প্রচার এবং বোনাস অফার</li>
                <li>✅ <strong>ভাল ৪জি কভারেজ</strong> - ৯৯% শহুরে এবং ৬০% গ্রামীণ কভারেজ</li>
                <li>✅ <strong>ব্যবহারকারী-বান্ধব অ্যাপ</strong> - সহজ অ্যাকাউন্ট ব্যবস্থাপনা ইন্টারফেস</li>
                <li>✅ <strong>ডিজিটাল পেমেন্ট সমর্থন</strong> - একাধিক পেমেন্ট প্ল্যাটফর্মের সাথে সমন্বয়</li>
                <li>✅ <strong>জরুরি ব্যালেন্স</strong> - ব্যালেন্স কম থাকলে জরুরি কলের সুবিধা</li>
                <li>✅ <strong>নমনীয় পরিকল্পনা</strong> - প্রিপেইড এবং পোস্টপেইড বিকল্প উপলব্ধ</li>
                <li>✅ <strong>ঘন ঘন প্রচার</strong> - নিয়মিত বিশেষ অফার এবং বোনাস</li>
                <li>✅ <strong>সোশ্যাল মিডিয়া প্যাকেজ</strong> - জনপ্রিয় প্ল্যাটফর্মে সীমাহীন অ্যাক্সেস</li>
                <li>✅ <strong>মানি সেবা</strong> - ব্যাপক ডিজিটাল পেমেন্ট ইকোসিস্টেম</li>
                <li>✅ <strong>একাধিক সহায়তা চ্যানেল</strong> - ফোন, অ্যাপ, ইউএসএসডি এবং সোশ্যাল মিডিয়া</li>
            </ul>
        </section>

        <section class="cons">
            <h2>অসুবিধাসমূহ (৭টি প্রধান সীমাবদ্ধতা)</h2>
            <ul>
                <li>⚠️ <strong>পিক আওয়ার ভিড়</strong> - সন্ধ্যা ৬-১০টায় নেটওয়ার্ক মন্থরতা</li>
                <li>⚠️ <strong>সীমিত গ্রামীণ কভারেজ</strong> - গ্রামে মাত্র ৬০% কভারেজ</li>
                <li>⚠️ <strong>দীর্ঘ গ্রাহক অপেক্ষা সময়</strong> - পিক সময়ে সহায়তা সারি</li>
                <li>⚠️ <strong>উচ্চতর রোমিং হার</strong> - আন্তর্জাতিকভাবে প্রতিযোগীদের চেয়ে বেশি ব্যয়বহুল</li>
                <li>⚠️ <strong>গতি থ্রটলিং</strong> - কিছু প্যাকেজে সাময়িক সীমাবদ্ধতা</li>
                <li>⚠️ <strong>জটিল সিম অ্যাক্টিভেশন</strong> - কেওয়াইসি প্রক্রিয়া আরও সহজ করা যেতে পারে</li>
                <li>⚠️ <strong>সীমাবদ্ধ বৈধতা</strong> - কিছু প্যাকেজের কঠোর মেয়াদ শেষ সময়</li>
            </ul>
        </section>

        <section class="verdict">
            <h2>চূড়ান্ত মূল্যায়ন</h2>
            <p><strong>✅ অত্যন্ত সুপারিশকৃত - সেরা মূল্য-গুণমান অনুপাত</strong></p>
            <p>বাংলালিংক বাংলাদেশের মোবাইল বাজারে সেরা মূল্য-গুণমান অনুপাত প্রদান করে। ৩৫+ মিলিয়ন গ্রাহক সহ দ্বিতীয় বৃহত্তম অপারেটর হিসেবে, এটি শহুরে এলাকায় নির্ভরযোগ্য ৪জি কভারেজ, গ্রামীণফোনের চেয়ে ১৪% কম মূল্য এবং ক্রমবর্ধমান ডিজিটাল সেবা ইকোসিস্টেম প্রদান করে।</p>
        </section>
    </article>
</body>
</html>`

	// Additional details JSON
	additionalDetails := datatypes.JSONMap{
		"market_rank":             2,
		"subscribers_millions":    35,
		"overall_rating":          4.1,
		"network_quality_rating":  4.0,
		"coverage_rating":         4.0,
		"data_speed_rating":       4.0,
		"pricing_rating":          4.5,
		"package_variety_rating":  4.0,
		"customer_support_rating": 3.5,
		"value_for_money_rating":  4.25,
		"avg_data_speed_mbps":     14,
		"peak_hour_speed_mbps":    8,
		"network_uptime_percent":  97.5,
		"latency_ms":              25,
		"urban_coverage_percent":  99.0,
		"rural_coverage_percent":  60.0,
		"hotline_number":          "111",
		"website_url":             "https://www.banglalink.com.bd",
		"verdict":                 "Highly Recommended - Best Value for Money",
		"recommendation_score":    8,
	}

	detailsJSON, err := json.Marshal(additionalDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Update English review
	result := db.Model(&map[string]interface{}{}).
		Table("product_reviews").
		Where("id = ? AND product_id = (SELECT id FROM products WHERE slug = ?)", 240, "banglalink").
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
		"market_rank":          2,
		"subscribers_millions": 35,
		"overall_rating":       4.1,
		"verdict":              "অত্যন্ত সুপারিশকৃত - সেরা মূল্য-গুণমান অনুপাত",
		"recommendation_score": 8,
	}

	bengaliJSON, err := json.Marshal(bengaliDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal Bengali JSON: %w", err)
	}

	result = db.Model(&map[string]interface{}{}).
		Table("product_review_translations").
		Where("product_review_id = ? AND locale = ?", 240, "bn").
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

	fmt.Println("✅ Banglalink reviews updated successfully!")
	fmt.Printf("English review: %d rows affected\n", result.RowsAffected)

	return nil
}
