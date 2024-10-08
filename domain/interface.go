package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Role is a type for user roles
type UserRepository interface {
	GetUserByID(c context.Context, id primitive.ObjectID) (*User, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserByUsername(c context.Context, username string) (*User, error)
	GetAllUsers(c context.Context) ([]*User, error)
	CreateUser(c context.Context, user *User) error
	UpdateUser(c context.Context, user *User) error
	UpdatePassword(c context.Context, user *User) error
	PromoteUser(c context.Context, id primitive.ObjectID)  error
	DemoteUser(c context.Context, id primitive.ObjectID)  error
	DeleteUser(c context.Context, id primitive.ObjectID) error
}

// SignupRepository is an interface that contains the CreateUser method

type UserUsecase interface {
	GetUserByID(c context.Context, id primitive.ObjectID) (*User, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserByUsername(c context.Context, username string) (*User, error)
	GetAllUsers(c context.Context) ([]*User, error)
	CreateUser(c context.Context, user *CreateUser,claims *JwtCustomClaims) error
	PromoteUser(c context.Context, id primitive.ObjectID,claims *JwtCustomClaims) error
	DemoteUser(c context.Context, id primitive.ObjectID,claims *JwtCustomClaims) error
	UpdateUser(c context.Context, user *User,claims *JwtCustomClaims, existinguser *User) (*User, error)
	DeleteUser(c context.Context, id primitive.ObjectID,claims *JwtCustomClaims) error
}

type SignupUsecase interface {
	RegisterUser(c context.Context, user *AuthSignup) (*primitive.ObjectID, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserByUsername(c context.Context, username string) (*User, error)
	CreateAccessToken(user *AuthSignup, secret string, expiry int) (string, error)
	CreateRefreshToken(user *AuthSignup, secret string, expiry int) (string, error)
	SaveRefreshToken(c context.Context, token string, id primitive.ObjectID) error
	VerifyOTP(c context.Context, otp *OTPRequest) (*OTP, error)
	SendOTP(c context.Context, user *AuthSignup, username, password string, devicePrint string) error
}

type ProfileUsecase interface {
	GetProfile(c context.Context, userid primitive.ObjectID) (*ProfileResponse, error)
}

type LoginUsecase interface {
	AuthenticateUser(c context.Context, login *AuthLogin) (*User, error)
	CreateAccessToken(user *User, secret string, expiry int) (string, error)
	CreateRefreshToken(user *User, secret string, expiry int) (string, error)
	SaveRefreshToken(c context.Context, token *Token) error
	CheckRefreshToken(c context.Context, refreshToken string) (*Token, error)
}
type TokenRepository interface {
	SaveToken(ctx context.Context, token *Token) error
	FindTokenByAccessToken(ctx context.Context, accessToken string) (*Token, error)
	DeleteToken(ctx context.Context, tokenID primitive.ObjectID) error
	FindTokenByRefreshToken(ctx context.Context, refreshToken string) (*Token, error)
}

type OTPRepository interface {
	GetOTPByEmail(ctx context.Context, email string) (*OTP, error)
	SaveOTP(c context.Context, otp *OTP) error
	DeleteOTP(c context.Context, email string) error
}

type LogoutUsecase interface {
	Logout(ctx context.Context, refreshToken string, deviceFingerprint string) error
}

type ForgotPasswordUsecase interface {
	SendResetOTP(c context.Context, email string, smtpUsername, smtpPassword string) error
	ResetPassword(c context.Context, email, otpValue, newPassword string) error
}
