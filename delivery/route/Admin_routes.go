package route

import (
	"loan/config"
	"loan/database"
	"loan/delivery/controller"
	"loan/domain"
	"loan/repository"
	"loan/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

// Setup sets up the routes for the application

func NewUserRouter(env *config.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lr := repository.NewLoanRepository(db, domain.CollectionLoan)
	uc := controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
		LoanUsecase: usecase.NewLoanUsecase(lr, timeout),
		Env:         env,
	}

	group.DELETE("/users/:id", uc.DeleteUser)
	group.GET("/users", uc.GetUsers)
	group.GET("/loans", uc.GetLoans)
	group.PATCH("/loans/:id",uc.UpdateLoan)
	group.DELETE("/loans/:id",uc.DeleteLoan)

}
