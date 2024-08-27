package route

import (
	"load/config"
	"load/database"
	"load/delivery/controller"
	"load/domain"
	"load/repository"
	"load/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

// Setup sets up the routes for the application

func NewSignupRouter(env *config.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	tr := repository.NewMongoTokenRepository(db, domain.TokenCollection)
	or := repository.NewOTPRepository(db, domain.CollectionOTP)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, tr, or, timeout),
		Env:           env,
	}
	group.POST("/register", sc.Signup)
	group.POST("/verify-email", sc.VerifyOTP)
	
}
