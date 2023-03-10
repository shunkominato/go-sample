// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go-gql-sample/app/ent/predicate"
	"go-gql-sample/app/ent/todo"
	"go-gql-sample/app/ent/todostatus"
	"go-gql-sample/app/ent/user"
)

// TodoUpdate is the builder for updating Todo entities.
type TodoUpdate struct {
	config
	hooks    []Hook
	mutation *TodoMutation
}

// Where appends a list predicates to the TodoUpdate builder.
func (tu *TodoUpdate) Where(ps ...predicate.Todo) *TodoUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTodo sets the "todo" field.
func (tu *TodoUpdate) SetTodo(s string) *TodoUpdate {
	tu.mutation.SetTodo(s)
	return tu
}

// SetUserID sets the "user_id" field.
func (tu *TodoUpdate) SetUserID(i int) *TodoUpdate {
	tu.mutation.SetUserID(i)
	return tu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableUserID(i *int) *TodoUpdate {
	if i != nil {
		tu.SetUserID(*i)
	}
	return tu
}

// ClearUserID clears the value of the "user_id" field.
func (tu *TodoUpdate) ClearUserID() *TodoUpdate {
	tu.mutation.ClearUserID()
	return tu
}

// SetTodoStatusesID sets the "todo_statuses_id" field.
func (tu *TodoUpdate) SetTodoStatusesID(i int) *TodoUpdate {
	tu.mutation.SetTodoStatusesID(i)
	return tu
}

// SetNillableTodoStatusesID sets the "todo_statuses_id" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableTodoStatusesID(i *int) *TodoUpdate {
	if i != nil {
		tu.SetTodoStatusesID(*i)
	}
	return tu
}

// ClearTodoStatusesID clears the value of the "todo_statuses_id" field.
func (tu *TodoUpdate) ClearTodoStatusesID() *TodoUpdate {
	tu.mutation.ClearTodoStatusesID()
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TodoUpdate) SetCreatedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TodoUpdate) SetUpdatedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetTodoStatuID sets the "todo_statu" edge to the TodoStatus entity by ID.
func (tu *TodoUpdate) SetTodoStatuID(id int) *TodoUpdate {
	tu.mutation.SetTodoStatuID(id)
	return tu
}

// SetNillableTodoStatuID sets the "todo_statu" edge to the TodoStatus entity by ID if the given value is not nil.
func (tu *TodoUpdate) SetNillableTodoStatuID(id *int) *TodoUpdate {
	if id != nil {
		tu = tu.SetTodoStatuID(*id)
	}
	return tu
}

// SetTodoStatu sets the "todo_statu" edge to the TodoStatus entity.
func (tu *TodoUpdate) SetTodoStatu(t *TodoStatus) *TodoUpdate {
	return tu.SetTodoStatuID(t.ID)
}

// SetUser sets the "user" edge to the User entity.
func (tu *TodoUpdate) SetUser(u *User) *TodoUpdate {
	return tu.SetUserID(u.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tu *TodoUpdate) Mutation() *TodoMutation {
	return tu.mutation
}

// ClearTodoStatu clears the "todo_statu" edge to the TodoStatus entity.
func (tu *TodoUpdate) ClearTodoStatu() *TodoUpdate {
	tu.mutation.ClearTodoStatu()
	return tu
}

// ClearUser clears the "user" edge to the User entity.
func (tu *TodoUpdate) ClearUser() *TodoUpdate {
	tu.mutation.ClearUser()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TodoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TodoMutation](ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TodoUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TodoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TodoUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TodoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Todo(); ok {
		_spec.SetField(todo.FieldTodo, field.TypeString, value)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(todo.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.TodoStatuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.TodoStatuTable,
			Columns: []string{todo.TodoStatuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todostatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TodoStatuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.TodoStatuTable,
			Columns: []string{todo.TodoStatuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todostatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.UserTable,
			Columns: []string{todo.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.UserTable,
			Columns: []string{todo.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TodoUpdateOne is the builder for updating a single Todo entity.
type TodoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoMutation
}

// SetTodo sets the "todo" field.
func (tuo *TodoUpdateOne) SetTodo(s string) *TodoUpdateOne {
	tuo.mutation.SetTodo(s)
	return tuo
}

// SetUserID sets the "user_id" field.
func (tuo *TodoUpdateOne) SetUserID(i int) *TodoUpdateOne {
	tuo.mutation.SetUserID(i)
	return tuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableUserID(i *int) *TodoUpdateOne {
	if i != nil {
		tuo.SetUserID(*i)
	}
	return tuo
}

// ClearUserID clears the value of the "user_id" field.
func (tuo *TodoUpdateOne) ClearUserID() *TodoUpdateOne {
	tuo.mutation.ClearUserID()
	return tuo
}

// SetTodoStatusesID sets the "todo_statuses_id" field.
func (tuo *TodoUpdateOne) SetTodoStatusesID(i int) *TodoUpdateOne {
	tuo.mutation.SetTodoStatusesID(i)
	return tuo
}

// SetNillableTodoStatusesID sets the "todo_statuses_id" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableTodoStatusesID(i *int) *TodoUpdateOne {
	if i != nil {
		tuo.SetTodoStatusesID(*i)
	}
	return tuo
}

// ClearTodoStatusesID clears the value of the "todo_statuses_id" field.
func (tuo *TodoUpdateOne) ClearTodoStatusesID() *TodoUpdateOne {
	tuo.mutation.ClearTodoStatusesID()
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TodoUpdateOne) SetCreatedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TodoUpdateOne) SetUpdatedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetTodoStatuID sets the "todo_statu" edge to the TodoStatus entity by ID.
func (tuo *TodoUpdateOne) SetTodoStatuID(id int) *TodoUpdateOne {
	tuo.mutation.SetTodoStatuID(id)
	return tuo
}

// SetNillableTodoStatuID sets the "todo_statu" edge to the TodoStatus entity by ID if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableTodoStatuID(id *int) *TodoUpdateOne {
	if id != nil {
		tuo = tuo.SetTodoStatuID(*id)
	}
	return tuo
}

// SetTodoStatu sets the "todo_statu" edge to the TodoStatus entity.
func (tuo *TodoUpdateOne) SetTodoStatu(t *TodoStatus) *TodoUpdateOne {
	return tuo.SetTodoStatuID(t.ID)
}

// SetUser sets the "user" edge to the User entity.
func (tuo *TodoUpdateOne) SetUser(u *User) *TodoUpdateOne {
	return tuo.SetUserID(u.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tuo *TodoUpdateOne) Mutation() *TodoMutation {
	return tuo.mutation
}

// ClearTodoStatu clears the "todo_statu" edge to the TodoStatus entity.
func (tuo *TodoUpdateOne) ClearTodoStatu() *TodoUpdateOne {
	tuo.mutation.ClearTodoStatu()
	return tuo
}

// ClearUser clears the "user" edge to the User entity.
func (tuo *TodoUpdateOne) ClearUser() *TodoUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// Where appends a list predicates to the TodoUpdate builder.
func (tuo *TodoUpdateOne) Where(ps ...predicate.Todo) *TodoUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TodoUpdateOne) Select(field string, fields ...string) *TodoUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Todo entity.
func (tuo *TodoUpdateOne) Save(ctx context.Context) (*Todo, error) {
	return withHooks[*Todo, TodoMutation](ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TodoUpdateOne) SaveX(ctx context.Context) *Todo {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TodoUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TodoUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TodoUpdateOne) sqlSave(ctx context.Context) (_node *Todo, err error) {
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Todo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for _, f := range fields {
			if !todo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Todo(); ok {
		_spec.SetField(todo.FieldTodo, field.TypeString, value)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(todo.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.TodoStatuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.TodoStatuTable,
			Columns: []string{todo.TodoStatuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todostatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TodoStatuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.TodoStatuTable,
			Columns: []string{todo.TodoStatuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todostatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.UserTable,
			Columns: []string{todo.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todo.UserTable,
			Columns: []string{todo.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Todo{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
