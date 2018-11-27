package config

// DatabaseConfig stores database connection options
type DatabaseConfig struct {
	Type         string `default:"mysql"`
	Host         string `default:"localhost"`
	Port         int    `default:"3306"`
	User         string `default:"gravitee"`
	Password     string `default:"gravitee"`
	DatabaseName string `default:"gravitee"`
	MaxIdleConns int    `default:"5"`
	MaxOpenConns int    `default:"5"`
}

// OauthConfig stores oauth service configuration options
type OauthConfig struct {
	AccessTokenLifetime  int `default:"3600"`    // default to 1 hour
	RefreshTokenLifetime int `default:"1209600"` // default to 14 days
	AuthCodeLifetime     int `default:"3600"`    // default to 1 hour
}

// SessionConfig stores session configuration for the web app
type SessionConfig struct {
	Secret string `default:"test_secret"`
	Path   string `default:"/"`
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge int `default:"604800"`
	// When you tag a cookie with the HttpOnly flag, it tells the browser that
	// this particular cookie should only be accessed by the server.
	// Any attempt to access the cookie from client script is strictly forbidden.
	HTTPOnly bool `default:"True"`
}

// Config stores all configuration options
type Config struct {
	Database      DatabaseConfig
	Oauth         OauthConfig
	Session       SessionConfig
	ServerPort    int  `default:"8080"`
	IsDevelopment bool `default:"True"`
}
