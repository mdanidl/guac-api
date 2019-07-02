package main

import (
	"fmt"

	guac "github.com/mdanidl/guac-api"
	. "github.com/mdanidl/guac-api/types"
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

	// CREATE Users
	// new_user := GuacUser{
	// 	Username: "Daniel3",
	// 	Password: "daniel3",
	// }
	// resp, err := gAPI.CreateUser(&new_user)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(resp.Username)
	// }

	// READ Users
	// read_user := GuacUser{
	// 	Username: "Daniel2",
	// }
	// read_response, err := gAPI.ReadUser(&read_user)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(read_response.Attributes.GuacFullName)
	// }

	// UPDATE User

	// update_user := GuacUser{
	// 	Username: "Daniel2",
	// 	Attributes: GuacUserAttributes{
	// 		GuacFullName: "Daniel Meszaros",
	// 	},
	// }
	// err = gAPI.UpdateUser(&update_user)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Updated. Re-reading info: ")
	// 	update_reread_resp, err := gAPI.ReadUser(&update_user)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(update_reread_resp.Attributes.GuacFullName)
	// 	}
	// }

	// DELETE User

	delete_user := GuacUser{
		Username: "Daniel2",
	}

	err = gAPI.DeleteUser(&delete_user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ok")
	}

	// LIST Users
	// users_resp, err := gAPI.ListUsers()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, u := range users_resp {
	// 		fmt.Println(u.Username)
	// 	}
	// }
}
