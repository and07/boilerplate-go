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
	GoogleKey          string `env:"GOOGLE_KEY" envDefault:"328290909614-jar104iq8508k7n2lhrj453up6oieo4j.apps.googleusercontent.com"`
	GoogleSecret       string `env:"GOOGLE_SECRET" envDefault:"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"`
	GoogleAuthCallback string `env:"GOOGLE_AUTH_CALLBACK" envDefault:"http://localhost:8080/auth/google/callback"`
}
