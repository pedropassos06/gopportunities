package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func (h *Handler) GoogleCallbackHandler(ctx *gin.Context) {
	// configure google oauth2
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	// retrieve auth code from request
	code := ctx.Query("code")
	if code == "" {
		sendError(ctx, http.StatusBadRequest, "authorization code not found")
		return
	}

	// exchange auth code for access token
	token, err := config.Exchange(ctx, code)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// retrieve user info
	client := config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	var userInfo map[string]interface{}
	if err := json.Unmarshal(data, &userInfo); err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// send success
	sendSuccess(ctx, "auth-endpoint", userInfo)

}
