package newsletter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
	"github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1

// @Summary Unsubscribe from Newsletter
// @Descriptions unsubscribe from newsletter given an email
// @Tags Newsletter
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param user_email path string true "Email of user to be unsubscribed"
// @Success 200 {object} NewsletterUnsubscribeResponse
// @Failure 400 {obect} utils.ErrorResponse
// @Failure 404 {obect} utils.ErrorResponse
// @Failure 500 {obect} utils.ErrorResponse
// @Route /newsletter/unsubscribe/{user_email} [put]
func (nh *NewsletterHandler) UnsubscribeHandler(ctx *gin.Context) {
	// get email from path
	email := ctx.Param("user_email")
	if email == "" {
		utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsRequired("user_email", "string").Error())
		return
	}

	var subscription schemas.NewsletterSubscription

	// check if email found
	if err := nh.DB.First(&subscription, "email = ?", email).Error; err != nil {
		utils.SendError(ctx, http.StatusNotFound, "email not found in newsletter list")
		return
	}

	// set subscribed param as false
	subscription.Subscribed = false

	// update record in db
	if err := nh.DB.Save(&subscription).Error; err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "unable to unsubscribe")
		return
	}

	utils.SendSuccess(ctx, "unsubscribe-handler", subscription)
}
