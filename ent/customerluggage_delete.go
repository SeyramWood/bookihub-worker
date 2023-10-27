// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/customerluggage"
	"github.com/SeyramWood/ent/predicate"
)

// CustomerLuggageDelete is the builder for deleting a CustomerLuggage entity.
type CustomerLuggageDelete struct {
	config
	hooks    []Hook
	mutation *CustomerLuggageMutation
}

// Where appends a list predicates to the CustomerLuggageDelete builder.
func (cld *CustomerLuggageDelete) Where(ps ...predicate.CustomerLuggage) *CustomerLuggageDelete {
	cld.mutation.Where(ps...)
	return cld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cld *CustomerLuggageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cld.sqlExec, cld.mutation, cld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cld *CustomerLuggageDelete) ExecX(ctx context.Context) int {
	n, err := cld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cld *CustomerLuggageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(customerluggage.Table, sqlgraph.NewFieldSpec(customerluggage.FieldID, field.TypeInt))
	if ps := cld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cld.mutation.done = true
	return affected, err
}

// CustomerLuggageDeleteOne is the builder for deleting a single CustomerLuggage entity.
type CustomerLuggageDeleteOne struct {
	cld *CustomerLuggageDelete
}

// Where appends a list predicates to the CustomerLuggageDelete builder.
func (cldo *CustomerLuggageDeleteOne) Where(ps ...predicate.CustomerLuggage) *CustomerLuggageDeleteOne {
	cldo.cld.mutation.Where(ps...)
	return cldo
}

// Exec executes the deletion query.
func (cldo *CustomerLuggageDeleteOne) Exec(ctx context.Context) error {
	n, err := cldo.cld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{customerluggage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cldo *CustomerLuggageDeleteOne) ExecX(ctx context.Context) {
	if err := cldo.Exec(ctx); err != nil {
		panic(err)
	}
}