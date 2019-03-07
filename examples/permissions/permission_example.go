package main

import (
	"fmt"

	guac "github.com/mdanidl/guac-api"
	. "github.com/mdanidl/guac-api/types"
)

func main() {
	// This is just a null declaration, so linter won't complain about the types package
	_ = GuacUserGroup{}

	// Initialise the Connection to guacamole using the built-in authentication mechanism
	gAPI := guac.Guac{
		URI:      "http://localhost:8088/guacamole",
		Username: "guacadmin",
		Password: "guacadmin",
	}
	err := gAPI.Connect()
	if err != nil {
		fmt.Println(err)
	}

	// UPDATE Permissions

	perm := []PermissionItem{}
	perm = append(perm,PermissionItem{
		Op: "add",
		Path: "/connections/1",
		Value: "READ"
	})

	// update_user_group := GuacUserGroup{
	// 	Identifier: "Daniels-group",
	// 	Attributes: GuacUserGroupAttributes{
	// 		Disabled: "true",
	// 	},
	// }
	// err = gAPI.UpdateUserGroup(&update_user_group)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Updated. Re-reading info: ")
	// 	update_reread_resp, err := gAPI.ReadUserGroup(&update_user_group)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(update_reread_resp.Attributes.Disabled)
	// 	}
	// }

	// DELETE UserGroup

	// delete_user_group := GuacUserGroup{
	// 	Identifier: "Daniels-group",
	// }

	// err = gAPI.DeleteUserGroup(&delete_user_group)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Ok")
	// }

	// LIST User Groups
	// groups_resp, err := gAPI.ListUserGroups()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, g := range groups_resp {
	// 		fmt.Println(g.Identifier)
	// 	}
	// }
}
