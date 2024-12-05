package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1

// @Summary Delete user
// @Description Delete a user by email
// @Tags Users
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} schemas.User
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /user/{email} [delete]
func (h *UserHandlerImpl) DeleteUserHandler(ctx *gin.Context) {
	// grab email parameter
	email := ctx.Param("email")
	if email == "" {
		utils.SendError(ctx, http.StatusBadRequest, "email is required")
		return
	}

	// find user
	user, err := h.usecase.GetUser(email)
	if err != nil {
		utils.SendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	// delete user
	err = h.usecase.DeleteUser(email)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "delete-user", user)
}
