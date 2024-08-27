package controller

import (
	"loan/config"
	"loan/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	LoanUsecase domain.LoanUsecase
	Env         *config.Env
}

// DeleteUser deletes a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	claims := c.MustGet("claim").(domain.JwtCustomClaims)
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	existingUser, errr := uc.UserUsecase.GetUserByID(c, objectID)
	if existingUser == nil || errr !=nil{
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
	if len(users) == 0 {
		c.JSON(http.StatusOK, domain.Response{
			Message: "No users found",
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Users fetched successfully",
		Data:    users,
	})
}

func (uc *UserController) GetLoans(c *gin.Context) {
	claims := c.MustGet("claim").(domain.JwtCustomClaims)
	loans, err := uc.LoanUsecase.GetLoans(c, &claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	if len(loans) == 0 {
		c.JSON(http.StatusOK, domain.Response{
			Message: "No loans found",
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Loans fetched successfully",
		Data:    loans,
	})
}

func (uc *UserController) UpdateLoan(c *gin.Context) {
	claims := c.MustGet("claim").(domain.JwtCustomClaims)
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	status := c.Query("status")
	existinloans, errr := uc.LoanUsecase.GetLoanByID(c, objectID, claims.UserID)
	if existinloans == nil || errr !=nil{
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     nil,
			Message: "laon not found",
		})
		return
	}
	err := uc.LoanUsecase.UpdateLoanStatus(c, &claims, status, objectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Loan status updated successfully",
	})
}

func (uc *UserController) DeleteLoan(c *gin.Context){
	claims := c.MustGet("claim").(domain.JwtCustomClaims)
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	existinloans, errr := uc.LoanUsecase.GetLoanByID(c, objectID, claims.UserID)
	if existinloans == nil || errr !=nil{
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     nil,
			Message: "laon not found",
		})
		return
	}
	err := uc.LoanUsecase.DeleteLoan(c, &claims,objectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Message: "Loan deleted successfully",
	})
}