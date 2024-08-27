package usecase

import (
	"context"
	"errors"
	"loan/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &ProfileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *ProfileUsecase) GetProfile(c context.Context, userid primitive.ObjectID) (*domain.ProfileResponse, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	User, err := pu.userRepository.GetUserByID(ctx, userid)
	if User == nil && err != nil {
		return nil, errors.New("failed to get profile")
	}
	return &domain.ProfileResponse{
		First_Name:      User.First_Name,
		Last_Name:       User.Last_Name,
		Bio:             User.Bio,
		Profile_Picture: User.Profile_Picture,
		Contact_Info:    User.Contact_Info,
	}, nil
}
