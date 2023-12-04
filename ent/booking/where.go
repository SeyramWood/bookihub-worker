// Code generated by ent, DO NOT EDIT.

package booking

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/bookibus/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldUpdatedAt, v))
}

// BookingNumber applies equality check predicate on the "booking_number" field. It's identical to BookingNumberEQ.
func BookingNumber(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldBookingNumber, v))
}

// SmsNotification applies equality check predicate on the "sms_notification" field. It's identical to SmsNotificationEQ.
func SmsNotification(v bool) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldSmsNotification, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldUpdatedAt, v))
}

// BookingNumberEQ applies the EQ predicate on the "booking_number" field.
func BookingNumberEQ(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldBookingNumber, v))
}

// BookingNumberNEQ applies the NEQ predicate on the "booking_number" field.
func BookingNumberNEQ(v string) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldBookingNumber, v))
}

// BookingNumberIn applies the In predicate on the "booking_number" field.
func BookingNumberIn(vs ...string) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldBookingNumber, vs...))
}

// BookingNumberNotIn applies the NotIn predicate on the "booking_number" field.
func BookingNumberNotIn(vs ...string) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldBookingNumber, vs...))
}

// BookingNumberGT applies the GT predicate on the "booking_number" field.
func BookingNumberGT(v string) predicate.Booking {
	return predicate.Booking(sql.FieldGT(FieldBookingNumber, v))
}

// BookingNumberGTE applies the GTE predicate on the "booking_number" field.
func BookingNumberGTE(v string) predicate.Booking {
	return predicate.Booking(sql.FieldGTE(FieldBookingNumber, v))
}

// BookingNumberLT applies the LT predicate on the "booking_number" field.
func BookingNumberLT(v string) predicate.Booking {
	return predicate.Booking(sql.FieldLT(FieldBookingNumber, v))
}

// BookingNumberLTE applies the LTE predicate on the "booking_number" field.
func BookingNumberLTE(v string) predicate.Booking {
	return predicate.Booking(sql.FieldLTE(FieldBookingNumber, v))
}

// BookingNumberContains applies the Contains predicate on the "booking_number" field.
func BookingNumberContains(v string) predicate.Booking {
	return predicate.Booking(sql.FieldContains(FieldBookingNumber, v))
}

// BookingNumberHasPrefix applies the HasPrefix predicate on the "booking_number" field.
func BookingNumberHasPrefix(v string) predicate.Booking {
	return predicate.Booking(sql.FieldHasPrefix(FieldBookingNumber, v))
}

// BookingNumberHasSuffix applies the HasSuffix predicate on the "booking_number" field.
func BookingNumberHasSuffix(v string) predicate.Booking {
	return predicate.Booking(sql.FieldHasSuffix(FieldBookingNumber, v))
}

// BookingNumberEqualFold applies the EqualFold predicate on the "booking_number" field.
func BookingNumberEqualFold(v string) predicate.Booking {
	return predicate.Booking(sql.FieldEqualFold(FieldBookingNumber, v))
}

// BookingNumberContainsFold applies the ContainsFold predicate on the "booking_number" field.
func BookingNumberContainsFold(v string) predicate.Booking {
	return predicate.Booking(sql.FieldContainsFold(FieldBookingNumber, v))
}

// SmsNotificationEQ applies the EQ predicate on the "sms_notification" field.
func SmsNotificationEQ(v bool) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldSmsNotification, v))
}

// SmsNotificationNEQ applies the NEQ predicate on the "sms_notification" field.
func SmsNotificationNEQ(v bool) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldSmsNotification, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Booking {
	return predicate.Booking(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Booking {
	return predicate.Booking(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Booking {
	return predicate.Booking(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Booking {
	return predicate.Booking(sql.FieldNotIn(FieldStatus, vs...))
}

// HasPassengers applies the HasEdge predicate on the "passengers" edge.
func HasPassengers() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PassengersTable, PassengersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPassengersWith applies the HasEdge predicate on the "passengers" edge with a given conditions (other predicates).
func HasPassengersWith(preds ...predicate.Passenger) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newPassengersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLuggages applies the HasEdge predicate on the "luggages" edge.
func HasLuggages() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LuggagesTable, LuggagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLuggagesWith applies the HasEdge predicate on the "luggages" edge with a given conditions (other predicates).
func HasLuggagesWith(preds ...predicate.CustomerLuggage) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newLuggagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasContact applies the HasEdge predicate on the "contact" edge.
func HasContact() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ContactTable, ContactColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContactWith applies the HasEdge predicate on the "contact" edge with a given conditions (other predicates).
func HasContactWith(preds ...predicate.CustomerContact) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newContactStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTransaction applies the HasEdge predicate on the "transaction" edge.
func HasTransaction() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, TransactionTable, TransactionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionWith applies the HasEdge predicate on the "transaction" edge with a given conditions (other predicates).
func HasTransactionWith(preds ...predicate.Transaction) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newTransactionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTrip applies the HasEdge predicate on the "trip" edge.
func HasTrip() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TripTable, TripColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTripWith applies the HasEdge predicate on the "trip" edge with a given conditions (other predicates).
func HasTripWith(preds ...predicate.Trip) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newTripStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := newCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Booking) predicate.Booking {
	return predicate.Booking(sql.NotPredicates(p))
}
