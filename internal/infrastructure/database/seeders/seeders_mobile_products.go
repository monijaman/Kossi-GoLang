package seeders

import (
	"fmt"
	"kossti/internal/infrastructure/database/models"
	"log"
	"time"

	"gorm.io/gorm"
)

// SeedMobileProducts seeds mobile phone products for category 79
func SeedMobileProducts(db *gorm.DB) error {
	fmt.Println("\n🔄 Seeding Mobile Products (Category 79)...")

	catID := uint(79)
	now := time.Now()

	// Mobile products data based on mobile.json
	products := []models.ProductModel{
		// Apple iPhones
		{Name: "iPhone 17 Pro Max", Description: ptrStr("Latest flagship Apple smartphone with advanced camera system"), Slug: "iphone-17-pro-max", StartPrice: ptrFloat(189999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 1, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 17 Pro", Description: ptrStr("Premium Apple flagship smartphone"), Slug: "iphone-17-pro", StartPrice: ptrFloat(149999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 2, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 17 Plus", Description: ptrStr("Large-screen Apple flagship"), Slug: "iphone-17-plus", StartPrice: ptrFloat(129999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 3, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 17", Description: ptrStr("Standard Apple flagship smartphone"), Slug: "iphone-17", StartPrice: ptrFloat(109999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 4, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 16", Description: ptrStr("Previous generation Apple flagship"), Slug: "iphone-16", Price: ptrFloat(89999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 5, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 15", Description: ptrStr("Latest generation standard iPhone"), Slug: "iphone-15", Price: ptrFloat(69999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 6, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone SE (3rd Gen)", Description: ptrStr("Affordable Apple smartphone"), Slug: "iphone-se-3rd-gen", Price: ptrFloat(44999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 7, CreatedAt: now, UpdatedAt: now},

		// Samsung Galaxy S Series
		{Name: "Samsung Galaxy S25 Ultra", Description: ptrStr("Premium Samsung flagship with S Pen"), Slug: "samsung-galaxy-s25-ultra", Price: ptrFloat(149999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 1, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S25+", Description: ptrStr("Large-screen Samsung flagship"), Slug: "samsung-galaxy-s25-plus", Price: ptrFloat(129999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 2, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S25", Description: ptrStr("Standard Samsung flagship"), Slug: "samsung-galaxy-s25", Price: ptrFloat(109999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 3, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S24 Ultra", Description: ptrStr("Previous generation Galaxy flagship"), Slug: "samsung-galaxy-s24-ultra", Price: ptrFloat(119999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 4, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S24+", Description: ptrStr("Previous generation Plus variant"), Slug: "samsung-galaxy-s24-plus", Price: ptrFloat(99999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 5, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S24", Description: ptrStr("Previous generation standard"), Slug: "samsung-galaxy-s24", Price: ptrFloat(84999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 6, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S24 FE", Description: ptrStr("Fan Edition Galaxy S24"), Slug: "samsung-galaxy-s24-fe", Price: ptrFloat(64999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 7, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S23 Ultra", Description: ptrStr("Older generation ultra"), Slug: "samsung-galaxy-s23-ultra", Price: ptrFloat(99999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 8, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S23+", Description: ptrStr("Older generation plus"), Slug: "samsung-galaxy-s23-plus", Price: ptrFloat(79999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 9, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S23", Description: ptrStr("Older generation standard"), Slug: "samsung-galaxy-s23", Price: ptrFloat(69999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 10, CreatedAt: now, UpdatedAt: now},

		// Samsung Galaxy Z Series (Foldables)
		{Name: "Samsung Galaxy Z Fold 7", Description: ptrStr("Premium foldable smartphone"), Slug: "samsung-galaxy-z-fold-7", Price: ptrFloat(189999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 11, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy Z Flip 7", Description: ptrStr("Compact foldable smartphone"), Slug: "samsung-galaxy-z-flip-7", Price: ptrFloat(129999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 12, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy Z Flip 7 FE", Description: ptrStr("Fan Edition foldable"), Slug: "samsung-galaxy-z-flip-7-fe", Price: ptrFloat(99999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 13, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy Z Fold 6", Description: ptrStr("Previous generation foldable"), Slug: "samsung-galaxy-z-fold-6", Price: ptrFloat(159999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 14, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy Z Flip 6", Description: ptrStr("Previous generation flip"), Slug: "samsung-galaxy-z-flip-6", Price: ptrFloat(109999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 15, CreatedAt: now, UpdatedAt: now},

		// Samsung Galaxy A Series (Mid-range)
		{Name: "Samsung Galaxy A56 5G", Description: ptrStr("Mid-range 5G smartphone"), Slug: "samsung-galaxy-a56-5g", Price: ptrFloat(44999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 16, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A55 5G", Description: ptrStr("Popular mid-range 5G"), Slug: "samsung-galaxy-a55-5g", Price: ptrFloat(39999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 17, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A54 5G", Description: ptrStr("Budget-friendly 5G"), Slug: "samsung-galaxy-a54-5g", Price: ptrFloat(34999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 18, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A25 5G", Description: ptrStr("Entry-level 5G"), Slug: "samsung-galaxy-a25-5g", Price: ptrFloat(26999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 19, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A15 5G", Description: ptrStr("Budget 5G smartphone"), Slug: "samsung-galaxy-a15-5g", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 20, CreatedAt: now, UpdatedAt: now},

		// Samsung Galaxy M Series
		{Name: "Samsung Galaxy M56 5G", Description: ptrStr("Mid-range M series 5G"), Slug: "samsung-galaxy-m56-5g", Price: ptrFloat(36999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 21, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy M36 5G", Description: ptrStr("Affordable M series"), Slug: "samsung-galaxy-m36-5g", Price: ptrFloat(27999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 22, CreatedAt: now, UpdatedAt: now},

		// Samsung Galaxy F Series
		{Name: "Samsung Galaxy F56 5G", Description: ptrStr("Premium F series phone"), Slug: "samsung-galaxy-f56-5g", Price: ptrFloat(42999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 23, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy F36 5G", Description: ptrStr("Budget F series 5G"), Slug: "samsung-galaxy-f36-5g", Price: ptrFloat(24999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 24, CreatedAt: now, UpdatedAt: now},

		// Xiaomi Flagship Series
		{Name: "Xiaomi 15 Ultra", Description: ptrStr("Premium Xiaomi flagship"), Slug: "xiaomi-15-ultra", Price: ptrFloat(99999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 25, CreatedAt: now, UpdatedAt: now},
		{Name: "Xiaomi 15", Description: ptrStr("Latest Xiaomi flagship"), Slug: "xiaomi-15", Price: ptrFloat(79999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 26, CreatedAt: now, UpdatedAt: now},
		{Name: "Xiaomi 15 Pro", Description: ptrStr("Pro variant Xiaomi 15"), Slug: "xiaomi-15-pro", Price: ptrFloat(89999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 27, CreatedAt: now, UpdatedAt: now},
		{Name: "Xiaomi 14 Ultra", Description: ptrStr("Previous generation ultra"), Slug: "xiaomi-14-ultra", Price: ptrFloat(89999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 28, CreatedAt: now, UpdatedAt: now},
		{Name: "Xiaomi 14 Pro", Description: ptrStr("Previous generation pro"), Slug: "xiaomi-14-pro", Price: ptrFloat(74999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 29, CreatedAt: now, UpdatedAt: now},

		// Redmi Note Series
		{Name: "Redmi Note 15", Description: ptrStr("Latest Redmi Note"), Slug: "redmi-note-15", Price: ptrFloat(39999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 30, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi Note 15 Pro 5G", Description: ptrStr("Pro variant Note 15"), Slug: "redmi-note-15-pro-5g", Price: ptrFloat(46999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 31, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi Note 14", Description: ptrStr("Mid-range Redmi Note"), Slug: "redmi-note-14", Price: ptrFloat(34999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 32, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi Note 14 Pro+ 5G", Description: ptrStr("Powerful Redmi Note"), Slug: "redmi-note-14-pro-plus-5g", Price: ptrFloat(42999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 33, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi Note 13", Description: ptrStr("Older Note generation"), Slug: "redmi-note-13", Price: ptrFloat(29999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 34, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi Note 12", Description: ptrStr("Older Note generation"), Slug: "redmi-note-12", Price: ptrFloat(24999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 35, CreatedAt: now, UpdatedAt: now},

		// Redmi Numbered Series
		{Name: "Redmi 15", Description: ptrStr("Latest Redmi numbered"), Slug: "redmi-15", Price: ptrFloat(22999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 36, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi 14", Description: ptrStr("Redmi numbered series"), Slug: "redmi-14", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 37, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi 13C", Description: ptrStr("Budget Redmi series"), Slug: "redmi-13c", Price: ptrFloat(16999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 38, CreatedAt: now, UpdatedAt: now},

		// POCO F Series
		{Name: "POCO F7", Description: ptrStr("Gaming POCO phone"), Slug: "poco-f7", Price: ptrFloat(54999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 39, CreatedAt: now, UpdatedAt: now},
		{Name: "POCO F7 Ultra", Description: ptrStr("Premium POCO gaming"), Slug: "poco-f7-ultra", Price: ptrFloat(64999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 40, CreatedAt: now, UpdatedAt: now},
		{Name: "POCO F6 5G", Description: ptrStr("Previous POCO flagship"), Slug: "poco-f6-5g", Price: ptrFloat(44999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 41, CreatedAt: now, UpdatedAt: now},

		// POCO X Series
		{Name: "POCO X7 Pro 5G", Description: ptrStr("Premium POCO X series"), Slug: "poco-x7-pro-5g", Price: ptrFloat(42999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 42, CreatedAt: now, UpdatedAt: now},
		{Name: "POCO X7 5G", Description: ptrStr("POCO X series 5G"), Slug: "poco-x7-5g", Price: ptrFloat(36999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 43, CreatedAt: now, UpdatedAt: now},
		{Name: "POCO X6 Pro 5G", Description: ptrStr("Older POCO X pro"), Slug: "poco-x6-pro-5g", Price: ptrFloat(32999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 44, CreatedAt: now, UpdatedAt: now},

		// POCO M Series
		{Name: "POCO M7 Pro 5G", Description: ptrStr("Budget-friendly POCO"), Slug: "poco-m7-pro-5g", Price: ptrFloat(24999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 45, CreatedAt: now, UpdatedAt: now},
		{Name: "POCO M6 Pro 5G", Description: ptrStr("Previous POCO M pro"), Slug: "poco-m6-pro-5g", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 46, CreatedAt: now, UpdatedAt: now},

		// Vivo X Series
		{Name: "Vivo X300 Pro", Description: ptrStr("Premium Vivo flagship"), Slug: "vivo-x300-pro", Price: ptrFloat(99999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 47, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo X300", Description: ptrStr("Latest Vivo X series"), Slug: "vivo-x300", Price: ptrFloat(84999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 48, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo X200 Pro", Description: ptrStr("Older X generation pro"), Slug: "vivo-x200-pro", Price: ptrFloat(74999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 49, CreatedAt: now, UpdatedAt: now},

		// Vivo V Series
		{Name: "Vivo V60 5G", Description: ptrStr("Latest Vivo V series"), Slug: "vivo-v60-5g", Price: ptrFloat(49999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 50, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo V60 Lite 5G", Description: ptrStr("Budget V60 variant"), Slug: "vivo-v60-lite-5g", Price: ptrFloat(39999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 51, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo V50", Description: ptrStr("Mid-range Vivo phone"), Slug: "vivo-v50", Price: ptrFloat(34999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 52, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo V40", Description: ptrStr("Older Vivo V series"), Slug: "vivo-v40", Price: ptrFloat(29999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 53, CreatedAt: now, UpdatedAt: now},

		// Vivo Y Series (Budget)
		{Name: "Vivo Y29", Description: ptrStr("Budget Vivo phone"), Slug: "vivo-y29", Price: ptrFloat(17999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 54, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo Y28", Description: ptrStr("Entry Vivo smartphone"), Slug: "vivo-y28", Price: ptrFloat(15999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 55, CreatedAt: now, UpdatedAt: now},

		// Oppo Find Series (Foldables)
		{Name: "Oppo Find N4 Fold", Description: ptrStr("Premium Oppo foldable"), Slug: "oppo-find-n4-fold", Price: ptrFloat(179999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 56, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo Find N4 Flip", Description: ptrStr("Oppo flip foldable"), Slug: "oppo-find-n4-flip", Price: ptrFloat(129999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 57, CreatedAt: now, UpdatedAt: now},

		// Oppo Reno Series
		{Name: "Oppo Reno 14 5G", Description: ptrStr("Latest Oppo Reno"), Slug: "oppo-reno-14-5g", Price: ptrFloat(69999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 58, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo Reno 14 F 5G", Description: ptrStr("Budget Reno 14"), Slug: "oppo-reno-14-f-5g", Price: ptrFloat(54999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 59, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo Reno 13 5G", Description: ptrStr("Older Reno generation"), Slug: "oppo-reno-13-5g", Price: ptrFloat(59999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 60, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo Reno 13 Pro 5G", Description: ptrStr("Pro Reno variant"), Slug: "oppo-reno-13-pro-5g", Price: ptrFloat(64999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 61, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo Reno 15", Description: ptrStr("Latest Reno numbered"), Slug: "oppo-reno-15", Price: ptrFloat(84999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 62, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo Reno 15 Pro", Description: ptrStr("Premium Reno 15"), Slug: "oppo-reno-15-pro", Price: ptrFloat(99999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 63, CreatedAt: now, UpdatedAt: now},

		// Oppo A Series
		{Name: "Oppo A5 Pro", Description: ptrStr("Premium Oppo A series"), Slug: "oppo-a5-pro", Price: ptrFloat(44999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 64, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo A5", Description: ptrStr("Latest A series"), Slug: "oppo-a5", Price: ptrFloat(34999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 65, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo A60", Description: ptrStr("Budget Oppo phone"), Slug: "oppo-a60", Price: ptrFloat(24999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 66, CreatedAt: now, UpdatedAt: now},

		// Realme Numbered Series
		{Name: "Realme 15 Pro 5G", Description: ptrStr("Performance Realme phone"), Slug: "realme-15-pro-5g", Price: ptrFloat(44999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 67, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme 15 5G", Description: ptrStr("Latest Realme 15"), Slug: "realme-15-5g", Price: ptrFloat(34999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 68, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme 14 5G", Description: ptrStr("Previous Realme 14"), Slug: "realme-14-5g", Price: ptrFloat(29999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 69, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme 13", Description: ptrStr("Mid-range Realme"), Slug: "realme-13", Price: ptrFloat(24999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 70, CreatedAt: now, UpdatedAt: now},

		// Realme C Series (Budget)
		{Name: "Realme C75x", Description: ptrStr("Budget-friendly Realme"), Slug: "realme-c75x", Price: ptrFloat(16999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 71, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme C71", Description: ptrStr("Entry-level Realme"), Slug: "realme-c71", Price: ptrFloat(14999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 72, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme C65", Description: ptrStr("Older C series"), Slug: "realme-c65", Price: ptrFloat(13999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 73, CreatedAt: now, UpdatedAt: now},

		// Realme GT Series
		{Name: "Realme GT Neo 3", Description: ptrStr("Gaming Realme phone"), Slug: "realme-gt-neo-3", Price: ptrFloat(49999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 74, CreatedAt: now, UpdatedAt: now},

		// Apple iPhone Series
		{Name: "iPhone 17 Pro Max", Description: ptrStr("Premium Apple flagship"), Slug: "iphone-17-pro-max", Price: ptrFloat(189999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 75, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 17 Pro", Description: ptrStr("Latest Apple Pro"), Slug: "iphone-17-pro", Price: ptrFloat(159999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 76, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 17 Plus", Description: ptrStr("Large iPhone 17"), Slug: "iphone-17-plus", Price: ptrFloat(129999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 77, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 17", Description: ptrStr("Standard iPhone 17"), Slug: "iphone-17", Price: ptrFloat(99999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 78, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 16 Pro Max", Description: ptrStr("Previous Apple flagship"), Slug: "iphone-16-pro-max", Price: ptrFloat(169999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 79, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 16 Pro", Description: ptrStr("Previous Apple Pro"), Slug: "iphone-16-pro", Price: ptrFloat(139999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 80, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 16", Description: ptrStr("Previous standard iPhone"), Slug: "iphone-16", Price: ptrFloat(89999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 81, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone SE", Description: ptrStr("Budget Apple iPhone"), Slug: "iphone-se", Price: ptrFloat(49999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 82, CreatedAt: now, UpdatedAt: now},

		// Google Pixel Series
		{Name: "Google Pixel 9 Pro XL", Description: ptrStr("Premium Google flagship"), Slug: "google-pixel-9-pro-xl", Price: ptrFloat(139999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Google")), ViewsCount: 0, Status: 1, Priority: 83, CreatedAt: now, UpdatedAt: now},
		{Name: "Google Pixel 9 Pro", Description: ptrStr("Latest Google Pro"), Slug: "google-pixel-9-pro", Price: ptrFloat(109999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Google")), ViewsCount: 0, Status: 1, Priority: 84, CreatedAt: now, UpdatedAt: now},
		{Name: "Google Pixel 9", Description: ptrStr("Standard Google flagship"), Slug: "google-pixel-9", Price: ptrFloat(79999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Google")), ViewsCount: 0, Status: 1, Priority: 85, CreatedAt: now, UpdatedAt: now},
		{Name: "Google Pixel 9 Pro Fold", Description: ptrStr("Foldable Google phone"), Slug: "google-pixel-9-pro-fold", Price: ptrFloat(159999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Google")), ViewsCount: 0, Status: 1, Priority: 86, CreatedAt: now, UpdatedAt: now},
		{Name: "Google Pixel 8 Pro", Description: ptrStr("Older Google Pro"), Slug: "google-pixel-8-pro", Price: ptrFloat(99999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Google")), ViewsCount: 0, Status: 1, Priority: 87, CreatedAt: now, UpdatedAt: now},

		// OnePlus Series
		{Name: "OnePlus 13", Description: ptrStr("Latest OnePlus flagship"), Slug: "oneplus-13", Price: ptrFloat(89999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 88, CreatedAt: now, UpdatedAt: now},
		{Name: "OnePlus 13R", Description: ptrStr("OnePlus value variant"), Slug: "oneplus-13r", Price: ptrFloat(74999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 89, CreatedAt: now, UpdatedAt: now},
		{Name: "OnePlus 12", Description: ptrStr("Previous OnePlus flagship"), Slug: "oneplus-12", Price: ptrFloat(69999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 90, CreatedAt: now, UpdatedAt: now},
		{Name: "OnePlus 12R", Description: ptrStr("Previous OnePlus R variant"), Slug: "oneplus-12r", Price: ptrFloat(54999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 91, CreatedAt: now, UpdatedAt: now},
		{Name: "OnePlus Nord 5", Description: ptrStr("OnePlus mid-range 5G"), Slug: "oneplus-nord-5", Price: ptrFloat(39999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 92, CreatedAt: now, UpdatedAt: now},
		{Name: "OnePlus Nord 4", Description: ptrStr("Previous Nord generation"), Slug: "oneplus-nord-4", Price: ptrFloat(34999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 93, CreatedAt: now, UpdatedAt: now},
		{Name: "OnePlus Nord CE5", Description: ptrStr("OnePlus Core Edition"), Slug: "oneplus-nord-ce5", Price: ptrFloat(29999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 94, CreatedAt: now, UpdatedAt: now},
		{Name: "OnePlus Open", Description: ptrStr("OnePlus foldable"), Slug: "oneplus-open", Price: ptrFloat(129999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 95, CreatedAt: now, UpdatedAt: now},

		// Walton (Bangladesh Local)
		{Name: "Walton XANON X1 Ultra", Description: ptrStr("Premium local Bangladesh phone"), Slug: "walton-xanon-x1-ultra", Price: ptrFloat(54999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 96, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton XANON X91", Description: ptrStr("Walton premium series"), Slug: "walton-xanon-x91", Price: ptrFloat(49999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 97, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton NEXG N75", Description: ptrStr("Walton mid-range phone"), Slug: "walton-nexg-n75", Price: ptrFloat(29999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 98, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton NEXG N27", Description: ptrStr("Budget Walton phone"), Slug: "walton-nexg-n27", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 99, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton Orbit Y13", Description: ptrStr("Entry Walton smartphone"), Slug: "walton-orbit-y13", Price: ptrFloat(14999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 100, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton ZENX 2", Description: ptrStr("Walton series phone"), Slug: "walton-zenx-2", Price: ptrFloat(24999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 101, CreatedAt: now, UpdatedAt: now},

		// Infinix
		{Name: "Infinix Note 50 Pro+ 5G", Description: ptrStr("Premium Infinix phone"), Slug: "infinix-note-50-pro-plus-5g", Price: ptrFloat(24999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Infinix")), ViewsCount: 0, Status: 1, Priority: 102, CreatedAt: now, UpdatedAt: now},
		{Name: "Infinix Note 50 Pro", Description: ptrStr("Mid-range Infinix Note"), Slug: "infinix-note-50-pro", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Infinix")), ViewsCount: 0, Status: 1, Priority: 103, CreatedAt: now, UpdatedAt: now},
		{Name: "Infinix Note 40 Pro", Description: ptrStr("Previous Infinix Note"), Slug: "infinix-note-40-pro", Price: ptrFloat(17999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Infinix")), ViewsCount: 0, Status: 1, Priority: 104, CreatedAt: now, UpdatedAt: now},
		{Name: "Infinix Hot 50", Description: ptrStr("Budget Infinix phone"), Slug: "infinix-hot-50", Price: ptrFloat(14999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Infinix")), ViewsCount: 0, Status: 1, Priority: 105, CreatedAt: now, UpdatedAt: now},
		{Name: "Infinix Smart 10", Description: ptrStr("Entry-level Infinix"), Slug: "infinix-smart-10", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Infinix")), ViewsCount: 0, Status: 1, Priority: 106, CreatedAt: now, UpdatedAt: now},

		// Tecno
		{Name: "Tecno CAMON 40 Pro 5G", Description: ptrStr("Camera-focused Tecno phone"), Slug: "tecno-camon-40-pro-5g", Price: ptrFloat(26999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 107, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno CAMON 40 Pro", Description: ptrStr("Tecno camera phone"), Slug: "tecno-camon-40-pro", Price: ptrFloat(23999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 108, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno CAMON 40", Description: ptrStr("Standard Tecno Camon"), Slug: "tecno-camon-40", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 109, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno SPARK 40 5G", Description: ptrStr("Value Tecno smartphone"), Slug: "tecno-spark-40-5g", Price: ptrFloat(16999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 110, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno SPARK 40", Description: ptrStr("Budget Tecno Spark"), Slug: "tecno-spark-40", Price: ptrFloat(13999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 111, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno POVA Slim 5G", Description: ptrStr("Slim Tecno 5G"), Slug: "tecno-pova-slim-5g", Price: ptrFloat(21999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 112, CreatedAt: now, UpdatedAt: now},

		// Symphony (Local)
		{Name: "Symphony G27", Description: ptrStr("Local Symphony phone"), Slug: "symphony-g27", Price: ptrFloat(18999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 113, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Atom 5", Description: ptrStr("Budget Symphony"), Slug: "symphony-atom-5", Price: ptrFloat(9999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 114, CreatedAt: now, UpdatedAt: now},

		// --- Additional models to mirror mobile.json ---
		// Samsung: Remaining S Series
		{Name: "Samsung Galaxy S23 FE", Description: ptrStr("Fan Edition Samsung flagship"), Slug: "samsung-galaxy-s23-fe", Price: ptrFloat(59999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 115, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S22 Ultra", Description: ptrStr("2022 Ultra Samsung flagship"), Slug: "samsung-galaxy-s22-ultra", Price: ptrFloat(109999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 116, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S22+", Description: ptrStr("2022 Plus Samsung flagship"), Slug: "samsung-galaxy-s22-plus", Price: ptrFloat(89999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 117, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy S22", Description: ptrStr("2022 Samsung flagship"), Slug: "samsung-galaxy-s22", Price: ptrFloat(74999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 118, CreatedAt: now, UpdatedAt: now},

		// Samsung: Remaining Z Series
		{Name: "Samsung Galaxy Z Fold 5", Description: ptrStr("Foldable Samsung flagship"), Slug: "samsung-galaxy-z-fold-5", Price: ptrFloat(159999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 119, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy Z Flip 5", Description: ptrStr("Compact foldable Samsung"), Slug: "samsung-galaxy-z-flip-5", Price: ptrFloat(109999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 120, CreatedAt: now, UpdatedAt: now},

		// Samsung: Additional A Series (per JSON)
		{Name: "Samsung Galaxy A36 5G", Description: ptrStr("Mid-range 5G smartphone"), Slug: "samsung-galaxy-a36-5g", Price: ptrFloat(32999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 121, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A26 5G", Description: ptrStr("Affordable 5G smartphone"), Slug: "samsung-galaxy-a26-5g", Price: ptrFloat(28999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 122, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A17 5G", Description: ptrStr("Entry-level 5G smartphone"), Slug: "samsung-galaxy-a17-5g", Price: ptrFloat(24999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 123, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A07", Description: ptrStr("Budget smartphone"), Slug: "samsung-galaxy-a07", Price: ptrFloat(15999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 124, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A06", Description: ptrStr("Entry-level smartphone"), Slug: "samsung-galaxy-a06", Price: ptrFloat(13999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 125, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A06 5G", Description: ptrStr("Entry-level 5G smartphone"), Slug: "samsung-galaxy-a06-5g", Price: ptrFloat(16999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 126, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A34 5G", Description: ptrStr("Balanced mid-range 5G"), Slug: "samsung-galaxy-a34-5g", Price: ptrFloat(29999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 127, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A24", Description: ptrStr("Popular budget Samsung"), Slug: "samsung-galaxy-a24", Price: ptrFloat(21999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 128, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A23", Description: ptrStr("Affordable Samsung A series"), Slug: "samsung-galaxy-a23", Price: ptrFloat(20999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 129, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy A14", Description: ptrStr("Budget Samsung A series"), Slug: "samsung-galaxy-a14", Price: ptrFloat(16999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 130, CreatedAt: now, UpdatedAt: now},

		// Samsung: Additional M Series
		{Name: "Samsung Galaxy M34 5G", Description: ptrStr("Mid-range M series"), Slug: "samsung-galaxy-m34-5g", Price: ptrFloat(28999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 131, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy M15 5G", Description: ptrStr("Entry 5G M series"), Slug: "samsung-galaxy-m15-5g", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 132, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy M14 5G", Description: ptrStr("Budget 5G M series"), Slug: "samsung-galaxy-m14-5g", Price: ptrFloat(17999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 133, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy M53 5G", Description: ptrStr("Previous M series"), Slug: "samsung-galaxy-m53-5g", Price: ptrFloat(32999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 134, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy M33 5G", Description: ptrStr("Affordable M series 5G"), Slug: "samsung-galaxy-m33-5g", Price: ptrFloat(21999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 135, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy M16 5G", Description: ptrStr("Entry M series 5G"), Slug: "samsung-galaxy-m16-5g", Price: ptrFloat(17999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 136, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy M06 5G", Description: ptrStr("Budget M series 5G"), Slug: "samsung-galaxy-m06-5g", Price: ptrFloat(15999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 137, CreatedAt: now, UpdatedAt: now},

		// Samsung: Additional F Series
		{Name: "Samsung Galaxy F16 5G", Description: ptrStr("Affordable F series 5G"), Slug: "samsung-galaxy-f16-5g", Price: ptrFloat(17999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 138, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy F55 5G", Description: ptrStr("Upper mid F series"), Slug: "samsung-galaxy-f55-5g", Price: ptrFloat(32999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 139, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy F54 5G", Description: ptrStr("Popular F series 5G"), Slug: "samsung-galaxy-f54-5g", Price: ptrFloat(29999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 140, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy F34 5G", Description: ptrStr("Mid-range F series"), Slug: "samsung-galaxy-f34-5g", Price: ptrFloat(24999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 141, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy F23", Description: ptrStr("Budget F series"), Slug: "samsung-galaxy-f23", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 142, CreatedAt: now, UpdatedAt: now},
		{Name: "Samsung Galaxy F15", Description: ptrStr("Entry F series"), Slug: "samsung-galaxy-f15", Price: ptrFloat(16999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Samsung")), ViewsCount: 0, Status: 1, Priority: 143, CreatedAt: now, UpdatedAt: now},

		// Xiaomi: Additional flagships
		{Name: "Xiaomi 15S Pro", Description: ptrStr("Special edition Xiaomi 15 Pro"), Slug: "xiaomi-15s-pro", Price: ptrFloat(94999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 144, CreatedAt: now, UpdatedAt: now},
		{Name: "Xiaomi 14", Description: ptrStr("Xiaomi flagship"), Slug: "xiaomi-14", Price: ptrFloat(69999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 145, CreatedAt: now, UpdatedAt: now},
		{Name: "Xiaomi 14T Pro", Description: ptrStr("Performance Xiaomi 14T Pro"), Slug: "xiaomi-14t-pro", Price: ptrFloat(59999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 146, CreatedAt: now, UpdatedAt: now},
		{Name: "Xiaomi 13 Ultra", Description: ptrStr("Previous-gen ultra"), Slug: "xiaomi-13-ultra", Price: ptrFloat(79999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 147, CreatedAt: now, UpdatedAt: now},
		{Name: "Xiaomi 13 Pro", Description: ptrStr("Previous-gen pro"), Slug: "xiaomi-13-pro", Price: ptrFloat(64999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 148, CreatedAt: now, UpdatedAt: now},
		{Name: "Xiaomi 13T Pro", Description: ptrStr("T series performance"), Slug: "xiaomi-13t-pro", Price: ptrFloat(54999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 149, CreatedAt: now, UpdatedAt: now},
		{Name: "Xiaomi 13T", Description: ptrStr("T series value"), Slug: "xiaomi-13t", Price: ptrFloat(49999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 150, CreatedAt: now, UpdatedAt: now},

		// Redmi Note: additional variants
		{Name: "Redmi Note 15 Pro+", Description: ptrStr("Top variant of Note 15"), Slug: "redmi-note-15-pro-plus", Price: ptrFloat(48999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 151, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi Note 14 Pro", Description: ptrStr("Pro variant Note 14"), Slug: "redmi-note-14-pro", Price: ptrFloat(37999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 152, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi Note 13 Pro", Description: ptrStr("Pro variant Note 13"), Slug: "redmi-note-13-pro", Price: ptrFloat(34999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 153, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi Note 13 Pro+ 5G", Description: ptrStr("Top Note 13 variant"), Slug: "redmi-note-13-pro-plus-5g", Price: ptrFloat(39999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 154, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi Note 12 Pro", Description: ptrStr("Pro variant Note 12"), Slug: "redmi-note-12-pro", Price: ptrFloat(29999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 155, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi Note 12 Pro+", Description: ptrStr("Top Note 12 variant"), Slug: "redmi-note-12-pro-plus", Price: ptrFloat(31999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 156, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi Note 12 4G", Description: ptrStr("Note 12 4G model"), Slug: "redmi-note-12-4g", Price: ptrFloat(22999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 157, CreatedAt: now, UpdatedAt: now},

		// Redmi Numbered & A Series
		{Name: "Redmi 15C", Description: ptrStr("Entry-level Redmi 15C"), Slug: "redmi-15c", Price: ptrFloat(14999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 158, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi 14C", Description: ptrStr("Entry-level Redmi 14C"), Slug: "redmi-14c", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 159, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi 12C", Description: ptrStr("Budget Redmi 12C"), Slug: "redmi-12c", Price: ptrFloat(10999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 160, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi 12", Description: ptrStr("Budget Redmi 12"), Slug: "redmi-12", Price: ptrFloat(11999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 161, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi A5", Description: ptrStr("Redmi A series"), Slug: "redmi-a5", Price: ptrFloat(9999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 162, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi A3", Description: ptrStr("Redmi A series"), Slug: "redmi-a3", Price: ptrFloat(8999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 163, CreatedAt: now, UpdatedAt: now},
		{Name: "Redmi A2+", Description: ptrStr("Redmi A series"), Slug: "redmi-a2-plus", Price: ptrFloat(7999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 164, CreatedAt: now, UpdatedAt: now},

		// POCO: additional series
		{Name: "POCO X6 Neo 5G", Description: ptrStr("X6 Neo 5G"), Slug: "poco-x6-neo-5g", Price: ptrFloat(27999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 165, CreatedAt: now, UpdatedAt: now},
		{Name: "POCO M6 Plus 5G", Description: ptrStr("M6 Plus 5G"), Slug: "poco-m6-plus-5g", Price: ptrFloat(21999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 166, CreatedAt: now, UpdatedAt: now},
		{Name: "POCO M6 Plus", Description: ptrStr("M6 Plus"), Slug: "poco-m6-plus", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 167, CreatedAt: now, UpdatedAt: now},
		{Name: "POCO M5", Description: ptrStr("POCO M5"), Slug: "poco-m5", Price: ptrFloat(16999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 168, CreatedAt: now, UpdatedAt: now},
		{Name: "POCO C65", Description: ptrStr("POCO C series"), Slug: "poco-c65", Price: ptrFloat(11999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 169, CreatedAt: now, UpdatedAt: now},
		{Name: "POCO C71", Description: ptrStr("POCO C series"), Slug: "poco-c71", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 170, CreatedAt: now, UpdatedAt: now},
		{Name: "POCO C75", Description: ptrStr("POCO C series"), Slug: "poco-c75", Price: ptrFloat(13999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Xiaomi")), ViewsCount: 0, Status: 1, Priority: 171, CreatedAt: now, UpdatedAt: now},

		// Vivo: additional models
		{Name: "Vivo X200 Pro Mini", Description: ptrStr("Compact X200 Pro"), Slug: "vivo-x200-pro-mini", Price: ptrFloat(59999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 172, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo V60 Lite 4G", Description: ptrStr("V60 Lite 4G"), Slug: "vivo-v60-lite-4g", Price: ptrFloat(34999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 173, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo V50 Lite 5G", Description: ptrStr("V50 Lite 5G"), Slug: "vivo-v50-lite-5g", Price: ptrFloat(32999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 174, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo V50 Lite", Description: ptrStr("V50 Lite"), Slug: "vivo-v50-lite", Price: ptrFloat(29999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 175, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo V40 Lite", Description: ptrStr("V40 Lite"), Slug: "vivo-v40-lite", Price: ptrFloat(27999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 176, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo Y400", Description: ptrStr("Vivo Y series"), Slug: "vivo-y400", Price: ptrFloat(21999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 177, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo Y21d", Description: ptrStr("Vivo budget series"), Slug: "vivo-y21d", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 178, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo Y19s Pro", Description: ptrStr("Vivo Y series"), Slug: "vivo-y19s-pro", Price: ptrFloat(14999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 179, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo Y04", Description: ptrStr("Vivo entry"), Slug: "vivo-y04", Price: ptrFloat(9999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 180, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo Y19s", Description: ptrStr("Vivo Y series"), Slug: "vivo-y19s", Price: ptrFloat(13999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 181, CreatedAt: now, UpdatedAt: now},
		{Name: "Vivo Y03t", Description: ptrStr("Vivo Y entry"), Slug: "vivo-y03t", Price: ptrFloat(9499), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Vivo")), ViewsCount: 0, Status: 1, Priority: 182, CreatedAt: now, UpdatedAt: now},

		// Oppo: additional Reno series per JSON
		{Name: "Oppo Reno 13 F 5G", Description: ptrStr("Reno 13 F 5G"), Slug: "oppo-reno-13-f-5g", Price: ptrFloat(51999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 183, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo Reno 12 5G", Description: ptrStr("Reno 12 5G"), Slug: "oppo-reno-12-5g", Price: ptrFloat(47999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 184, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo Reno 12 F 5G", Description: ptrStr("Reno 12 F 5G"), Slug: "oppo-reno-12-f-5g", Price: ptrFloat(42999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 185, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo Reno 12 Pro 5G", Description: ptrStr("Reno 12 Pro 5G"), Slug: "oppo-reno-12-pro-5g", Price: ptrFloat(54999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 186, CreatedAt: now, UpdatedAt: now},

		// Oppo: additional A series per JSON
		{Name: "Oppo A5x", Description: ptrStr("Oppo A series"), Slug: "oppo-a5x", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 187, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo A3x", Description: ptrStr("Oppo A series"), Slug: "oppo-a3x", Price: ptrFloat(17999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 188, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo A78", Description: ptrStr("Oppo A series"), Slug: "oppo-a78", Price: ptrFloat(22999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 189, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo A38", Description: ptrStr("Oppo A series"), Slug: "oppo-a38", Price: ptrFloat(16999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 190, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo A18", Description: ptrStr("Oppo A series"), Slug: "oppo-a18", Price: ptrFloat(14999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 191, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo A57", Description: ptrStr("Oppo A series"), Slug: "oppo-a57", Price: ptrFloat(14999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 192, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo A55 5G", Description: ptrStr("Oppo A series 5G"), Slug: "oppo-a55-5g", Price: ptrFloat(17999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 193, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo A74 5G", Description: ptrStr("Oppo A series 5G"), Slug: "oppo-a74-5g", Price: ptrFloat(21999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 194, CreatedAt: now, UpdatedAt: now},
		{Name: "Oppo A12", Description: ptrStr("Oppo A series"), Slug: "oppo-a12", Price: ptrFloat(10999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Oppo")), ViewsCount: 0, Status: 1, Priority: 195, CreatedAt: now, UpdatedAt: now},

		// Realme: expand per JSON
		{Name: "Realme 15T 5G", Description: ptrStr("Realme numbered series"), Slug: "realme-15t-5g", Price: ptrFloat(32999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 196, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme 14T 5G", Description: ptrStr("Realme numbered series"), Slug: "realme-14t-5g", Price: ptrFloat(29999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 197, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme 12", Description: ptrStr("Realme numbered series"), Slug: "realme-12", Price: ptrFloat(21999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 198, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme 13 Pro", Description: ptrStr("Realme numbered series"), Slug: "realme-13-pro", Price: ptrFloat(27999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 199, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme C63", Description: ptrStr("Realme C series"), Slug: "realme-c63", Price: ptrFloat(13999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 200, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme C61", Description: ptrStr("Realme C series"), Slug: "realme-c61", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 201, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme Note 70", Description: ptrStr("Realme Note series"), Slug: "realme-note-70", Price: ptrFloat(17999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 202, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme Note 60x", Description: ptrStr("Realme Note series"), Slug: "realme-note-60x", Price: ptrFloat(15999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 203, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme Note 60", Description: ptrStr("Realme Note series"), Slug: "realme-note-60", Price: ptrFloat(14999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 204, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme GT Neo2", Description: ptrStr("Gaming Realme phone"), Slug: "realme-gt-neo2", Price: ptrFloat(39999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 205, CreatedAt: now, UpdatedAt: now},
		{Name: "Realme GT Master Edition", Description: ptrStr("Design-focused GT Master"), Slug: "realme-gt-master-edition", Price: ptrFloat(34999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Realme")), ViewsCount: 0, Status: 1, Priority: 206, CreatedAt: now, UpdatedAt: now},

		// Apple: add missing families from JSON
		{Name: "iPhone 16 Pro Max", Description: ptrStr("iPhone 16 Pro Max"), Slug: "iphone-16-pro-max", Price: ptrFloat(169999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 207, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 16 Pro", Description: ptrStr("iPhone 16 Pro"), Slug: "iphone-16-pro", Price: ptrFloat(139999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 208, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 15 Pro Max", Description: ptrStr("iPhone 15 Pro Max"), Slug: "iphone-15-pro-max", Price: ptrFloat(149999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 209, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 15 Pro", Description: ptrStr("iPhone 15 Pro"), Slug: "iphone-15-pro", Price: ptrFloat(129999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 210, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 15 Plus", Description: ptrStr("iPhone 15 Plus"), Slug: "iphone-15-plus", Price: ptrFloat(99999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 211, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 14 Pro Max", Description: ptrStr("iPhone 14 Pro Max"), Slug: "iphone-14-pro-max", Price: ptrFloat(109999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 212, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 14 Pro", Description: ptrStr("iPhone 14 Pro"), Slug: "iphone-14-pro", Price: ptrFloat(99999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 213, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 14 Plus", Description: ptrStr("iPhone 14 Plus"), Slug: "iphone-14-plus", Price: ptrFloat(79999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 214, CreatedAt: now, UpdatedAt: now},
		{Name: "iPhone 14", Description: ptrStr("iPhone 14"), Slug: "iphone-14", Price: ptrFloat(69999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Apple")), ViewsCount: 0, Status: 1, Priority: 215, CreatedAt: now, UpdatedAt: now},

		// Google: add remaining popular models
		{Name: "Google Pixel 8", Description: ptrStr("Google Pixel 8"), Slug: "google-pixel-8", Price: ptrFloat(79999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Google")), ViewsCount: 0, Status: 1, Priority: 216, CreatedAt: now, UpdatedAt: now},
		{Name: "Google Pixel 8a", Description: ptrStr("Google Pixel 8a"), Slug: "google-pixel-8a", Price: ptrFloat(54999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Google")), ViewsCount: 0, Status: 1, Priority: 217, CreatedAt: now, UpdatedAt: now},

		// OnePlus: additional models
		{Name: "OnePlus 11", Description: ptrStr("Previous OnePlus flagship"), Slug: "oneplus-11", Price: ptrFloat(64999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 218, CreatedAt: now, UpdatedAt: now},
		{Name: "OnePlus Nord CE4 Lite 5G", Description: ptrStr("Nord CE4 Lite 5G"), Slug: "oneplus-nord-ce4-lite-5g", Price: ptrFloat(24999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 219, CreatedAt: now, UpdatedAt: now},
		{Name: "OnePlus Nord N30 SE 5G", Description: ptrStr("Nord N30 SE 5G"), Slug: "oneplus-nord-n30-se-5g", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 220, CreatedAt: now, UpdatedAt: now},
		{Name: "OnePlus N20 SE", Description: ptrStr("OnePlus budget N series"), Slug: "oneplus-n20-se", Price: ptrFloat(16999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "OnePlus")), ViewsCount: 0, Status: 1, Priority: 221, CreatedAt: now, UpdatedAt: now},

		// Walton: add additional distinct models (no uppercase dupes)
		{Name: "Walton XANON X21", Description: ptrStr("Walton XANON series"), Slug: "walton-xanon-x21", Price: ptrFloat(36999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 222, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton NEXG N10 ULTRA", Description: ptrStr("Walton NEXG series"), Slug: "walton-nexg-n10-ultra", Price: ptrFloat(32999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 223, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton NEXG N26", Description: ptrStr("Walton NEXG series"), Slug: "walton-nexg-n26", Price: ptrFloat(21999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 224, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton NEXG N74", Description: ptrStr("Walton NEXG series"), Slug: "walton-nexg-n74", Price: ptrFloat(23999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 225, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton ORBIT Y70C", Description: ptrStr("Walton Orbit series"), Slug: "walton-orbit-y70c", Price: ptrFloat(13999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 226, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton ZENX 1", Description: ptrStr("Walton ZENX series"), Slug: "walton-zenx-1", Price: ptrFloat(19999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 227, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton ZENX 1T", Description: ptrStr("Walton ZENX series"), Slug: "walton-zenx-1t", Price: ptrFloat(21999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 228, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton NEXG N71 PLUS", Description: ptrStr("Walton NEXG series"), Slug: "walton-nexg-n71-plus", Price: ptrFloat(25999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 229, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton NEXG N9", Description: ptrStr("Walton NEXG series"), Slug: "walton-nexg-n9", Price: ptrFloat(14999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 230, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton ORBIT Y12", Description: ptrStr("Walton Orbit series"), Slug: "walton-orbit-y12", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 231, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton ORBIT Y11", Description: ptrStr("Walton Orbit series"), Slug: "walton-orbit-y11", Price: ptrFloat(11999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 232, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton Primo S8", Description: ptrStr("Walton Primo series"), Slug: "walton-primo-s8", Price: ptrFloat(21999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 233, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton Primo GH11", Description: ptrStr("Walton Primo series"), Slug: "walton-primo-gh11", Price: ptrFloat(14999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 234, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton Primo H10", Description: ptrStr("Walton Primo series"), Slug: "walton-primo-h10", Price: ptrFloat(17999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 235, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton Primo ZX4", Description: ptrStr("Walton Primo series"), Slug: "walton-primo-zx4", Price: ptrFloat(26999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 236, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton Primo R8", Description: ptrStr("Walton Primo series"), Slug: "walton-primo-r8", Price: ptrFloat(15999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 237, CreatedAt: now, UpdatedAt: now},
		{Name: "Walton Primo F10", Description: ptrStr("Walton Primo series"), Slug: "walton-primo-f10", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Walton")), ViewsCount: 0, Status: 1, Priority: 238, CreatedAt: now, UpdatedAt: now},

		// Infinix: add per JSON
		{Name: "Infinix Zero 30 5G", Description: ptrStr("Infinix Zero series"), Slug: "infinix-zero-30-5g", Price: ptrFloat(26999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Infinix")), ViewsCount: 0, Status: 1, Priority: 239, CreatedAt: now, UpdatedAt: now},
		{Name: "Infinix Note 50", Description: ptrStr("Infinix Note series"), Slug: "infinix-note-50", Price: ptrFloat(17999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Infinix")), ViewsCount: 0, Status: 1, Priority: 240, CreatedAt: now, UpdatedAt: now},
		{Name: "Infinix Smart 10 Plus", Description: ptrStr("Infinix Smart series"), Slug: "infinix-smart-10-plus", Price: ptrFloat(13999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Infinix")), ViewsCount: 0, Status: 1, Priority: 241, CreatedAt: now, UpdatedAt: now},
		{Name: "Infinix Hot 60i 4G", Description: ptrStr("Infinix Hot series"), Slug: "infinix-hot-60i-4g", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Infinix")), ViewsCount: 0, Status: 1, Priority: 242, CreatedAt: now, UpdatedAt: now},

		// Tecno: expand SPARK/POVA/PHANTOM/CAMON
		{Name: "Tecno SPARK 40C", Description: ptrStr("Tecno SPARK series"), Slug: "tecno-spark-40c", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 243, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno SPARK 40 Pro+", Description: ptrStr("Tecno SPARK top variant"), Slug: "tecno-spark-40-pro-plus", Price: ptrFloat(17999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 244, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno SPARK 40 Pro", Description: ptrStr("Tecno SPARK pro"), Slug: "tecno-spark-40-pro", Price: ptrFloat(15999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 245, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno SPARK 40", Description: ptrStr("Tecno SPARK"), Slug: "tecno-spark-40", Price: ptrFloat(13999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 246, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno SPARK Go 2", Description: ptrStr("Tecno SPARK Go series"), Slug: "tecno-spark-go-2", Price: ptrFloat(10999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 247, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno SPARK 30 Pro", Description: ptrStr("Tecno SPARK 30 Pro"), Slug: "tecno-spark-30-pro", Price: ptrFloat(14999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 248, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno SPARK 30", Description: ptrStr("Tecno SPARK 30"), Slug: "tecno-spark-30", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 249, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno SPARK 30C", Description: ptrStr("Tecno SPARK 30C"), Slug: "tecno-spark-30c", Price: ptrFloat(11999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 250, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno SPARK Go 1", Description: ptrStr("Tecno SPARK Go 1"), Slug: "tecno-spark-go-1", Price: ptrFloat(9999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 251, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno POVA 7 Pro 5G", Description: ptrStr("Tecno POVA series"), Slug: "tecno-pova-7-pro-5g", Price: ptrFloat(22999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 252, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno POVA Curve 5G", Description: ptrStr("Tecno POVA series"), Slug: "tecno-pova-curve-5g", Price: ptrFloat(23999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 253, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno PHANTOM V Fold2 5G", Description: ptrStr("Tecno foldable"), Slug: "tecno-phantom-v-fold2-5g", Price: ptrFloat(129999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 254, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno CAMON 30S", Description: ptrStr("Tecno CAMON 30S"), Slug: "tecno-camon-30s", Price: ptrFloat(16999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 255, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno CAMON 30", Description: ptrStr("Tecno CAMON 30"), Slug: "tecno-camon-30", Price: ptrFloat(15999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 256, CreatedAt: now, UpdatedAt: now},
		{Name: "Tecno CAMON 30 Premier 5G", Description: ptrStr("Tecno CAMON 30 Premier"), Slug: "tecno-camon-30-premier-5g", Price: ptrFloat(21999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Tecno")), ViewsCount: 0, Status: 1, Priority: 257, CreatedAt: now, UpdatedAt: now},

		// Symphony: expand per JSON
		{Name: "Symphony Innova 30", Description: ptrStr("Symphony Innova series"), Slug: "symphony-innova-30", Price: ptrFloat(11999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 258, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Helio 90", Description: ptrStr("Symphony Helio series"), Slug: "symphony-helio-90", Price: ptrFloat(15999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 259, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Helio 50", Description: ptrStr("Symphony Helio series"), Slug: "symphony-helio-50", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 260, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Helio 40", Description: ptrStr("Symphony Helio series"), Slug: "symphony-helio-40", Price: ptrFloat(10999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 261, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Max 60", Description: ptrStr("Symphony Max series"), Slug: "symphony-max-60", Price: ptrFloat(13999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 262, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Innova 40", Description: ptrStr("Symphony Innova series"), Slug: "symphony-innova-40", Price: ptrFloat(9999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 263, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Z72", Description: ptrStr("Symphony Z series"), Slug: "symphony-z72", Price: ptrFloat(12999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 264, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Z70", Description: ptrStr("Symphony Z series"), Slug: "symphony-z70", Price: ptrFloat(11999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 265, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony S72", Description: ptrStr("Symphony S series"), Slug: "symphony-s72", Price: ptrFloat(10999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 266, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony S75", Description: ptrStr("Symphony S series"), Slug: "symphony-s75", Price: ptrFloat(11999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 267, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Xplorer ZV", Description: ptrStr("Symphony Xplorer series"), Slug: "symphony-xplorer-zv", Price: ptrFloat(9999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 268, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Xplorer V55", Description: ptrStr("Symphony Xplorer V55"), Slug: "symphony-xplorer-v55", Price: ptrFloat(8999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 269, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Xplorer W20", Description: ptrStr("Symphony Xplorer W20"), Slug: "symphony-xplorer-w20", Price: ptrFloat(7999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 270, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Xplorer W68", Description: ptrStr("Symphony Xplorer W68"), Slug: "symphony-xplorer-w68", Price: ptrFloat(8999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 271, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony Xplorer W92", Description: ptrStr("Symphony Xplorer W92"), Slug: "symphony-xplorer-w92", Price: ptrFloat(10999), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 272, CreatedAt: now, UpdatedAt: now},
		{Name: "Symphony i10", Description: ptrStr("Symphony i series"), Slug: "symphony-i10", Price: ptrFloat(7499), CategoryID: &catID, BrandID: ptrUint(getBrandID(db, "Symphony")), ViewsCount: 0, Status: 1, Priority: 273, CreatedAt: now, UpdatedAt: now},
	}

	// Remove duplicate slugs within the seeder list and skip products
	// that already exist in the database (prevents unique constraint errors).
	seen := make(map[string]bool)
	var toInsert []models.ProductModel
	skippedExisting := 0
	skippedDuplicateInList := 0

	for _, p := range products {
		if p.Slug == "" {
			// if slug is empty, skip (shouldn't happen for these seeders)
			log.Printf("Skipping product with empty slug: %s", p.Name)
			skippedDuplicateInList++
			continue
		}

		if seen[p.Slug] {
			skippedDuplicateInList++
			log.Printf("Skipping duplicate slug in seeder list: %s", p.Slug)
			continue
		}

		// mark as seen for list-level dedupe
		seen[p.Slug] = true

		// check if slug already exists in DB
		var existing models.ProductModel
		err := db.Where("slug = ?", p.Slug).First(&existing).Error
		if err == nil {
			skippedExisting++
			log.Printf("Skipping existing product with slug: %s", p.Slug)
			continue
		}
		if err != nil && err != gorm.ErrRecordNotFound {
			// unexpected DB error while checking
			log.Printf("❌ Error checking existing product slug %s: %v", p.Slug, err)
			return err
		}

		toInsert = append(toInsert, p)
	}

	if len(toInsert) == 0 {
		fmt.Printf("✅ No new mobile products to seed (skipped %d existing, %d duplicates in list)\n", skippedExisting, skippedDuplicateInList)
		return nil
	}

	// Batch create only the new items
	result := db.CreateInBatches(toInsert, 100)
	if result.Error != nil {
		log.Printf("❌ Error seeding products: %v", result.Error)
		return result.Error
	}

	fmt.Printf("✅ Successfully seeded %d mobile products (skipped %d existing, %d duplicates in list)\n", result.RowsAffected, skippedExisting, skippedDuplicateInList)
	return nil
}

// Helper functions
func getBrandID(db *gorm.DB, brandName string) uint {
	var brand models.BrandModel
	db.Where("name = ?", brandName).First(&brand)
	return brand.ID
}

func ptrStr(s string) *string {
	return &s
}

func ptrFloat(f float64) *float64 {
	return &f
}

func ptrUint(u uint) *uint {
	return &u
}
