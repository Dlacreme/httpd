package wdb

import "database/sql"

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// GetIdsFromResult will build a []int from sql.Result
func GetIdsFromResult(sqlRes sql.Result) []int {

	lastIdTmp, _ := sqlRes.LastInsertId()
	numberTmp, _ := sqlRes.RowsAffected()

	lastId := int(lastIdTmp)
	number := int(numberTmp)
	r := []int{}

	for i := lastId; i > lastId-number; i-- {
		r = append(r, i)
	}

	return r
}

func ParseDateLayout() string {
	return "2006-01-02T15:04:05Z"
}
