// Code generated by ent, DO NOT EDIT.

package customercontact

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the customercontact type in the database.
	Label = "customer_contact"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// EdgeBooking holds the string denoting the booking edge name in mutations.
	EdgeBooking = "booking"
	// Table holds the table name of the customercontact in the database.
	Table = "customer_contacts"
	// BookingTable is the table that holds the booking relation/edge.
	BookingTable = "customer_contacts"
	// BookingInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingInverseTable = "bookings"
	// BookingColumn is the table column denoting the booking relation/edge.
	BookingColumn = "booking_contact"
)

// Columns holds all SQL columns for customercontact fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldFullName,
	FieldEmail,
	FieldPhone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "customer_contacts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"booking_contact",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	FullNameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
)

// OrderOption defines the ordering options for the CustomerContact queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByFullName orders the results by the full_name field.
func ByFullName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByBookingField orders the results by booking field.
func ByBookingField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBookingStep(), sql.OrderByField(field, opts...))
	}
}
func newBookingStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookingInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, BookingTable, BookingColumn),
	)
}
