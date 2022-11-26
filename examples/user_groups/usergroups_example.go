package main

import (
	"fmt"

	guac "github.com/mdanidl/guac-api"
)

func main() {
	// This is just a null declaration, so linter won't complain about the types package
	_ = guac.GuacUserGroup{}

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

	// CREATE UsersGroups
	// new_user_group := GuacUserGroup{
	// 	Identifier: "Daniels-group",
	// }
	// resp, err := gAPI.CreateUserGroup(&new_user_group)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(resp.Identifier)
	// }

	// READ UsersGroup
	// read_user_group := GuacUserGroup{
	// 	Identifier: "Daniels-group",
	// }
	// read_response, err := gAPI.ReadUserGroup(&read_user_group)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(read_response.Identifier)
	// }

	// UPDATE UserGroup

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
