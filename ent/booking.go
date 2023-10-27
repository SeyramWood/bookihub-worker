// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/ent/booking"
	"github.com/SeyramWood/ent/company"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/customercontact"
	"github.com/SeyramWood/ent/trip"
)

// Booking is the model entity for the Booking schema.
type Booking struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// BookingNumber holds the value of the "booking_number" field.
	BookingNumber string `json:"booking_number,omitempty"`
	// BoardingPoint holds the value of the "boarding_point" field.
	BoardingPoint string `json:"boarding_point,omitempty"`
	// Vat holds the value of the "vat" field.
	Vat float64 `json:"vat,omitempty"`
	// SmsFee holds the value of the "sms_fee" field.
	SmsFee float64 `json:"sms_fee,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// RefundAmount holds the value of the "refund_amount" field.
	RefundAmount float64 `json:"refund_amount,omitempty"`
	// RefundAt holds the value of the "refund_at" field.
	RefundAt time.Time `json:"refund_at,omitempty"`
	// TansType holds the value of the "tans_type" field.
	TansType booking.TansType `json:"tans_type,omitempty"`
	// SmsNotification holds the value of the "sms_notification" field.
	SmsNotification bool `json:"sms_notification,omitempty"`
	// Status holds the value of the "status" field.
	Status booking.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookingQuery when eager-loading is set.
	Edges             BookingEdges `json:"edges"`
	company_bookings  *int
	customer_bookings *int
	trip_bookings     *int
	selectValues      sql.SelectValues
}

// BookingEdges holds the relations/edges for other nodes in the graph.
type BookingEdges struct {
	// Passengers holds the value of the passengers edge.
	Passengers []*Passenger `json:"passengers,omitempty"`
	// Luggages holds the value of the luggages edge.
	Luggages []*CustomerLuggage `json:"luggages,omitempty"`
	// Contact holds the value of the contact edge.
	Contact *CustomerContact `json:"contact,omitempty"`
	// Trip holds the value of the trip edge.
	Trip *Trip `json:"trip,omitempty"`
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// PassengersOrErr returns the Passengers value or an error if the edge
// was not loaded in eager-loading.
func (e BookingEdges) PassengersOrErr() ([]*Passenger, error) {
	if e.loadedTypes[0] {
		return e.Passengers, nil
	}
	return nil, &NotLoadedError{edge: "passengers"}
}

// LuggagesOrErr returns the Luggages value or an error if the edge
// was not loaded in eager-loading.
func (e BookingEdges) LuggagesOrErr() ([]*CustomerLuggage, error) {
	if e.loadedTypes[1] {
		return e.Luggages, nil
	}
	return nil, &NotLoadedError{edge: "luggages"}
}

// ContactOrErr returns the Contact value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) ContactOrErr() (*CustomerContact, error) {
	if e.loadedTypes[2] {
		if e.Contact == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customercontact.Label}
		}
		return e.Contact, nil
	}
	return nil, &NotLoadedError{edge: "contact"}
}

// TripOrErr returns the Trip value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) TripOrErr() (*Trip, error) {
	if e.loadedTypes[3] {
		if e.Trip == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: trip.Label}
		}
		return e.Trip, nil
	}
	return nil, &NotLoadedError{edge: "trip"}
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) CompanyOrErr() (*Company, error) {
	if e.loadedTypes[4] {
		if e.Company == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: company.Label}
		}
		return e.Company, nil
	}
	return nil, &NotLoadedError{edge: "company"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[5] {
		if e.Customer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Booking) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case booking.FieldSmsNotification:
			values[i] = new(sql.NullBool)
		case booking.FieldVat, booking.FieldSmsFee, booking.FieldAmount, booking.FieldRefundAmount:
			values[i] = new(sql.NullFloat64)
		case booking.FieldID:
			values[i] = new(sql.NullInt64)
		case booking.FieldBookingNumber, booking.FieldBoardingPoint, booking.FieldTansType, booking.FieldStatus:
			values[i] = new(sql.NullString)
		case booking.FieldCreatedAt, booking.FieldUpdatedAt, booking.FieldRefundAt:
			values[i] = new(sql.NullTime)
		case booking.ForeignKeys[0]: // company_bookings
			values[i] = new(sql.NullInt64)
		case booking.ForeignKeys[1]: // customer_bookings
			values[i] = new(sql.NullInt64)
		case booking.ForeignKeys[2]: // trip_bookings
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Booking fields.
func (b *Booking) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case booking.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case booking.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case booking.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case booking.FieldBookingNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field booking_number", values[i])
			} else if value.Valid {
				b.BookingNumber = value.String
			}
		case booking.FieldBoardingPoint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field boarding_point", values[i])
			} else if value.Valid {
				b.BoardingPoint = value.String
			}
		case booking.FieldVat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field vat", values[i])
			} else if value.Valid {
				b.Vat = value.Float64
			}
		case booking.FieldSmsFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field sms_fee", values[i])
			} else if value.Valid {
				b.SmsFee = value.Float64
			}
		case booking.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				b.Amount = value.Float64
			}
		case booking.FieldRefundAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field refund_amount", values[i])
			} else if value.Valid {
				b.RefundAmount = value.Float64
			}
		case booking.FieldRefundAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field refund_at", values[i])
			} else if value.Valid {
				b.RefundAt = value.Time
			}
		case booking.FieldTansType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tans_type", values[i])
			} else if value.Valid {
				b.TansType = booking.TansType(value.String)
			}
		case booking.FieldSmsNotification:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field sms_notification", values[i])
			} else if value.Valid {
				b.SmsNotification = value.Bool
			}
		case booking.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = booking.Status(value.String)
			}
		case booking.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_bookings", value)
			} else if value.Valid {
				b.company_bookings = new(int)
				*b.company_bookings = int(value.Int64)
			}
		case booking.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field customer_bookings", value)
			} else if value.Valid {
				b.customer_bookings = new(int)
				*b.customer_bookings = int(value.Int64)
			}
		case booking.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field trip_bookings", value)
			} else if value.Valid {
				b.trip_bookings = new(int)
				*b.trip_bookings = int(value.Int64)
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Booking.
// This includes values selected through modifiers, order, etc.
func (b *Booking) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryPassengers queries the "passengers" edge of the Booking entity.
func (b *Booking) QueryPassengers() *PassengerQuery {
	return NewBookingClient(b.config).QueryPassengers(b)
}

// QueryLuggages queries the "luggages" edge of the Booking entity.
func (b *Booking) QueryLuggages() *CustomerLuggageQuery {
	return NewBookingClient(b.config).QueryLuggages(b)
}

// QueryContact queries the "contact" edge of the Booking entity.
func (b *Booking) QueryContact() *CustomerContactQuery {
	return NewBookingClient(b.config).QueryContact(b)
}

// QueryTrip queries the "trip" edge of the Booking entity.
func (b *Booking) QueryTrip() *TripQuery {
	return NewBookingClient(b.config).QueryTrip(b)
}

// QueryCompany queries the "company" edge of the Booking entity.
func (b *Booking) QueryCompany() *CompanyQuery {
	return NewBookingClient(b.config).QueryCompany(b)
}

// QueryCustomer queries the "customer" edge of the Booking entity.
func (b *Booking) QueryCustomer() *CustomerQuery {
	return NewBookingClient(b.config).QueryCustomer(b)
}

// Update returns a builder for updating this Booking.
// Note that you need to call Booking.Unwrap() before calling this method if this Booking
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Booking) Update() *BookingUpdateOne {
	return NewBookingClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Booking entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Booking) Unwrap() *Booking {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Booking is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Booking) String() string {
	var builder strings.Builder
	builder.WriteString("Booking(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("booking_number=")
	builder.WriteString(b.BookingNumber)
	builder.WriteString(", ")
	builder.WriteString("boarding_point=")
	builder.WriteString(b.BoardingPoint)
	builder.WriteString(", ")
	builder.WriteString("vat=")
	builder.WriteString(fmt.Sprintf("%v", b.Vat))
	builder.WriteString(", ")
	builder.WriteString("sms_fee=")
	builder.WriteString(fmt.Sprintf("%v", b.SmsFee))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", b.Amount))
	builder.WriteString(", ")
	builder.WriteString("refund_amount=")
	builder.WriteString(fmt.Sprintf("%v", b.RefundAmount))
	builder.WriteString(", ")
	builder.WriteString("refund_at=")
	builder.WriteString(b.RefundAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("tans_type=")
	builder.WriteString(fmt.Sprintf("%v", b.TansType))
	builder.WriteString(", ")
	builder.WriteString("sms_notification=")
	builder.WriteString(fmt.Sprintf("%v", b.SmsNotification))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Bookings is a parsable slice of Booking.
type Bookings []*Booking