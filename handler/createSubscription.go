package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropassos06/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Subscribe to newsletter
// @Description Subscribes user to newsletter
// @Accept json
// @Produce json
// @Tags Newsletter
// @Param request body NewsletterSubscriptionRequest true "Newsletter subscription details"
// @Success 200 {object} NewsletterSubscriptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /newsletter/subscribe [post]
func (h *Handler) SubscribeHandler(ctx *gin.Context) {
	var request NewsletterSubscriptionRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		sendError(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	subscription := schemas.NewsletterSubscription{
		UserID:     request.UserID,
		Email:      request.Email,
		Subscribed: true,
	}

	if err := h.DB.Create(&subscription).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to subscribe")
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

	sendSuccess(ctx, "subscribe-handler", response)
}
