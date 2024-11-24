package envs

// Хранение данных значений ENV
var ServerEnvs Envs

// Структура для хранения значений ENV
type Envs struct {
	SERVER_PORT       string
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_PORT     string
	POSTGRES_NAME     string
	POSTGRES_HOST     string
	POSTGRES_USE_SSL  string
}

// / Инициализация значений ENV
func LoadEnvs() error {
	ServerEnvs.SERVER_PORT = "3000"
	ServerEnvs.POSTGRES_USER = "evans"
	ServerEnvs.POSTGRES_PASSWORD = "evans"
	ServerEnvs.POSTGRES_PORT = "3010"
	ServerEnvs.POSTGRES_NAME = "postgres"
	ServerEnvs.POSTGRES_HOST = "localhost"
	ServerEnvs.POSTGRES_USE_SSL = "disable"

	return nil
}
