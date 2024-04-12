package auth

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
)

type GoogleTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GoogleUserBasicInfo struct {
	Picture         string `json:"picture"`
	IsEmailVerified bool   `json:"verified_email"`
	Id              string `json:"id"`
	Email           string `json:"email"`
}

func getGoogleUserInfoFromCallback(r *http.Request) (GoogleUserBasicInfo, error) {
	tokenResponse, err := getGoogleAccessToken(r)
	if err != nil {
		return GoogleUserBasicInfo{}, err
	}

	bearer := "Bearer " + tokenResponse.AccessToken
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/userinfo/v2/me", nil)
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GoogleUserBasicInfo{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	var userInfo GoogleUserBasicInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return GoogleUserBasicInfo{}, err
	}

	return userInfo, nil

}

func getGoogleAccessToken(r *http.Request) (GoogleTokenResponse, error) {
	queryParams := r.URL.Query()
	data := url.Values{
		"code":          {queryParams.Get("code")},
		"client_id":     {os.Getenv("GOOGLE_CLIENT_ID")},
		"client_secret": {os.Getenv("GOOGLE_CLIENT_SECRET")},
		"redirect_uri":  {"http://" + r.Host + "/auth/o/google/callback/"},
		"grant_type":    {"authorization_code"},
	}

	resp, err := http.PostForm("https://oauth2.googleapis.com/token", data)
	if err != nil {
		return GoogleTokenResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	var res GoogleTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return GoogleTokenResponse{}, errors.New("Error while decoding response from google: " + err.Error())
	}
	return res, nil
}
