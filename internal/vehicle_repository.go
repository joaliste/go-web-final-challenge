package internal

import "errors"

var (
	// ErrVehicleIdAlreadyExists is the error returned when a vehicle id already exists
	ErrVehicleIdAlreadyExists = errors.New("vehicle id already exists")
	// ErrVehicleRegistrationAlreadyExists is the error returned when a vehicle registration already exists
	ErrVehicleRegistrationAlreadyExists = errors.New("vehicle registration already exists")
	// ErrVehiclesNotFound is the error returned when a vehicle registration already exists
	ErrVehiclesNotFound = errors.New("vehicles not found")
	// ErrVehicleIdNotFound is the error returned when a vehicle registration already exists
	ErrVehicleIdNotFound = errors.New("vehicle not found")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Add is a method that adds a new vehicle to the repository
	Add(v *Vehicle) (err error)
	// GetByColorAndYear is a method that returns a map of vehicles with a specific color and year
	GetByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	// GetByBrandAndYears is a method that returns a map of vehicles with a specific brand
	// and between two years
	GetByBrandAndYears(brand string, startYear, endYear int) (v map[int]Vehicle, err error)
	// GetByBrand is a method that returns the vehicles of a brand
	GetByBrand(brand string) (v map[int]Vehicle, err error)
	// AddBatch is a method that adds a new vehicles to the repository
	AddBatch(vSlice []*Vehicle) (err error)
	// UpdateSpeed is a method that updates the speed of a vehicle
	UpdateSpeed(speed float64, id int) (err error)
	// GetByFuelType is a method that returns a map of vehicles with a type of fuel
	GetByFuelType(fuelType string) (v map[int]Vehicle, err error)
	// DeleteVehicle is a method that deletes a vehicle
	DeleteVehicle(id int) (err error)
}
