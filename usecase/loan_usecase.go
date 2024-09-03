package usecase

import (
	"context"
	"errors"
	"loan/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanUsecase struct {
	LoanRepository domain.LoanRepository
	contextTimeout time.Duration
}

// CreateLoan implements domain.LoanUsecase.
func (l *LoanUsecase) CreateLoan(c context.Context, loan *domain.LoanRequest) (*domain.Loan,error) {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	newloan := &domain.Loan{
		ID:          primitive.NewObjectID(),
		UserID:      loan.UserID,
		Amount:      loan.Amount,
		Status:      "submitted",
		SubmittedAt: time.Now(),
		ReviewedBy:  "",
	}
	err := l.LoanRepository.CreateLoan(ctx, newloan)
	if err != nil {
		return nil,err
	}
	return newloan,nil
}

// GetLoans implements domain.LoanUsecase.
func (l *LoanUsecase) GetLoans(c context.Context, claims *domain.JwtCustomClaims) ([]*domain.Loan, error) {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	if claims.Role == "admin" || claims.Role == "root" {
		return l.LoanRepository.GetLoans(ctx)
	}
	return nil, errors.New("only admins and root users can view all loans")

}

// DeleteLoan implements domain.LoanUsecase.
func (l *LoanUsecase) DeleteLoan(c context.Context, claims *domain.JwtCustomClaims, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	if claims.Role == "admin" || claims.Role == "root" {
		return l.LoanRepository.DeleteLoan(ctx, id)
	}
	return errors.New("only admins and root users can delete loans")
}

// GetLoanByID implements domain.LoanUsecase.
func (l *LoanUsecase) GetLoanByID(c context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*domain.Loan, error) {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()

	loan, err := l.LoanRepository.GetLoanByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	return loan, nil
}

// GetLoansByUserID implements domain.LoanUsecase.
func (l *LoanUsecase) GetLoansByUserID(c context.Context, userid primitive.ObjectID) ([]*domain.Loan, error) {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()

	loans, err := l.LoanRepository.GetLoansByUserID(ctx, userid)
	if err != nil {
		return nil, err
	}
	return loans, nil
}

// UpdateLoanStatus implements domain.LoanUsecase.
func (l *LoanUsecase) UpdateLoanStatus(c context.Context, claims *domain.JwtCustomClaims, status string, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	if claims.Role == "admin" || claims.Role == "root" {
		return l.LoanRepository.UpdateLoanStatus(ctx, id,status)
	}
	return  errors.New("only admins and root users can update loans")
}

// checkRefreshToken implements domain.LoginUsecase.

func NewLoanUsecase(LoanRepository domain.LoanRepository, timeout time.Duration) domain.LoanUsecase {
	return &LoanUsecase{
		LoanRepository: LoanRepository,
		contextTimeout: timeout,
	}
}
