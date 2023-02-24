package options

import (
	"github.com/marmotedu/iam/pkg/log"

	genericoptions "github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/options"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/server"
	"github.com/CSU-Apple-Lab/tinyTikTok/pkg/app"
)

func NewOptions() *Options {
	opt := Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		GRPCOptions:             genericoptions.NewGRPCOptions(),
		MySQLOptions:            genericoptions.NewMySQLOptions(),
		SecureOptions:           genericoptions.NewSecureOptions(),
		Log:                     log.NewOptions(),
	}

	return &opt
}

// Options runs an iam api server.
type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions `json:"server" mapstructure:"server"`
	GRPCOptions             *genericoptions.GRPCOptions      `json:"grpc"   mapstructure:"grpc"`
	MySQLOptions            *genericoptions.MySQLOptions     `json:"mysql"  mapstructure:"mysql"`
	SecureOptions           *genericoptions.SecureOptions    `json:"secure" mapstructure:"secure"`
	Log                     *log.Options                     `json:"log"    mapstructure:"log"`
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *server.Config) error {
	return nil
}

func (o Options) Flags() (fss app.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.SecureOptions.AddFlags(fss.FlagSet("secure"))
	o.Log.AddFlags(fss.FlagSet("logs"))
	return fss
}

// Complete set default Options.
func (o *Options) Complete() error {
	return o.SecureOptions.Complete()
}
