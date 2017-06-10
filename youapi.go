package youapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	API_TOKEN_URL = "https://www.googleapis.com/oauth2/v4/token"
)

// Получение токена
func (you *Api) GetToken(d TokenData) (ans GetTokenAns, err error) {
	q := url.Values{}

	q.Add("client_id", d.ClientId)
	q.Add("client_secret", d.ClientSecret)

	if d.Code != "" {
		q.Add("grant_type", "authorization_code")
		q.Add("redirect_uri", d.RedirectUri)
		q.Add("code", d.Code)
	} else {
		q.Add("grant_type", "refresh_token")
	}

	if d.RefreshToken != "" {
		q.Add("refresh_token", d.RefreshToken)
	}

	req, err := http.NewRequest("POST", API_TOKEN_URL, strings.NewReader(q.Encode()))
	if err != nil {
		log.Println("[error]", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println("[error]", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	log.Println(string(content))

	err = json.Unmarshal(content, &ans)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	return
}

func (you *Api) GetVideoInfo(id string) (ans GetVideoInfoAns, err error) {

	q := url.Values{}
	q.Add("part", "snippet,statistics")
	q.Add("key", you.AccessToken)
	q.Add("id", id)

	resp, err := http.Get("https://www.googleapis.com/youtube/v3/videos?" + q.Encode())
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println("[error]", err, string(content))
		return
	}

	err = json.Unmarshal(content, &ans)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	return
}
