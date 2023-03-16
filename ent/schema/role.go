// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{field.Uint64("id"), field.String("name"), field.Time("created_at").Optional(), field.Time("updated_at").Optional(), field.String("text"), field.Time("deleted_at").Optional()}
}
func (Role) Edges() []ent.Edge {
	return nil
}
func (Role) Annotations() []schema.Annotation {
	return nil
}
