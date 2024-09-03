package route

import (
	"load/config"
	"load/database"
	"load/delivery/controller"

	// "load/delivery/middleware"
	"load/domain"
	"load/repository"
	"load/usecase"
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
