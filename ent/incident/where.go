// Code generated by ent, DO NOT EDIT.

package incident

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/SeyramWood/bookibus/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Incident {
	return predicate.Incident(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Incident {
	return predicate.Incident(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Incident {
	return predicate.Incident(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Incident {
	return predicate.Incident(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Incident {
	return predicate.Incident(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Incident {
	return predicate.Incident(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Incident {
	return predicate.Incident(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldUpdatedAt, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldTime, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldLocation, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldDescription, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldType, v))
}

// Audio applies equality check predicate on the "audio" field. It's identical to AudioEQ.
func Audio(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldAudio, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldLTE(FieldUpdatedAt, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v time.Time) predicate.Incident {
	return predicate.Incident(sql.FieldLTE(FieldTime, v))
}

// TimeIsNil applies the IsNil predicate on the "time" field.
func TimeIsNil() predicate.Incident {
	return predicate.Incident(sql.FieldIsNull(FieldTime))
}

// TimeNotNil applies the NotNil predicate on the "time" field.
func TimeNotNil() predicate.Incident {
	return predicate.Incident(sql.FieldNotNull(FieldTime))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.Incident {
	return predicate.Incident(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.Incident {
	return predicate.Incident(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.Incident {
	return predicate.Incident(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.Incident {
	return predicate.Incident(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.Incident {
	return predicate.Incident(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.Incident {
	return predicate.Incident(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.Incident {
	return predicate.Incident(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.Incident {
	return predicate.Incident(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.Incident {
	return predicate.Incident(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.Incident {
	return predicate.Incident(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.Incident {
	return predicate.Incident(sql.FieldContainsFold(FieldLocation, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Incident {
	return predicate.Incident(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Incident {
	return predicate.Incident(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Incident {
	return predicate.Incident(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Incident {
	return predicate.Incident(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Incident {
	return predicate.Incident(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Incident {
	return predicate.Incident(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Incident {
	return predicate.Incident(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Incident {
	return predicate.Incident(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Incident {
	return predicate.Incident(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Incident {
	return predicate.Incident(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Incident {
	return predicate.Incident(sql.FieldContainsFold(FieldDescription, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Incident {
	return predicate.Incident(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Incident {
	return predicate.Incident(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Incident {
	return predicate.Incident(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Incident {
	return predicate.Incident(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Incident {
	return predicate.Incident(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Incident {
	return predicate.Incident(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Incident {
	return predicate.Incident(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Incident {
	return predicate.Incident(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Incident {
	return predicate.Incident(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Incident {
	return predicate.Incident(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Incident {
	return predicate.Incident(sql.FieldContainsFold(FieldType, v))
}

// AudioEQ applies the EQ predicate on the "audio" field.
func AudioEQ(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldAudio, v))
}

// AudioNEQ applies the NEQ predicate on the "audio" field.
func AudioNEQ(v string) predicate.Incident {
	return predicate.Incident(sql.FieldNEQ(FieldAudio, v))
}

// AudioIn applies the In predicate on the "audio" field.
func AudioIn(vs ...string) predicate.Incident {
	return predicate.Incident(sql.FieldIn(FieldAudio, vs...))
}

// AudioNotIn applies the NotIn predicate on the "audio" field.
func AudioNotIn(vs ...string) predicate.Incident {
	return predicate.Incident(sql.FieldNotIn(FieldAudio, vs...))
}

// AudioGT applies the GT predicate on the "audio" field.
func AudioGT(v string) predicate.Incident {
	return predicate.Incident(sql.FieldGT(FieldAudio, v))
}

// AudioGTE applies the GTE predicate on the "audio" field.
func AudioGTE(v string) predicate.Incident {
	return predicate.Incident(sql.FieldGTE(FieldAudio, v))
}

// AudioLT applies the LT predicate on the "audio" field.
func AudioLT(v string) predicate.Incident {
	return predicate.Incident(sql.FieldLT(FieldAudio, v))
}

// AudioLTE applies the LTE predicate on the "audio" field.
func AudioLTE(v string) predicate.Incident {
	return predicate.Incident(sql.FieldLTE(FieldAudio, v))
}

// AudioContains applies the Contains predicate on the "audio" field.
func AudioContains(v string) predicate.Incident {
	return predicate.Incident(sql.FieldContains(FieldAudio, v))
}

// AudioHasPrefix applies the HasPrefix predicate on the "audio" field.
func AudioHasPrefix(v string) predicate.Incident {
	return predicate.Incident(sql.FieldHasPrefix(FieldAudio, v))
}

// AudioHasSuffix applies the HasSuffix predicate on the "audio" field.
func AudioHasSuffix(v string) predicate.Incident {
	return predicate.Incident(sql.FieldHasSuffix(FieldAudio, v))
}

// AudioIsNil applies the IsNil predicate on the "audio" field.
func AudioIsNil() predicate.Incident {
	return predicate.Incident(sql.FieldIsNull(FieldAudio))
}

// AudioNotNil applies the NotNil predicate on the "audio" field.
func AudioNotNil() predicate.Incident {
	return predicate.Incident(sql.FieldNotNull(FieldAudio))
}

// AudioEqualFold applies the EqualFold predicate on the "audio" field.
func AudioEqualFold(v string) predicate.Incident {
	return predicate.Incident(sql.FieldEqualFold(FieldAudio, v))
}

// AudioContainsFold applies the ContainsFold predicate on the "audio" field.
func AudioContainsFold(v string) predicate.Incident {
	return predicate.Incident(sql.FieldContainsFold(FieldAudio, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Incident {
	return predicate.Incident(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Incident {
	return predicate.Incident(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Incident {
	return predicate.Incident(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Incident {
	return predicate.Incident(sql.FieldNotIn(FieldStatus, vs...))
}

// HasImages applies the HasEdge predicate on the "images" edge.
func HasImages() predicate.Incident {
	return predicate.Incident(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImagesWith applies the HasEdge predicate on the "images" edge with a given conditions (other predicates).
func HasImagesWith(preds ...predicate.IncidentImage) predicate.Incident {
	return predicate.Incident(func(s *sql.Selector) {
		step := newImagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTrip applies the HasEdge predicate on the "trip" edge.
func HasTrip() predicate.Incident {
	return predicate.Incident(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TripTable, TripColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTripWith applies the HasEdge predicate on the "trip" edge with a given conditions (other predicates).
func HasTripWith(preds ...predicate.Trip) predicate.Incident {
	return predicate.Incident(func(s *sql.Selector) {
		step := newTripStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.Incident {
	return predicate.Incident(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.Incident {
	return predicate.Incident(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDriver applies the HasEdge predicate on the "driver" edge.
func HasDriver() predicate.Incident {
	return predicate.Incident(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DriverTable, DriverColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDriverWith applies the HasEdge predicate on the "driver" edge with a given conditions (other predicates).
func HasDriverWith(preds ...predicate.CompanyUser) predicate.Incident {
	return predicate.Incident(func(s *sql.Selector) {
		step := newDriverStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Incident) predicate.Incident {
	return predicate.Incident(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Incident) predicate.Incident {
	return predicate.Incident(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Incident) predicate.Incident {
	return predicate.Incident(sql.NotPredicates(p))
}
