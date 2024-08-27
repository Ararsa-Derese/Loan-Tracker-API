package repository

import (
	"context"
	"loan/database"
	"loan/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanRepository struct {
	database   database.Database
	collection string
}

// CreateLoan implements domain.LoanRepository.
func (l *LoanRepository) CreateLoan(c context.Context, loan *domain.Loan) error {
	collection := l.database.Collection(l.collection)
	_, err := collection.InsertOne(c, loan)
	return err
}

// GetLoans implements domain.LoanRepository.

func (l *LoanRepository) GetLoans(c context.Context) ([]*domain.Loan, error) {
	collection := l.database.Collection(l.collection)
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(c)

	var loans []*domain.Loan
	for cursor.Next(c) {
		var loan domain.Loan
		if err := cursor.Decode(&loan); err != nil {
			return nil, err
		}
		loans = append(loans, &loan)
	}
	return loans, nil
}

// DeleteLoan implements domain.LoanRepository.
func (l *LoanRepository) DeleteLoan(c context.Context, id primitive.ObjectID) error {
	collection := l.database.Collection(l.collection)
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(c, filter)
	return err
}

// GetLoanByID implements domain.LoanRepository.
func (l *LoanRepository) GetLoanByID(c context.Context, id primitive.ObjectID,userID primitive.ObjectID) (*domain.Loan, error) {
	collection := l.database.Collection(l.collection)
	filter := bson.M{"_id": id,"user_id":userID}
	loan := &domain.Loan{}
	err := collection.FindOne(c, filter).Decode(loan)
	if err != nil {
		return nil, err
	}
	return loan, err
}
// GetLoansByUserID implements domain.LoanRepository.
func (l *LoanRepository) GetLoansByUserID(c context.Context, userid primitive.ObjectID) ([]*domain.Loan, error) {
	collection := l.database.Collection(l.collection)
	filter := bson.M{"user_id": userid}
	cursor, err := collection.Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var loans []*domain.Loan
	for cursor.Next(c) {
		var loan domain.Loan
		if err := cursor.Decode(&loan); err != nil {
			return nil, err
		}
		loans = append(loans, &loan)
	}
	return loans, nil
}

// UpdateLoanStatus implements domain.LoanRepository.
func (l *LoanRepository) UpdateLoanStatus(c context.Context, id primitive.ObjectID,status string) error {
	collection := l.database.Collection(l.collection)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status}}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func NewLoanRepository(db database.Database, collection string) domain.LoanRepository {
	return &LoanRepository{
		database:   db,
		collection: collection,
	}
}
