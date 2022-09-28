package dent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

type WithQuery struct {
	fromkey string
	single  bool
	query   *DynamicQuery
}

// DynamicQuery is the builder for querying Dynamic entities.
type DynamicQuery struct {
	limit  *int
	offset *int
	unique *bool
	order  []OrderFunc
	fields []string

	predicates []PredicateDynamic
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
	// eager-loading edges.
	withData map[string]*WithQuery
	// build
	table *Table
}

// Where adds a new predicate for the DynamicQuery builder.
func (dq *DynamicQuery) Where(ps ...PredicateDynamic) *DynamicQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DynamicQuery) Limit(limit int) *DynamicQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DynamicQuery) Offset(offset int) *DynamicQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DynamicQuery) Unique(unique bool) *DynamicQuery {
	dq.unique = &unique
	return dq
}

// Order adds an order step to the query.
func (dq *DynamicQuery) Order(o ...OrderFunc) *DynamicQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// First returns the first Dynamic entity from the query.
// Returns a *NotFoundError when no Dynamic was found.
func (dq *DynamicQuery) First(ctx context.Context) (*Dynamic, error) {
	nodes, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dq.table.Name}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DynamicQuery) FirstX(ctx context.Context) *Dynamic {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Dynamic ID from the query.
// Returns a *NotFoundError when no Dynamic ID was found.
func (dq *DynamicQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dq.table.Name}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DynamicQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Dynamic entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Dynamic entity is found.
// Returns a *NotFoundError when no Dynamic entities are found.
func (dq *DynamicQuery) Only(ctx context.Context) (*Dynamic, error) {
	nodes, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dq.table.Name}
	default:
		return nil, &NotSingularError{dq.table.Name}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DynamicQuery) OnlyX(ctx context.Context) *Dynamic {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Dynamic ID in the query.
// Returns a *NotSingularError when more than one Dynamic ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DynamicQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dq.table.Name}
	default:
		err = &NotSingularError{dq.table.Name}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DynamicQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Dynamics.
func (dq *DynamicQuery) All(ctx context.Context) ([]*Dynamic, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DynamicQuery) AllX(ctx context.Context) []*Dynamic {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Dynamic IDs.
func (dq *DynamicQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dq.Select(FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DynamicQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DynamicQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DynamicQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DynamicQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DynamicQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DynamicQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DynamicQuery) Clone() *DynamicQuery {
	if dq == nil {
		return nil
	}

	withData := make(map[string]*WithQuery)
	for k, dq2 := range dq.withData {
		withData[k] = &WithQuery{
			fromkey: dq2.fromkey,
			query:   dq2.query.Clone(),
		}
	}

	return &DynamicQuery{
		limit:      dq.limit,
		offset:     dq.offset,
		order:      append([]OrderFunc{}, dq.order...),
		predicates: append([]PredicateDynamic{}, dq.predicates...),
		// clone intermediate query.
		sql:      dq.sql.Clone(),
		withData: withData,
		path:     dq.path,
		unique:   dq.unique,
		table:    dq.table,
	}
}

// WithData 关联查询某个边的值，一对一关系
// 通过当前数据的fromKey字段对应的ID的去查询对应的关联表中的值
// 如查询用户的创建者
func (dq *DynamicQuery) WithData(table, storeKey, fromKey string, opts ...func(*DynamicQuery)) *DynamicQuery {
	query := &DynamicQuery{table: dq.table.client.Table(table), withData: make(map[string]*WithQuery)}
	for _, opt := range opts {
		opt(query)
	}
	dq.withData[storeKey] = &WithQuery{
		fromkey: fromKey,
		single:  true,
		query:   query,
	}
	return dq
}

// WithListData 关联查询列表
// 通过当前数据的ID去查询fromKey字段对应的ID的所有列表值
// 如查询用户的角色列表
func (dq *DynamicQuery) WithListData(table, storeKey, fromKey string, opts ...func(*DynamicQuery)) *DynamicQuery {
	query := &DynamicQuery{table: dq.table.client.Table(table), withData: make(map[string]*WithQuery)}
	for _, opt := range opts {
		opt(query)
	}
	dq.withData[storeKey] = &WithQuery{
		fromkey: fromKey,
		single:  false,
		query:   query,
	}
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatorID int `json:"creator_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Dynamic.Query().
//		GroupBy(dynamic.FieldCreatorID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DynamicQuery) GroupBy(field string, fields ...string) *DynamicGroupBy {
	grbuild := &DynamicGroupBy{table: dq.table}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(ctx), nil
	}
	grbuild.label = dq.table.Name
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatorID int `json:"creator_id,omitempty"`
//	}
//
//	client.Dynamic.Query().
//		Select(dynamic.FieldCreatorID).
//		Scan(ctx, &v)
func (dq *DynamicQuery) Select(fields ...string) *DynamicSelect {
	dq.fields = append(dq.fields, fields...)
	selbuild := &DynamicSelect{DynamicQuery: dq}
	selbuild.label = dq.table.Name
	selbuild.flds, selbuild.scan = &dq.fields, selbuild.Scan
	return selbuild
}

func (dq *DynamicQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dq.fields {
		if !dq.table.HasColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DynamicQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Dynamic, error) {
	var (
		nodes = []*Dynamic{}
		_spec = dq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return dq.table.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Dynamic{table: dq.table, Row: make(map[string]ent.Value)}
		node.Edges.SingleMap = make(map[string]*Dynamic)
		node.Edges.ListMap = make(map[string][]*Dynamic)
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.table.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	for key, dq2 := range dq.withData {
		if dq2.single {
			ids := make([]int, 0, len(nodes))
			nodeids := make(map[int][]*Dynamic)
			for i := range nodes {
				fk := int(nodes[i].Row[dq2.fromkey].(int64))
				if _, ok := nodeids[fk]; !ok {
					ids = append(ids, fk)
				}
				nodeids[fk] = append(nodeids[fk], nodes[i])
			}
			query := dq2.query
			query.Where(IDIn(ids...))
			neighbors, err := query.All(ctx)
			if err != nil {
				return nil, err
			}
			for _, n := range neighbors {
				nodes, ok := nodeids[n.ID]
				if !ok {
					return nil, fmt.Errorf(`unexpected foreign-key "editor_id" returned %v`, n.ID)
				}
				for i := range nodes {
					nodes[i].Edges.SingleMap[key] = n
				}
			}
		} else {
			fks := make([]driver.Value, 0, len(nodes))
			nodeids := make(map[int]*Dynamic)
			for i := range nodes {
				fks = append(fks, nodes[i].ID)
				nodeids[nodes[i].ID] = nodes[i]
			}
			query := dq2.query
			query.Where(PredicateDynamic(func(s *sql.Selector) {
				s.Where(sql.InValues(dq2.fromkey, fks...))
			}))
			neighbors, err := query.All(ctx)
			if err != nil {
				return nil, err
			}
			for _, n := range neighbors {
				fk := int(n.Row[dq2.fromkey].(int64))
				node, ok := nodeids[fk]
				if !ok {
					return nil, fmt.Errorf(`unexpected foreign-key "d_id" returned %v for node %v`, fk, n.ID)
				}
				node.Edges.ListMap[key] = append(node.Edges.ListMap[key], n)
			}
		}
	}

	return nodes, nil
}

func (dq *DynamicQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.table.driver, _spec)
}

func (dq *DynamicQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dq *DynamicQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dq.table.Name,
			Columns: dq.table.GetColumns(),
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if unique := dq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, FieldID)
		for i := range fields {
			if fields[i] != FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DynamicQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.table.driver.Dialect())
	t1 := builder.Table(dq.table.Name)
	columns := dq.fields
	if len(columns) == 0 {
		columns = dq.table.GetColumns()
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.unique != nil && *dq.unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DynamicGroupBy is the group-by builder for Dynamic entities.
type DynamicGroupBy struct {
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql   *sql.Selector
	path  func(context.Context) (*sql.Selector, error)
	table *Table
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DynamicGroupBy) Aggregate(fns ...AggregateFunc) *DynamicGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dgb *DynamicGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

func (dgb *DynamicGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dgb.fields {
		if !dgb.table.HasColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.table.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DynamicGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql.Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
		for _, f := range dgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dgb.fields...)...)
}

// DynamicSelect is the builder for selecting fields of Dynamic entities.
type DynamicSelect struct {
	*DynamicQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DynamicSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	ds.sql = ds.DynamicQuery.sqlQuery(ctx)
	return ds.sqlScan(ctx, v)
}

func (ds *DynamicSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sql.Query()
	if err := ds.table.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
