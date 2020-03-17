package gql

import (
	"fmt"
)

type query struct {
	//where   where
	fields []*field
}

type aux struct {
	Offset     int
	Limit      int
	DistinctOn []string
}

func Query(fields ...*field) *query {
	return &query{
		fields: fields,
	}
}

func (q *query) String() string {
	var result string
	for _, v := range q.fields {
		result = result + v.string()
	}
	return fmt.Sprintf("{%s}", result)
}

func (f *field) Offset(count int) *field {
	f.aux.Offset = count
	return f
}

func (f *field) Limit(count int) *field {
	f.aux.Limit = count
	return f
}

func (f *field) Distinct(field ...string) *field {
	for _, v := range field {
		f.aux.DistinctOn = append(f.aux.DistinctOn, v)
	}
	return f
}

func (f *field) Asc(field string) *field {
	f.orderBy = append(f.orderBy, orderBy{
		field:      field,
		descending: false,
	})
	return f
}

func (f *field) Desc(field string) *field {
	f.orderBy = append(f.orderBy, orderBy{
		field:      field,
		descending: true,
	})
	return f
}
func (f *field) string() string {
	var conStr string
	for _, condition := range f.where.condition {
		conStr = conStr + condition.String() + ","
	}
	conStr = conStr[:len(conStr)-1] + " "
	single := fmt.Sprintf("\n  %s ", f.table)
	if len(conStr) > 0 {
		single = single + "( where: { " + conStr + "} )"
	}
	single = single + "{\n"
	for _, column := range f.columns {
		single = single + fmt.Sprintf("   %s\n", column)
	}
	if f.subField != nil {
		single = single + f.subField.string()
	}

	single = single + "  }\n"

	return single
}

/////

type orderBy struct {
	field      string
	descending bool
}

type field struct {
	table    string
	columns  []string
	where    where
	aux      aux
	orderBy  []orderBy
	subField *field
}

func Field(table string, columns ...string) (t *field) {
	t = &field{
		table: table,
	}
	for _, v := range columns {
		t.columns = append(t.columns, v)
	}
	return
}

func (f *field) WithSubField(subField *field) *field {
	f.subField = subField
	return f
}

func (f *field) Where(c *Condition) *field {
	f.where.condition = []*Condition{c}
	return f
}

func (f *field) And(c *Condition) *field {
	f.where.condition = append(f.where.condition, &Condition{
		field:    c.field,
		operator: c.operator,
		value:    c.value,
		combine:  and,
	})
	return f
}

func (f *field) Or(c *Condition) *field {
	f.where.condition = append(f.where.condition, &Condition{
		field:    c.field,
		operator: c.operator,
		value:    c.value,
		combine:  or,
	})
	return f
}

type Operator int32

const equalTo Operator = 1
const notEqualTo Operator = 2
const greaterThan Operator = 3
const lessThan Operator = 4
const greaterThanOrEqualTo Operator = 5
const lessThanOrEqualTo Operator = 6
const inList Operator = 7
const notInList Operator = 8
const isNull Operator = 9
const like Operator = 10
const similar Operator = 11
const contains Operator = 12
const hasKey Operator = 13
const and Operator = 14
const or Operator = 15

var operator_symbol = map[Operator]string{
	1:  "_eq",
	2:  "_neq",
	3:  "_gt",
	4:  "_lt",
	5:  "_gte",
	6:  "_lte",
	7:  "_in",
	8:  "_nin",
	9:  "_is_null",
	10: "_like",
	11: "_similar",
	12: "_contains",
	13: "_has_key",
	14: "_and",
	15: "_or",
}

type where struct {
	condition []*Condition
	subWhere  *where
}

type Condition struct {
	field    string
	operator Operator
	value    string
	combine  Operator
}

func (c *Condition) String() string {
	return fmt.Sprintf("%s: {%s: %s}", c.field, operator_symbol[c.operator], c.value)
}

//func NewWhere(c *Condition) *where {
//	return &where{
//		condition: []*Condition{
//			c,
//		},
//	}
//}

func GT(field string, value string) *Condition {
	return &Condition{
		field:    field,
		operator: greaterThan,
		value:    value,
	}
}

func EQ(field string, value string) *Condition {
	return &Condition{
		field:    field,
		operator: equalTo,
		value:    value,
	}
}

func NEQ(field string, value string) *Condition {
	return &Condition{
		field:    field,
		operator: notEqualTo,
		value:    value,
	}
}

func LT(field string, value string) *Condition {
	return &Condition{
		field:    field,
		operator: lessThan,
		value:    value,
	}
}
