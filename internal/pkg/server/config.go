package server

import (
	"net"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Config is a structure used to configure a GenericAPIServer.
// Its members are sorted roughly in order of importance for composers.
type Config struct {
	SecureServing *SecureServingInfo
	Mode          string
	Id            int64
}

// SecureServingInfo holds configuration of the TLS server.
type SecureServingInfo struct {
	BindAddress string
	BindPort    int
}

// Address join host IP address and host port number into a address string, like: 0.0.0.0:8443.
func (s *SecureServingInfo) Address() string {
	return net.JoinHostPort(s.BindAddress, strconv.Itoa(s.BindPort))
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *Config {
	return &Config{
		Mode: gin.ReleaseMode,
		Id:   1,
	}
}

// New returns a new instance of GenericAPIServer from the given configs.
func (c CompletedConfig) New() (*GenericAPIServer, error) {
	// setMode before gin.New()
	gin.SetMode(c.Mode)

	s := &GenericAPIServer{
		SecureServingInfo: c.SecureServing,
		Engine:            gin.New(),
	}

	initGenericAPIServer(s)

	return s, nil
}

// CompletedConfig is the completed configuration for GenericAPIServer.
type CompletedConfig struct {
	*Config
}

// Complete fills in any fields not set that are required to have valid data and can be derived
// from other fields. If you're going to `ApplyOptions`, do that first. It's mutating the receiver.
func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}
