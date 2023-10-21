package models

type Settings struct {
	AppParams      Params
	PostgresParams PostgresSettings
}

type Params struct {
	ServerURL     string
	ServerName    string
	AppVersion    string
	PortRun       string
	LogInfo       string
	LogError      string
	LogDebug      string
	LogWarning    string
	LogMaxSize    int
	LogMaxBackups int
	LogMaxAge     int
	LogCompress   bool
	SecretKey     string
	TokenTTL      int
}

type PostgresSettings struct {
	User     string
	Password string
	Server   string
	Port     string
	Database string
	SSLMode  string
}
