package postgresql

import (
	"context"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

// establish_connecion() is a private method, which opens a connectiona and saves it to the struct.
// there is potential an issue with openign a connection like this, due to the fact, that now we could open multiple connections, due to async stuff
// hence this fucntion may only be used, for when we only want one conenction, even when it is used by different
// Accoring to stackoverflow, it is fine to leave thje connection open, and sue the sam conection for all calls.
func (m *SQL_Connection) establish_connecion() (conn *pgx.Conn, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connection_string := "postgres://" + m.connection_data.Username + ":" + m.connection_data.Password + "@" + m.connection_data.Host + ":" + strconv.FormatInt(m.connection_data.Port, 10) + "/" + m.connection_data.Database_name
	conn, err = pgx.Connect(ctx, connection_string)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
