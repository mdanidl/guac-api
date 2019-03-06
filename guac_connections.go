package guacapi

import (
	"encoding/json"

	. "github.com/mdanidl/guac-api/types"
)

func (g *Guac) CreateConnection(conn *GuacConnection) (GuacConnection, error) {
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

func (g *Guac) ReadConnection(conn *GuacConnection) (GuacConnection, error) {
	ret := GuacConnection{}
	resp, err := g.Call("GET", "/api/session/data/mysql/connections/"+conn.Identifier, nil, nil)
	if err != nil {
		return GuacConnection{}, err
	}
	err = json.Unmarshal(resp, &ret)
	return ret, nil
}
