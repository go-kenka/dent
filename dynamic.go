package dent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Dynamic is the model entity for the Dynamic schema.
type Dynamic struct {
	table *Table `json:"-"`
	// ID of the ent.
	ID  int                  `json:"id,omitempty"`
	Row map[string]ent.Value `json:"row,omitempty"`
	// edges
	Edges DynamicEdges `json:"edges"`
}

// DatastoreEdges holds the relations/edges for other nodes in the graph.
type DynamicEdges struct {
	SingleMap map[string]*Dynamic   `json:"single_map,omitempty"`
	ListMap   map[string][]*Dynamic `json:"list_map,omitempty"`
}

func (d *DynamicEdges) Get(key string) *Dynamic {
	return d.SingleMap[key]
}

func (d *DynamicEdges) List(key string) []*Dynamic {
	return d.ListMap[key]
}

// scanValues returns the types for scanning values from sql.Rows.
func (d *Dynamic) scanValues(columns []string) ([]interface{}, error) {
	return d.table.scanValues(columns)
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dynamic fields.
func (d *Dynamic) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i, v := range columns {

		switch value := values[i].(type) {
		case *sql.NullInt64:
			if v == FieldID {
				d.ID = int(value.Int64)
			} else {
				d.Row[v] = value.Int64
			}
		case *sql.NullBool:
			d.Row[v] = value.Bool
		case *sql.NullString:
			d.Row[v] = value.String
		case *sql.NullFloat64:
			d.Row[v] = value.Float64
		case *sql.NullTime:
			d.Row[v] = value.Time
		case *[]byte:
			d.Row[v] = string(*value)
		default:
			d.Row[v] = value
		}
	}
	return nil
}

// // Update returns a builder for updating this Dynamic.
// // Note that you need to call Dynamic.Unwrap() before calling this method if this Dynamic
// // was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dynamic) Update() *DUpdateOne {
	return d.table.UpdateOne(d)
}

// // Update returns a builder for updating this Dynamic.
// // Note that you need to call Dynamic.Unwrap() before calling this method if this Dynamic
// // was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dynamic) Delete() *DDeleteOne {
	return d.table.DeleteOneID(d.ID)
}

// Unwrap unwraps the Dynamic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dynamic) Unwrap() *Dynamic {
	_tx, ok := d.table.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dynamic is not a transactional entity")
	}
	d.table.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dynamic) String() string {
	var builder strings.Builder
	builder.WriteString("Dynamic(")
	i := 0
	len := len(d.Row) - 1
	for k, v := range d.Row {
		if i < len {
			builder.WriteString(fmt.Sprintf("%s=%v, ", k, v))
		} else {
			builder.WriteString(fmt.Sprintf("%s=%v", k, v))
		}
		i++
	}
	builder.WriteByte(')')

	return builder.String()
}
