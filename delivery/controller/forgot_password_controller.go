package controller

import (
	"context"
	"load/config"
	"load/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordController struct {
	ForgotPasswordUsecase domain.ForgotPasswordUsecase
	Env                   *config.Env
}

func (fpc *ForgotPasswordController) ForgotPassword(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	// Pass SMTP credentials directly to the usecase
	smtpUsername := fpc.Env.SMTPUsername // Replace with actual SMTP username
	smtpPassword := fpc.Env.SMTPPassword // Replace with actual SMTP password
	err := fpc.ForgotPasswordUsecase.SendResetOTP(context.Background(), request.Email, smtpUsername, smtpPassword)
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

func (fpc *ForgotPasswordController) ResetPassword(c *gin.Context) {
	var request struct {
		Email       string `json:"email"`
		OTPValue    string `json:"otp_value"`
		NewPassword string `json:"new_password"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})

		return
	}

	err := fpc.ForgotPasswordUsecase.ResetPassword(context.Background(), request.Email, request.OTPValue, request.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Password reset successfully",
	})
}
