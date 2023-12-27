// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/bookibus/ent/company"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/route"
	"github.com/SeyramWood/bookibus/ent/trip"
)

// RouteQuery is the builder for querying Route entities.
type RouteQuery struct {
	config
	ctx         *QueryContext
	order       []route.OrderOption
	inters      []Interceptor
	predicates  []predicate.Route
	withCompany *CompanyQuery
	withTrips   *TripQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RouteQuery builder.
func (rq *RouteQuery) Where(ps ...predicate.Route) *RouteQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RouteQuery) Limit(limit int) *RouteQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RouteQuery) Offset(offset int) *RouteQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RouteQuery) Unique(unique bool) *RouteQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RouteQuery) Order(o ...route.OrderOption) *RouteQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryCompany chains the current query on the "company" edge.
func (rq *RouteQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(route.Table, route.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, route.CompanyTable, route.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTrips chains the current query on the "trips" edge.
func (rq *RouteQuery) QueryTrips() *TripQuery {
	query := (&TripClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(route.Table, route.FieldID, selector),
			sqlgraph.To(trip.Table, trip.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, route.TripsTable, route.TripsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Route entity from the query.
// Returns a *NotFoundError when no Route was found.
func (rq *RouteQuery) First(ctx context.Context) (*Route, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{route.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RouteQuery) FirstX(ctx context.Context) *Route {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Route ID from the query.
// Returns a *NotFoundError when no Route ID was found.
func (rq *RouteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{route.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RouteQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Route entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Route entity is found.
// Returns a *NotFoundError when no Route entities are found.
func (rq *RouteQuery) Only(ctx context.Context) (*Route, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{route.Label}
	default:
		return nil, &NotSingularError{route.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RouteQuery) OnlyX(ctx context.Context) *Route {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Route ID in the query.
// Returns a *NotSingularError when more than one Route ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RouteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{route.Label}
	default:
		err = &NotSingularError{route.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RouteQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Routes.
func (rq *RouteQuery) All(ctx context.Context) ([]*Route, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Route, *RouteQuery]()
	return withInterceptors[[]*Route](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RouteQuery) AllX(ctx context.Context) []*Route {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Route IDs.
func (rq *RouteQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(route.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RouteQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RouteQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RouteQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RouteQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RouteQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RouteQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RouteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RouteQuery) Clone() *RouteQuery {
	if rq == nil {
		return nil
	}
	return &RouteQuery{
		config:      rq.config,
		ctx:         rq.ctx.Clone(),
		order:       append([]route.OrderOption{}, rq.order...),
		inters:      append([]Interceptor{}, rq.inters...),
		predicates:  append([]predicate.Route{}, rq.predicates...),
		withCompany: rq.withCompany.Clone(),
		withTrips:   rq.withTrips.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RouteQuery) WithCompany(opts ...func(*CompanyQuery)) *RouteQuery {
	query := (&CompanyClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withCompany = query
	return rq
}

// WithTrips tells the query-builder to eager-load the nodes that are connected to
// the "trips" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RouteQuery) WithTrips(opts ...func(*TripQuery)) *RouteQuery {
	query := (&TripClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withTrips = query
	return rq
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
//	client.Route.Query().
//		GroupBy(route.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RouteQuery) GroupBy(field string, fields ...string) *RouteGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RouteGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = route.Label
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
//	client.Route.Query().
//		Select(route.FieldCreatedAt).
//		Scan(ctx, &v)
func (rq *RouteQuery) Select(fields ...string) *RouteSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RouteSelect{RouteQuery: rq}
	sbuild.label = route.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RouteSelect configured with the given aggregations.
func (rq *RouteQuery) Aggregate(fns ...AggregateFunc) *RouteSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RouteQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !route.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RouteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Route, error) {
	var (
		nodes       = []*Route{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withCompany != nil,
			rq.withTrips != nil,
		}
	)
	if rq.withCompany != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, route.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Route).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Route{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withCompany; query != nil {
		if err := rq.loadCompany(ctx, query, nodes, nil,
			func(n *Route, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withTrips; query != nil {
		if err := rq.loadTrips(ctx, query, nodes,
			func(n *Route) { n.Edges.Trips = []*Trip{} },
			func(n *Route, e *Trip) { n.Edges.Trips = append(n.Edges.Trips, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RouteQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*Route, init func(*Route), assign func(*Route, *Company)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Route)
	for i := range nodes {
		if nodes[i].company_routes == nil {
			continue
		}
		fk := *nodes[i].company_routes
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(company.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "company_routes" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RouteQuery) loadTrips(ctx context.Context, query *TripQuery, nodes []*Route, init func(*Route), assign func(*Route, *Trip)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Route)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Trip(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(route.TripsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.route_trips
		if fk == nil {
			return fmt.Errorf(`foreign-key "route_trips" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "route_trips" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *RouteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RouteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(route.Table, route.Columns, sqlgraph.NewFieldSpec(route.FieldID, field.TypeInt))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, route.FieldID)
		for i := range fields {
			if fields[i] != route.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RouteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(route.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = route.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range rq.modifiers {
		m(selector)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rq *RouteQuery) Modify(modifiers ...func(s *sql.Selector)) *RouteSelect {
	rq.modifiers = append(rq.modifiers, modifiers...)
	return rq.Select()
}

// RouteGroupBy is the group-by builder for Route entities.
type RouteGroupBy struct {
	selector
	build *RouteQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RouteGroupBy) Aggregate(fns ...AggregateFunc) *RouteGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RouteGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RouteQuery, *RouteGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RouteGroupBy) sqlScan(ctx context.Context, root *RouteQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RouteSelect is the builder for selecting fields of Route entities.
type RouteSelect struct {
	*RouteQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RouteSelect) Aggregate(fns ...AggregateFunc) *RouteSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RouteSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RouteQuery, *RouteSelect](ctx, rs.RouteQuery, rs, rs.inters, v)
}

func (rs *RouteSelect) sqlScan(ctx context.Context, root *RouteQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rs *RouteSelect) Modify(modifiers ...func(s *sql.Selector)) *RouteSelect {
	rs.modifiers = append(rs.modifiers, modifiers...)
	return rs
}
