package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Object holds the schema definition for the Object entity.
type Object struct {
	ent.Schema
}

// Fields of the Object.
func (Object) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),                                          // 对象ID
		field.String("object_name"),                                          // 对象名称（对应原始文件名）
		field.String("content_type"),                                         // 对象类型
		field.Uint8("object_location"),                                       // 存储位置 1 对象存储 2 文本存储
		field.Int64("object_size"),                                           // 对象大小
		field.String("object_sha256"),                                        // 对象Sha256的值
		field.Time("created_at").Default(time.Now),                           // 对象创建时间
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now), // 对象更新时间
	}
}

// Edges of the Object.
func (Object) Edges() []ent.Edge {
	return nil
}

func (Object) Indexes() []ent.Index {
	return []ent.Index{index.Fields("id")} // 给对象ID创建一个索引
}
