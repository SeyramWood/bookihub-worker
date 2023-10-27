// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/route"
	"github.com/SeyramWood/ent/routestop"
)

// RouteStopCreate is the builder for creating a RouteStop entity.
type RouteStopCreate struct {
	config
	mutation *RouteStopMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rsc *RouteStopCreate) SetCreatedAt(t time.Time) *RouteStopCreate {
	rsc.mutation.SetCreatedAt(t)
	return rsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rsc *RouteStopCreate) SetNillableCreatedAt(t *time.Time) *RouteStopCreate {
	if t != nil {
		rsc.SetCreatedAt(*t)
	}
	return rsc
}

// SetUpdatedAt sets the "updated_at" field.
func (rsc *RouteStopCreate) SetUpdatedAt(t time.Time) *RouteStopCreate {
	rsc.mutation.SetUpdatedAt(t)
	return rsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rsc *RouteStopCreate) SetNillableUpdatedAt(t *time.Time) *RouteStopCreate {
	if t != nil {
		rsc.SetUpdatedAt(*t)
	}
	return rsc
}

// SetLatitude sets the "latitude" field.
func (rsc *RouteStopCreate) SetLatitude(f float64) *RouteStopCreate {
	rsc.mutation.SetLatitude(f)
	return rsc
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (rsc *RouteStopCreate) SetNillableLatitude(f *float64) *RouteStopCreate {
	if f != nil {
		rsc.SetLatitude(*f)
	}
	return rsc
}

// SetLongitude sets the "longitude" field.
func (rsc *RouteStopCreate) SetLongitude(f float64) *RouteStopCreate {
	rsc.mutation.SetLongitude(f)
	return rsc
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (rsc *RouteStopCreate) SetNillableLongitude(f *float64) *RouteStopCreate {
	if f != nil {
		rsc.SetLongitude(*f)
	}
	return rsc
}

// SetRouteID sets the "route" edge to the Route entity by ID.
func (rsc *RouteStopCreate) SetRouteID(id int) *RouteStopCreate {
	rsc.mutation.SetRouteID(id)
	return rsc
}

// SetNillableRouteID sets the "route" edge to the Route entity by ID if the given value is not nil.
func (rsc *RouteStopCreate) SetNillableRouteID(id *int) *RouteStopCreate {
	if id != nil {
		rsc = rsc.SetRouteID(*id)
	}
	return rsc
}

// SetRoute sets the "route" edge to the Route entity.
func (rsc *RouteStopCreate) SetRoute(r *Route) *RouteStopCreate {
	return rsc.SetRouteID(r.ID)
}

// Mutation returns the RouteStopMutation object of the builder.
func (rsc *RouteStopCreate) Mutation() *RouteStopMutation {
	return rsc.mutation
}

// Save creates the RouteStop in the database.
func (rsc *RouteStopCreate) Save(ctx context.Context) (*RouteStop, error) {
	rsc.defaults()
	return withHooks(ctx, rsc.sqlSave, rsc.mutation, rsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rsc *RouteStopCreate) SaveX(ctx context.Context) *RouteStop {
	v, err := rsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rsc *RouteStopCreate) Exec(ctx context.Context) error {
	_, err := rsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsc *RouteStopCreate) ExecX(ctx context.Context) {
	if err := rsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rsc *RouteStopCreate) defaults() {
	if _, ok := rsc.mutation.CreatedAt(); !ok {
		v := routestop.DefaultCreatedAt()
		rsc.mutation.SetCreatedAt(v)
	}
	if _, ok := rsc.mutation.UpdatedAt(); !ok {
		v := routestop.DefaultUpdatedAt()
		rsc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rsc *RouteStopCreate) check() error {
	if _, ok := rsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RouteStop.created_at"`)}
	}
	if _, ok := rsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "RouteStop.updated_at"`)}
	}
	return nil
}

func (rsc *RouteStopCreate) sqlSave(ctx context.Context) (*RouteStop, error) {
	if err := rsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rsc.mutation.id = &_node.ID
	rsc.mutation.done = true
	return _node, nil
}

func (rsc *RouteStopCreate) createSpec() (*RouteStop, *sqlgraph.CreateSpec) {
	var (
		_node = &RouteStop{config: rsc.config}
		_spec = sqlgraph.NewCreateSpec(routestop.Table, sqlgraph.NewFieldSpec(routestop.FieldID, field.TypeInt))
	)
	if value, ok := rsc.mutation.CreatedAt(); ok {
		_spec.SetField(routestop.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rsc.mutation.UpdatedAt(); ok {
		_spec.SetField(routestop.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rsc.mutation.Latitude(); ok {
		_spec.SetField(routestop.FieldLatitude, field.TypeFloat64, value)
		_node.Latitude = value
	}
	if value, ok := rsc.mutation.Longitude(); ok {
		_spec.SetField(routestop.FieldLongitude, field.TypeFloat64, value)
		_node.Longitude = value
	}
	if nodes := rsc.mutation.RouteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routestop.RouteTable,
			Columns: []string{routestop.RouteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.route_stops = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RouteStopCreateBulk is the builder for creating many RouteStop entities in bulk.
type RouteStopCreateBulk struct {
	config
	err      error
	builders []*RouteStopCreate
}

// Save creates the RouteStop entities in the database.
func (rscb *RouteStopCreateBulk) Save(ctx context.Context) ([]*RouteStop, error) {
	if rscb.err != nil {
		return nil, rscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rscb.builders))
	nodes := make([]*RouteStop, len(rscb.builders))
	mutators := make([]Mutator, len(rscb.builders))
	for i := range rscb.builders {
		func(i int, root context.Context) {
			builder := rscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RouteStopMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rscb *RouteStopCreateBulk) SaveX(ctx context.Context) []*RouteStop {
	v, err := rscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rscb *RouteStopCreateBulk) Exec(ctx context.Context) error {
	_, err := rscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rscb *RouteStopCreateBulk) ExecX(ctx context.Context) {
	if err := rscb.Exec(ctx); err != nil {
		panic(err)
	}
}
