package guacapi

import (
	"fmt"

	. "github.com/mdanidl/guac-api/types"
)

//                                 userGroups       memberUsers    daniels-group, permission_ops
func (g *Guac) UpdatePermissions(for_entity_type, to_entity_type, entity string, p []PermissionItem) {
	url := "/api/session/data/mysql/" + for_entity_type + "/" + entity + "/" + to_entity_type
	resp, err := g.Call("PATCH", url, nil, p)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(string(resp))
	}

}

func (g *Guac) UpdateUserPermission
