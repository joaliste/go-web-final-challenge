package handler

import (
	"app/internal"
	"errors"
	"github.com/bootcamp-go/web/request"
	"net/http"

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
