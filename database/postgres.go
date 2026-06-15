package database

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func Connect(cfg Config) error {
	return nil
}

func Close() error {
	return nil
}
