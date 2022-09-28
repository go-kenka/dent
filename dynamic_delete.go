package dent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DynamicDelete is the builder for deleting a Dynamic entity.
type DynamicDelete struct {
	mutation *DynamicMutation
}

// Where appends a list predicates to the DynamicDelete builder.
func (dd *DynamicDelete) Where(ps ...PredicateDynamic) *DynamicDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DynamicDelete) Exec(ctx context.Context) (int, error) {
	return dd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DynamicDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DynamicDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: dd.mutation.table.Name,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: FieldID,
			},
		},
	}
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.mutation.table.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DynamicDeleteOne is the builder for deleting a single Dynamic entity.
type DynamicDeleteOne struct {
	dd *DynamicDelete
}

// Exec executes the deletion query.
func (ddo *DynamicDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ddo.dd.mutation.table.Name}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DynamicDeleteOne) ExecX(ctx context.Context) {
	ddo.dd.ExecX(ctx)
}
