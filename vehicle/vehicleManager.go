package vehicle

// VehicleManager defines the methods for vehicle management
type VehicleManager interface {
	AddVehicle(owner, model, plate string) error
	GetVehicles(owner string) ([]Vehicle, error)
}
