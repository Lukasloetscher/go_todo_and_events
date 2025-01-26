package postgresql

//Initialise_postgresql() creates a new connection
//for the moment the connection will be hard_coded, but we will change this to a dynamic connection soon
//Note that the "sql server" currently runs local in a contaimner using docker desktop.
//TODO as soon as i use a real database, the password should no longer be pushed to github, for obvios reasons.
//TODO look up enviromental variables for this. (or how this is doen in profesional projects)
func Initialise_postgresql() (_ *SQL_Connection, err error) {
	var m SQL_Connection
	m.Setting = SQL_connection_settings{}
	m.Setting.Close_connection_after_each_use = true
	m.connection_data = SQL_connection_data{}
	m.connection_data.Database_name = "test_website_go"
	m.connection_data.Host = "localhost"
	m.connection_data.Password = "password"
	m.connection_data.Port = 5432
	m.connection_data.Username = "postgres"

	return &m, nil
}
