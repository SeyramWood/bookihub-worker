// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BookibusUsersColumns holds the columns for the "bookibus_users" table.
	BookibusUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_name", Type: field.TypeString},
		{Name: "other_name", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "other_phone", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"super_admin", "admin"}, Default: "super_admin"},
	}
	// BookibusUsersTable holds the schema information for the "bookibus_users" table.
	BookibusUsersTable = &schema.Table{
		Name:       "bookibus_users",
		Columns:    BookibusUsersColumns,
		PrimaryKey: []*schema.Column{BookibusUsersColumns[0]},
	}
	// BookingsColumns holds the columns for the "bookings" table.
	BookingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "reference", Type: field.TypeString, Nullable: true},
		{Name: "booking_number", Type: field.TypeString},
		{Name: "vat", Type: field.TypeFloat64, Default: 0},
		{Name: "sms_fee", Type: field.TypeFloat64, Default: 0},
		{Name: "amount", Type: field.TypeFloat64, Default: 0},
		{Name: "refund_amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "paid_at", Type: field.TypeTime, Nullable: true},
		{Name: "refund_at", Type: field.TypeTime, Nullable: true},
		{Name: "tans_type", Type: field.TypeEnum, Enums: []string{"momo", "card", "cash"}, Default: "cash"},
		{Name: "sms_notification", Type: field.TypeBool, Default: false},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"successful", "canceled"}, Default: "successful"},
		{Name: "company_bookings", Type: field.TypeInt, Nullable: true},
		{Name: "customer_bookings", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "trip_bookings", Type: field.TypeInt, Nullable: true},
	}
	// BookingsTable holds the schema information for the "bookings" table.
	BookingsTable = &schema.Table{
		Name:       "bookings",
		Columns:    BookingsColumns,
		PrimaryKey: []*schema.Column{BookingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bookings_companies_bookings",
				Columns:    []*schema.Column{BookingsColumns[14]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "bookings_customers_bookings",
				Columns:    []*schema.Column{BookingsColumns[15]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "bookings_trips_bookings",
				Columns:    []*schema.Column{BookingsColumns[16]},
				RefColumns: []*schema.Column{TripsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CompaniesColumns holds the columns for the "companies" table.
	CompaniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "other_phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString},
	}
	// CompaniesTable holds the schema information for the "companies" table.
	CompaniesTable = &schema.Table{
		Name:       "companies",
		Columns:    CompaniesColumns,
		PrimaryKey: []*schema.Column{CompaniesColumns[0]},
	}
	// CompanyUsersColumns holds the columns for the "company_users" table.
	CompanyUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_name", Type: field.TypeString, Nullable: true},
		{Name: "other_name", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "other_phone", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "manager", "teller", "driver"}, Default: "admin"},
		{Name: "company_profile", Type: field.TypeInt, Nullable: true},
	}
	// CompanyUsersTable holds the schema information for the "company_users" table.
	CompanyUsersTable = &schema.Table{
		Name:       "company_users",
		Columns:    CompanyUsersColumns,
		PrimaryKey: []*schema.Column{CompanyUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "company_users_companies_profile",
				Columns:    []*schema.Column{CompanyUsersColumns[8]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_name", Type: field.TypeString},
		{Name: "other_name", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "other_phone", Type: field.TypeString, Nullable: true},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
	}
	// CustomerContactsColumns holds the columns for the "customer_contacts" table.
	CustomerContactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "full_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "booking_contact", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CustomerContactsTable holds the schema information for the "customer_contacts" table.
	CustomerContactsTable = &schema.Table{
		Name:       "customer_contacts",
		Columns:    CustomerContactsColumns,
		PrimaryKey: []*schema.Column{CustomerContactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "customer_contacts_bookings_contact",
				Columns:    []*schema.Column{CustomerContactsColumns[6]},
				RefColumns: []*schema.Column{BookingsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CustomerLuggagesColumns holds the columns for the "customer_luggages" table.
	CustomerLuggagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "baggage", Type: field.TypeEnum, Nullable: true, Enums: []string{"excess", "bulky"}},
		{Name: "quantity", Type: field.TypeInt, Default: 0},
		{Name: "amount", Type: field.TypeFloat64, Default: 0},
		{Name: "booking_luggages", Type: field.TypeInt, Nullable: true},
	}
	// CustomerLuggagesTable holds the schema information for the "customer_luggages" table.
	CustomerLuggagesTable = &schema.Table{
		Name:       "customer_luggages",
		Columns:    CustomerLuggagesColumns,
		PrimaryKey: []*schema.Column{CustomerLuggagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "customer_luggages_bookings_luggages",
				Columns:    []*schema.Column{CustomerLuggagesColumns[6]},
				RefColumns: []*schema.Column{BookingsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// IncidentsColumns holds the columns for the "incidents" table.
	IncidentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "time", Type: field.TypeTime, Nullable: true},
		{Name: "location", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "type", Type: field.TypeString},
		{Name: "audio", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "resolved"}, Default: "pending"},
		{Name: "company_incidents", Type: field.TypeInt, Nullable: true},
		{Name: "company_user_incidents", Type: field.TypeInt, Nullable: true},
		{Name: "trip_incidents", Type: field.TypeInt, Nullable: true},
	}
	// IncidentsTable holds the schema information for the "incidents" table.
	IncidentsTable = &schema.Table{
		Name:       "incidents",
		Columns:    IncidentsColumns,
		PrimaryKey: []*schema.Column{IncidentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "incidents_companies_incidents",
				Columns:    []*schema.Column{IncidentsColumns[9]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "incidents_company_users_incidents",
				Columns:    []*schema.Column{IncidentsColumns[10]},
				RefColumns: []*schema.Column{CompanyUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "incidents_trips_incidents",
				Columns:    []*schema.Column{IncidentsColumns[11]},
				RefColumns: []*schema.Column{TripsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// IncidentImagesColumns holds the columns for the "incident_images" table.
	IncidentImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "incident_images", Type: field.TypeInt, Nullable: true},
	}
	// IncidentImagesTable holds the schema information for the "incident_images" table.
	IncidentImagesTable = &schema.Table{
		Name:       "incident_images",
		Columns:    IncidentImagesColumns,
		PrimaryKey: []*schema.Column{IncidentImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "incident_images_incidents_images",
				Columns:    []*schema.Column{IncidentImagesColumns[4]},
				RefColumns: []*schema.Column{IncidentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// NotificationsColumns holds the columns for the "notifications" table.
	NotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "event", Type: field.TypeString},
		{Name: "activity", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "subject_type", Type: field.TypeString},
		{Name: "subject_id", Type: field.TypeInt, Nullable: true},
		{Name: "creator_type", Type: field.TypeString},
		{Name: "creator_id", Type: field.TypeInt, Nullable: true},
		{Name: "customer_read_at", Type: field.TypeString, Nullable: true},
		{Name: "bookibus_read_at", Type: field.TypeJSON, Nullable: true},
		{Name: "company_read_at", Type: field.TypeJSON, Nullable: true},
		{Name: "data", Type: field.TypeJSON, Nullable: true},
		{Name: "company_notifications", Type: field.TypeInt, Nullable: true},
	}
	// NotificationsTable holds the schema information for the "notifications" table.
	NotificationsTable = &schema.Table{
		Name:       "notifications",
		Columns:    NotificationsColumns,
		PrimaryKey: []*schema.Column{NotificationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notifications_companies_notifications",
				Columns:    []*schema.Column{NotificationsColumns[12]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ParcelsColumns holds the columns for the "parcels" table.
	ParcelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "parcel_code", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "sender_name", Type: field.TypeString},
		{Name: "sender_phone", Type: field.TypeString},
		{Name: "sender_email", Type: field.TypeString},
		{Name: "recipient_name", Type: field.TypeString},
		{Name: "recipient_phone", Type: field.TypeString},
		{Name: "recipient_location", Type: field.TypeString},
		{Name: "weight", Type: field.TypeFloat32, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Default: 0},
		{Name: "paid_at", Type: field.TypeTime, Nullable: true},
		{Name: "tans_type", Type: field.TypeEnum, Enums: []string{"momo", "card", "cash"}, Default: "cash"},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"outgoing", "delivered"}, Default: "outgoing"},
		{Name: "company_parcels", Type: field.TypeInt, Nullable: true},
		{Name: "company_user_parcels", Type: field.TypeInt, Nullable: true},
		{Name: "trip_parcels", Type: field.TypeInt, Nullable: true},
	}
	// ParcelsTable holds the schema information for the "parcels" table.
	ParcelsTable = &schema.Table{
		Name:       "parcels",
		Columns:    ParcelsColumns,
		PrimaryKey: []*schema.Column{ParcelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "parcels_companies_parcels",
				Columns:    []*schema.Column{ParcelsColumns[16]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "parcels_company_users_parcels",
				Columns:    []*schema.Column{ParcelsColumns[17]},
				RefColumns: []*schema.Column{CompanyUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "parcels_trips_parcels",
				Columns:    []*schema.Column{ParcelsColumns[18]},
				RefColumns: []*schema.Column{TripsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ParcelImagesColumns holds the columns for the "parcel_images" table.
	ParcelImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"parcel", "recipient"}, Default: "parcel"},
		{Name: "parcel_images", Type: field.TypeInt, Nullable: true},
	}
	// ParcelImagesTable holds the schema information for the "parcel_images" table.
	ParcelImagesTable = &schema.Table{
		Name:       "parcel_images",
		Columns:    ParcelImagesColumns,
		PrimaryKey: []*schema.Column{ParcelImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "parcel_images_parcels_images",
				Columns:    []*schema.Column{ParcelImagesColumns[5]},
				RefColumns: []*schema.Column{ParcelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PassengersColumns holds the columns for the "passengers" table.
	PassengersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "full_name", Type: field.TypeString},
		{Name: "amount", Type: field.TypeFloat64, Default: 0},
		{Name: "maturity", Type: field.TypeEnum, Enums: []string{"adult", "child"}, Default: "adult"},
		{Name: "gender", Type: field.TypeEnum, Nullable: true, Enums: []string{"male", "female", "other"}},
		{Name: "booking_passengers", Type: field.TypeInt},
	}
	// PassengersTable holds the schema information for the "passengers" table.
	PassengersTable = &schema.Table{
		Name:       "passengers",
		Columns:    PassengersColumns,
		PrimaryKey: []*schema.Column{PassengersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "passengers_bookings_passengers",
				Columns:    []*schema.Column{PassengersColumns[7]},
				RefColumns: []*schema.Column{BookingsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RoutesColumns holds the columns for the "routes" table.
	RoutesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "from_location", Type: field.TypeString},
		{Name: "to_location", Type: field.TypeString},
		{Name: "from_latitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "from_longitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "to_latitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "to_longitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "rate", Type: field.TypeFloat64, Default: 0},
		{Name: "discount", Type: field.TypeFloat32, Default: 0},
		{Name: "popularity", Type: field.TypeInt, Default: 0},
		{Name: "company_routes", Type: field.TypeInt, Nullable: true},
	}
	// RoutesTable holds the schema information for the "routes" table.
	RoutesTable = &schema.Table{
		Name:       "routes",
		Columns:    RoutesColumns,
		PrimaryKey: []*schema.Column{RoutesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "routes_companies_routes",
				Columns:    []*schema.Column{RoutesColumns[12]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RouteStopsColumns holds the columns for the "route_stops" table.
	RouteStopsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "latitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "longitude", Type: field.TypeFloat64, Nullable: true},
		{Name: "route_stops", Type: field.TypeInt, Nullable: true},
	}
	// RouteStopsTable holds the schema information for the "route_stops" table.
	RouteStopsTable = &schema.Table{
		Name:       "route_stops",
		Columns:    RouteStopsColumns,
		PrimaryKey: []*schema.Column{RouteStopsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "route_stops_routes_stops",
				Columns:    []*schema.Column{RouteStopsColumns[5]},
				RefColumns: []*schema.Column{RoutesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TerminalsColumns holds the columns for the "terminals" table.
	TerminalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "company_terminals", Type: field.TypeInt, Nullable: true},
	}
	// TerminalsTable holds the schema information for the "terminals" table.
	TerminalsTable = &schema.Table{
		Name:       "terminals",
		Columns:    TerminalsColumns,
		PrimaryKey: []*schema.Column{TerminalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "terminals_companies_terminals",
				Columns:    []*schema.Column{TerminalsColumns[4]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TripsColumns holds the columns for the "trips" table.
	TripsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "departure_date", Type: field.TypeTime, Nullable: true},
		{Name: "arrival_date", Type: field.TypeTime, Nullable: true},
		{Name: "return_date", Type: field.TypeTime, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"oneway", "round"}, Default: "oneway"},
		{Name: "exterior_inspected", Type: field.TypeBool, Default: false},
		{Name: "interior_inspected", Type: field.TypeBool, Default: false},
		{Name: "engine_compartment_inspected", Type: field.TypeBool, Default: false},
		{Name: "brake_and_steering_inspected", Type: field.TypeBool, Default: false},
		{Name: "emergency_equipment_inspected", Type: field.TypeBool, Default: false},
		{Name: "fuel_and_fluids_inspected", Type: field.TypeBool, Default: false},
		{Name: "scheduled", Type: field.TypeBool, Default: false},
		{Name: "seat_left", Type: field.TypeInt, Default: 0},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"started", "ended"}},
		{Name: "company_trips", Type: field.TypeInt, Nullable: true},
		{Name: "company_user_trips", Type: field.TypeInt, Nullable: true},
		{Name: "route_trips", Type: field.TypeInt, Nullable: true},
		{Name: "terminal_from", Type: field.TypeInt, Nullable: true},
		{Name: "terminal_to", Type: field.TypeInt, Nullable: true},
		{Name: "vehicle_trips", Type: field.TypeInt, Nullable: true},
	}
	// TripsTable holds the schema information for the "trips" table.
	TripsTable = &schema.Table{
		Name:       "trips",
		Columns:    TripsColumns,
		PrimaryKey: []*schema.Column{TripsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "trips_companies_trips",
				Columns:    []*schema.Column{TripsColumns[16]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "trips_company_users_trips",
				Columns:    []*schema.Column{TripsColumns[17]},
				RefColumns: []*schema.Column{CompanyUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "trips_routes_trips",
				Columns:    []*schema.Column{TripsColumns[18]},
				RefColumns: []*schema.Column{RoutesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "trips_terminals_from",
				Columns:    []*schema.Column{TripsColumns[19]},
				RefColumns: []*schema.Column{TerminalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "trips_terminals_to",
				Columns:    []*schema.Column{TripsColumns[20]},
				RefColumns: []*schema.Column{TerminalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "trips_vehicles_trips",
				Columns:    []*schema.Column{TripsColumns[21]},
				RefColumns: []*schema.Column{VehiclesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeBytes},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"bookibus", "company", "customer"}},
		{Name: "bookibus_user_profile", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "company_user_profile", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "customer_profile", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_bookibus_users_profile",
				Columns:    []*schema.Column{UsersColumns[6]},
				RefColumns: []*schema.Column{BookibusUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "users_company_users_profile",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{CompanyUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "users_customers_profile",
				Columns:    []*schema.Column{UsersColumns[8]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// VehiclesColumns holds the columns for the "vehicles" table.
	VehiclesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "registration_number", Type: field.TypeString},
		{Name: "model", Type: field.TypeString},
		{Name: "seat", Type: field.TypeInt, Default: 0},
		{Name: "company_vehicles", Type: field.TypeInt, Nullable: true},
	}
	// VehiclesTable holds the schema information for the "vehicles" table.
	VehiclesTable = &schema.Table{
		Name:       "vehicles",
		Columns:    VehiclesColumns,
		PrimaryKey: []*schema.Column{VehiclesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vehicles_companies_vehicles",
				Columns:    []*schema.Column{VehiclesColumns[6]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// VehicleImagesColumns holds the columns for the "vehicle_images" table.
	VehicleImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "image", Type: field.TypeString},
		{Name: "vehicle_images", Type: field.TypeInt, Nullable: true},
	}
	// VehicleImagesTable holds the schema information for the "vehicle_images" table.
	VehicleImagesTable = &schema.Table{
		Name:       "vehicle_images",
		Columns:    VehicleImagesColumns,
		PrimaryKey: []*schema.Column{VehicleImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vehicle_images_vehicles_images",
				Columns:    []*schema.Column{VehicleImagesColumns[4]},
				RefColumns: []*schema.Column{VehiclesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// BookibusUserNotificationsColumns holds the columns for the "bookibus_user_notifications" table.
	BookibusUserNotificationsColumns = []*schema.Column{
		{Name: "bookibus_user_id", Type: field.TypeInt},
		{Name: "notification_id", Type: field.TypeInt},
	}
	// BookibusUserNotificationsTable holds the schema information for the "bookibus_user_notifications" table.
	BookibusUserNotificationsTable = &schema.Table{
		Name:       "bookibus_user_notifications",
		Columns:    BookibusUserNotificationsColumns,
		PrimaryKey: []*schema.Column{BookibusUserNotificationsColumns[0], BookibusUserNotificationsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bookibus_user_notifications_bookibus_user_id",
				Columns:    []*schema.Column{BookibusUserNotificationsColumns[0]},
				RefColumns: []*schema.Column{BookibusUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "bookibus_user_notifications_notification_id",
				Columns:    []*schema.Column{BookibusUserNotificationsColumns[1]},
				RefColumns: []*schema.Column{NotificationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CompanyUserNotificationsColumns holds the columns for the "company_user_notifications" table.
	CompanyUserNotificationsColumns = []*schema.Column{
		{Name: "company_user_id", Type: field.TypeInt},
		{Name: "notification_id", Type: field.TypeInt},
	}
	// CompanyUserNotificationsTable holds the schema information for the "company_user_notifications" table.
	CompanyUserNotificationsTable = &schema.Table{
		Name:       "company_user_notifications",
		Columns:    CompanyUserNotificationsColumns,
		PrimaryKey: []*schema.Column{CompanyUserNotificationsColumns[0], CompanyUserNotificationsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "company_user_notifications_company_user_id",
				Columns:    []*schema.Column{CompanyUserNotificationsColumns[0]},
				RefColumns: []*schema.Column{CompanyUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "company_user_notifications_notification_id",
				Columns:    []*schema.Column{CompanyUserNotificationsColumns[1]},
				RefColumns: []*schema.Column{NotificationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CustomerNotificationsColumns holds the columns for the "customer_notifications" table.
	CustomerNotificationsColumns = []*schema.Column{
		{Name: "customer_id", Type: field.TypeInt},
		{Name: "notification_id", Type: field.TypeInt},
	}
	// CustomerNotificationsTable holds the schema information for the "customer_notifications" table.
	CustomerNotificationsTable = &schema.Table{
		Name:       "customer_notifications",
		Columns:    CustomerNotificationsColumns,
		PrimaryKey: []*schema.Column{CustomerNotificationsColumns[0], CustomerNotificationsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "customer_notifications_customer_id",
				Columns:    []*schema.Column{CustomerNotificationsColumns[0]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "customer_notifications_notification_id",
				Columns:    []*schema.Column{CustomerNotificationsColumns[1]},
				RefColumns: []*schema.Column{NotificationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BookibusUsersTable,
		BookingsTable,
		CompaniesTable,
		CompanyUsersTable,
		CustomersTable,
		CustomerContactsTable,
		CustomerLuggagesTable,
		IncidentsTable,
		IncidentImagesTable,
		NotificationsTable,
		ParcelsTable,
		ParcelImagesTable,
		PassengersTable,
		RoutesTable,
		RouteStopsTable,
		TerminalsTable,
		TripsTable,
		UsersTable,
		VehiclesTable,
		VehicleImagesTable,
		BookibusUserNotificationsTable,
		CompanyUserNotificationsTable,
		CustomerNotificationsTable,
	}
)

func init() {
	BookingsTable.ForeignKeys[0].RefTable = CompaniesTable
	BookingsTable.ForeignKeys[1].RefTable = CustomersTable
	BookingsTable.ForeignKeys[2].RefTable = TripsTable
	CompanyUsersTable.ForeignKeys[0].RefTable = CompaniesTable
	CustomerContactsTable.ForeignKeys[0].RefTable = BookingsTable
	CustomerLuggagesTable.ForeignKeys[0].RefTable = BookingsTable
	IncidentsTable.ForeignKeys[0].RefTable = CompaniesTable
	IncidentsTable.ForeignKeys[1].RefTable = CompanyUsersTable
	IncidentsTable.ForeignKeys[2].RefTable = TripsTable
	IncidentImagesTable.ForeignKeys[0].RefTable = IncidentsTable
	NotificationsTable.ForeignKeys[0].RefTable = CompaniesTable
	ParcelsTable.ForeignKeys[0].RefTable = CompaniesTable
	ParcelsTable.ForeignKeys[1].RefTable = CompanyUsersTable
	ParcelsTable.ForeignKeys[2].RefTable = TripsTable
	ParcelImagesTable.ForeignKeys[0].RefTable = ParcelsTable
	PassengersTable.ForeignKeys[0].RefTable = BookingsTable
	RoutesTable.ForeignKeys[0].RefTable = CompaniesTable
	RouteStopsTable.ForeignKeys[0].RefTable = RoutesTable
	TerminalsTable.ForeignKeys[0].RefTable = CompaniesTable
	TripsTable.ForeignKeys[0].RefTable = CompaniesTable
	TripsTable.ForeignKeys[1].RefTable = CompanyUsersTable
	TripsTable.ForeignKeys[2].RefTable = RoutesTable
	TripsTable.ForeignKeys[3].RefTable = TerminalsTable
	TripsTable.ForeignKeys[4].RefTable = TerminalsTable
	TripsTable.ForeignKeys[5].RefTable = VehiclesTable
	UsersTable.ForeignKeys[0].RefTable = BookibusUsersTable
	UsersTable.ForeignKeys[1].RefTable = CompanyUsersTable
	UsersTable.ForeignKeys[2].RefTable = CustomersTable
	VehiclesTable.ForeignKeys[0].RefTable = CompaniesTable
	VehicleImagesTable.ForeignKeys[0].RefTable = VehiclesTable
	BookibusUserNotificationsTable.ForeignKeys[0].RefTable = BookibusUsersTable
	BookibusUserNotificationsTable.ForeignKeys[1].RefTable = NotificationsTable
	CompanyUserNotificationsTable.ForeignKeys[0].RefTable = CompanyUsersTable
	CompanyUserNotificationsTable.ForeignKeys[1].RefTable = NotificationsTable
	CustomerNotificationsTable.ForeignKeys[0].RefTable = CustomersTable
	CustomerNotificationsTable.ForeignKeys[1].RefTable = NotificationsTable
}
