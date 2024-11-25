package envs

// Хранение данных значений ENV
var ServerEnvs Envs

// Структура для хранения значений ENV
type Envs struct {
	SONGBOOK_PORT     string
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_PORT     string
	POSTGRES_NAME     string
	POSTGRES_HOST     string
	POSTGRES_USE_SSL  string
}

// / Инициализация значений ENV
func LoadEnvs() error {
	ServerEnvs.POSTGRES_USER = "evans"
	ServerEnvs.POSTGRES_PASSWORD = "evans"
	ServerEnvs.POSTGRES_PORT = "3010"
	ServerEnvs.POSTGRES_NAME = "postgres"
	ServerEnvs.POSTGRES_HOST = "localhost"
	ServerEnvs.POSTGRES_USE_SSL = "disable"

	ServerEnvs.SONGBOOK_PORT = "3000"

	// для запуска через докер, все, что выше нужно добавить в коммиты, все ниже - коммиты убрать и ипортировать пакет OS

	// ServerEnvs.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	// ServerEnvs.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	// ServerEnvs.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	// ServerEnvs.POSTGRES_NAME = os.Getenv("POSTGRES_NAME")
	// ServerEnvs.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	// ServerEnvs.POSTGRES_USE_SSL = os.Getenv("POSTGRES_USE_SSL")

	// ServerEnvs.SONGBOOK_PORT = os.Getenv("SONGBOOK_PORT")

	return nil
}
