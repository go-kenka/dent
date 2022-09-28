package dent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)

var (
	// WithGlobalUniqueID sets the universal ids options to the migration.
	// If this option is enabled, ent migration will allocate a 1<<32 range
	// for the ids of each entity (table).
	// Note that this option cannot be applied on tables that already exist.
	WithGlobalUniqueID = schema.WithGlobalUniqueID
	// WithDropColumn sets the drop column option to the migration.
	// If this option is enabled, ent migration will drop old columns
	// that were used for both fields and edges. This defaults to false.
	WithDropColumn = schema.WithDropColumn
	// WithDropIndex sets the drop index option to the migration.
	// If this option is enabled, ent migration will drop old indexes
	// that were defined in the schema. This defaults to false.
	// Note that unique constraints are defined using `UNIQUE INDEX`,
	// and therefore, it's recommended to enable this option to get more
	// flexibility in the schema changes.
	WithDropIndex = schema.WithDropIndex
	// WithForeignKeys enables creating foreign-key in schema DDL. This defaults to true.
	WithForeignKeys = schema.WithForeignKeys
)

// Schema is the API for creating, migrating and dropping a schema.
type Schema struct {
	drv dialect.Driver
}

// NewSchema creates a new schema client.
func NewSchema(drv dialect.Driver) *Schema { return &Schema{drv: drv} }

// Create creates all schema resources.
func (s *Schema) Create(ctx context.Context, table *schema.Table, opts ...schema.MigrateOption) error {
	return Create(ctx, s, []*schema.Table{table}, opts...)
}

// Create creates all table resources using the given schema driver.
func Create(ctx context.Context, s *Schema, tables []*schema.Table, opts ...schema.MigrateOption) error {
	migrate, err := schema.NewMigrate(s.drv, opts...)
	if err != nil {
		return fmt.Errorf("ent/migrate: %w", err)
	}
	return migrate.Create(ctx, tables...)
}
