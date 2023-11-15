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
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/parcel"
	"github.com/SeyramWood/bookibus/ent/parcelimage"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/trip"
)

// ParcelUpdate is the builder for updating Parcel entities.
type ParcelUpdate struct {
	config
	hooks     []Hook
	mutation  *ParcelMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ParcelUpdate builder.
func (pu *ParcelUpdate) Where(ps ...predicate.Parcel) *ParcelUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ParcelUpdate) SetUpdatedAt(t time.Time) *ParcelUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetParcelCode sets the "parcel_code" field.
func (pu *ParcelUpdate) SetParcelCode(s string) *ParcelUpdate {
	pu.mutation.SetParcelCode(s)
	return pu
}

// SetType sets the "type" field.
func (pu *ParcelUpdate) SetType(s string) *ParcelUpdate {
	pu.mutation.SetType(s)
	return pu
}

// SetSenderName sets the "sender_name" field.
func (pu *ParcelUpdate) SetSenderName(s string) *ParcelUpdate {
	pu.mutation.SetSenderName(s)
	return pu
}

// SetSenderPhone sets the "sender_phone" field.
func (pu *ParcelUpdate) SetSenderPhone(s string) *ParcelUpdate {
	pu.mutation.SetSenderPhone(s)
	return pu
}

// SetSenderEmail sets the "sender_email" field.
func (pu *ParcelUpdate) SetSenderEmail(s string) *ParcelUpdate {
	pu.mutation.SetSenderEmail(s)
	return pu
}

// SetRecipientName sets the "recipient_name" field.
func (pu *ParcelUpdate) SetRecipientName(s string) *ParcelUpdate {
	pu.mutation.SetRecipientName(s)
	return pu
}

// SetRecipientPhone sets the "recipient_phone" field.
func (pu *ParcelUpdate) SetRecipientPhone(s string) *ParcelUpdate {
	pu.mutation.SetRecipientPhone(s)
	return pu
}

// SetRecipientLocation sets the "recipient_location" field.
func (pu *ParcelUpdate) SetRecipientLocation(s string) *ParcelUpdate {
	pu.mutation.SetRecipientLocation(s)
	return pu
}

// SetWeight sets the "weight" field.
func (pu *ParcelUpdate) SetWeight(f float32) *ParcelUpdate {
	pu.mutation.ResetWeight()
	pu.mutation.SetWeight(f)
	return pu
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (pu *ParcelUpdate) SetNillableWeight(f *float32) *ParcelUpdate {
	if f != nil {
		pu.SetWeight(*f)
	}
	return pu
}

// AddWeight adds f to the "weight" field.
func (pu *ParcelUpdate) AddWeight(f float32) *ParcelUpdate {
	pu.mutation.AddWeight(f)
	return pu
}

// ClearWeight clears the value of the "weight" field.
func (pu *ParcelUpdate) ClearWeight() *ParcelUpdate {
	pu.mutation.ClearWeight()
	return pu
}

// SetAmount sets the "amount" field.
func (pu *ParcelUpdate) SetAmount(f float64) *ParcelUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(f)
	return pu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pu *ParcelUpdate) SetNillableAmount(f *float64) *ParcelUpdate {
	if f != nil {
		pu.SetAmount(*f)
	}
	return pu
}

// AddAmount adds f to the "amount" field.
func (pu *ParcelUpdate) AddAmount(f float64) *ParcelUpdate {
	pu.mutation.AddAmount(f)
	return pu
}

// SetPaidAt sets the "paid_at" field.
func (pu *ParcelUpdate) SetPaidAt(t time.Time) *ParcelUpdate {
	pu.mutation.SetPaidAt(t)
	return pu
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (pu *ParcelUpdate) SetNillablePaidAt(t *time.Time) *ParcelUpdate {
	if t != nil {
		pu.SetPaidAt(*t)
	}
	return pu
}

// ClearPaidAt clears the value of the "paid_at" field.
func (pu *ParcelUpdate) ClearPaidAt() *ParcelUpdate {
	pu.mutation.ClearPaidAt()
	return pu
}

// SetTansType sets the "tans_type" field.
func (pu *ParcelUpdate) SetTansType(pt parcel.TansType) *ParcelUpdate {
	pu.mutation.SetTansType(pt)
	return pu
}

// SetNillableTansType sets the "tans_type" field if the given value is not nil.
func (pu *ParcelUpdate) SetNillableTansType(pt *parcel.TansType) *ParcelUpdate {
	if pt != nil {
		pu.SetTansType(*pt)
	}
	return pu
}

// SetStatus sets the "status" field.
func (pu *ParcelUpdate) SetStatus(pa parcel.Status) *ParcelUpdate {
	pu.mutation.SetStatus(pa)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *ParcelUpdate) SetNillableStatus(pa *parcel.Status) *ParcelUpdate {
	if pa != nil {
		pu.SetStatus(*pa)
	}
	return pu
}

// AddImageIDs adds the "images" edge to the ParcelImage entity by IDs.
func (pu *ParcelUpdate) AddImageIDs(ids ...int) *ParcelUpdate {
	pu.mutation.AddImageIDs(ids...)
	return pu
}

// AddImages adds the "images" edges to the ParcelImage entity.
func (pu *ParcelUpdate) AddImages(p ...*ParcelImage) *ParcelUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddImageIDs(ids...)
}

// SetTripID sets the "trip" edge to the Trip entity by ID.
func (pu *ParcelUpdate) SetTripID(id int) *ParcelUpdate {
	pu.mutation.SetTripID(id)
	return pu
}

// SetNillableTripID sets the "trip" edge to the Trip entity by ID if the given value is not nil.
func (pu *ParcelUpdate) SetNillableTripID(id *int) *ParcelUpdate {
	if id != nil {
		pu = pu.SetTripID(*id)
	}
	return pu
}

// SetTrip sets the "trip" edge to the Trip entity.
func (pu *ParcelUpdate) SetTrip(t *Trip) *ParcelUpdate {
	return pu.SetTripID(t.ID)
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (pu *ParcelUpdate) SetCompanyID(id int) *ParcelUpdate {
	pu.mutation.SetCompanyID(id)
	return pu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (pu *ParcelUpdate) SetNillableCompanyID(id *int) *ParcelUpdate {
	if id != nil {
		pu = pu.SetCompanyID(*id)
	}
	return pu
}

// SetCompany sets the "company" edge to the Company entity.
func (pu *ParcelUpdate) SetCompany(c *Company) *ParcelUpdate {
	return pu.SetCompanyID(c.ID)
}

// SetDriverID sets the "driver" edge to the CompanyUser entity by ID.
func (pu *ParcelUpdate) SetDriverID(id int) *ParcelUpdate {
	pu.mutation.SetDriverID(id)
	return pu
}

// SetNillableDriverID sets the "driver" edge to the CompanyUser entity by ID if the given value is not nil.
func (pu *ParcelUpdate) SetNillableDriverID(id *int) *ParcelUpdate {
	if id != nil {
		pu = pu.SetDriverID(*id)
	}
	return pu
}

// SetDriver sets the "driver" edge to the CompanyUser entity.
func (pu *ParcelUpdate) SetDriver(c *CompanyUser) *ParcelUpdate {
	return pu.SetDriverID(c.ID)
}

// Mutation returns the ParcelMutation object of the builder.
func (pu *ParcelUpdate) Mutation() *ParcelMutation {
	return pu.mutation
}

// ClearImages clears all "images" edges to the ParcelImage entity.
func (pu *ParcelUpdate) ClearImages() *ParcelUpdate {
	pu.mutation.ClearImages()
	return pu
}

// RemoveImageIDs removes the "images" edge to ParcelImage entities by IDs.
func (pu *ParcelUpdate) RemoveImageIDs(ids ...int) *ParcelUpdate {
	pu.mutation.RemoveImageIDs(ids...)
	return pu
}

// RemoveImages removes "images" edges to ParcelImage entities.
func (pu *ParcelUpdate) RemoveImages(p ...*ParcelImage) *ParcelUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveImageIDs(ids...)
}

// ClearTrip clears the "trip" edge to the Trip entity.
func (pu *ParcelUpdate) ClearTrip() *ParcelUpdate {
	pu.mutation.ClearTrip()
	return pu
}

// ClearCompany clears the "company" edge to the Company entity.
func (pu *ParcelUpdate) ClearCompany() *ParcelUpdate {
	pu.mutation.ClearCompany()
	return pu
}

// ClearDriver clears the "driver" edge to the CompanyUser entity.
func (pu *ParcelUpdate) ClearDriver() *ParcelUpdate {
	pu.mutation.ClearDriver()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ParcelUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ParcelUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ParcelUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ParcelUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ParcelUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := parcel.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ParcelUpdate) check() error {
	if v, ok := pu.mutation.ParcelCode(); ok {
		if err := parcel.ParcelCodeValidator(v); err != nil {
			return &ValidationError{Name: "parcel_code", err: fmt.Errorf(`ent: validator failed for field "Parcel.parcel_code": %w`, err)}
		}
	}
	if v, ok := pu.mutation.GetType(); ok {
		if err := parcel.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Parcel.type": %w`, err)}
		}
	}
	if v, ok := pu.mutation.SenderName(); ok {
		if err := parcel.SenderNameValidator(v); err != nil {
			return &ValidationError{Name: "sender_name", err: fmt.Errorf(`ent: validator failed for field "Parcel.sender_name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.SenderPhone(); ok {
		if err := parcel.SenderPhoneValidator(v); err != nil {
			return &ValidationError{Name: "sender_phone", err: fmt.Errorf(`ent: validator failed for field "Parcel.sender_phone": %w`, err)}
		}
	}
	if v, ok := pu.mutation.SenderEmail(); ok {
		if err := parcel.SenderEmailValidator(v); err != nil {
			return &ValidationError{Name: "sender_email", err: fmt.Errorf(`ent: validator failed for field "Parcel.sender_email": %w`, err)}
		}
	}
	if v, ok := pu.mutation.RecipientName(); ok {
		if err := parcel.RecipientNameValidator(v); err != nil {
			return &ValidationError{Name: "recipient_name", err: fmt.Errorf(`ent: validator failed for field "Parcel.recipient_name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.RecipientPhone(); ok {
		if err := parcel.RecipientPhoneValidator(v); err != nil {
			return &ValidationError{Name: "recipient_phone", err: fmt.Errorf(`ent: validator failed for field "Parcel.recipient_phone": %w`, err)}
		}
	}
	if v, ok := pu.mutation.RecipientLocation(); ok {
		if err := parcel.RecipientLocationValidator(v); err != nil {
			return &ValidationError{Name: "recipient_location", err: fmt.Errorf(`ent: validator failed for field "Parcel.recipient_location": %w`, err)}
		}
	}
	if v, ok := pu.mutation.TansType(); ok {
		if err := parcel.TansTypeValidator(v); err != nil {
			return &ValidationError{Name: "tans_type", err: fmt.Errorf(`ent: validator failed for field "Parcel.tans_type": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := parcel.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Parcel.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *ParcelUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ParcelUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *ParcelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(parcel.Table, parcel.Columns, sqlgraph.NewFieldSpec(parcel.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(parcel.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.ParcelCode(); ok {
		_spec.SetField(parcel.FieldParcelCode, field.TypeString, value)
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(parcel.FieldType, field.TypeString, value)
	}
	if value, ok := pu.mutation.SenderName(); ok {
		_spec.SetField(parcel.FieldSenderName, field.TypeString, value)
	}
	if value, ok := pu.mutation.SenderPhone(); ok {
		_spec.SetField(parcel.FieldSenderPhone, field.TypeString, value)
	}
	if value, ok := pu.mutation.SenderEmail(); ok {
		_spec.SetField(parcel.FieldSenderEmail, field.TypeString, value)
	}
	if value, ok := pu.mutation.RecipientName(); ok {
		_spec.SetField(parcel.FieldRecipientName, field.TypeString, value)
	}
	if value, ok := pu.mutation.RecipientPhone(); ok {
		_spec.SetField(parcel.FieldRecipientPhone, field.TypeString, value)
	}
	if value, ok := pu.mutation.RecipientLocation(); ok {
		_spec.SetField(parcel.FieldRecipientLocation, field.TypeString, value)
	}
	if value, ok := pu.mutation.Weight(); ok {
		_spec.SetField(parcel.FieldWeight, field.TypeFloat32, value)
	}
	if value, ok := pu.mutation.AddedWeight(); ok {
		_spec.AddField(parcel.FieldWeight, field.TypeFloat32, value)
	}
	if pu.mutation.WeightCleared() {
		_spec.ClearField(parcel.FieldWeight, field.TypeFloat32)
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(parcel.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.AddField(parcel.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.PaidAt(); ok {
		_spec.SetField(parcel.FieldPaidAt, field.TypeTime, value)
	}
	if pu.mutation.PaidAtCleared() {
		_spec.ClearField(parcel.FieldPaidAt, field.TypeTime)
	}
	if value, ok := pu.mutation.TansType(); ok {
		_spec.SetField(parcel.FieldTansType, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(parcel.FieldStatus, field.TypeEnum, value)
	}
	if pu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   parcel.ImagesTable,
			Columns: []string{parcel.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcelimage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedImagesIDs(); len(nodes) > 0 && !pu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   parcel.ImagesTable,
			Columns: []string{parcel.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcelimage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   parcel.ImagesTable,
			Columns: []string{parcel.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcelimage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.TripCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.TripTable,
			Columns: []string{parcel.TripColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TripIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.TripTable,
			Columns: []string{parcel.TripColumn},
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
	if pu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.CompanyTable,
			Columns: []string{parcel.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.CompanyTable,
			Columns: []string{parcel.CompanyColumn},
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
	if pu.mutation.DriverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.DriverTable,
			Columns: []string{parcel.DriverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyuser.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DriverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.DriverTable,
			Columns: []string{parcel.DriverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{parcel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ParcelUpdateOne is the builder for updating a single Parcel entity.
type ParcelUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ParcelMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ParcelUpdateOne) SetUpdatedAt(t time.Time) *ParcelUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetParcelCode sets the "parcel_code" field.
func (puo *ParcelUpdateOne) SetParcelCode(s string) *ParcelUpdateOne {
	puo.mutation.SetParcelCode(s)
	return puo
}

// SetType sets the "type" field.
func (puo *ParcelUpdateOne) SetType(s string) *ParcelUpdateOne {
	puo.mutation.SetType(s)
	return puo
}

// SetSenderName sets the "sender_name" field.
func (puo *ParcelUpdateOne) SetSenderName(s string) *ParcelUpdateOne {
	puo.mutation.SetSenderName(s)
	return puo
}

// SetSenderPhone sets the "sender_phone" field.
func (puo *ParcelUpdateOne) SetSenderPhone(s string) *ParcelUpdateOne {
	puo.mutation.SetSenderPhone(s)
	return puo
}

// SetSenderEmail sets the "sender_email" field.
func (puo *ParcelUpdateOne) SetSenderEmail(s string) *ParcelUpdateOne {
	puo.mutation.SetSenderEmail(s)
	return puo
}

// SetRecipientName sets the "recipient_name" field.
func (puo *ParcelUpdateOne) SetRecipientName(s string) *ParcelUpdateOne {
	puo.mutation.SetRecipientName(s)
	return puo
}

// SetRecipientPhone sets the "recipient_phone" field.
func (puo *ParcelUpdateOne) SetRecipientPhone(s string) *ParcelUpdateOne {
	puo.mutation.SetRecipientPhone(s)
	return puo
}

// SetRecipientLocation sets the "recipient_location" field.
func (puo *ParcelUpdateOne) SetRecipientLocation(s string) *ParcelUpdateOne {
	puo.mutation.SetRecipientLocation(s)
	return puo
}

// SetWeight sets the "weight" field.
func (puo *ParcelUpdateOne) SetWeight(f float32) *ParcelUpdateOne {
	puo.mutation.ResetWeight()
	puo.mutation.SetWeight(f)
	return puo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (puo *ParcelUpdateOne) SetNillableWeight(f *float32) *ParcelUpdateOne {
	if f != nil {
		puo.SetWeight(*f)
	}
	return puo
}

// AddWeight adds f to the "weight" field.
func (puo *ParcelUpdateOne) AddWeight(f float32) *ParcelUpdateOne {
	puo.mutation.AddWeight(f)
	return puo
}

// ClearWeight clears the value of the "weight" field.
func (puo *ParcelUpdateOne) ClearWeight() *ParcelUpdateOne {
	puo.mutation.ClearWeight()
	return puo
}

// SetAmount sets the "amount" field.
func (puo *ParcelUpdateOne) SetAmount(f float64) *ParcelUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(f)
	return puo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (puo *ParcelUpdateOne) SetNillableAmount(f *float64) *ParcelUpdateOne {
	if f != nil {
		puo.SetAmount(*f)
	}
	return puo
}

// AddAmount adds f to the "amount" field.
func (puo *ParcelUpdateOne) AddAmount(f float64) *ParcelUpdateOne {
	puo.mutation.AddAmount(f)
	return puo
}

// SetPaidAt sets the "paid_at" field.
func (puo *ParcelUpdateOne) SetPaidAt(t time.Time) *ParcelUpdateOne {
	puo.mutation.SetPaidAt(t)
	return puo
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (puo *ParcelUpdateOne) SetNillablePaidAt(t *time.Time) *ParcelUpdateOne {
	if t != nil {
		puo.SetPaidAt(*t)
	}
	return puo
}

// ClearPaidAt clears the value of the "paid_at" field.
func (puo *ParcelUpdateOne) ClearPaidAt() *ParcelUpdateOne {
	puo.mutation.ClearPaidAt()
	return puo
}

// SetTansType sets the "tans_type" field.
func (puo *ParcelUpdateOne) SetTansType(pt parcel.TansType) *ParcelUpdateOne {
	puo.mutation.SetTansType(pt)
	return puo
}

// SetNillableTansType sets the "tans_type" field if the given value is not nil.
func (puo *ParcelUpdateOne) SetNillableTansType(pt *parcel.TansType) *ParcelUpdateOne {
	if pt != nil {
		puo.SetTansType(*pt)
	}
	return puo
}

// SetStatus sets the "status" field.
func (puo *ParcelUpdateOne) SetStatus(pa parcel.Status) *ParcelUpdateOne {
	puo.mutation.SetStatus(pa)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *ParcelUpdateOne) SetNillableStatus(pa *parcel.Status) *ParcelUpdateOne {
	if pa != nil {
		puo.SetStatus(*pa)
	}
	return puo
}

// AddImageIDs adds the "images" edge to the ParcelImage entity by IDs.
func (puo *ParcelUpdateOne) AddImageIDs(ids ...int) *ParcelUpdateOne {
	puo.mutation.AddImageIDs(ids...)
	return puo
}

// AddImages adds the "images" edges to the ParcelImage entity.
func (puo *ParcelUpdateOne) AddImages(p ...*ParcelImage) *ParcelUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddImageIDs(ids...)
}

// SetTripID sets the "trip" edge to the Trip entity by ID.
func (puo *ParcelUpdateOne) SetTripID(id int) *ParcelUpdateOne {
	puo.mutation.SetTripID(id)
	return puo
}

// SetNillableTripID sets the "trip" edge to the Trip entity by ID if the given value is not nil.
func (puo *ParcelUpdateOne) SetNillableTripID(id *int) *ParcelUpdateOne {
	if id != nil {
		puo = puo.SetTripID(*id)
	}
	return puo
}

// SetTrip sets the "trip" edge to the Trip entity.
func (puo *ParcelUpdateOne) SetTrip(t *Trip) *ParcelUpdateOne {
	return puo.SetTripID(t.ID)
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (puo *ParcelUpdateOne) SetCompanyID(id int) *ParcelUpdateOne {
	puo.mutation.SetCompanyID(id)
	return puo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (puo *ParcelUpdateOne) SetNillableCompanyID(id *int) *ParcelUpdateOne {
	if id != nil {
		puo = puo.SetCompanyID(*id)
	}
	return puo
}

// SetCompany sets the "company" edge to the Company entity.
func (puo *ParcelUpdateOne) SetCompany(c *Company) *ParcelUpdateOne {
	return puo.SetCompanyID(c.ID)
}

// SetDriverID sets the "driver" edge to the CompanyUser entity by ID.
func (puo *ParcelUpdateOne) SetDriverID(id int) *ParcelUpdateOne {
	puo.mutation.SetDriverID(id)
	return puo
}

// SetNillableDriverID sets the "driver" edge to the CompanyUser entity by ID if the given value is not nil.
func (puo *ParcelUpdateOne) SetNillableDriverID(id *int) *ParcelUpdateOne {
	if id != nil {
		puo = puo.SetDriverID(*id)
	}
	return puo
}

// SetDriver sets the "driver" edge to the CompanyUser entity.
func (puo *ParcelUpdateOne) SetDriver(c *CompanyUser) *ParcelUpdateOne {
	return puo.SetDriverID(c.ID)
}

// Mutation returns the ParcelMutation object of the builder.
func (puo *ParcelUpdateOne) Mutation() *ParcelMutation {
	return puo.mutation
}

// ClearImages clears all "images" edges to the ParcelImage entity.
func (puo *ParcelUpdateOne) ClearImages() *ParcelUpdateOne {
	puo.mutation.ClearImages()
	return puo
}

// RemoveImageIDs removes the "images" edge to ParcelImage entities by IDs.
func (puo *ParcelUpdateOne) RemoveImageIDs(ids ...int) *ParcelUpdateOne {
	puo.mutation.RemoveImageIDs(ids...)
	return puo
}

// RemoveImages removes "images" edges to ParcelImage entities.
func (puo *ParcelUpdateOne) RemoveImages(p ...*ParcelImage) *ParcelUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveImageIDs(ids...)
}

// ClearTrip clears the "trip" edge to the Trip entity.
func (puo *ParcelUpdateOne) ClearTrip() *ParcelUpdateOne {
	puo.mutation.ClearTrip()
	return puo
}

// ClearCompany clears the "company" edge to the Company entity.
func (puo *ParcelUpdateOne) ClearCompany() *ParcelUpdateOne {
	puo.mutation.ClearCompany()
	return puo
}

// ClearDriver clears the "driver" edge to the CompanyUser entity.
func (puo *ParcelUpdateOne) ClearDriver() *ParcelUpdateOne {
	puo.mutation.ClearDriver()
	return puo
}

// Where appends a list predicates to the ParcelUpdate builder.
func (puo *ParcelUpdateOne) Where(ps ...predicate.Parcel) *ParcelUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ParcelUpdateOne) Select(field string, fields ...string) *ParcelUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Parcel entity.
func (puo *ParcelUpdateOne) Save(ctx context.Context) (*Parcel, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ParcelUpdateOne) SaveX(ctx context.Context) *Parcel {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ParcelUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ParcelUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ParcelUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := parcel.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ParcelUpdateOne) check() error {
	if v, ok := puo.mutation.ParcelCode(); ok {
		if err := parcel.ParcelCodeValidator(v); err != nil {
			return &ValidationError{Name: "parcel_code", err: fmt.Errorf(`ent: validator failed for field "Parcel.parcel_code": %w`, err)}
		}
	}
	if v, ok := puo.mutation.GetType(); ok {
		if err := parcel.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Parcel.type": %w`, err)}
		}
	}
	if v, ok := puo.mutation.SenderName(); ok {
		if err := parcel.SenderNameValidator(v); err != nil {
			return &ValidationError{Name: "sender_name", err: fmt.Errorf(`ent: validator failed for field "Parcel.sender_name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.SenderPhone(); ok {
		if err := parcel.SenderPhoneValidator(v); err != nil {
			return &ValidationError{Name: "sender_phone", err: fmt.Errorf(`ent: validator failed for field "Parcel.sender_phone": %w`, err)}
		}
	}
	if v, ok := puo.mutation.SenderEmail(); ok {
		if err := parcel.SenderEmailValidator(v); err != nil {
			return &ValidationError{Name: "sender_email", err: fmt.Errorf(`ent: validator failed for field "Parcel.sender_email": %w`, err)}
		}
	}
	if v, ok := puo.mutation.RecipientName(); ok {
		if err := parcel.RecipientNameValidator(v); err != nil {
			return &ValidationError{Name: "recipient_name", err: fmt.Errorf(`ent: validator failed for field "Parcel.recipient_name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.RecipientPhone(); ok {
		if err := parcel.RecipientPhoneValidator(v); err != nil {
			return &ValidationError{Name: "recipient_phone", err: fmt.Errorf(`ent: validator failed for field "Parcel.recipient_phone": %w`, err)}
		}
	}
	if v, ok := puo.mutation.RecipientLocation(); ok {
		if err := parcel.RecipientLocationValidator(v); err != nil {
			return &ValidationError{Name: "recipient_location", err: fmt.Errorf(`ent: validator failed for field "Parcel.recipient_location": %w`, err)}
		}
	}
	if v, ok := puo.mutation.TansType(); ok {
		if err := parcel.TansTypeValidator(v); err != nil {
			return &ValidationError{Name: "tans_type", err: fmt.Errorf(`ent: validator failed for field "Parcel.tans_type": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := parcel.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Parcel.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *ParcelUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ParcelUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *ParcelUpdateOne) sqlSave(ctx context.Context) (_node *Parcel, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(parcel.Table, parcel.Columns, sqlgraph.NewFieldSpec(parcel.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Parcel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, parcel.FieldID)
		for _, f := range fields {
			if !parcel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != parcel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(parcel.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.ParcelCode(); ok {
		_spec.SetField(parcel.FieldParcelCode, field.TypeString, value)
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(parcel.FieldType, field.TypeString, value)
	}
	if value, ok := puo.mutation.SenderName(); ok {
		_spec.SetField(parcel.FieldSenderName, field.TypeString, value)
	}
	if value, ok := puo.mutation.SenderPhone(); ok {
		_spec.SetField(parcel.FieldSenderPhone, field.TypeString, value)
	}
	if value, ok := puo.mutation.SenderEmail(); ok {
		_spec.SetField(parcel.FieldSenderEmail, field.TypeString, value)
	}
	if value, ok := puo.mutation.RecipientName(); ok {
		_spec.SetField(parcel.FieldRecipientName, field.TypeString, value)
	}
	if value, ok := puo.mutation.RecipientPhone(); ok {
		_spec.SetField(parcel.FieldRecipientPhone, field.TypeString, value)
	}
	if value, ok := puo.mutation.RecipientLocation(); ok {
		_spec.SetField(parcel.FieldRecipientLocation, field.TypeString, value)
	}
	if value, ok := puo.mutation.Weight(); ok {
		_spec.SetField(parcel.FieldWeight, field.TypeFloat32, value)
	}
	if value, ok := puo.mutation.AddedWeight(); ok {
		_spec.AddField(parcel.FieldWeight, field.TypeFloat32, value)
	}
	if puo.mutation.WeightCleared() {
		_spec.ClearField(parcel.FieldWeight, field.TypeFloat32)
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(parcel.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.AddField(parcel.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.PaidAt(); ok {
		_spec.SetField(parcel.FieldPaidAt, field.TypeTime, value)
	}
	if puo.mutation.PaidAtCleared() {
		_spec.ClearField(parcel.FieldPaidAt, field.TypeTime)
	}
	if value, ok := puo.mutation.TansType(); ok {
		_spec.SetField(parcel.FieldTansType, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(parcel.FieldStatus, field.TypeEnum, value)
	}
	if puo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   parcel.ImagesTable,
			Columns: []string{parcel.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcelimage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedImagesIDs(); len(nodes) > 0 && !puo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   parcel.ImagesTable,
			Columns: []string{parcel.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcelimage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   parcel.ImagesTable,
			Columns: []string{parcel.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(parcelimage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.TripCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.TripTable,
			Columns: []string{parcel.TripColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TripIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.TripTable,
			Columns: []string{parcel.TripColumn},
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
	if puo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.CompanyTable,
			Columns: []string{parcel.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.CompanyTable,
			Columns: []string{parcel.CompanyColumn},
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
	if puo.mutation.DriverCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.DriverTable,
			Columns: []string{parcel.DriverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyuser.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DriverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   parcel.DriverTable,
			Columns: []string{parcel.DriverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Parcel{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{parcel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
