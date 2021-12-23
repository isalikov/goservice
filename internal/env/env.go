package env

type Config struct {
	Env			string	`env:"ENV,default=dev"`
	Port     	int    	`env:"PORT,default=50000"`
	Release     string  `env:"RELEASE,default=0.0.0"`
	SentryDsn	string	`env:"SENTRY_DSN,default=https://..."`
}
