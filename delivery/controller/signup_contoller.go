package controller

import (
	"errors"
	"load/config"
	"load/domain"
	"load/internal/userutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *config.Env
}

// Signup creates a new user
func (sc *SignupController) Signup(c *gin.Context) {
	var user domain.AuthSignup
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()
	deviceFingerprint := userutil.GenerateDeviceFingerprint(ipAddress, userAgent)

	returnedUser, _ := sc.SignupUsecase.GetUserByEmail(c, user.Email)
	if returnedUser != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     errors.New("ivalid Request"),
			Message: "Email already exists",
		})
		return
	}
	returnedUser, _ = sc.SignupUsecase.GetUserByUsername(c, user.Username)
	if returnedUser != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     errors.New("ivalid Request"),
			Message: "Username already exists",
		})
		return
	}
	err := sc.SignupUsecase.SendOTP(c, &user, sc.Env.SMTPUsername, sc.Env.SMTPPassword, deviceFingerprint)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Message: "OTP sent successfully",
	})

}

// VerifyOTP verifies the OTP
func (sc *SignupController) VerifyOTP(c *gin.Context) {
	var otp domain.OTPRequest
	email := c.Query("email")
	otpValue := c.Query("otp")
	otp.Email = email
	otp.Value = otpValue
	otpresponse, err := sc.SignupUsecase.VerifyOTP(c, &otp)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	user := domain.AuthSignup{
		Username: otpresponse.Username,
		Email:    otpresponse.Email,
		Password: otpresponse.Password,
		Role:     "user",
	}
	sc.Register(c, user)
}
func (sc *SignupController) Register(c *gin.Context, user domain.AuthSignup) {
	userID, err := sc.SignupUsecase.RegisterUser(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	sc.Token(c, user, *userID)
}
func (sc *SignupController) Token(c *gin.Context, user domain.AuthSignup, userID primitive.ObjectID) {
	user.UserID = userID
	Accesstoken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	RefreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),

	})
		return
	}
	err = sc.SignupUsecase.SaveRefreshToken(c, RefreshToken, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
	})
		return
	}
	resp := domain.SignUpResponse{
		ID:           userID,
		AcessToken:   Accesstoken,
		RefreshToken: RefreshToken,
	}
	c.JSON(http.StatusOK, domain.Response{
		Message: "User created successfully",
		Data:    resp,
	})
}
