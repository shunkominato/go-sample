// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"story.com/story/app/ent/todo"
	"story.com/story/app/ent/todostatus"
	"story.com/story/app/ent/user"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Todo holds the value of the "todo" field.
	Todo string `json:"todo,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// TodoStatusesID holds the value of the "todo_statuses_id" field.
	TodoStatusesID int `json:"todo_statuses_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TodoQuery when eager-loading is set.
	Edges TodoEdges `json:"edges"`
}

// TodoEdges holds the relations/edges for other nodes in the graph.
type TodoEdges struct {
	// TodoStatu holds the value of the todo_statu edge.
	TodoStatu *TodoStatus `json:"todo_statu,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TodoStatuOrErr returns the TodoStatu value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TodoEdges) TodoStatuOrErr() (*TodoStatus, error) {
	if e.loadedTypes[0] {
		if e.TodoStatu == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: todostatus.Label}
		}
		return e.TodoStatu, nil
	}
	return nil, &NotLoadedError{edge: "todo_statu"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TodoEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case todo.FieldID, todo.FieldUserID, todo.FieldTodoStatusesID:
			values[i] = new(sql.NullInt64)
		case todo.FieldTodo:
			values[i] = new(sql.NullString)
		case todo.FieldCreatedAt, todo.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Todo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case todo.FieldTodo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field todo", values[i])
			} else if value.Valid {
				t.Todo = value.String
			}
		case todo.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				t.UserID = int(value.Int64)
			}
		case todo.FieldTodoStatusesID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field todo_statuses_id", values[i])
			} else if value.Valid {
				t.TodoStatusesID = int(value.Int64)
			}
		case todo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case todo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTodoStatu queries the "todo_statu" edge of the Todo entity.
func (t *Todo) QueryTodoStatu() *TodoStatusQuery {
	return NewTodoClient(t.config).QueryTodoStatu(t)
}

// QueryUser queries the "user" edge of the Todo entity.
func (t *Todo) QueryUser() *UserQuery {
	return NewTodoClient(t.config).QueryUser(t)
}

// Update returns a builder for updating this Todo.
// Note that you need to call Todo.Unwrap() before calling this method if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return NewTodoClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Todo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("todo=")
	builder.WriteString(t.Todo)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", t.UserID))
	builder.WriteString(", ")
	builder.WriteString("todo_statuses_id=")
	builder.WriteString(fmt.Sprintf("%v", t.TodoStatusesID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Todos is a parsable slice of Todo.
type Todos []*Todo
