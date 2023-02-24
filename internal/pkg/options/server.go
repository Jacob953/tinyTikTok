package options

import (
	"github.com/spf13/pflag"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/server"
)

// ServerRunOptions contains the options while running a generic api server.
type ServerRunOptions struct {
	Mode string `json:"mode" mapstructure:"mode"`
	Id   int64  `json:"id"   mapstructure:"id"`
}

// NewServerRunOptions creates a new ServerRunOptions object with default parameters.
func NewServerRunOptions() *ServerRunOptions {
	defaults := server.NewConfig()

	return &ServerRunOptions{
		Mode: defaults.Mode,
		Id:   defaults.Id,
	}
}

// Validate checks validation of ServerRunOptions.
func (s *ServerRunOptions) Validate() []error {
	errors := []error{}

	return errors
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet.
func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Mode, "server.mode", s.Mode, ""+
		"Start the server in a specified server mode. Supported server mode: debug, test, release.")
	fs.Int64Var(&s.Id, "server.id", s.Id, ""+
		"Start the server in a specified server id. Supported server id is unique.")
}

// ApplyTo applies the run options to the method receiver and returns self.
func (s *ServerRunOptions) ApplyTo(c *server.Config) error {
	c.Mode = s.Mode
	c.Id = s.Id

	return nil
}
