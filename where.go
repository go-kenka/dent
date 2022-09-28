package dent

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

var FieldID = "id"

// ID filters vertices based on their ID field.
func ID(id int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IntEQ applies the EQ predicate on the ID field.
func IntEQ(fleld string, val int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(fleld), val))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IntNEQ(fleld string, val int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(fleld), val))
	})
}

// IntIn applies the In predicate on the ID field.
func IntIn(fleld string, vals ...int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		v := make([]interface{}, len(vals))
		for i := range v {
			v[i] = vals[i]
		}
		s.Where(sql.In(s.C(fleld), v...))
	})
}

// IntNotIn applies the NotIn predicate on the ID field.
func IntNotIn(fleld string, vals ...int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		v := make([]interface{}, len(vals))
		for i := range v {
			v[i] = vals[i]
		}
		s.Where(sql.NotIn(s.C(fleld), v...))
	})
}

// IntGT applies the GT predicate on the ID field.
func IntGT(fleld string, val int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(fleld), val))
	})
}

// IntGTE applies the GTE predicate on the ID field.
func IntGTE(fleld string, val int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(fleld), val))
	})
}

// IntLT applies the LT predicate on the ID field.
func IntLT(fleld string, val int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(fleld), val))
	})
}

// IntLTE applies the LTE predicate on the ID field.
func IntLTE(fleld string, val int) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(fleld), val))
	})
}

// FloatEQ applies the EQ predicate on the ID field.
func FloatEQ(fleld string, val float64) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(fleld), val))
	})
}

// FloatNEQ applies the NEQ predicate on the ID field.
func FloatNEQ(fleld string, val float64) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(fleld), val))
	})
}

// FloatIn applies the In predicate on the ID field.
func FloatIn(fleld string, vals ...float64) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		v := make([]interface{}, len(vals))
		for i := range v {
			v[i] = vals[i]
		}
		s.Where(sql.In(s.C(fleld), v...))
	})
}

// FloatNotIn applies the NotIn predicate on the ID field.
func FloatNotIn(fleld string, vals ...float64) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		v := make([]interface{}, len(vals))
		for i := range v {
			v[i] = vals[i]
		}
		s.Where(sql.NotIn(s.C(fleld), v...))
	})
}

// FloatGT applies the GT predicate on the ID field.
func FloatGT(fleld string, val float64) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(fleld), val))
	})
}

// FloatGTE applies the GTE predicate on the ID field.
func FloatGTE(fleld string, val float64) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(fleld), val))
	})
}

// FloatLT applies the LT predicate on the ID field.
func FloatLT(fleld string, val float64) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(fleld), val))
	})
}

// FloatLTE applies the LTE predicate on the ID field.
func FloatLTE(fleld string, val float64) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(fleld), val))
	})
}

// TimeEQ applies the EQ predicate on the "created_at" field.
func TimeEQ(fleld string, v time.Time) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(fleld), v))
	})
}

// TimeNEQ applies the NEQ predicate on the "created_at" field.
func TimeNEQ(fleld string, v time.Time) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(fleld), v))
	})
}

// TimeIn applies the In predicate on the "created_at" field.
func TimeIn(fleld string, vs ...time.Time) PredicateDynamic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.In(s.C(fleld), v...))
	})
}

// TimeNotIn applies the NotIn predicate on the "created_at" field.
func TimeNotIn(fleld string, vs ...time.Time) PredicateDynamic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(fleld), v...))
	})
}

// TimeGT applies the GT predicate on the "created_at" field.
func TimeGT(fleld string, v time.Time) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(fleld), v))
	})
}

// TimeGTE applies the GTE predicate on the "created_at" field.
func TimeGTE(fleld string, v time.Time) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(fleld), v))
	})
}

// TimeLT applies the LT predicate on the "created_at" field.
func TimeLT(fleld string, v time.Time) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(fleld), v))
	})
}

// TimeLTE applies the LTE predicate on the "created_at" field.
func TimeLTE(fleld string, v time.Time) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(fleld), v))
	})
}

// StringEQ applies the EQ predicate on the "name" field.
func StringEQ(fleld string, v string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(fleld), v))
	})
}

// StringNEQ applies the NEQ predicate on the "name" field.
func StringNEQ(fleld string, v string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(fleld), v))
	})
}

// StringIn applies the In predicate on the "name" field.
func StringIn(fleld string, vs ...string) PredicateDynamic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.In(s.C(fleld), v...))
	})
}

// StringNotIn applies the NotIn predicate on the "name" field.
func StringNotIn(fleld string, vs ...string) PredicateDynamic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(fleld), v...))
	})
}

// StringGT applies the GT predicate on the "name" field.
func StringGT(fleld string, v string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(fleld), v))
	})
}

// StringGTE applies the GTE predicate on the "name" field.
func StringGTE(fleld string, v string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(fleld), v))
	})
}

// StringLT applies the LT predicate on the "name" field.
func StringLT(fleld string, v string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(fleld), v))
	})
}

// StringLTE applies the LTE predicate on the "name" field.
func StringLTE(fleld string, v string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(fleld), v))
	})
}

// StringContains applies the Contains predicate on the "name" field.
func StringContains(fleld string, v string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(fleld), v))
	})
}

// StringHasPrefix applies the HasPrefix predicate on the "name" field.
func StringHasPrefix(fleld string, v string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(fleld), v))
	})
}

// StringHasSuffix applies the HasSuffix predicate on the "name" field.
func StringHasSuffix(fleld string, v string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(fleld), v))
	})
}

// StringEqualFold applies the EqualFold predicate on the "name" field.
func StringEqualFold(fleld string, v string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(fleld), v))
	})
}

// StringContainsFold applies the ContainsFold predicate on the "name" field.
func StringContainsFold(fleld string, v string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(fleld), v))
	})
}

// BoolEQ applies the EQ predicate on the "name" field.
func BoolEQ(fleld string, v bool) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(fleld), v))
	})
}

// BoolNEQ applies the NEQ predicate on the "name" field.
func BoolNEQ(fleld string, v bool) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(fleld), v))
	})
}

func HasTable(rel sqlgraph.Rel, inverse bool, table string, columns ...string) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(s.TableName(), FieldID),
			sqlgraph.To(table, FieldID),
			sqlgraph.Edge(rel, inverse, table, columns...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...PredicateDynamic) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...PredicateDynamic) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
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
func Not(p PredicateDynamic) PredicateDynamic {
	return PredicateDynamic(func(s *sql.Selector) {
		p(s.Not())
	})
}
