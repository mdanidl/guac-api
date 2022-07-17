package guacamole

import (
	"encoding/json"
)

func (g *Guac) SendUserConnectionPermissionChanges(user string, p []GuacPermissionItem) error {
	url := "/api/session/data/{{ .Datasource }}/users/" + user + "/permissions"
	_, err := g.Call("PATCH", url, nil, p)
	if err != nil {
		return err
	}
	return nil
}

func (g *Guac) GetUserConnectionPermissions(user string) (GuacPermissionData, error) {
	ret := GuacPermissionData{}
	url := "/api/session/data/{{ .Datasource }}/users/" + user + "/permissions"
	resp, err := g.Call("GET", url, nil, nil)
	if err != nil {
		return GuacPermissionData{}, err
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return GuacPermissionData{}, err
	}

	return ret, nil
}
