package db

import (
	"context"
	"fmt"
	"time"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/config"
	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ProductCode     string `gorm:"primaryKey"`
	Name            string
	Description     string
	UnitPriceCents  int32
	QuantityInStock int32
	CreatedAtMillis int64
	UpdatedAtMillis int64
}

type PostgresAdapter struct {
	db *gorm.DB
}

func NewPostgresAdapter(cfg *config.Config) (*PostgresAdapter, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  buildDSN(&cfg.DB),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("database connection error: %v", err)
	}

	// Auto migrate db schema only in development or testing environments.
	// In production it is recommended to manage db schema versioning explicitly
	// and to use dedicated migration tools.
	if cfg.IsDevelopmentMode() || cfg.IsTestingMode() {
		err = db.AutoMigrate(&Product{})
		if err != nil {
			return nil, fmt.Errorf("database migration error: %v", err)
		}
	}

	return &PostgresAdapter{db: db}, nil
}

func buildDSN(cfg *config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode,
	)
}

func (a *PostgresAdapter) GetProductsByCode(ctx context.Context, productCodes []string) ([]*domain.Product, error) {
	var productModels []Product
	res := a.db.WithContext(ctx).Where("product_code IN ?", productCodes).Find(&productModels)
	if res.Error != nil {
		return nil, res.Error
	}
	var products []*domain.Product
	for _, pm := range productModels {
		products = append(products, &domain.Product{
			ProductCode:     pm.ProductCode,
			Name:            pm.Name,
			Description:     pm.Description,
			UnitPriceCents:  pm.UnitPriceCents,
			QuantityInStock: pm.QuantityInStock,
			CreatedAtMillis: pm.CreatedAtMillis,
			UpdatedAtMillis: pm.UpdatedAtMillis,
		})
	}
	return products, res.Error
}

func (a *PostgresAdapter) UpdateProductStockQuantities(ctx context.Context, products []*domain.ProductQuantity) error {
	// Bulk update transaction
	return a.db.Transaction(func(tx *gorm.DB) error {
		for _, product := range products {
			now := time.Now().UnixMilli()
			if err := tx.WithContext(ctx).Model(&Product{}).Where("product_code = ?", product.ProductCode).Updates(Product{
				QuantityInStock: product.Quantity,
				UpdatedAtMillis: now,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (a *PostgresAdapter) SaveProducts(ctx context.Context, products []*domain.Product) error {
	// Bulk insert transaction
	return a.db.Transaction(func(tx *gorm.DB) error {
		for _, product := range products {
			productModel := Product{
				ProductCode:     product.ProductCode,
				Name:            product.Name,
				Description:     product.Description,
				UnitPriceCents:  product.UnitPriceCents,
				QuantityInStock: product.QuantityInStock,
				CreatedAtMillis: product.CreatedAtMillis,
				UpdatedAtMillis: product.UpdatedAtMillis,
			}
			if err := tx.WithContext(ctx).Create(&productModel).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
