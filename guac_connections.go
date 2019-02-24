package guacapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RootGuacConnection struct {
	Name              string                        `json:"name"`
	Identifier        string                        `json:"identifier"`
	Type              string                        `json:"type"`
	ActiveConnections int                           `json:"activeConnections"`
	ChildConnections  []GuacConnection              `json:"childConnections"`
	ChildGroups       []GuacConnectionGroups        `json:"childConnectionGroups"`
	Attributes        GuacConnectionGroupAttributes `json:"attributes"`
}
type GuacConnectionGroups struct {
	Name              string                        `json:"name"`
	Identifier        string                        `json:"identifier"`
	ParentIdentifier  string                        `json:"parentIdentifier"`
	Type              string                        `json:"type"`
	ActiveConnections int                           `json:"activeConnections"`
	ChildConnections  []GuacConnection              `json:"childConnections"`
	ChildGroups       []GuacConnectionGroups        `json:"childConnectionGroups"`
	Attributes        GuacConnectionGroupAttributes `json:"attributes"`
}
type GuacConnection struct {
	Name              string                   `json:"name"`
	Identifier        string                   `json:"identifier"`
	ParentIdentifier  string                   `json:"parentIdentifier"`
	Protocol          string                   `json:"protocol"`
	Attributes        GuacConnectionAttributes `json:"attributes"`
	ActiveConnections int                      `json:"activeConnections"`
}
type GuacConnectionGroupAttributes struct {
	MaxConnections        string `json:"max-connections"`
	MaxConnectionsPerUser string `json:"max-connections-per-user"`
	EnableSessionAffinity string `json:"enable-session-affinity"`
}

type GuacConnectionAttributes struct {
	GuacdEncryption       string `json:"guacd-encryption"`
	FailoverOnly          string `json:"failover-only"`
	Weight                string `json:"weight"`
	MaxConnections        string `json:"max-connections"`
	GuacdHostname         string `json:"guacd-hostname"`
	GuacdPort             string `json:"guacd-port"`
	MaxConnectionsPerUser string `json:"max-connections-per-user"`
}

func (g *Guac) GetConnections() (RootGuacConnection, error) {
	req, _ := http.NewRequest("GET", g.URI+"/api/session/data/mysql/connectionGroups/ROOT/tree", nil)

	q := req.URL.Query()
	q.Add("token", g.Token)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return RootGuacConnection{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RootGuacConnection{}, err
	}

	var connresp RootGuacConnection

	err = json.Unmarshal(body, &connresp)
	if err != nil {
		return RootGuacConnection{}, err
	}

	return connresp, err
}
