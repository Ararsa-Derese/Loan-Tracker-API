package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionLogs = "logs"


type Log struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`
	Activity string `json:"activity" bson:"activity"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
}

type LogUsecase interface {
	CreateLog(c context.Context,activity string, userID primitive.ObjectID) error
	GetLogs(c context.Context,claims *JwtCustomClaims) ([]*Log, error)
}

type LogRepository interface {
	CreateLog(c context.Context,log *Log) error
	GetLogs(c context.Context) ([]*Log, error)
}