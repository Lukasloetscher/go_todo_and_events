package postgresql

type SQL_Connection struct {
	connection_data SQL_connection_data
	Setting         SQL_connection_settings
}

type SQL_connection_data struct {
	Host          string
	Username      string
	Password      string
	Port          int64
	Database_name string
}

type SQL_connection_settings struct {
	Close_connection_after_each_use bool //only implemented true yet
}
