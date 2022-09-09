package postgresql

type Config struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func NewConfig(host, port, database, username, password string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		Database: database,
		Username: username,
		Password: password,
	}
}
