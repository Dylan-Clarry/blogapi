package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect string
	Host 	string
	Port 	int
	User 	string
	Pass 	string
	DBName 	string
}

func CreateConfig() *Config {
	return &Config {
		DB: &DBConfig {
			Dialect: 	"postgres",
			Host: 		"localhost",
			Port: 		5432,
			User: 		"testuser",
			Pass: 		"password",
			DBName: 	"userdb",
		},
	}
}