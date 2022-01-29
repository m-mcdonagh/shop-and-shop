package db

import (
	"time"

	"gorm.io/gorm"
)

type namedModel struct {
	gorm.Model
	Name string `gorm:"collate:nocase;"`
}

type Product struct {
	namedModel

	Stores  []*Store  `gorm:"many2many:store_products;"`
	Recipes []*Recipe `gorm:"many2many:ingredients;"`
}

type Store struct {
	namedModel

	Products []*Product `gorm:"many2many:store_products;"`
}

type Recipe struct {
	namedModel

	Products []*Product `gorm:"many2many:ingredients;"`
}

type StoreProduct struct {
	StoreID   int `gorm:"primaryKey"`
	ProductID int `gorm:"primaryKey"`
	Amount    float64
	Price     int // Price in cents
}

type Ingredient struct {
	ProductID int `gorm:"primaryKey"`
	RecipeID  int `gorm:"primaryKey"`
	Amount    float64
}

type MealPlan struct {
	namedModel

	Recipes []*Recipe `gorm:"many2many:meals;"`
}

type Meal struct {
	MealPlanID int `gorm:"primaryKey"`
	RecipeID   int `gorm:"primaryKey"`

	Date time.Time
}
