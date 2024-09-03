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

func NewUserRouter(env *config.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	uc := controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
		Env:         env,
	}
	
	group.DELETE("/users/:id", uc.DeleteUser)
	group.GET("/users", uc.GetUsers)

}
