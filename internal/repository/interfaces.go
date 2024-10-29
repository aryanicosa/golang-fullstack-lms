package repository

import (
    "context"
    "time"

    "github.com/aryanicosa/golang-fullstack-lms/internal/models"
)

type ExamRepository interface {
    // Database methods
    GetExamByID(ctx context.Context, id uint) (*models.Exam, error)
    ListExams(ctx context.Context) ([]*models.Exam, error)
    CreateExam(ctx context.Context, exam *models.Exam) error
    UpdateExam(ctx context.Context, exam *models.Exam) error
    DeleteExam(ctx context.Context, id uint) error

    // Cache methods
    GetFromCache(ctx context.Context, key string) (interface{}, error)
    SetInCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error
    DeleteFromCache(ctx context.Context, key string) error
}
