package internal

import "errors"

var (
	// ErrVehicleIdAlreadyExists is the error returned when a vehicle id already exists
	ErrVehicleIdAlreadyExists = errors.New("vehicle id already exists")
	// ErrVehicleRegistrationAlreadyExists is the error returned when a vehicle registration already exists
	ErrVehicleRegistrationAlreadyExists = errors.New("vehicle registration already exists")
	// ErrVehiclesNotFound is the error returned when a vehicle registration already exists
	ErrVehiclesNotFound = errors.New("vehicles not found")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Add is a method that adds a new vehicle to the repository
	Add(v *Vehicle) error
	// GetByColorAndYear is a method that returns a map of vehicles with a specific color and year
	GetByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	// GetByBrandAndYears is a method that returns a map of vehicles with a specific brand
	// and between two years
	GetByBrandAndYears(brand string, startYear, endYear int) (v map[int]Vehicle, err error)
}
