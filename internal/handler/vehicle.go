package handler

import (
	"app/internal"
	"errors"
	"github.com/bootcamp-go/web/request"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/response"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// SpeedUpdateRequest is a struct that represents the speed update request.
type SpeedUpdateRequest struct {
	MaxSpeed float64 `json:"max_speed"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// AddVehicle is a method that adds a new vehicle to the vehicles map for the route post /vehicles
func (h *VehicleDefault) AddVehicle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody VehicleJSON
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}
		// process
		// - deserialize to vehicle
		v := internal.Vehicle{
			Id: reqBody.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           reqBody.Brand,
				Model:           reqBody.Model,
				Registration:    reqBody.Registration,
				Color:           reqBody.Color,
				FabricationYear: reqBody.FabricationYear,
				Capacity:        reqBody.Capacity,
				MaxSpeed:        reqBody.MaxSpeed,
				FuelType:        reqBody.FuelType,
				Transmission:    reqBody.Transmission,
				Weight:          reqBody.Weight,
				Dimensions: internal.Dimensions{
					Height: reqBody.Height,
					Length: reqBody.Length,
					Width:  reqBody.Width,
				},
			},
		}
		err = h.sv.Add(&v)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleAlreadyExists):
				response.Error(w, http.StatusConflict, "Vehicle already exists")
			case errors.Is(err, internal.ErrFieldRequired):
				response.Error(w, http.StatusBadRequest, "Some fields are missing")
			case errors.Is(err, internal.ErrInvalidFieldValue):
				response.Error(w, http.StatusBadRequest, "Some fields have invalid values")
			default:
				response.Error(w, http.StatusInternalServerError, "Internal error")
			}
			return
		}
		// response
		// - serialize to VehicleJSON
		responseData := VehicleJSON{
			ID:              v.Id,
			Brand:           v.Brand,
			Model:           v.Model,
			Registration:    v.Registration,
			Color:           v.Color,
			FabricationYear: v.FabricationYear,
			Capacity:        v.Capacity,
			MaxSpeed:        v.MaxSpeed,
			FuelType:        v.FuelType,
			Transmission:    v.Transmission,
			Weight:          v.Weight,
			Height:          v.Height,
			Length:          v.Length,
			Width:           v.Width,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "vehicle created",
			"data":    responseData,
		})
	}
}

// GetByColorAndYear is a method that returns a map of vehicles with a specific color and year.
// Pattern GET /vehicles/color/{color}/year/{year}
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		color := chi.URLParam(r, "color")
		year, err := strconv.Atoi(chi.URLParam(r, "year"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid year")
			return
		}

		// process
		// - get all vehicles
		v, err := h.sv.GetByColorAndYear(color, year)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFound):
				response.Error(w, http.StatusNotFound, "No vehicles found with that color and year")
			default:
				response.Error(w, http.StatusInternalServerError, "Internal error")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetByBrandAndYears is a method that returns a map of vehicles with a specific brand
// and between two years.
// Pattern GET /brand/{brand}/between/{start_year}/{end_year}
func (h *VehicleDefault) GetByBrandAndYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")
		startYear, err := strconv.Atoi(chi.URLParam(r, "start_year"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid start year")
			return
		}
		endYear, err := strconv.Atoi(chi.URLParam(r, "end_year"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid end year")
			return
		}

		// process
		// - get all vehicles
		v, err := h.sv.GetByBrandAndYears(brand, startYear, endYear)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFound):
				response.Error(
					w,
					http.StatusNotFound,
					"No vehicles found with that brand and between those years",
				)
			default:
				response.Error(w, http.StatusInternalServerError, "Internal error")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetAverageSpeedByBrand is a method that returns a float with the average speed for a brand
// Pattern GET /average_speed/brand/{brand}
func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")

		// process
		// - get average speed for a brand
		averageSpeed, err := h.sv.GetAverageSpeedByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFound):
				response.Error(
					w,
					http.StatusNotFound,
					"No vehicles found with that brand",
				)
			default:
				response.Error(w, http.StatusInternalServerError, "Internal error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message":       "success",
			"average_speed": averageSpeed,
		})
	}
}

// AddVehiclesByBatch is a method that adds batch of vehicles. Pattern /bach
func (h *VehicleDefault) AddVehiclesByBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody []VehicleJSON
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}
		// process
		// - deserialize to vehicle
		var deserializedData []*internal.Vehicle
		for _, v := range reqBody {
			deserializedV := internal.Vehicle{
				Id: v.ID,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           v.Brand,
					Model:           v.Model,
					Registration:    v.Registration,
					Color:           v.Color,
					FabricationYear: v.FabricationYear,
					Capacity:        v.Capacity,
					MaxSpeed:        v.MaxSpeed,
					FuelType:        v.FuelType,
					Transmission:    v.Transmission,
					Weight:          v.Weight,
					Dimensions: internal.Dimensions{
						Height: v.Height,
						Length: v.Length,
						Width:  v.Width,
					},
				},
			}
			deserializedData = append(deserializedData, &deserializedV)
		}
		err = h.sv.AddBatch(deserializedData)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleAlreadyExists):
				response.Error(w, http.StatusConflict, "Some of the vehicles already exist")
			case errors.Is(err, internal.ErrFieldRequired):
				response.Error(w, http.StatusBadRequest, "Some vehicles have missing fields")
			default:
				response.Error(w, http.StatusInternalServerError, "Internal error")
			}
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "vehicles successfully created",
		})
	}
}

// UpdateSpeed is a method that update the speed of a vehicle
func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil || id <= 0 {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}
		var reqBody SpeedUpdateRequest
		err = request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}
		// process

		err = h.sv.UpdateSpeed(reqBody.MaxSpeed, id)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleIdNotFound):
				response.Error(w, http.StatusConflict, "Vehicle with that id not found")
			default:
				response.Error(w, http.StatusInternalServerError, "Internal error")
			}
			return
		}
		// response

		response.JSON(w, http.StatusCreated, map[string]any{
			"message":   "max speed updated",
			"id":        id,
			"max_speed": reqBody.MaxSpeed,
		})
	}
}

// GetByFuelType is a method that returns a map of vehicles with a specific fuel type.
func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		fuelType := chi.URLParam(r, "type")

		// process
		// - get all vehicles
		v, err := h.sv.GetByFuelType(fuelType)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFound):
				response.Error(w, http.StatusNotFound, "No vehicles found with that fuel type")
			default:
				response.Error(w, http.StatusInternalServerError, "Internal error")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// DeleteVehicle is a method that deletes a vehicle
func (h *VehicleDefault) DeleteVehicle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil || id <= 0 {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}
		// process

		err = h.sv.DeleteVehicle(id)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleIdNotFound):
				response.Error(w, http.StatusConflict, "Vehicle with that id not found")
			default:
				response.Error(w, http.StatusInternalServerError, "Internal error")
			}
			return
		}
		// response
		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "vehicle successfully deleted",
			"id":      id,
		})
	}
}

// GetAverageCapacityByBrand is a method that returns a float with the average capacity for a brand
// Pattern GET /average_capacity/brand/{brand}
func (h *VehicleDefault) GetAverageCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")

		// process
		// - get average speed for a brand
		averageCapacity, err := h.sv.GetAverageCapacityByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFound):
				response.Error(
					w,
					http.StatusNotFound,
					"No vehicles found with that brand",
				)
			default:
				response.Error(w, http.StatusInternalServerError, "Internal error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message":          "success",
			"average_capacity": averageCapacity,
		})
	}
}

// GetByDimensions is a method that returns a map of vehicles with a specific dimension.
func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		length := strings.Split(r.URL.Query().Get("length"), "-")
		width := strings.Split(r.URL.Query().Get("width"), "-")

		minLength, err := strconv.ParseFloat(length[0], 64)
		if err != nil || minLength < 0 {
			response.Text(w, http.StatusBadRequest, "invalid min_length")
			return
		}
		maxLength, err := strconv.ParseFloat(length[1], 64)
		if err != nil || maxLength < 0 || minLength > maxLength {
			response.Text(w, http.StatusBadRequest, "invalid max_length")
			return
		}
		minWidth, err := strconv.ParseFloat(width[0], 64)
		if err != nil || minWidth < 0 {
			response.Text(w, http.StatusBadRequest, "invalid min_width")
			return
		}
		maxWidth, err := strconv.ParseFloat(width[1], 64)
		if err != nil || maxWidth < 0 || minWidth > maxWidth {
			response.Text(w, http.StatusBadRequest, "invalid max_width")
			return
		}

		// process
		// - get all vehicles by dimension
		v, err := h.sv.GetByDimensions(minLength, maxLength, minWidth, maxWidth)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehiclesNotFound):
				response.Error(w, http.StatusNotFound, "No vehicles found with that dimensions")
			default:
				response.Error(w, http.StatusInternalServerError, "Internal error")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
