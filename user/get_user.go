package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1

// @Summary Get User
// @Description Returns a user based on email
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id query string true "User email"
// @Success 200 {object} schemas.User
// @Failure 400 {object} utils.ErrorResponse "email parameter is missing"
// @Failure 404 {object} utils.ErrorResponse "User not found"
// @Router /user [get]
func (h *UserHandlerImpl) GetUserHandler(ctx *gin.Context) {
	// grab email parameter
	email := ctx.Query("email")
	if email == "" {
		utils.SendError(ctx, http.StatusBadRequest, "email is required")
		return
	}

	// find user
	user, err := h.Usecase.GetUser(email)
	if err != nil {
		utils.SendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	utils.SendSuccess(ctx, "get-user", user)
}
