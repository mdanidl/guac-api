package main

import (
	"fmt"

	guac "github.com/mdanidl/guac-api"
)

func main() {
	// Keeping the types import in
	_ = guac.GuacConnection{}

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

	// CREATE ConnectionGroup

	// newgrp := GuacConnectionGroup{
	// 	Name:             "newConnectionGroup3",
	// 	Type:             "ORGANIZATIONAL",
	// 	ParentIdentifier: "4",
	// }
	// _, err = gAPI.CreateConnectionGroup(&newgrp)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//------------------------

	// READ ConnectionGroup

	// conn := GuacConnectionGroup{
	// 	Identifier: "4",
	// }
	// conn_grp, err := gAPI.ReadConnectionGroup(&conn)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(conn_grp.Identifier + " :: " + conn_grp.Name)
	// }
	//------------------------

	// UPDATE ConnectionGroup

	// conngrp_upd := GuacConnectionGroup{
	// 	Identifier:       "4",
	// 	Name:             "imgood",
	// 	Type:             "ORGANIZATIONAL",
	// 	ParentIdentifier: "ROOT",
	// }
	// err = gAPI.UpdateConnectionGroup(&conngrp_upd)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Updated, reading back")
	// 	updated_grp, err := gAPI.ReadConnectionGroup(&conngrp_upd)
	// 	if err != nil {
	// 		fmt.Println("Readback failed: ", err)
	// 	} else {
	// 		fmt.Println(updated_grp.Identifier + " :: " + updated_grp.Name)
	// 	}
	// }

	//------------------------

	// DELETE ConnectionGroup

	// conngrp_to_delete := GuacConnectionGroup{
	// 	Identifier: "5",
	// }

	// err = gAPI.DeleteConnectionGroup(&conngrp_to_delete)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// } else {
	// 	fmt.Println("Deleted: " + conngrp_to_delete.Identifier)
	// }
	//------------------------

	// LIST ConnectionGroups

	// resp, err := gAPI.ListConnectionGroups()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, y := range resp {
	// 		fmt.Println(y.Identifier, " :: ", y.Name)
	// 	}
	// }

}
