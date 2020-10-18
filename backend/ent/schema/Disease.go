package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)

// Disease schema.
type Disease struct {
    ent.Schema
}

// Fields of the disease.
func (Disease) Fields() []ent.Field {
    return []ent.Field{
        field.String("disease_name").NotEmpty(),
    }
}

// Edges of the Disease.
func (Disease) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("patients", Patient.Type).StorageKey(edge.Column("disease_id")),
    }
}