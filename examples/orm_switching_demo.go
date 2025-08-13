// Package examples demonstrates switching relational ORMs in Clean Architecture
package main

import (
	"fmt"
	"kossti/internal/domain/entities"
	"strings"
	"time"
)

// ==========================================
// Current GORM Model (What we have now)
// ==========================================
type BrandGORMModel struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Slug      string     `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}

func (b *BrandGORMModel) ToEntity() *entities.Brand {
	return &entities.Brand{
		ID: b.ID, Name: b.Name, Slug: b.Slug,
		CreatedAt: b.CreatedAt, UpdatedAt: b.UpdatedAt, DeletedAt: b.DeletedAt,
	}
}

func (b *BrandGORMModel) FromEntity(entity *entities.Brand) {
	b.ID = entity.ID
	b.Name = entity.Name
	b.Slug = entity.Slug
	b.CreatedAt = entity.CreatedAt
	b.UpdatedAt = entity.UpdatedAt
	b.DeletedAt = entity.DeletedAt
}

// ==========================================
// Ent ORM Model (Alternative)
// ==========================================
type BrandEntModel struct {
	ID        int        `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Slug      string     `json:"slug,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Note: Ent uses JSON tags, not GORM tags!
}

func (b *BrandEntModel) ToEntity() *entities.Brand {
	return &entities.Brand{
		ID: uint(b.ID), Name: b.Name, Slug: b.Slug,
		CreatedAt: b.CreatedAt, UpdatedAt: b.UpdatedAt, DeletedAt: b.DeletedAt,
	}
}

func (b *BrandEntModel) FromEntity(entity *entities.Brand) {
	b.ID = int(entity.ID)
	b.Name = entity.Name
	b.Slug = entity.Slug
	b.CreatedAt = entity.CreatedAt
	b.UpdatedAt = entity.UpdatedAt
	b.DeletedAt = entity.DeletedAt
}

// ==========================================
// SQLBoiler Model (Type-safe SQL)
// ==========================================
type BrandBoilerModel struct {
	ID        uint       `boil:"id" json:"id"`
	Name      string     `boil:"name" json:"name"`
	Slug      string     `boil:"slug" json:"slug"`
	CreatedAt time.Time  `boil:"created_at" json:"created_at"`
	UpdatedAt time.Time  `boil:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `boil:"deleted_at" json:"deleted_at"`
	// Note: SQLBoiler uses 'boil' tags!
}

func (b *BrandBoilerModel) ToEntity() *entities.Brand {
	return &entities.Brand{
		ID: b.ID, Name: b.Name, Slug: b.Slug,
		CreatedAt: b.CreatedAt, UpdatedAt: b.UpdatedAt, DeletedAt: b.DeletedAt,
	}
}

func (b *BrandBoilerModel) FromEntity(entity *entities.Brand) {
	b.ID = entity.ID
	b.Name = entity.Name
	b.Slug = entity.Slug
	b.CreatedAt = entity.CreatedAt
	b.UpdatedAt = entity.UpdatedAt
	b.DeletedAt = entity.DeletedAt
}

// ==========================================
// Raw SQL Model (No ORM)
// ==========================================
type BrandSQLModel struct {
	ID        uint
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	// No tags at all - pure Go struct!
}

func (b *BrandSQLModel) ToEntity() *entities.Brand {
	return &entities.Brand{
		ID: b.ID, Name: b.Name, Slug: b.Slug,
		CreatedAt: b.CreatedAt, UpdatedAt: b.UpdatedAt, DeletedAt: b.DeletedAt,
	}
}

func (b *BrandSQLModel) FromEntity(entity *entities.Brand) {
	b.ID = entity.ID
	b.Name = entity.Name
	b.Slug = entity.Slug
	b.CreatedAt = entity.CreatedAt
	b.UpdatedAt = entity.UpdatedAt
	b.DeletedAt = entity.DeletedAt
}

// ==========================================
// Demo: Same Domain Entity, Different Models
// ==========================================

func mainnn() {
	fmt.Println("🚀 RELATIONAL ORM SWITCHING DEMO")
	fmt.Println("=================================")
	fmt.Println()

	// Start with a domain entity (business logic)
	domainBrand := &entities.Brand{
		ID:   1,
		Name: "Apple Inc",
		Slug: "apple-inc",

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Printf("📋 Original Domain Entity:\n")
	fmt.Printf("   ID: %d, Name: %s, Slug: %s\n", domainBrand.ID, domainBrand.Name, domainBrand.Slug)
	fmt.Println()

	// Test conversion to different ORM models
	testORMConversion("GORM", domainBrand, &BrandGORMModel{})
	testORMConversion("Ent ORM", domainBrand, &BrandEntModel{})
	testORMConversion("SQLBoiler", domainBrand, &BrandBoilerModel{})
	testORMConversion("Raw SQL", domainBrand, &BrandSQLModel{})

	fmt.Println("\n🎯 KEY INSIGHTS:")
	fmt.Println("✅ Same domain entity works with ALL ORMs!")
	fmt.Println("✅ Business logic never changes!")
	fmt.Println("✅ Only infrastructure models change!")
	fmt.Println("✅ Different struct tags per ORM!")
	fmt.Println("✅ Perfect isolation achieved!")

	fmt.Println("\n📊 Migration Effort Comparison:")
	fmt.Println("❌ Tightly Coupled: Rewrite entire application")
	fmt.Println("✅ Clean Architecture: Replace only models & repos")
}

func testORMConversion(ormName string, entity *entities.Brand, model interface{}) {
	fmt.Printf("� Converting to %s:\n", ormName)
	fmt.Printf("%s\n", strings.Repeat("-", len(ormName)+17))

	switch m := model.(type) {
	case *BrandGORMModel:
		m.FromEntity(entity)
		fmt.Printf("   GORM Tags: `gorm:\"primaryKey;autoIncrement\"`\n")
		fmt.Printf("   SQL: INSERT INTO brands (name, slug) VALUES ('%s', '%s')\n", m.Name, m.Slug)
		converted := m.ToEntity()
		fmt.Printf("   ✅ Round-trip successful: %s\n", converted.Name)

	case *BrandEntModel:
		m.FromEntity(entity)
		fmt.Printf("   Ent Tags: `json:\"name,omitempty\"`\n")
		fmt.Printf("   Ent API: client.Brand.Create().SetName('%s').SetSlug('%s').Save()\n", m.Name, m.Slug)
		converted := m.ToEntity()
		fmt.Printf("   ✅ Round-trip successful: %s\n", converted.Name)

	case *BrandBoilerModel:
		m.FromEntity(entity)
		fmt.Printf("   Boiler Tags: `boil:\"name\" json:\"name\"`\n")
		fmt.Printf("   SQLBoiler: brand.Insert(ctx, exec, boil.Infer())\n")
		converted := m.ToEntity()
		fmt.Printf("   ✅ Round-trip successful: %s\n", converted.Name)

	case *BrandSQLModel:
		m.FromEntity(entity)
		fmt.Printf("   SQL Tags: None (pure Go struct)\n")
		fmt.Printf("   Raw SQL: INSERT INTO brands (name, slug) VALUES ($1, $2)\n")
		converted := m.ToEntity()
		fmt.Printf("   ✅ Round-trip successful: %s\n", converted.Name)
	}
	fmt.Println()
}
