package postgres

import (
	"fmt"

	"github.com/aryanicosa/golang-fullstack-lms/internal/config"
	"github.com/aryanicosa/golang-fullstack-lms/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.DBHost,
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBName,
        cfg.DBPort,
    )
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Auto migrate the schema
    err = db.AutoMigrate(
        &models.User{},
    )
    if err != nil {
        return nil, err
    }

    return db, nil
}

func ClosePostgresDB(db *gorm.DB) error {
    sqlDB, err := db.DB()
    if err != nil {
        return err
    }

    return sqlDB.Close()
}