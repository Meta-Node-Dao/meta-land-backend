package config

// Ceres configuration from the .toml file

// Minio configuration for ceres to access the minio server.
var Minio *MinioConfig

// Github Oauth configuration.
var Github *GithubOauth

// Google Oauth configuration.
var Google *GoogleOauth

// Facebook Oauth configuration.
var Facebook *FacebookOauth

// JWT configuration.
var JWT *JWTConfig

// Seq configuration.
var Seq *Sequence

// Mysql configuration.
var Mysql *MysqlConfig

// Aws configuration.
var Aws *AwsConfig

// Eth configuration.
var Eth *EthConfig

// Log configuration.
var Log *LogConfig

// MinioConfig from the .toml file
type MinioConfig struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
}

// JWTConfig from the .toml file
type JWTConfig struct {
	Expired int
	Secret  string
}

// GithubOauth from the .toml file
type GithubOauth struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

// GoogleOauth from the .toml file
type GoogleOauth struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

// FacebookOauth from the .toml file
type FacebookOauth struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

// Sequence from .toml file
type Sequence struct {
	Epoch int64
}

// MysqlConfig from .toml file
type MysqlConfig struct {
	ConnMaxLifetime int
	Debug           bool
	Dsn             string
	Level           string
	MaxIdleConns    int
	MaxOpenConns    int
}

// AwsConfig from .toml file
type AwsConfig struct {
	AccessKey    string
	AccessSecret string
	Bucket       string
	EndPoint     string
	Region       string
	MaxSize      int64
}

// EthConfig from .toml file
type EthConfig struct {
	Epoch                  int64
	EndPoint               string
	InfuraKey              string
	StartupContractAddress string
	ChainID                uint64
	WSSEndPoint            string
	RPCEndPoint            string
	EthereumChainID        uint64
}

type LogConfig struct {
	Debug bool
	Level string
}
