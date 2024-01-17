package internal

import "errors"

var (
	// ErrFieldRequired is an error returned when a field is missing
	ErrFieldRequired = errors.New("field required")
	// ErrVehicleAlreadyExists is an error returned when a vehicle already exists
	ErrVehicleAlreadyExists = errors.New("vehicle already exists")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Add is a method that adds a new vehicle to the repository
	Add(v *Vehicle) error
}
