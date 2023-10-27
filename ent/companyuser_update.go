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
	"github.com/SeyramWood/ent/companyuser"
	"github.com/SeyramWood/ent/notification"
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/trip"
	"github.com/SeyramWood/ent/user"
)

// CompanyUserUpdate is the builder for updating CompanyUser entities.
type CompanyUserUpdate struct {
	config
	hooks     []Hook
	mutation  *CompanyUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CompanyUserUpdate builder.
func (cuu *CompanyUserUpdate) Where(ps ...predicate.CompanyUser) *CompanyUserUpdate {
	cuu.mutation.Where(ps...)
	return cuu
}

// SetUpdatedAt sets the "updated_at" field.
func (cuu *CompanyUserUpdate) SetUpdatedAt(t time.Time) *CompanyUserUpdate {
	cuu.mutation.SetUpdatedAt(t)
	return cuu
}

// SetLastName sets the "last_name" field.
func (cuu *CompanyUserUpdate) SetLastName(s string) *CompanyUserUpdate {
	cuu.mutation.SetLastName(s)
	return cuu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (cuu *CompanyUserUpdate) SetNillableLastName(s *string) *CompanyUserUpdate {
	if s != nil {
		cuu.SetLastName(*s)
	}
	return cuu
}

// ClearLastName clears the value of the "last_name" field.
func (cuu *CompanyUserUpdate) ClearLastName() *CompanyUserUpdate {
	cuu.mutation.ClearLastName()
	return cuu
}

// SetOtherName sets the "other_name" field.
func (cuu *CompanyUserUpdate) SetOtherName(s string) *CompanyUserUpdate {
	cuu.mutation.SetOtherName(s)
	return cuu
}

// SetNillableOtherName sets the "other_name" field if the given value is not nil.
func (cuu *CompanyUserUpdate) SetNillableOtherName(s *string) *CompanyUserUpdate {
	if s != nil {
		cuu.SetOtherName(*s)
	}
	return cuu
}

// ClearOtherName clears the value of the "other_name" field.
func (cuu *CompanyUserUpdate) ClearOtherName() *CompanyUserUpdate {
	cuu.mutation.ClearOtherName()
	return cuu
}

// SetPhone sets the "phone" field.
func (cuu *CompanyUserUpdate) SetPhone(s string) *CompanyUserUpdate {
	cuu.mutation.SetPhone(s)
	return cuu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (cuu *CompanyUserUpdate) SetNillablePhone(s *string) *CompanyUserUpdate {
	if s != nil {
		cuu.SetPhone(*s)
	}
	return cuu
}

// ClearPhone clears the value of the "phone" field.
func (cuu *CompanyUserUpdate) ClearPhone() *CompanyUserUpdate {
	cuu.mutation.ClearPhone()
	return cuu
}

// SetOtherPhone sets the "other_phone" field.
func (cuu *CompanyUserUpdate) SetOtherPhone(s string) *CompanyUserUpdate {
	cuu.mutation.SetOtherPhone(s)
	return cuu
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (cuu *CompanyUserUpdate) SetNillableOtherPhone(s *string) *CompanyUserUpdate {
	if s != nil {
		cuu.SetOtherPhone(*s)
	}
	return cuu
}

// ClearOtherPhone clears the value of the "other_phone" field.
func (cuu *CompanyUserUpdate) ClearOtherPhone() *CompanyUserUpdate {
	cuu.mutation.ClearOtherPhone()
	return cuu
}

// SetRole sets the "role" field.
func (cuu *CompanyUserUpdate) SetRole(c companyuser.Role) *CompanyUserUpdate {
	cuu.mutation.SetRole(c)
	return cuu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (cuu *CompanyUserUpdate) SetNillableRole(c *companyuser.Role) *CompanyUserUpdate {
	if c != nil {
		cuu.SetRole(*c)
	}
	return cuu
}

// SetProfileID sets the "profile" edge to the User entity by ID.
func (cuu *CompanyUserUpdate) SetProfileID(id int) *CompanyUserUpdate {
	cuu.mutation.SetProfileID(id)
	return cuu
}

// SetNillableProfileID sets the "profile" edge to the User entity by ID if the given value is not nil.
func (cuu *CompanyUserUpdate) SetNillableProfileID(id *int) *CompanyUserUpdate {
	if id != nil {
		cuu = cuu.SetProfileID(*id)
	}
	return cuu
}

// SetProfile sets the "profile" edge to the User entity.
func (cuu *CompanyUserUpdate) SetProfile(u *User) *CompanyUserUpdate {
	return cuu.SetProfileID(u.ID)
}

// AddTripIDs adds the "trips" edge to the Trip entity by IDs.
func (cuu *CompanyUserUpdate) AddTripIDs(ids ...int) *CompanyUserUpdate {
	cuu.mutation.AddTripIDs(ids...)
	return cuu
}

// AddTrips adds the "trips" edges to the Trip entity.
func (cuu *CompanyUserUpdate) AddTrips(t ...*Trip) *CompanyUserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuu.AddTripIDs(ids...)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (cuu *CompanyUserUpdate) AddNotificationIDs(ids ...int) *CompanyUserUpdate {
	cuu.mutation.AddNotificationIDs(ids...)
	return cuu
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (cuu *CompanyUserUpdate) AddNotifications(n ...*Notification) *CompanyUserUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cuu.AddNotificationIDs(ids...)
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (cuu *CompanyUserUpdate) SetCompanyID(id int) *CompanyUserUpdate {
	cuu.mutation.SetCompanyID(id)
	return cuu
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (cuu *CompanyUserUpdate) SetNillableCompanyID(id *int) *CompanyUserUpdate {
	if id != nil {
		cuu = cuu.SetCompanyID(*id)
	}
	return cuu
}

// SetCompany sets the "company" edge to the Company entity.
func (cuu *CompanyUserUpdate) SetCompany(c *Company) *CompanyUserUpdate {
	return cuu.SetCompanyID(c.ID)
}

// Mutation returns the CompanyUserMutation object of the builder.
func (cuu *CompanyUserUpdate) Mutation() *CompanyUserMutation {
	return cuu.mutation
}

// ClearProfile clears the "profile" edge to the User entity.
func (cuu *CompanyUserUpdate) ClearProfile() *CompanyUserUpdate {
	cuu.mutation.ClearProfile()
	return cuu
}

// ClearTrips clears all "trips" edges to the Trip entity.
func (cuu *CompanyUserUpdate) ClearTrips() *CompanyUserUpdate {
	cuu.mutation.ClearTrips()
	return cuu
}

// RemoveTripIDs removes the "trips" edge to Trip entities by IDs.
func (cuu *CompanyUserUpdate) RemoveTripIDs(ids ...int) *CompanyUserUpdate {
	cuu.mutation.RemoveTripIDs(ids...)
	return cuu
}

// RemoveTrips removes "trips" edges to Trip entities.
func (cuu *CompanyUserUpdate) RemoveTrips(t ...*Trip) *CompanyUserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuu.RemoveTripIDs(ids...)
}

// ClearNotifications clears all "notifications" edges to the Notification entity.
func (cuu *CompanyUserUpdate) ClearNotifications() *CompanyUserUpdate {
	cuu.mutation.ClearNotifications()
	return cuu
}

// RemoveNotificationIDs removes the "notifications" edge to Notification entities by IDs.
func (cuu *CompanyUserUpdate) RemoveNotificationIDs(ids ...int) *CompanyUserUpdate {
	cuu.mutation.RemoveNotificationIDs(ids...)
	return cuu
}

// RemoveNotifications removes "notifications" edges to Notification entities.
func (cuu *CompanyUserUpdate) RemoveNotifications(n ...*Notification) *CompanyUserUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cuu.RemoveNotificationIDs(ids...)
}

// ClearCompany clears the "company" edge to the Company entity.
func (cuu *CompanyUserUpdate) ClearCompany() *CompanyUserUpdate {
	cuu.mutation.ClearCompany()
	return cuu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cuu *CompanyUserUpdate) Save(ctx context.Context) (int, error) {
	cuu.defaults()
	return withHooks(ctx, cuu.sqlSave, cuu.mutation, cuu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuu *CompanyUserUpdate) SaveX(ctx context.Context) int {
	affected, err := cuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cuu *CompanyUserUpdate) Exec(ctx context.Context) error {
	_, err := cuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuu *CompanyUserUpdate) ExecX(ctx context.Context) {
	if err := cuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuu *CompanyUserUpdate) defaults() {
	if _, ok := cuu.mutation.UpdatedAt(); !ok {
		v := companyuser.UpdateDefaultUpdatedAt()
		cuu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuu *CompanyUserUpdate) check() error {
	if v, ok := cuu.mutation.Role(); ok {
		if err := companyuser.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "CompanyUser.role": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuu *CompanyUserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CompanyUserUpdate {
	cuu.modifiers = append(cuu.modifiers, modifiers...)
	return cuu
}

func (cuu *CompanyUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cuu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(companyuser.Table, companyuser.Columns, sqlgraph.NewFieldSpec(companyuser.FieldID, field.TypeInt))
	if ps := cuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuu.mutation.UpdatedAt(); ok {
		_spec.SetField(companyuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuu.mutation.LastName(); ok {
		_spec.SetField(companyuser.FieldLastName, field.TypeString, value)
	}
	if cuu.mutation.LastNameCleared() {
		_spec.ClearField(companyuser.FieldLastName, field.TypeString)
	}
	if value, ok := cuu.mutation.OtherName(); ok {
		_spec.SetField(companyuser.FieldOtherName, field.TypeString, value)
	}
	if cuu.mutation.OtherNameCleared() {
		_spec.ClearField(companyuser.FieldOtherName, field.TypeString)
	}
	if value, ok := cuu.mutation.Phone(); ok {
		_spec.SetField(companyuser.FieldPhone, field.TypeString, value)
	}
	if cuu.mutation.PhoneCleared() {
		_spec.ClearField(companyuser.FieldPhone, field.TypeString)
	}
	if value, ok := cuu.mutation.OtherPhone(); ok {
		_spec.SetField(companyuser.FieldOtherPhone, field.TypeString, value)
	}
	if cuu.mutation.OtherPhoneCleared() {
		_spec.ClearField(companyuser.FieldOtherPhone, field.TypeString)
	}
	if value, ok := cuu.mutation.Role(); ok {
		_spec.SetField(companyuser.FieldRole, field.TypeEnum, value)
	}
	if cuu.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   companyuser.ProfileTable,
			Columns: []string{companyuser.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuu.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   companyuser.ProfileTable,
			Columns: []string{companyuser.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuu.mutation.TripsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyuser.TripsTable,
			Columns: []string{companyuser.TripsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuu.mutation.RemovedTripsIDs(); len(nodes) > 0 && !cuu.mutation.TripsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyuser.TripsTable,
			Columns: []string{companyuser.TripsColumn},
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
	if nodes := cuu.mutation.TripsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyuser.TripsTable,
			Columns: []string{companyuser.TripsColumn},
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
	if cuu.mutation.NotificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   companyuser.NotificationsTable,
			Columns: companyuser.NotificationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuu.mutation.RemovedNotificationsIDs(); len(nodes) > 0 && !cuu.mutation.NotificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   companyuser.NotificationsTable,
			Columns: companyuser.NotificationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuu.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   companyuser.NotificationsTable,
			Columns: companyuser.NotificationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuu.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companyuser.CompanyTable,
			Columns: []string{companyuser.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuu.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companyuser.CompanyTable,
			Columns: []string{companyuser.CompanyColumn},
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
	_spec.AddModifiers(cuu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{companyuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cuu.mutation.done = true
	return n, nil
}

// CompanyUserUpdateOne is the builder for updating a single CompanyUser entity.
type CompanyUserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CompanyUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (cuuo *CompanyUserUpdateOne) SetUpdatedAt(t time.Time) *CompanyUserUpdateOne {
	cuuo.mutation.SetUpdatedAt(t)
	return cuuo
}

// SetLastName sets the "last_name" field.
func (cuuo *CompanyUserUpdateOne) SetLastName(s string) *CompanyUserUpdateOne {
	cuuo.mutation.SetLastName(s)
	return cuuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (cuuo *CompanyUserUpdateOne) SetNillableLastName(s *string) *CompanyUserUpdateOne {
	if s != nil {
		cuuo.SetLastName(*s)
	}
	return cuuo
}

// ClearLastName clears the value of the "last_name" field.
func (cuuo *CompanyUserUpdateOne) ClearLastName() *CompanyUserUpdateOne {
	cuuo.mutation.ClearLastName()
	return cuuo
}

// SetOtherName sets the "other_name" field.
func (cuuo *CompanyUserUpdateOne) SetOtherName(s string) *CompanyUserUpdateOne {
	cuuo.mutation.SetOtherName(s)
	return cuuo
}

// SetNillableOtherName sets the "other_name" field if the given value is not nil.
func (cuuo *CompanyUserUpdateOne) SetNillableOtherName(s *string) *CompanyUserUpdateOne {
	if s != nil {
		cuuo.SetOtherName(*s)
	}
	return cuuo
}

// ClearOtherName clears the value of the "other_name" field.
func (cuuo *CompanyUserUpdateOne) ClearOtherName() *CompanyUserUpdateOne {
	cuuo.mutation.ClearOtherName()
	return cuuo
}

// SetPhone sets the "phone" field.
func (cuuo *CompanyUserUpdateOne) SetPhone(s string) *CompanyUserUpdateOne {
	cuuo.mutation.SetPhone(s)
	return cuuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (cuuo *CompanyUserUpdateOne) SetNillablePhone(s *string) *CompanyUserUpdateOne {
	if s != nil {
		cuuo.SetPhone(*s)
	}
	return cuuo
}

// ClearPhone clears the value of the "phone" field.
func (cuuo *CompanyUserUpdateOne) ClearPhone() *CompanyUserUpdateOne {
	cuuo.mutation.ClearPhone()
	return cuuo
}

// SetOtherPhone sets the "other_phone" field.
func (cuuo *CompanyUserUpdateOne) SetOtherPhone(s string) *CompanyUserUpdateOne {
	cuuo.mutation.SetOtherPhone(s)
	return cuuo
}

// SetNillableOtherPhone sets the "other_phone" field if the given value is not nil.
func (cuuo *CompanyUserUpdateOne) SetNillableOtherPhone(s *string) *CompanyUserUpdateOne {
	if s != nil {
		cuuo.SetOtherPhone(*s)
	}
	return cuuo
}

// ClearOtherPhone clears the value of the "other_phone" field.
func (cuuo *CompanyUserUpdateOne) ClearOtherPhone() *CompanyUserUpdateOne {
	cuuo.mutation.ClearOtherPhone()
	return cuuo
}

// SetRole sets the "role" field.
func (cuuo *CompanyUserUpdateOne) SetRole(c companyuser.Role) *CompanyUserUpdateOne {
	cuuo.mutation.SetRole(c)
	return cuuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (cuuo *CompanyUserUpdateOne) SetNillableRole(c *companyuser.Role) *CompanyUserUpdateOne {
	if c != nil {
		cuuo.SetRole(*c)
	}
	return cuuo
}

// SetProfileID sets the "profile" edge to the User entity by ID.
func (cuuo *CompanyUserUpdateOne) SetProfileID(id int) *CompanyUserUpdateOne {
	cuuo.mutation.SetProfileID(id)
	return cuuo
}

// SetNillableProfileID sets the "profile" edge to the User entity by ID if the given value is not nil.
func (cuuo *CompanyUserUpdateOne) SetNillableProfileID(id *int) *CompanyUserUpdateOne {
	if id != nil {
		cuuo = cuuo.SetProfileID(*id)
	}
	return cuuo
}

// SetProfile sets the "profile" edge to the User entity.
func (cuuo *CompanyUserUpdateOne) SetProfile(u *User) *CompanyUserUpdateOne {
	return cuuo.SetProfileID(u.ID)
}

// AddTripIDs adds the "trips" edge to the Trip entity by IDs.
func (cuuo *CompanyUserUpdateOne) AddTripIDs(ids ...int) *CompanyUserUpdateOne {
	cuuo.mutation.AddTripIDs(ids...)
	return cuuo
}

// AddTrips adds the "trips" edges to the Trip entity.
func (cuuo *CompanyUserUpdateOne) AddTrips(t ...*Trip) *CompanyUserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuuo.AddTripIDs(ids...)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (cuuo *CompanyUserUpdateOne) AddNotificationIDs(ids ...int) *CompanyUserUpdateOne {
	cuuo.mutation.AddNotificationIDs(ids...)
	return cuuo
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (cuuo *CompanyUserUpdateOne) AddNotifications(n ...*Notification) *CompanyUserUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cuuo.AddNotificationIDs(ids...)
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (cuuo *CompanyUserUpdateOne) SetCompanyID(id int) *CompanyUserUpdateOne {
	cuuo.mutation.SetCompanyID(id)
	return cuuo
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (cuuo *CompanyUserUpdateOne) SetNillableCompanyID(id *int) *CompanyUserUpdateOne {
	if id != nil {
		cuuo = cuuo.SetCompanyID(*id)
	}
	return cuuo
}

// SetCompany sets the "company" edge to the Company entity.
func (cuuo *CompanyUserUpdateOne) SetCompany(c *Company) *CompanyUserUpdateOne {
	return cuuo.SetCompanyID(c.ID)
}

// Mutation returns the CompanyUserMutation object of the builder.
func (cuuo *CompanyUserUpdateOne) Mutation() *CompanyUserMutation {
	return cuuo.mutation
}

// ClearProfile clears the "profile" edge to the User entity.
func (cuuo *CompanyUserUpdateOne) ClearProfile() *CompanyUserUpdateOne {
	cuuo.mutation.ClearProfile()
	return cuuo
}

// ClearTrips clears all "trips" edges to the Trip entity.
func (cuuo *CompanyUserUpdateOne) ClearTrips() *CompanyUserUpdateOne {
	cuuo.mutation.ClearTrips()
	return cuuo
}

// RemoveTripIDs removes the "trips" edge to Trip entities by IDs.
func (cuuo *CompanyUserUpdateOne) RemoveTripIDs(ids ...int) *CompanyUserUpdateOne {
	cuuo.mutation.RemoveTripIDs(ids...)
	return cuuo
}

// RemoveTrips removes "trips" edges to Trip entities.
func (cuuo *CompanyUserUpdateOne) RemoveTrips(t ...*Trip) *CompanyUserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuuo.RemoveTripIDs(ids...)
}

// ClearNotifications clears all "notifications" edges to the Notification entity.
func (cuuo *CompanyUserUpdateOne) ClearNotifications() *CompanyUserUpdateOne {
	cuuo.mutation.ClearNotifications()
	return cuuo
}

// RemoveNotificationIDs removes the "notifications" edge to Notification entities by IDs.
func (cuuo *CompanyUserUpdateOne) RemoveNotificationIDs(ids ...int) *CompanyUserUpdateOne {
	cuuo.mutation.RemoveNotificationIDs(ids...)
	return cuuo
}

// RemoveNotifications removes "notifications" edges to Notification entities.
func (cuuo *CompanyUserUpdateOne) RemoveNotifications(n ...*Notification) *CompanyUserUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cuuo.RemoveNotificationIDs(ids...)
}

// ClearCompany clears the "company" edge to the Company entity.
func (cuuo *CompanyUserUpdateOne) ClearCompany() *CompanyUserUpdateOne {
	cuuo.mutation.ClearCompany()
	return cuuo
}

// Where appends a list predicates to the CompanyUserUpdate builder.
func (cuuo *CompanyUserUpdateOne) Where(ps ...predicate.CompanyUser) *CompanyUserUpdateOne {
	cuuo.mutation.Where(ps...)
	return cuuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuuo *CompanyUserUpdateOne) Select(field string, fields ...string) *CompanyUserUpdateOne {
	cuuo.fields = append([]string{field}, fields...)
	return cuuo
}

// Save executes the query and returns the updated CompanyUser entity.
func (cuuo *CompanyUserUpdateOne) Save(ctx context.Context) (*CompanyUser, error) {
	cuuo.defaults()
	return withHooks(ctx, cuuo.sqlSave, cuuo.mutation, cuuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuuo *CompanyUserUpdateOne) SaveX(ctx context.Context) *CompanyUser {
	node, err := cuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuuo *CompanyUserUpdateOne) Exec(ctx context.Context) error {
	_, err := cuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuuo *CompanyUserUpdateOne) ExecX(ctx context.Context) {
	if err := cuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuuo *CompanyUserUpdateOne) defaults() {
	if _, ok := cuuo.mutation.UpdatedAt(); !ok {
		v := companyuser.UpdateDefaultUpdatedAt()
		cuuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuuo *CompanyUserUpdateOne) check() error {
	if v, ok := cuuo.mutation.Role(); ok {
		if err := companyuser.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "CompanyUser.role": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuuo *CompanyUserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CompanyUserUpdateOne {
	cuuo.modifiers = append(cuuo.modifiers, modifiers...)
	return cuuo
}

func (cuuo *CompanyUserUpdateOne) sqlSave(ctx context.Context) (_node *CompanyUser, err error) {
	if err := cuuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(companyuser.Table, companyuser.Columns, sqlgraph.NewFieldSpec(companyuser.FieldID, field.TypeInt))
	id, ok := cuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CompanyUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, companyuser.FieldID)
		for _, f := range fields {
			if !companyuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != companyuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuuo.mutation.UpdatedAt(); ok {
		_spec.SetField(companyuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuuo.mutation.LastName(); ok {
		_spec.SetField(companyuser.FieldLastName, field.TypeString, value)
	}
	if cuuo.mutation.LastNameCleared() {
		_spec.ClearField(companyuser.FieldLastName, field.TypeString)
	}
	if value, ok := cuuo.mutation.OtherName(); ok {
		_spec.SetField(companyuser.FieldOtherName, field.TypeString, value)
	}
	if cuuo.mutation.OtherNameCleared() {
		_spec.ClearField(companyuser.FieldOtherName, field.TypeString)
	}
	if value, ok := cuuo.mutation.Phone(); ok {
		_spec.SetField(companyuser.FieldPhone, field.TypeString, value)
	}
	if cuuo.mutation.PhoneCleared() {
		_spec.ClearField(companyuser.FieldPhone, field.TypeString)
	}
	if value, ok := cuuo.mutation.OtherPhone(); ok {
		_spec.SetField(companyuser.FieldOtherPhone, field.TypeString, value)
	}
	if cuuo.mutation.OtherPhoneCleared() {
		_spec.ClearField(companyuser.FieldOtherPhone, field.TypeString)
	}
	if value, ok := cuuo.mutation.Role(); ok {
		_spec.SetField(companyuser.FieldRole, field.TypeEnum, value)
	}
	if cuuo.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   companyuser.ProfileTable,
			Columns: []string{companyuser.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuuo.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   companyuser.ProfileTable,
			Columns: []string{companyuser.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuuo.mutation.TripsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyuser.TripsTable,
			Columns: []string{companyuser.TripsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuuo.mutation.RemovedTripsIDs(); len(nodes) > 0 && !cuuo.mutation.TripsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyuser.TripsTable,
			Columns: []string{companyuser.TripsColumn},
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
	if nodes := cuuo.mutation.TripsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companyuser.TripsTable,
			Columns: []string{companyuser.TripsColumn},
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
	if cuuo.mutation.NotificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   companyuser.NotificationsTable,
			Columns: companyuser.NotificationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuuo.mutation.RemovedNotificationsIDs(); len(nodes) > 0 && !cuuo.mutation.NotificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   companyuser.NotificationsTable,
			Columns: companyuser.NotificationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuuo.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   companyuser.NotificationsTable,
			Columns: companyuser.NotificationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuuo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companyuser.CompanyTable,
			Columns: []string{companyuser.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuuo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companyuser.CompanyTable,
			Columns: []string{companyuser.CompanyColumn},
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
	_spec.AddModifiers(cuuo.modifiers...)
	_node = &CompanyUser{config: cuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{companyuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuuo.mutation.done = true
	return _node, nil
}