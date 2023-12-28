package config

// environment variable, will be .upperCase()
const (
	MySqlUsernameConfig string = "db_username"
	MySqlPasswordConfig string = "db_password"
	MySqlAddrConfig     string = "db_address"
	MySqlDatabaseName   string = "db_name"

	HttpHandlerPort string = "http_handler_port"
	Workdir         string = "workdir"
)
