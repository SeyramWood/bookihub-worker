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
	"github.com/SeyramWood/bookibus/ent/booking"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/customer"
	"github.com/SeyramWood/bookibus/ent/customercontact"
	"github.com/SeyramWood/bookibus/ent/customerluggage"
	"github.com/SeyramWood/bookibus/ent/passenger"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/transaction"
	"github.com/SeyramWood/bookibus/ent/trip"
)

// BookingUpdate is the builder for updating Booking entities.
type BookingUpdate struct {
	config
	hooks     []Hook
	mutation  *BookingMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the BookingUpdate builder.
func (bu *BookingUpdate) Where(ps ...predicate.Booking) *BookingUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BookingUpdate) SetUpdatedAt(t time.Time) *BookingUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetBookingNumber sets the "booking_number" field.
func (bu *BookingUpdate) SetBookingNumber(s string) *BookingUpdate {
	bu.mutation.SetBookingNumber(s)
	return bu
}

// SetSmsNotification sets the "sms_notification" field.
func (bu *BookingUpdate) SetSmsNotification(b bool) *BookingUpdate {
	bu.mutation.SetSmsNotification(b)
	return bu
}

// SetNillableSmsNotification sets the "sms_notification" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableSmsNotification(b *bool) *BookingUpdate {
	if b != nil {
		bu.SetSmsNotification(*b)
	}
	return bu
}

// SetStatus sets the "status" field.
func (bu *BookingUpdate) SetStatus(b booking.Status) *BookingUpdate {
	bu.mutation.SetStatus(b)
	return bu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableStatus(b *booking.Status) *BookingUpdate {
	if b != nil {
		bu.SetStatus(*b)
	}
	return bu
}

// AddPassengerIDs adds the "passengers" edge to the Passenger entity by IDs.
func (bu *BookingUpdate) AddPassengerIDs(ids ...int) *BookingUpdate {
	bu.mutation.AddPassengerIDs(ids...)
	return bu
}

// AddPassengers adds the "passengers" edges to the Passenger entity.
func (bu *BookingUpdate) AddPassengers(p ...*Passenger) *BookingUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.AddPassengerIDs(ids...)
}

// AddLuggageIDs adds the "luggages" edge to the CustomerLuggage entity by IDs.
func (bu *BookingUpdate) AddLuggageIDs(ids ...int) *BookingUpdate {
	bu.mutation.AddLuggageIDs(ids...)
	return bu
}

// AddLuggages adds the "luggages" edges to the CustomerLuggage entity.
func (bu *BookingUpdate) AddLuggages(c ...*CustomerLuggage) *BookingUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.AddLuggageIDs(ids...)
}

// SetContactID sets the "contact" edge to the CustomerContact entity by ID.
func (bu *BookingUpdate) SetContactID(id int) *BookingUpdate {
	bu.mutation.SetContactID(id)
	return bu
}

// SetNillableContactID sets the "contact" edge to the CustomerContact entity by ID if the given value is not nil.
func (bu *BookingUpdate) SetNillableContactID(id *int) *BookingUpdate {
	if id != nil {
		bu = bu.SetContactID(*id)
	}
	return bu
}

// SetContact sets the "contact" edge to the CustomerContact entity.
func (bu *BookingUpdate) SetContact(c *CustomerContact) *BookingUpdate {
	return bu.SetContactID(c.ID)
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (bu *BookingUpdate) SetTransactionID(id int) *BookingUpdate {
	bu.mutation.SetTransactionID(id)
	return bu
}

// SetNillableTransactionID sets the "transaction" edge to the Transaction entity by ID if the given value is not nil.
func (bu *BookingUpdate) SetNillableTransactionID(id *int) *BookingUpdate {
	if id != nil {
		bu = bu.SetTransactionID(*id)
	}
	return bu
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (bu *BookingUpdate) SetTransaction(t *Transaction) *BookingUpdate {
	return bu.SetTransactionID(t.ID)
}

// SetTripID sets the "trip" edge to the Trip entity by ID.
func (bu *BookingUpdate) SetTripID(id int) *BookingUpdate {
	bu.mutation.SetTripID(id)
	return bu
}

// SetNillableTripID sets the "trip" edge to the Trip entity by ID if the given value is not nil.
func (bu *BookingUpdate) SetNillableTripID(id *int) *BookingUpdate {
	if id != nil {
		bu = bu.SetTripID(*id)
	}
	return bu
}

// SetTrip sets the "trip" edge to the Trip entity.
func (bu *BookingUpdate) SetTrip(t *Trip) *BookingUpdate {
	return bu.SetTripID(t.ID)
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (bu *BookingUpdate) SetCompanyID(id int) *BookingUpdate {
	bu.mutation.SetCompanyID(id)
	return bu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (bu *BookingUpdate) SetNillableCompanyID(id *int) *BookingUpdate {
	if id != nil {
		bu = bu.SetCompanyID(*id)
	}
	return bu
}

// SetCompany sets the "company" edge to the Company entity.
func (bu *BookingUpdate) SetCompany(c *Company) *BookingUpdate {
	return bu.SetCompanyID(c.ID)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (bu *BookingUpdate) SetCustomerID(id int) *BookingUpdate {
	bu.mutation.SetCustomerID(id)
	return bu
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (bu *BookingUpdate) SetNillableCustomerID(id *int) *BookingUpdate {
	if id != nil {
		bu = bu.SetCustomerID(*id)
	}
	return bu
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (bu *BookingUpdate) SetCustomer(c *Customer) *BookingUpdate {
	return bu.SetCustomerID(c.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (bu *BookingUpdate) Mutation() *BookingMutation {
	return bu.mutation
}

// ClearPassengers clears all "passengers" edges to the Passenger entity.
func (bu *BookingUpdate) ClearPassengers() *BookingUpdate {
	bu.mutation.ClearPassengers()
	return bu
}

// RemovePassengerIDs removes the "passengers" edge to Passenger entities by IDs.
func (bu *BookingUpdate) RemovePassengerIDs(ids ...int) *BookingUpdate {
	bu.mutation.RemovePassengerIDs(ids...)
	return bu
}

// RemovePassengers removes "passengers" edges to Passenger entities.
func (bu *BookingUpdate) RemovePassengers(p ...*Passenger) *BookingUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.RemovePassengerIDs(ids...)
}

// ClearLuggages clears all "luggages" edges to the CustomerLuggage entity.
func (bu *BookingUpdate) ClearLuggages() *BookingUpdate {
	bu.mutation.ClearLuggages()
	return bu
}

// RemoveLuggageIDs removes the "luggages" edge to CustomerLuggage entities by IDs.
func (bu *BookingUpdate) RemoveLuggageIDs(ids ...int) *BookingUpdate {
	bu.mutation.RemoveLuggageIDs(ids...)
	return bu
}

// RemoveLuggages removes "luggages" edges to CustomerLuggage entities.
func (bu *BookingUpdate) RemoveLuggages(c ...*CustomerLuggage) *BookingUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.RemoveLuggageIDs(ids...)
}

// ClearContact clears the "contact" edge to the CustomerContact entity.
func (bu *BookingUpdate) ClearContact() *BookingUpdate {
	bu.mutation.ClearContact()
	return bu
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (bu *BookingUpdate) ClearTransaction() *BookingUpdate {
	bu.mutation.ClearTransaction()
	return bu
}

// ClearTrip clears the "trip" edge to the Trip entity.
func (bu *BookingUpdate) ClearTrip() *BookingUpdate {
	bu.mutation.ClearTrip()
	return bu
}

// ClearCompany clears the "company" edge to the Company entity.
func (bu *BookingUpdate) ClearCompany() *BookingUpdate {
	bu.mutation.ClearCompany()
	return bu
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (bu *BookingUpdate) ClearCustomer() *BookingUpdate {
	bu.mutation.ClearCustomer()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookingUpdate) Save(ctx context.Context) (int, error) {
	bu.defaults()
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookingUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookingUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookingUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BookingUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := booking.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookingUpdate) check() error {
	if v, ok := bu.mutation.BookingNumber(); ok {
		if err := booking.BookingNumberValidator(v); err != nil {
			return &ValidationError{Name: "booking_number", err: fmt.Errorf(`ent: validator failed for field "Booking.booking_number": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Status(); ok {
		if err := booking.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Booking.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (bu *BookingUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *BookingUpdate {
	bu.modifiers = append(bu.modifiers, modifiers...)
	return bu
}

func (bu *BookingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(booking.Table, booking.Columns, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(booking.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.BookingNumber(); ok {
		_spec.SetField(booking.FieldBookingNumber, field.TypeString, value)
	}
	if value, ok := bu.mutation.SmsNotification(); ok {
		_spec.SetField(booking.FieldSmsNotification, field.TypeBool, value)
	}
	if value, ok := bu.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeEnum, value)
	}
	if bu.mutation.PassengersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.PassengersTable,
			Columns: []string{booking.PassengersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passenger.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedPassengersIDs(); len(nodes) > 0 && !bu.mutation.PassengersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.PassengersTable,
			Columns: []string{booking.PassengersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passenger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.PassengersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.PassengersTable,
			Columns: []string{booking.PassengersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passenger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.LuggagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.LuggagesTable,
			Columns: []string{booking.LuggagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customerluggage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedLuggagesIDs(); len(nodes) > 0 && !bu.mutation.LuggagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.LuggagesTable,
			Columns: []string{booking.LuggagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customerluggage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.LuggagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.LuggagesTable,
			Columns: []string{booking.LuggagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customerluggage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.ContactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   booking.ContactTable,
			Columns: []string{booking.ContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customercontact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ContactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   booking.ContactTable,
			Columns: []string{booking.ContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customercontact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   booking.TransactionTable,
			Columns: []string{booking.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   booking.TransactionTable,
			Columns: []string{booking.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.TripCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.TripTable,
			Columns: []string{booking.TripColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.TripIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.TripTable,
			Columns: []string{booking.TripColumn},
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
	if bu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CompanyTable,
			Columns: []string{booking.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CompanyTable,
			Columns: []string{booking.CompanyColumn},
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
	if bu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CustomerTable,
			Columns: []string{booking.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CustomerTable,
			Columns: []string{booking.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(bu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookingUpdateOne is the builder for updating a single Booking entity.
type BookingUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *BookingMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BookingUpdateOne) SetUpdatedAt(t time.Time) *BookingUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetBookingNumber sets the "booking_number" field.
func (buo *BookingUpdateOne) SetBookingNumber(s string) *BookingUpdateOne {
	buo.mutation.SetBookingNumber(s)
	return buo
}

// SetSmsNotification sets the "sms_notification" field.
func (buo *BookingUpdateOne) SetSmsNotification(b bool) *BookingUpdateOne {
	buo.mutation.SetSmsNotification(b)
	return buo
}

// SetNillableSmsNotification sets the "sms_notification" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableSmsNotification(b *bool) *BookingUpdateOne {
	if b != nil {
		buo.SetSmsNotification(*b)
	}
	return buo
}

// SetStatus sets the "status" field.
func (buo *BookingUpdateOne) SetStatus(b booking.Status) *BookingUpdateOne {
	buo.mutation.SetStatus(b)
	return buo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableStatus(b *booking.Status) *BookingUpdateOne {
	if b != nil {
		buo.SetStatus(*b)
	}
	return buo
}

// AddPassengerIDs adds the "passengers" edge to the Passenger entity by IDs.
func (buo *BookingUpdateOne) AddPassengerIDs(ids ...int) *BookingUpdateOne {
	buo.mutation.AddPassengerIDs(ids...)
	return buo
}

// AddPassengers adds the "passengers" edges to the Passenger entity.
func (buo *BookingUpdateOne) AddPassengers(p ...*Passenger) *BookingUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.AddPassengerIDs(ids...)
}

// AddLuggageIDs adds the "luggages" edge to the CustomerLuggage entity by IDs.
func (buo *BookingUpdateOne) AddLuggageIDs(ids ...int) *BookingUpdateOne {
	buo.mutation.AddLuggageIDs(ids...)
	return buo
}

// AddLuggages adds the "luggages" edges to the CustomerLuggage entity.
func (buo *BookingUpdateOne) AddLuggages(c ...*CustomerLuggage) *BookingUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.AddLuggageIDs(ids...)
}

// SetContactID sets the "contact" edge to the CustomerContact entity by ID.
func (buo *BookingUpdateOne) SetContactID(id int) *BookingUpdateOne {
	buo.mutation.SetContactID(id)
	return buo
}

// SetNillableContactID sets the "contact" edge to the CustomerContact entity by ID if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableContactID(id *int) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetContactID(*id)
	}
	return buo
}

// SetContact sets the "contact" edge to the CustomerContact entity.
func (buo *BookingUpdateOne) SetContact(c *CustomerContact) *BookingUpdateOne {
	return buo.SetContactID(c.ID)
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (buo *BookingUpdateOne) SetTransactionID(id int) *BookingUpdateOne {
	buo.mutation.SetTransactionID(id)
	return buo
}

// SetNillableTransactionID sets the "transaction" edge to the Transaction entity by ID if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableTransactionID(id *int) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetTransactionID(*id)
	}
	return buo
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (buo *BookingUpdateOne) SetTransaction(t *Transaction) *BookingUpdateOne {
	return buo.SetTransactionID(t.ID)
}

// SetTripID sets the "trip" edge to the Trip entity by ID.
func (buo *BookingUpdateOne) SetTripID(id int) *BookingUpdateOne {
	buo.mutation.SetTripID(id)
	return buo
}

// SetNillableTripID sets the "trip" edge to the Trip entity by ID if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableTripID(id *int) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetTripID(*id)
	}
	return buo
}

// SetTrip sets the "trip" edge to the Trip entity.
func (buo *BookingUpdateOne) SetTrip(t *Trip) *BookingUpdateOne {
	return buo.SetTripID(t.ID)
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (buo *BookingUpdateOne) SetCompanyID(id int) *BookingUpdateOne {
	buo.mutation.SetCompanyID(id)
	return buo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableCompanyID(id *int) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetCompanyID(*id)
	}
	return buo
}

// SetCompany sets the "company" edge to the Company entity.
func (buo *BookingUpdateOne) SetCompany(c *Company) *BookingUpdateOne {
	return buo.SetCompanyID(c.ID)
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (buo *BookingUpdateOne) SetCustomerID(id int) *BookingUpdateOne {
	buo.mutation.SetCustomerID(id)
	return buo
}

// SetNillableCustomerID sets the "customer" edge to the Customer entity by ID if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableCustomerID(id *int) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetCustomerID(*id)
	}
	return buo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (buo *BookingUpdateOne) SetCustomer(c *Customer) *BookingUpdateOne {
	return buo.SetCustomerID(c.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (buo *BookingUpdateOne) Mutation() *BookingMutation {
	return buo.mutation
}

// ClearPassengers clears all "passengers" edges to the Passenger entity.
func (buo *BookingUpdateOne) ClearPassengers() *BookingUpdateOne {
	buo.mutation.ClearPassengers()
	return buo
}

// RemovePassengerIDs removes the "passengers" edge to Passenger entities by IDs.
func (buo *BookingUpdateOne) RemovePassengerIDs(ids ...int) *BookingUpdateOne {
	buo.mutation.RemovePassengerIDs(ids...)
	return buo
}

// RemovePassengers removes "passengers" edges to Passenger entities.
func (buo *BookingUpdateOne) RemovePassengers(p ...*Passenger) *BookingUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.RemovePassengerIDs(ids...)
}

// ClearLuggages clears all "luggages" edges to the CustomerLuggage entity.
func (buo *BookingUpdateOne) ClearLuggages() *BookingUpdateOne {
	buo.mutation.ClearLuggages()
	return buo
}

// RemoveLuggageIDs removes the "luggages" edge to CustomerLuggage entities by IDs.
func (buo *BookingUpdateOne) RemoveLuggageIDs(ids ...int) *BookingUpdateOne {
	buo.mutation.RemoveLuggageIDs(ids...)
	return buo
}

// RemoveLuggages removes "luggages" edges to CustomerLuggage entities.
func (buo *BookingUpdateOne) RemoveLuggages(c ...*CustomerLuggage) *BookingUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.RemoveLuggageIDs(ids...)
}

// ClearContact clears the "contact" edge to the CustomerContact entity.
func (buo *BookingUpdateOne) ClearContact() *BookingUpdateOne {
	buo.mutation.ClearContact()
	return buo
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (buo *BookingUpdateOne) ClearTransaction() *BookingUpdateOne {
	buo.mutation.ClearTransaction()
	return buo
}

// ClearTrip clears the "trip" edge to the Trip entity.
func (buo *BookingUpdateOne) ClearTrip() *BookingUpdateOne {
	buo.mutation.ClearTrip()
	return buo
}

// ClearCompany clears the "company" edge to the Company entity.
func (buo *BookingUpdateOne) ClearCompany() *BookingUpdateOne {
	buo.mutation.ClearCompany()
	return buo
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (buo *BookingUpdateOne) ClearCustomer() *BookingUpdateOne {
	buo.mutation.ClearCustomer()
	return buo
}

// Where appends a list predicates to the BookingUpdate builder.
func (buo *BookingUpdateOne) Where(ps ...predicate.Booking) *BookingUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookingUpdateOne) Select(field string, fields ...string) *BookingUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Booking entity.
func (buo *BookingUpdateOne) Save(ctx context.Context) (*Booking, error) {
	buo.defaults()
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookingUpdateOne) SaveX(ctx context.Context) *Booking {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookingUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookingUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BookingUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := booking.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookingUpdateOne) check() error {
	if v, ok := buo.mutation.BookingNumber(); ok {
		if err := booking.BookingNumberValidator(v); err != nil {
			return &ValidationError{Name: "booking_number", err: fmt.Errorf(`ent: validator failed for field "Booking.booking_number": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Status(); ok {
		if err := booking.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Booking.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (buo *BookingUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *BookingUpdateOne {
	buo.modifiers = append(buo.modifiers, modifiers...)
	return buo
}

func (buo *BookingUpdateOne) sqlSave(ctx context.Context) (_node *Booking, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(booking.Table, booking.Columns, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Booking.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, booking.FieldID)
		for _, f := range fields {
			if !booking.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != booking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(booking.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.BookingNumber(); ok {
		_spec.SetField(booking.FieldBookingNumber, field.TypeString, value)
	}
	if value, ok := buo.mutation.SmsNotification(); ok {
		_spec.SetField(booking.FieldSmsNotification, field.TypeBool, value)
	}
	if value, ok := buo.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeEnum, value)
	}
	if buo.mutation.PassengersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.PassengersTable,
			Columns: []string{booking.PassengersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passenger.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedPassengersIDs(); len(nodes) > 0 && !buo.mutation.PassengersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.PassengersTable,
			Columns: []string{booking.PassengersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passenger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.PassengersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.PassengersTable,
			Columns: []string{booking.PassengersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(passenger.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.LuggagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.LuggagesTable,
			Columns: []string{booking.LuggagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customerluggage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedLuggagesIDs(); len(nodes) > 0 && !buo.mutation.LuggagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.LuggagesTable,
			Columns: []string{booking.LuggagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customerluggage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.LuggagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.LuggagesTable,
			Columns: []string{booking.LuggagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customerluggage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.ContactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   booking.ContactTable,
			Columns: []string{booking.ContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customercontact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ContactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   booking.ContactTable,
			Columns: []string{booking.ContactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customercontact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   booking.TransactionTable,
			Columns: []string{booking.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   booking.TransactionTable,
			Columns: []string{booking.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.TripCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.TripTable,
			Columns: []string{booking.TripColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.TripIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.TripTable,
			Columns: []string{booking.TripColumn},
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
	if buo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CompanyTable,
			Columns: []string{booking.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CompanyTable,
			Columns: []string{booking.CompanyColumn},
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
	if buo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CustomerTable,
			Columns: []string{booking.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CustomerTable,
			Columns: []string{booking.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(buo.modifiers...)
	_node = &Booking{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
