// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/route"
	"github.com/SeyramWood/ent/routestop"
)

// RouteStop is the model entity for the RouteStop schema.
type RouteStop struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude float64 `json:"longitude,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RouteStopQuery when eager-loading is set.
	Edges        RouteStopEdges `json:"edges"`
	route_stops  *int
	selectValues sql.SelectValues
}

// RouteStopEdges holds the relations/edges for other nodes in the graph.
type RouteStopEdges struct {
	// Route holds the value of the route edge.
	Route *Route `json:"route,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RouteOrErr returns the Route value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RouteStopEdges) RouteOrErr() (*Route, error) {
	if e.loadedTypes[0] {
		if e.Route == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: route.Label}
		}
		return e.Route, nil
	}
	return nil, &NotLoadedError{edge: "route"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RouteStop) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case routestop.FieldLatitude, routestop.FieldLongitude:
			values[i] = new(sql.NullFloat64)
		case routestop.FieldID:
			values[i] = new(sql.NullInt64)
		case routestop.FieldCreatedAt, routestop.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case routestop.ForeignKeys[0]: // route_stops
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RouteStop fields.
func (rs *RouteStop) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case routestop.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rs.ID = int(value.Int64)
		case routestop.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rs.CreatedAt = value.Time
			}
		case routestop.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rs.UpdatedAt = value.Time
			}
		case routestop.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				rs.Latitude = value.Float64
			}
		case routestop.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				rs.Longitude = value.Float64
			}
		case routestop.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field route_stops", value)
			} else if value.Valid {
				rs.route_stops = new(int)
				*rs.route_stops = int(value.Int64)
			}
		default:
			rs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RouteStop.
// This includes values selected through modifiers, order, etc.
func (rs *RouteStop) Value(name string) (ent.Value, error) {
	return rs.selectValues.Get(name)
}

// QueryRoute queries the "route" edge of the RouteStop entity.
func (rs *RouteStop) QueryRoute() *RouteQuery {
	return NewRouteStopClient(rs.config).QueryRoute(rs)
}

// Update returns a builder for updating this RouteStop.
// Note that you need to call RouteStop.Unwrap() before calling this method if this RouteStop
// was returned from a transaction, and the transaction was committed or rolled back.
func (rs *RouteStop) Update() *RouteStopUpdateOne {
	return NewRouteStopClient(rs.config).UpdateOne(rs)
}

// Unwrap unwraps the RouteStop entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rs *RouteStop) Unwrap() *RouteStop {
	_tx, ok := rs.config.driver.(*txDriver)
	if !ok {
		panic("ent: RouteStop is not a transactional entity")
	}
	rs.config.driver = _tx.drv
	return rs
}

// String implements the fmt.Stringer.
func (rs *RouteStop) String() string {
	var builder strings.Builder
	builder.WriteString("RouteStop(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rs.ID))
	builder.WriteString("created_at=")
	builder.WriteString(rs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rs.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(fmt.Sprintf("%v", rs.Latitude))
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(fmt.Sprintf("%v", rs.Longitude))
	builder.WriteByte(')')
	return builder.String()
}

// RouteStops is a parsable slice of RouteStop.
type RouteStops []*RouteStop