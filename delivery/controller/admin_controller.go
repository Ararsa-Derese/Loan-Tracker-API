package controller

import (
	"load/config"
	"load/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Env         *config.Env
}

// DeleteUser deletes a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	claims := c.MustGet("claim").(domain.JwtCustomClaims)
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	existingUser, _ := uc.UserUsecase.GetUserByID(c, objectID)
	if existingUser == nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     nil,
			Message: "User not found",
		})
		return
	}

	err := uc.UserUsecase.DeleteUser(c, objectID, &claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Message: "User deleted successfully",
	})
}



// GetUsers gets all users
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.UserUsecase.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Message: "Users fetched successfully",
		Data:    users,
	})
}




