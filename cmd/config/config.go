package config

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Dialect:  "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "secret1234",
		Name:     "library",
		Charset:  "utf8",
	}
}

type RouteConfig struct {
	Port    string
	Version string
}

func GetRouterConfig() *RouteConfig {
	return &RouteConfig{
		Port:    ":8080",
		Version: "/v1",
	}
}

type Config struct {
	DB    *DBConfig
	Route *RouteConfig
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "secret1234",
			Name:     "library",
			Charset:  "utf8",
		},
		Route: &RouteConfig{
			Port:    ":8080",
			Version: "/v1",
		},
	}
}
