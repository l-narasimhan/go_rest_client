package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgresql",
			Host:     "172.18.0.2",
			Port:     5555,
			Username: "postgres",
			Password: "postgres",
			Name:     "twitter_postgres",
			Charset:  "utf8",
		},
	}
}


