package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	userInfoEmailEndpoint   = "https://www.googleapis.com/auth/userinfo.email"
	userInfoProfileEndpoint = "https://www.googleapis.com/auth/userinfo.profile"
)

func (h *AuthHandlerImpl) GoogleAuthHandler(ctx *gin.Context) {
	//configure google oauth2 settings
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{userInfoProfileEndpoint, userInfoEmailEndpoint},
		Endpoint:     google.Endpoint,
	}

	// generate oauth2 url
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Println(url)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
