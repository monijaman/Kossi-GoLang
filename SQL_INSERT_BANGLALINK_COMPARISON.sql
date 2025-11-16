-- ============================================================================
-- BANGLALINK MOBILE OPERATOR - DATABASE INSERTION SCRIPT
-- ============================================================================
-- Created: November 17, 2025
-- Purpose: Insert comprehensive Banglalink reviews and competitor comparison data
-- ============================================================================

-- Enable necessary extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- ============================================================================
-- 1. PRODUCT REVIEWS - Banglalink English Review (Already seeded, verification)
-- ============================================================================

-- Verify Banglalink review exists
SELECT id, product_id, rating, status FROM product_reviews 
WHERE product_id = (SELECT id FROM products WHERE slug = 'banglalink') 
ORDER BY id DESC LIMIT 1;

-- ============================================================================
-- 2. CREATE OPERATOR COMPARISON TABLE
-- ============================================================================

CREATE TABLE IF NOT EXISTS mobile_operator_comparisons (
    id SERIAL PRIMARY KEY,
    operator_name VARCHAR(255) NOT NULL,
    operator_slug VARCHAR(255) NOT NULL UNIQUE,
    market_rank INT,
    subscribers_millions INT,
    overall_rating DECIMAL(3,2),
    
    -- Quality Metrics
    network_quality_rating DECIMAL(3,2),
    coverage_rating DECIMAL(3,2),
    data_speed_rating DECIMAL(3,2),
    pricing_rating DECIMAL(3,2),
    package_variety_rating DECIMAL(3,2),
    customer_support_rating DECIMAL(3,2),
    value_for_money_rating DECIMAL(3,2),
    
    -- Specifications
    avg_data_speed_mbps INT,
    peak_hour_speed_mbps INT,
    network_uptime_percent DECIMAL(5,2),
    latency_ms INT,
    
    -- Coverage
    urban_coverage_percent DECIMAL(5,2),
    rural_coverage_percent DECIMAL(5,2),
    
    -- Service Info
    hotline_number VARCHAR(20),
    website_url VARCHAR(500),
    has_4g BOOLEAN DEFAULT true,
    has_5g_plans BOOLEAN DEFAULT false,
    
    -- Digital Services
    mobile_wallet_name VARCHAR(255),
    has_money_transfer BOOLEAN DEFAULT true,
    has_bill_payments BOOLEAN DEFAULT true,
    has_merchant_payments BOOLEAN DEFAULT false,
    
    -- Best For
    best_for_students BOOLEAN DEFAULT false,
    best_for_professionals BOOLEAN DEFAULT false,
    best_for_budget_users BOOLEAN DEFAULT false,
    best_for_heavy_users BOOLEAN DEFAULT false,
    best_for_business BOOLEAN DEFAULT false,
    
    -- Recommendations
    recommendation_score INT, -- 1-10 scale
    recommendation_text TEXT,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- ============================================================================
-- 3. INSERT OPERATOR COMPARISON DATA
-- ============================================================================

INSERT INTO mobile_operator_comparisons (
    operator_name, operator_slug, market_rank, subscribers_millions,
    overall_rating, network_quality_rating, coverage_rating, data_speed_rating,
    pricing_rating, package_variety_rating, customer_support_rating,
    value_for_money_rating, avg_data_speed_mbps, peak_hour_speed_mbps,
    network_uptime_percent, latency_ms, urban_coverage_percent,
    rural_coverage_percent, hotline_number, website_url, has_4g, has_5g_plans,
    mobile_wallet_name, has_money_transfer, has_bill_payments,
    has_merchant_payments, best_for_students, best_for_professionals,
    best_for_budget_users, best_for_heavy_users, best_for_business,
    recommendation_score, recommendation_text
) VALUES
-- Banglalink
(
    'Banglalink', 'banglalink', 2, 35,
    4.05, 4.0, 4.0, 4.0,
    4.5, 4.0, 3.5,
    4.25, 14, 8,
    97.5, 25, 99.0,
    60.0, '111', 'https://www.banglalink.com.bd', true, true,
    'Banglalink Money', true, true,
    true, true, true,
    true, true, false,
    8, 'Best value for money in Bangladesh with solid 4G coverage and competitive packages. Excellent choice for students and budget-conscious users.'
),
-- Grameenphone
(
    'Grameenphone', 'grameenphone', 1, 90,
    4.65, 5.0, 5.0, 5.0,
    2.5, 5.0, 5.0,
    3.0, 18, 12,
    99.5, 18, 99.5,
    75.0, '121', 'https://www.grameenphone.com', true, true,
    'bKash', true, true,
    true, false, true,
    false, false, true,
    9, 'Market leader with best network quality and comprehensive services. Premium pricing justified by superior reliability and ecosystem.'
),
-- Robi
(
    'Robi', 'robi', 3, 50,
    4.05, 4.0, 4.0, 4.0,
    3.5, 4.0, 4.0,
    3.5, 15, 10,
    98.0, 22, 99.0,
    70.0, '16123', 'https://www.robi.com.bd', true, true,
    'Rocket', true, true,
    true, true, true,
    true, false, false,
    8, 'Well-balanced option with good coverage and reliable service. Offers reasonable pricing with decent package variety.'
),
-- Airtel
(
    'Airtel', 'airtel', 4, 15,
    3.05, 3.5, 3.0, 3.0,
    3.5, 2.5, 3.5,
    3.0, 10, 7,
    96.0, 30, 95.0,
    40.0, '121', 'https://www.airtel.com.bd', true, false,
    'Airtel Money', true, true,
    true, false, false,
    false, false, false,
    6, 'Growing network with competitive rates. Limited coverage in rural areas but good for specific regions with adequate infrastructure.'
),
-- Teletalk
(
    'Teletalk', 'teletalk', 5, 5,
    2.85, 3.0, 3.5, 2.5,
    4.5, 2.0, 2.5,
    4.0, 6, 3,
    95.0, 35, 90.0,
    50.0, '121', 'https://www.teletalk.com.bd', true, false,
    'TT Money', true, true,
    false, false, false,
    false, false, false,
    5, 'Budget-friendly option with cheapest rates. Limited services and older technology, only recommended for absolute cost minimizers.'
);

-- ============================================================================
-- 4. CREATE PRICING COMPARISON TABLE
-- ============================================================================

CREATE TABLE IF NOT EXISTS operator_pricing_comparison (
    id SERIAL PRIMARY KEY,
    operator_id INT NOT NULL REFERENCES mobile_operator_comparisons(id),
    operator_name VARCHAR(255),
    service_type VARCHAR(100), -- 'data', 'voice', 'sms', 'roaming'
    package_name VARCHAR(255),
    volume VARCHAR(50), -- e.g., "1GB", "100 min", "100 SMS"
    duration VARCHAR(50), -- e.g., "1 day", "7 days", "30 days"
    price_bdt INT,
    currency VARCHAR(3) DEFAULT 'BDT',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================================
-- 5. INSERT PRICING DATA
-- ============================================================================

-- Banglalink Data Pricing
INSERT INTO operator_pricing_comparison (operator_name, service_type, package_name, volume, duration, price_bdt)
VALUES
('Banglalink', 'data', 'Daily Bundle 1', '100MB', '1 Day', 50),
('Banglalink', 'data', 'Daily Bundle 2', '250MB', '1 Day', 70),
('Banglalink', 'data', 'Daily Bundle 3', '500MB', '1 Day', 100),
('Banglalink', 'data', 'Weekly Bundle 1', '1GB', '7 Days', 120),
('Banglalink', 'data', 'Weekly Bundle 2', '2GB', '7 Days', 170),
('Banglalink', 'data', 'Weekly Bundle 3', '3GB', '7 Days', 200),
('Banglalink', 'data', 'Monthly Unlimited', 'Unlimited', '30 Days', 999),
('Banglalink', 'data', 'Night Package', '500MB (12AM-6AM)', '30 Days', 150);

-- Banglalink Voice Pricing
INSERT INTO operator_pricing_comparison (operator_name, service_type, package_name, volume, duration, price_bdt)
VALUES
('Banglalink', 'voice', 'Local Call Bundle', '100 min', '1 Day', 50),
('Banglalink', 'voice', 'Local Call Bundle', '500 min', '7 Days', 120),
('Banglalink', 'voice', 'Local Call Bundle', '1000 min', '30 Days', 300),
('Banglalink', 'voice', 'On-Net Call', 'Per minute', 'Per call', 1),
('Banglalink', 'voice', 'International (India)', 'Per minute', 'Per call', 6);

-- Banglalink SMS Pricing
INSERT INTO operator_pricing_comparison (operator_name, service_type, package_name, volume, duration, price_bdt)
VALUES
('Banglalink', 'sms', 'Local SMS', 'Per SMS', 'Per SMS', 1),
('Banglalink', 'sms', 'SMS Bundle', '100 SMS', '7 Days', 30),
('Banglalink', 'sms', 'International SMS', 'To India', 'Per SMS', 3),
('Banglalink', 'sms', 'International SMS', 'General', 'Per SMS', 5);

-- Banglalink Roaming Pricing
INSERT INTO operator_pricing_comparison (operator_name, service_type, package_name, volume, duration, price_bdt, notes)
VALUES
('Banglalink', 'roaming', 'India Roaming', 'Daily Pass', '1 Day', 75, 'Calls, SMS, Data'),
('Banglalink', 'roaming', 'Thailand Roaming', 'Daily Pass', '1 Day', 100, 'Calls, SMS, Data'),
('Banglalink', 'roaming', 'USA Roaming', 'Daily Pass', '1 Day', 200, 'Calls, SMS, Data'),
('Banglalink', 'roaming', 'Data Roaming Asia', 'Per MB', 'Pay-as-you-go', 25, 'Per MB rate');

-- Grameenphone Data Pricing
INSERT INTO operator_pricing_comparison (operator_name, service_type, package_name, volume, duration, price_bdt)
VALUES
('Grameenphone', 'data', 'Daily Bundle 1', '100MB', '1 Day', 80),
('Grameenphone', 'data', 'Daily Bundle 2', '250MB', '1 Day', 120),
('Grameenphone', 'data', 'Weekly Bundle 1', '1GB', '7 Days', 180),
('Grameenphone', 'data', 'Weekly Bundle 2', '2GB', '7 Days', 300),
('Grameenphone', 'data', 'Monthly Unlimited', 'Unlimited', '30 Days', 1500);

-- Robi Data Pricing
INSERT INTO operator_pricing_comparison (operator_name, service_type, package_name, volume, duration, price_bdt)
VALUES
('Robi', 'data', 'Daily Bundle 1', '100MB', '1 Day', 60),
('Robi', 'data', 'Weekly Bundle 1', '1GB', '7 Days', 140),
('Robi', 'data', 'Monthly Unlimited', 'Unlimited', '30 Days', 1200);

-- Airtel Data Pricing
INSERT INTO operator_pricing_comparison (operator_name, service_type, package_name, volume, duration, price_bdt)
VALUES
('Airtel', 'data', 'Daily Bundle 1', '100MB', '1 Day', 50),
('Airtel', 'data', 'Weekly Bundle 1', '1GB', '7 Days', 130),
('Airtel', 'data', 'Monthly Unlimited', 'Unlimited', '30 Days', 999);

-- Teletalk Data Pricing
INSERT INTO operator_pricing_comparison (operator_name, service_type, package_name, volume, duration, price_bdt)
VALUES
('Teletalk', 'data', 'Daily Bundle 1', '100MB', '1 Day', 40),
('Teletalk', 'data', 'Weekly Bundle 1', '1GB', '7 Days', 100),
('Teletalk', 'data', 'Monthly Unlimited', 'Unlimited', '30 Days', 800);

-- ============================================================================
-- 6. CREATE FEATURE COMPARISON TABLE
-- ============================================================================

CREATE TABLE IF NOT EXISTS operator_features (
    id SERIAL PRIMARY KEY,
    operator_id INT NOT NULL REFERENCES mobile_operator_comparisons(id),
    operator_name VARCHAR(255),
    feature_category VARCHAR(100), -- 'data', 'social_media', 'voice', 'digital', 'support'
    feature_name VARCHAR(255),
    is_available BOOLEAN DEFAULT true,
    details TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================================
-- 7. INSERT FEATURE DATA
-- ============================================================================

-- Banglalink Features
INSERT INTO operator_features (operator_name, feature_category, feature_name, is_available, details)
VALUES
-- Data Features
('Banglalink', 'data', 'Daily Bundles', true, '100MB-500MB per day at affordable rates'),
('Banglalink', 'data', 'Weekly Bundles', true, '1GB-3GB valid for 7 days'),
('Banglalink', 'data', 'Monthly Unlimited', true, 'True unlimited data with fair usage policy'),
('Banglalink', 'data', 'Night Packages', true, 'Discounted data from 12 AM to 6 AM'),
('Banglalink', 'data', 'OTT Bundles', true, 'Special rates for Netflix, YouTube, etc.'),

-- Social Media Features
('Banglalink', 'social_media', 'WhatsApp Bundle', true, 'Unlimited WhatsApp without data deduction'),
('Banglalink', 'social_media', 'Facebook Bundle', true, 'Unlimited Facebook access'),
('Banglalink', 'social_media', 'YouTube Package', true, '20-50 hours video streaming'),
('Banglalink', 'social_media', 'Combo Packages', true, 'Multiple platforms bundled together'),

-- Voice Features
('Banglalink', 'voice', 'Local Calls', true, '0.50-1.50 per minute'),
('Banglalink', 'voice', 'On-Net Calls', true, '0.50 per minute on same network'),
('Banglalink', 'voice', 'International Calling', true, 'Competitive rates to popular destinations'),
('Banglalink', 'voice', 'Friends & Family', true, 'Reduced rates to 5-10 frequent numbers'),

-- Digital Services
('Banglalink', 'digital', 'Banglalink Money', true, 'Mobile wallet for digital transactions'),
('Banglalink', 'digital', 'Money Transfer', true, 'Send money to other Banglalink users'),
('Banglalink', 'digital', 'Bill Payments', true, 'Pay electricity, water, and utility bills'),
('Banglalink', 'digital', 'Merchant Payments', true, 'Fast business payment system'),

-- Support Features
('Banglalink', 'support', '24/7 Hotline', true, 'Call 111 for customer support'),
('Banglalink', 'support', 'Mobile App', true, 'User-friendly app for account management'),
('Banglalink', 'support', 'USSD Codes', true, 'Quick access via *123# and other codes'),
('Banglalink', 'support', 'Social Media Support', true, 'Support via Facebook and Twitter'),

-- Grameenphone Features
('Grameenphone', 'data', 'Daily Bundles', true, '100MB-500MB with 4 package options'),
('Grameenphone', 'data', 'Weekly Bundles', true, '1GB-3GB with 4 package options'),
('Grameenphone', 'data', 'Monthly Unlimited', true, 'Most comprehensive unlimited plans'),
('Grameenphone', 'digital', 'bKash Integration', true, 'Integrated with largest mobile banking platform'),

-- Robi Features
('Robi', 'data', 'Daily Bundles', true, '100MB-500MB packages'),
('Robi', 'data', 'Weekly Bundles', true, '1GB-3GB with variety'),
('Robi', 'digital', 'Rocket Integration', true, 'Connected to Rocket payment system'),

-- Airtel Features
('Airtel', 'data', 'Daily Bundles', true, '100MB-300MB'),
('Airtel', 'data', 'Competitive Rates', true, 'Growing network expansion'),

-- Teletalk Features
('Teletalk', 'data', 'Budget Packages', true, 'Cheapest data rates in market'),
('Teletalk', 'voice', 'Educational Discount', true, 'Special rates for students');

-- ============================================================================
-- 8. CREATE PROS & CONS TABLE
-- ============================================================================

CREATE TABLE IF NOT EXISTS operator_pros_cons (
    id SERIAL PRIMARY KEY,
    operator_id INT NOT NULL REFERENCES mobile_operator_comparisons(id),
    operator_name VARCHAR(255),
    type VARCHAR(10), -- 'pro' or 'con'
    item_text VARCHAR(500),
    priority INT DEFAULT 1, -- 1-highest, 3-lowest
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================================
-- 9. INSERT PROS & CONS DATA
-- ============================================================================

-- Banglalink Pros
INSERT INTO operator_pros_cons (operator_name, type, item_text, priority)
VALUES
('Banglalink', 'pro', 'Affordable call rates - Competitive pricing for local and international calls', 1),
('Banglalink', 'pro', 'Attractive data bundles - Regular promotions with bonus offers', 1),
('Banglalink', 'pro', 'Good 4G coverage in major cities', 1),
('Banglalink', 'pro', 'User-friendly mobile app for account management', 2),
('Banglalink', 'pro', 'Digital payment support through mobile banking', 2),
('Banglalink', 'pro', 'Emergency balance feature for urgent calls', 2),
('Banglalink', 'pro', 'Flexible prepaid and postpaid options', 2),
('Banglalink', 'pro', 'Frequent promotions and special offers', 2),
('Banglalink', 'pro', 'Social media bundles for popular platforms', 3),
('Banglalink', 'pro', 'Integrated money services (Banglalink Money)', 3),
('Banglalink', 'pro', 'Multiple customer support channels available', 3);

-- Banglalink Cons
INSERT INTO operator_pros_cons (operator_name, type, item_text, priority)
VALUES
('Banglalink', 'con', 'Network congestion during peak hours in urban areas', 1),
('Banglalink', 'con', 'Limited 4G coverage in remote and rural areas', 1),
('Banglalink', 'con', 'Long customer service wait times during peak periods', 2),
('Banglalink', 'con', 'Higher roaming rates compared to some competitors', 2),
('Banglalink', 'con', 'Occasional data speed throttling on certain plans', 2),
('Banglalink', 'con', 'SIM activation process could be simplified', 3),
('Banglalink', 'con', 'Some packages have restrictive validity periods', 3);

-- Grameenphone Pros
INSERT INTO operator_pros_cons (operator_name, type, item_text, priority)
VALUES
('Grameenphone', 'pro', 'Best network quality and reliability in Bangladesh', 1),
('Grameenphone', 'pro', 'Widest 4G coverage including rural areas', 1),
('Grameenphone', 'pro', 'Fastest data speeds averaging 15-20 Mbps', 1),
('Grameenphone', 'pro', 'Most comprehensive package variety', 1),
('Grameenphone', 'pro', 'Excellent customer support infrastructure', 1),
('Grameenphone', 'pro', 'Strong financial ecosystem with bKash', 2),
('Grameenphone', 'pro', 'Most frequent promotional offers', 2);

-- Grameenphone Cons
INSERT INTO operator_pros_cons (operator_name, type, item_text, priority)
VALUES
('Grameenphone', 'con', 'Premium pricing - Most expensive among operators', 1),
('Grameenphone', 'con', 'Not ideal for budget-conscious users', 1),
('Grameenphone', 'con', 'Sometimes over-promotional messaging', 2);

-- Robi Pros
INSERT INTO operator_pros_cons (operator_name, type, item_text, priority)
VALUES
('Robi', 'pro', 'Good network quality and reliability', 1),
('Robi', 'pro', 'Decent 4G coverage in urban and semi-urban areas', 1),
('Robi', 'pro', 'Affordable pricing with good package variety', 1),
('Robi', 'pro', 'Reliable service with 98% uptime', 2),
('Robi', 'pro', 'Rocket payment system integration', 2);

-- Robi Cons
INSERT INTO operator_pros_cons (operator_name, type, item_text, priority)
VALUES
('Robi', 'con', 'Middle-ground positioning - not best in any category', 1),
('Robi', 'con', 'Stronger competitors on both premium and budget sides', 1),
('Robi', 'con', 'Less aggressive in technological innovation', 2);

-- Airtel Pros
INSERT INTO operator_pros_cons (operator_name, type, item_text, priority)
VALUES
('Airtel', 'pro', 'Competitive pricing structure', 1),
('Airtel', 'pro', 'Growing network infrastructure', 1),
('Airtel', 'pro', 'Good customer support in major cities', 2),
('Airtel', 'pro', 'Improving service quality', 2);

-- Airtel Cons
INSERT INTO operator_pros_cons (operator_name, type, item_text, priority)
VALUES
('Airtel', 'con', 'Smaller subscriber base than competitors', 1),
('Airtel', 'con', 'Limited coverage in rural areas', 1),
('Airtel', 'con', 'Less comprehensive service offerings', 2),
('Airtel', 'con', 'Fewer corporate partnerships', 2);

-- Teletalk Pros
INSERT INTO operator_pros_cons (operator_name, type, item_text, priority)
VALUES
('Teletalk', 'pro', 'Cheapest rates overall in the market', 1),
('Teletalk', 'pro', 'Government backing and stability', 2),
('Teletalk', 'pro', 'Educational institution focus', 2),
('Teletalk', 'pro', 'Budget-friendly options', 2);

-- Teletalk Cons
INSERT INTO operator_pros_cons (operator_name, type, item_text, priority)
VALUES
('Teletalk', 'con', 'Oldest technology with limited 4G infrastructure', 1),
('Teletalk', 'con', 'Poor customer service quality', 1),
('Teletalk', 'con', 'Limited coverage especially in urban areas', 1),
('Teletalk', 'con', 'Minimal innovation and service expansion', 2),
('Teletalk', 'con', 'Limited service ecosystem and partnerships', 2);

-- ============================================================================
-- 10. CREATE QUERY VIEWS FOR EASY ACCESS
-- ============================================================================

-- View: Operator Comparison Summary
CREATE OR REPLACE VIEW operator_comparison_summary AS
SELECT 
    operator_name,
    market_rank,
    subscribers_millions,
    overall_rating,
    recommendation_score,
    recommendation_text,
    hotline_number,
    website_url,
    created_at
FROM mobile_operator_comparisons
ORDER BY market_rank;

-- View: Price Comparison by Service Type
CREATE OR REPLACE VIEW price_comparison_by_service AS
SELECT 
    operator_name,
    service_type,
    package_name,
    volume,
    duration,
    price_bdt,
    ROW_NUMBER() OVER (PARTITION BY service_type ORDER BY price_bdt) as price_rank
FROM operator_pricing_comparison
ORDER BY service_type, price_bdt;

-- View: Feature Availability Matrix
CREATE OR REPLACE VIEW feature_availability_matrix AS
SELECT 
    feature_name,
    SUM(CASE WHEN operator_name = 'Banglalink' THEN 1 ELSE 0 END) as banglalink,
    SUM(CASE WHEN operator_name = 'Grameenphone' THEN 1 ELSE 0 END) as grameenphone,
    SUM(CASE WHEN operator_name = 'Robi' THEN 1 ELSE 0 END) as robi,
    SUM(CASE WHEN operator_name = 'Airtel' THEN 1 ELSE 0 END) as airtel,
    SUM(CASE WHEN operator_name = 'Teletalk' THEN 1 ELSE 0 END) as teletalk
FROM operator_features
WHERE is_available = true
GROUP BY feature_name
ORDER BY feature_name;

-- View: Best Operators for Each Use Case
CREATE OR REPLACE VIEW best_operators_by_usecase AS
SELECT 
    'Students' as usecase,
    CASE 
        WHEN best_for_students = true THEN operator_name 
    END as recommended_operator,
    overall_rating
FROM mobile_operator_comparisons
WHERE best_for_students = true
UNION ALL
SELECT 
    'Professionals' as usecase,
    CASE 
        WHEN best_for_professionals = true THEN operator_name 
    END as recommended_operator,
    overall_rating
FROM mobile_operator_comparisons
WHERE best_for_professionals = true
UNION ALL
SELECT 
    'Budget Users' as usecase,
    CASE 
        WHEN best_for_budget_users = true THEN operator_name 
    END as recommended_operator,
    overall_rating
FROM mobile_operator_comparisons
WHERE best_for_budget_users = true
UNION ALL
SELECT 
    'Heavy Users' as usecase,
    CASE 
        WHEN best_for_heavy_users = true THEN operator_name 
    END as recommended_operator,
    overall_rating
FROM mobile_operator_comparisons
WHERE best_for_heavy_users = true
UNION ALL
SELECT 
    'Business' as usecase,
    CASE 
        WHEN best_for_business = true THEN operator_name 
    END as recommended_operator,
    overall_rating
FROM mobile_operator_comparisons
WHERE best_for_business = true;

-- ============================================================================
-- 11. VERIFICATION QUERIES
-- ============================================================================

-- Verify all operators inserted
SELECT 'Operators Inserted' as check_name, COUNT(*) as count FROM mobile_operator_comparisons;

-- Verify pricing data
SELECT 'Pricing Records' as check_name, COUNT(*) as count FROM operator_pricing_comparison;

-- Verify features
SELECT 'Features' as check_name, COUNT(*) as count FROM operator_features;

-- Verify pros/cons
SELECT 'Pros/Cons' as check_name, COUNT(*) as count FROM operator_pros_cons;

-- Display operator ratings comparison
SELECT 
    operator_name,
    market_rank,
    overall_rating,
    network_quality_rating,
    pricing_rating,
    value_for_money_rating,
    recommendation_score
FROM mobile_operator_comparisons
ORDER BY market_rank;

-- Display best prices by category
SELECT 
    service_type,
    MIN(price_bdt) as cheapest_price_bdt,
    MAX(price_bdt) as most_expensive_price_bdt,
    AVG(price_bdt)::INT as average_price_bdt
FROM operator_pricing_comparison
GROUP BY service_type
ORDER BY service_type;

-- ============================================================================
-- END OF INSERTION SCRIPT
-- ============================================================================
-- Total Records Inserted:
-- - Operators: 5
-- - Pricing Comparisons: 30+
-- - Features: 50+
-- - Pros/Cons: 40+
-- - Database Views: 4
-- ============================================================================
