// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/company"
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/route"
	"github.com/SeyramWood/ent/routestop"
	"github.com/SeyramWood/ent/trip"
)

// RouteUpdate is the builder for updating Route entities.
type RouteUpdate struct {
	config
	hooks     []Hook
	mutation  *RouteMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RouteUpdate builder.
func (ru *RouteUpdate) Where(ps ...predicate.Route) *RouteUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RouteUpdate) SetUpdatedAt(t time.Time) *RouteUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetFromLocation sets the "from_location" field.
func (ru *RouteUpdate) SetFromLocation(s string) *RouteUpdate {
	ru.mutation.SetFromLocation(s)
	return ru
}

// SetToLocation sets the "to_location" field.
func (ru *RouteUpdate) SetToLocation(s string) *RouteUpdate {
	ru.mutation.SetToLocation(s)
	return ru
}

// SetFromLatitude sets the "from_latitude" field.
func (ru *RouteUpdate) SetFromLatitude(f float64) *RouteUpdate {
	ru.mutation.ResetFromLatitude()
	ru.mutation.SetFromLatitude(f)
	return ru
}

// SetNillableFromLatitude sets the "from_latitude" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableFromLatitude(f *float64) *RouteUpdate {
	if f != nil {
		ru.SetFromLatitude(*f)
	}
	return ru
}

// AddFromLatitude adds f to the "from_latitude" field.
func (ru *RouteUpdate) AddFromLatitude(f float64) *RouteUpdate {
	ru.mutation.AddFromLatitude(f)
	return ru
}

// ClearFromLatitude clears the value of the "from_latitude" field.
func (ru *RouteUpdate) ClearFromLatitude() *RouteUpdate {
	ru.mutation.ClearFromLatitude()
	return ru
}

// SetFromLongitude sets the "from_longitude" field.
func (ru *RouteUpdate) SetFromLongitude(f float64) *RouteUpdate {
	ru.mutation.ResetFromLongitude()
	ru.mutation.SetFromLongitude(f)
	return ru
}

// SetNillableFromLongitude sets the "from_longitude" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableFromLongitude(f *float64) *RouteUpdate {
	if f != nil {
		ru.SetFromLongitude(*f)
	}
	return ru
}

// AddFromLongitude adds f to the "from_longitude" field.
func (ru *RouteUpdate) AddFromLongitude(f float64) *RouteUpdate {
	ru.mutation.AddFromLongitude(f)
	return ru
}

// ClearFromLongitude clears the value of the "from_longitude" field.
func (ru *RouteUpdate) ClearFromLongitude() *RouteUpdate {
	ru.mutation.ClearFromLongitude()
	return ru
}

// SetToLatitude sets the "to_latitude" field.
func (ru *RouteUpdate) SetToLatitude(f float64) *RouteUpdate {
	ru.mutation.ResetToLatitude()
	ru.mutation.SetToLatitude(f)
	return ru
}

// SetNillableToLatitude sets the "to_latitude" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableToLatitude(f *float64) *RouteUpdate {
	if f != nil {
		ru.SetToLatitude(*f)
	}
	return ru
}

// AddToLatitude adds f to the "to_latitude" field.
func (ru *RouteUpdate) AddToLatitude(f float64) *RouteUpdate {
	ru.mutation.AddToLatitude(f)
	return ru
}

// ClearToLatitude clears the value of the "to_latitude" field.
func (ru *RouteUpdate) ClearToLatitude() *RouteUpdate {
	ru.mutation.ClearToLatitude()
	return ru
}

// SetToLongitude sets the "to_longitude" field.
func (ru *RouteUpdate) SetToLongitude(f float64) *RouteUpdate {
	ru.mutation.ResetToLongitude()
	ru.mutation.SetToLongitude(f)
	return ru
}

// SetNillableToLongitude sets the "to_longitude" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableToLongitude(f *float64) *RouteUpdate {
	if f != nil {
		ru.SetToLongitude(*f)
	}
	return ru
}

// AddToLongitude adds f to the "to_longitude" field.
func (ru *RouteUpdate) AddToLongitude(f float64) *RouteUpdate {
	ru.mutation.AddToLongitude(f)
	return ru
}

// ClearToLongitude clears the value of the "to_longitude" field.
func (ru *RouteUpdate) ClearToLongitude() *RouteUpdate {
	ru.mutation.ClearToLongitude()
	return ru
}

// SetRate sets the "rate" field.
func (ru *RouteUpdate) SetRate(f float64) *RouteUpdate {
	ru.mutation.ResetRate()
	ru.mutation.SetRate(f)
	return ru
}

// SetNillableRate sets the "rate" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableRate(f *float64) *RouteUpdate {
	if f != nil {
		ru.SetRate(*f)
	}
	return ru
}

// AddRate adds f to the "rate" field.
func (ru *RouteUpdate) AddRate(f float64) *RouteUpdate {
	ru.mutation.AddRate(f)
	return ru
}

// SetDiscount sets the "discount" field.
func (ru *RouteUpdate) SetDiscount(f float32) *RouteUpdate {
	ru.mutation.ResetDiscount()
	ru.mutation.SetDiscount(f)
	return ru
}

// SetNillableDiscount sets the "discount" field if the given value is not nil.
func (ru *RouteUpdate) SetNillableDiscount(f *float32) *RouteUpdate {
	if f != nil {
		ru.SetDiscount(*f)
	}
	return ru
}

// AddDiscount adds f to the "discount" field.
func (ru *RouteUpdate) AddDiscount(f float32) *RouteUpdate {
	ru.mutation.AddDiscount(f)
	return ru
}

// SetPopularity sets the "popularity" field.
func (ru *RouteUpdate) SetPopularity(i int) *RouteUpdate {
	ru.mutation.ResetPopularity()
	ru.mutation.SetPopularity(i)
	return ru
}

// SetNillablePopularity sets the "popularity" field if the given value is not nil.
func (ru *RouteUpdate) SetNillablePopularity(i *int) *RouteUpdate {
	if i != nil {
		ru.SetPopularity(*i)
	}
	return ru
}

// AddPopularity adds i to the "popularity" field.
func (ru *RouteUpdate) AddPopularity(i int) *RouteUpdate {
	ru.mutation.AddPopularity(i)
	return ru
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (ru *RouteUpdate) SetCompanyID(id int) *RouteUpdate {
	ru.mutation.SetCompanyID(id)
	return ru
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (ru *RouteUpdate) SetNillableCompanyID(id *int) *RouteUpdate {
	if id != nil {
		ru = ru.SetCompanyID(*id)
	}
	return ru
}

// SetCompany sets the "company" edge to the Company entity.
func (ru *RouteUpdate) SetCompany(c *Company) *RouteUpdate {
	return ru.SetCompanyID(c.ID)
}

// AddStopIDs adds the "stops" edge to the RouteStop entity by IDs.
func (ru *RouteUpdate) AddStopIDs(ids ...int) *RouteUpdate {
	ru.mutation.AddStopIDs(ids...)
	return ru
}

// AddStops adds the "stops" edges to the RouteStop entity.
func (ru *RouteUpdate) AddStops(r ...*RouteStop) *RouteUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddStopIDs(ids...)
}

// AddTripIDs adds the "trips" edge to the Trip entity by IDs.
func (ru *RouteUpdate) AddTripIDs(ids ...int) *RouteUpdate {
	ru.mutation.AddTripIDs(ids...)
	return ru
}

// AddTrips adds the "trips" edges to the Trip entity.
func (ru *RouteUpdate) AddTrips(t ...*Trip) *RouteUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.AddTripIDs(ids...)
}

// Mutation returns the RouteMutation object of the builder.
func (ru *RouteUpdate) Mutation() *RouteMutation {
	return ru.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (ru *RouteUpdate) ClearCompany() *RouteUpdate {
	ru.mutation.ClearCompany()
	return ru
}

// ClearStops clears all "stops" edges to the RouteStop entity.
func (ru *RouteUpdate) ClearStops() *RouteUpdate {
	ru.mutation.ClearStops()
	return ru
}

// RemoveStopIDs removes the "stops" edge to RouteStop entities by IDs.
func (ru *RouteUpdate) RemoveStopIDs(ids ...int) *RouteUpdate {
	ru.mutation.RemoveStopIDs(ids...)
	return ru
}

// RemoveStops removes "stops" edges to RouteStop entities.
func (ru *RouteUpdate) RemoveStops(r ...*RouteStop) *RouteUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveStopIDs(ids...)
}

// ClearTrips clears all "trips" edges to the Trip entity.
func (ru *RouteUpdate) ClearTrips() *RouteUpdate {
	ru.mutation.ClearTrips()
	return ru
}

// RemoveTripIDs removes the "trips" edge to Trip entities by IDs.
func (ru *RouteUpdate) RemoveTripIDs(ids ...int) *RouteUpdate {
	ru.mutation.RemoveTripIDs(ids...)
	return ru
}

// RemoveTrips removes "trips" edges to Trip entities.
func (ru *RouteUpdate) RemoveTrips(t ...*Trip) *RouteUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.RemoveTripIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RouteUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RouteUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RouteUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RouteUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RouteUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := route.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RouteUpdate) check() error {
	if v, ok := ru.mutation.FromLocation(); ok {
		if err := route.FromLocationValidator(v); err != nil {
			return &ValidationError{Name: "from_location", err: fmt.Errorf(`ent: validator failed for field "Route.from_location": %w`, err)}
		}
	}
	if v, ok := ru.mutation.ToLocation(); ok {
		if err := route.ToLocationValidator(v); err != nil {
			return &ValidationError{Name: "to_location", err: fmt.Errorf(`ent: validator failed for field "Route.to_location": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *RouteUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RouteUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *RouteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(route.Table, route.Columns, sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(route.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.FromLocation(); ok {
		_spec.SetField(route.FieldFromLocation, field.TypeString, value)
	}
	if value, ok := ru.mutation.ToLocation(); ok {
		_spec.SetField(route.FieldToLocation, field.TypeString, value)
	}
	if value, ok := ru.mutation.FromLatitude(); ok {
		_spec.SetField(route.FieldFromLatitude, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.AddedFromLatitude(); ok {
		_spec.AddField(route.FieldFromLatitude, field.TypeFloat64, value)
	}
	if ru.mutation.FromLatitudeCleared() {
		_spec.ClearField(route.FieldFromLatitude, field.TypeFloat64)
	}
	if value, ok := ru.mutation.FromLongitude(); ok {
		_spec.SetField(route.FieldFromLongitude, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.AddedFromLongitude(); ok {
		_spec.AddField(route.FieldFromLongitude, field.TypeFloat64, value)
	}
	if ru.mutation.FromLongitudeCleared() {
		_spec.ClearField(route.FieldFromLongitude, field.TypeFloat64)
	}
	if value, ok := ru.mutation.ToLatitude(); ok {
		_spec.SetField(route.FieldToLatitude, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.AddedToLatitude(); ok {
		_spec.AddField(route.FieldToLatitude, field.TypeFloat64, value)
	}
	if ru.mutation.ToLatitudeCleared() {
		_spec.ClearField(route.FieldToLatitude, field.TypeFloat64)
	}
	if value, ok := ru.mutation.ToLongitude(); ok {
		_spec.SetField(route.FieldToLongitude, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.AddedToLongitude(); ok {
		_spec.AddField(route.FieldToLongitude, field.TypeFloat64, value)
	}
	if ru.mutation.ToLongitudeCleared() {
		_spec.ClearField(route.FieldToLongitude, field.TypeFloat64)
	}
	if value, ok := ru.mutation.Rate(); ok {
		_spec.SetField(route.FieldRate, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.AddedRate(); ok {
		_spec.AddField(route.FieldRate, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.Discount(); ok {
		_spec.SetField(route.FieldDiscount, field.TypeFloat32, value)
	}
	if value, ok := ru.mutation.AddedDiscount(); ok {
		_spec.AddField(route.FieldDiscount, field.TypeFloat32, value)
	}
	if value, ok := ru.mutation.Popularity(); ok {
		_spec.SetField(route.FieldPopularity, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedPopularity(); ok {
		_spec.AddField(route.FieldPopularity, field.TypeInt, value)
	}
	if ru.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.CompanyTable,
			Columns: []string{route.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.CompanyTable,
			Columns: []string{route.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.StopsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.StopsTable,
			Columns: []string{route.StopsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routestop.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedStopsIDs(); len(nodes) > 0 && !ru.mutation.StopsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.StopsTable,
			Columns: []string{route.StopsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routestop.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.StopsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.StopsTable,
			Columns: []string{route.StopsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routestop.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.TripsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.TripsTable,
			Columns: []string{route.TripsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedTripsIDs(); len(nodes) > 0 && !ru.mutation.TripsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.TripsTable,
			Columns: []string{route.TripsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.TripsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.TripsTable,
			Columns: []string{route.TripsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{route.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RouteUpdateOne is the builder for updating a single Route entity.
type RouteUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RouteMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RouteUpdateOne) SetUpdatedAt(t time.Time) *RouteUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetFromLocation sets the "from_location" field.
func (ruo *RouteUpdateOne) SetFromLocation(s string) *RouteUpdateOne {
	ruo.mutation.SetFromLocation(s)
	return ruo
}

// SetToLocation sets the "to_location" field.
func (ruo *RouteUpdateOne) SetToLocation(s string) *RouteUpdateOne {
	ruo.mutation.SetToLocation(s)
	return ruo
}

// SetFromLatitude sets the "from_latitude" field.
func (ruo *RouteUpdateOne) SetFromLatitude(f float64) *RouteUpdateOne {
	ruo.mutation.ResetFromLatitude()
	ruo.mutation.SetFromLatitude(f)
	return ruo
}

// SetNillableFromLatitude sets the "from_latitude" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableFromLatitude(f *float64) *RouteUpdateOne {
	if f != nil {
		ruo.SetFromLatitude(*f)
	}
	return ruo
}

// AddFromLatitude adds f to the "from_latitude" field.
func (ruo *RouteUpdateOne) AddFromLatitude(f float64) *RouteUpdateOne {
	ruo.mutation.AddFromLatitude(f)
	return ruo
}

// ClearFromLatitude clears the value of the "from_latitude" field.
func (ruo *RouteUpdateOne) ClearFromLatitude() *RouteUpdateOne {
	ruo.mutation.ClearFromLatitude()
	return ruo
}

// SetFromLongitude sets the "from_longitude" field.
func (ruo *RouteUpdateOne) SetFromLongitude(f float64) *RouteUpdateOne {
	ruo.mutation.ResetFromLongitude()
	ruo.mutation.SetFromLongitude(f)
	return ruo
}

// SetNillableFromLongitude sets the "from_longitude" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableFromLongitude(f *float64) *RouteUpdateOne {
	if f != nil {
		ruo.SetFromLongitude(*f)
	}
	return ruo
}

// AddFromLongitude adds f to the "from_longitude" field.
func (ruo *RouteUpdateOne) AddFromLongitude(f float64) *RouteUpdateOne {
	ruo.mutation.AddFromLongitude(f)
	return ruo
}

// ClearFromLongitude clears the value of the "from_longitude" field.
func (ruo *RouteUpdateOne) ClearFromLongitude() *RouteUpdateOne {
	ruo.mutation.ClearFromLongitude()
	return ruo
}

// SetToLatitude sets the "to_latitude" field.
func (ruo *RouteUpdateOne) SetToLatitude(f float64) *RouteUpdateOne {
	ruo.mutation.ResetToLatitude()
	ruo.mutation.SetToLatitude(f)
	return ruo
}

// SetNillableToLatitude sets the "to_latitude" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableToLatitude(f *float64) *RouteUpdateOne {
	if f != nil {
		ruo.SetToLatitude(*f)
	}
	return ruo
}

// AddToLatitude adds f to the "to_latitude" field.
func (ruo *RouteUpdateOne) AddToLatitude(f float64) *RouteUpdateOne {
	ruo.mutation.AddToLatitude(f)
	return ruo
}

// ClearToLatitude clears the value of the "to_latitude" field.
func (ruo *RouteUpdateOne) ClearToLatitude() *RouteUpdateOne {
	ruo.mutation.ClearToLatitude()
	return ruo
}

// SetToLongitude sets the "to_longitude" field.
func (ruo *RouteUpdateOne) SetToLongitude(f float64) *RouteUpdateOne {
	ruo.mutation.ResetToLongitude()
	ruo.mutation.SetToLongitude(f)
	return ruo
}

// SetNillableToLongitude sets the "to_longitude" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableToLongitude(f *float64) *RouteUpdateOne {
	if f != nil {
		ruo.SetToLongitude(*f)
	}
	return ruo
}

// AddToLongitude adds f to the "to_longitude" field.
func (ruo *RouteUpdateOne) AddToLongitude(f float64) *RouteUpdateOne {
	ruo.mutation.AddToLongitude(f)
	return ruo
}

// ClearToLongitude clears the value of the "to_longitude" field.
func (ruo *RouteUpdateOne) ClearToLongitude() *RouteUpdateOne {
	ruo.mutation.ClearToLongitude()
	return ruo
}

// SetRate sets the "rate" field.
func (ruo *RouteUpdateOne) SetRate(f float64) *RouteUpdateOne {
	ruo.mutation.ResetRate()
	ruo.mutation.SetRate(f)
	return ruo
}

// SetNillableRate sets the "rate" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableRate(f *float64) *RouteUpdateOne {
	if f != nil {
		ruo.SetRate(*f)
	}
	return ruo
}

// AddRate adds f to the "rate" field.
func (ruo *RouteUpdateOne) AddRate(f float64) *RouteUpdateOne {
	ruo.mutation.AddRate(f)
	return ruo
}

// SetDiscount sets the "discount" field.
func (ruo *RouteUpdateOne) SetDiscount(f float32) *RouteUpdateOne {
	ruo.mutation.ResetDiscount()
	ruo.mutation.SetDiscount(f)
	return ruo
}

// SetNillableDiscount sets the "discount" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableDiscount(f *float32) *RouteUpdateOne {
	if f != nil {
		ruo.SetDiscount(*f)
	}
	return ruo
}

// AddDiscount adds f to the "discount" field.
func (ruo *RouteUpdateOne) AddDiscount(f float32) *RouteUpdateOne {
	ruo.mutation.AddDiscount(f)
	return ruo
}

// SetPopularity sets the "popularity" field.
func (ruo *RouteUpdateOne) SetPopularity(i int) *RouteUpdateOne {
	ruo.mutation.ResetPopularity()
	ruo.mutation.SetPopularity(i)
	return ruo
}

// SetNillablePopularity sets the "popularity" field if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillablePopularity(i *int) *RouteUpdateOne {
	if i != nil {
		ruo.SetPopularity(*i)
	}
	return ruo
}

// AddPopularity adds i to the "popularity" field.
func (ruo *RouteUpdateOne) AddPopularity(i int) *RouteUpdateOne {
	ruo.mutation.AddPopularity(i)
	return ruo
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (ruo *RouteUpdateOne) SetCompanyID(id int) *RouteUpdateOne {
	ruo.mutation.SetCompanyID(id)
	return ruo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (ruo *RouteUpdateOne) SetNillableCompanyID(id *int) *RouteUpdateOne {
	if id != nil {
		ruo = ruo.SetCompanyID(*id)
	}
	return ruo
}

// SetCompany sets the "company" edge to the Company entity.
func (ruo *RouteUpdateOne) SetCompany(c *Company) *RouteUpdateOne {
	return ruo.SetCompanyID(c.ID)
}

// AddStopIDs adds the "stops" edge to the RouteStop entity by IDs.
func (ruo *RouteUpdateOne) AddStopIDs(ids ...int) *RouteUpdateOne {
	ruo.mutation.AddStopIDs(ids...)
	return ruo
}

// AddStops adds the "stops" edges to the RouteStop entity.
func (ruo *RouteUpdateOne) AddStops(r ...*RouteStop) *RouteUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddStopIDs(ids...)
}

// AddTripIDs adds the "trips" edge to the Trip entity by IDs.
func (ruo *RouteUpdateOne) AddTripIDs(ids ...int) *RouteUpdateOne {
	ruo.mutation.AddTripIDs(ids...)
	return ruo
}

// AddTrips adds the "trips" edges to the Trip entity.
func (ruo *RouteUpdateOne) AddTrips(t ...*Trip) *RouteUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.AddTripIDs(ids...)
}

// Mutation returns the RouteMutation object of the builder.
func (ruo *RouteUpdateOne) Mutation() *RouteMutation {
	return ruo.mutation
}

// ClearCompany clears the "company" edge to the Company entity.
func (ruo *RouteUpdateOne) ClearCompany() *RouteUpdateOne {
	ruo.mutation.ClearCompany()
	return ruo
}

// ClearStops clears all "stops" edges to the RouteStop entity.
func (ruo *RouteUpdateOne) ClearStops() *RouteUpdateOne {
	ruo.mutation.ClearStops()
	return ruo
}

// RemoveStopIDs removes the "stops" edge to RouteStop entities by IDs.
func (ruo *RouteUpdateOne) RemoveStopIDs(ids ...int) *RouteUpdateOne {
	ruo.mutation.RemoveStopIDs(ids...)
	return ruo
}

// RemoveStops removes "stops" edges to RouteStop entities.
func (ruo *RouteUpdateOne) RemoveStops(r ...*RouteStop) *RouteUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveStopIDs(ids...)
}

// ClearTrips clears all "trips" edges to the Trip entity.
func (ruo *RouteUpdateOne) ClearTrips() *RouteUpdateOne {
	ruo.mutation.ClearTrips()
	return ruo
}

// RemoveTripIDs removes the "trips" edge to Trip entities by IDs.
func (ruo *RouteUpdateOne) RemoveTripIDs(ids ...int) *RouteUpdateOne {
	ruo.mutation.RemoveTripIDs(ids...)
	return ruo
}

// RemoveTrips removes "trips" edges to Trip entities.
func (ruo *RouteUpdateOne) RemoveTrips(t ...*Trip) *RouteUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.RemoveTripIDs(ids...)
}

// Where appends a list predicates to the RouteUpdate builder.
func (ruo *RouteUpdateOne) Where(ps ...predicate.Route) *RouteUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RouteUpdateOne) Select(field string, fields ...string) *RouteUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Route entity.
func (ruo *RouteUpdateOne) Save(ctx context.Context) (*Route, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RouteUpdateOne) SaveX(ctx context.Context) *Route {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RouteUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RouteUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RouteUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := route.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RouteUpdateOne) check() error {
	if v, ok := ruo.mutation.FromLocation(); ok {
		if err := route.FromLocationValidator(v); err != nil {
			return &ValidationError{Name: "from_location", err: fmt.Errorf(`ent: validator failed for field "Route.from_location": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.ToLocation(); ok {
		if err := route.ToLocationValidator(v); err != nil {
			return &ValidationError{Name: "to_location", err: fmt.Errorf(`ent: validator failed for field "Route.to_location": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *RouteUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RouteUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *RouteUpdateOne) sqlSave(ctx context.Context) (_node *Route, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(route.Table, route.Columns, sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Route.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, route.FieldID)
		for _, f := range fields {
			if !route.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != route.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(route.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.FromLocation(); ok {
		_spec.SetField(route.FieldFromLocation, field.TypeString, value)
	}
	if value, ok := ruo.mutation.ToLocation(); ok {
		_spec.SetField(route.FieldToLocation, field.TypeString, value)
	}
	if value, ok := ruo.mutation.FromLatitude(); ok {
		_spec.SetField(route.FieldFromLatitude, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.AddedFromLatitude(); ok {
		_spec.AddField(route.FieldFromLatitude, field.TypeFloat64, value)
	}
	if ruo.mutation.FromLatitudeCleared() {
		_spec.ClearField(route.FieldFromLatitude, field.TypeFloat64)
	}
	if value, ok := ruo.mutation.FromLongitude(); ok {
		_spec.SetField(route.FieldFromLongitude, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.AddedFromLongitude(); ok {
		_spec.AddField(route.FieldFromLongitude, field.TypeFloat64, value)
	}
	if ruo.mutation.FromLongitudeCleared() {
		_spec.ClearField(route.FieldFromLongitude, field.TypeFloat64)
	}
	if value, ok := ruo.mutation.ToLatitude(); ok {
		_spec.SetField(route.FieldToLatitude, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.AddedToLatitude(); ok {
		_spec.AddField(route.FieldToLatitude, field.TypeFloat64, value)
	}
	if ruo.mutation.ToLatitudeCleared() {
		_spec.ClearField(route.FieldToLatitude, field.TypeFloat64)
	}
	if value, ok := ruo.mutation.ToLongitude(); ok {
		_spec.SetField(route.FieldToLongitude, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.AddedToLongitude(); ok {
		_spec.AddField(route.FieldToLongitude, field.TypeFloat64, value)
	}
	if ruo.mutation.ToLongitudeCleared() {
		_spec.ClearField(route.FieldToLongitude, field.TypeFloat64)
	}
	if value, ok := ruo.mutation.Rate(); ok {
		_spec.SetField(route.FieldRate, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.AddedRate(); ok {
		_spec.AddField(route.FieldRate, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.Discount(); ok {
		_spec.SetField(route.FieldDiscount, field.TypeFloat32, value)
	}
	if value, ok := ruo.mutation.AddedDiscount(); ok {
		_spec.AddField(route.FieldDiscount, field.TypeFloat32, value)
	}
	if value, ok := ruo.mutation.Popularity(); ok {
		_spec.SetField(route.FieldPopularity, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedPopularity(); ok {
		_spec.AddField(route.FieldPopularity, field.TypeInt, value)
	}
	if ruo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.CompanyTable,
			Columns: []string{route.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   route.CompanyTable,
			Columns: []string{route.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.StopsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.StopsTable,
			Columns: []string{route.StopsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routestop.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedStopsIDs(); len(nodes) > 0 && !ruo.mutation.StopsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.StopsTable,
			Columns: []string{route.StopsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routestop.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.StopsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.StopsTable,
			Columns: []string{route.StopsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routestop.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.TripsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.TripsTable,
			Columns: []string{route.TripsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedTripsIDs(); len(nodes) > 0 && !ruo.mutation.TripsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.TripsTable,
			Columns: []string{route.TripsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.TripsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   route.TripsTable,
			Columns: []string{route.TripsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Route{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{route.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}