package apiserver

import (
	"fmt"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/config"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store/mysql"
	genericoptions "github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/options"
	genericapiserver "github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type apiServer struct {
	gRPCAPIServer    *grpcAPIServer
	genericAPIServer *genericapiserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

func (srv *apiServer) PrepareRun() preparedAPIServer {
	initRouter(srv.genericAPIServer.Engine)
	return preparedAPIServer{srv}
}

func (psrv preparedAPIServer) Run() error {
	go psrv.gRPCAPIServer.Run()

	return psrv.genericAPIServer.Run()
}

// ExtraConfig defines extra configuration for the iam-apiserver.
type ExtraConfig struct {
	Addr         string
	MaxMsgSize   int
	mysqlOptions *genericoptions.MySQLOptions
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}
	extraConfig, err := buildExtraConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}
	extraServer, err := extraConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	server := &apiServer{
		genericAPIServer: genericServer,
		gRPCAPIServer:    extraServer,
	}

	genericapiserver.SnowflakeSrv, _ = genericapiserver.NewSnowflake(genericConfig.Id)

	return server, nil
}

// New create a grpcAPIServer instance.
func (c *completedExtraConfig) New() (*grpcAPIServer, error) {
	grpcServer := grpc.NewServer()

	storeIns, _ := mysql.GetMySQLFactory(c.mysqlOptions)

	store.SetClient(storeIns)

	reflection.Register(grpcServer)

	return &grpcAPIServer{grpcServer, c.Addr}, nil
}

// Complete fills in any fields not set that are required to have valid data and can be derived from other fields.
func (c *ExtraConfig) Complete() *completedExtraConfig {
	if c.Addr == "" {
		c.Addr = "127.0.0.1:8081"
	}

	return &completedExtraConfig{c}
}

type completedExtraConfig struct {
	*ExtraConfig
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	if lastErr = cfg.SecureOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	return
}

// nolint: unparam
func buildExtraConfig(cfg *config.Config) (*ExtraConfig, error) {
	return &ExtraConfig{
		Addr:         fmt.Sprintf("%s:%d", cfg.GRPCOptions.BindAddress, cfg.GRPCOptions.BindPort),
		MaxMsgSize:   cfg.GRPCOptions.MaxMsgSize,
		mysqlOptions: cfg.MySQLOptions,
	}, nil
}
