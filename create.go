package dent

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DCreate is the builder for creating a Dynamic entity.
type DCreate struct {
	mutation *DMutation
}

// SetCreatorID sets the "creator_id" field.
func (du *DCreate) SetValue(name string, val interface{}) *DCreate {
	du.mutation.SetValue(name, val)
	return du
}

// Mutation returns the DMutation object of the builder.
func (dc *DCreate) Mutation() *DMutation {
	return dc.mutation
}

// Save creates the Dynamic in the database.
func (dc *DCreate) Save(ctx context.Context) (*Dynamic, error) {
	var (
		err  error
		node *Dynamic
	)
	if err := dc.defaults(); err != nil {
		return nil, err
	}
	if err = dc.check(); err != nil {
		return nil, err
	}
	node, err = dc.sqlSave(ctx)

	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DCreate) SaveX(ctx context.Context) *Dynamic {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DCreate) defaults() error {
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dc *DCreate) check() error {
	return nil
}

func (dc *DCreate) sqlSave(ctx context.Context) (*Dynamic, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.mutation.table.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}

	_node.table = dc.mutation.table
	return _node, nil
}

func (dc *DCreate) createSpec() (*Dynamic, *sqlgraph.CreateSpec) {
	var (
		_node = &Dynamic{table: dc.mutation.table, Row: make(map[string]ent.Value)}
		_spec = &sqlgraph.CreateSpec{
			Table: dc.mutation.table.Name,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	for k, v := range dc.mutation.data {
		col, _ := dc.mutation.table.Column(k)
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   col.Type,
			Value:  v,
			Column: k,
		})
	}

	_node.Row = dc.mutation.data

	return _node, _spec
}

// DCreateBulk is the builder for creating many Dynamic entities in bulk.
type DCreateBulk struct {
	config
	builders []*DCreate
}

// Save creates the Dynamic entities in the database.
func (dcb *DCreateBulk) Save(ctx context.Context) ([]*Dynamic, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dynamic, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DCreateBulk) SaveX(ctx context.Context) []*Dynamic {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
