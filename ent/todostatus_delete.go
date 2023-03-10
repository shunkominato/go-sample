// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"go-gql-sample/app/ent/predicate"
	"go-gql-sample/app/ent/todostatus"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoStatusDelete is the builder for deleting a TodoStatus entity.
type TodoStatusDelete struct {
	config
	hooks    []Hook
	mutation *TodoStatusMutation
}

// Where appends a list predicates to the TodoStatusDelete builder.
func (tsd *TodoStatusDelete) Where(ps ...predicate.TodoStatus) *TodoStatusDelete {
	tsd.mutation.Where(ps...)
	return tsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tsd *TodoStatusDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, TodoStatusMutation](ctx, tsd.sqlExec, tsd.mutation, tsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tsd *TodoStatusDelete) ExecX(ctx context.Context) int {
	n, err := tsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tsd *TodoStatusDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(todostatus.Table, sqlgraph.NewFieldSpec(todostatus.FieldID, field.TypeInt))
	if ps := tsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tsd.mutation.done = true
	return affected, err
}

// TodoStatusDeleteOne is the builder for deleting a single TodoStatus entity.
type TodoStatusDeleteOne struct {
	tsd *TodoStatusDelete
}

// Where appends a list predicates to the TodoStatusDelete builder.
func (tsdo *TodoStatusDeleteOne) Where(ps ...predicate.TodoStatus) *TodoStatusDeleteOne {
	tsdo.tsd.mutation.Where(ps...)
	return tsdo
}

// Exec executes the deletion query.
func (tsdo *TodoStatusDeleteOne) Exec(ctx context.Context) error {
	n, err := tsdo.tsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{todostatus.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tsdo *TodoStatusDeleteOne) ExecX(ctx context.Context) {
	if err := tsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
