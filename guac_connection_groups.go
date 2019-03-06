package guacapi

import (
	"encoding/json"

	. "github.com/mdanidl/guac-api/types"
)

func (g *Guac) GetConnectionTree() (GuacConnectionGroup, error) {
	body, err := g.Call("GET", "/api/session/data/mysql/connectionGroups/ROOT/tree", nil, nil)

	var connresp GuacConnectionGroup

	err = json.Unmarshal(body, &connresp)
	if err != nil {
		return GuacConnectionGroup{}, err
	}

	return connresp, err
}

func (g *Guac) GetAllConnectionsFlat() ([]GuacConnection, error) {
	ret := []GuacConnection{}
	conn_tree, err := g.GetConnectionTree()
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

func flatten(nested []GuacConnectionGroup) ([]GuacConnection, error) {
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
