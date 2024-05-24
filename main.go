package main

import (
	"fmt"
	"ride_sharing_application/ride"
	"ride_sharing_application/user"
	"ride_sharing_application/vehicle"
)

//Write tests here only

func main() {
	um := user.NewUserManager()
	vm := vehicle.NewVehicleManager(um)
	rm := ride.NewRideManager(um, vm)

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

	// Offer rides
	err = rm.OfferRide("Rohan", "Hyderabad", "Bangalore", "Swift", "KA-01-12345", 1)
	if err != nil {
		fmt.Println("Error offering ride for Rohan:", err)
	}

	err = rm.OfferRide("Shipra", "Bangalore", "Mysore", "Activa", "KA-12-12245", 1)
	if err != nil {
		fmt.Println("Error offering ride for Shipra (Activa):", err)
	}

	err = rm.OfferRide("Shipra", "Bangalore", "Mysore", "Polo", "KA-05-19845", 2)
	if err != nil {
		fmt.Println("Error offering ride for Shipra (Polo):", err)
	}

	err = rm.OfferRide("Shashank", "Hyderabad", "Bangalore", "Baleno", "TS-04-22345", 2)
	if err != nil {
		fmt.Println("Error offering ride for Shashank:", err)
	}

	err = rm.OfferRide("Rahul", "Hyderabad", "Bangalore", "XUV", "TS-10-22895", 5)
	if err != nil {
		fmt.Println("Error offering ride for Rahul:", err)
	}

	err = rm.OfferRide("Rohan", "Bangalore", "Pune", "Swift", "KA-01-12345", 1)
	if err != nil {
		fmt.Println("Error offering ride for Rohan (Bangalore to Pune):", err)
	}

	err = rm.OfferRide("Rahul", "Hyderabad", "Bangalore", "XUV", "TS-10-22895", 5)
	if err != nil {
		fmt.Println("Error offering ride for Rahul:", err)
	}

	err = rm.OfferRide("Rohan", "Bangalore", "Pune", "Swift", "KA-01-12345", 1)
	if err != nil {
		fmt.Println("Error offering ride for Rohan (Bangalore to Pune):", err)
	}

	err = rm.OfferRide("Armaan", "Mumbai", "karjat", "Polo", "KA-05-19088", 1)
	if err != nil {
		fmt.Println("Error offering ride for Rohan (Bangalore to Pune):", err)
	}

	err = rm.OfferRide("Shawn", "Karjat", "Lonavala", "WagonR", "KA-01-10005", 1)
	if err != nil {
		fmt.Println("Error offering ride for Rohan (Bangalore to Pune):", err)
	}

	fmt.Println("Select Ride")
	//  Select rides
	ride, err := rm.SelectRide("Nandini", "Bangalore", "Mysore", 1, "Most Vacant")
	if err != nil {
		fmt.Println("Error selecting ride for Nandini:", err)
	} else {
		fmt.Printf("Ride selected for Nandini: %+v\n", ride)
	}

	ride, err = rm.SelectRide("Gaurav", "Bangalore", "Mysore", 1, "Activa")
	if err != nil {
		fmt.Println("Error selecting ride for Gaurav:", err)
	} else {
		fmt.Printf("Ride selected for Gaurav: %+v\n", ride)
	}

	ride, err = rm.SelectRide("Shashank", "Mumbai", "Bangalore", 1, "Most Vacant")
	if err != nil {
		fmt.Println("Error selecting ride for Shashank:", err)
	} else {
		fmt.Printf("Ride selected for Shashank: %+v\n", ride)
	}

	ride, err = rm.SelectRide("Rohan", "Hyderabad", "Bangalore", 1, "Baleno")
	if err != nil {
		fmt.Println("Error selecting ride for Rohan:", err)
	} else {
		fmt.Printf("Ride selected for Rohan: %+v\n", ride)
	}

	ride, err = rm.SelectRide("Shashank", "Hyderabad", "Bangalore", 1, "Polo")
	if err != nil {
		fmt.Println("Error selecting ride for Shashank:", err)
	} else {
		fmt.Printf("Ride selected for Shashank: %+v\n", ride)
	}

	ride, err = rm.SelectRide("Tina", "Mumbai", "Lonavala", 1, "")
	if err != nil {
		fmt.Println("Error selecting ride for Shashank:", err)
	} else {
		fmt.Printf("Ride selected for Tina: %+v\n", ride)
	}

	// End rides
	fmt.Println("End Rides")

	err = rm.EndRide("Rohan:KA-01-12345")
	if err != nil {
		fmt.Println("Error ending ride for Rohan:", err)
	}

	ride, err = rm.GetRide("Rohan:KA-01-12345")
	if err != nil {
		fmt.Println("Error getting ride:", err)
	} else {
		fmt.Printf("Details of ride %s: %+v\n", "Rohan:KA-01-12345", ride)
	}

	err = rm.EndRide("Shipra:KA-12-12245")
	if err != nil {
		fmt.Println("Error ending ride for Shipra (Activa):", err)
	}

	err = rm.EndRide("Shipra:KA-05-19845")
	if err != nil {
		fmt.Println("Error ending ride for Shipra (Polo):", err)
	}

	err = rm.EndRide("Shashank:TS-04-22345")
	if err != nil {
		fmt.Println("Error ending ride for Shashank:", err)
	}

	rm.PrintRideStats()
}
