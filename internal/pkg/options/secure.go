package options

import (
	"fmt"

	"github.com/spf13/pflag"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/server"
)

// SecureOptions contains configuration items related to HTTPS server startup.
type SecureOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
	Required    bool
}

// NewSecureOptions creates a SecureOptions object with default parameters.
func NewSecureOptions() *SecureOptions {
	return &SecureOptions{
		BindAddress: "0.0.0.0",
		BindPort:    8443,
		Required:    true,
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (s *SecureOptions) ApplyTo(c *server.Config) error {
	c.SecureServing = &server.SecureServingInfo{
		BindAddress: s.BindAddress,
		BindPort:    s.BindPort,
	}

	return nil
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (s *SecureOptions) Validate() []error {
	if s == nil {
		return nil
	}

	errors := []error{}

	if s.Required && s.BindPort < 1 || s.BindPort > 65535 {
		errors = append(
			errors,
			fmt.Errorf(
				"--secure.bind-port %v must be between 1 and 65535, inclusive. It cannot be turned off with 0",
				s.BindPort,
			),
		)
	} else if s.BindPort < 0 || s.BindPort > 65535 {
		errors = append(errors, fmt.Errorf("--secure.bind-port %v must be between 0 and 65535, inclusive. 0 for turning off secure port", s.BindPort))
	}

	return errors
}

// AddFlags adds flags related to HTTPS server for a specific APIServer to the
// specified FlagSet.
func (s *SecureOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddress, "secure.bind-address", s.BindAddress, ""+
		"The IP address on which to listen for the --secure.bind-port port. The "+
		"associated interface(s) must be reachable by the rest of the engine, and by CLI/web "+
		"clients. If blank, all interfaces will be used (0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")
	desc := "The port on which to serve HTTPS with authentication and authorization."
	if s.Required {
		desc += " It cannot be switched off with 0."
	} else {
		desc += " If 0, don't serve HTTPS at all."
	}
	fs.IntVar(&s.BindPort, "secure.bind-port", s.BindPort, desc)
}

// Complete fills in any fields not set that are required to have valid data.
func (s *SecureOptions) Complete() error {
	if s == nil || s.BindPort == 0 {
		return nil
	}

	return nil
}
