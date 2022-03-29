package utils

import (
	"github.com/hashicorp/go-hclog"
	"github.com/lib/pq"
	"github.com/spf13/viper"
)

// Configurations wraps all the config variables required by the auth service
type Configurations struct {
	DBHost string
	DBName string
	DBUser string
	DBPass string
	DBPort string
	DBConn string

	AccessTokenPrivateKeyPath  string
	AccessTokenPublicKeyPath   string
	RefreshTokenPrivateKeyPath string
	RefreshTokenPublicKeyPath  string
	JwtExpiration              int // in minutes

	MailVerifCodeExpiration int // in hours
	PassResetCodeExpiration int // in minutes
	MailVerifTemplate       string
	PassResetTemplate       string

	SendGridApiKey string

	MailGunDomain        string
	MailGunPrivateAPIKey string
}

// NewConfigurations returns a new Configuration object
func NewConfigurations(logger hclog.Logger) *Configurations {

	viper.AutomaticEnv()

	dbURL := viper.GetString("DATABASE_URL")
	conn, _ := pq.ParseURL(dbURL)
	logger.Debug("found database url in env, connection string is formed by parsing it")
	logger.Debug("db connection string", conn)

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "bookite")
	viper.SetDefault("DB_USER", "vignesh")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("ACCESS_TOKEN_PRIVATE_KEY_PATH", "./access-private.pem")
	viper.SetDefault("ACCESS_TOKEN_PUBLIC_KEY_PATH", "./access-public.pem")
	viper.SetDefault("REFRESH_TOKEN_PRIVATE_KEY_PATH", "./refresh-private.pem")
	viper.SetDefault("REFRESH_TOKEN_PUBLIC_KEY_PATH", "./refresh-public.pem")
	viper.SetDefault("JWT_EXPIRATION", 30)
	viper.SetDefault("MAIL_VERIFICATION_CODE_EXPIRATION", 24)
	viper.SetDefault("PASSWORD_RESET_CODE_EXPIRATION", 15)
	viper.SetDefault("MAIL_VERIFICATION_TEMPLATE", "d-5ecbea6e38764af3b703daf03f139b48")
	viper.SetDefault("PASSWORD_RESET_TEMPLATE", "d-3fc222d11809441abaa8ed459bb44319")

	configs := &Configurations{
		DBHost:                     viper.GetString("DB_HOST"),
		DBName:                     viper.GetString("DB_NAME"),
		DBUser:                     viper.GetString("DB_USER"),
		DBPass:                     viper.GetString("DB_PASSWORD"),
		DBPort:                     viper.GetString("DB_PORT"),
		DBConn:                     conn,
		JwtExpiration:              viper.GetInt("JWT_EXPIRATION"),
		AccessTokenPrivateKeyPath:  viper.GetString("ACCESS_TOKEN_PRIVATE_KEY_PATH"),
		AccessTokenPublicKeyPath:   viper.GetString("ACCESS_TOKEN_PUBLIC_KEY_PATH"),
		RefreshTokenPrivateKeyPath: viper.GetString("REFRESH_TOKEN_PRIVATE_KEY_PATH"),
		RefreshTokenPublicKeyPath:  viper.GetString("REFRESH_TOKEN_PUBLIC_KEY_PATH"),

		MailVerifCodeExpiration: viper.GetInt("MAIL_VERIFICATION_CODE_EXPIRATION"),
		PassResetCodeExpiration: viper.GetInt("PASSWORD_RESET_CODE_EXPIRATION"),
		MailVerifTemplate:       viper.GetString("MAIL_VERIFICATION_TEMPLATE"),
		PassResetTemplate:       viper.GetString("PASSWORD_RESET_TEMPLATE"),

		SendGridApiKey:       viper.GetString("SENDGRID_API_KEY"),
		MailGunDomain:        viper.GetString("MAILGUN_DOMAIN"),
		MailGunPrivateAPIKey: viper.GetString("MAILGUN_PRIVATE_API_KEY"),
	}

	logger.Debug("db host", configs.DBHost)
	logger.Debug("db name", configs.DBName)
	logger.Debug("db port", configs.DBPort)
	logger.Debug("jwt expiration", configs.JwtExpiration)

	return configs
}
