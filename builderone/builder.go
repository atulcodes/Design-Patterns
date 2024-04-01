package main

import "fmt"

type ServerConfig struct {
	Port           int
	Timeout        int
	MaxConnections int
	Logging        bool
}

type ServerConfigBuilder struct {
	*ServerConfig
}

func NewServerConfigBuilder() *ServerConfigBuilder {
	return &ServerConfigBuilder{
		ServerConfig: &ServerConfig{
			Port:           8080, // Default Port
			Timeout:        30,   // Default timeout in seconds
			MaxConnections: 100,  // Default no of concurrent connections
			Logging:        true, // Logging enabled by default
		},
	}
}

func (s *ServerConfigBuilder) SetPort(port int) *ServerConfigBuilder {
	s.Port = port
	return s
}

func (s *ServerConfigBuilder) SetTimeout(timeout int) *ServerConfigBuilder {
	s.Timeout = timeout
	return s
}

func (s *ServerConfigBuilder) SetMaxConnections(maxConnections int) *ServerConfigBuilder {
	s.MaxConnections = maxConnections
	return s
}

func (s *ServerConfigBuilder) EnableLogging() *ServerConfigBuilder {
	s.Logging = true
	return s
}

func (s *ServerConfigBuilder) DisableLogging() *ServerConfigBuilder {
	s.Logging = false
	return s
}

func (s *ServerConfigBuilder) Build() *ServerConfig {
	return s.ServerConfig
}

func main() {
	serverConfig := NewServerConfigBuilder().
		SetPort(8080).
		SetTimeout(10).
		SetMaxConnections(50).
		Build()

	fmt.Println(*serverConfig)
	fmt.Printf("%v", *serverConfig)
}
