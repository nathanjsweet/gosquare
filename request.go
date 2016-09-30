package gosquare

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	_V1_ENDPOINT       = "https://connect.squareup.com/v1/%s"
	_OAUTH_URL         = "https://connect.squareup.com/oauth2/"
	_OAUTH_PERM        = _OAUTH_URL + "authorize?client_id=%s&scope=%s&session=%t"
	_OAUTH_TOKEN       = _OAUTH_URL + "token"
	_OAUTH_RENEW_TOKEN = _OAUTH_URL + "clients/%s/access-token/renew"
)

func v1Request(method, action, token string, reqObj interface{}, result interface{}) error {
	var body io.Reader = nil
	if reqObj != nil {
		bts, err := json.Marshal(reqObj)
		if err != nil {
			return err
		}
		body = bytes.NewReader(bts)
	}
	req, err := http.NewRequest(method, fmt.Sprintf(_V1_ENDPOINT, action), body)
	if err != nil {
		return err
	}
	req.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", token)}
	req.Header["Accept"] = []string{"application/json"}
	if method == "POST" || method == "PUT" {
		req.Header["Content-Type"] = []string{"application/json"}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(result); err != nil {
		return err
	}
	return nil
}

// Generate a url to pass to a user to gain permisson to their account.
// Cf. https://docs.connect.squareup.com/api/oauth/
// "locale", "state" are optional leave them as empty strings to omit them.
// "session", if "false", will require the merchant to sign into their account
// even if they have a valid session in their account. It is likely that you want
// to pass "true".
// Scope should be a space seperated list of permissions, see the above url
// for details on what permissions are available.
// This function will escape all your arguments so don't pass uri-escaped values.
func GeneratePermissionURL(clientId, scope string, session bool, locale, state string) string {
	uri := fmt.Sprintf(_OAUTH_PERM, url.QueryEscape(clientId), url.QueryEscape(scope), session)
	if len(locale) > 0 {
		uri += fmt.Sprintf("&locale=%s", url.QueryEscape(locale))
	}
	if len(state) > 0 {
		uri += fmt.Sprintf("&state=%s", url.QueryEscape(state))
	}
	return uri
}

// Response from GetToken and Renew Token
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   string `json:"expires_at"`
	MerchantId  string `json:"merchant_id"`
}

// Get first token from new merchant's authorization code.
func GetToken(authorizationCode, applicationId, applicationSecret string) (*Token, error) {
	bts, err := json.Marshal(&map[string]string{
		"code":          authorizationCode,
		"client_id":     applicationId,
		"client_secret": applicationSecret,
	})
	if err != nil {
		return nil, err
	}
	return tokenRequest(_OAUTH_TOKEN, bts, applicationSecret)
}

// Renew token from expired token. If the token is older than 30 days this won't work.
func RenewToken(expiredToken, applicationId, applicationSecret string) (*Token, error) {
	bts, err := json.Marshal(&map[string]string{
		"access_token": expiredToken,
	})
	if err != nil {
		return nil, err
	}
	return tokenRequest(fmt.Sprintf(_OAUTH_RENEW_TOKEN, applicationId), bts, applicationSecret)
}

func tokenRequest(url string, body []byte, applicationSecret string) (*Token, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header["Authorization"] = []string{fmt.Sprintf("Client %s", applicationSecret)}
	req.Header["Accept"] = []string{"application/json"}
	req.Header["Content-Type"] = []string{"application/json"}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	token := new(Token)
	if err = dec.Decode(token); err != nil {
		return nil, err
	}
	return token, nil
}

// This method validates that the "X-Square-Signature" header is valid
// and that the webook is, therefore, a valide request from sqaure and not an attack.
// Cf https://docs.connect.squareup.com/api/connect/v1/#validating-notifications
// The first argument should be the url that handles the incoming webhooks
// The second argument is your webhook signature, the third is the body of the request,
// the fourth is the header "X-Square-Signatrue"
func ValidateWebHook(webhookURL, webhookSignatureKey, body, squareSignature string) bool {
	mac := hmac.New(sha1.New, []byte(webhookSignatureKey))
	// Hash writes don't return errors
	mac.Write([]byte(webhookURL + body))
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(expectedMAC, []byte(squareSignature))
}
