package utils

// ServerConfig defines the configuration for the server
type ServerConfig struct {
	Host          string
	Port          string
	URISchema     string
	Version       string
	Mode          string
	SessionSecret string
	Database      DatabaseConfig
	GraphQL       GraphQLConfig
	JWT           JWTConfig
}

//JWTConfig defines the options for JWT tokens
type JWTConfig struct {
	Secret    string
	Algorithm string
}

// GraphQLConfig defines the configuration for the GQL Server
type GraphQLConfig struct {
	Path                string
	PlaygroundPath      string
	IsPlaygroundEnabled bool
}

// DatabaseConfig defines the configuration for the DB config
type DatabaseConfig struct {
	Dialect      string
	Host         string
	Port         string
	Password     string
	Username     string
	DatabaseName string
	SSLMode      string
	SeedDB       bool
	LogMode      bool
	AutoMigrate  bool
}

// ListenEndpoint builds the endpoint string (host + port)
func (s *ServerConfig) ListenEndpoint() string {
	if s.Port == "80" {
		return s.Host
	}
	return s.Host + ":" + s.Port
}

// VersionedEndpoint builds the endpoint string (host + port + version)
func (s *ServerConfig) VersionedEndpoint(path string) string {
	return "/" + s.Version + path
}

// SchemaVersionedEndpoint builds the schema endpoint string (schema + host + port + version)
func (s *ServerConfig) SchemaVersionedEndpoint(path string) string {
	if s.Port == "80" {
		return s.URISchema + s.Host + "/" + s.Version + path
	}
	return s.URISchema + s.Host + ":" + s.Port + "/" + s.Version + path
}
