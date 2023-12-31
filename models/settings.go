package models

type Settings struct {
	AppParams      Params
	PostgresParams PostgresSettings
	SMTPParams     SMTPParams
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
	JWTIssuer     string
}

type PostgresSettings struct {
	User     string
	Password string
	Server   string
	Port     string
	Database string
	SSLMode  string
}

type SMTPParams struct {
	Host     string
	Username string
	Port     string
	Password string
}
