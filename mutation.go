package dent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeDynamic = "Dynamic"
)

// DynamicMutation represents an operation that mutates the Dynamic nodes in the graph.
type DynamicMutation struct {
	table         *Table
	op            Op
	typ           string
	id            *int
	data          map[string]ent.Value
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Dynamic, error)
	predicates    []PredicateDynamic
}

var _ ent.Mutation = (*DynamicMutation)(nil)

// dynamicOption allows management of the mutation configuration using functional options.
type dynamicOption func(*DynamicMutation)

// newDynamicMutation creates new mutation for the Dynamic entity.
func newDynamicMutation(t *Table, op Op, opts ...dynamicOption) *DynamicMutation {
	m := &DynamicMutation{
		table:         t,
		op:            op,
		typ:           TypeDynamic,
		data:          make(map[string]ent.Value),
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withID sets the ID field of the mutation.
func withID(id int) dynamicOption {
	return func(m *DynamicMutation) {
		var (
			err   error
			once  sync.Once
			value *Dynamic
		)
		m.oldValue = func(ctx context.Context) (*Dynamic, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Table(m.table.Name).Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEntity sets the old Dynamic of the mutation.
func withEntity(node *Dynamic) dynamicOption {
	return func(m *DynamicMutation) {
		m.oldValue = func(context.Context) (*Dynamic, error) {
			return node, nil
		}

		m.id = &node.ID
		m.data = node.Row
	}
}

// withField sets the old Dynamic of the mutation.
func withField(fields ...*schema.Column) dynamicOption {
	return func(m *DynamicMutation) {
		for _, c := range fields {
			if c.Name == FieldID {
				continue
			}
			var v interface{}
			switch c.Type {
			case field.TypeBool:
				v = new(sql.NullBool)
			case field.TypeTime:
				v = new(sql.NullTime)
			case field.TypeJSON, field.TypeBytes:
				v = new([]byte)
			case field.TypeUUID, field.TypeString:
				v = new(sql.NullString)
			case field.TypeEnum, field.TypeInt8, field.TypeInt16, field.TypeInt32, field.TypeInt, field.TypeInt64, field.TypeUint8, field.TypeUint16, field.TypeUint32, field.TypeUint, field.TypeUint64:
				v = new(sql.NullInt64)
			case field.TypeFloat32, field.TypeFloat64:
				v = new(sql.NullFloat64)
			case field.TypeOther:
				v = new(interface{})
			}

			m.data[c.Name] = v
		}
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DynamicMutation) Client() *Client {
	client := &Client{config: m.table.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DynamicMutation) Tx() (*Tx, error) {
	if _, ok := m.table.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.table.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Access entities.
func (m *DynamicMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DynamicMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetValue sets the value of the id field. Note that this
// operation is only accepted on creation of Dynamic entities.
func (m *DynamicMutation) SetValue(field string, val interface{}) {
	if m.table.HasColumn(field) {
		m.data[field] = &val
	}
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DynamicMutation) Value(field string) (id interface{}, exists bool) {
	if _, ok := m.data[field]; !ok {
		return
	}

	v := m.data[field]
	return v, true
}

// OldValue returns the old "field" field's value of the Dynamic entity.
// If the Dynamic object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DynamicMutation) OldValue(ctx context.Context, field string) (v interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatorID is only allowed on UpdateOne operations")
	}
	if _, ok := m.data[field]; !ok || m.oldValue == nil {
		return v, errors.New("OldCreatorID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatorID: %w", err)
	}
	return oldValue.Row[field], nil
}

// ClearEditorID clears the value of the "field" field.
func (m *DynamicMutation) ClearValue(field string) {
	delete(m.data, field)
	m.clearedFields[field] = struct{}{}
}

// ResetName resets all changes to the "field" field.
func (m *DynamicMutation) ResetValue(field string) {
	delete(m.data, field)
}

// AddedValue returns the value that was added to the "field" field in this mutation.
func (m *DynamicMutation) AddedValue(field string) (ent.Value, bool) {
	if col, ok := m.table.Column(field); ok {
		if col.Type.Integer() {
			v, ok := m.data[field]
			if !ok {
				return nil, false
			}
			return v, true
		}
	}

	return nil, false
}

// AddValue returns the value that was added to the "field" field in this mutation.
func (m *DynamicMutation) AddValue(name string, i int) {
	if col, ok := m.table.Column(name); ok {
		if col.Type.Integer() {
			v, ok := m.data[name]
			if !ok {
				m.data[name] = i
			}

			m.data[name] = v.(int) + i
		}
	}
}

// Where appends a list predicates to the DynamicMutation builder.
func (m *DynamicMutation) Where(ps ...PredicateDynamic) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *DynamicMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Dynamic).
func (m *DynamicMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DynamicMutation) Fields() []string {
	return m.table.GetColumns()
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DynamicMutation) Field(name string) (ent.Value, bool) {
	if m.table.HasColumn(name) {
		return m.data[name], true
	}

	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DynamicMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return m.OldValue(ctx, name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DynamicMutation) SetField(name string, value ent.Value) error {
	m.data[name] = value
	return nil
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DynamicMutation) AddedFields() []string {
	var fields []string
	for _, c := range m.table.Columns {
		if c.Type.Integer() {

			fields = append(fields, c.Name)
		}
	}

	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DynamicMutation) AddedField(name string) (ent.Value, bool) {
	return m.AddedValue(name)
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DynamicMutation) AddField(name string, value ent.Value) error {

	v, ok := value.(int)
	if !ok {
		return fmt.Errorf("unexpected type %T for field %s", value, name)
	}

	m.AddValue(name, v)
	return nil
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DynamicMutation) ClearedFields() []string {
	var fields []string
	// if m.FieldCleared(dynamic.FieldCreatorID) {
	// 	fields = append(fields, dynamic.FieldCreatorID)
	// }
	// if m.FieldCleared(dynamic.FieldEditorID) {
	// 	fields = append(fields, dynamic.FieldEditorID)
	// }
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DynamicMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DynamicMutation) ClearField(name string) error {
	m.ClearValue(name)
	return nil
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DynamicMutation) ResetField(name string) error {
	m.ResetValue(name)
	return nil
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DynamicMutation) AddedEdges() []string {
	edges := make([]string, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DynamicMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DynamicMutation) RemovedEdges() []string {
	edges := make([]string, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DynamicMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DynamicMutation) ClearedEdges() []string {
	edges := make([]string, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DynamicMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DynamicMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Dynamic unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DynamicMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Dynamic edge %s", name)
}
