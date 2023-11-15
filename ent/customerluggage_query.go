// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/bookibus/ent/booking"
	"github.com/SeyramWood/bookibus/ent/customerluggage"
	"github.com/SeyramWood/bookibus/ent/predicate"
)

// CustomerLuggageQuery is the builder for querying CustomerLuggage entities.
type CustomerLuggageQuery struct {
	config
	ctx         *QueryContext
	order       []customerluggage.OrderOption
	inters      []Interceptor
	predicates  []predicate.CustomerLuggage
	withBooking *BookingQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CustomerLuggageQuery builder.
func (clq *CustomerLuggageQuery) Where(ps ...predicate.CustomerLuggage) *CustomerLuggageQuery {
	clq.predicates = append(clq.predicates, ps...)
	return clq
}

// Limit the number of records to be returned by this query.
func (clq *CustomerLuggageQuery) Limit(limit int) *CustomerLuggageQuery {
	clq.ctx.Limit = &limit
	return clq
}

// Offset to start from.
func (clq *CustomerLuggageQuery) Offset(offset int) *CustomerLuggageQuery {
	clq.ctx.Offset = &offset
	return clq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (clq *CustomerLuggageQuery) Unique(unique bool) *CustomerLuggageQuery {
	clq.ctx.Unique = &unique
	return clq
}

// Order specifies how the records should be ordered.
func (clq *CustomerLuggageQuery) Order(o ...customerluggage.OrderOption) *CustomerLuggageQuery {
	clq.order = append(clq.order, o...)
	return clq
}

// QueryBooking chains the current query on the "booking" edge.
func (clq *CustomerLuggageQuery) QueryBooking() *BookingQuery {
	query := (&BookingClient{config: clq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := clq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := clq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customerluggage.Table, customerluggage.FieldID, selector),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, customerluggage.BookingTable, customerluggage.BookingColumn),
		)
		fromU = sqlgraph.SetNeighbors(clq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CustomerLuggage entity from the query.
// Returns a *NotFoundError when no CustomerLuggage was found.
func (clq *CustomerLuggageQuery) First(ctx context.Context) (*CustomerLuggage, error) {
	nodes, err := clq.Limit(1).All(setContextOp(ctx, clq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{customerluggage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (clq *CustomerLuggageQuery) FirstX(ctx context.Context) *CustomerLuggage {
	node, err := clq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CustomerLuggage ID from the query.
// Returns a *NotFoundError when no CustomerLuggage ID was found.
func (clq *CustomerLuggageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = clq.Limit(1).IDs(setContextOp(ctx, clq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{customerluggage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (clq *CustomerLuggageQuery) FirstIDX(ctx context.Context) int {
	id, err := clq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CustomerLuggage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CustomerLuggage entity is found.
// Returns a *NotFoundError when no CustomerLuggage entities are found.
func (clq *CustomerLuggageQuery) Only(ctx context.Context) (*CustomerLuggage, error) {
	nodes, err := clq.Limit(2).All(setContextOp(ctx, clq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{customerluggage.Label}
	default:
		return nil, &NotSingularError{customerluggage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (clq *CustomerLuggageQuery) OnlyX(ctx context.Context) *CustomerLuggage {
	node, err := clq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CustomerLuggage ID in the query.
// Returns a *NotSingularError when more than one CustomerLuggage ID is found.
// Returns a *NotFoundError when no entities are found.
func (clq *CustomerLuggageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = clq.Limit(2).IDs(setContextOp(ctx, clq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{customerluggage.Label}
	default:
		err = &NotSingularError{customerluggage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (clq *CustomerLuggageQuery) OnlyIDX(ctx context.Context) int {
	id, err := clq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CustomerLuggages.
func (clq *CustomerLuggageQuery) All(ctx context.Context) ([]*CustomerLuggage, error) {
	ctx = setContextOp(ctx, clq.ctx, "All")
	if err := clq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CustomerLuggage, *CustomerLuggageQuery]()
	return withInterceptors[[]*CustomerLuggage](ctx, clq, qr, clq.inters)
}

// AllX is like All, but panics if an error occurs.
func (clq *CustomerLuggageQuery) AllX(ctx context.Context) []*CustomerLuggage {
	nodes, err := clq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CustomerLuggage IDs.
func (clq *CustomerLuggageQuery) IDs(ctx context.Context) (ids []int, err error) {
	if clq.ctx.Unique == nil && clq.path != nil {
		clq.Unique(true)
	}
	ctx = setContextOp(ctx, clq.ctx, "IDs")
	if err = clq.Select(customerluggage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (clq *CustomerLuggageQuery) IDsX(ctx context.Context) []int {
	ids, err := clq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (clq *CustomerLuggageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, clq.ctx, "Count")
	if err := clq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, clq, querierCount[*CustomerLuggageQuery](), clq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (clq *CustomerLuggageQuery) CountX(ctx context.Context) int {
	count, err := clq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (clq *CustomerLuggageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, clq.ctx, "Exist")
	switch _, err := clq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (clq *CustomerLuggageQuery) ExistX(ctx context.Context) bool {
	exist, err := clq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CustomerLuggageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (clq *CustomerLuggageQuery) Clone() *CustomerLuggageQuery {
	if clq == nil {
		return nil
	}
	return &CustomerLuggageQuery{
		config:      clq.config,
		ctx:         clq.ctx.Clone(),
		order:       append([]customerluggage.OrderOption{}, clq.order...),
		inters:      append([]Interceptor{}, clq.inters...),
		predicates:  append([]predicate.CustomerLuggage{}, clq.predicates...),
		withBooking: clq.withBooking.Clone(),
		// clone intermediate query.
		sql:  clq.sql.Clone(),
		path: clq.path,
	}
}

// WithBooking tells the query-builder to eager-load the nodes that are connected to
// the "booking" edge. The optional arguments are used to configure the query builder of the edge.
func (clq *CustomerLuggageQuery) WithBooking(opts ...func(*BookingQuery)) *CustomerLuggageQuery {
	query := (&BookingClient{config: clq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	clq.withBooking = query
	return clq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CustomerLuggage.Query().
//		GroupBy(customerluggage.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (clq *CustomerLuggageQuery) GroupBy(field string, fields ...string) *CustomerLuggageGroupBy {
	clq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CustomerLuggageGroupBy{build: clq}
	grbuild.flds = &clq.ctx.Fields
	grbuild.label = customerluggage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.CustomerLuggage.Query().
//		Select(customerluggage.FieldCreatedAt).
//		Scan(ctx, &v)
func (clq *CustomerLuggageQuery) Select(fields ...string) *CustomerLuggageSelect {
	clq.ctx.Fields = append(clq.ctx.Fields, fields...)
	sbuild := &CustomerLuggageSelect{CustomerLuggageQuery: clq}
	sbuild.label = customerluggage.Label
	sbuild.flds, sbuild.scan = &clq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CustomerLuggageSelect configured with the given aggregations.
func (clq *CustomerLuggageQuery) Aggregate(fns ...AggregateFunc) *CustomerLuggageSelect {
	return clq.Select().Aggregate(fns...)
}

func (clq *CustomerLuggageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range clq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, clq); err != nil {
				return err
			}
		}
	}
	for _, f := range clq.ctx.Fields {
		if !customerluggage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if clq.path != nil {
		prev, err := clq.path(ctx)
		if err != nil {
			return err
		}
		clq.sql = prev
	}
	return nil
}

func (clq *CustomerLuggageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CustomerLuggage, error) {
	var (
		nodes       = []*CustomerLuggage{}
		withFKs     = clq.withFKs
		_spec       = clq.querySpec()
		loadedTypes = [1]bool{
			clq.withBooking != nil,
		}
	)
	if clq.withBooking != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, customerluggage.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CustomerLuggage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CustomerLuggage{config: clq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(clq.modifiers) > 0 {
		_spec.Modifiers = clq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, clq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := clq.withBooking; query != nil {
		if err := clq.loadBooking(ctx, query, nodes, nil,
			func(n *CustomerLuggage, e *Booking) { n.Edges.Booking = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (clq *CustomerLuggageQuery) loadBooking(ctx context.Context, query *BookingQuery, nodes []*CustomerLuggage, init func(*CustomerLuggage), assign func(*CustomerLuggage, *Booking)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CustomerLuggage)
	for i := range nodes {
		if nodes[i].booking_luggages == nil {
			continue
		}
		fk := *nodes[i].booking_luggages
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(booking.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "booking_luggages" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (clq *CustomerLuggageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := clq.querySpec()
	if len(clq.modifiers) > 0 {
		_spec.Modifiers = clq.modifiers
	}
	_spec.Node.Columns = clq.ctx.Fields
	if len(clq.ctx.Fields) > 0 {
		_spec.Unique = clq.ctx.Unique != nil && *clq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, clq.driver, _spec)
}

func (clq *CustomerLuggageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(customerluggage.Table, customerluggage.Columns, sqlgraph.NewFieldSpec(customerluggage.FieldID, field.TypeInt))
	_spec.From = clq.sql
	if unique := clq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if clq.path != nil {
		_spec.Unique = true
	}
	if fields := clq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customerluggage.FieldID)
		for i := range fields {
			if fields[i] != customerluggage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := clq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := clq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := clq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := clq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (clq *CustomerLuggageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(clq.driver.Dialect())
	t1 := builder.Table(customerluggage.Table)
	columns := clq.ctx.Fields
	if len(columns) == 0 {
		columns = customerluggage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if clq.sql != nil {
		selector = clq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if clq.ctx.Unique != nil && *clq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range clq.modifiers {
		m(selector)
	}
	for _, p := range clq.predicates {
		p(selector)
	}
	for _, p := range clq.order {
		p(selector)
	}
	if offset := clq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := clq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (clq *CustomerLuggageQuery) Modify(modifiers ...func(s *sql.Selector)) *CustomerLuggageSelect {
	clq.modifiers = append(clq.modifiers, modifiers...)
	return clq.Select()
}

// CustomerLuggageGroupBy is the group-by builder for CustomerLuggage entities.
type CustomerLuggageGroupBy struct {
	selector
	build *CustomerLuggageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (clgb *CustomerLuggageGroupBy) Aggregate(fns ...AggregateFunc) *CustomerLuggageGroupBy {
	clgb.fns = append(clgb.fns, fns...)
	return clgb
}

// Scan applies the selector query and scans the result into the given value.
func (clgb *CustomerLuggageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, clgb.build.ctx, "GroupBy")
	if err := clgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerLuggageQuery, *CustomerLuggageGroupBy](ctx, clgb.build, clgb, clgb.build.inters, v)
}

func (clgb *CustomerLuggageGroupBy) sqlScan(ctx context.Context, root *CustomerLuggageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(clgb.fns))
	for _, fn := range clgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*clgb.flds)+len(clgb.fns))
		for _, f := range *clgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*clgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := clgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CustomerLuggageSelect is the builder for selecting fields of CustomerLuggage entities.
type CustomerLuggageSelect struct {
	*CustomerLuggageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cls *CustomerLuggageSelect) Aggregate(fns ...AggregateFunc) *CustomerLuggageSelect {
	cls.fns = append(cls.fns, fns...)
	return cls
}

// Scan applies the selector query and scans the result into the given value.
func (cls *CustomerLuggageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cls.ctx, "Select")
	if err := cls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerLuggageQuery, *CustomerLuggageSelect](ctx, cls.CustomerLuggageQuery, cls, cls.inters, v)
}

func (cls *CustomerLuggageSelect) sqlScan(ctx context.Context, root *CustomerLuggageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cls.fns))
	for _, fn := range cls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cls *CustomerLuggageSelect) Modify(modifiers ...func(s *sql.Selector)) *CustomerLuggageSelect {
	cls.modifiers = append(cls.modifiers, modifiers...)
	return cls
}
