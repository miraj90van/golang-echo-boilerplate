// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"myapp/helper"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {

	schema := []ent.Field{
		//WARNING : mandatory value, make sure on repository on creat() function for model always set below field
		field.Uint64("id"),
		field.String("name"),
		field.String("password"),
		field.String("avatar"),
		field.Int32("role_id"),


		//WARNING : becarefull with bool value, it always have default value as FALSE,
		//make sure when do testing DTO / request and actual mock or return value have same value
		field.Bool("is_verified"),

		//OPTIONAL VALUE
		field.String("email").Optional().Unique(),
		field.Time("email_verified_at").Optional(),
		field.String("remember_token").Optional(),
		field.String("social_media_id").Optional(),
		field.String("login_type").Optional(),
		field.String("sub_specialist").Optional(),
		field.String("firebase_token").Optional(),
		field.String("info").Optional(),
		field.String("description").Optional(),
		field.String("specialist").Optional(),
		field.String("phone").Optional().Unique(),
		field.Time("last_access_at").Optional(),
		field.Bool("pregnancy_mode").Optional(),
		field.Time("latest_skip_update").Optional(),
		field.Time("latest_deleted_at").Optional()}

	return helper.InitBaseSchema(schema)
}
func (User) Edges() []ent.Edge {
	return nil
}
func (User) Annotations() []schema.Annotation {
	return nil
}
