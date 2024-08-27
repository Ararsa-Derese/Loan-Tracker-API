package controller

import (
	"loan/config"
	"loan/domain"
	"loan/internal/tokenutil"
	"loan/internal/userutil"
	"net/http"

	// "errors"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *config.Env
}

// Login authenticates a user and returns tokens
func (lc *LoginController) Login(c *gin.Context) {
	var loginUser domain.AuthLogin
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()
	deviceFingerprint := userutil.GenerateDeviceFingerprint(ipAddress, userAgent)

	user, err := lc.LoginUsecase.AuthenticateUser(c, &loginUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	tkn := domain.Token{
		ID:                primitive.NewObjectID(),
		UserID:            user.ID,
		RefreshToken:      refreshToken,
		ExpiresAt:         time.Now().Add(time.Hour * 24 * time.Duration(lc.Env.RefreshTokenExpiryHour)),
		CreatedAt:         time.Now(),
		DeviceFingerprint: deviceFingerprint,
	}
	err = lc.LoginUsecase.SaveRefreshToken(c, &tkn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	resp := domain.LoginResponse{
		ID:           user.ID,
		AcessToken:   accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Logged in Succefully",
		Data:    resp,
	})
}

func (lc *LoginController) RefreshTokenHandler(c *gin.Context) {

	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	claims, err := tokenutil.VerifyToken(req.RefreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	_, err = lc.LoginUsecase.CheckRefreshToken(c, req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	user := domain.AuthSignup{
		Username: claims.Username,
		Email:    claims.Email,
		UserID:   claims.UserID,
	}
	newaccessToken, err := tokenutil.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "New access Created",
		Data:    map[string]string{"access_token": newaccessToken},
	})
}
