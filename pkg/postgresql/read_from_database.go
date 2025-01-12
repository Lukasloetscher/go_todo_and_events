package postgresql

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

// Read_table allows for simple reads from table.
// this function does not allow for complicated stuff
func (m *SQL_Connection) Read_table(table_name string, schema string, col []string, cond map[string]string) (data []map[string]string, err error) {
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
	fmt.Println(sql_string)
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
	fmt.Println("start of for each")
	re, err := pgx.ForEachRow(rows, s, func() error {
		fmt.Println(*(s[0].(*interface{})), *(s[1].(*interface{})), *(s[2].(*interface{})))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(re)
	fmt.Println("end of for each")

	return nil, nil
}
