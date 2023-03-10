// Code generated by ent, DO NOT EDIT.

package arinternalmetadatum

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"go-gql-sample/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldLTE(FieldID, id))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldEQ(FieldValue, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldEQ(FieldUpdatedAt, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldHasSuffix(FieldValue, v))
}

// ValueIsNil applies the IsNil predicate on the "value" field.
func ValueIsNil() predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldIsNull(FieldValue))
}

// ValueNotNil applies the NotNil predicate on the "value" field.
func ValueNotNil() predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldNotNull(FieldValue))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldContainsFold(FieldValue, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ArInternalMetadatum) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ArInternalMetadatum) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ArInternalMetadatum) predicate.ArInternalMetadatum {
	return predicate.ArInternalMetadatum(func(s *sql.Selector) {
		p(s.Not())
	})
}
