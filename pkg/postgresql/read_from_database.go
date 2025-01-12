package postgresql

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

//just as a note:
// https://pkg.go.dev/github.com/jackc/pgx/v5

// Read_table allows for simple reads from table.
// this function does not allow for complicated stuff
func (m *SQL_Connection) Read_table(table_name string, schema string, col []string, cond map[string]string) (_ []map[string]string, err error) {
	conn, err := m.establish_connecion()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	sql_string := "SELECT "
	if len(col) != 0 {
		for i, column := range col {
			if i == len(col)-1 {
				sql_string += column + " "
			} else {
				sql_string += column + ", "
			}
		}
	} else { //when there are no columns specified we read all of them
		sql_string += "* "
	}
	sql_string += "FROM " + schema + "." + table_name
	if len(cond) != 0 {
		sql_string += " Where "
		i := 0
		for key, value := range cond {
			if i == len(cond)-1 {
				sql_string += key + " " + value
			} else {
				sql_string += key + " " + value + " AND "
			}
			i++
		}
	}
	rows, err := conn.Query(context.Background(), sql_string)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	defer rows.Close()
	s := make([]any, len(col))
	for i := range s {
		s[i] = new(interface{})
	}
	var data []map[string]string
	i := 0
	_, err = pgx.ForEachRow(rows, s, func() error {
		data = append(data, make(map[string]string))
		for ii, val := range s {
			//value_string := (*(val.(*interface{}))).(string) //type assertion see https://go.dev/ref/spec#Type_assertions
			//this is not the cleanest -> i should code this specific for tables i'm goingto use.
			value_string := fmt.Sprintf("%v", *(val.(*interface{})))
			data[i][col[ii]] = value_string
		}
		i += 1
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return data, nil
}
