package common

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConn(db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:password@tcp(10.100.104.21:3307)/dbname?charset=utf8mb4")
	return
}

func GetResultRow(rows *sql.Rows) map[string]string {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	record := make(map[string]string)
	for rows.Next() {
		rows.Scan(scanArgs...)
		for i, v := range values {
			if v != nil {
				record[columns[i]] = string(v.([]byte))
			}
		}
	}
	return record
}

// func GetResultRows(rows *sql.Rows) map[string]string {
// columns, _ := rows.Columns()
// values := make([][]byte, len(columns))
// scans := make([]interface{}, len(columns))
// for k, _ := range values {
// 	scans[k] = &values[k.([]byte)]
// }
// //i :=0
// record := make(map[string]string)

// for rows.Next() {
// 	rows.Scan(scans...)
// 	for i, v := range values {
// 		if v != nil {
// 			record[columns[i]] = string(v)
// 		}
// 	}
// }
// return record
// }
