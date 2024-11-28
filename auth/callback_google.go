package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	schemas "github.com/pedropassos06/gopportunities/schemas"
	utils "github.com/pedropassos06/gopportunities/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	userInfoURL  = "https://www.googleapis.com/oauth2/v2/userinfo"
	profileScope = "https://www.googleapis.com/auth/userinfo.profile"
	emailScope   = "https://www.googleapis.com/auth/userinfo.email"
)

func (h *AuthHandler) GoogleCallbackHandler(ctx *gin.Context) {
	// load OAuth2 configuration
	config, err := setUpGoogleAuthClient()
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error setting up google auth client")
		return
	}

	// retrieve auth code from request
	code := ctx.Query("code")
	if code == "" {
		utils.SendError(ctx, http.StatusBadRequest, "authorization code not found")
		return
	}

	// exchange auth code for access token
	authToken, err := config.Exchange(ctx, code)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// retrieve user info
	client := config.Client(ctx, authToken)
	resp, err := client.Get(userInfoURL)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	var userInfo map[string]interface{}
	if err := json.Unmarshal(data, &userInfo); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// generate JWT token
	jwt, err := utils.GenerateJWT(userInfo)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "failed to generate JWT")
		return
	}

	// create token response
	token := &schemas.Token{Token: jwt}

	// send success
	ctx.JSON(http.StatusOK, token)
}

// sets up google auth client configuration
func setUpGoogleAuthClient() (*oauth2.Config, error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := os.Getenv("GOOGLE_REDIRECT_URL")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{profileScope, emailScope},
		Endpoint:     google.Endpoint,
	}, nil
}
