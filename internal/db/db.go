package db

import (
	"fmt"

	"github.com/Rodco-off/checkCode/domain"
	"github.com/Rodco-off/checkCode/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Migrate(conn *gorm.DB) error {
	return conn.AutoMigrate(domain.Task{}, domain.Submission{})

}
