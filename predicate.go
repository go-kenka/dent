package dent

import (
	"entgo.io/ent/dialect/sql"
)

// Predicate is the predicate function for dynamic builders.
type Predicate func(*sql.Selector)
