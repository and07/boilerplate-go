package configs

// Configs ...
type Configs struct {
	Port            string `env:"PORT" envDefault:"8080"`
	PortDebug       string `env:"PORT_DEBUG" envDefault:"8888"`
	PortGrpc        string `env:"PORT_GRPC" envDefault:"8842"`
	PortGrpcGateway string `env:"PORT_GRPC_GATEWAY" envDefault:"8843"`
}

// GoogleAuthConfigs ...
type GoogleAuthConfigs struct {
	GoogleKey          string `env:"GOOGLE_KEY,notEmpty" envDefault:"328290909614-jar104iq8508k7n2lhrj453up6oieo4j.apps.googleusercontent.com"`
	GoogleSecret       string `env:"GOOGLE_SECRET,notEmpty" envDefault:"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"`
	GoogleAuthCallback string `env:"GOOGLE_AUTH_CALLBACK,notEmpty" envDefault:"http://localhost:8080/auth/google/callback"`
}

type S3Configs struct {
	Region    string `env:"S3_REGION,notEmpty"`
	Host      string `env:"S3_HOST",notEmpty`
	AccessKey string `env:"S3_ACCESS_KEY,notEmpty"`
	SecretKey string `env:"S3_SECRET_KEY,notEmpty"`
}
