package ride

// Ride represents a ride in the ride-sharing application
type Ride struct {
	Owner          string
	Origin         string
	Destination    string
	AvailableSeats int
	Vehicle        string
	Plate          string
}
