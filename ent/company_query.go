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
	"github.com/SeyramWood/ent/booking"
	"github.com/SeyramWood/ent/company"
	"github.com/SeyramWood/ent/companyuser"
	"github.com/SeyramWood/ent/notification"
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/route"
	"github.com/SeyramWood/ent/trip"
	"github.com/SeyramWood/ent/vehicle"
)

// CompanyQuery is the builder for querying Company entities.
type CompanyQuery struct {
	config
	ctx               *QueryContext
	order             []company.OrderOption
	inters            []Interceptor
	predicates        []predicate.Company
	withProfile       *CompanyUserQuery
	withVehicles      *VehicleQuery
	withRoutes        *RouteQuery
	withTrips         *TripQuery
	withBookings      *BookingQuery
	withNotifications *NotificationQuery
	modifiers         []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CompanyQuery builder.
func (cq *CompanyQuery) Where(ps ...predicate.Company) *CompanyQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CompanyQuery) Limit(limit int) *CompanyQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CompanyQuery) Offset(offset int) *CompanyQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CompanyQuery) Unique(unique bool) *CompanyQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CompanyQuery) Order(o ...company.OrderOption) *CompanyQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryProfile chains the current query on the "profile" edge.
func (cq *CompanyQuery) QueryProfile() *CompanyUserQuery {
	query := (&CompanyUserClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(company.Table, company.FieldID, selector),
			sqlgraph.To(companyuser.Table, companyuser.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, company.ProfileTable, company.ProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVehicles chains the current query on the "vehicles" edge.
func (cq *CompanyQuery) QueryVehicles() *VehicleQuery {
	query := (&VehicleClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(company.Table, company.FieldID, selector),
			sqlgraph.To(vehicle.Table, vehicle.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, company.VehiclesTable, company.VehiclesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoutes chains the current query on the "routes" edge.
func (cq *CompanyQuery) QueryRoutes() *RouteQuery {
	query := (&RouteClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(company.Table, company.FieldID, selector),
			sqlgraph.To(route.Table, route.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, company.RoutesTable, company.RoutesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTrips chains the current query on the "trips" edge.
func (cq *CompanyQuery) QueryTrips() *TripQuery {
	query := (&TripClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(company.Table, company.FieldID, selector),
			sqlgraph.To(trip.Table, trip.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, company.TripsTable, company.TripsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBookings chains the current query on the "bookings" edge.
func (cq *CompanyQuery) QueryBookings() *BookingQuery {
	query := (&BookingClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(company.Table, company.FieldID, selector),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, company.BookingsTable, company.BookingsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifications chains the current query on the "notifications" edge.
func (cq *CompanyQuery) QueryNotifications() *NotificationQuery {
	query := (&NotificationClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(company.Table, company.FieldID, selector),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, company.NotificationsTable, company.NotificationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Company entity from the query.
// Returns a *NotFoundError when no Company was found.
func (cq *CompanyQuery) First(ctx context.Context) (*Company, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{company.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CompanyQuery) FirstX(ctx context.Context) *Company {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Company ID from the query.
// Returns a *NotFoundError when no Company ID was found.
func (cq *CompanyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{company.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CompanyQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Company entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Company entity is found.
// Returns a *NotFoundError when no Company entities are found.
func (cq *CompanyQuery) Only(ctx context.Context) (*Company, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{company.Label}
	default:
		return nil, &NotSingularError{company.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CompanyQuery) OnlyX(ctx context.Context) *Company {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Company ID in the query.
// Returns a *NotSingularError when more than one Company ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CompanyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{company.Label}
	default:
		err = &NotSingularError{company.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CompanyQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Companies.
func (cq *CompanyQuery) All(ctx context.Context) ([]*Company, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Company, *CompanyQuery]()
	return withInterceptors[[]*Company](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CompanyQuery) AllX(ctx context.Context) []*Company {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Company IDs.
func (cq *CompanyQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(company.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CompanyQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CompanyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CompanyQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CompanyQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CompanyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CompanyQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CompanyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CompanyQuery) Clone() *CompanyQuery {
	if cq == nil {
		return nil
	}
	return &CompanyQuery{
		config:            cq.config,
		ctx:               cq.ctx.Clone(),
		order:             append([]company.OrderOption{}, cq.order...),
		inters:            append([]Interceptor{}, cq.inters...),
		predicates:        append([]predicate.Company{}, cq.predicates...),
		withProfile:       cq.withProfile.Clone(),
		withVehicles:      cq.withVehicles.Clone(),
		withRoutes:        cq.withRoutes.Clone(),
		withTrips:         cq.withTrips.Clone(),
		withBookings:      cq.withBookings.Clone(),
		withNotifications: cq.withNotifications.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithProfile tells the query-builder to eager-load the nodes that are connected to
// the "profile" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompanyQuery) WithProfile(opts ...func(*CompanyUserQuery)) *CompanyQuery {
	query := (&CompanyUserClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withProfile = query
	return cq
}

// WithVehicles tells the query-builder to eager-load the nodes that are connected to
// the "vehicles" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompanyQuery) WithVehicles(opts ...func(*VehicleQuery)) *CompanyQuery {
	query := (&VehicleClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withVehicles = query
	return cq
}

// WithRoutes tells the query-builder to eager-load the nodes that are connected to
// the "routes" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompanyQuery) WithRoutes(opts ...func(*RouteQuery)) *CompanyQuery {
	query := (&RouteClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withRoutes = query
	return cq
}

// WithTrips tells the query-builder to eager-load the nodes that are connected to
// the "trips" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompanyQuery) WithTrips(opts ...func(*TripQuery)) *CompanyQuery {
	query := (&TripClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withTrips = query
	return cq
}

// WithBookings tells the query-builder to eager-load the nodes that are connected to
// the "bookings" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompanyQuery) WithBookings(opts ...func(*BookingQuery)) *CompanyQuery {
	query := (&BookingClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withBookings = query
	return cq
}

// WithNotifications tells the query-builder to eager-load the nodes that are connected to
// the "notifications" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompanyQuery) WithNotifications(opts ...func(*NotificationQuery)) *CompanyQuery {
	query := (&NotificationClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withNotifications = query
	return cq
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
//	client.Company.Query().
//		GroupBy(company.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CompanyQuery) GroupBy(field string, fields ...string) *CompanyGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CompanyGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = company.Label
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
//	client.Company.Query().
//		Select(company.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *CompanyQuery) Select(fields ...string) *CompanySelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CompanySelect{CompanyQuery: cq}
	sbuild.label = company.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CompanySelect configured with the given aggregations.
func (cq *CompanyQuery) Aggregate(fns ...AggregateFunc) *CompanySelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CompanyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !company.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CompanyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Company, error) {
	var (
		nodes       = []*Company{}
		_spec       = cq.querySpec()
		loadedTypes = [6]bool{
			cq.withProfile != nil,
			cq.withVehicles != nil,
			cq.withRoutes != nil,
			cq.withTrips != nil,
			cq.withBookings != nil,
			cq.withNotifications != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Company).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Company{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withProfile; query != nil {
		if err := cq.loadProfile(ctx, query, nodes,
			func(n *Company) { n.Edges.Profile = []*CompanyUser{} },
			func(n *Company, e *CompanyUser) { n.Edges.Profile = append(n.Edges.Profile, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withVehicles; query != nil {
		if err := cq.loadVehicles(ctx, query, nodes,
			func(n *Company) { n.Edges.Vehicles = []*Vehicle{} },
			func(n *Company, e *Vehicle) { n.Edges.Vehicles = append(n.Edges.Vehicles, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withRoutes; query != nil {
		if err := cq.loadRoutes(ctx, query, nodes,
			func(n *Company) { n.Edges.Routes = []*Route{} },
			func(n *Company, e *Route) { n.Edges.Routes = append(n.Edges.Routes, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withTrips; query != nil {
		if err := cq.loadTrips(ctx, query, nodes,
			func(n *Company) { n.Edges.Trips = []*Trip{} },
			func(n *Company, e *Trip) { n.Edges.Trips = append(n.Edges.Trips, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withBookings; query != nil {
		if err := cq.loadBookings(ctx, query, nodes,
			func(n *Company) { n.Edges.Bookings = []*Booking{} },
			func(n *Company, e *Booking) { n.Edges.Bookings = append(n.Edges.Bookings, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withNotifications; query != nil {
		if err := cq.loadNotifications(ctx, query, nodes,
			func(n *Company) { n.Edges.Notifications = []*Notification{} },
			func(n *Company, e *Notification) { n.Edges.Notifications = append(n.Edges.Notifications, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CompanyQuery) loadProfile(ctx context.Context, query *CompanyUserQuery, nodes []*Company, init func(*Company), assign func(*Company, *CompanyUser)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Company)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CompanyUser(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(company.ProfileColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.company_profile
		if fk == nil {
			return fmt.Errorf(`foreign-key "company_profile" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "company_profile" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CompanyQuery) loadVehicles(ctx context.Context, query *VehicleQuery, nodes []*Company, init func(*Company), assign func(*Company, *Vehicle)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Company)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(company.VehiclesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.company_vehicles
		if fk == nil {
			return fmt.Errorf(`foreign-key "company_vehicles" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "company_vehicles" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CompanyQuery) loadRoutes(ctx context.Context, query *RouteQuery, nodes []*Company, init func(*Company), assign func(*Company, *Route)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Company)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Route(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(company.RoutesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.company_routes
		if fk == nil {
			return fmt.Errorf(`foreign-key "company_routes" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "company_routes" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CompanyQuery) loadTrips(ctx context.Context, query *TripQuery, nodes []*Company, init func(*Company), assign func(*Company, *Trip)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Company)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Trip(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(company.TripsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.company_trips
		if fk == nil {
			return fmt.Errorf(`foreign-key "company_trips" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "company_trips" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CompanyQuery) loadBookings(ctx context.Context, query *BookingQuery, nodes []*Company, init func(*Company), assign func(*Company, *Booking)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Company)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(company.BookingsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.company_bookings
		if fk == nil {
			return fmt.Errorf(`foreign-key "company_bookings" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "company_bookings" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CompanyQuery) loadNotifications(ctx context.Context, query *NotificationQuery, nodes []*Company, init func(*Company), assign func(*Company, *Notification)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Company)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(company.NotificationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.company_notifications
		if fk == nil {
			return fmt.Errorf(`foreign-key "company_notifications" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "company_notifications" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CompanyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CompanyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(company.Table, company.Columns, sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, company.FieldID)
		for i := range fields {
			if fields[i] != company.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CompanyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(company.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = company.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *CompanyQuery) Modify(modifiers ...func(s *sql.Selector)) *CompanySelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// CompanyGroupBy is the group-by builder for Company entities.
type CompanyGroupBy struct {
	selector
	build *CompanyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CompanyGroupBy) Aggregate(fns ...AggregateFunc) *CompanyGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CompanyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompanyQuery, *CompanyGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CompanyGroupBy) sqlScan(ctx context.Context, root *CompanyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CompanySelect is the builder for selecting fields of Company entities.
type CompanySelect struct {
	*CompanyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CompanySelect) Aggregate(fns ...AggregateFunc) *CompanySelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CompanySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompanyQuery, *CompanySelect](ctx, cs.CompanyQuery, cs, cs.inters, v)
}

func (cs *CompanySelect) sqlScan(ctx context.Context, root *CompanyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *CompanySelect) Modify(modifiers ...func(s *sql.Selector)) *CompanySelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}