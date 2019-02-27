package guacapi

import (
	"encoding/json"
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
	body, err := g.Call("GET", "/api/session/data/mysql/connectionGroups/ROOT/tree", map[string]string{})

	var connresp RootGuacConnection

	err = json.Unmarshal(body, &connresp)
	if err != nil {
		return RootGuacConnection{}, err
	}

	return connresp, err
}

// TODO
func (g *Guac) GetConnectionsFlat() ([]GuacConnection, error) {
	return []GuacConnection{}, nil
}

func (g *Guac) AddConnection(conn GuacConnection) (GuacConnection, error) {
	return GuacConnection{}, nil
}
func (g *Guac) DeleteConnection(conn GuacConnection) (GuacConnection, error) {
	return GuacConnection{}, nil
}
func (g *Guac) GetConnection(conn GuacConnection) (GuacConnection, error) {
	return GuacConnection{}, nil
}
