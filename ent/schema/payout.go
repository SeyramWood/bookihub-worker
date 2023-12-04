package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Payout holds the schema definition for the Payout entity.
type Payout struct {
	ent.Schema
}

// Fields of the Payout.
func (Payout) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_name").NotEmpty(),
		field.String("account_number").NotEmpty(),
		field.String("bank").NotEmpty(),
		field.String("branch").NotEmpty(),
	}
}

// Edges of the Payout.
func (Payout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transaction", Transaction.Type).Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
