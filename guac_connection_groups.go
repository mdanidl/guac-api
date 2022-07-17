package guacamole

import (
	"encoding/json"

)

func (g *Guac) GetConnectionTree() (GuacConnectionGroup, error) {
	body, err := g.Call("GET", "/api/session/data/mysql/connectionGroups/ROOT/tree", nil, nil)
	if err != nil {
		return GuacConnectionGroup{}, err
	}

	var connresp GuacConnectionGroup

	err = json.Unmarshal(body, &connresp)
	if err != nil {
		return GuacConnectionGroup{}, err
	}

	return connresp, err
}

func flatten(nested []GuacConnectionGroup) ([]GuacConnection, []GuacConnectionGroup, error) {
	flat_conns := []GuacConnection{}
	flat_grps := []GuacConnectionGroup{}
	for _, groups := range nested {
		flat_grps = append(flat_grps, groups)
		if len(groups.ChildGroups) > 0 {
			conns, subgrps, err := flatten(groups.ChildGroups)
			if err != nil {
				return nil, nil, err
			}
			flat_conns = append(flat_conns, conns...)
			flat_grps = append(flat_grps, subgrps...)
		}
		if len(groups.ChildConnections) > 0 {
			flat_conns = append(flat_conns, groups.ChildConnections...)
		}
	}
	return flat_conns, flat_grps, nil
}

func (g *Guac) CreateConnectionGroup(conn *GuacConnectionGroup) (GuacConnectionGroup, error) {
	ret := GuacConnectionGroup{}
	resp, err := g.Call("POST", "/api/session/data/mysql/connectionGroups", nil, conn)
	if err != nil {
		return GuacConnectionGroup{}, err
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return GuacConnectionGroup{}, err
	} else {
		return ret, nil
	}
}

func (g *Guac) ReadConnectionGroup(conn *GuacConnectionGroup) (GuacConnectionGroup, error) {
	ret := GuacConnectionGroup{}
	resp, err := g.Call("GET", "/api/session/data/mysql/connectionGroups/"+conn.Identifier, nil, nil)
	if err != nil {
		return GuacConnectionGroup{}, err
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return GuacConnectionGroup{}, err
	} else {
		return ret, nil
	}

}

func (g *Guac) UpdateConnectionGroup(conn *GuacConnectionGroup) error {
	_, err := g.Call("PUT", "/api/session/data/mysql/connectionGroups/"+conn.Identifier, nil, conn)
	if err != nil {
		return err
	} else {
		return nil
	}

}

func (g *Guac) DeleteConnectionGroup(conn *GuacConnectionGroup) error {
	_, err := g.Call("DELETE", "/api/session/data/mysql/connectionGroups/"+conn.Identifier, nil, conn)
	if err != nil {
		return err
	} else {
		return nil
	}

}

func (g *Guac) ListConnectionGroups() ([]GuacConnectionGroup, error) {
	ret := []GuacConnectionGroup{}
	conn_tree, err := g.GetConnectionTree()
	if err != nil {
		return []GuacConnectionGroup{}, err
	}

	_, flattened_grps, err := flatten([]GuacConnectionGroup{conn_tree})
	if err != nil {
		return []GuacConnectionGroup{}, err
	}
	ret = append(ret, flattened_grps...)

	return ret, nil
}
