package gopocket

import (
	"net/http"
	"net/url"
)

type Pocket struct {
	ConsumerKey string
}

const (
	POCKET_BASE_PATH = "https://getpocket.com"
	V3_OAUTH_REQUEST = POCKET_BASE_PATH + "/v3/oauth/request"
	AUTH_AUTHORIZE = POCKET_BASE_PATH + "/auth/authorize"
)


func (p Pocket) OAuthRequest(redirectUri string) (url.URL, error) {
	requestBody := map[string]interface{} {
		"consumer_key": p.ConsumerKey,
		"redirect_uri": redirectUri,
	}

	pocketRequest, err := buildPocketRequest(V3_OAUTH_REQUEST, requestBody)
	if err != nil {
		return url.URL{}, err
	}

	oauthRequestResponse, err := http.DefaultClient.Do(pocketRequest)
	if err != nil {
		return url.URL{}, err
	}

	pocketOAuthReqResponse := struct {
		Code string `json:"code"`
	}{}

	err = readPocketResponse(oauthRequestResponse, 200, &pocketOAuthReqResponse)
	if err != nil {
		return url.URL{}, err
	}

	redirectUrlWithCode, err := url.Parse(redirectUri)
	if err != nil {
		return url.URL{}, err
	}

	redirectUrlWithCode.Query().Add("code", pocketOAuthReqResponse.Code)
	
	pocketUrl, err := url.Parse(AUTH_AUTHORIZE)
	q := pocketUrl.Query()
	q.Add("request_token", pocketOAuthReqResponse.Code)
	q.Add("redirect_uri", url.QueryEscape(redirectUrlWithCode.String()))
	pocketUrl.RawQuery = q.Encode()

	return *pocketUrl, err
}