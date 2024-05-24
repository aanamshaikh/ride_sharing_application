package main

import (
	"fmt"
	"ride_sharing_application/user"
)

//Write tests here only

func main() {
	um := user.NewUserManager()

	um.AddUser("Rohan", "M", 24)
	um.AddUser("Shashank", "M", 29)
	um.AddUser("Nandini", "F", 29)
	um.AddUser("Shipra", "F", 34)
	um.AddUser("Gaurav", "M", 24)
	um.AddUser("Rahul", "M", 39)
	um.AddUser("Armaan", "M", 24)
	um.AddUser("Shawn", "M", 39)
	um.AddUser("Tina", "F", 29)

	// Retrieve and print a user
	user, err := um.GetUser("Rohan")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v\n", user)
	}

}
