package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/CrowderSoup/gozette/config"
	"github.com/CrowderSoup/gozette/micropub"
)

// IndieAuthRes the auth response
type IndieAuthRes struct {
	Me       string `json:"me"`
	ClientID string `json:"client_id"`
	Scope    string `json:"scope"`
	Issue    int    `json:"issued_at"`
	Nonce    int    `json:"nonce"`
}

func checkAccess(token, indieAuthMe, indieAuthTokenURL string) (bool, error) {
	if token == "" {
		return false,
			errors.New("Token string is empty")
	}
	// form the request to check the token
	client := &http.Client{}
	req, err := http.NewRequest("GET", indieAuthTokenURL, nil)
	if err != nil {
		return false,
			errors.New("Error making the request for checking token access")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", token)

	// send the request
	res, err := client.Do(req)
	if err != nil {
		return false,
			errors.New("Error sending the request for checking token access")
	}
	defer res.Body.Close()
	// parse the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false,
			errors.New("Error parsing the response for checking token access")
	}
	var indieAuthRes = new(IndieAuthRes)
	err = json.Unmarshal(body, &indieAuthRes)
	if err != nil {
		return false,
			errors.New("Error parsing the response into json for checking token access " + err.Error())
	}

	// verify results of the response
	if indieAuthRes.Me != indieAuthMe {
		return false,
			errors.New("Me does not match")
	}
	scopes := strings.Fields(indieAuthRes.Scope)
	postPresent := false
	for _, scope := range scopes {
		if scope == "post" || scope == "create" || scope == "update" {
			postPresent = true
			break
		}
	}
	if !postPresent {
		return false,
			errors.New("Post is not present in the scope")
	}
	return true, nil
}

// CheckAuthorization checks that the request is authorized
func CheckAuthorization(config *config.Config, entry *micropub.Entry, headers map[string]string) bool {
	token, ok := headers["authorization"]
	if !ok && len(entry.Token) == 0 { // there is no token provided
		return false
	} else if ok {
		entry.Token = token
	}

	if ok, err := checkAccess(entry.Token, config.IndieAuthMe, config.IndieAuthTokenURL); ok {
		return true
	} else if err != nil {
		return false
	} else {
		return false
	}
}
