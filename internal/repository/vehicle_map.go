package repository

import "app/internal"

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Add is a method that adds a new vehicle to the repository
func (r *VehicleMap) Add(v *internal.Vehicle) error {
	for _, vdb := range r.db {
		if vdb.Id == v.Id {
			return internal.ErrVehicleIdAlreadyExists
		}

		if vdb.Registration == v.Registration {
			return internal.ErrVehicleRegistrationAlreadyExists
		}
	}
	r.db[v.Id] = *v

	return nil
}

// GetByColorAndYear is a method that returns a map of vehicles with a specific color and year
func (r *VehicleMap) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db with the specific vehicles considering color and year
	for key, value := range r.db {
		if value.FabricationYear == year && value.Color == color {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrVehiclesNotFound
	}

	return
}

// GetByBrandAndYears is a method that returns a map of vehicles with a specific brand
// and between two years
func (r *VehicleMap) GetByBrandAndYears(brand string, startYear, endYear int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db with the specific vehicles considering brand and between two years
	for key, value := range r.db {
		if value.FabricationYear >= startYear && value.FabricationYear <= endYear && value.Brand == brand {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrVehiclesNotFound
	}

	return
}

// GetByBrand is a method that returns a map with vehicles from an specific brand
func (r *VehicleMap) GetByBrand(brand string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db with the specific vehicles considering color and year
	for key, value := range r.db {
		if value.Brand == brand {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrVehiclesNotFound
	}

	return
}
