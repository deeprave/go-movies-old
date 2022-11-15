package helpers

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Port:    9000,
		Host:    "localhost",
		Env:     "dev",
		Version: "1.0",
	}
}
