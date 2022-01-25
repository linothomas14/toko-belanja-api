package domain

import (
	"time"
)

type Product struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title      string    `json:"title" gorm:"notNull"`
	Price      int64     `json:"price" gorm:"notNull"`
	Stock      int64     `json:"stock" gorm:"notNull"`
	CategoryID int64     `json:"category_Id"  gorm:"notNull"`
	CreatedAt  time.Time `json:"created_at" gorm:"notNull"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"notNull"`
}
