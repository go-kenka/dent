package dent

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

func NewTable(name string) *schema.Table {

	t := schema.NewTable(name)

	defaultColumns := []*schema.Column{
		{Name: FieldID, Type: field.TypeInt, Increment: true},
	}

	// 添加主键
	t.AddPrimary(defaultColumns[0])
	return t
}
