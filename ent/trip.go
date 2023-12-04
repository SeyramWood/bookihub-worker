// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/route"
	"github.com/SeyramWood/bookibus/ent/terminal"
	"github.com/SeyramWood/bookibus/ent/trip"
	"github.com/SeyramWood/bookibus/ent/vehicle"
)

// Trip is the model entity for the Trip schema.
type Trip struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DepartureDate holds the value of the "departure_date" field.
	DepartureDate time.Time `json:"departure_date,omitempty"`
	// ArrivalDate holds the value of the "arrival_date" field.
	ArrivalDate time.Time `json:"arrival_date,omitempty"`
	// ReturnDate holds the value of the "return_date" field.
	ReturnDate time.Time `json:"return_date,omitempty"`
	// Type holds the value of the "type" field.
	Type trip.Type `json:"type,omitempty"`
	// ExteriorInspected holds the value of the "exterior_inspected" field.
	ExteriorInspected bool `json:"exterior_inspected,omitempty"`
	// InteriorInspected holds the value of the "interior_inspected" field.
	InteriorInspected bool `json:"interior_inspected,omitempty"`
	// EngineCompartmentInspected holds the value of the "engine_compartment_inspected" field.
	EngineCompartmentInspected bool `json:"engine_compartment_inspected,omitempty"`
	// BrakeAndSteeringInspected holds the value of the "brake_and_steering_inspected" field.
	BrakeAndSteeringInspected bool `json:"brake_and_steering_inspected,omitempty"`
	// EmergencyEquipmentInspected holds the value of the "emergency_equipment_inspected" field.
	EmergencyEquipmentInspected bool `json:"emergency_equipment_inspected,omitempty"`
	// FuelAndFluidsInspected holds the value of the "fuel_and_fluids_inspected" field.
	FuelAndFluidsInspected bool `json:"fuel_and_fluids_inspected,omitempty"`
	// Scheduled holds the value of the "scheduled" field.
	Scheduled bool `json:"scheduled,omitempty"`
	// SeatLeft holds the value of the "seat_left" field.
	SeatLeft int `json:"seat_left,omitempty"`
	// Rate holds the value of the "rate" field.
	Rate float64 `json:"rate,omitempty"`
	// Discount holds the value of the "discount" field.
	Discount float32 `json:"discount,omitempty"`
	// Status holds the value of the "status" field.
	Status trip.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TripQuery when eager-loading is set.
	Edges              TripEdges `json:"edges"`
	company_trips      *int
	company_user_trips *int
	route_trips        *int
	terminal_from      *int
	terminal_to        *int
	vehicle_trips      *int
	selectValues       sql.SelectValues
}

// TripEdges holds the relations/edges for other nodes in the graph.
type TripEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// Driver holds the value of the driver edge.
	Driver *CompanyUser `json:"driver,omitempty"`
	// FromTerminal holds the value of the from_terminal edge.
	FromTerminal *Terminal `json:"from_terminal,omitempty"`
	// ToTerminal holds the value of the to_terminal edge.
	ToTerminal *Terminal `json:"to_terminal,omitempty"`
	// Vehicle holds the value of the vehicle edge.
	Vehicle *Vehicle `json:"vehicle,omitempty"`
	// Route holds the value of the route edge.
	Route *Route `json:"route,omitempty"`
	// Bookings holds the value of the bookings edge.
	Bookings []*Booking `json:"bookings,omitempty"`
	// Incidents holds the value of the incidents edge.
	Incidents []*Incident `json:"incidents,omitempty"`
	// Parcels holds the value of the parcels edge.
	Parcels []*Parcel `json:"parcels,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TripEdges) CompanyOrErr() (*Company, error) {
	if e.loadedTypes[0] {
		if e.Company == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: company.Label}
		}
		return e.Company, nil
	}
	return nil, &NotLoadedError{edge: "company"}
}

// DriverOrErr returns the Driver value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TripEdges) DriverOrErr() (*CompanyUser, error) {
	if e.loadedTypes[1] {
		if e.Driver == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: companyuser.Label}
		}
		return e.Driver, nil
	}
	return nil, &NotLoadedError{edge: "driver"}
}

// FromTerminalOrErr returns the FromTerminal value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TripEdges) FromTerminalOrErr() (*Terminal, error) {
	if e.loadedTypes[2] {
		if e.FromTerminal == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: terminal.Label}
		}
		return e.FromTerminal, nil
	}
	return nil, &NotLoadedError{edge: "from_terminal"}
}

// ToTerminalOrErr returns the ToTerminal value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TripEdges) ToTerminalOrErr() (*Terminal, error) {
	if e.loadedTypes[3] {
		if e.ToTerminal == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: terminal.Label}
		}
		return e.ToTerminal, nil
	}
	return nil, &NotLoadedError{edge: "to_terminal"}
}

// VehicleOrErr returns the Vehicle value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TripEdges) VehicleOrErr() (*Vehicle, error) {
	if e.loadedTypes[4] {
		if e.Vehicle == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: vehicle.Label}
		}
		return e.Vehicle, nil
	}
	return nil, &NotLoadedError{edge: "vehicle"}
}

// RouteOrErr returns the Route value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TripEdges) RouteOrErr() (*Route, error) {
	if e.loadedTypes[5] {
		if e.Route == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: route.Label}
		}
		return e.Route, nil
	}
	return nil, &NotLoadedError{edge: "route"}
}

// BookingsOrErr returns the Bookings value or an error if the edge
// was not loaded in eager-loading.
func (e TripEdges) BookingsOrErr() ([]*Booking, error) {
	if e.loadedTypes[6] {
		return e.Bookings, nil
	}
	return nil, &NotLoadedError{edge: "bookings"}
}

// IncidentsOrErr returns the Incidents value or an error if the edge
// was not loaded in eager-loading.
func (e TripEdges) IncidentsOrErr() ([]*Incident, error) {
	if e.loadedTypes[7] {
		return e.Incidents, nil
	}
	return nil, &NotLoadedError{edge: "incidents"}
}

// ParcelsOrErr returns the Parcels value or an error if the edge
// was not loaded in eager-loading.
func (e TripEdges) ParcelsOrErr() ([]*Parcel, error) {
	if e.loadedTypes[8] {
		return e.Parcels, nil
	}
	return nil, &NotLoadedError{edge: "parcels"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Trip) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case trip.FieldExteriorInspected, trip.FieldInteriorInspected, trip.FieldEngineCompartmentInspected, trip.FieldBrakeAndSteeringInspected, trip.FieldEmergencyEquipmentInspected, trip.FieldFuelAndFluidsInspected, trip.FieldScheduled:
			values[i] = new(sql.NullBool)
		case trip.FieldRate, trip.FieldDiscount:
			values[i] = new(sql.NullFloat64)
		case trip.FieldID, trip.FieldSeatLeft:
			values[i] = new(sql.NullInt64)
		case trip.FieldType, trip.FieldStatus:
			values[i] = new(sql.NullString)
		case trip.FieldCreatedAt, trip.FieldUpdatedAt, trip.FieldDepartureDate, trip.FieldArrivalDate, trip.FieldReturnDate:
			values[i] = new(sql.NullTime)
		case trip.ForeignKeys[0]: // company_trips
			values[i] = new(sql.NullInt64)
		case trip.ForeignKeys[1]: // company_user_trips
			values[i] = new(sql.NullInt64)
		case trip.ForeignKeys[2]: // route_trips
			values[i] = new(sql.NullInt64)
		case trip.ForeignKeys[3]: // terminal_from
			values[i] = new(sql.NullInt64)
		case trip.ForeignKeys[4]: // terminal_to
			values[i] = new(sql.NullInt64)
		case trip.ForeignKeys[5]: // vehicle_trips
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Trip fields.
func (t *Trip) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case trip.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case trip.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case trip.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case trip.FieldDepartureDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field departure_date", values[i])
			} else if value.Valid {
				t.DepartureDate = value.Time
			}
		case trip.FieldArrivalDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field arrival_date", values[i])
			} else if value.Valid {
				t.ArrivalDate = value.Time
			}
		case trip.FieldReturnDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field return_date", values[i])
			} else if value.Valid {
				t.ReturnDate = value.Time
			}
		case trip.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = trip.Type(value.String)
			}
		case trip.FieldExteriorInspected:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field exterior_inspected", values[i])
			} else if value.Valid {
				t.ExteriorInspected = value.Bool
			}
		case trip.FieldInteriorInspected:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field interior_inspected", values[i])
			} else if value.Valid {
				t.InteriorInspected = value.Bool
			}
		case trip.FieldEngineCompartmentInspected:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field engine_compartment_inspected", values[i])
			} else if value.Valid {
				t.EngineCompartmentInspected = value.Bool
			}
		case trip.FieldBrakeAndSteeringInspected:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field brake_and_steering_inspected", values[i])
			} else if value.Valid {
				t.BrakeAndSteeringInspected = value.Bool
			}
		case trip.FieldEmergencyEquipmentInspected:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field emergency_equipment_inspected", values[i])
			} else if value.Valid {
				t.EmergencyEquipmentInspected = value.Bool
			}
		case trip.FieldFuelAndFluidsInspected:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field fuel_and_fluids_inspected", values[i])
			} else if value.Valid {
				t.FuelAndFluidsInspected = value.Bool
			}
		case trip.FieldScheduled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field scheduled", values[i])
			} else if value.Valid {
				t.Scheduled = value.Bool
			}
		case trip.FieldSeatLeft:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field seat_left", values[i])
			} else if value.Valid {
				t.SeatLeft = int(value.Int64)
			}
		case trip.FieldRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field rate", values[i])
			} else if value.Valid {
				t.Rate = value.Float64
			}
		case trip.FieldDiscount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field discount", values[i])
			} else if value.Valid {
				t.Discount = float32(value.Float64)
			}
		case trip.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = trip.Status(value.String)
			}
		case trip.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_trips", value)
			} else if value.Valid {
				t.company_trips = new(int)
				*t.company_trips = int(value.Int64)
			}
		case trip.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_user_trips", value)
			} else if value.Valid {
				t.company_user_trips = new(int)
				*t.company_user_trips = int(value.Int64)
			}
		case trip.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field route_trips", value)
			} else if value.Valid {
				t.route_trips = new(int)
				*t.route_trips = int(value.Int64)
			}
		case trip.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field terminal_from", value)
			} else if value.Valid {
				t.terminal_from = new(int)
				*t.terminal_from = int(value.Int64)
			}
		case trip.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field terminal_to", value)
			} else if value.Valid {
				t.terminal_to = new(int)
				*t.terminal_to = int(value.Int64)
			}
		case trip.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field vehicle_trips", value)
			} else if value.Valid {
				t.vehicle_trips = new(int)
				*t.vehicle_trips = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Trip.
// This includes values selected through modifiers, order, etc.
func (t *Trip) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryCompany queries the "company" edge of the Trip entity.
func (t *Trip) QueryCompany() *CompanyQuery {
	return NewTripClient(t.config).QueryCompany(t)
}

// QueryDriver queries the "driver" edge of the Trip entity.
func (t *Trip) QueryDriver() *CompanyUserQuery {
	return NewTripClient(t.config).QueryDriver(t)
}

// QueryFromTerminal queries the "from_terminal" edge of the Trip entity.
func (t *Trip) QueryFromTerminal() *TerminalQuery {
	return NewTripClient(t.config).QueryFromTerminal(t)
}

// QueryToTerminal queries the "to_terminal" edge of the Trip entity.
func (t *Trip) QueryToTerminal() *TerminalQuery {
	return NewTripClient(t.config).QueryToTerminal(t)
}

// QueryVehicle queries the "vehicle" edge of the Trip entity.
func (t *Trip) QueryVehicle() *VehicleQuery {
	return NewTripClient(t.config).QueryVehicle(t)
}

// QueryRoute queries the "route" edge of the Trip entity.
func (t *Trip) QueryRoute() *RouteQuery {
	return NewTripClient(t.config).QueryRoute(t)
}

// QueryBookings queries the "bookings" edge of the Trip entity.
func (t *Trip) QueryBookings() *BookingQuery {
	return NewTripClient(t.config).QueryBookings(t)
}

// QueryIncidents queries the "incidents" edge of the Trip entity.
func (t *Trip) QueryIncidents() *IncidentQuery {
	return NewTripClient(t.config).QueryIncidents(t)
}

// QueryParcels queries the "parcels" edge of the Trip entity.
func (t *Trip) QueryParcels() *ParcelQuery {
	return NewTripClient(t.config).QueryParcels(t)
}

// Update returns a builder for updating this Trip.
// Note that you need to call Trip.Unwrap() before calling this method if this Trip
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Trip) Update() *TripUpdateOne {
	return NewTripClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Trip entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Trip) Unwrap() *Trip {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Trip is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Trip) String() string {
	var builder strings.Builder
	builder.WriteString("Trip(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("departure_date=")
	builder.WriteString(t.DepartureDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("arrival_date=")
	builder.WriteString(t.ArrivalDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("return_date=")
	builder.WriteString(t.ReturnDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteString(", ")
	builder.WriteString("exterior_inspected=")
	builder.WriteString(fmt.Sprintf("%v", t.ExteriorInspected))
	builder.WriteString(", ")
	builder.WriteString("interior_inspected=")
	builder.WriteString(fmt.Sprintf("%v", t.InteriorInspected))
	builder.WriteString(", ")
	builder.WriteString("engine_compartment_inspected=")
	builder.WriteString(fmt.Sprintf("%v", t.EngineCompartmentInspected))
	builder.WriteString(", ")
	builder.WriteString("brake_and_steering_inspected=")
	builder.WriteString(fmt.Sprintf("%v", t.BrakeAndSteeringInspected))
	builder.WriteString(", ")
	builder.WriteString("emergency_equipment_inspected=")
	builder.WriteString(fmt.Sprintf("%v", t.EmergencyEquipmentInspected))
	builder.WriteString(", ")
	builder.WriteString("fuel_and_fluids_inspected=")
	builder.WriteString(fmt.Sprintf("%v", t.FuelAndFluidsInspected))
	builder.WriteString(", ")
	builder.WriteString("scheduled=")
	builder.WriteString(fmt.Sprintf("%v", t.Scheduled))
	builder.WriteString(", ")
	builder.WriteString("seat_left=")
	builder.WriteString(fmt.Sprintf("%v", t.SeatLeft))
	builder.WriteString(", ")
	builder.WriteString("rate=")
	builder.WriteString(fmt.Sprintf("%v", t.Rate))
	builder.WriteString(", ")
	builder.WriteString("discount=")
	builder.WriteString(fmt.Sprintf("%v", t.Discount))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Trips is a parsable slice of Trip.
type Trips []*Trip
