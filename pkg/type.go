package pkg

type (
	ConfigRes struct {
		Port  int
		Env   EnvMode
		DbUrl string
		Mail  MailConfig
	}

	MailConfig struct {
		Host     string
		Port     int
		User     string
		Password string
	}
)

type EnvMode string

const (
	EnvProduction  EnvMode = "production"
	EnvDevelopment EnvMode = "development"
)
