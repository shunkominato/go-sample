// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-gql-sample/app/ent/predicate"
	"go-gql-sample/app/ent/todo"
	"go-gql-sample/app/ent/todostatus"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoStatusUpdate is the builder for updating TodoStatus entities.
type TodoStatusUpdate struct {
	config
	hooks    []Hook
	mutation *TodoStatusMutation
}

// Where appends a list predicates to the TodoStatusUpdate builder.
func (tsu *TodoStatusUpdate) Where(ps ...predicate.TodoStatus) *TodoStatusUpdate {
	tsu.mutation.Where(ps...)
	return tsu
}

// SetStatus sets the "status" field.
func (tsu *TodoStatusUpdate) SetStatus(s string) *TodoStatusUpdate {
	tsu.mutation.SetStatus(s)
	return tsu
}

// SetCreatedAt sets the "created_at" field.
func (tsu *TodoStatusUpdate) SetCreatedAt(t time.Time) *TodoStatusUpdate {
	tsu.mutation.SetCreatedAt(t)
	return tsu
}

// SetUpdatedAt sets the "updated_at" field.
func (tsu *TodoStatusUpdate) SetUpdatedAt(t time.Time) *TodoStatusUpdate {
	tsu.mutation.SetUpdatedAt(t)
	return tsu
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (tsu *TodoStatusUpdate) AddTodoIDs(ids ...int) *TodoStatusUpdate {
	tsu.mutation.AddTodoIDs(ids...)
	return tsu
}

// AddTodos adds the "todos" edges to the Todo entity.
func (tsu *TodoStatusUpdate) AddTodos(t ...*Todo) *TodoStatusUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tsu.AddTodoIDs(ids...)
}

// Mutation returns the TodoStatusMutation object of the builder.
func (tsu *TodoStatusUpdate) Mutation() *TodoStatusMutation {
	return tsu.mutation
}

// ClearTodos clears all "todos" edges to the Todo entity.
func (tsu *TodoStatusUpdate) ClearTodos() *TodoStatusUpdate {
	tsu.mutation.ClearTodos()
	return tsu
}

// RemoveTodoIDs removes the "todos" edge to Todo entities by IDs.
func (tsu *TodoStatusUpdate) RemoveTodoIDs(ids ...int) *TodoStatusUpdate {
	tsu.mutation.RemoveTodoIDs(ids...)
	return tsu
}

// RemoveTodos removes "todos" edges to Todo entities.
func (tsu *TodoStatusUpdate) RemoveTodos(t ...*Todo) *TodoStatusUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tsu.RemoveTodoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tsu *TodoStatusUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TodoStatusMutation](ctx, tsu.sqlSave, tsu.mutation, tsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsu *TodoStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := tsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tsu *TodoStatusUpdate) Exec(ctx context.Context) error {
	_, err := tsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsu *TodoStatusUpdate) ExecX(ctx context.Context) {
	if err := tsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tsu *TodoStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(todostatus.Table, todostatus.Columns, sqlgraph.NewFieldSpec(todostatus.FieldID, field.TypeInt))
	if ps := tsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsu.mutation.Status(); ok {
		_spec.SetField(todostatus.FieldStatus, field.TypeString, value)
	}
	if value, ok := tsu.mutation.CreatedAt(); ok {
		_spec.SetField(todostatus.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tsu.mutation.UpdatedAt(); ok {
		_spec.SetField(todostatus.FieldUpdatedAt, field.TypeTime, value)
	}
	if tsu.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todostatus.TodosTable,
			Columns: []string{todostatus.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tsu.mutation.RemovedTodosIDs(); len(nodes) > 0 && !tsu.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todostatus.TodosTable,
			Columns: []string{todostatus.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tsu.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todostatus.TodosTable,
			Columns: []string{todostatus.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todostatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tsu.mutation.done = true
	return n, nil
}

// TodoStatusUpdateOne is the builder for updating a single TodoStatus entity.
type TodoStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoStatusMutation
}

// SetStatus sets the "status" field.
func (tsuo *TodoStatusUpdateOne) SetStatus(s string) *TodoStatusUpdateOne {
	tsuo.mutation.SetStatus(s)
	return tsuo
}

// SetCreatedAt sets the "created_at" field.
func (tsuo *TodoStatusUpdateOne) SetCreatedAt(t time.Time) *TodoStatusUpdateOne {
	tsuo.mutation.SetCreatedAt(t)
	return tsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tsuo *TodoStatusUpdateOne) SetUpdatedAt(t time.Time) *TodoStatusUpdateOne {
	tsuo.mutation.SetUpdatedAt(t)
	return tsuo
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (tsuo *TodoStatusUpdateOne) AddTodoIDs(ids ...int) *TodoStatusUpdateOne {
	tsuo.mutation.AddTodoIDs(ids...)
	return tsuo
}

// AddTodos adds the "todos" edges to the Todo entity.
func (tsuo *TodoStatusUpdateOne) AddTodos(t ...*Todo) *TodoStatusUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tsuo.AddTodoIDs(ids...)
}

// Mutation returns the TodoStatusMutation object of the builder.
func (tsuo *TodoStatusUpdateOne) Mutation() *TodoStatusMutation {
	return tsuo.mutation
}

// ClearTodos clears all "todos" edges to the Todo entity.
func (tsuo *TodoStatusUpdateOne) ClearTodos() *TodoStatusUpdateOne {
	tsuo.mutation.ClearTodos()
	return tsuo
}

// RemoveTodoIDs removes the "todos" edge to Todo entities by IDs.
func (tsuo *TodoStatusUpdateOne) RemoveTodoIDs(ids ...int) *TodoStatusUpdateOne {
	tsuo.mutation.RemoveTodoIDs(ids...)
	return tsuo
}

// RemoveTodos removes "todos" edges to Todo entities.
func (tsuo *TodoStatusUpdateOne) RemoveTodos(t ...*Todo) *TodoStatusUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tsuo.RemoveTodoIDs(ids...)
}

// Where appends a list predicates to the TodoStatusUpdate builder.
func (tsuo *TodoStatusUpdateOne) Where(ps ...predicate.TodoStatus) *TodoStatusUpdateOne {
	tsuo.mutation.Where(ps...)
	return tsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tsuo *TodoStatusUpdateOne) Select(field string, fields ...string) *TodoStatusUpdateOne {
	tsuo.fields = append([]string{field}, fields...)
	return tsuo
}

// Save executes the query and returns the updated TodoStatus entity.
func (tsuo *TodoStatusUpdateOne) Save(ctx context.Context) (*TodoStatus, error) {
	return withHooks[*TodoStatus, TodoStatusMutation](ctx, tsuo.sqlSave, tsuo.mutation, tsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsuo *TodoStatusUpdateOne) SaveX(ctx context.Context) *TodoStatus {
	node, err := tsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tsuo *TodoStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := tsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsuo *TodoStatusUpdateOne) ExecX(ctx context.Context) {
	if err := tsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tsuo *TodoStatusUpdateOne) sqlSave(ctx context.Context) (_node *TodoStatus, err error) {
	_spec := sqlgraph.NewUpdateSpec(todostatus.Table, todostatus.Columns, sqlgraph.NewFieldSpec(todostatus.FieldID, field.TypeInt))
	id, ok := tsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TodoStatus.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todostatus.FieldID)
		for _, f := range fields {
			if !todostatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todostatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsuo.mutation.Status(); ok {
		_spec.SetField(todostatus.FieldStatus, field.TypeString, value)
	}
	if value, ok := tsuo.mutation.CreatedAt(); ok {
		_spec.SetField(todostatus.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(todostatus.FieldUpdatedAt, field.TypeTime, value)
	}
	if tsuo.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todostatus.TodosTable,
			Columns: []string{todostatus.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tsuo.mutation.RemovedTodosIDs(); len(nodes) > 0 && !tsuo.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todostatus.TodosTable,
			Columns: []string{todostatus.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tsuo.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   todostatus.TodosTable,
			Columns: []string{todostatus.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TodoStatus{config: tsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todostatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tsuo.mutation.done = true
	return _node, nil
}
