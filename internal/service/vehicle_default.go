package service

import (
	"app/internal"
	"errors"
	"fmt"
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

	return nil
}
