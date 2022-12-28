package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"myapp/commons"
)

// System_parameter holds the schema definition for the System_parameter entity.
type System_parameter struct {
	ent.Schema
}

// Fields of the System_parameter.
func (System_parameter) Fields() []ent.Field {
	schema := []ent.Field{
		field.String("key").NotEmpty().Unique(),
		field.String("value").NotEmpty(),
	}

	return commons.InitBaseSchema(schema)
}

// Edges of the System_parameter.
func (System_parameter) Edges() []ent.Edge {
	return nil
}
