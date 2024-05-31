package adapters

import (
	"github.com/fiatfour/go-unit/core"
	"gorm.io/gorm"
)

// * Secondary adapter (gorm_adapter.go)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order core.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		// Handle database errors
		return result.Error
	}
	return nil
}
