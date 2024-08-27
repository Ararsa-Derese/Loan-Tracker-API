package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionLoan = "loans"

type Loan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	Amount      float64            `json:"amount" bson:"amount"`
	Status      string             `json:"status" bson:"status"`
	SubmittedAt time.Time          `json:"submitted_at" bson:"submitted_at"`
	ReviewedAt  time.Time          `json:"reviewed_at" bson:"reviewed_at"`
	ReviewedBy  string             `json:"reviewed_by" bson:"reviewed_by"`
}

type LoanRequest struct {
	UserID primitive.ObjectID `bson:"user_id,omitempty"`
	Amount float64            `json:"amount" bson:"amount"`
}

type LoanUsecase interface {
	CreateLoan(c context.Context, loan *LoanRequest) (*Loan,error)
	GetLoans(c context.Context, claims *JwtCustomClaims) ([]*Loan, error)
	GetLoanByID(c context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*Loan, error)
	GetLoansByUserID(c context.Context, userid primitive.ObjectID) ([]*Loan, error)
	UpdateLoanStatus(c context.Context, claims *JwtCustomClaims, status string, id primitive.ObjectID) error
	DeleteLoan(c context.Context,  claims *JwtCustomClaims, id primitive.ObjectID) error
}

type LoanRepository interface {
	CreateLoan(c context.Context, loan *Loan) error
	GetLoans(c context.Context) ([]*Loan, error)
	GetLoanByID(c context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*Loan, error)
	GetLoansByUserID(c context.Context, userid primitive.ObjectID) ([]*Loan, error)
	UpdateLoanStatus(c context.Context, id primitive.ObjectID, status string) error
	DeleteLoan(c context.Context, id primitive.ObjectID) error
}
