// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/route"
	"github.com/SeyramWood/ent/routestop"
)

// RouteStopQuery is the builder for querying RouteStop entities.
type RouteStopQuery struct {
	config
	ctx        *QueryContext
	order      []routestop.OrderOption
	inters     []Interceptor
	predicates []predicate.RouteStop
	withRoute  *RouteQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RouteStopQuery builder.
func (rsq *RouteStopQuery) Where(ps ...predicate.RouteStop) *RouteStopQuery {
	rsq.predicates = append(rsq.predicates, ps...)
	return rsq
}

// Limit the number of records to be returned by this query.
func (rsq *RouteStopQuery) Limit(limit int) *RouteStopQuery {
	rsq.ctx.Limit = &limit
	return rsq
}

// Offset to start from.
func (rsq *RouteStopQuery) Offset(offset int) *RouteStopQuery {
	rsq.ctx.Offset = &offset
	return rsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rsq *RouteStopQuery) Unique(unique bool) *RouteStopQuery {
	rsq.ctx.Unique = &unique
	return rsq
}

// Order specifies how the records should be ordered.
func (rsq *RouteStopQuery) Order(o ...routestop.OrderOption) *RouteStopQuery {
	rsq.order = append(rsq.order, o...)
	return rsq
}

// QueryRoute chains the current query on the "route" edge.
func (rsq *RouteStopQuery) QueryRoute() *RouteQuery {
	query := (&RouteClient{config: rsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routestop.Table, routestop.FieldID, selector),
			sqlgraph.To(route.Table, route.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, routestop.RouteTable, routestop.RouteColumn),
		)
		fromU = sqlgraph.SetNeighbors(rsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RouteStop entity from the query.
// Returns a *NotFoundError when no RouteStop was found.
func (rsq *RouteStopQuery) First(ctx context.Context) (*RouteStop, error) {
	nodes, err := rsq.Limit(1).All(setContextOp(ctx, rsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{routestop.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rsq *RouteStopQuery) FirstX(ctx context.Context) *RouteStop {
	node, err := rsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RouteStop ID from the query.
// Returns a *NotFoundError when no RouteStop ID was found.
func (rsq *RouteStopQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rsq.Limit(1).IDs(setContextOp(ctx, rsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{routestop.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rsq *RouteStopQuery) FirstIDX(ctx context.Context) int {
	id, err := rsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RouteStop entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RouteStop entity is found.
// Returns a *NotFoundError when no RouteStop entities are found.
func (rsq *RouteStopQuery) Only(ctx context.Context) (*RouteStop, error) {
	nodes, err := rsq.Limit(2).All(setContextOp(ctx, rsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{routestop.Label}
	default:
		return nil, &NotSingularError{routestop.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rsq *RouteStopQuery) OnlyX(ctx context.Context) *RouteStop {
	node, err := rsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RouteStop ID in the query.
// Returns a *NotSingularError when more than one RouteStop ID is found.
// Returns a *NotFoundError when no entities are found.
func (rsq *RouteStopQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rsq.Limit(2).IDs(setContextOp(ctx, rsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{routestop.Label}
	default:
		err = &NotSingularError{routestop.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rsq *RouteStopQuery) OnlyIDX(ctx context.Context) int {
	id, err := rsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RouteStops.
func (rsq *RouteStopQuery) All(ctx context.Context) ([]*RouteStop, error) {
	ctx = setContextOp(ctx, rsq.ctx, "All")
	if err := rsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RouteStop, *RouteStopQuery]()
	return withInterceptors[[]*RouteStop](ctx, rsq, qr, rsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rsq *RouteStopQuery) AllX(ctx context.Context) []*RouteStop {
	nodes, err := rsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RouteStop IDs.
func (rsq *RouteStopQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rsq.ctx.Unique == nil && rsq.path != nil {
		rsq.Unique(true)
	}
	ctx = setContextOp(ctx, rsq.ctx, "IDs")
	if err = rsq.Select(routestop.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rsq *RouteStopQuery) IDsX(ctx context.Context) []int {
	ids, err := rsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rsq *RouteStopQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rsq.ctx, "Count")
	if err := rsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rsq, querierCount[*RouteStopQuery](), rsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rsq *RouteStopQuery) CountX(ctx context.Context) int {
	count, err := rsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rsq *RouteStopQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rsq.ctx, "Exist")
	switch _, err := rsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rsq *RouteStopQuery) ExistX(ctx context.Context) bool {
	exist, err := rsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RouteStopQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rsq *RouteStopQuery) Clone() *RouteStopQuery {
	if rsq == nil {
		return nil
	}
	return &RouteStopQuery{
		config:     rsq.config,
		ctx:        rsq.ctx.Clone(),
		order:      append([]routestop.OrderOption{}, rsq.order...),
		inters:     append([]Interceptor{}, rsq.inters...),
		predicates: append([]predicate.RouteStop{}, rsq.predicates...),
		withRoute:  rsq.withRoute.Clone(),
		// clone intermediate query.
		sql:  rsq.sql.Clone(),
		path: rsq.path,
	}
}

// WithRoute tells the query-builder to eager-load the nodes that are connected to
// the "route" edge. The optional arguments are used to configure the query builder of the edge.
func (rsq *RouteStopQuery) WithRoute(opts ...func(*RouteQuery)) *RouteStopQuery {
	query := (&RouteClient{config: rsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rsq.withRoute = query
	return rsq
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
//	client.RouteStop.Query().
//		GroupBy(routestop.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rsq *RouteStopQuery) GroupBy(field string, fields ...string) *RouteStopGroupBy {
	rsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RouteStopGroupBy{build: rsq}
	grbuild.flds = &rsq.ctx.Fields
	grbuild.label = routestop.Label
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
//	client.RouteStop.Query().
//		Select(routestop.FieldCreatedAt).
//		Scan(ctx, &v)
func (rsq *RouteStopQuery) Select(fields ...string) *RouteStopSelect {
	rsq.ctx.Fields = append(rsq.ctx.Fields, fields...)
	sbuild := &RouteStopSelect{RouteStopQuery: rsq}
	sbuild.label = routestop.Label
	sbuild.flds, sbuild.scan = &rsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RouteStopSelect configured with the given aggregations.
func (rsq *RouteStopQuery) Aggregate(fns ...AggregateFunc) *RouteStopSelect {
	return rsq.Select().Aggregate(fns...)
}

func (rsq *RouteStopQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rsq); err != nil {
				return err
			}
		}
	}
	for _, f := range rsq.ctx.Fields {
		if !routestop.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rsq.path != nil {
		prev, err := rsq.path(ctx)
		if err != nil {
			return err
		}
		rsq.sql = prev
	}
	return nil
}

func (rsq *RouteStopQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RouteStop, error) {
	var (
		nodes       = []*RouteStop{}
		withFKs     = rsq.withFKs
		_spec       = rsq.querySpec()
		loadedTypes = [1]bool{
			rsq.withRoute != nil,
		}
	)
	if rsq.withRoute != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, routestop.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RouteStop).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RouteStop{config: rsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rsq.modifiers) > 0 {
		_spec.Modifiers = rsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rsq.withRoute; query != nil {
		if err := rsq.loadRoute(ctx, query, nodes, nil,
			func(n *RouteStop, e *Route) { n.Edges.Route = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rsq *RouteStopQuery) loadRoute(ctx context.Context, query *RouteQuery, nodes []*RouteStop, init func(*RouteStop), assign func(*RouteStop, *Route)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*RouteStop)
	for i := range nodes {
		if nodes[i].route_stops == nil {
			continue
		}
		fk := *nodes[i].route_stops
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(route.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "route_stops" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rsq *RouteStopQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rsq.querySpec()
	if len(rsq.modifiers) > 0 {
		_spec.Modifiers = rsq.modifiers
	}
	_spec.Node.Columns = rsq.ctx.Fields
	if len(rsq.ctx.Fields) > 0 {
		_spec.Unique = rsq.ctx.Unique != nil && *rsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rsq.driver, _spec)
}

func (rsq *RouteStopQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(routestop.Table, routestop.Columns, sqlgraph.NewFieldSpec(routestop.FieldID, field.TypeInt))
	_spec.From = rsq.sql
	if unique := rsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rsq.path != nil {
		_spec.Unique = true
	}
	if fields := rsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routestop.FieldID)
		for i := range fields {
			if fields[i] != routestop.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rsq *RouteStopQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rsq.driver.Dialect())
	t1 := builder.Table(routestop.Table)
	columns := rsq.ctx.Fields
	if len(columns) == 0 {
		columns = routestop.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rsq.sql != nil {
		selector = rsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rsq.ctx.Unique != nil && *rsq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range rsq.modifiers {
		m(selector)
	}
	for _, p := range rsq.predicates {
		p(selector)
	}
	for _, p := range rsq.order {
		p(selector)
	}
	if offset := rsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rsq *RouteStopQuery) Modify(modifiers ...func(s *sql.Selector)) *RouteStopSelect {
	rsq.modifiers = append(rsq.modifiers, modifiers...)
	return rsq.Select()
}

// RouteStopGroupBy is the group-by builder for RouteStop entities.
type RouteStopGroupBy struct {
	selector
	build *RouteStopQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rsgb *RouteStopGroupBy) Aggregate(fns ...AggregateFunc) *RouteStopGroupBy {
	rsgb.fns = append(rsgb.fns, fns...)
	return rsgb
}

// Scan applies the selector query and scans the result into the given value.
func (rsgb *RouteStopGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rsgb.build.ctx, "GroupBy")
	if err := rsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RouteStopQuery, *RouteStopGroupBy](ctx, rsgb.build, rsgb, rsgb.build.inters, v)
}

func (rsgb *RouteStopGroupBy) sqlScan(ctx context.Context, root *RouteStopQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rsgb.fns))
	for _, fn := range rsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rsgb.flds)+len(rsgb.fns))
		for _, f := range *rsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RouteStopSelect is the builder for selecting fields of RouteStop entities.
type RouteStopSelect struct {
	*RouteStopQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rss *RouteStopSelect) Aggregate(fns ...AggregateFunc) *RouteStopSelect {
	rss.fns = append(rss.fns, fns...)
	return rss
}

// Scan applies the selector query and scans the result into the given value.
func (rss *RouteStopSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rss.ctx, "Select")
	if err := rss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RouteStopQuery, *RouteStopSelect](ctx, rss.RouteStopQuery, rss, rss.inters, v)
}

func (rss *RouteStopSelect) sqlScan(ctx context.Context, root *RouteStopQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rss.fns))
	for _, fn := range rss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rss *RouteStopSelect) Modify(modifiers ...func(s *sql.Selector)) *RouteStopSelect {
	rss.modifiers = append(rss.modifiers, modifiers...)
	return rss
}
