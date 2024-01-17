package internal

import "errors"

var (
	// ErrFieldRequired is an error returned when a field is missing
	ErrFieldRequired = errors.New("field required")
	// ErrVehicleAlreadyExists is an error returned when a vehicle already exists
	ErrVehicleAlreadyExists = errors.New("vehicle already exists")
	// ErrInvalidFieldValue is an error returned when a field have an invalid value
	ErrInvalidFieldValue = errors.New("field with invalid value")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Add is a method that adds a new vehicle to the repository
	Add(v *Vehicle) error
	// GetByColorAndYear is a method that returns a map of vehicles with a specific color and year
	GetByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	// GetByBrandAndYears is a method that returns a map of vehicles with a specific brand
	// and between two years
	GetByBrandAndYears(brand string, startYear, endYear int) (v map[int]Vehicle, err error)
	// GetAverageSpeedByBrand is a method that returns the average speed of the vehicles of a brand
	GetAverageSpeedByBrand(brand string) (s float64, err error)
	// AddBatch is a method that adds a new vehicles to the repository
	AddBatch(vSlice []*Vehicle) error
	// UpdateSpeed is a method that updates the speed of a vehicle
	UpdateSpeed(speed float64, id int) error
	// GetByFuelType is a method that returns a map of vehicles with a type of fuel
	GetByFuelType(fuelType string) (v map[int]Vehicle, err error)
}
