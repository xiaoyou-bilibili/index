// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"index.data/ent/object"
	"index.data/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeObject = "Object"
)

// ObjectMutation represents an operation that mutates the Object nodes in the graph.
type ObjectMutation struct {
	config
	op                 Op
	typ                string
	id                 *string
	object_name        *string
	content_type       *string
	object_location    *uint8
	addobject_location *int8
	object_size        *int64
	addobject_size     *int64
	object_sha256      *string
	created_at         *time.Time
	updated_at         *time.Time
	clearedFields      map[string]struct{}
	done               bool
	oldValue           func(context.Context) (*Object, error)
	predicates         []predicate.Object
}

var _ ent.Mutation = (*ObjectMutation)(nil)

// objectOption allows management of the mutation configuration using functional options.
type objectOption func(*ObjectMutation)

// newObjectMutation creates new mutation for the Object entity.
func newObjectMutation(c config, op Op, opts ...objectOption) *ObjectMutation {
	m := &ObjectMutation{
		config:        c,
		op:            op,
		typ:           TypeObject,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withObjectID sets the ID field of the mutation.
func withObjectID(id string) objectOption {
	return func(m *ObjectMutation) {
		var (
			err   error
			once  sync.Once
			value *Object
		)
		m.oldValue = func(ctx context.Context) (*Object, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Object.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withObject sets the old Object of the mutation.
func withObject(node *Object) objectOption {
	return func(m *ObjectMutation) {
		m.oldValue = func(context.Context) (*Object, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ObjectMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ObjectMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Object entities.
func (m *ObjectMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ObjectMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ObjectMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Object.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetObjectName sets the "object_name" field.
func (m *ObjectMutation) SetObjectName(s string) {
	m.object_name = &s
}

// ObjectName returns the value of the "object_name" field in the mutation.
func (m *ObjectMutation) ObjectName() (r string, exists bool) {
	v := m.object_name
	if v == nil {
		return
	}
	return *v, true
}

// OldObjectName returns the old "object_name" field's value of the Object entity.
// If the Object object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ObjectMutation) OldObjectName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldObjectName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldObjectName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldObjectName: %w", err)
	}
	return oldValue.ObjectName, nil
}

// ResetObjectName resets all changes to the "object_name" field.
func (m *ObjectMutation) ResetObjectName() {
	m.object_name = nil
}

// SetContentType sets the "content_type" field.
func (m *ObjectMutation) SetContentType(s string) {
	m.content_type = &s
}

// ContentType returns the value of the "content_type" field in the mutation.
func (m *ObjectMutation) ContentType() (r string, exists bool) {
	v := m.content_type
	if v == nil {
		return
	}
	return *v, true
}

// OldContentType returns the old "content_type" field's value of the Object entity.
// If the Object object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ObjectMutation) OldContentType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContentType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContentType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContentType: %w", err)
	}
	return oldValue.ContentType, nil
}

// ResetContentType resets all changes to the "content_type" field.
func (m *ObjectMutation) ResetContentType() {
	m.content_type = nil
}

// SetObjectLocation sets the "object_location" field.
func (m *ObjectMutation) SetObjectLocation(u uint8) {
	m.object_location = &u
	m.addobject_location = nil
}

// ObjectLocation returns the value of the "object_location" field in the mutation.
func (m *ObjectMutation) ObjectLocation() (r uint8, exists bool) {
	v := m.object_location
	if v == nil {
		return
	}
	return *v, true
}

// OldObjectLocation returns the old "object_location" field's value of the Object entity.
// If the Object object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ObjectMutation) OldObjectLocation(ctx context.Context) (v uint8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldObjectLocation is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldObjectLocation requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldObjectLocation: %w", err)
	}
	return oldValue.ObjectLocation, nil
}

// AddObjectLocation adds u to the "object_location" field.
func (m *ObjectMutation) AddObjectLocation(u int8) {
	if m.addobject_location != nil {
		*m.addobject_location += u
	} else {
		m.addobject_location = &u
	}
}

// AddedObjectLocation returns the value that was added to the "object_location" field in this mutation.
func (m *ObjectMutation) AddedObjectLocation() (r int8, exists bool) {
	v := m.addobject_location
	if v == nil {
		return
	}
	return *v, true
}

// ResetObjectLocation resets all changes to the "object_location" field.
func (m *ObjectMutation) ResetObjectLocation() {
	m.object_location = nil
	m.addobject_location = nil
}

// SetObjectSize sets the "object_size" field.
func (m *ObjectMutation) SetObjectSize(i int64) {
	m.object_size = &i
	m.addobject_size = nil
}

// ObjectSize returns the value of the "object_size" field in the mutation.
func (m *ObjectMutation) ObjectSize() (r int64, exists bool) {
	v := m.object_size
	if v == nil {
		return
	}
	return *v, true
}

// OldObjectSize returns the old "object_size" field's value of the Object entity.
// If the Object object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ObjectMutation) OldObjectSize(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldObjectSize is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldObjectSize requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldObjectSize: %w", err)
	}
	return oldValue.ObjectSize, nil
}

// AddObjectSize adds i to the "object_size" field.
func (m *ObjectMutation) AddObjectSize(i int64) {
	if m.addobject_size != nil {
		*m.addobject_size += i
	} else {
		m.addobject_size = &i
	}
}

// AddedObjectSize returns the value that was added to the "object_size" field in this mutation.
func (m *ObjectMutation) AddedObjectSize() (r int64, exists bool) {
	v := m.addobject_size
	if v == nil {
		return
	}
	return *v, true
}

// ResetObjectSize resets all changes to the "object_size" field.
func (m *ObjectMutation) ResetObjectSize() {
	m.object_size = nil
	m.addobject_size = nil
}

// SetObjectSha256 sets the "object_sha256" field.
func (m *ObjectMutation) SetObjectSha256(s string) {
	m.object_sha256 = &s
}

// ObjectSha256 returns the value of the "object_sha256" field in the mutation.
func (m *ObjectMutation) ObjectSha256() (r string, exists bool) {
	v := m.object_sha256
	if v == nil {
		return
	}
	return *v, true
}

// OldObjectSha256 returns the old "object_sha256" field's value of the Object entity.
// If the Object object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ObjectMutation) OldObjectSha256(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldObjectSha256 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldObjectSha256 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldObjectSha256: %w", err)
	}
	return oldValue.ObjectSha256, nil
}

// ResetObjectSha256 resets all changes to the "object_sha256" field.
func (m *ObjectMutation) ResetObjectSha256() {
	m.object_sha256 = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *ObjectMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ObjectMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Object entity.
// If the Object object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ObjectMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ObjectMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *ObjectMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *ObjectMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Object entity.
// If the Object object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ObjectMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *ObjectMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the ObjectMutation builder.
func (m *ObjectMutation) Where(ps ...predicate.Object) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ObjectMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Object).
func (m *ObjectMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ObjectMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.object_name != nil {
		fields = append(fields, object.FieldObjectName)
	}
	if m.content_type != nil {
		fields = append(fields, object.FieldContentType)
	}
	if m.object_location != nil {
		fields = append(fields, object.FieldObjectLocation)
	}
	if m.object_size != nil {
		fields = append(fields, object.FieldObjectSize)
	}
	if m.object_sha256 != nil {
		fields = append(fields, object.FieldObjectSha256)
	}
	if m.created_at != nil {
		fields = append(fields, object.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, object.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ObjectMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case object.FieldObjectName:
		return m.ObjectName()
	case object.FieldContentType:
		return m.ContentType()
	case object.FieldObjectLocation:
		return m.ObjectLocation()
	case object.FieldObjectSize:
		return m.ObjectSize()
	case object.FieldObjectSha256:
		return m.ObjectSha256()
	case object.FieldCreatedAt:
		return m.CreatedAt()
	case object.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ObjectMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case object.FieldObjectName:
		return m.OldObjectName(ctx)
	case object.FieldContentType:
		return m.OldContentType(ctx)
	case object.FieldObjectLocation:
		return m.OldObjectLocation(ctx)
	case object.FieldObjectSize:
		return m.OldObjectSize(ctx)
	case object.FieldObjectSha256:
		return m.OldObjectSha256(ctx)
	case object.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case object.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Object field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ObjectMutation) SetField(name string, value ent.Value) error {
	switch name {
	case object.FieldObjectName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetObjectName(v)
		return nil
	case object.FieldContentType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContentType(v)
		return nil
	case object.FieldObjectLocation:
		v, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetObjectLocation(v)
		return nil
	case object.FieldObjectSize:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetObjectSize(v)
		return nil
	case object.FieldObjectSha256:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetObjectSha256(v)
		return nil
	case object.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case object.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Object field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ObjectMutation) AddedFields() []string {
	var fields []string
	if m.addobject_location != nil {
		fields = append(fields, object.FieldObjectLocation)
	}
	if m.addobject_size != nil {
		fields = append(fields, object.FieldObjectSize)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ObjectMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case object.FieldObjectLocation:
		return m.AddedObjectLocation()
	case object.FieldObjectSize:
		return m.AddedObjectSize()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ObjectMutation) AddField(name string, value ent.Value) error {
	switch name {
	case object.FieldObjectLocation:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddObjectLocation(v)
		return nil
	case object.FieldObjectSize:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddObjectSize(v)
		return nil
	}
	return fmt.Errorf("unknown Object numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ObjectMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ObjectMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ObjectMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Object nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ObjectMutation) ResetField(name string) error {
	switch name {
	case object.FieldObjectName:
		m.ResetObjectName()
		return nil
	case object.FieldContentType:
		m.ResetContentType()
		return nil
	case object.FieldObjectLocation:
		m.ResetObjectLocation()
		return nil
	case object.FieldObjectSize:
		m.ResetObjectSize()
		return nil
	case object.FieldObjectSha256:
		m.ResetObjectSha256()
		return nil
	case object.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case object.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Object field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ObjectMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ObjectMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ObjectMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ObjectMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ObjectMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ObjectMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ObjectMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Object unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ObjectMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Object edge %s", name)
}
