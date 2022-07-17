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

	// READ User ConnectionPermissions

	resp, err := gAPI.GetUserConnectionPermissions("tf-test-1")
	if err != nil {
		fmt.Println("fuck", err)
	}
	fmt.Println(resp)

	// UPDATE Permissions

	// items := []GuacPermissionItem{}
	// items = append(items, GuacPermissionItem{
	// 	Op:    "add",
	// 	Path:  "/connectionPermissions/19",
	// 	Value: "READ",
	// })
	// _ = gAPI.SendUserConnectionPermissionChanges("username", items)

}
