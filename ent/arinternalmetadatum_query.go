// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"story.com/story/app/ent/arinternalmetadatum"
	"story.com/story/app/ent/predicate"
)

// ArInternalMetadatumQuery is the builder for querying ArInternalMetadatum entities.
type ArInternalMetadatumQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.ArInternalMetadatum
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArInternalMetadatumQuery builder.
func (aimq *ArInternalMetadatumQuery) Where(ps ...predicate.ArInternalMetadatum) *ArInternalMetadatumQuery {
	aimq.predicates = append(aimq.predicates, ps...)
	return aimq
}

// Limit the number of records to be returned by this query.
func (aimq *ArInternalMetadatumQuery) Limit(limit int) *ArInternalMetadatumQuery {
	aimq.ctx.Limit = &limit
	return aimq
}

// Offset to start from.
func (aimq *ArInternalMetadatumQuery) Offset(offset int) *ArInternalMetadatumQuery {
	aimq.ctx.Offset = &offset
	return aimq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aimq *ArInternalMetadatumQuery) Unique(unique bool) *ArInternalMetadatumQuery {
	aimq.ctx.Unique = &unique
	return aimq
}

// Order specifies how the records should be ordered.
func (aimq *ArInternalMetadatumQuery) Order(o ...OrderFunc) *ArInternalMetadatumQuery {
	aimq.order = append(aimq.order, o...)
	return aimq
}

// First returns the first ArInternalMetadatum entity from the query.
// Returns a *NotFoundError when no ArInternalMetadatum was found.
func (aimq *ArInternalMetadatumQuery) First(ctx context.Context) (*ArInternalMetadatum, error) {
	nodes, err := aimq.Limit(1).All(setContextOp(ctx, aimq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{arinternalmetadatum.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aimq *ArInternalMetadatumQuery) FirstX(ctx context.Context) *ArInternalMetadatum {
	node, err := aimq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ArInternalMetadatum ID from the query.
// Returns a *NotFoundError when no ArInternalMetadatum ID was found.
func (aimq *ArInternalMetadatumQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aimq.Limit(1).IDs(setContextOp(ctx, aimq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{arinternalmetadatum.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aimq *ArInternalMetadatumQuery) FirstIDX(ctx context.Context) string {
	id, err := aimq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ArInternalMetadatum entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ArInternalMetadatum entity is found.
// Returns a *NotFoundError when no ArInternalMetadatum entities are found.
func (aimq *ArInternalMetadatumQuery) Only(ctx context.Context) (*ArInternalMetadatum, error) {
	nodes, err := aimq.Limit(2).All(setContextOp(ctx, aimq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{arinternalmetadatum.Label}
	default:
		return nil, &NotSingularError{arinternalmetadatum.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aimq *ArInternalMetadatumQuery) OnlyX(ctx context.Context) *ArInternalMetadatum {
	node, err := aimq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ArInternalMetadatum ID in the query.
// Returns a *NotSingularError when more than one ArInternalMetadatum ID is found.
// Returns a *NotFoundError when no entities are found.
func (aimq *ArInternalMetadatumQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aimq.Limit(2).IDs(setContextOp(ctx, aimq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{arinternalmetadatum.Label}
	default:
		err = &NotSingularError{arinternalmetadatum.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aimq *ArInternalMetadatumQuery) OnlyIDX(ctx context.Context) string {
	id, err := aimq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ArInternalMetadata.
func (aimq *ArInternalMetadatumQuery) All(ctx context.Context) ([]*ArInternalMetadatum, error) {
	ctx = setContextOp(ctx, aimq.ctx, "All")
	if err := aimq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ArInternalMetadatum, *ArInternalMetadatumQuery]()
	return withInterceptors[[]*ArInternalMetadatum](ctx, aimq, qr, aimq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aimq *ArInternalMetadatumQuery) AllX(ctx context.Context) []*ArInternalMetadatum {
	nodes, err := aimq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ArInternalMetadatum IDs.
func (aimq *ArInternalMetadatumQuery) IDs(ctx context.Context) (ids []string, err error) {
	if aimq.ctx.Unique == nil && aimq.path != nil {
		aimq.Unique(true)
	}
	ctx = setContextOp(ctx, aimq.ctx, "IDs")
	if err = aimq.Select(arinternalmetadatum.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aimq *ArInternalMetadatumQuery) IDsX(ctx context.Context) []string {
	ids, err := aimq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aimq *ArInternalMetadatumQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aimq.ctx, "Count")
	if err := aimq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aimq, querierCount[*ArInternalMetadatumQuery](), aimq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aimq *ArInternalMetadatumQuery) CountX(ctx context.Context) int {
	count, err := aimq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aimq *ArInternalMetadatumQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aimq.ctx, "Exist")
	switch _, err := aimq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aimq *ArInternalMetadatumQuery) ExistX(ctx context.Context) bool {
	exist, err := aimq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArInternalMetadatumQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aimq *ArInternalMetadatumQuery) Clone() *ArInternalMetadatumQuery {
	if aimq == nil {
		return nil
	}
	return &ArInternalMetadatumQuery{
		config:     aimq.config,
		ctx:        aimq.ctx.Clone(),
		order:      append([]OrderFunc{}, aimq.order...),
		inters:     append([]Interceptor{}, aimq.inters...),
		predicates: append([]predicate.ArInternalMetadatum{}, aimq.predicates...),
		// clone intermediate query.
		sql:  aimq.sql.Clone(),
		path: aimq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ArInternalMetadatum.Query().
//		GroupBy(arinternalmetadatum.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aimq *ArInternalMetadatumQuery) GroupBy(field string, fields ...string) *ArInternalMetadatumGroupBy {
	aimq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ArInternalMetadatumGroupBy{build: aimq}
	grbuild.flds = &aimq.ctx.Fields
	grbuild.label = arinternalmetadatum.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//	}
//
//	client.ArInternalMetadatum.Query().
//		Select(arinternalmetadatum.FieldValue).
//		Scan(ctx, &v)
func (aimq *ArInternalMetadatumQuery) Select(fields ...string) *ArInternalMetadatumSelect {
	aimq.ctx.Fields = append(aimq.ctx.Fields, fields...)
	sbuild := &ArInternalMetadatumSelect{ArInternalMetadatumQuery: aimq}
	sbuild.label = arinternalmetadatum.Label
	sbuild.flds, sbuild.scan = &aimq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ArInternalMetadatumSelect configured with the given aggregations.
func (aimq *ArInternalMetadatumQuery) Aggregate(fns ...AggregateFunc) *ArInternalMetadatumSelect {
	return aimq.Select().Aggregate(fns...)
}

func (aimq *ArInternalMetadatumQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aimq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aimq); err != nil {
				return err
			}
		}
	}
	for _, f := range aimq.ctx.Fields {
		if !arinternalmetadatum.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aimq.path != nil {
		prev, err := aimq.path(ctx)
		if err != nil {
			return err
		}
		aimq.sql = prev
	}
	return nil
}

func (aimq *ArInternalMetadatumQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ArInternalMetadatum, error) {
	var (
		nodes = []*ArInternalMetadatum{}
		_spec = aimq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ArInternalMetadatum).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ArInternalMetadatum{config: aimq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aimq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aimq *ArInternalMetadatumQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aimq.querySpec()
	_spec.Node.Columns = aimq.ctx.Fields
	if len(aimq.ctx.Fields) > 0 {
		_spec.Unique = aimq.ctx.Unique != nil && *aimq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aimq.driver, _spec)
}

func (aimq *ArInternalMetadatumQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(arinternalmetadatum.Table, arinternalmetadatum.Columns, sqlgraph.NewFieldSpec(arinternalmetadatum.FieldID, field.TypeString))
	_spec.From = aimq.sql
	if unique := aimq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aimq.path != nil {
		_spec.Unique = true
	}
	if fields := aimq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, arinternalmetadatum.FieldID)
		for i := range fields {
			if fields[i] != arinternalmetadatum.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aimq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aimq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aimq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aimq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aimq *ArInternalMetadatumQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aimq.driver.Dialect())
	t1 := builder.Table(arinternalmetadatum.Table)
	columns := aimq.ctx.Fields
	if len(columns) == 0 {
		columns = arinternalmetadatum.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aimq.sql != nil {
		selector = aimq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aimq.ctx.Unique != nil && *aimq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aimq.predicates {
		p(selector)
	}
	for _, p := range aimq.order {
		p(selector)
	}
	if offset := aimq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aimq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ArInternalMetadatumGroupBy is the group-by builder for ArInternalMetadatum entities.
type ArInternalMetadatumGroupBy struct {
	selector
	build *ArInternalMetadatumQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aimgb *ArInternalMetadatumGroupBy) Aggregate(fns ...AggregateFunc) *ArInternalMetadatumGroupBy {
	aimgb.fns = append(aimgb.fns, fns...)
	return aimgb
}

// Scan applies the selector query and scans the result into the given value.
func (aimgb *ArInternalMetadatumGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aimgb.build.ctx, "GroupBy")
	if err := aimgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArInternalMetadatumQuery, *ArInternalMetadatumGroupBy](ctx, aimgb.build, aimgb, aimgb.build.inters, v)
}

func (aimgb *ArInternalMetadatumGroupBy) sqlScan(ctx context.Context, root *ArInternalMetadatumQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aimgb.fns))
	for _, fn := range aimgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aimgb.flds)+len(aimgb.fns))
		for _, f := range *aimgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aimgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aimgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ArInternalMetadatumSelect is the builder for selecting fields of ArInternalMetadatum entities.
type ArInternalMetadatumSelect struct {
	*ArInternalMetadatumQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aims *ArInternalMetadatumSelect) Aggregate(fns ...AggregateFunc) *ArInternalMetadatumSelect {
	aims.fns = append(aims.fns, fns...)
	return aims
}

// Scan applies the selector query and scans the result into the given value.
func (aims *ArInternalMetadatumSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aims.ctx, "Select")
	if err := aims.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArInternalMetadatumQuery, *ArInternalMetadatumSelect](ctx, aims.ArInternalMetadatumQuery, aims, aims.inters, v)
}

func (aims *ArInternalMetadatumSelect) sqlScan(ctx context.Context, root *ArInternalMetadatumQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aims.fns))
	for _, fn := range aims.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aims.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aims.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
