package ride

// RideManager defines the methods for ride management
type RideManager interface {
	OfferRide(owner, origin, destination, vehicle, plate string, availableSeats int) error
	SelectRide(user, origin, destination string, seats int, selectionStrategy string) (Ride, error)
	EndRide(rideKey string) error
	GetRide(rideKey string) (Ride, error)
	PrintRideStats()
}
