package controller

import (
	"load/config"
	"load/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
	Env            *config.Env
}

// UpdateProfile updates a user's profile
func (pc *ProfileController) GetProfile(c *gin.Context) {
	var profile domain.Profile
	if err := c.ShouldBind(&profile); err != nil {
		c.JSON(http.StatusBadRequest,domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}
	claims := c.MustGet("claim").(domain.JwtCustomClaims)
	id, _ := primitive.ObjectIDFromHex(claims.Id)
	resp, err := pc.ProfileUsecase.GetProfile(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Err:     err,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Profile fetched successfully",
		Data:    resp,
	})
}
