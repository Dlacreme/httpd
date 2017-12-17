package wenum

import (
	"database/sql"
	"fmt"

	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

// genum build enum from DB table

type item struct {
	Id    int    `db:"id"`
	Label string `db:"label"`
}

// BuildEnum :
func BuildEnum(db wdb.Connection, tableName string, idField string, labelField string) (map[int]string, *werror.Error) {
	r := []item{}
	err := db.Select(&r, fmt.Sprintf(`
		SELECT %s AS 'id', %s AS 'label' FROM %s	
	`, idField, labelField, tableName))

	if err != nil && err != sql.ErrNoRows {
		return nil, werror.New(500, fmt.Sprintf("Cannot build enum %s : %s", tableName, err.Error()))
	}

	res := make(map[int]string, len(r))
	for i := range r {
		res[r[i].Id] = r[i].Label
	}

	return res, nil
}
