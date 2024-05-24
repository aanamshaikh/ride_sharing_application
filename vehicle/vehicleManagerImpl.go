package vehicle

import (
	"errors"
	"ride_sharing_application/user"
)

// vehicleManagerImpl implements the VehicleManager interface
type vehicleManagerImpl struct {
	Vehicles    map[string][]Vehicle
	UserManager user.UserManager
}

// NewVehicleManager creates a new instance of vehicleManagerImpl
func NewVehicleManager(userManager user.UserManager) VehicleManager {
	return &vehicleManagerImpl{
		Vehicles:    make(map[string][]Vehicle),
		UserManager: userManager,
	}
}

func (vm *vehicleManagerImpl) AddVehicle(owner, model, plate string) error {
	_, err := vm.UserManager.GetUser(owner)
	if err != nil {
		return errors.New("user not found")
	}

	vehicle := Vehicle{Owner: owner, Model: model, Plate: plate}
	vm.Vehicles[owner] = append(vm.Vehicles[owner], vehicle)
	return nil
}

func (vm *vehicleManagerImpl) GetVehicles(owner string) ([]Vehicle, error) {
	vehicles, exists := vm.Vehicles[owner]
	if !exists {
		return nil, errors.New("no vehicles found for this owner")
	}
	return vehicles, nil
}
