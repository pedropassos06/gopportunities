package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/pedropassos06/gopportunities/helper"
	schemas "github.com/pedropassos06/gopportunities/schemas"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	userInfoURL  = "https://www.googleapis.com/oauth2/v2/userinfo"
	profileScope = "https://www.googleapis.com/auth/userinfo.profile"
	emailScope   = "https://www.googleapis.com/auth/userinfo.email"
)

func (h *Handler) GoogleCallbackHandler(ctx *gin.Context) {
	// load OAuth2 configuration
	config, err := setUpGoogleAuthClient()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error setting up google auth client")
		return
	}

	// retrieve auth code from request
	code := ctx.Query("code")
	if code == "" {
		sendError(ctx, http.StatusBadRequest, "authorization code not found")
		return
	}

	// exchange auth code for access token
	authToken, err := config.Exchange(ctx, code)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// retrieve user info
	client := config.Client(ctx, authToken)
	resp, err := client.Get(userInfoURL)
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

	// generate JWT token
	jwt, err := helper.GenerateJWT(userInfo)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "failed to generate JWT")
		return
	}

	// create token response
	token := &schemas.Token{Token: jwt}

	// send success
	ctx.JSON(http.StatusOK, token)
}

// sets up google auth client configuration
func setUpGoogleAuthClient() (*oauth2.Config, error) {
	clientID := helper.LoadEnvVariable("GOOGLE_CLIENT_ID")
	clientSecret := helper.LoadEnvVariable("GOOGLE_CLIENT_SECRET")
	redirectURL := helper.LoadEnvVariable("GOOGLE_REDIRECT_URL")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{profileScope, emailScope},
		Endpoint:     google.Endpoint,
	}, nil
}