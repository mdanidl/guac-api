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
	_, err := gAPI.Connect()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(gAPI.Token)
}
