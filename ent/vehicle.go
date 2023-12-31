// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/vehicle"
)

// Vehicle is the model entity for the Vehicle schema.
type Vehicle struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// RegistrationNumber holds the value of the "registration_number" field.
	RegistrationNumber string `json:"registration_number,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// Seat holds the value of the "seat" field.
	Seat int `json:"seat,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VehicleQuery when eager-loading is set.
	Edges            VehicleEdges `json:"edges"`
	company_vehicles *int
	selectValues     sql.SelectValues
}

// VehicleEdges holds the relations/edges for other nodes in the graph.
type VehicleEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// Images holds the value of the images edge.
	Images []*VehicleImage `json:"images,omitempty"`
	// Trips holds the value of the trips edge.
	Trips []*Trip `json:"trips,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VehicleEdges) CompanyOrErr() (*Company, error) {
	if e.loadedTypes[0] {
		if e.Company == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: company.Label}
		}
		return e.Company, nil
	}
	return nil, &NotLoadedError{edge: "company"}
}

// ImagesOrErr returns the Images value or an error if the edge
// was not loaded in eager-loading.
func (e VehicleEdges) ImagesOrErr() ([]*VehicleImage, error) {
	if e.loadedTypes[1] {
		return e.Images, nil
	}
	return nil, &NotLoadedError{edge: "images"}
}

// TripsOrErr returns the Trips value or an error if the edge
// was not loaded in eager-loading.
func (e VehicleEdges) TripsOrErr() ([]*Trip, error) {
	if e.loadedTypes[2] {
		return e.Trips, nil
	}
	return nil, &NotLoadedError{edge: "trips"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vehicle) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vehicle.FieldID, vehicle.FieldSeat:
			values[i] = new(sql.NullInt64)
		case vehicle.FieldRegistrationNumber, vehicle.FieldModel:
			values[i] = new(sql.NullString)
		case vehicle.FieldCreatedAt, vehicle.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case vehicle.ForeignKeys[0]: // company_vehicles
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vehicle fields.
func (v *Vehicle) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vehicle.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case vehicle.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		case vehicle.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				v.UpdatedAt = value.Time
			}
		case vehicle.FieldRegistrationNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field registration_number", values[i])
			} else if value.Valid {
				v.RegistrationNumber = value.String
			}
		case vehicle.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				v.Model = value.String
			}
		case vehicle.FieldSeat:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field seat", values[i])
			} else if value.Valid {
				v.Seat = int(value.Int64)
			}
		case vehicle.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_vehicles", value)
			} else if value.Valid {
				v.company_vehicles = new(int)
				*v.company_vehicles = int(value.Int64)
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Vehicle.
// This includes values selected through modifiers, order, etc.
func (v *Vehicle) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// QueryCompany queries the "company" edge of the Vehicle entity.
func (v *Vehicle) QueryCompany() *CompanyQuery {
	return NewVehicleClient(v.config).QueryCompany(v)
}

// QueryImages queries the "images" edge of the Vehicle entity.
func (v *Vehicle) QueryImages() *VehicleImageQuery {
	return NewVehicleClient(v.config).QueryImages(v)
}

// QueryTrips queries the "trips" edge of the Vehicle entity.
func (v *Vehicle) QueryTrips() *TripQuery {
	return NewVehicleClient(v.config).QueryTrips(v)
}

// Update returns a builder for updating this Vehicle.
// Note that you need to call Vehicle.Unwrap() before calling this method if this Vehicle
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vehicle) Update() *VehicleUpdateOne {
	return NewVehicleClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Vehicle entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vehicle) Unwrap() *Vehicle {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vehicle is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vehicle) String() string {
	var builder strings.Builder
	builder.WriteString("Vehicle(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(v.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("registration_number=")
	builder.WriteString(v.RegistrationNumber)
	builder.WriteString(", ")
	builder.WriteString("model=")
	builder.WriteString(v.Model)
	builder.WriteString(", ")
	builder.WriteString("seat=")
	builder.WriteString(fmt.Sprintf("%v", v.Seat))
	builder.WriteByte(')')
	return builder.String()
}

// Vehicles is a parsable slice of Vehicle.
type Vehicles []*Vehicle
