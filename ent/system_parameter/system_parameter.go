// ErrorCode generated by ent, DO NOT EDIT.

package system_parameter

const (
	// Label holds the string label denoting the system_parameter type in the database.
	Label = "system_parameter"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// Table holds the table name of the system_parameter in the database.
	Table = "system_parameters"
)

// Columns holds all SQL columns for system_parameter fields.
var Columns = []string{
	FieldID,
	FieldKey,
	FieldValue,
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

var (
	// KeyValidator is a validator for the "key" field. It is called by the builders before save.
	KeyValidator func(string) error
	// ValueValidator is a validator for the "value" field. It is called by the builders before save.
	ValueValidator func(string) error
)
