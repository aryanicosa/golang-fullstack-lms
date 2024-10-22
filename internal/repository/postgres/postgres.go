package postgres

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "fullstack-lms-go/internal/config"
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
        &models.Exam{},
        &models.Question{},
        &models.Result{},
        &models.Subscription{},
    )
    if err != nil {
        return nil, err
    }

    return db, nil
}