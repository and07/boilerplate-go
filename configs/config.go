package configs

// Configs ...
type Configs struct {
	Port            string `env:"PORT" envDefault:"8080"`
	PortDebug       string `env:"PORT_DEBUG" envDefault:"8888"`
	PortGrpc        string `env:"PORT_GRPC" envDefault:"8842"`
	PortGrpcGateway string `env:"PORT_GRPC_GATEWAY" envDefault:"8843"`
}
