package dent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DynamicUpdate is the builder for updating Dynamic entities.
type DynamicUpdate struct {
	mutation *DynamicMutation
}

// Where appends a list predicates to the DynamicUpdate builder.
func (du *DynamicUpdate) Where(ps ...PredicateDynamic) *DynamicUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetCreatorID sets the "creator_id" field.
func (du *DynamicUpdate) SetValue(name string, val interface{}) *DynamicUpdate {
	du.mutation.SetValue(name, val)
	return du
}

// AddCreatorID adds i to the "creator_id" field.
func (du *DynamicUpdate) AddValue(name string, i int) *DynamicUpdate {
	du.mutation.AddValue(name, i)
	return du
}

// ClearCreatorID clears the value of the "creator_id" field.
func (du *DynamicUpdate) ClearValue(name string) *DynamicUpdate {
	du.mutation.ClearValue(name)
	return du
}

// Mutation returns the DynamicMutation object of the builder.
func (du *DynamicUpdate) Mutation() *DynamicMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DynamicUpdate) Save(ctx context.Context) (int, error) {
	if err := du.defaults(); err != nil {
		return 0, err
	}
	return du.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DynamicUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DynamicUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DynamicUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DynamicUpdate) defaults() error {
	return nil
}

func (du *DynamicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   du.mutation.table.Name,
			Columns: du.mutation.table.GetColumns(),
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	for k, v := range du.mutation.data {
		col, _ := du.mutation.table.Column(k)
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   col.Type,
			Value:  v,
			Column: k,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.mutation.table.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{du.mutation.table.Name}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DynamicUpdateOne is the builder for updating a single Dynamic entity.
type DynamicUpdateOne struct {
	fields   []string
	mutation *DynamicMutation
}

// SetCreatorID sets the "creator_id" field.
func (duo *DynamicUpdateOne) SetValue(name string, val interface{}) *DynamicUpdateOne {
	duo.mutation.SetValue(name, val)
	return duo
}

// AddCreatorID adds i to the "creator_id" field.
func (duo *DynamicUpdateOne) AddValue(name string, i int) *DynamicUpdateOne {
	duo.mutation.AddValue(name, i)
	return duo
}

// ClearCreatorID clears the value of the "creator_id" field.
func (duo *DynamicUpdateOne) ClearValue(name string) *DynamicUpdateOne {
	duo.mutation.ClearValue(name)
	return duo
}

// Mutation returns the DynamicMutation object of the builder.
func (duo *DynamicUpdateOne) Mutation() *DynamicMutation {
	return duo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DynamicUpdateOne) Select(field string, fields ...string) *DynamicUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dynamic entity.
func (duo *DynamicUpdateOne) Save(ctx context.Context) (*Dynamic, error) {
	if err := duo.defaults(); err != nil {
		return nil, err
	}

	return duo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DynamicUpdateOne) SaveX(ctx context.Context) *Dynamic {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DynamicUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DynamicUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DynamicUpdateOne) defaults() error {
	return nil
}

func (duo *DynamicUpdateOne) sqlSave(ctx context.Context) (_node *Dynamic, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   duo.mutation.table.Name,
			Columns: duo.mutation.table.GetColumns(),
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: FieldID, err: errors.New(`ent: missing "Dynamic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, FieldID)
		for _, f := range fields {
			if !duo.mutation.table.HasColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}

	for k, v := range duo.mutation.data {
		col, _ := duo.mutation.table.Column(k)
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   col.Type,
			Value:  v,
			Column: k,
		})
	}

	_node = &Dynamic{table: duo.mutation.table}
	_node.ID = id
	_node.Row = duo.mutation.data
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.mutation.table.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{duo.mutation.table.Name}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
