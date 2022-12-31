// ErrorCode generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SystemParametersColumns holds the columns for the "system_parameters" table.
	SystemParametersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "value", Type: field.TypeString},
	}
	// SystemParametersTable holds the schema information for the "system_parameters" table.
	SystemParametersTable = &schema.Table{
		Name:       "system_parameters",
		Columns:    SystemParametersColumns,
		PrimaryKey: []*schema.Column{SystemParametersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SystemParametersTable,
	}
)

func init() {
}
