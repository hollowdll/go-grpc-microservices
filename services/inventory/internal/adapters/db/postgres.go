package db

import (
	"fmt"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/config"
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
