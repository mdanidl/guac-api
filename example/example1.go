package main

import (
	"fmt"

	guac "github.com/mdanidl/guac-api"
)

func main() {
	gAPI := guac.Guac{
		URI:      "http://localhost:8088/guacamole",
		Username: "guacadmin",
		Password: "guacadmin",
	}
	err := gAPI.Connect()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(gAPI.Token)

	conns, err := gAPI.GetConnections()
	if err != nil {
		fmt.Println("Err: ", err)
	}
	for _, x := range conns.ChildConnections {
		fmt.Println(x.ParentIdentifier)
		fmt.Println("--", x.Identifier, " :: ", x.Name)
	}

	for _, y := range conns.ChildGroups {
		fmt.Println("----", y.Name)
		for _, z := range y.ChildConnections {
			fmt.Println("--------", z.Identifier, " :: ", z.Name)
		}
	}
}
