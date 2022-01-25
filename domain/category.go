package domain

import (
	"context"
	"time"
)

type Category struct {
	ID                  int64     `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Type                string    `json:"type" gorm:"notNull"`
	Sold_Product_Amount int64     `json:"sold_product_amount" gorm:"notNull"`
	CreatedAt           time.Time `json:"created_at" gorm:"notNull"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"notNull"`
	Products            []Product `gorm:"foreignKey:category_Id`
}

type CategoryUsecase interface {
	GetCategories(ctx context.Context) (interface{}, error)
	StoreCategory(ctx context.Context, category *Category) (Category, error)
	UpdateCategory(ctx context.Context, category *Category) (Category, error)
	DeleteCategory(ctx context.Context, id int64) error
}

type CategoryRepository interface {
	GetCategories(ctx context.Context) (interface{}, error)
	StoreCategory(ctx context.Context, category *Category) (categoryId int64, err error)
	GetCategoryByID(ctx context.Context, id int64) (Category, error)
	UpdateCategory(ctx context.Context, category *Category) error
	DeleteCategory(ctx context.Context, id int64) error
}
