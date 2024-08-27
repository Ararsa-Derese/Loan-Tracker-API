package route

import (
	"loan/config"
	"loan/database"
	"loan/delivery/controller"
	"loan/domain"
	"loan/repository"
	"loan/usecase"

	// "go/token"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoanRouter(env *config.Env, timeout time.Duration, db database.Database, router *gin.RouterGroup) {
	LoanRepo := repository.NewLoanRepository(db, domain.CollectionLoan)

	LoanController := &controller.LoanController{
		LoanUsecase: usecase.NewLoanUsecase(LoanRepo,  timeout),
		Env: env,
	}
	router.POST("/", LoanController.ApplyLoan)
	router.GET("/:id",LoanController.GetLoanByID)
}