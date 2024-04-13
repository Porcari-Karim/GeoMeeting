package auth

import (
	"fmt"
	"github.com/Porcari-Karim/GeoMeeting/internal/authentication"
	"net/http"
	"os"
)

func googleOAuthHandler(w http.ResponseWriter, r *http.Request) {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	redirectURI := "http://" + r.Host + "/auth/o/google/callback/"
	scope := "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile"
	authRedirectURL := fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?response_type=code&scope=%s&redirect_uri=%s&client_id=%s", scope, redirectURI, clientId)
	http.Redirect(w, r, authRedirectURL, http.StatusTemporaryRedirect)
}

func googleOAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	userInfo, err := authentication.GetGoogleUserInfoFromCallback(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userInfo.Email))
}
