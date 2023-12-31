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
	"github.com/SeyramWood/bookibus/ent/passenger"
	"github.com/SeyramWood/bookibus/ent/predicate"
)

// PassengerQuery is the builder for querying Passenger entities.
type PassengerQuery struct {
	config
	ctx         *QueryContext
	order       []passenger.OrderOption
	inters      []Interceptor
	predicates  []predicate.Passenger
	withBooking *BookingQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PassengerQuery builder.
func (pq *PassengerQuery) Where(ps ...predicate.Passenger) *PassengerQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PassengerQuery) Limit(limit int) *PassengerQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PassengerQuery) Offset(offset int) *PassengerQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PassengerQuery) Unique(unique bool) *PassengerQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PassengerQuery) Order(o ...passenger.OrderOption) *PassengerQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryBooking chains the current query on the "booking" edge.
func (pq *PassengerQuery) QueryBooking() *BookingQuery {
	query := (&BookingClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(passenger.Table, passenger.FieldID, selector),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, passenger.BookingTable, passenger.BookingColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Passenger entity from the query.
// Returns a *NotFoundError when no Passenger was found.
func (pq *PassengerQuery) First(ctx context.Context) (*Passenger, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{passenger.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PassengerQuery) FirstX(ctx context.Context) *Passenger {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Passenger ID from the query.
// Returns a *NotFoundError when no Passenger ID was found.
func (pq *PassengerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{passenger.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PassengerQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Passenger entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Passenger entity is found.
// Returns a *NotFoundError when no Passenger entities are found.
func (pq *PassengerQuery) Only(ctx context.Context) (*Passenger, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{passenger.Label}
	default:
		return nil, &NotSingularError{passenger.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PassengerQuery) OnlyX(ctx context.Context) *Passenger {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Passenger ID in the query.
// Returns a *NotSingularError when more than one Passenger ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PassengerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{passenger.Label}
	default:
		err = &NotSingularError{passenger.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PassengerQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Passengers.
func (pq *PassengerQuery) All(ctx context.Context) ([]*Passenger, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Passenger, *PassengerQuery]()
	return withInterceptors[[]*Passenger](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PassengerQuery) AllX(ctx context.Context) []*Passenger {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Passenger IDs.
func (pq *PassengerQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(passenger.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PassengerQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PassengerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PassengerQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PassengerQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PassengerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PassengerQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PassengerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PassengerQuery) Clone() *PassengerQuery {
	if pq == nil {
		return nil
	}
	return &PassengerQuery{
		config:      pq.config,
		ctx:         pq.ctx.Clone(),
		order:       append([]passenger.OrderOption{}, pq.order...),
		inters:      append([]Interceptor{}, pq.inters...),
		predicates:  append([]predicate.Passenger{}, pq.predicates...),
		withBooking: pq.withBooking.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithBooking tells the query-builder to eager-load the nodes that are connected to
// the "booking" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PassengerQuery) WithBooking(opts ...func(*BookingQuery)) *PassengerQuery {
	query := (&BookingClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withBooking = query
	return pq
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
//	client.Passenger.Query().
//		GroupBy(passenger.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PassengerQuery) GroupBy(field string, fields ...string) *PassengerGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PassengerGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = passenger.Label
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
//	client.Passenger.Query().
//		Select(passenger.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *PassengerQuery) Select(fields ...string) *PassengerSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PassengerSelect{PassengerQuery: pq}
	sbuild.label = passenger.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PassengerSelect configured with the given aggregations.
func (pq *PassengerQuery) Aggregate(fns ...AggregateFunc) *PassengerSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PassengerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !passenger.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PassengerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Passenger, error) {
	var (
		nodes       = []*Passenger{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withBooking != nil,
		}
	)
	if pq.withBooking != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, passenger.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Passenger).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Passenger{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withBooking; query != nil {
		if err := pq.loadBooking(ctx, query, nodes, nil,
			func(n *Passenger, e *Booking) { n.Edges.Booking = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PassengerQuery) loadBooking(ctx context.Context, query *BookingQuery, nodes []*Passenger, init func(*Passenger), assign func(*Passenger, *Booking)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Passenger)
	for i := range nodes {
		if nodes[i].booking_passengers == nil {
			continue
		}
		fk := *nodes[i].booking_passengers
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
			return fmt.Errorf(`unexpected foreign-key "booking_passengers" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PassengerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PassengerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(passenger.Table, passenger.Columns, sqlgraph.NewFieldSpec(passenger.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passenger.FieldID)
		for i := range fields {
			if fields[i] != passenger.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PassengerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(passenger.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = passenger.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range pq.modifiers {
		m(selector)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pq *PassengerQuery) Modify(modifiers ...func(s *sql.Selector)) *PassengerSelect {
	pq.modifiers = append(pq.modifiers, modifiers...)
	return pq.Select()
}

// PassengerGroupBy is the group-by builder for Passenger entities.
type PassengerGroupBy struct {
	selector
	build *PassengerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PassengerGroupBy) Aggregate(fns ...AggregateFunc) *PassengerGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PassengerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PassengerQuery, *PassengerGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PassengerGroupBy) sqlScan(ctx context.Context, root *PassengerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PassengerSelect is the builder for selecting fields of Passenger entities.
type PassengerSelect struct {
	*PassengerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PassengerSelect) Aggregate(fns ...AggregateFunc) *PassengerSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PassengerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PassengerQuery, *PassengerSelect](ctx, ps.PassengerQuery, ps, ps.inters, v)
}

func (ps *PassengerSelect) sqlScan(ctx context.Context, root *PassengerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ps *PassengerSelect) Modify(modifiers ...func(s *sql.Selector)) *PassengerSelect {
	ps.modifiers = append(ps.modifiers, modifiers...)
	return ps
}
