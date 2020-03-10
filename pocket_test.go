package gopocket

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jarcoal/httpmock"
)

const CONSUMER_KEY = "POCKET_CONSUMER_KEY"

func Test_OAuthRequestSuccess(t *testing.T) {
	pocket := Pocket{ ConsumerKey: CONSUMER_KEY }
	redirectUrl := "THE_REDIRECT_URL"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://getpocket.com/v3/oauth/request",
		httpmock.NewStringResponder(200, `{"consumer_key": "`+CONSUMER_KEY+`", "code": "a-fake-code"}`))

	result, err := pocket.OAuthRequest(redirectUrl)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, "getpocket.com", result.Hostname())
	assert.Equal(t, "/auth/authorize", result.Path)
	assert.NotEmpty(t, "a-fake-code", result.Query().Get("request_token"), "Request token missing from redirected URI")
	assert.Equal(t, redirectUrl, result.Query().Get("redirect_uri"), "Redirect url does not match")
}