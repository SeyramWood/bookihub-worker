// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/bookibus/ent/customer"
	"github.com/SeyramWood/bookibus/ent/user"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// OtherName holds the value of the "other_name" field.
	OtherName string `json:"other_name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// OtherPhone holds the value of the "other_phone" field.
	OtherPhone string `json:"other_phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges        CustomerEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Profile holds the value of the profile edge.
	Profile *User `json:"profile,omitempty"`
	// Bookings holds the value of the bookings edge.
	Bookings []*Booking `json:"bookings,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) ProfileOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Profile == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Profile, nil
	}
	return nil, &NotLoadedError{edge: "profile"}
}

// BookingsOrErr returns the Bookings value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) BookingsOrErr() ([]*Booking, error) {
	if e.loadedTypes[1] {
		return e.Bookings, nil
	}
	return nil, &NotLoadedError{edge: "bookings"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[2] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			values[i] = new(sql.NullInt64)
		case customer.FieldLastName, customer.FieldOtherName, customer.FieldPhone, customer.FieldOtherPhone:
			values[i] = new(sql.NullString)
		case customer.FieldCreatedAt, customer.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case customer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case customer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case customer.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				c.LastName = value.String
			}
		case customer.FieldOtherName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_name", values[i])
			} else if value.Valid {
				c.OtherName = value.String
			}
		case customer.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = value.String
			}
		case customer.FieldOtherPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_phone", values[i])
			} else if value.Valid {
				c.OtherPhone = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Customer.
// This includes values selected through modifiers, order, etc.
func (c *Customer) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryProfile queries the "profile" edge of the Customer entity.
func (c *Customer) QueryProfile() *UserQuery {
	return NewCustomerClient(c.config).QueryProfile(c)
}

// QueryBookings queries the "bookings" edge of the Customer entity.
func (c *Customer) QueryBookings() *BookingQuery {
	return NewCustomerClient(c.config).QueryBookings(c)
}

// QueryNotifications queries the "notifications" edge of the Customer entity.
func (c *Customer) QueryNotifications() *NotificationQuery {
	return NewCustomerClient(c.config).QueryNotifications(c)
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return NewCustomerClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(c.LastName)
	builder.WriteString(", ")
	builder.WriteString("other_name=")
	builder.WriteString(c.OtherName)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(c.Phone)
	builder.WriteString(", ")
	builder.WriteString("other_phone=")
	builder.WriteString(c.OtherPhone)
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer
