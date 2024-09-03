package route

import (
	"loan/config"
	"loan/database"
	"loan/delivery/controller"
	"loan/repository"
	"loan/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewForgotPasswordRouter(env *config.Env, db database.Database, router *gin.RouterGroup) {
	otpRepo := repository.NewOTPRepository(db, "otp_collection")
	userRepo := repository.NewUserRepository(db, "users")
	forgotPasswordUsecase := usecase.NewForgotPasswordUsecase(userRepo, otpRepo, time.Minute*15)

	forgotPasswordController := &controller.ForgotPasswordController{
		ForgotPasswordUsecase: forgotPasswordUsecase,
		Env:                   env,
	}

	router.POST("/password-reset", forgotPasswordController.ForgotPassword)
	router.POST("/password-update", forgotPasswordController.ResetPassword)
}
