## dent - 一个强大的Go语言实体框架

dent是一个简单而又功能强大的Go语言实体框架，dent易于构建和维护应用程序与大数据模型。

- **图就是代码** - 将任何数据库表建模为Go对象。
- **轻松地遍历任何图形** - 可以轻松地运行查询、聚合和遍历任何图形结构。
- **静态类型和显式API** - 使用代码生成静态类型和显式API，查询数据更加便捷。
- **多存储驱动程序** - 支持MySQL, PostgreSQL, SQLite 和 Gremlin。
- **可扩展** - 简单地扩展和使用Go模板自定义。

## 快速安装
```console
go get github.com/go-kenka/dent
```
## 示例
```go
client, err := dent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
	if err != nil {
		t.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	table := dent.NewTable("test_aaa")
	table.AddColumn(&schema.Column{
		Name: "field1",
		Type: field.TypeInt,
	})
	table.AddColumn(&schema.Column{
		Name: "field2",
		Type: field.TypeInt,
	})

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), table); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	dclient := client.Dynamic
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

	result, err := dclient.Table("test_aaa").Create().SetValue("field3", 1).SetValue("field2", 2).Save(context.Background())
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

```

## 声明
dent使用Apache 2.0协议授权，可以在[LICENSE文件](LICENSE)中找到。
