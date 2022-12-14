// Code generated by entc, DO NOT EDIT.

package object

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"index.data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ObjectName applies equality check predicate on the "object_name" field. It's identical to ObjectNameEQ.
func ObjectName(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectName), v))
	})
}

// ContentType applies equality check predicate on the "content_type" field. It's identical to ContentTypeEQ.
func ContentType(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentType), v))
	})
}

// ObjectLocation applies equality check predicate on the "object_location" field. It's identical to ObjectLocationEQ.
func ObjectLocation(v uint8) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectLocation), v))
	})
}

// ObjectSize applies equality check predicate on the "object_size" field. It's identical to ObjectSizeEQ.
func ObjectSize(v int64) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectSize), v))
	})
}

// ObjectSha256 applies equality check predicate on the "object_sha256" field. It's identical to ObjectSha256EQ.
func ObjectSha256(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectSha256), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// ObjectNameEQ applies the EQ predicate on the "object_name" field.
func ObjectNameEQ(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectName), v))
	})
}

// ObjectNameNEQ applies the NEQ predicate on the "object_name" field.
func ObjectNameNEQ(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldObjectName), v))
	})
}

// ObjectNameIn applies the In predicate on the "object_name" field.
func ObjectNameIn(vs ...string) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldObjectName), v...))
	})
}

// ObjectNameNotIn applies the NotIn predicate on the "object_name" field.
func ObjectNameNotIn(vs ...string) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldObjectName), v...))
	})
}

// ObjectNameGT applies the GT predicate on the "object_name" field.
func ObjectNameGT(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldObjectName), v))
	})
}

// ObjectNameGTE applies the GTE predicate on the "object_name" field.
func ObjectNameGTE(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldObjectName), v))
	})
}

// ObjectNameLT applies the LT predicate on the "object_name" field.
func ObjectNameLT(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldObjectName), v))
	})
}

// ObjectNameLTE applies the LTE predicate on the "object_name" field.
func ObjectNameLTE(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldObjectName), v))
	})
}

// ObjectNameContains applies the Contains predicate on the "object_name" field.
func ObjectNameContains(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldObjectName), v))
	})
}

// ObjectNameHasPrefix applies the HasPrefix predicate on the "object_name" field.
func ObjectNameHasPrefix(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldObjectName), v))
	})
}

// ObjectNameHasSuffix applies the HasSuffix predicate on the "object_name" field.
func ObjectNameHasSuffix(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldObjectName), v))
	})
}

// ObjectNameEqualFold applies the EqualFold predicate on the "object_name" field.
func ObjectNameEqualFold(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldObjectName), v))
	})
}

// ObjectNameContainsFold applies the ContainsFold predicate on the "object_name" field.
func ObjectNameContainsFold(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldObjectName), v))
	})
}

// ContentTypeEQ applies the EQ predicate on the "content_type" field.
func ContentTypeEQ(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentType), v))
	})
}

// ContentTypeNEQ applies the NEQ predicate on the "content_type" field.
func ContentTypeNEQ(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContentType), v))
	})
}

// ContentTypeIn applies the In predicate on the "content_type" field.
func ContentTypeIn(vs ...string) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContentType), v...))
	})
}

// ContentTypeNotIn applies the NotIn predicate on the "content_type" field.
func ContentTypeNotIn(vs ...string) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContentType), v...))
	})
}

// ContentTypeGT applies the GT predicate on the "content_type" field.
func ContentTypeGT(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContentType), v))
	})
}

// ContentTypeGTE applies the GTE predicate on the "content_type" field.
func ContentTypeGTE(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContentType), v))
	})
}

// ContentTypeLT applies the LT predicate on the "content_type" field.
func ContentTypeLT(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContentType), v))
	})
}

// ContentTypeLTE applies the LTE predicate on the "content_type" field.
func ContentTypeLTE(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContentType), v))
	})
}

// ContentTypeContains applies the Contains predicate on the "content_type" field.
func ContentTypeContains(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContentType), v))
	})
}

// ContentTypeHasPrefix applies the HasPrefix predicate on the "content_type" field.
func ContentTypeHasPrefix(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContentType), v))
	})
}

// ContentTypeHasSuffix applies the HasSuffix predicate on the "content_type" field.
func ContentTypeHasSuffix(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContentType), v))
	})
}

// ContentTypeEqualFold applies the EqualFold predicate on the "content_type" field.
func ContentTypeEqualFold(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContentType), v))
	})
}

// ContentTypeContainsFold applies the ContainsFold predicate on the "content_type" field.
func ContentTypeContainsFold(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContentType), v))
	})
}

// ObjectLocationEQ applies the EQ predicate on the "object_location" field.
func ObjectLocationEQ(v uint8) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectLocation), v))
	})
}

// ObjectLocationNEQ applies the NEQ predicate on the "object_location" field.
func ObjectLocationNEQ(v uint8) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldObjectLocation), v))
	})
}

// ObjectLocationIn applies the In predicate on the "object_location" field.
func ObjectLocationIn(vs ...uint8) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldObjectLocation), v...))
	})
}

// ObjectLocationNotIn applies the NotIn predicate on the "object_location" field.
func ObjectLocationNotIn(vs ...uint8) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldObjectLocation), v...))
	})
}

// ObjectLocationGT applies the GT predicate on the "object_location" field.
func ObjectLocationGT(v uint8) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldObjectLocation), v))
	})
}

// ObjectLocationGTE applies the GTE predicate on the "object_location" field.
func ObjectLocationGTE(v uint8) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldObjectLocation), v))
	})
}

// ObjectLocationLT applies the LT predicate on the "object_location" field.
func ObjectLocationLT(v uint8) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldObjectLocation), v))
	})
}

// ObjectLocationLTE applies the LTE predicate on the "object_location" field.
func ObjectLocationLTE(v uint8) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldObjectLocation), v))
	})
}

// ObjectSizeEQ applies the EQ predicate on the "object_size" field.
func ObjectSizeEQ(v int64) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectSize), v))
	})
}

// ObjectSizeNEQ applies the NEQ predicate on the "object_size" field.
func ObjectSizeNEQ(v int64) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldObjectSize), v))
	})
}

// ObjectSizeIn applies the In predicate on the "object_size" field.
func ObjectSizeIn(vs ...int64) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldObjectSize), v...))
	})
}

// ObjectSizeNotIn applies the NotIn predicate on the "object_size" field.
func ObjectSizeNotIn(vs ...int64) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldObjectSize), v...))
	})
}

// ObjectSizeGT applies the GT predicate on the "object_size" field.
func ObjectSizeGT(v int64) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldObjectSize), v))
	})
}

// ObjectSizeGTE applies the GTE predicate on the "object_size" field.
func ObjectSizeGTE(v int64) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldObjectSize), v))
	})
}

// ObjectSizeLT applies the LT predicate on the "object_size" field.
func ObjectSizeLT(v int64) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldObjectSize), v))
	})
}

// ObjectSizeLTE applies the LTE predicate on the "object_size" field.
func ObjectSizeLTE(v int64) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldObjectSize), v))
	})
}

// ObjectSha256EQ applies the EQ predicate on the "object_sha256" field.
func ObjectSha256EQ(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObjectSha256), v))
	})
}

// ObjectSha256NEQ applies the NEQ predicate on the "object_sha256" field.
func ObjectSha256NEQ(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldObjectSha256), v))
	})
}

// ObjectSha256In applies the In predicate on the "object_sha256" field.
func ObjectSha256In(vs ...string) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldObjectSha256), v...))
	})
}

// ObjectSha256NotIn applies the NotIn predicate on the "object_sha256" field.
func ObjectSha256NotIn(vs ...string) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldObjectSha256), v...))
	})
}

// ObjectSha256GT applies the GT predicate on the "object_sha256" field.
func ObjectSha256GT(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldObjectSha256), v))
	})
}

// ObjectSha256GTE applies the GTE predicate on the "object_sha256" field.
func ObjectSha256GTE(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldObjectSha256), v))
	})
}

// ObjectSha256LT applies the LT predicate on the "object_sha256" field.
func ObjectSha256LT(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldObjectSha256), v))
	})
}

// ObjectSha256LTE applies the LTE predicate on the "object_sha256" field.
func ObjectSha256LTE(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldObjectSha256), v))
	})
}

// ObjectSha256Contains applies the Contains predicate on the "object_sha256" field.
func ObjectSha256Contains(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldObjectSha256), v))
	})
}

// ObjectSha256HasPrefix applies the HasPrefix predicate on the "object_sha256" field.
func ObjectSha256HasPrefix(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldObjectSha256), v))
	})
}

// ObjectSha256HasSuffix applies the HasSuffix predicate on the "object_sha256" field.
func ObjectSha256HasSuffix(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldObjectSha256), v))
	})
}

// ObjectSha256EqualFold applies the EqualFold predicate on the "object_sha256" field.
func ObjectSha256EqualFold(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldObjectSha256), v))
	})
}

// ObjectSha256ContainsFold applies the ContainsFold predicate on the "object_sha256" field.
func ObjectSha256ContainsFold(v string) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldObjectSha256), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Object {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Object(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Object) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Object) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
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
func Not(p predicate.Object) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		p(s.Not())
	})
}
