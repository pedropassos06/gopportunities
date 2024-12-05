package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
	"github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1

// @Summary Update user
// @Description Update a user
// @Tags Users
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} schemas.User
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /user/{email} [put]
func (h *UserHandlerImpl) UpdateUserHandler(ctx *gin.Context) {
	// grab email parameter
	var user schemas.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "invalid user body")
		return
	}

	// update user
	err := h.Usecase.UpdateUser(user)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "update-user", user)
}
