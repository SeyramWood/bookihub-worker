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
	"github.com/SeyramWood/bookibus/ent/incident"
	"github.com/SeyramWood/bookibus/ent/incidentimage"
	"github.com/SeyramWood/bookibus/ent/predicate"
)

// IncidentImageUpdate is the builder for updating IncidentImage entities.
type IncidentImageUpdate struct {
	config
	hooks     []Hook
	mutation  *IncidentImageMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the IncidentImageUpdate builder.
func (iiu *IncidentImageUpdate) Where(ps ...predicate.IncidentImage) *IncidentImageUpdate {
	iiu.mutation.Where(ps...)
	return iiu
}

// SetUpdatedAt sets the "updated_at" field.
func (iiu *IncidentImageUpdate) SetUpdatedAt(t time.Time) *IncidentImageUpdate {
	iiu.mutation.SetUpdatedAt(t)
	return iiu
}

// SetImage sets the "image" field.
func (iiu *IncidentImageUpdate) SetImage(s string) *IncidentImageUpdate {
	iiu.mutation.SetImage(s)
	return iiu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (iiu *IncidentImageUpdate) SetNillableImage(s *string) *IncidentImageUpdate {
	if s != nil {
		iiu.SetImage(*s)
	}
	return iiu
}

// ClearImage clears the value of the "image" field.
func (iiu *IncidentImageUpdate) ClearImage() *IncidentImageUpdate {
	iiu.mutation.ClearImage()
	return iiu
}

// SetIncidentID sets the "incident" edge to the Incident entity by ID.
func (iiu *IncidentImageUpdate) SetIncidentID(id int) *IncidentImageUpdate {
	iiu.mutation.SetIncidentID(id)
	return iiu
}

// SetNillableIncidentID sets the "incident" edge to the Incident entity by ID if the given value is not nil.
func (iiu *IncidentImageUpdate) SetNillableIncidentID(id *int) *IncidentImageUpdate {
	if id != nil {
		iiu = iiu.SetIncidentID(*id)
	}
	return iiu
}

// SetIncident sets the "incident" edge to the Incident entity.
func (iiu *IncidentImageUpdate) SetIncident(i *Incident) *IncidentImageUpdate {
	return iiu.SetIncidentID(i.ID)
}

// Mutation returns the IncidentImageMutation object of the builder.
func (iiu *IncidentImageUpdate) Mutation() *IncidentImageMutation {
	return iiu.mutation
}

// ClearIncident clears the "incident" edge to the Incident entity.
func (iiu *IncidentImageUpdate) ClearIncident() *IncidentImageUpdate {
	iiu.mutation.ClearIncident()
	return iiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iiu *IncidentImageUpdate) Save(ctx context.Context) (int, error) {
	iiu.defaults()
	return withHooks(ctx, iiu.sqlSave, iiu.mutation, iiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iiu *IncidentImageUpdate) SaveX(ctx context.Context) int {
	affected, err := iiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iiu *IncidentImageUpdate) Exec(ctx context.Context) error {
	_, err := iiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iiu *IncidentImageUpdate) ExecX(ctx context.Context) {
	if err := iiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iiu *IncidentImageUpdate) defaults() {
	if _, ok := iiu.mutation.UpdatedAt(); !ok {
		v := incidentimage.UpdateDefaultUpdatedAt()
		iiu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iiu *IncidentImageUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *IncidentImageUpdate {
	iiu.modifiers = append(iiu.modifiers, modifiers...)
	return iiu
}

func (iiu *IncidentImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(incidentimage.Table, incidentimage.Columns, sqlgraph.NewFieldSpec(incidentimage.FieldID, field.TypeInt))
	if ps := iiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iiu.mutation.UpdatedAt(); ok {
		_spec.SetField(incidentimage.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iiu.mutation.Image(); ok {
		_spec.SetField(incidentimage.FieldImage, field.TypeString, value)
	}
	if iiu.mutation.ImageCleared() {
		_spec.ClearField(incidentimage.FieldImage, field.TypeString)
	}
	if iiu.mutation.IncidentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incidentimage.IncidentTable,
			Columns: []string{incidentimage.IncidentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(incident.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iiu.mutation.IncidentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incidentimage.IncidentTable,
			Columns: []string{incidentimage.IncidentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(incident.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(iiu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, iiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{incidentimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iiu.mutation.done = true
	return n, nil
}

// IncidentImageUpdateOne is the builder for updating a single IncidentImage entity.
type IncidentImageUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *IncidentImageMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (iiuo *IncidentImageUpdateOne) SetUpdatedAt(t time.Time) *IncidentImageUpdateOne {
	iiuo.mutation.SetUpdatedAt(t)
	return iiuo
}

// SetImage sets the "image" field.
func (iiuo *IncidentImageUpdateOne) SetImage(s string) *IncidentImageUpdateOne {
	iiuo.mutation.SetImage(s)
	return iiuo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (iiuo *IncidentImageUpdateOne) SetNillableImage(s *string) *IncidentImageUpdateOne {
	if s != nil {
		iiuo.SetImage(*s)
	}
	return iiuo
}

// ClearImage clears the value of the "image" field.
func (iiuo *IncidentImageUpdateOne) ClearImage() *IncidentImageUpdateOne {
	iiuo.mutation.ClearImage()
	return iiuo
}

// SetIncidentID sets the "incident" edge to the Incident entity by ID.
func (iiuo *IncidentImageUpdateOne) SetIncidentID(id int) *IncidentImageUpdateOne {
	iiuo.mutation.SetIncidentID(id)
	return iiuo
}

// SetNillableIncidentID sets the "incident" edge to the Incident entity by ID if the given value is not nil.
func (iiuo *IncidentImageUpdateOne) SetNillableIncidentID(id *int) *IncidentImageUpdateOne {
	if id != nil {
		iiuo = iiuo.SetIncidentID(*id)
	}
	return iiuo
}

// SetIncident sets the "incident" edge to the Incident entity.
func (iiuo *IncidentImageUpdateOne) SetIncident(i *Incident) *IncidentImageUpdateOne {
	return iiuo.SetIncidentID(i.ID)
}

// Mutation returns the IncidentImageMutation object of the builder.
func (iiuo *IncidentImageUpdateOne) Mutation() *IncidentImageMutation {
	return iiuo.mutation
}

// ClearIncident clears the "incident" edge to the Incident entity.
func (iiuo *IncidentImageUpdateOne) ClearIncident() *IncidentImageUpdateOne {
	iiuo.mutation.ClearIncident()
	return iiuo
}

// Where appends a list predicates to the IncidentImageUpdate builder.
func (iiuo *IncidentImageUpdateOne) Where(ps ...predicate.IncidentImage) *IncidentImageUpdateOne {
	iiuo.mutation.Where(ps...)
	return iiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iiuo *IncidentImageUpdateOne) Select(field string, fields ...string) *IncidentImageUpdateOne {
	iiuo.fields = append([]string{field}, fields...)
	return iiuo
}

// Save executes the query and returns the updated IncidentImage entity.
func (iiuo *IncidentImageUpdateOne) Save(ctx context.Context) (*IncidentImage, error) {
	iiuo.defaults()
	return withHooks(ctx, iiuo.sqlSave, iiuo.mutation, iiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iiuo *IncidentImageUpdateOne) SaveX(ctx context.Context) *IncidentImage {
	node, err := iiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iiuo *IncidentImageUpdateOne) Exec(ctx context.Context) error {
	_, err := iiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iiuo *IncidentImageUpdateOne) ExecX(ctx context.Context) {
	if err := iiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iiuo *IncidentImageUpdateOne) defaults() {
	if _, ok := iiuo.mutation.UpdatedAt(); !ok {
		v := incidentimage.UpdateDefaultUpdatedAt()
		iiuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iiuo *IncidentImageUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *IncidentImageUpdateOne {
	iiuo.modifiers = append(iiuo.modifiers, modifiers...)
	return iiuo
}

func (iiuo *IncidentImageUpdateOne) sqlSave(ctx context.Context) (_node *IncidentImage, err error) {
	_spec := sqlgraph.NewUpdateSpec(incidentimage.Table, incidentimage.Columns, sqlgraph.NewFieldSpec(incidentimage.FieldID, field.TypeInt))
	id, ok := iiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "IncidentImage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, incidentimage.FieldID)
		for _, f := range fields {
			if !incidentimage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != incidentimage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(incidentimage.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iiuo.mutation.Image(); ok {
		_spec.SetField(incidentimage.FieldImage, field.TypeString, value)
	}
	if iiuo.mutation.ImageCleared() {
		_spec.ClearField(incidentimage.FieldImage, field.TypeString)
	}
	if iiuo.mutation.IncidentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incidentimage.IncidentTable,
			Columns: []string{incidentimage.IncidentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(incident.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iiuo.mutation.IncidentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   incidentimage.IncidentTable,
			Columns: []string{incidentimage.IncidentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(incident.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(iiuo.modifiers...)
	_node = &IncidentImage{config: iiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{incidentimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iiuo.mutation.done = true
	return _node, nil
}
