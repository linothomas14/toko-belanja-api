package repository

import (
	"context"
	"time"

	"github.com/linothomas14/toko-belanja-api/domain"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	Conn *gorm.DB
}

func NewCategoryRepository(Conn *gorm.DB) domain.CategoryRepository {
	return &CategoryRepository{Conn}
}

func (c *CategoryRepository) GetCategories(ctx context.Context) (interface{}, error) {

	type Product struct {
		ID        int64     `json:"id" gorm:"primaryKey;autoIncrement:true"`
		Title     string    `json:"title" gorm:"notNull"`
		Price     int64     `json:"price" gorm:"notNull"`
		Stock     int64     `json:"stock" gorm:"notNull"`
		CreatedAt time.Time `json:"created_at" gorm:"notNull"`
		UpdatedAt time.Time `json:"updated_at" gorm:"notNull"`
	}

	var categoriesProducts []domain.Category
	err := c.Conn.Order("id").Preload("Products").Find(&categoriesProducts).Error
	if err != nil {
		return []domain.Category{}, err
	}
	return categoriesProducts, nil
}

func (c *CategoryRepository) StoreCategory(ctx context.Context, category *domain.Category) (categoryId int64, err error) {
	err = c.Conn.Create(category).Error
	if err != nil {
		return
	}
	categoryId = category.ID
	return
}

func (c *CategoryRepository) GetCategoryByID(ctx context.Context, id int64) (domain.Category, error) {
	var category domain.Category
	err := c.Conn.First(&category, "id=?", id).Error
	if err != nil {
		return domain.Category{}, err
	}
	return category, nil
}

func (c *CategoryRepository) UpdateCategory(ctx context.Context, category *domain.Category) error {
	err := c.Conn.Model(category).Updates(category).Error
	return err
}

func (c *CategoryRepository) DeleteCategory(ctx context.Context, id int64) error {
	var category domain.Category
	category.ID = id
	err := c.Conn.Delete(&category).Error
	return err
}
