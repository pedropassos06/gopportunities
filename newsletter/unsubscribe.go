package newsletter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1

// @Summary Unsubscribe from Newsletter
// @Description unsubscribe from newsletter given an email
// @Tags Newsletter
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param user_email path string true "Email of user to be unsubscribed"
// @Success 200 {object} schemas.NewsletterSubscription
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /newsletter/unsubscribe/{user_email} [put]
func (h *NewsletterHandlerImpl) UnsubscribeHandler(ctx *gin.Context) {
	// get email from path
	email := ctx.Param("user_email")
	if email == "" {
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("user_email", "string").Error())
		return
	}

	// check if email found
	subscription, err := h.Usecase.Find(email)
	if err != nil {
		utils.SendError(ctx, http.StatusNotFound, "email not found in newsletter list")
		return
	}

	// update record in db
	if err := h.Usecase.Unsubscribe(subscription).Error; err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "unable to unsubscribe")
		return
	}

	utils.SendSuccess(ctx, "unsubscribe-handler", subscription)
}
