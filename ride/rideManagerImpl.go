package ride

import (
	"errors"
	"fmt"
	"ride_sharing_application/user"
	"ride_sharing_application/vehicle"
	"sort"
)

// RideManagerImpl implements the RideManager interface
type rideManagerImpl struct {
	UserManager    user.UserManager
	VehicleManager vehicle.VehicleManager
	Rides          map[string]Ride
	RidesTaken     map[string]int
	RidesOffered   map[string]int // Add RidesOffered map
}

// NewRideManager creates a new instance of rideManagerImpl
func NewRideManager(userManager user.UserManager, vehicleManager vehicle.VehicleManager) RideManager {
	return &rideManagerImpl{
		Rides:          make(map[string]Ride),
		RidesOffered:   make(map[string]int),
		RidesTaken:     make(map[string]int),
		UserManager:    userManager,
		VehicleManager: vehicleManager,
	}
}

func (rm *rideManagerImpl) OfferRide(owner, origin, destination, vehicle, plate string, availableSeats int) error {
	// Check if user exists
	if _, err := rm.UserManager.GetUser(owner); err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	// Check if vehicle exists and belongs to the user
	if !rm.isVehicleOwnedByUser(owner, vehicle, plate) {
		return errors.New("vehicle details do not match")
	}

	// Check if the ride is already offered for this vehicle
	rideKey := fmt.Sprintf("%s:%s", owner, plate)
	if _, exists := rm.Rides[rideKey]; exists {
		return errors.New("ride already offered for this vehicle")
	}

	// Offer the ride
	rm.Rides[rideKey] = Ride{
		Owner:          owner,
		Origin:         origin,
		Destination:    destination,
		AvailableSeats: availableSeats,
		Vehicle:        vehicle,
		Plate:          plate,
	}

	// Update the count of rides offered by the user
	rm.RidesOffered[owner]++
	return nil
}

// Helper method to check if the vehicle exists and belongs to the user
func (rm *rideManagerImpl) isVehicleOwnedByUser(owner, vehicle, plate string) bool {
	vehicles, err := rm.VehicleManager.GetVehicles(owner)
	if err != nil {
		return false
	}

	for _, v := range vehicles {
		if v.Model == vehicle && v.Plate == plate {
			return true
		}
	}
	return false
}

// check number of available seats while selecting
func (rm *rideManagerImpl) SelectRide(user, origin, destination string, seats int, selectionStrategy string) (Ride, error) {
	// Function to check if a ride matches the origin and destination
	rideMatchesRoute := func(ride Ride) bool {
		return ride.Origin == origin && ride.Destination == destination && ride.AvailableSeats >= seats
	}

	// Search for rides
	directRides, indirectRides := rm.findRides(origin, destination, rideMatchesRoute)

	// Select the best direct ride if available
	if len(directRides) > 0 {
		selectedRide, err := rm.selectBestRide(directRides, selectionStrategy, seats)
		if err == nil {
			rm.RidesTaken[user]++
			return selectedRide, nil
		}
	}

	// Select the best indirect ride if available
	if len(indirectRides) > 0 {
		selectedRide := rm.selectMostVacantRide(indirectRides, seats)
		rm.RidesTaken[user]++
		return selectedRide, nil
	}

	// If no rides are found, return an error
	return Ride{}, errors.New("no rides found")
}

// findRides searches for direct and indirect rides based on the origin and destination.
func (rm *rideManagerImpl) findRides(origin, destination string, rideMatchesRoute func(Ride) bool) ([]Ride, []Ride) {
	directRides := []Ride{}
	indirectRides := []Ride{}
	for _, ride := range rm.Rides {
		if rideMatchesRoute(ride) {
			directRides = append(directRides, ride)
		} else if ride.Origin == origin && ride.Destination != destination {
			indirectRides = append(indirectRides, ride)
		}
	}
	return directRides, indirectRides
}

// selectBestRide selects the best ride from a list of direct rides based on the selection strategy.
func (rm *rideManagerImpl) selectBestRide(rides []Ride, selectionStrategy string, seats int) (Ride, error) {
	switch selectionStrategy {
	case "Most Vacant":
		sort.Slice(rides, func(i, j int) bool {
			return rides[i].AvailableSeats > rides[j].AvailableSeats
		})
	case "":
		// No sorting needed for default strategy
	default:
		for _, ride := range rides {
			if ride.Vehicle == selectionStrategy {
				rm.updateAvailableSeats(&ride, seats)
				return ride, nil
			}
		}
		return Ride{}, errors.New("no preferred vehicle found")
	}

	selectedRide := rides[0]
	rm.updateAvailableSeats(&selectedRide, seats)
	return selectedRide, nil
}

// selectMostVacantRide selects the most vacant ride from a list of indirect rides.
func (rm *rideManagerImpl) selectMostVacantRide(rides []Ride, seats int) Ride {
	sort.Slice(rides, func(i, j int) bool {
		return rides[i].AvailableSeats > rides[j].AvailableSeats
	})
	selectedRide := rides[0]
	rm.updateAvailableSeats(&selectedRide, seats)
	return selectedRide
}

// updateAvailableSeats updates the available seats for the selected ride.
func (rm *rideManagerImpl) updateAvailableSeats(ride *Ride, seats int) {
	ride.AvailableSeats -= seats
	rm.Rides[fmt.Sprintf("%s:%s", ride.Owner, ride.Plate)] = *ride
}

func (rm *rideManagerImpl) EndRide(rideKey string) error {
	if _, exists := rm.Rides[rideKey]; !exists {
		return errors.New("ride not found")
	}
	delete(rm.Rides, rideKey)
	return nil
}

func (rm *rideManagerImpl) GetRide(rideKey string) (Ride, error) {
	ride, exists := rm.Rides[rideKey]
	if !exists {
		return Ride{}, errors.New("ride not found")
	}
	return ride, nil
}

func (rm *rideManagerImpl) PrintRideStats() {
	fmt.Println("Ride stats:")
	allUsers := make(map[string]bool)
	for user := range rm.RidesTaken {
		allUsers[user] = true
	}

	for user := range rm.RidesOffered {
		allUsers[user] = true
	}

	for user := range allUsers {
		ridesTaken := rm.RidesTaken[user]
		ridesOffered := rm.RidesOffered[user]
		fmt.Printf("%s: %d Taken, %d Offered", user, ridesTaken, ridesOffered)
	}
}
