package utils

import (
	"../configs"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

func Auth(code string) (accessToken string, refreshToken string, err error) {
	data := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {configs.CLIENT_ID},
		"client_secret": {configs.CLIENT_SECRET},
		"code":          {code},
		"redirect_uri":  {"https://w0rng.ru"},
	}

	resp, err := http.PostForm(configs.AUTH_URL, data)

	if err != nil {
		return "", "", err
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	if resp.StatusCode != 200 {
		return "", "", errors.New(res["error"].(string))
	}

	return res["access_token"].(string), res["refresh_token"].(string), nil
}

func Reauth(refreshToken string) (accessToken string, newRefreshToken string, err error) {
	data := url.Values{
		"grant_type":     {"refresh_token"},
		"refresh_token ": {refreshToken},
	}

	resp, err := http.PostForm(configs.AUTH_URL, data)
	if err != nil {
		return "", "", err
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	if resp.StatusCode != 200 {
		return "", "", errors.New(res["error"].(string))
	}
	return res["access_token"].(string), res["refresh_token"].(string), nil
}
