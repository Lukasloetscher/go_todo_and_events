package postgresql

//Initialise_postgresql() creates a new connection
//for the moment the connection will be hard_coded, but we will change this to a dynamic connection soon
func Initialise_postgresql() (_ *SQL_Connection, err error) {
	var m SQL_Connection
	m.Setting = SQL_connection_settings{}
	m.Setting.Close_connection_after_each_use = true
	m.connection_data = SQL_connection_data{}
	m.connection_data.Database_name = "udemy_testing"
	m.connection_data.Host = "localhost"
	m.connection_data.Password = "admin"
	m.connection_data.Port = 5432
	m.connection_data.Username = "postgres"

	return &m, nil
}
