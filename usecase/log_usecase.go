package usecase

import (
	"context"
	"errors"
	"loan/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogUsecase struct {
	LogRepository domain.LogRepository
	contextTimeout time.Duration
}

func NewLogUsecase(logRepository domain.LogRepository, timeout time.Duration) domain.LogUsecase {
	return &LogUsecase{
		LogRepository: logRepository,
		contextTimeout: timeout,
	}
}

func (l *LogUsecase) CreateLog(c context.Context, activity string, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	newlog := &domain.Log{
		ID:        primitive.NewObjectID(),
		UserID:    userID,
		Activity:  activity,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := l.LogRepository.CreateLog(ctx, newlog)
	if err != nil {
		return err
	}
	return nil
}

func (l *LogUsecase) GetLogs(c context.Context, claims *domain.JwtCustomClaims) ([]*domain.Log, error) {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	if claims.Role == "admin" || claims.Role == "root" {
		return l.LogRepository.GetLogs(ctx)
	}
	return nil, errors.New("only admins and root users can view all logs")

}

