package config

type Configuration struct {
	HTTPServerConfiguration HTTPServerConfiguration
}

type HTTPServerConfiguration struct {
	Port string
}
