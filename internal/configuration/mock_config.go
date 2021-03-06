package configuration

import "time"

// Mock represents configuration required to run jenkins-proxy
type Mock struct {
	PostgresHost              string
	PostgresPort              int
	PostgresUser              string
	PostgresDatabase          string
	PostgresPassword          string
	PostgresSSLMode           string
	PostgresConnectionTimeout int
	PostgresConnectionMaxIdle int
	PostgresConnectionMaxOpen int
	IdlerURL                  string
	AuthURL                   string
	TenantURL                 string
	WitURL                    string
	AuthToken                 string
	RedirectURL               string
	IndexPath                 string
	MaxRequestRetry           int
	DebugMode                 bool
	HTTPSEnabled              bool
	GatewayTimeout            time.Duration
	AllowedOrigins            []string
	Clusters                  map[string]string
}

// NewMock creates an instance of configuration
func NewMock() Mock {
	c := Mock{}

	// set some defaults which can be overridden as needed
	c.PostgresHost = "localhost"
	c.PostgresPort = 5432
	c.PostgresUser = "postgres"
	c.PostgresPassword = "postgres"
	c.PostgresDatabase = "tenant"
	c.PostgresSSLMode = "disable"
	c.PostgresConnectionMaxIdle = -1
	c.PostgresConnectionMaxOpen = -1
	c.PostgresConnectionTimeout = 5

	c.RedirectURL = "https://localhost:8443/"
	c.IndexPath = "static/html/index.html"
	c.GatewayTimeout = 25 * time.Second
	c.AllowedOrigins = []string{"https://*openshift.io", "https://localhost:*", "http://localhost:*"}

	return c
}

// GetPostgresHost return hardcoded postgres host from test configuration.
func (c *Mock) GetPostgresHost() string {
	return c.PostgresHost
}

// GetPostgresPort return hardcoded postgres porn from test configuration.
func (c *Mock) GetPostgresPort() int {
	return c.PostgresPort
}

// GetPostgresUser return hardcoded postgres user from test configuration.
func (c *Mock) GetPostgresUser() string {
	return c.PostgresUser
}

// GetPostgresDatabase return hardcoded postgres database name from test configuration.
func (c *Mock) GetPostgresDatabase() string {
	return c.PostgresDatabase
}

// GetPostgresPassword return hardcoded postgres password from test configuration.
func (c *Mock) GetPostgresPassword() string {
	return c.PostgresPassword
}

// GetPostgresSSLMode return hardcoded postgres ssl mode (ssl enabled or disabled) from test configuration.
func (c *Mock) GetPostgresSSLMode() string {
	return c.PostgresSSLMode
}

// GetPostgresConnectionTimeout return hardcoded postgres connection timeout from test configuration.
func (c *Mock) GetPostgresConnectionTimeout() int {
	return c.PostgresConnectionTimeout
}

// GetPostgresConnectionMaxIdle return hardcoded number of connections that should be kept alive in the database connection pool at any given time from test configuration.
// Here it is set to -1, which represents no restrictions/default behavior.
func (c *Mock) GetPostgresConnectionMaxIdle() int {
	return c.PostgresConnectionMaxIdle
}

// GetPostgresConnectionMaxOpen return hardcoded max number of open connections that should be open in the database connection pool from test configuration.
// Here it is set to -1 represents no restrictions/default behavior.
func (c *Mock) GetPostgresConnectionMaxOpen() int {
	return c.PostgresConnectionMaxOpen
}

// GetIdlerURL return hardcoded url of idler service from test configuration.
func (c *Mock) GetIdlerURL() string {
	return c.IdlerURL
}

// GetAuthURL return hardcoded url of the fabric8-auth service from test configuration.
func (c *Mock) GetAuthURL() string {
	return c.AuthURL
}

// GetTenantURL return hardcoded url of tenant service from test configuration.
func (c *Mock) GetTenantURL() string {
	return c.TenantURL
}

// GetWitURL return hardcoded url of wit service from test configuration.
func (c *Mock) GetWitURL() string {
	return c.WitURL
}

// GetAuthToken return hardcoded value of auth token from test configuration.
func (c *Mock) GetAuthToken() string {
	return c.AuthToken
}

// GetRedirectURL return hardcoded redirect url to be passed to the auth service from test configuration.
func (c *Mock) GetRedirectURL() string {
	return c.RedirectURL
}

// GetIndexPath return hardcoded path to loading page template from test configuration.
func (c *Mock) GetIndexPath() string {
	return c.IndexPath
}

// GetMaxRequestRetry return hardcoded number of retries for webhook request forwarding from test configuration.
func (c *Mock) GetMaxRequestRetry() int {
	return c.MaxRequestRetry
}

// GetDebugMode return hardcoded debug mode from test configuration.
func (c *Mock) GetDebugMode() bool {
	return c.DebugMode
}

// GetHTTPSEnabled return hardcoded http-enabled from test configuration.
func (c *Mock) GetHTTPSEnabled() bool {
	return c.HTTPSEnabled
}

// GetGatewayTimeout returns hardcoded gateway timeout from test configuration.
func (c *Mock) GetGatewayTimeout() time.Duration {
	return c.GatewayTimeout
}

// GetAllowedOrigins returns hardcoded allowed origins
func (c *Mock) GetAllowedOrigins() []string {
	return c.AllowedOrigins
}

func (c *Mock) String() string {
	return "mockConfig"
}
