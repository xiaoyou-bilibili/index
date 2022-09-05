// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"index.data/ent/object"
)

// ObjectCreate is the builder for creating a Object entity.
type ObjectCreate struct {
	config
	mutation *ObjectMutation
	hooks    []Hook
}

// SetObjectName sets the "object_name" field.
func (oc *ObjectCreate) SetObjectName(s string) *ObjectCreate {
	oc.mutation.SetObjectName(s)
	return oc
}

// SetContentType sets the "content_type" field.
func (oc *ObjectCreate) SetContentType(s string) *ObjectCreate {
	oc.mutation.SetContentType(s)
	return oc
}

// SetObjectLocation sets the "object_location" field.
func (oc *ObjectCreate) SetObjectLocation(u uint8) *ObjectCreate {
	oc.mutation.SetObjectLocation(u)
	return oc
}

// SetObjectSize sets the "object_size" field.
func (oc *ObjectCreate) SetObjectSize(i int64) *ObjectCreate {
	oc.mutation.SetObjectSize(i)
	return oc
}

// SetObjectSha256 sets the "object_sha256" field.
func (oc *ObjectCreate) SetObjectSha256(s string) *ObjectCreate {
	oc.mutation.SetObjectSha256(s)
	return oc
}

// SetCreatedAt sets the "created_at" field.
func (oc *ObjectCreate) SetCreatedAt(t time.Time) *ObjectCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *ObjectCreate) SetNillableCreatedAt(t *time.Time) *ObjectCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *ObjectCreate) SetUpdatedAt(t time.Time) *ObjectCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *ObjectCreate) SetNillableUpdatedAt(t *time.Time) *ObjectCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetID sets the "id" field.
func (oc *ObjectCreate) SetID(s string) *ObjectCreate {
	oc.mutation.SetID(s)
	return oc
}

// Mutation returns the ObjectMutation object of the builder.
func (oc *ObjectCreate) Mutation() *ObjectMutation {
	return oc.mutation
}

// Save creates the Object in the database.
func (oc *ObjectCreate) Save(ctx context.Context) (*Object, error) {
	var (
		err  error
		node *Object
	)
	oc.defaults()
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			if node, err = oc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			if oc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *ObjectCreate) SaveX(ctx context.Context) *Object {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *ObjectCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *ObjectCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *ObjectCreate) defaults() {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := object.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := object.DefaultUpdatedAt
		oc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *ObjectCreate) check() error {
	if _, ok := oc.mutation.ObjectName(); !ok {
		return &ValidationError{Name: "object_name", err: errors.New(`ent: missing required field "Object.object_name"`)}
	}
	if _, ok := oc.mutation.ContentType(); !ok {
		return &ValidationError{Name: "content_type", err: errors.New(`ent: missing required field "Object.content_type"`)}
	}
	if _, ok := oc.mutation.ObjectLocation(); !ok {
		return &ValidationError{Name: "object_location", err: errors.New(`ent: missing required field "Object.object_location"`)}
	}
	if _, ok := oc.mutation.ObjectSize(); !ok {
		return &ValidationError{Name: "object_size", err: errors.New(`ent: missing required field "Object.object_size"`)}
	}
	if _, ok := oc.mutation.ObjectSha256(); !ok {
		return &ValidationError{Name: "object_sha256", err: errors.New(`ent: missing required field "Object.object_sha256"`)}
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Object.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Object.updated_at"`)}
	}
	return nil
}

func (oc *ObjectCreate) sqlSave(ctx context.Context) (*Object, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Object.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (oc *ObjectCreate) createSpec() (*Object, *sqlgraph.CreateSpec) {
	var (
		_node = &Object{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: object.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: object.FieldID,
			},
		}
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.ObjectName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: object.FieldObjectName,
		})
		_node.ObjectName = value
	}
	if value, ok := oc.mutation.ContentType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: object.FieldContentType,
		})
		_node.ContentType = value
	}
	if value, ok := oc.mutation.ObjectLocation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: object.FieldObjectLocation,
		})
		_node.ObjectLocation = value
	}
	if value, ok := oc.mutation.ObjectSize(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: object.FieldObjectSize,
		})
		_node.ObjectSize = value
	}
	if value, ok := oc.mutation.ObjectSha256(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: object.FieldObjectSha256,
		})
		_node.ObjectSha256 = value
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: object.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: object.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// ObjectCreateBulk is the builder for creating many Object entities in bulk.
type ObjectCreateBulk struct {
	config
	builders []*ObjectCreate
}

// Save creates the Object entities in the database.
func (ocb *ObjectCreateBulk) Save(ctx context.Context) ([]*Object, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Object, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ObjectMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *ObjectCreateBulk) SaveX(ctx context.Context) []*Object {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *ObjectCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *ObjectCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}