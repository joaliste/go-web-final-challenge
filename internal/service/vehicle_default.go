package service

import (
	"app/internal"
	"errors"
	"fmt"
	"time"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Add is a method that adds a vehicle to the repository
func (s *VehicleDefault) Add(v *internal.Vehicle) error {
	err := validateAddVehicleRequestData(v)
	if err != nil {
		return err
	}

	err = s.rp.Add(v)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrVehicleIdAlreadyExists):
			err = fmt.Errorf("%w: id", internal.ErrVehicleAlreadyExists)
		case errors.Is(err, internal.ErrVehicleRegistrationAlreadyExists):
			err = fmt.Errorf("%w: registration", internal.ErrVehicleAlreadyExists)
		}
		return err
	}

	return nil
}

// validateAddVehicleRequestData is a function that validates the required fields of a vehicle
func validateAddVehicleRequestData(v *internal.Vehicle) error {
	// missing fields
	if v.Brand == "" {
		return fmt.Errorf("%w: brand", internal.ErrFieldRequired)
	}
	if v.Model == "" {
		return fmt.Errorf("%w: model", internal.ErrFieldRequired)
	}
	if v.Registration == "" {
		return fmt.Errorf("%w: brand", internal.ErrFieldRequired)
	}
	if v.Color == "" {
		return fmt.Errorf("%w: color", internal.ErrFieldRequired)
	}
	if v.FabricationYear == 0 {
		return fmt.Errorf("%w: year", internal.ErrFieldRequired)
	}
	if v.Capacity == 0 {
		return fmt.Errorf("%w: passenger", internal.ErrFieldRequired)
	}
	if v.MaxSpeed == 0 {
		return fmt.Errorf("%w: max_speed", internal.ErrFieldRequired)
	}

	if v.FuelType == "" {
		return fmt.Errorf("%w: fuel_type", internal.ErrFieldRequired)
	}

	if v.Transmission == "" {
		return fmt.Errorf("%w: transmission", internal.ErrFieldRequired)
	}

	if v.Weight == 0 {
		return fmt.Errorf("%w: weight", internal.ErrFieldRequired)
	}

	if v.Height == 0 {
		return fmt.Errorf("%w: height", internal.ErrFieldRequired)
	}

	if v.Length == 0 {
		return fmt.Errorf("%w: length", internal.ErrFieldRequired)
	}

	if v.Width == 0 {
		return fmt.Errorf("%w: width", internal.ErrFieldRequired)
	}
	// valid field values
	year, _, _ := time.Now().Date()
	if v.FabricationYear < 1900 || v.FabricationYear > year {
		return fmt.Errorf("%w: year", internal.ErrInvalidFieldValue)
	}

	if v.Capacity <= 0 || v.Capacity > 6 {
		return fmt.Errorf("%w: passenger", internal.ErrInvalidFieldValue)
	}

	if v.MaxSpeed < 0 || v.MaxSpeed > 300 {
		return fmt.Errorf("%w: max_speed", internal.ErrInvalidFieldValue)
	}

	if v.Weight < 0 || v.Weight > 500 {
		return fmt.Errorf("%w: weight", internal.ErrInvalidFieldValue)
	}

	if v.Height < 0 || v.Height > 500 {
		return fmt.Errorf("%w: height", internal.ErrInvalidFieldValue)
	}

	if v.Length < 0 || v.Length > 500 {
		return fmt.Errorf("%w: length", internal.ErrInvalidFieldValue)
	}

	if v.Width < 0 || v.Width > 0 {
		return fmt.Errorf("%w: width", internal.ErrInvalidFieldValue)
	}

	return nil
}

// GetByColorAndYear is a method that returns a map of vehicles with a specific color and year
func (s *VehicleDefault) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByColorAndYear(color, year)
	return
}

// GetByBrandAndYears is a method that returns a map of vehicles with a specific brand
// and between two years
func (s *VehicleDefault) GetByBrandAndYears(color string, startYear, endYear int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByBrandAndYears(color, startYear, endYear)
	return
}

// GetAverageSpeedByBrand is a method that returns the average speed of the vehicles of a brand
func (s *VehicleDefault) GetAverageSpeedByBrand(brand string) (as float64, err error) {
	vm, err := s.rp.GetByBrand(brand)

	if err != nil {
		return
	}

	count := len(vm)
	speedSum := 0.0
	for _, v := range vm {
		speedSum += v.MaxSpeed
	}
	as = speedSum / float64(count)

	return
}

// AddBatch is a method that adds a new vehicles to the repository
func (s *VehicleDefault) AddBatch(vSlice []*internal.Vehicle) error {
	for _, v := range vSlice {
		err := validateAddVehicleRequestData(v)
		if err != nil {
			return err
		}
	}

	err := s.rp.AddBatch(vSlice)

	if err != nil {
		switch {
		case errors.Is(err, internal.ErrVehicleIdAlreadyExists):
			err = fmt.Errorf("%w: id", internal.ErrVehicleAlreadyExists)
		case errors.Is(err, internal.ErrVehicleRegistrationAlreadyExists):
			err = fmt.Errorf("%w: registration", internal.ErrVehicleAlreadyExists)
		}
		return err
	}

	return nil
}

// UpdateSpeed is a method that updates the max speed of a vehicle
func (s *VehicleDefault) UpdateSpeed(speed float64, id int) error {
	if speed < 0 || speed > 300 {
		return internal.ErrInvalidFieldValue
	}

	err := s.rp.UpdateSpeed(speed, id)

	if err != nil {
		return err
	}

	return nil
}

// GetByFuelType is a method that returns a map of vehicles with a type of fuel
func (s *VehicleDefault) GetByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByFuelType(fuelType)
	return
}

// DeleteVehicle is a method that deletes a vehicle
func (s *VehicleDefault) DeleteVehicle(id int) (err error) {
	err = s.rp.DeleteVehicle(id)
	return
}

// GetAverageCapacityByBrand is a method that returns the average speed of the vehicles of a brand
func (s *VehicleDefault) GetAverageCapacityByBrand(brand string) (ac float64, err error) {
	vm, err := s.rp.GetByBrand(brand)

	if err != nil {
		return
	}

	count := len(vm)
	capacitySum := 0.0
	for _, v := range vm {
		capacitySum += float64(v.Capacity)
	}
	ac = capacitySum / float64(count)

	return
}

// GetByDimensions is a method that returns a map of vehicles with a specific dimension
func (s *VehicleDefault) GetByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByDimensions(minLength, maxLength, minWidth, maxWidth)
	return
}
