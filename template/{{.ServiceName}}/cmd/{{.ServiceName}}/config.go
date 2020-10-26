package main

// Config ...
type Config struct {
	Port      string `env:"PORT"`
	PortDebug string `env:"PORT_DEBUG"`
	PortGRPC  string `env:"PORT_GRPC"`
}
