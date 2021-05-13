// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/predicate"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/typeconfig"
)

// TypeConfigQuery is the builder for querying TypeConfig entities.
type TypeConfigQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TypeConfig
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TypeConfigQuery builder.
func (tcq *TypeConfigQuery) Where(ps ...predicate.TypeConfig) *TypeConfigQuery {
	tcq.predicates = append(tcq.predicates, ps...)
	return tcq
}

// Limit adds a limit step to the query.
func (tcq *TypeConfigQuery) Limit(limit int) *TypeConfigQuery {
	tcq.limit = &limit
	return tcq
}

// Offset adds an offset step to the query.
func (tcq *TypeConfigQuery) Offset(offset int) *TypeConfigQuery {
	tcq.offset = &offset
	return tcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tcq *TypeConfigQuery) Unique(unique bool) *TypeConfigQuery {
	tcq.unique = &unique
	return tcq
}

// Order adds an order step to the query.
func (tcq *TypeConfigQuery) Order(o ...OrderFunc) *TypeConfigQuery {
	tcq.order = append(tcq.order, o...)
	return tcq
}

// First returns the first TypeConfig entity from the query.
// Returns a *NotFoundError when no TypeConfig was found.
func (tcq *TypeConfigQuery) First(ctx context.Context) (*TypeConfig, error) {
	nodes, err := tcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{typeconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tcq *TypeConfigQuery) FirstX(ctx context.Context) *TypeConfig {
	node, err := tcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TypeConfig ID from the query.
// Returns a *NotFoundError when no TypeConfig ID was found.
func (tcq *TypeConfigQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{typeconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tcq *TypeConfigQuery) FirstIDX(ctx context.Context) int {
	id, err := tcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TypeConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TypeConfig entity is not found.
// Returns a *NotFoundError when no TypeConfig entities are found.
func (tcq *TypeConfigQuery) Only(ctx context.Context) (*TypeConfig, error) {
	nodes, err := tcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{typeconfig.Label}
	default:
		return nil, &NotSingularError{typeconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tcq *TypeConfigQuery) OnlyX(ctx context.Context) *TypeConfig {
	node, err := tcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TypeConfig ID in the query.
// Returns a *NotSingularError when exactly one TypeConfig ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tcq *TypeConfigQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{typeconfig.Label}
	default:
		err = &NotSingularError{typeconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tcq *TypeConfigQuery) OnlyIDX(ctx context.Context) int {
	id, err := tcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TypeConfigs.
func (tcq *TypeConfigQuery) All(ctx context.Context) ([]*TypeConfig, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tcq *TypeConfigQuery) AllX(ctx context.Context) []*TypeConfig {
	nodes, err := tcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TypeConfig IDs.
func (tcq *TypeConfigQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tcq.Select(typeconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tcq *TypeConfigQuery) IDsX(ctx context.Context) []int {
	ids, err := tcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tcq *TypeConfigQuery) Count(ctx context.Context) (int, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tcq *TypeConfigQuery) CountX(ctx context.Context) int {
	count, err := tcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tcq *TypeConfigQuery) Exist(ctx context.Context) (bool, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tcq *TypeConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := tcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TypeConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tcq *TypeConfigQuery) Clone() *TypeConfigQuery {
	if tcq == nil {
		return nil
	}
	return &TypeConfigQuery{
		config:     tcq.config,
		limit:      tcq.limit,
		offset:     tcq.offset,
		order:      append([]OrderFunc{}, tcq.order...),
		predicates: append([]predicate.TypeConfig{}, tcq.predicates...),
		// clone intermediate query.
		sql:  tcq.sql.Clone(),
		path: tcq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type string `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TypeConfig.Query().
//		GroupBy(typeconfig.FieldType).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (tcq *TypeConfigQuery) GroupBy(field string, fields ...string) *TypeConfigGroupBy {
	group := &TypeConfigGroupBy{config: tcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tcq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type string `json:"type,omitempty"`
//	}
//
//	client.TypeConfig.Query().
//		Select(typeconfig.FieldType).
//		Scan(ctx, &v)
//
func (tcq *TypeConfigQuery) Select(field string, fields ...string) *TypeConfigSelect {
	tcq.fields = append([]string{field}, fields...)
	return &TypeConfigSelect{TypeConfigQuery: tcq}
}

func (tcq *TypeConfigQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tcq.fields {
		if !typeconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if tcq.path != nil {
		prev, err := tcq.path(ctx)
		if err != nil {
			return err
		}
		tcq.sql = prev
	}
	return nil
}

func (tcq *TypeConfigQuery) sqlAll(ctx context.Context) ([]*TypeConfig, error) {
	var (
		nodes = []*TypeConfig{}
		_spec = tcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TypeConfig{config: tcq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("model: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, tcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tcq *TypeConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tcq.querySpec()
	return sqlgraph.CountNodes(ctx, tcq.driver, _spec)
}

func (tcq *TypeConfigQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (tcq *TypeConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   typeconfig.Table,
			Columns: typeconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typeconfig.FieldID,
			},
		},
		From:   tcq.sql,
		Unique: true,
	}
	if unique := tcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, typeconfig.FieldID)
		for i := range fields {
			if fields[i] != typeconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tcq *TypeConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tcq.driver.Dialect())
	t1 := builder.Table(typeconfig.Table)
	selector := builder.Select(t1.Columns(typeconfig.Columns...)...).From(t1)
	if tcq.sql != nil {
		selector = tcq.sql
		selector.Select(selector.Columns(typeconfig.Columns...)...)
	}
	for _, p := range tcq.predicates {
		p(selector)
	}
	for _, p := range tcq.order {
		p(selector)
	}
	if offset := tcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TypeConfigGroupBy is the group-by builder for TypeConfig entities.
type TypeConfigGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tcgb *TypeConfigGroupBy) Aggregate(fns ...AggregateFunc) *TypeConfigGroupBy {
	tcgb.fns = append(tcgb.fns, fns...)
	return tcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tcgb *TypeConfigGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tcgb.path(ctx)
	if err != nil {
		return err
	}
	tcgb.sql = query
	return tcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tcgb *TypeConfigGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TypeConfigGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("model: TypeConfigGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tcgb *TypeConfigGroupBy) StringsX(ctx context.Context) []string {
	v, err := tcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TypeConfigGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{typeconfig.Label}
	default:
		err = fmt.Errorf("model: TypeConfigGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tcgb *TypeConfigGroupBy) StringX(ctx context.Context) string {
	v, err := tcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TypeConfigGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("model: TypeConfigGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tcgb *TypeConfigGroupBy) IntsX(ctx context.Context) []int {
	v, err := tcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TypeConfigGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{typeconfig.Label}
	default:
		err = fmt.Errorf("model: TypeConfigGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tcgb *TypeConfigGroupBy) IntX(ctx context.Context) int {
	v, err := tcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TypeConfigGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("model: TypeConfigGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tcgb *TypeConfigGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TypeConfigGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{typeconfig.Label}
	default:
		err = fmt.Errorf("model: TypeConfigGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tcgb *TypeConfigGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TypeConfigGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tcgb.fields) > 1 {
		return nil, errors.New("model: TypeConfigGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tcgb *TypeConfigGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tcgb *TypeConfigGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{typeconfig.Label}
	default:
		err = fmt.Errorf("model: TypeConfigGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tcgb *TypeConfigGroupBy) BoolX(ctx context.Context) bool {
	v, err := tcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tcgb *TypeConfigGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tcgb.fields {
		if !typeconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tcgb *TypeConfigGroupBy) sqlQuery() *sql.Selector {
	selector := tcgb.sql
	columns := make([]string, 0, len(tcgb.fields)+len(tcgb.fns))
	columns = append(columns, tcgb.fields...)
	for _, fn := range tcgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(tcgb.fields...)
}

// TypeConfigSelect is the builder for selecting fields of TypeConfig entities.
type TypeConfigSelect struct {
	*TypeConfigQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tcs *TypeConfigSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tcs.prepareQuery(ctx); err != nil {
		return err
	}
	tcs.sql = tcs.TypeConfigQuery.sqlQuery(ctx)
	return tcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tcs *TypeConfigSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tcs *TypeConfigSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("model: TypeConfigSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tcs *TypeConfigSelect) StringsX(ctx context.Context) []string {
	v, err := tcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tcs *TypeConfigSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{typeconfig.Label}
	default:
		err = fmt.Errorf("model: TypeConfigSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tcs *TypeConfigSelect) StringX(ctx context.Context) string {
	v, err := tcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tcs *TypeConfigSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("model: TypeConfigSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tcs *TypeConfigSelect) IntsX(ctx context.Context) []int {
	v, err := tcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tcs *TypeConfigSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{typeconfig.Label}
	default:
		err = fmt.Errorf("model: TypeConfigSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tcs *TypeConfigSelect) IntX(ctx context.Context) int {
	v, err := tcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tcs *TypeConfigSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("model: TypeConfigSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tcs *TypeConfigSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tcs *TypeConfigSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{typeconfig.Label}
	default:
		err = fmt.Errorf("model: TypeConfigSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tcs *TypeConfigSelect) Float64X(ctx context.Context) float64 {
	v, err := tcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tcs *TypeConfigSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tcs.fields) > 1 {
		return nil, errors.New("model: TypeConfigSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tcs *TypeConfigSelect) BoolsX(ctx context.Context) []bool {
	v, err := tcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tcs *TypeConfigSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{typeconfig.Label}
	default:
		err = fmt.Errorf("model: TypeConfigSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tcs *TypeConfigSelect) BoolX(ctx context.Context) bool {
	v, err := tcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tcs *TypeConfigSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tcs.sqlQuery().Query()
	if err := tcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tcs *TypeConfigSelect) sqlQuery() sql.Querier {
	selector := tcs.sql
	selector.Select(selector.Columns(tcs.fields...)...)
	return selector
}
