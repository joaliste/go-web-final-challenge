package internal

import "errors"

var (
	// ErrVehicleIdAlreadyExists is the error returned when a vehicle id already exists
	ErrVehicleIdAlreadyExists = errors.New("vehicle id already exists")
	// ErrVehicleRegistrationAlreadyExists is the error returned when a vehicle registration already exists
	ErrVehicleRegistrationAlreadyExists = errors.New("vehicle registration already exists")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Add is a method that adds a new vehicle to the repository
	Add(v *Vehicle) error
}
