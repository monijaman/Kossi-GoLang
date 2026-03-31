# Banglalink Database Integration Guide

## Quick Start

### Option 1: Direct PostgreSQL Execution (Recommended)

```bash
# Navigate to server directory
cd d:/GO/gocrit/gocrit_server

# Execute the insertion script
psql -U postgres -h localhost -d gocrit < SQL_INSERT_BANGLALINK_COMPARISON.sql

# OR if using a specific database connection
psql postgresql://user:password@localhost:5432/gocrit < SQL_INSERT_BANGLALINK_COMPARISON.sql
```

### Option 2: Using pgAdmin

1. Open pgAdmin
2. Connect to your PostgreSQL database
3. Open Query Tool
4. Copy and paste contents of `SQL_INSERT_BANGLALINK_COMPARISON.sql`
5. Execute (F5 or Execute button)

### Option 3: Using DBeaver

1. Open DBeaver
2. Connect to your PostgreSQL database
3. File → Open SQL Script → `SQL_INSERT_BANGLALINK_COMPARISON.sql`
4. Execute (Ctrl+Enter or Run All)

---

## Database Schema Created

### 1. `mobile_operator_comparisons` Table
**Purpose:** Store operator overview and ratings

**Columns:**
```
- id: PRIMARY KEY
- operator_name: Unique operator name
- operator_slug: URL-friendly slug
- market_rank: Position in market
- subscribers_millions: User base size
- overall_rating: Composite rating (0-5)
- network_quality_rating: Technical quality (0-5)
- coverage_rating: Geographic reach (0-5)
- data_speed_rating: Data speed performance (0-5)
- pricing_rating: Price competitiveness (0-5)
- package_variety_rating: Plan options (0-5)
- customer_support_rating: Support quality (0-5)
- value_for_money_rating: Overall value (0-5)
- avg_data_speed_mbps: Average download speed
- peak_hour_speed_mbps: Speed during peak hours
- network_uptime_percent: Reliability percentage
- latency_ms: Typical latency in milliseconds
- urban_coverage_percent: City coverage
- rural_coverage_percent: Village coverage
- hotline_number: Customer service number
- website_url: Official website
- has_4g, has_5g_plans: Technology availability
- mobile_wallet_name: Digital wallet name
- has_money_transfer, has_bill_payments, has_merchant_payments: Services
- best_for_*: Boolean flags for use cases
- recommendation_score: 1-10 score
- recommendation_text: Reason for recommendation
```

**Operators Inserted:**
1. Banglalink (ID: 2)
2. Grameenphone (ID: 1)
3. Robi (ID: 3)
4. Airtel (ID: 4)
5. Teletalk (ID: 5)

---

### 2. `operator_pricing_comparison` Table
**Purpose:** Store detailed pricing for each operator and service

**Record Count:** 30+ pricing entries

**Service Types:**
- `data`: Mobile data plans
- `voice`: Call/minute packages
- `sms`: Text message packages
- `roaming`: International packages

**Sample Data Structure:**
```
Operator: Banglalink
Service: data
Package: Daily Bundle 1
Volume: 100MB
Duration: 1 Day
Price: 50 BDT

Package: Weekly Bundle 1
Volume: 1GB
Duration: 7 Days
Price: 120 BDT
```

---

### 3. `operator_features` Table
**Purpose:** Track features and capabilities

**Record Count:** 50+ features

**Feature Categories:**
- `data`: Data packages and speed features
- `social_media`: Social platform bundles
- `voice`: Calling features
- `digital`: Digital services
- `support`: Customer support options

**Examples:**
```
- Banglalink: WhatsApp Bundle (Unlimited WhatsApp)
- Grameenphone: bKash Integration
- Robi: Rocket Payment System
- Airtel: Growing network expansion
- Teletalk: Educational Discount
```

---

### 4. `operator_pros_cons` Table
**Purpose:** Structured pros and cons for decision-making

**Record Count:** 40+ entries

**Fields:**
```
- operator_name: Which operator
- type: 'pro' or 'con'
- item_text: Description
- priority: 1 (highest), 2 (medium), 3 (lowest)
```

**Banglalink Data:**
```
PROS (11 items):
- Affordable call rates (Priority 1)
- Attractive data bundles (Priority 1)
- Good 4G coverage (Priority 1)
- User-friendly app (Priority 2)
- Digital payment support (Priority 2)
... (6 more)

CONS (7 items):
- Network congestion during peak hours (Priority 1)
- Limited rural coverage (Priority 1)
- Long customer service wait times (Priority 2)
... (4 more)
```

---

## Views Created for Easy Access

### 1. `operator_comparison_summary`
Quick overview of all operators with key metrics

**Query Example:**
```sql
SELECT * FROM operator_comparison_summary ORDER BY market_rank;
```

---

### 2. `price_comparison_by_service`
Compare prices across operators for each service type

**Query Example:**
```sql
SELECT * FROM price_comparison_by_service WHERE service_type = 'data';
```

---

### 3. `feature_availability_matrix`
See which features each operator offers

**Query Example:**
```sql
SELECT * FROM feature_availability_matrix;
```

---

### 4. `best_operators_by_usecase`
Find best operator for specific use cases

**Query Example:**
```sql
SELECT * FROM best_operators_by_usecase WHERE usecase = 'Students';
```

---

## Using Reporting Queries

Execute the `SQL_BANGLALINK_REPORTS.sql` file for ready-made reports:

### Market Reports
```sql
-- Report 1: Market Ranking
SELECT * FROM mobile_operator_comparisons ORDER BY market_rank;

-- Report 2: Rating Comparison
-- Shows all ratings in table format
```

### Pricing Reports
```sql
-- Report 5: Cheapest Options per Service
-- Shows lowest price for each service type

-- Report 6: Data Package Price Comparison
-- Calculates price per day for data packages
```

### Feature Reports
```sql
-- Report 9: Feature Availability Summary
-- Count of features per operator

-- Report 10: Feature Comparison Matrix
-- ✓/✗ matrix of all features
```

### Decision Reports
```sql
-- Report 14: Best Operator by Use Case
-- Recommendations for different scenarios

-- Report 15: Quality vs Price Analysis
-- Value assessment for each operator
```

---

## Data Integration with Go Application

### Option 1: Direct Query in Code

```go
// Get all operators
var operators []MobileOperatorComparison
db.Find(&operators)

// Get best operator by use case
var student_best MobileOperatorComparison
db.Where("best_for_students = ?", true).First(&student_best)

// Get pricing for specific operator
var prices []OperatorPricingComparison
db.Where("operator_name = ? AND service_type = ?", "Banglalink", "data").Find(&prices)
```

### Option 2: GORM Models

Create corresponding GORM models:

```go
type MobileOperatorComparison struct {
    ID                      uint
    OperatorName            string
    OperatorSlug            string
    MarketRank              int
    SubscribersMilions      int
    OverallRating           float32
    NetworkQualityRating    float32
    CoverageRating          float32
    DataSpeedRating         float32
    PricingRating           float32
    // ... other fields
    CreatedAt               time.Time
    UpdatedAt               time.Time
}

type OperatorPricingComparison struct {
    ID            uint
    OperatorID    uint
    OperatorName  string
    ServiceType   string
    PackageName   string
    Volume        string
    Duration      string
    PriceBdt      int
    Currency      string
    Notes         string
    CreatedAt     time.Time
}

type OperatorFeatures struct {
    ID              uint
    OperatorID      uint
    OperatorName    string
    FeatureCategory string
    FeatureName     string
    IsAvailable     bool
    Details         string
    CreatedAt       time.Time
}

type OperatorProsCons struct {
    ID            uint
    OperatorID    uint
    OperatorName  string
    Type          string // 'pro' or 'con'
    ItemText      string
    Priority      int
    CreatedAt     time.Time
}
```

### Option 3: Create API Endpoints

```go
// GET /api/operators - List all operators
// GET /api/operators/:id - Get specific operator
// GET /api/operators/:id/pricing - Get pricing
// GET /api/operators/:id/features - Get features
// GET /api/operators/:id/pros-cons - Get pros and cons
// GET /api/comparison - Compare multiple operators
// GET /api/best-for/:usecase - Get recommendation
```

---

## Verification Steps

After insertion, verify the data:

```bash
# Check if tables were created
psql -U postgres -d gocrit -c "\dt"

# Count records in each table
psql -U postgres -d gocrit -c "
  SELECT 'Operators' as table_name, COUNT(*) FROM mobile_operator_comparisons
  UNION ALL
  SELECT 'Pricing', COUNT(*) FROM operator_pricing_comparison
  UNION ALL
  SELECT 'Features', COUNT(*) FROM operator_features
  UNION ALL
  SELECT 'Pros/Cons', COUNT(*) FROM operator_pros_cons;
"

# View sample data
psql -U postgres -d gocrit -c "SELECT operator_name, overall_rating, recommendation_score FROM mobile_operator_comparisons ORDER BY market_rank;"
```

---

## Expected Results

**Tables Created:** 4
- mobile_operator_comparisons
- operator_pricing_comparison
- operator_features
- operator_pros_cons

**Views Created:** 4
- operator_comparison_summary
- price_comparison_by_service
- feature_availability_matrix
- best_operators_by_usecase

**Records Inserted:**
- Operators: 5
- Pricing Comparisons: 30+
- Features: 50+
- Pros/Cons: 40+

**Total Rows Added:** 130+

---

## Common Queries for Frontend

### Get All Operators with Ratings
```sql
SELECT operator_name, market_rank, overall_rating, subscribers_millions, hotline_number, website_url
FROM mobile_operator_comparisons
ORDER BY market_rank;
```

### Get Banglalink Details
```sql
SELECT * FROM mobile_operator_comparisons WHERE operator_slug = 'banglalink';
```

### Get Banglalink Pricing
```sql
SELECT service_type, package_name, volume, duration, price_bdt
FROM operator_pricing_comparison
WHERE operator_name = 'Banglalink'
ORDER BY service_type, price_bdt;
```

### Get Banglalink Features
```sql
SELECT feature_category, feature_name, details
FROM operator_features
WHERE operator_name = 'Banglalink' AND is_available = true
ORDER BY feature_category, feature_name;
```

### Get Banglalink Pros and Cons
```sql
SELECT type, item_text, priority
FROM operator_pros_cons
WHERE operator_name = 'Banglalink'
ORDER BY type DESC, priority;
```

### Compare Two Operators
```sql
SELECT 
    service_type,
    'Banglalink' as operator1,
    'Grameenphone' as operator2,
    (SELECT MIN(price_bdt) FROM operator_pricing_comparison WHERE operator_name = 'Banglalink' AND service_type = operator_pricing_comparison.service_type) as banglalink_price,
    (SELECT MIN(price_bdt) FROM operator_pricing_comparison WHERE operator_name = 'Grameenphone' AND service_type = operator_pricing_comparison.service_type) as grameenphone_price
FROM operator_pricing_comparison
GROUP BY service_type;
```

---

## Troubleshooting

### Issue: "Extension uuid-ossp does not exist"
**Solution:**
```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

### Issue: "Permission denied"
**Solution:** Use superuser or account with CREATE TABLE privileges:
```bash
psql -U postgres (instead of current user)
```

### Issue: Table already exists
**Solution:** Drop and recreate or use `IF NOT EXISTS` clauses (already included in script)

### Issue: Foreign key constraint errors
**Solution:** Ensure parent tables are created before child tables. The script has correct order.

---

## Next Steps

1. **Execute SQL scripts** to create tables and insert data
2. **Create GORM models** in your Go application
3. **Build API endpoints** to query the data
4. **Create frontend views** using the views provided
5. **Add filtering and sorting** capabilities
6. **Implement caching** for frequently accessed queries

---

## File References

- **Insertion Script:** `SQL_INSERT_BANGLALINK_COMPARISON.sql`
- **Reporting Queries:** `SQL_BANGLALINK_REPORTS.sql`
- **Documentation:** This file

---

## Support

For issues or additional queries needed, refer to:
- BANGLALINK_MASTER_INDEX.md (documentation structure)
- BANGLALINK_COMPETITOR_COMPARISON.md (detailed comparison logic)
- BANGLALINK_COMPARISON_QUICK_TABLES.md (data reference)

