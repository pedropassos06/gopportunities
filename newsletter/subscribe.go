package newsletter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
	utils "github.com/pedropassos06/gopportunities/utils"
)

// @BasePath /api/v1

// @Summary Subscribe to newsletter
// @Description Subscribes user to newsletter
// @Accept json
// @Produce json
// @Tags Newsletter
// @Param Authorization header string true "Bearer Token"
// @Param request body NewsletterSubscriptionRequest true "Newsletter subscription details"
// @Success 200 {object} NewsletterSubscriptionResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /newsletter/subscribe [post]
func (h *NewsletterHandler) SubscribeHandler(ctx *gin.Context) {
	var request NewsletterSubscriptionRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := request.Validate(); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	subscription := schemas.NewsletterSubscription{
		UserID:     request.UserID,
		Email:      request.Email,
		Subscribed: true,
	}

	if err := h.DB.Create(&subscription).Error; err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "Failed to subscribe")
		return
	}

	response := schemas.NewsletterSubscriptionResponse{
		ID:         subscription.ID,
		CreatedAt:  subscription.CreatedAt,
		UpdatedAt:  subscription.UpdatedAt,
		UserID:     subscription.UserID,
		Email:      subscription.Email,
		Subscribed: subscription.Subscribed,
	}

	utils.SendSuccess(ctx, "subscribe-handler", response)
}
