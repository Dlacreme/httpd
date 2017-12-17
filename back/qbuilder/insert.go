package qbuilder

import (
	"fmt"
	"strconv"
	"strings"
)

// InsertInto allows to push value into a table
func InsertInto(table string, fields []string) string {
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES ", table, strings.Join(fields, ","))
}

// AddInto will add values to an insert into statement
func AddInto(q string, params []string) string {
	if strings.Count(q, "(") > 1 {
		return fmt.Sprintf(", (%s) ", strings.Join(params, ","))
	}
	return fmt.Sprintf("(%s) ", strings.Join(params, ","))
}

// FormatStr will add ' to any param
func FormatStr(item string) string {
	s := []string{"'", item, "'"}
	return strings.Join(s, "")
}

func ToStringOrNull(i int) string {
	if i < 1 {
		return "NULL"
	}
	return strconv.Itoa(i)
}

func BoolAsIntString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
