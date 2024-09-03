package controller

import (
	"loan/config"
	"loan/domain"
	"net/http"

	// "errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanController struct {
	LoanUsecase domain.LoanUsecase
	Env         *config.Env
}

// ApplyLoan applies for a loan

func (lc *LoanController) ApplyLoan(c *gin.Context) {
	var loan domain.LoanRequest
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	claims := c.MustGet("claim").(domain.JwtCustomClaims)
	loan.UserID = claims.UserID
	addedloan,err := lc.LoanUsecase.CreateLoan(c, &loan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, domain.Response{
		Message: "Loan applied successfully",
		Data: addedloan,
	})
}

func (lc *LoanController) GetLoanByID(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	claims := c.MustGet("claim").(domain.JwtCustomClaims)
	userID := claims.UserID
	loan, err := lc.LoanUsecase.GetLoanByID(c, id, userID)
	if loan == nil || err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Err:     err,
			Message: "No Loan found with that ID",
		})
	}
	c.JSON(http.StatusFound, domain.Response{
		Message : "Fetched Successfully",
		Data: loan,
	})
}
