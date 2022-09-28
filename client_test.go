package dent

import (
	"context"
	"testing"

	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
	_ "github.com/go-sql-driver/mysql"
)

func TestNewClient(t *testing.T) {
	dclient, err := Open("mysql", "root:pass@tcp(localhost:3306)/test")
	if err != nil {
		t.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer dclient.Close()

	table := NewTable("test_aaa")
	table.AddColumn(&schema.Column{
		Name: "field1",
		Type: field.TypeInt,
	})
	table.AddColumn(&schema.Column{
		Name: "field2",
		Type: field.TypeInt,
	})

	// Run the auto migration tool.
	if err := dclient.Schema.Create(context.Background(), table); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	dclient.AddTable(table)
	{
		table := NewTable("user")
		table.AddColumn(&schema.Column{
			Name: "username",
			Type: field.TypeString,
		})
		table.AddColumn(&schema.Column{
			Name: "creator_id",
			Type: field.TypeInt,
		})
		dclient.AddTable(table)
	}

	result, err := dclient.Table("test_aaa").Create().SetValue("field1", 1).SetValue("field2", 2).Save(context.Background())
	if err != nil {
		t.Fatalf("failed create value: %v", err)
	}

	t.Logf("create result: %v", result.Row)

	result, err = result.Update().SetValue("field1", 4).Save(context.Background())
	if err != nil {
		t.Fatalf("failed update value: %v", err)
	}

	t.Logf("update result: %v", result.Row)

	err = result.Delete().Exec(context.Background())
	if err != nil {
		t.Fatalf("failed delete value: %v", err)
	}

	result, err = dclient.Table("user").Query().Where(And(IntEQ("id", 1))).WithListData("test_aaa", "test_list", "field1").First(context.Background())
	if err != nil {
		t.Fatalf("failed create value: %v", err)
	}
	t.Logf("query result: %v", result)

	count, err := dclient.Table("test_aaa").Query().Where(And(IntEQ("field1", 1))).Count(context.Background())
	if err != nil {
		t.Fatalf("failed create value: %v", err)
	}
	t.Logf("count result: %v", count)
}
