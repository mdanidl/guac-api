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

func (g *Guac) Connect() error {
	resp, err := http.PostForm(g.URI+"/api/tokens",
		url.Values{
			"username": {g.Username},
			"password": {g.Password},
		})
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var tokenresp ConnectResponse

	err = json.Unmarshal(body, &tokenresp)
	if err != nil {
		return err
	}
	g.Token = tokenresp.Authoken
	return nil
}

func (g *Guac) Call(m, u string, xq map[string]string) ([]byte, error) {
	req, _ := http.NewRequest(m, g.URI+u, nil)

	q := req.URL.Query()
	q.Add("token", g.Token)
	if len(xq) > 0 {
		for k, v := range xq {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
