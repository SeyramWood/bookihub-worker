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
	"github.com/SeyramWood/bookibus/ent/companyuser"
	"github.com/SeyramWood/bookibus/ent/parcel"
	"github.com/SeyramWood/bookibus/ent/parcelimage"
	"github.com/SeyramWood/bookibus/ent/predicate"
	"github.com/SeyramWood/bookibus/ent/transaction"
	"github.com/SeyramWood/bookibus/ent/trip"
)

// ParcelQuery is the builder for querying Parcel entities.
type ParcelQuery struct {
	config
	ctx             *QueryContext
	order           []parcel.OrderOption
	inters          []Interceptor
	predicates      []predicate.Parcel
	withImages      *ParcelImageQuery
	withTransaction *TransactionQuery
	withTrip        *TripQuery
	withCompany     *CompanyQuery
	withDriver      *CompanyUserQuery
	withFKs         bool
	modifiers       []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ParcelQuery builder.
func (pq *ParcelQuery) Where(ps ...predicate.Parcel) *ParcelQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *ParcelQuery) Limit(limit int) *ParcelQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *ParcelQuery) Offset(offset int) *ParcelQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ParcelQuery) Unique(unique bool) *ParcelQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *ParcelQuery) Order(o ...parcel.OrderOption) *ParcelQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryImages chains the current query on the "images" edge.
func (pq *ParcelQuery) QueryImages() *ParcelImageQuery {
	query := (&ParcelImageClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(parcel.Table, parcel.FieldID, selector),
			sqlgraph.To(parcelimage.Table, parcelimage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, parcel.ImagesTable, parcel.ImagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTransaction chains the current query on the "transaction" edge.
func (pq *ParcelQuery) QueryTransaction() *TransactionQuery {
	query := (&TransactionClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(parcel.Table, parcel.FieldID, selector),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, parcel.TransactionTable, parcel.TransactionColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTrip chains the current query on the "trip" edge.
func (pq *ParcelQuery) QueryTrip() *TripQuery {
	query := (&TripClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(parcel.Table, parcel.FieldID, selector),
			sqlgraph.To(trip.Table, trip.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, parcel.TripTable, parcel.TripColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCompany chains the current query on the "company" edge.
func (pq *ParcelQuery) QueryCompany() *CompanyQuery {
	query := (&CompanyClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(parcel.Table, parcel.FieldID, selector),
			sqlgraph.To(company.Table, company.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, parcel.CompanyTable, parcel.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDriver chains the current query on the "driver" edge.
func (pq *ParcelQuery) QueryDriver() *CompanyUserQuery {
	query := (&CompanyUserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(parcel.Table, parcel.FieldID, selector),
			sqlgraph.To(companyuser.Table, companyuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, parcel.DriverTable, parcel.DriverColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Parcel entity from the query.
// Returns a *NotFoundError when no Parcel was found.
func (pq *ParcelQuery) First(ctx context.Context) (*Parcel, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{parcel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ParcelQuery) FirstX(ctx context.Context) *Parcel {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Parcel ID from the query.
// Returns a *NotFoundError when no Parcel ID was found.
func (pq *ParcelQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{parcel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ParcelQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Parcel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Parcel entity is found.
// Returns a *NotFoundError when no Parcel entities are found.
func (pq *ParcelQuery) Only(ctx context.Context) (*Parcel, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{parcel.Label}
	default:
		return nil, &NotSingularError{parcel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ParcelQuery) OnlyX(ctx context.Context) *Parcel {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Parcel ID in the query.
// Returns a *NotSingularError when more than one Parcel ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ParcelQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{parcel.Label}
	default:
		err = &NotSingularError{parcel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ParcelQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Parcels.
func (pq *ParcelQuery) All(ctx context.Context) ([]*Parcel, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Parcel, *ParcelQuery]()
	return withInterceptors[[]*Parcel](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *ParcelQuery) AllX(ctx context.Context) []*Parcel {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Parcel IDs.
func (pq *ParcelQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(parcel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ParcelQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ParcelQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*ParcelQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ParcelQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ParcelQuery) Exist(ctx context.Context) (bool, error) {
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
func (pq *ParcelQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ParcelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ParcelQuery) Clone() *ParcelQuery {
	if pq == nil {
		return nil
	}
	return &ParcelQuery{
		config:          pq.config,
		ctx:             pq.ctx.Clone(),
		order:           append([]parcel.OrderOption{}, pq.order...),
		inters:          append([]Interceptor{}, pq.inters...),
		predicates:      append([]predicate.Parcel{}, pq.predicates...),
		withImages:      pq.withImages.Clone(),
		withTransaction: pq.withTransaction.Clone(),
		withTrip:        pq.withTrip.Clone(),
		withCompany:     pq.withCompany.Clone(),
		withDriver:      pq.withDriver.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithImages tells the query-builder to eager-load the nodes that are connected to
// the "images" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ParcelQuery) WithImages(opts ...func(*ParcelImageQuery)) *ParcelQuery {
	query := (&ParcelImageClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withImages = query
	return pq
}

// WithTransaction tells the query-builder to eager-load the nodes that are connected to
// the "transaction" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ParcelQuery) WithTransaction(opts ...func(*TransactionQuery)) *ParcelQuery {
	query := (&TransactionClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withTransaction = query
	return pq
}

// WithTrip tells the query-builder to eager-load the nodes that are connected to
// the "trip" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ParcelQuery) WithTrip(opts ...func(*TripQuery)) *ParcelQuery {
	query := (&TripClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withTrip = query
	return pq
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ParcelQuery) WithCompany(opts ...func(*CompanyQuery)) *ParcelQuery {
	query := (&CompanyClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withCompany = query
	return pq
}

// WithDriver tells the query-builder to eager-load the nodes that are connected to
// the "driver" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ParcelQuery) WithDriver(opts ...func(*CompanyUserQuery)) *ParcelQuery {
	query := (&CompanyUserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withDriver = query
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
//	client.Parcel.Query().
//		GroupBy(parcel.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *ParcelQuery) GroupBy(field string, fields ...string) *ParcelGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ParcelGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = parcel.Label
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
//	client.Parcel.Query().
//		Select(parcel.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *ParcelQuery) Select(fields ...string) *ParcelSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &ParcelSelect{ParcelQuery: pq}
	sbuild.label = parcel.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ParcelSelect configured with the given aggregations.
func (pq *ParcelQuery) Aggregate(fns ...AggregateFunc) *ParcelSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ParcelQuery) prepareQuery(ctx context.Context) error {
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
		if !parcel.ValidColumn(f) {
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

func (pq *ParcelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Parcel, error) {
	var (
		nodes       = []*Parcel{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [5]bool{
			pq.withImages != nil,
			pq.withTransaction != nil,
			pq.withTrip != nil,
			pq.withCompany != nil,
			pq.withDriver != nil,
		}
	)
	if pq.withTrip != nil || pq.withCompany != nil || pq.withDriver != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, parcel.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Parcel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Parcel{config: pq.config}
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
	if query := pq.withImages; query != nil {
		if err := pq.loadImages(ctx, query, nodes,
			func(n *Parcel) { n.Edges.Images = []*ParcelImage{} },
			func(n *Parcel, e *ParcelImage) { n.Edges.Images = append(n.Edges.Images, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withTransaction; query != nil {
		if err := pq.loadTransaction(ctx, query, nodes, nil,
			func(n *Parcel, e *Transaction) { n.Edges.Transaction = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withTrip; query != nil {
		if err := pq.loadTrip(ctx, query, nodes, nil,
			func(n *Parcel, e *Trip) { n.Edges.Trip = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withCompany; query != nil {
		if err := pq.loadCompany(ctx, query, nodes, nil,
			func(n *Parcel, e *Company) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withDriver; query != nil {
		if err := pq.loadDriver(ctx, query, nodes, nil,
			func(n *Parcel, e *CompanyUser) { n.Edges.Driver = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ParcelQuery) loadImages(ctx context.Context, query *ParcelImageQuery, nodes []*Parcel, init func(*Parcel), assign func(*Parcel, *ParcelImage)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Parcel)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ParcelImage(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(parcel.ImagesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.parcel_images
		if fk == nil {
			return fmt.Errorf(`foreign-key "parcel_images" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "parcel_images" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *ParcelQuery) loadTransaction(ctx context.Context, query *TransactionQuery, nodes []*Parcel, init func(*Parcel), assign func(*Parcel, *Transaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Parcel)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(parcel.TransactionColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.parcel_transaction
		if fk == nil {
			return fmt.Errorf(`foreign-key "parcel_transaction" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "parcel_transaction" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *ParcelQuery) loadTrip(ctx context.Context, query *TripQuery, nodes []*Parcel, init func(*Parcel), assign func(*Parcel, *Trip)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Parcel)
	for i := range nodes {
		if nodes[i].trip_parcels == nil {
			continue
		}
		fk := *nodes[i].trip_parcels
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(trip.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "trip_parcels" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *ParcelQuery) loadCompany(ctx context.Context, query *CompanyQuery, nodes []*Parcel, init func(*Parcel), assign func(*Parcel, *Company)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Parcel)
	for i := range nodes {
		if nodes[i].company_parcels == nil {
			continue
		}
		fk := *nodes[i].company_parcels
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
			return fmt.Errorf(`unexpected foreign-key "company_parcels" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *ParcelQuery) loadDriver(ctx context.Context, query *CompanyUserQuery, nodes []*Parcel, init func(*Parcel), assign func(*Parcel, *CompanyUser)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Parcel)
	for i := range nodes {
		if nodes[i].company_user_parcels == nil {
			continue
		}
		fk := *nodes[i].company_user_parcels
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(companyuser.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "company_user_parcels" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *ParcelQuery) sqlCount(ctx context.Context) (int, error) {
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

func (pq *ParcelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(parcel.Table, parcel.Columns, sqlgraph.NewFieldSpec(parcel.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, parcel.FieldID)
		for i := range fields {
			if fields[i] != parcel.FieldID {
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

func (pq *ParcelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(parcel.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = parcel.Columns
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
func (pq *ParcelQuery) Modify(modifiers ...func(s *sql.Selector)) *ParcelSelect {
	pq.modifiers = append(pq.modifiers, modifiers...)
	return pq.Select()
}

// ParcelGroupBy is the group-by builder for Parcel entities.
type ParcelGroupBy struct {
	selector
	build *ParcelQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ParcelGroupBy) Aggregate(fns ...AggregateFunc) *ParcelGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ParcelGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ParcelQuery, *ParcelGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ParcelGroupBy) sqlScan(ctx context.Context, root *ParcelQuery, v any) error {
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

// ParcelSelect is the builder for selecting fields of Parcel entities.
type ParcelSelect struct {
	*ParcelQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ParcelSelect) Aggregate(fns ...AggregateFunc) *ParcelSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ParcelSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ParcelQuery, *ParcelSelect](ctx, ps.ParcelQuery, ps, ps.inters, v)
}

func (ps *ParcelSelect) sqlScan(ctx context.Context, root *ParcelQuery, v any) error {
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
func (ps *ParcelSelect) Modify(modifiers ...func(s *sql.Selector)) *ParcelSelect {
	ps.modifiers = append(ps.modifiers, modifiers...)
	return ps
}
