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

func (g *Guac) ReadConnection(conn *GuacConnection) (GuacConnection, error) {
	ret := GuacConnection{}
	resp, err := g.Call("GET", "/api/session/data/mysql/connections/"+conn.Identifier, nil, nil)
	if err != nil {
		return GuacConnection{}, err
	}
	err = json.Unmarshal(resp, &ret)
	return ret, nil
}

func (g *Guac) UpdateConnection(conn *GuacConnection) error {
	_, err := g.Call("PUT", "/api/session/data/mysql/connections", nil, conn)
	if err != nil {
		return err
	}
	return nil
}

func (g *Guac) DeleteConnection(conn *GuacConnection) error {
	_, err := g.Call("DELETE", "/api/session/data/mysql/connections/"+conn.Identifier, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (g *Guac) ListConnections() ([]GuacConnection, error) {
	ret := []GuacConnection{}
	conn_tree, err := g.GetConnectionTree()
	if err != nil {
		return []GuacConnection{}, err
	}

	for _, root_conns := range conn_tree.ChildConnections {
		ret = append(ret, root_conns)
	}

	flat_conns_from_groups, _, err := flatten(conn_tree.ChildGroups)
	for _, conns_from_grps := range flat_conns_from_groups {
		ret = append(ret, conns_from_grps)
	}

	return ret, nil
}
