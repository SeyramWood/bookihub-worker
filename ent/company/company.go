// Code generated by ent, DO NOT EDIT.

package company

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the company type in the database.
	Label = "company"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldCertificate holds the string denoting the certificate field in the database.
	FieldCertificate = "certificate"
	// FieldBankAccount holds the string denoting the bank_account field in the database.
	FieldBankAccount = "bank_account"
	// FieldContactPerson holds the string denoting the contact_person field in the database.
	FieldContactPerson = "contact_person"
	// FieldOnboardingStatus holds the string denoting the onboarding_status field in the database.
	FieldOnboardingStatus = "onboarding_status"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// EdgeTerminals holds the string denoting the terminals edge name in mutations.
	EdgeTerminals = "terminals"
	// EdgeVehicles holds the string denoting the vehicles edge name in mutations.
	EdgeVehicles = "vehicles"
	// EdgeRoutes holds the string denoting the routes edge name in mutations.
	EdgeRoutes = "routes"
	// EdgeTrips holds the string denoting the trips edge name in mutations.
	EdgeTrips = "trips"
	// EdgeBookings holds the string denoting the bookings edge name in mutations.
	EdgeBookings = "bookings"
	// EdgeIncidents holds the string denoting the incidents edge name in mutations.
	EdgeIncidents = "incidents"
	// EdgeParcels holds the string denoting the parcels edge name in mutations.
	EdgeParcels = "parcels"
	// EdgeTransactions holds the string denoting the transactions edge name in mutations.
	EdgeTransactions = "transactions"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// Table holds the table name of the company in the database.
	Table = "companies"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "company_users"
	// ProfileInverseTable is the table name for the CompanyUser entity.
	// It exists in this package in order to avoid circular dependency with the "companyuser" package.
	ProfileInverseTable = "company_users"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "company_profile"
	// TerminalsTable is the table that holds the terminals relation/edge.
	TerminalsTable = "terminals"
	// TerminalsInverseTable is the table name for the Terminal entity.
	// It exists in this package in order to avoid circular dependency with the "terminal" package.
	TerminalsInverseTable = "terminals"
	// TerminalsColumn is the table column denoting the terminals relation/edge.
	TerminalsColumn = "company_terminals"
	// VehiclesTable is the table that holds the vehicles relation/edge.
	VehiclesTable = "vehicles"
	// VehiclesInverseTable is the table name for the Vehicle entity.
	// It exists in this package in order to avoid circular dependency with the "vehicle" package.
	VehiclesInverseTable = "vehicles"
	// VehiclesColumn is the table column denoting the vehicles relation/edge.
	VehiclesColumn = "company_vehicles"
	// RoutesTable is the table that holds the routes relation/edge.
	RoutesTable = "routes"
	// RoutesInverseTable is the table name for the Route entity.
	// It exists in this package in order to avoid circular dependency with the "route" package.
	RoutesInverseTable = "routes"
	// RoutesColumn is the table column denoting the routes relation/edge.
	RoutesColumn = "company_routes"
	// TripsTable is the table that holds the trips relation/edge.
	TripsTable = "trips"
	// TripsInverseTable is the table name for the Trip entity.
	// It exists in this package in order to avoid circular dependency with the "trip" package.
	TripsInverseTable = "trips"
	// TripsColumn is the table column denoting the trips relation/edge.
	TripsColumn = "company_trips"
	// BookingsTable is the table that holds the bookings relation/edge.
	BookingsTable = "bookings"
	// BookingsInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingsInverseTable = "bookings"
	// BookingsColumn is the table column denoting the bookings relation/edge.
	BookingsColumn = "company_bookings"
	// IncidentsTable is the table that holds the incidents relation/edge.
	IncidentsTable = "incidents"
	// IncidentsInverseTable is the table name for the Incident entity.
	// It exists in this package in order to avoid circular dependency with the "incident" package.
	IncidentsInverseTable = "incidents"
	// IncidentsColumn is the table column denoting the incidents relation/edge.
	IncidentsColumn = "company_incidents"
	// ParcelsTable is the table that holds the parcels relation/edge.
	ParcelsTable = "parcels"
	// ParcelsInverseTable is the table name for the Parcel entity.
	// It exists in this package in order to avoid circular dependency with the "parcel" package.
	ParcelsInverseTable = "parcels"
	// ParcelsColumn is the table column denoting the parcels relation/edge.
	ParcelsColumn = "company_parcels"
	// TransactionsTable is the table that holds the transactions relation/edge.
	TransactionsTable = "transactions"
	// TransactionsInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionsInverseTable = "transactions"
	// TransactionsColumn is the table column denoting the transactions relation/edge.
	TransactionsColumn = "company_transactions"
	// NotificationsTable is the table that holds the notifications relation/edge.
	NotificationsTable = "notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
	// NotificationsColumn is the table column denoting the notifications relation/edge.
	NotificationsColumn = "company_notifications"
)

// Columns holds all SQL columns for company fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldPhone,
	FieldEmail,
	FieldCertificate,
	FieldBankAccount,
	FieldContactPerson,
	FieldOnboardingStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
)

// OnboardingStatus defines the type for the "onboarding_status" enum field.
type OnboardingStatus string

// OnboardingStatusPending is the default value of the OnboardingStatus enum.
const DefaultOnboardingStatus = OnboardingStatusPending

// OnboardingStatus values.
const (
	OnboardingStatusPending  OnboardingStatus = "pending"
	OnboardingStatusApproved OnboardingStatus = "approved"
	OnboardingStatusRejected OnboardingStatus = "rejected"
)

func (os OnboardingStatus) String() string {
	return string(os)
}

// OnboardingStatusValidator is a validator for the "onboarding_status" field enum values. It is called by the builders before save.
func OnboardingStatusValidator(os OnboardingStatus) error {
	switch os {
	case OnboardingStatusPending, OnboardingStatusApproved, OnboardingStatusRejected:
		return nil
	default:
		return fmt.Errorf("company: invalid enum value for onboarding_status field: %q", os)
	}
}

// OrderOption defines the ordering options for the Company queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByCertificate orders the results by the certificate field.
func ByCertificate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCertificate, opts...).ToFunc()
}

// ByOnboardingStatus orders the results by the onboarding_status field.
func ByOnboardingStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOnboardingStatus, opts...).ToFunc()
}

// ByProfileCount orders the results by profile count.
func ByProfileCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProfileStep(), opts...)
	}
}

// ByProfile orders the results by profile terms.
func ByProfile(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfileStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTerminalsCount orders the results by terminals count.
func ByTerminalsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTerminalsStep(), opts...)
	}
}

// ByTerminals orders the results by terminals terms.
func ByTerminals(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTerminalsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVehiclesCount orders the results by vehicles count.
func ByVehiclesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVehiclesStep(), opts...)
	}
}

// ByVehicles orders the results by vehicles terms.
func ByVehicles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVehiclesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRoutesCount orders the results by routes count.
func ByRoutesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoutesStep(), opts...)
	}
}

// ByRoutes orders the results by routes terms.
func ByRoutes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoutesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTripsCount orders the results by trips count.
func ByTripsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTripsStep(), opts...)
	}
}

// ByTrips orders the results by trips terms.
func ByTrips(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTripsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBookingsCount orders the results by bookings count.
func ByBookingsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBookingsStep(), opts...)
	}
}

// ByBookings orders the results by bookings terms.
func ByBookings(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBookingsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIncidentsCount orders the results by incidents count.
func ByIncidentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIncidentsStep(), opts...)
	}
}

// ByIncidents orders the results by incidents terms.
func ByIncidents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIncidentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByParcelsCount orders the results by parcels count.
func ByParcelsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newParcelsStep(), opts...)
	}
}

// ByParcels orders the results by parcels terms.
func ByParcels(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParcelsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTransactionsCount orders the results by transactions count.
func ByTransactionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTransactionsStep(), opts...)
	}
}

// ByTransactions orders the results by transactions terms.
func ByTransactions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransactionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNotificationsCount orders the results by notifications count.
func ByNotificationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationsStep(), opts...)
	}
}

// ByNotifications orders the results by notifications terms.
func ByNotifications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProfileTable, ProfileColumn),
	)
}
func newTerminalsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TerminalsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TerminalsTable, TerminalsColumn),
	)
}
func newVehiclesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VehiclesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VehiclesTable, VehiclesColumn),
	)
}
func newRoutesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoutesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RoutesTable, RoutesColumn),
	)
}
func newTripsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TripsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TripsTable, TripsColumn),
	)
}
func newBookingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BookingsTable, BookingsColumn),
	)
}
func newIncidentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IncidentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IncidentsTable, IncidentsColumn),
	)
}
func newParcelsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ParcelsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ParcelsTable, ParcelsColumn),
	)
}
func newTransactionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransactionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
	)
}
func newNotificationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
	)
}
