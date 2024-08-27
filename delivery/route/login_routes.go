package route

import (
	"load/config"
	"load/database"
	"load/delivery/controller"
	"load/domain"
	"load/repository"
	"load/usecase"

	// "go/token"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *config.Env, timeout time.Duration, db database.Database, router *gin.RouterGroup) {
	loginRepo := repository.NewUserRepository(db, domain.CollectionUser)
	tokenRepo := repository.NewMongoTokenRepository(db, domain.TokenCollection)

	loginController := &controller.LoginController{

		LoginUsecase: usecase.NewLoginUsecase(loginRepo, tokenRepo, timeout),

		Env: env,
	}

	router.POST("/login", loginController.Login)
	router.POST("/token/refresh", loginController.RefreshTokenHandler)
}
