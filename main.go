package main

import (
	"fmt"
	"ride_sharing_application/user"
	"ride_sharing_application/vehicle"
)

//Write tests here only

func main() {
	um := user.NewUserManager()
	vm := vehicle.NewVehicleManager(um)

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

	// Add vehicles
	err = vm.AddVehicle("Rohan", "Swift", "KA-01-12345")
	if err != nil {
		fmt.Println("Error adding vehicle for Rohan:", err)
	}

	err = vm.AddVehicle("Shashank", "Baleno", "TS-04-22345")
	if err != nil {
		fmt.Println("Error adding vehicle for Shashank:", err)
	}

	err = vm.AddVehicle("Rahul", "XUV", "TS-10-22895")
	if err != nil {
		fmt.Println("Error adding vehicle for Rahul:", err)
	}

	err = vm.AddVehicle("Shipra", "Activa", "KA-12-12245")
	if err != nil {
		fmt.Println("Error adding vehicle for Shipra:", err)
	}

	err = vm.AddVehicle("Shipra", "Polo", "KA-05-19845")
	if err != nil {
		fmt.Println("Error adding vehicle for Shipra:", err)
	}

	err = vm.AddVehicle("Armaan", "Polo", "KA-05-19088")
	if err != nil {
		fmt.Println("Error adding vehicle for Shipra:", err)
	}

	err = vm.AddVehicle("Shawn", "WagonR", "KA-05-10005")
	if err != nil {
		fmt.Println("Error adding vehicle for Shipra:", err)
	}

	// // Attempt to add a vehicle for a non-existent user
	err = vm.AddVehicle("UnknownUser", "UnknownModel", "UNKNOWN-12345")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
