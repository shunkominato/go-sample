// Code generated by ent, DO NOT EDIT.

package schemamigration

const (
	// Label holds the string label denoting the schemamigration type in the database.
	Label = "schema_migration"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "version"
	// Table holds the table name of the schemamigration in the database.
	Table = "schema_migrations"
)

// Columns holds all SQL columns for schemamigration fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
