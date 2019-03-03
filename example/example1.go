package main

import (
	"fmt"

	guac "github.com/mdanidl/guac-api"
)

func main() {
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

	// Retrieving all connections preserving the group hierarchy

	// conns, err := gAPI.GetAllConnections()
	// if err != nil {
	// 	fmt.Println("Err: ", err)
	// }
	// for _, x := range conns.ChildConnections {
	// 	fmt.Println(x.ParentIdentifier)
	// 	fmt.Println("--", x.Identifier, " :: ", x.Name)
	// }

	// for _, y := range conns.ChildGroups {
	// 	fmt.Println("----", y.Name)
	// 	for _, z := range y.ChildConnections {
	// 		fmt.Println("--------", z.Identifier, " :: ", z.Name)
	// 	}
	// }
	//------------------------

	// Retrieving all connections into a flat slice

	// conns, err := gAPI.GetAllConnectionsFlat()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _, c := range conns {
	// 	fmt.Println("Parent: " + c.ParentIdentifier + " - " + "Name: " + c.Name)
	// }
	//------------------------

	// Adding new connection to the list

	// newconn := guac.GuacConnection{
	// 	Name:             "random2",
	// 	ParentIdentifier: "ROOT",
	// 	Protocol:         "vnc",
	// }
	// _, err = gAPI.AddConnection(&newconn)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//------------------------

	// Deleting a single connection

	// conn_to_delete := guac.GuacConnection{
	// 	Name:       "vmiasd",
	// 	Identifier: "9",
	// }

	// err = gAPI.DeleteConnection(&conn_to_delete)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// } else {
	// 	fmt.Println("Deleted: " + conn_to_delete.Name)
	// }
	//------------------------

	// blah, err := json.Marshal(ok)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(blah))

}
