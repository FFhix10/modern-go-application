// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo/todoadapter/ent/predicate"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo/todoadapter/ent/todoitem"
)

// TodoItemQuery is the builder for querying TodoItem entities.
type TodoItemQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.TodoItem
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (tiq *TodoItemQuery) Where(ps ...predicate.TodoItem) *TodoItemQuery {
	tiq.predicates = append(tiq.predicates, ps...)
	return tiq
}

// Limit adds a limit step to the query.
func (tiq *TodoItemQuery) Limit(limit int) *TodoItemQuery {
	tiq.limit = &limit
	return tiq
}

// Offset adds an offset step to the query.
func (tiq *TodoItemQuery) Offset(offset int) *TodoItemQuery {
	tiq.offset = &offset
	return tiq
}

// Order adds an order step to the query.
func (tiq *TodoItemQuery) Order(o ...Order) *TodoItemQuery {
	tiq.order = append(tiq.order, o...)
	return tiq
}

// First returns the first TodoItem entity in the query. Returns *NotFoundError when no todoitem was found.
func (tiq *TodoItemQuery) First(ctx context.Context) (*TodoItem, error) {
	tis, err := tiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(tis) == 0 {
		return nil, &NotFoundError{todoitem.Label}
	}
	return tis[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tiq *TodoItemQuery) FirstX(ctx context.Context) *TodoItem {
	ti, err := tiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ti
}

// FirstID returns the first TodoItem id in the query. Returns *NotFoundError when no id was found.
func (tiq *TodoItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{todoitem.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (tiq *TodoItemQuery) FirstXID(ctx context.Context) int {
	id, err := tiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only TodoItem entity in the query, returns an error if not exactly one entity was returned.
func (tiq *TodoItemQuery) Only(ctx context.Context) (*TodoItem, error) {
	tis, err := tiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(tis) {
	case 1:
		return tis[0], nil
	case 0:
		return nil, &NotFoundError{todoitem.Label}
	default:
		return nil, &NotSingularError{todoitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tiq *TodoItemQuery) OnlyX(ctx context.Context) *TodoItem {
	ti, err := tiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ti
}

// OnlyID returns the only TodoItem id in the query, returns an error if not exactly one id was returned.
func (tiq *TodoItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{todoitem.Label}
	default:
		err = &NotSingularError{todoitem.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (tiq *TodoItemQuery) OnlyXID(ctx context.Context) int {
	id, err := tiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TodoItems.
func (tiq *TodoItemQuery) All(ctx context.Context) ([]*TodoItem, error) {
	if err := tiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tiq *TodoItemQuery) AllX(ctx context.Context) []*TodoItem {
	tis, err := tiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return tis
}

// IDs executes the query and returns a list of TodoItem ids.
func (tiq *TodoItemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tiq.Select(todoitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tiq *TodoItemQuery) IDsX(ctx context.Context) []int {
	ids, err := tiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tiq *TodoItemQuery) Count(ctx context.Context) (int, error) {
	if err := tiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tiq *TodoItemQuery) CountX(ctx context.Context) int {
	count, err := tiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tiq *TodoItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := tiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tiq *TodoItemQuery) ExistX(ctx context.Context) bool {
	exist, err := tiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tiq *TodoItemQuery) Clone() *TodoItemQuery {
	return &TodoItemQuery{
		config:     tiq.config,
		limit:      tiq.limit,
		offset:     tiq.offset,
		order:      append([]Order{}, tiq.order...),
		unique:     append([]string{}, tiq.unique...),
		predicates: append([]predicate.TodoItem{}, tiq.predicates...),
		// clone intermediate query.
		sql:  tiq.sql.Clone(),
		path: tiq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UID string `json:"uid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TodoItem.Query().
//		GroupBy(todoitem.FieldUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tiq *TodoItemQuery) GroupBy(field string, fields ...string) *TodoItemGroupBy {
	group := &TodoItemGroupBy{config: tiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tiq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		UID string `json:"uid,omitempty"`
//	}
//
//	client.TodoItem.Query().
//		Select(todoitem.FieldUID).
//		Scan(ctx, &v)
//
func (tiq *TodoItemQuery) Select(field string, fields ...string) *TodoItemSelect {
	selector := &TodoItemSelect{config: tiq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tiq.sqlQuery(), nil
	}
	return selector
}

func (tiq *TodoItemQuery) prepareQuery(ctx context.Context) error {
	if tiq.path != nil {
		prev, err := tiq.path(ctx)
		if err != nil {
			return err
		}
		tiq.sql = prev
	}
	return nil
}

func (tiq *TodoItemQuery) sqlAll(ctx context.Context) ([]*TodoItem, error) {
	var (
		nodes = []*TodoItem{}
		_spec = tiq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &TodoItem{config: tiq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, tiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tiq *TodoItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tiq.querySpec()
	return sqlgraph.CountNodes(ctx, tiq.driver, _spec)
}

func (tiq *TodoItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (tiq *TodoItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   todoitem.Table,
			Columns: todoitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todoitem.FieldID,
			},
		},
		From:   tiq.sql,
		Unique: true,
	}
	if ps := tiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tiq *TodoItemQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(tiq.driver.Dialect())
	t1 := builder.Table(todoitem.Table)
	selector := builder.Select(t1.Columns(todoitem.Columns...)...).From(t1)
	if tiq.sql != nil {
		selector = tiq.sql
		selector.Select(selector.Columns(todoitem.Columns...)...)
	}
	for _, p := range tiq.predicates {
		p(selector)
	}
	for _, p := range tiq.order {
		p(selector)
	}
	if offset := tiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TodoItemGroupBy is the builder for group-by TodoItem entities.
type TodoItemGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tigb *TodoItemGroupBy) Aggregate(fns ...Aggregate) *TodoItemGroupBy {
	tigb.fns = append(tigb.fns, fns...)
	return tigb
}

// Scan applies the group-by query and scan the result into the given value.
func (tigb *TodoItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tigb.path(ctx)
	if err != nil {
		return err
	}
	tigb.sql = query
	return tigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tigb *TodoItemGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (tigb *TodoItemGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tigb.fields) > 1 {
		return nil, errors.New("ent: TodoItemGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tigb *TodoItemGroupBy) StringsX(ctx context.Context) []string {
	v, err := tigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (tigb *TodoItemGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tigb.fields) > 1 {
		return nil, errors.New("ent: TodoItemGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tigb *TodoItemGroupBy) IntsX(ctx context.Context) []int {
	v, err := tigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (tigb *TodoItemGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tigb.fields) > 1 {
		return nil, errors.New("ent: TodoItemGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tigb *TodoItemGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (tigb *TodoItemGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tigb.fields) > 1 {
		return nil, errors.New("ent: TodoItemGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tigb *TodoItemGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tigb *TodoItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tigb.sqlQuery().Query()
	if err := tigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tigb *TodoItemGroupBy) sqlQuery() *sql.Selector {
	selector := tigb.sql
	columns := make([]string, 0, len(tigb.fields)+len(tigb.fns))
	columns = append(columns, tigb.fields...)
	for _, fn := range tigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(tigb.fields...)
}

// TodoItemSelect is the builder for select fields of TodoItem entities.
type TodoItemSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (tis *TodoItemSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := tis.path(ctx)
	if err != nil {
		return err
	}
	tis.sql = query
	return tis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tis *TodoItemSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (tis *TodoItemSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tis.fields) > 1 {
		return nil, errors.New("ent: TodoItemSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tis *TodoItemSelect) StringsX(ctx context.Context) []string {
	v, err := tis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (tis *TodoItemSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tis.fields) > 1 {
		return nil, errors.New("ent: TodoItemSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tis *TodoItemSelect) IntsX(ctx context.Context) []int {
	v, err := tis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (tis *TodoItemSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tis.fields) > 1 {
		return nil, errors.New("ent: TodoItemSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tis *TodoItemSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (tis *TodoItemSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tis.fields) > 1 {
		return nil, errors.New("ent: TodoItemSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tis *TodoItemSelect) BoolsX(ctx context.Context) []bool {
	v, err := tis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tis *TodoItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tis.sqlQuery().Query()
	if err := tis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tis *TodoItemSelect) sqlQuery() sql.Querier {
	selector := tis.sql
	selector.Select(selector.Columns(tis.fields...)...)
	return selector
}
