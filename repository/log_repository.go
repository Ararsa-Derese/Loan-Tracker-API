package repository

import (
	"context"
	"loan/database"
	"loan/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type LogRepository struct {
	database   database.Database
	collection string
}

func NewLogRepository(database database.Database) domain.LogRepository {
	return &LogRepository{
		database:   database,
		collection: domain.CollectionLogs,
	}
}

// CreateLog implements domain.LogRepository.
func (l *LogRepository) CreateLog(c context.Context, log *domain.Log) error {
	collection := l.database.Collection(l.collection)
	_, err := collection.InsertOne(c, log)
	return err
}

// GetLogs implements domain.LogRepository.
func (l *LogRepository) GetLogs(c context.Context) ([]*domain.Log, error) {
	collection := l.database.Collection(l.collection)
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(c)

	var logs []*domain.Log
	for cursor.Next(c) {
		var log domain.Log
		if err := cursor.Decode(&log); err != nil {
			return nil, err
		}
		logs = append(logs, &log)
	}
	return logs, nil
}
