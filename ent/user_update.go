// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-gql-sample/app/ent/predicate"
	"go-gql-sample/app/ent/todo"
	"go-gql-sample/app/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetProvider sets the "provider" field.
func (uu *UserUpdate) SetProvider(s string) *UserUpdate {
	uu.mutation.SetProvider(s)
	return uu
}

// SetUID sets the "uid" field.
func (uu *UserUpdate) SetUID(s string) *UserUpdate {
	uu.mutation.SetUID(s)
	return uu
}

// SetEncryptedPassword sets the "encrypted_password" field.
func (uu *UserUpdate) SetEncryptedPassword(s string) *UserUpdate {
	uu.mutation.SetEncryptedPassword(s)
	return uu
}

// SetResetPasswordToken sets the "reset_password_token" field.
func (uu *UserUpdate) SetResetPasswordToken(s string) *UserUpdate {
	uu.mutation.SetResetPasswordToken(s)
	return uu
}

// SetNillableResetPasswordToken sets the "reset_password_token" field if the given value is not nil.
func (uu *UserUpdate) SetNillableResetPasswordToken(s *string) *UserUpdate {
	if s != nil {
		uu.SetResetPasswordToken(*s)
	}
	return uu
}

// ClearResetPasswordToken clears the value of the "reset_password_token" field.
func (uu *UserUpdate) ClearResetPasswordToken() *UserUpdate {
	uu.mutation.ClearResetPasswordToken()
	return uu
}

// SetResetPasswordSentAt sets the "reset_password_sent_at" field.
func (uu *UserUpdate) SetResetPasswordSentAt(t time.Time) *UserUpdate {
	uu.mutation.SetResetPasswordSentAt(t)
	return uu
}

// SetNillableResetPasswordSentAt sets the "reset_password_sent_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableResetPasswordSentAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetResetPasswordSentAt(*t)
	}
	return uu
}

// ClearResetPasswordSentAt clears the value of the "reset_password_sent_at" field.
func (uu *UserUpdate) ClearResetPasswordSentAt() *UserUpdate {
	uu.mutation.ClearResetPasswordSentAt()
	return uu
}

// SetAllowPasswordChange sets the "allow_password_change" field.
func (uu *UserUpdate) SetAllowPasswordChange(b bool) *UserUpdate {
	uu.mutation.SetAllowPasswordChange(b)
	return uu
}

// SetNillableAllowPasswordChange sets the "allow_password_change" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAllowPasswordChange(b *bool) *UserUpdate {
	if b != nil {
		uu.SetAllowPasswordChange(*b)
	}
	return uu
}

// ClearAllowPasswordChange clears the value of the "allow_password_change" field.
func (uu *UserUpdate) ClearAllowPasswordChange() *UserUpdate {
	uu.mutation.ClearAllowPasswordChange()
	return uu
}

// SetRememberCreatedAt sets the "remember_created_at" field.
func (uu *UserUpdate) SetRememberCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetRememberCreatedAt(t)
	return uu
}

// SetNillableRememberCreatedAt sets the "remember_created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRememberCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetRememberCreatedAt(*t)
	}
	return uu
}

// ClearRememberCreatedAt clears the value of the "remember_created_at" field.
func (uu *UserUpdate) ClearRememberCreatedAt() *UserUpdate {
	uu.mutation.ClearRememberCreatedAt()
	return uu
}

// SetConfirmationToken sets the "confirmation_token" field.
func (uu *UserUpdate) SetConfirmationToken(s string) *UserUpdate {
	uu.mutation.SetConfirmationToken(s)
	return uu
}

// SetNillableConfirmationToken sets the "confirmation_token" field if the given value is not nil.
func (uu *UserUpdate) SetNillableConfirmationToken(s *string) *UserUpdate {
	if s != nil {
		uu.SetConfirmationToken(*s)
	}
	return uu
}

// ClearConfirmationToken clears the value of the "confirmation_token" field.
func (uu *UserUpdate) ClearConfirmationToken() *UserUpdate {
	uu.mutation.ClearConfirmationToken()
	return uu
}

// SetConfirmedAt sets the "confirmed_at" field.
func (uu *UserUpdate) SetConfirmedAt(t time.Time) *UserUpdate {
	uu.mutation.SetConfirmedAt(t)
	return uu
}

// SetNillableConfirmedAt sets the "confirmed_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableConfirmedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetConfirmedAt(*t)
	}
	return uu
}

// ClearConfirmedAt clears the value of the "confirmed_at" field.
func (uu *UserUpdate) ClearConfirmedAt() *UserUpdate {
	uu.mutation.ClearConfirmedAt()
	return uu
}

// SetConfirmationSentAt sets the "confirmation_sent_at" field.
func (uu *UserUpdate) SetConfirmationSentAt(t time.Time) *UserUpdate {
	uu.mutation.SetConfirmationSentAt(t)
	return uu
}

// SetNillableConfirmationSentAt sets the "confirmation_sent_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableConfirmationSentAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetConfirmationSentAt(*t)
	}
	return uu
}

// ClearConfirmationSentAt clears the value of the "confirmation_sent_at" field.
func (uu *UserUpdate) ClearConfirmationSentAt() *UserUpdate {
	uu.mutation.ClearConfirmationSentAt()
	return uu
}

// SetUnconfirmedEmail sets the "unconfirmed_email" field.
func (uu *UserUpdate) SetUnconfirmedEmail(s string) *UserUpdate {
	uu.mutation.SetUnconfirmedEmail(s)
	return uu
}

// SetNillableUnconfirmedEmail sets the "unconfirmed_email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUnconfirmedEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetUnconfirmedEmail(*s)
	}
	return uu
}

// ClearUnconfirmedEmail clears the value of the "unconfirmed_email" field.
func (uu *UserUpdate) ClearUnconfirmedEmail() *UserUpdate {
	uu.mutation.ClearUnconfirmedEmail()
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// ClearName clears the value of the "name" field.
func (uu *UserUpdate) ClearName() *UserUpdate {
	uu.mutation.ClearName()
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickname(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickname(*s)
	}
	return uu
}

// ClearNickname clears the value of the "nickname" field.
func (uu *UserUpdate) ClearNickname() *UserUpdate {
	uu.mutation.ClearNickname()
	return uu
}

// SetImage sets the "image" field.
func (uu *UserUpdate) SetImage(s string) *UserUpdate {
	uu.mutation.SetImage(s)
	return uu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (uu *UserUpdate) SetNillableImage(s *string) *UserUpdate {
	if s != nil {
		uu.SetImage(*s)
	}
	return uu
}

// ClearImage clears the value of the "image" field.
func (uu *UserUpdate) ClearImage() *UserUpdate {
	uu.mutation.ClearImage()
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of the "email" field.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetTokens sets the "tokens" field.
func (uu *UserUpdate) SetTokens(s []string) *UserUpdate {
	uu.mutation.SetTokens(s)
	return uu
}

// AppendTokens appends s to the "tokens" field.
func (uu *UserUpdate) AppendTokens(s []string) *UserUpdate {
	uu.mutation.AppendTokens(s)
	return uu
}

// ClearTokens clears the value of the "tokens" field.
func (uu *UserUpdate) ClearTokens() *UserUpdate {
	uu.mutation.ClearTokens()
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (uu *UserUpdate) AddTodoIDs(ids ...int) *UserUpdate {
	uu.mutation.AddTodoIDs(ids...)
	return uu
}

// AddTodos adds the "todos" edges to the Todo entity.
func (uu *UserUpdate) AddTodos(t ...*Todo) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTodoIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearTodos clears all "todos" edges to the Todo entity.
func (uu *UserUpdate) ClearTodos() *UserUpdate {
	uu.mutation.ClearTodos()
	return uu
}

// RemoveTodoIDs removes the "todos" edge to Todo entities by IDs.
func (uu *UserUpdate) RemoveTodoIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveTodoIDs(ids...)
	return uu
}

// RemoveTodos removes "todos" edges to Todo entities.
func (uu *UserUpdate) RemoveTodos(t ...*Todo) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTodoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Provider(); ok {
		_spec.SetField(user.FieldProvider, field.TypeString, value)
	}
	if value, ok := uu.mutation.UID(); ok {
		_spec.SetField(user.FieldUID, field.TypeString, value)
	}
	if value, ok := uu.mutation.EncryptedPassword(); ok {
		_spec.SetField(user.FieldEncryptedPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.ResetPasswordToken(); ok {
		_spec.SetField(user.FieldResetPasswordToken, field.TypeString, value)
	}
	if uu.mutation.ResetPasswordTokenCleared() {
		_spec.ClearField(user.FieldResetPasswordToken, field.TypeString)
	}
	if value, ok := uu.mutation.ResetPasswordSentAt(); ok {
		_spec.SetField(user.FieldResetPasswordSentAt, field.TypeTime, value)
	}
	if uu.mutation.ResetPasswordSentAtCleared() {
		_spec.ClearField(user.FieldResetPasswordSentAt, field.TypeTime)
	}
	if value, ok := uu.mutation.AllowPasswordChange(); ok {
		_spec.SetField(user.FieldAllowPasswordChange, field.TypeBool, value)
	}
	if uu.mutation.AllowPasswordChangeCleared() {
		_spec.ClearField(user.FieldAllowPasswordChange, field.TypeBool)
	}
	if value, ok := uu.mutation.RememberCreatedAt(); ok {
		_spec.SetField(user.FieldRememberCreatedAt, field.TypeTime, value)
	}
	if uu.mutation.RememberCreatedAtCleared() {
		_spec.ClearField(user.FieldRememberCreatedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.ConfirmationToken(); ok {
		_spec.SetField(user.FieldConfirmationToken, field.TypeString, value)
	}
	if uu.mutation.ConfirmationTokenCleared() {
		_spec.ClearField(user.FieldConfirmationToken, field.TypeString)
	}
	if value, ok := uu.mutation.ConfirmedAt(); ok {
		_spec.SetField(user.FieldConfirmedAt, field.TypeTime, value)
	}
	if uu.mutation.ConfirmedAtCleared() {
		_spec.ClearField(user.FieldConfirmedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.ConfirmationSentAt(); ok {
		_spec.SetField(user.FieldConfirmationSentAt, field.TypeTime, value)
	}
	if uu.mutation.ConfirmationSentAtCleared() {
		_spec.ClearField(user.FieldConfirmationSentAt, field.TypeTime)
	}
	if value, ok := uu.mutation.UnconfirmedEmail(); ok {
		_spec.SetField(user.FieldUnconfirmedEmail, field.TypeString, value)
	}
	if uu.mutation.UnconfirmedEmailCleared() {
		_spec.ClearField(user.FieldUnconfirmedEmail, field.TypeString)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if uu.mutation.NameCleared() {
		_spec.ClearField(user.FieldName, field.TypeString)
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if uu.mutation.NicknameCleared() {
		_spec.ClearField(user.FieldNickname, field.TypeString)
	}
	if value, ok := uu.mutation.Image(); ok {
		_spec.SetField(user.FieldImage, field.TypeString, value)
	}
	if uu.mutation.ImageCleared() {
		_spec.ClearField(user.FieldImage, field.TypeString)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uu.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uu.mutation.Tokens(); ok {
		_spec.SetField(user.FieldTokens, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedTokens(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldTokens, value)
		})
	}
	if uu.mutation.TokensCleared() {
		_spec.ClearField(user.FieldTokens, field.TypeJSON)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
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
	if nodes := uu.mutation.RemovedTodosIDs(); len(nodes) > 0 && !uu.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
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
	if nodes := uu.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetProvider sets the "provider" field.
func (uuo *UserUpdateOne) SetProvider(s string) *UserUpdateOne {
	uuo.mutation.SetProvider(s)
	return uuo
}

// SetUID sets the "uid" field.
func (uuo *UserUpdateOne) SetUID(s string) *UserUpdateOne {
	uuo.mutation.SetUID(s)
	return uuo
}

// SetEncryptedPassword sets the "encrypted_password" field.
func (uuo *UserUpdateOne) SetEncryptedPassword(s string) *UserUpdateOne {
	uuo.mutation.SetEncryptedPassword(s)
	return uuo
}

// SetResetPasswordToken sets the "reset_password_token" field.
func (uuo *UserUpdateOne) SetResetPasswordToken(s string) *UserUpdateOne {
	uuo.mutation.SetResetPasswordToken(s)
	return uuo
}

// SetNillableResetPasswordToken sets the "reset_password_token" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableResetPasswordToken(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetResetPasswordToken(*s)
	}
	return uuo
}

// ClearResetPasswordToken clears the value of the "reset_password_token" field.
func (uuo *UserUpdateOne) ClearResetPasswordToken() *UserUpdateOne {
	uuo.mutation.ClearResetPasswordToken()
	return uuo
}

// SetResetPasswordSentAt sets the "reset_password_sent_at" field.
func (uuo *UserUpdateOne) SetResetPasswordSentAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetResetPasswordSentAt(t)
	return uuo
}

// SetNillableResetPasswordSentAt sets the "reset_password_sent_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableResetPasswordSentAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetResetPasswordSentAt(*t)
	}
	return uuo
}

// ClearResetPasswordSentAt clears the value of the "reset_password_sent_at" field.
func (uuo *UserUpdateOne) ClearResetPasswordSentAt() *UserUpdateOne {
	uuo.mutation.ClearResetPasswordSentAt()
	return uuo
}

// SetAllowPasswordChange sets the "allow_password_change" field.
func (uuo *UserUpdateOne) SetAllowPasswordChange(b bool) *UserUpdateOne {
	uuo.mutation.SetAllowPasswordChange(b)
	return uuo
}

// SetNillableAllowPasswordChange sets the "allow_password_change" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAllowPasswordChange(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetAllowPasswordChange(*b)
	}
	return uuo
}

// ClearAllowPasswordChange clears the value of the "allow_password_change" field.
func (uuo *UserUpdateOne) ClearAllowPasswordChange() *UserUpdateOne {
	uuo.mutation.ClearAllowPasswordChange()
	return uuo
}

// SetRememberCreatedAt sets the "remember_created_at" field.
func (uuo *UserUpdateOne) SetRememberCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetRememberCreatedAt(t)
	return uuo
}

// SetNillableRememberCreatedAt sets the "remember_created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRememberCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetRememberCreatedAt(*t)
	}
	return uuo
}

// ClearRememberCreatedAt clears the value of the "remember_created_at" field.
func (uuo *UserUpdateOne) ClearRememberCreatedAt() *UserUpdateOne {
	uuo.mutation.ClearRememberCreatedAt()
	return uuo
}

// SetConfirmationToken sets the "confirmation_token" field.
func (uuo *UserUpdateOne) SetConfirmationToken(s string) *UserUpdateOne {
	uuo.mutation.SetConfirmationToken(s)
	return uuo
}

// SetNillableConfirmationToken sets the "confirmation_token" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableConfirmationToken(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetConfirmationToken(*s)
	}
	return uuo
}

// ClearConfirmationToken clears the value of the "confirmation_token" field.
func (uuo *UserUpdateOne) ClearConfirmationToken() *UserUpdateOne {
	uuo.mutation.ClearConfirmationToken()
	return uuo
}

// SetConfirmedAt sets the "confirmed_at" field.
func (uuo *UserUpdateOne) SetConfirmedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetConfirmedAt(t)
	return uuo
}

// SetNillableConfirmedAt sets the "confirmed_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableConfirmedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetConfirmedAt(*t)
	}
	return uuo
}

// ClearConfirmedAt clears the value of the "confirmed_at" field.
func (uuo *UserUpdateOne) ClearConfirmedAt() *UserUpdateOne {
	uuo.mutation.ClearConfirmedAt()
	return uuo
}

// SetConfirmationSentAt sets the "confirmation_sent_at" field.
func (uuo *UserUpdateOne) SetConfirmationSentAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetConfirmationSentAt(t)
	return uuo
}

// SetNillableConfirmationSentAt sets the "confirmation_sent_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableConfirmationSentAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetConfirmationSentAt(*t)
	}
	return uuo
}

// ClearConfirmationSentAt clears the value of the "confirmation_sent_at" field.
func (uuo *UserUpdateOne) ClearConfirmationSentAt() *UserUpdateOne {
	uuo.mutation.ClearConfirmationSentAt()
	return uuo
}

// SetUnconfirmedEmail sets the "unconfirmed_email" field.
func (uuo *UserUpdateOne) SetUnconfirmedEmail(s string) *UserUpdateOne {
	uuo.mutation.SetUnconfirmedEmail(s)
	return uuo
}

// SetNillableUnconfirmedEmail sets the "unconfirmed_email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUnconfirmedEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUnconfirmedEmail(*s)
	}
	return uuo
}

// ClearUnconfirmedEmail clears the value of the "unconfirmed_email" field.
func (uuo *UserUpdateOne) ClearUnconfirmedEmail() *UserUpdateOne {
	uuo.mutation.ClearUnconfirmedEmail()
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// ClearName clears the value of the "name" field.
func (uuo *UserUpdateOne) ClearName() *UserUpdateOne {
	uuo.mutation.ClearName()
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickname(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickname(*s)
	}
	return uuo
}

// ClearNickname clears the value of the "nickname" field.
func (uuo *UserUpdateOne) ClearNickname() *UserUpdateOne {
	uuo.mutation.ClearNickname()
	return uuo
}

// SetImage sets the "image" field.
func (uuo *UserUpdateOne) SetImage(s string) *UserUpdateOne {
	uuo.mutation.SetImage(s)
	return uuo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableImage(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetImage(*s)
	}
	return uuo
}

// ClearImage clears the value of the "image" field.
func (uuo *UserUpdateOne) ClearImage() *UserUpdateOne {
	uuo.mutation.ClearImage()
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of the "email" field.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetTokens sets the "tokens" field.
func (uuo *UserUpdateOne) SetTokens(s []string) *UserUpdateOne {
	uuo.mutation.SetTokens(s)
	return uuo
}

// AppendTokens appends s to the "tokens" field.
func (uuo *UserUpdateOne) AppendTokens(s []string) *UserUpdateOne {
	uuo.mutation.AppendTokens(s)
	return uuo
}

// ClearTokens clears the value of the "tokens" field.
func (uuo *UserUpdateOne) ClearTokens() *UserUpdateOne {
	uuo.mutation.ClearTokens()
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (uuo *UserUpdateOne) AddTodoIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddTodoIDs(ids...)
	return uuo
}

// AddTodos adds the "todos" edges to the Todo entity.
func (uuo *UserUpdateOne) AddTodos(t ...*Todo) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTodoIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearTodos clears all "todos" edges to the Todo entity.
func (uuo *UserUpdateOne) ClearTodos() *UserUpdateOne {
	uuo.mutation.ClearTodos()
	return uuo
}

// RemoveTodoIDs removes the "todos" edge to Todo entities by IDs.
func (uuo *UserUpdateOne) RemoveTodoIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveTodoIDs(ids...)
	return uuo
}

// RemoveTodos removes "todos" edges to Todo entities.
func (uuo *UserUpdateOne) RemoveTodos(t ...*Todo) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTodoIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Provider(); ok {
		_spec.SetField(user.FieldProvider, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UID(); ok {
		_spec.SetField(user.FieldUID, field.TypeString, value)
	}
	if value, ok := uuo.mutation.EncryptedPassword(); ok {
		_spec.SetField(user.FieldEncryptedPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.ResetPasswordToken(); ok {
		_spec.SetField(user.FieldResetPasswordToken, field.TypeString, value)
	}
	if uuo.mutation.ResetPasswordTokenCleared() {
		_spec.ClearField(user.FieldResetPasswordToken, field.TypeString)
	}
	if value, ok := uuo.mutation.ResetPasswordSentAt(); ok {
		_spec.SetField(user.FieldResetPasswordSentAt, field.TypeTime, value)
	}
	if uuo.mutation.ResetPasswordSentAtCleared() {
		_spec.ClearField(user.FieldResetPasswordSentAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.AllowPasswordChange(); ok {
		_spec.SetField(user.FieldAllowPasswordChange, field.TypeBool, value)
	}
	if uuo.mutation.AllowPasswordChangeCleared() {
		_spec.ClearField(user.FieldAllowPasswordChange, field.TypeBool)
	}
	if value, ok := uuo.mutation.RememberCreatedAt(); ok {
		_spec.SetField(user.FieldRememberCreatedAt, field.TypeTime, value)
	}
	if uuo.mutation.RememberCreatedAtCleared() {
		_spec.ClearField(user.FieldRememberCreatedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.ConfirmationToken(); ok {
		_spec.SetField(user.FieldConfirmationToken, field.TypeString, value)
	}
	if uuo.mutation.ConfirmationTokenCleared() {
		_spec.ClearField(user.FieldConfirmationToken, field.TypeString)
	}
	if value, ok := uuo.mutation.ConfirmedAt(); ok {
		_spec.SetField(user.FieldConfirmedAt, field.TypeTime, value)
	}
	if uuo.mutation.ConfirmedAtCleared() {
		_spec.ClearField(user.FieldConfirmedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.ConfirmationSentAt(); ok {
		_spec.SetField(user.FieldConfirmationSentAt, field.TypeTime, value)
	}
	if uuo.mutation.ConfirmationSentAtCleared() {
		_spec.ClearField(user.FieldConfirmationSentAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.UnconfirmedEmail(); ok {
		_spec.SetField(user.FieldUnconfirmedEmail, field.TypeString, value)
	}
	if uuo.mutation.UnconfirmedEmailCleared() {
		_spec.ClearField(user.FieldUnconfirmedEmail, field.TypeString)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if uuo.mutation.NameCleared() {
		_spec.ClearField(user.FieldName, field.TypeString)
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if uuo.mutation.NicknameCleared() {
		_spec.ClearField(user.FieldNickname, field.TypeString)
	}
	if value, ok := uuo.mutation.Image(); ok {
		_spec.SetField(user.FieldImage, field.TypeString, value)
	}
	if uuo.mutation.ImageCleared() {
		_spec.ClearField(user.FieldImage, field.TypeString)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uuo.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uuo.mutation.Tokens(); ok {
		_spec.SetField(user.FieldTokens, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedTokens(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldTokens, value)
		})
	}
	if uuo.mutation.TokensCleared() {
		_spec.ClearField(user.FieldTokens, field.TypeJSON)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
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
	if nodes := uuo.mutation.RemovedTodosIDs(); len(nodes) > 0 && !uuo.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
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
	if nodes := uuo.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
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
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}