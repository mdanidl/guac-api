package guacamole

import (
	"encoding/json"

	"github.com/imdario/mergo"

)

func (g *Guac) CreateConnection(conn *GuacConnection) (GuacConnection, error) {
	ret := GuacConnection{}
	resp, err := g.Call("POST", "/api/session/data/{{ .Datasource }}/connections", nil, conn)
	if err != nil {
		return ret, err
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return GuacConnection{}, err
	}
	return ret, nil
}

func (g *Guac) ReadConnection(conn *GuacConnection) (GuacConnection, error) {
	ret := GuacConnection{}
	retParams := GuacConnectionParameters{}

	connData, err := g.Call("GET", "/api/session/data/{{ .Datasource }}/connections/"+conn.Identifier, nil, nil)
	if err != nil {
		return GuacConnection{}, err
	}
	err = json.Unmarshal(connData, &ret)
	if err != nil {
		return GuacConnection{}, err
	}

	connParams, err := g.Call("GET", "/api/session/data/{{ .Datasource }}/connections/"+conn.Identifier+"/parameters", nil, nil)
	if err != nil {
		return GuacConnection{}, err
	}
	err = json.Unmarshal(connParams, &retParams)
	if err != nil {
		return GuacConnection{}, err
	}
	ret2 := GuacConnection{
		Properties: retParams,
	}
	err = mergo.Merge(&ret, &ret2)
	if err != nil {
		return GuacConnection{}, err
	}
	return ret, nil
}

func (g *Guac) UpdateConnection(conn *GuacConnection) error {
	_, err := g.Call("PUT", "/api/session/data/{{ .Datasource }}/connections/"+conn.Identifier, nil, conn)
	if err != nil {
		return err
	}
	return nil
}

func (g *Guac) DeleteConnection(conn *GuacConnection) error {
	_, err := g.Call("DELETE", "/api/session/data/{{ .Datasource }}/connections/"+conn.Identifier, nil, nil)
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

	ret = append(ret, conn_tree.ChildConnections...)

	flat_conns_from_groups, _, err := flatten(conn_tree.ChildGroups)
	if err != nil {
		return []GuacConnection{}, err
	}
	ret = append(ret, flat_conns_from_groups...)

	return ret, nil
}
