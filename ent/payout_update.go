// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/bookibus/ent/payout"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/transaction"
)

// PayoutUpdate is the builder for updating Payout entities.
type PayoutUpdate struct {
	config
	hooks     []Hook
	mutation  *PayoutMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PayoutUpdate builder.
func (pu *PayoutUpdate) Where(ps ...predicate.Payout) *PayoutUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetAccountName sets the "account_name" field.
func (pu *PayoutUpdate) SetAccountName(s string) *PayoutUpdate {
	pu.mutation.SetAccountName(s)
	return pu
}

// SetAccountNumber sets the "account_number" field.
func (pu *PayoutUpdate) SetAccountNumber(s string) *PayoutUpdate {
	pu.mutation.SetAccountNumber(s)
	return pu
}

// SetBank sets the "bank" field.
func (pu *PayoutUpdate) SetBank(s string) *PayoutUpdate {
	pu.mutation.SetBank(s)
	return pu
}

// SetBranch sets the "branch" field.
func (pu *PayoutUpdate) SetBranch(s string) *PayoutUpdate {
	pu.mutation.SetBranch(s)
	return pu
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (pu *PayoutUpdate) SetTransactionID(id int) *PayoutUpdate {
	pu.mutation.SetTransactionID(id)
	return pu
}

// SetNillableTransactionID sets the "transaction" edge to the Transaction entity by ID if the given value is not nil.
func (pu *PayoutUpdate) SetNillableTransactionID(id *int) *PayoutUpdate {
	if id != nil {
		pu = pu.SetTransactionID(*id)
	}
	return pu
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (pu *PayoutUpdate) SetTransaction(t *Transaction) *PayoutUpdate {
	return pu.SetTransactionID(t.ID)
}

// Mutation returns the PayoutMutation object of the builder.
func (pu *PayoutUpdate) Mutation() *PayoutMutation {
	return pu.mutation
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (pu *PayoutUpdate) ClearTransaction() *PayoutUpdate {
	pu.mutation.ClearTransaction()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PayoutUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PayoutUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PayoutUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PayoutUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PayoutUpdate) check() error {
	if v, ok := pu.mutation.AccountName(); ok {
		if err := payout.AccountNameValidator(v); err != nil {
			return &ValidationError{Name: "account_name", err: fmt.Errorf(`ent: validator failed for field "Payout.account_name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.AccountNumber(); ok {
		if err := payout.AccountNumberValidator(v); err != nil {
			return &ValidationError{Name: "account_number", err: fmt.Errorf(`ent: validator failed for field "Payout.account_number": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Bank(); ok {
		if err := payout.BankValidator(v); err != nil {
			return &ValidationError{Name: "bank", err: fmt.Errorf(`ent: validator failed for field "Payout.bank": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Branch(); ok {
		if err := payout.BranchValidator(v); err != nil {
			return &ValidationError{Name: "branch", err: fmt.Errorf(`ent: validator failed for field "Payout.branch": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PayoutUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PayoutUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PayoutUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(payout.Table, payout.Columns, sqlgraph.NewFieldSpec(payout.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.AccountName(); ok {
		_spec.SetField(payout.FieldAccountName, field.TypeString, value)
	}
	if value, ok := pu.mutation.AccountNumber(); ok {
		_spec.SetField(payout.FieldAccountNumber, field.TypeString, value)
	}
	if value, ok := pu.mutation.Bank(); ok {
		_spec.SetField(payout.FieldBank, field.TypeString, value)
	}
	if value, ok := pu.mutation.Branch(); ok {
		_spec.SetField(payout.FieldBranch, field.TypeString, value)
	}
	if pu.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payout.TransactionTable,
			Columns: []string{payout.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payout.TransactionTable,
			Columns: []string{payout.TransactionColumn},
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
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payout.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PayoutUpdateOne is the builder for updating a single Payout entity.
type PayoutUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PayoutMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetAccountName sets the "account_name" field.
func (puo *PayoutUpdateOne) SetAccountName(s string) *PayoutUpdateOne {
	puo.mutation.SetAccountName(s)
	return puo
}

// SetAccountNumber sets the "account_number" field.
func (puo *PayoutUpdateOne) SetAccountNumber(s string) *PayoutUpdateOne {
	puo.mutation.SetAccountNumber(s)
	return puo
}

// SetBank sets the "bank" field.
func (puo *PayoutUpdateOne) SetBank(s string) *PayoutUpdateOne {
	puo.mutation.SetBank(s)
	return puo
}

// SetBranch sets the "branch" field.
func (puo *PayoutUpdateOne) SetBranch(s string) *PayoutUpdateOne {
	puo.mutation.SetBranch(s)
	return puo
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (puo *PayoutUpdateOne) SetTransactionID(id int) *PayoutUpdateOne {
	puo.mutation.SetTransactionID(id)
	return puo
}

// SetNillableTransactionID sets the "transaction" edge to the Transaction entity by ID if the given value is not nil.
func (puo *PayoutUpdateOne) SetNillableTransactionID(id *int) *PayoutUpdateOne {
	if id != nil {
		puo = puo.SetTransactionID(*id)
	}
	return puo
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (puo *PayoutUpdateOne) SetTransaction(t *Transaction) *PayoutUpdateOne {
	return puo.SetTransactionID(t.ID)
}

// Mutation returns the PayoutMutation object of the builder.
func (puo *PayoutUpdateOne) Mutation() *PayoutMutation {
	return puo.mutation
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (puo *PayoutUpdateOne) ClearTransaction() *PayoutUpdateOne {
	puo.mutation.ClearTransaction()
	return puo
}

// Where appends a list predicates to the PayoutUpdate builder.
func (puo *PayoutUpdateOne) Where(ps ...predicate.Payout) *PayoutUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PayoutUpdateOne) Select(field string, fields ...string) *PayoutUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payout entity.
func (puo *PayoutUpdateOne) Save(ctx context.Context) (*Payout, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PayoutUpdateOne) SaveX(ctx context.Context) *Payout {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PayoutUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PayoutUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PayoutUpdateOne) check() error {
	if v, ok := puo.mutation.AccountName(); ok {
		if err := payout.AccountNameValidator(v); err != nil {
			return &ValidationError{Name: "account_name", err: fmt.Errorf(`ent: validator failed for field "Payout.account_name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.AccountNumber(); ok {
		if err := payout.AccountNumberValidator(v); err != nil {
			return &ValidationError{Name: "account_number", err: fmt.Errorf(`ent: validator failed for field "Payout.account_number": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Bank(); ok {
		if err := payout.BankValidator(v); err != nil {
			return &ValidationError{Name: "bank", err: fmt.Errorf(`ent: validator failed for field "Payout.bank": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Branch(); ok {
		if err := payout.BranchValidator(v); err != nil {
			return &ValidationError{Name: "branch", err: fmt.Errorf(`ent: validator failed for field "Payout.branch": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PayoutUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PayoutUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PayoutUpdateOne) sqlSave(ctx context.Context) (_node *Payout, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(payout.Table, payout.Columns, sqlgraph.NewFieldSpec(payout.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Payout.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payout.FieldID)
		for _, f := range fields {
			if !payout.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payout.FieldID {
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
	if value, ok := puo.mutation.AccountName(); ok {
		_spec.SetField(payout.FieldAccountName, field.TypeString, value)
	}
	if value, ok := puo.mutation.AccountNumber(); ok {
		_spec.SetField(payout.FieldAccountNumber, field.TypeString, value)
	}
	if value, ok := puo.mutation.Bank(); ok {
		_spec.SetField(payout.FieldBank, field.TypeString, value)
	}
	if value, ok := puo.mutation.Branch(); ok {
		_spec.SetField(payout.FieldBranch, field.TypeString, value)
	}
	if puo.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payout.TransactionTable,
			Columns: []string{payout.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payout.TransactionTable,
			Columns: []string{payout.TransactionColumn},
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
	_spec.AddModifiers(puo.modifiers...)
	_node = &Payout{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payout.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
