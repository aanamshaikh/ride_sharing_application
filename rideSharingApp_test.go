package main

import (
	"ride_sharing_application/ride"
	"ride_sharing_application/user"
	"ride_sharing_application/vehicle"
	"testing"
)

func TestUserManagement(t *testing.T) {
	// Initialize UserManager
	um := user.NewUserManager()

	// Test adding users
	um.AddUser("Rohan", "M", 24)
	um.AddUser("Shashank", "M", 29)

	// Test getting existing user
	user, err := um.GetUser("Rohan")
	if err != nil {
		t.Errorf("Failed to get existing user: %v", err)
	}
	if user.Name != "Rohan" || user.Age != 24 || user.Gender != "M" {
		t.Errorf("User details mismatch")
	}

	// Test getting non-existing user
	_, err = um.GetUser("UnknownUser")
	if err == nil {
		t.Errorf("Expected error for non-existing user")
	}
}

func TestRideOffering(t *testing.T) {
	um := user.NewUserManager()
	vm := vehicle.NewVehicleManager(um)
	rm := ride.NewRideManager(um, vm)

	// Offer rides
	rm.OfferRide("Rohan", "Hyderabad", "Bangalore", "Swift", "KA-01-12345", 1)

	// Attempt to offer a ride for a vehicle already offered
	err := rm.OfferRide("Rohan", "Hyderabad", "Bangalore", "Swift", "KA-01-12345", 1)
	if err == nil {
		t.Errorf("Expected error for offering duplicate ride")
	}
}

func TestRideSelection(t *testing.T) {
	um := user.NewUserManager()
	vm := vehicle.NewVehicleManager(um)
	rm := ride.NewRideManager(um, vm)

	// Offer rides
	rm.OfferRide("Shipra", "Hyderabad", "Bangalore", "swift", "KA-05-19845", 2)

	// Select rides
	_, err := rm.SelectRide("Nandini", "Bangalore", "Mysore", 1, "Most Vacant")
	if err == nil {
		t.Errorf("Expected error for no available rides")
	}

	// Test selecting ride with preferred vehicle not available
	_, err = rm.SelectRide("Rohan", "Hyderabad", "Bangalore", 1, "Polo")
	if err == nil {
		t.Errorf("Expected error for preferred vehicle not available")
	}

	//Test selecting ride with avaiblae rides
	_, err = rm.SelectRide("Rohan", "Hyderabad", "Bangalore", 1, "swift")
	if err == nil {
		t.Errorf("Expected error for preferred vehicle not available")
	}
}

func TestRideEnding(t *testing.T) {
	um := user.NewUserManager()
	vm := vehicle.NewVehicleManager(um)
	rm := ride.NewRideManager(um, vm)

	// Offer a ride
	rm.OfferRide("Rohan", "Hyderabad", "Bangalore", "Swift", "KA-01-12345", 1)

	// End the ride
	err := rm.EndRide("Rohan:KA-01-12345")
	if err == nil {
		t.Errorf("Error ending ride for Rohan: %v", err)
	}

	// Attempt to end a ride that does not exist
	err = rm.EndRide("Unknown:UNKNOWN-12345")
	if err == nil {
		t.Errorf("Expected error for ending non-existent ride")
	}
}
