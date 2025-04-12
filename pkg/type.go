package pkg

type (
	ConfigRes struct {
		Port  int
		Env   string
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
