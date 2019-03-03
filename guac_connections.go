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
	Identifier        string                   `json:"identifier,omitempty"`
	ParentIdentifier  string                   `json:"parentIdentifier"`
	Protocol          string                   `json:"protocol"`
	Attributes        GuacConnectionAttributes `json:"attributes"`
	Properties        GuacConnectionParameters `json:"parameters"`
	ActiveConnections int                      `json:"activeConnections,omitempty"`
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
	GuacdHostname         string `json:"guacd-hostname,omitempty"`
	GuacdPort             string `json:"guacd-port"`
	MaxConnectionsPerUser string `json:"max-connections-per-user"`
}
type GuacConnectionParameters struct {
	Port                    string `json:"port"`
	ReadOnly                string `json:"read-only"`
	SwapRedBlue             string `json:"swap-red-blue"`
	Cursor                  string `json:"cursor"`
	ColorDepth              string `json:"color-depth"`
	ClipboardEncoding       string `json:"clipboard-encoding"`
	RecordingExcludeOutput  string `json:"recording-exclude-output"`
	RecordingExcludeMouse   string `json:"recording-exclude-mouse"`
	RecordingIncludeKeys    string `json:"recording-include-keys"`
	CreateRecordingPath     string `json:"create-recording-path"`
	DestPort                string `json:"dest-port"`
	EnableSftp              string `json:"enable-sftp"`
	SftpPort                string `json:"sftp-port"`
	SftpServerAliveInterval string `json:"sftp-server-alive-interval"`
	EnableAudio             string `json:"enable-audio"`
}

func (g *Guac) GetAllConnections() (RootGuacConnection, error) {
	body, err := g.Call("GET", "/api/session/data/mysql/connectionGroups/ROOT/tree", nil, nil)

	var connresp RootGuacConnection

	err = json.Unmarshal(body, &connresp)
	if err != nil {
		return RootGuacConnection{}, err
	}

	return connresp, err
}

func (g *Guac) AddConnection(conn *GuacConnection) (GuacConnection, error) {
	ret := GuacConnection{}
	resp, err := g.Call("POST", "/api/session/data/mysql/connections", nil, conn)
	if err != nil {
		return ret, err
	}
	err = json.Unmarshal(resp, &ret)
	return ret, nil
}

func (g *Guac) DeleteConnection(conn *GuacConnection) error {
	_, err := g.Call("DELETE", "/api/session/data/mysql/connections/"+conn.Identifier, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// TODO
func (g *Guac) GetAllConnectionsFlat() ([]GuacConnection, error) {
	ret := []GuacConnection{}
	conn_tree, err := g.GetAllConnections()
	if err != nil {
		return []GuacConnection{}, err
	}

	for _, root_conns := range conn_tree.ChildConnections {
		ret = append(ret, root_conns)
	}

	flat_conns_from_groups, err := flatten(conn_tree.ChildGroups)
	for _, conns_from_grps := range flat_conns_from_groups {
		ret = append(ret, conns_from_grps)
	}

	return ret, nil
}

func flatten(nested []GuacConnectionGroups) ([]GuacConnection, error) {
	flat_conns := []GuacConnection{}
	for _, groups := range nested {
		if len(groups.ChildGroups) > 0 {
			conns, err := flatten(groups.ChildGroups)
			if err != nil {
				return []GuacConnection{}, err
			}
			for _, c := range conns {
				flat_conns = append(flat_conns, c)
			}
		}
		if len(groups.ChildConnections) > 0 {
			for _, c := range groups.ChildConnections {
				flat_conns = append(flat_conns, c)
			}
		}
	}
	return flat_conns, nil
}

func (g *Guac) GetConnection(conn GuacConnection) (GuacConnection, error) {
	return GuacConnection{}, nil
}
