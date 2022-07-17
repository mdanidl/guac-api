package guacamole

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

)

type Guac struct {
	URI      string
	Username string
	Password string
	Token    string
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
	g.Token = tokenresp.AuthToken
	return nil
}

func (g *Guac) RefreshToken() error {
	resp, err := http.PostForm(g.URI+"/api/tokens",
		url.Values{
			"token": {g.Token},
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
	g.Token = tokenresp.AuthToken
	return nil
}

func (g *Guac) Call(m, u string, xq map[string]string, b interface{}) ([]byte, error) {
	err := g.RefreshToken()
	if err != nil {
		return nil, err
	}

	body_byte, err := json.Marshal(b)
	if err != nil {
		return []byte{}, err
	}
	// fmt.Println("Request: ", string(body_byte))
	body_reader := bytes.NewReader(body_byte)
	req, err := http.NewRequest(m, g.URI+u, body_reader)
	if err != nil {
		return []byte{}, err
	}
	// http headers

	req.Header.Set("Content-Type", "application/json")

	// URL query params
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
	if !(resp.StatusCode == 200 || resp.StatusCode == 204) {
		// fmt.Println(resp.StatusCode)
		errbody, _ := ioutil.ReadAll(resp.Body)
		errmsg := ErrorResponse{}
		umerr := json.Unmarshal(errbody, &errmsg)
		if umerr != nil {
			return nil, errors.New("JSON Unmarshal error")
		}
		msg := strconv.Itoa(resp.StatusCode) + "-" + errmsg.Type + ": " + errmsg.Message
		e := errors.New(msg)
		return nil, e
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
