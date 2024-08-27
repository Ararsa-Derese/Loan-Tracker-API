package route

import (
	"loan/config"
	"loan/database"
	"loan/delivery/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *config.Env, timeout time.Duration, db database.Database, gin *gin.Engine) {
	UserRouter := gin.Group("/users")
	adminRouter := gin.Group("/admin")
	loanRouter := gin.Group("/loans")
	// All Public APIs
	NewSignupRouter(env, timeout, db, UserRouter)
	NewLoginRouter(env, timeout, db, UserRouter)
	NewForgotPasswordRouter(env, db, UserRouter)

	protectedRouter := gin.Group("/users")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.AuthMidd)
	adminRouter.Use(middleware.AuthMidd)
	loanRouter.Use(middleware.AuthMidd)
	// All Private APIs
	NewUserRouter(env, timeout, db, adminRouter)
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewLoanRouter(env, timeout, db, loanRouter)

}
