package dent

import (
	"entgo.io/ent/dialect/sql"
)

// Dynamic is the predicate function for dynamic builders.
type PredicateDynamic func(*sql.Selector)
