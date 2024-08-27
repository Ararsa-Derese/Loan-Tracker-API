package route

import (
	"loan/config"
	"loan/database"
	"loan/delivery/controller"

	// "loan/delivery/middleware"
	"loan/domain"
	"loan/repository"
	"loan/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *config.Env, timeout time.Duration, db database.Database, router *gin.RouterGroup) {
	Profilerepo := repository.NewUserRepository(db, domain.CollectionUser)
	Profileusecase := usecase.NewProfileUsecase(Profilerepo, timeout)
	ProfileController := &controller.ProfileController{
		ProfileUsecase: Profileusecase,
		Env:            env,
	}

	router.GET("/profile", ProfileController.GetProfile)
}
