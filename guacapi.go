package guacapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Guac blabla
type Guac struct {
	URI      string
	Username string
	Password string
	Token    string
}
type ConnectResponse struct {
	Authoken             string   `json:"authToken"`
	Username             string   `json:"username"`
	Datasource           string   `json:"dataSource"`
	Availabledatasources []string `json:"availableDataSource"`
}

func (g *Guac) Connect() (string, error) {
	resp, err := http.PostForm(g.URI+"/api/tokens",
		url.Values{
			"username": {g.Username},
			"password": {g.Password},
		})
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var tokenresp ConnectResponse

	err = json.Unmarshal(body, &tokenresp)
	if err != nil {
		return "", err
	}
	g.Token = tokenresp.Authoken
	// fmt.Println(g.Token)
	return "ok", nil
}
