package client

import (
	"context"
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/pomerium/enterprise-client-go/pb"
)

// AuthorizationHeader turns a service account token into the correct format for the Authorization header
func AuthorizationHeader(token string) string {
	return fmt.Sprintf("Pomerium %s", token)
}

// PomeriumAuthCredentials implements grpc.PerRPCCredentials for communication with Pomerium Enterprise
type PomeriumAuthCredentials struct {
	metadata map[string]string
}

// NewPomeriumAuthCredentials returns a new PomeriumAuthCredentials for the provided service account authToken
func NewPomeriumAuthCredentials(authToken string) *PomeriumAuthCredentials {
	return &PomeriumAuthCredentials{
		metadata: map[string]string{"authorization": AuthorizationHeader(authToken)},
	}
}

func (p *PomeriumAuthCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return p.metadata, nil
}

func (p *PomeriumAuthCredentials) RequireTransportSecurity() bool {
	return true
}

type options struct {
	tlsConfig *tls.Config
	dialOpts  []grpc.DialOption
}

type Option func(*options)

// Client provides a wrapper interface for all API calls
type Client struct {
	conn *grpc.ClientConn

	ActivityLogService            pb.ActivityLogServiceClient
	ClustersService               pb.ClustersServiceClient
	DeviceService                 pb.DeviceServiceClient
	ExternalDataSourceService     pb.ExternalDataSourceServiceClient
	KeyChainService               pb.KeyChainServiceClient
	NamespacePermissionService    pb.NamespacePermissionServiceClient
	NamespaceService              pb.NamespaceServiceClient
	PolicyService                 pb.PolicyServiceClient
	PomeriumServiceAccountService pb.PomeriumServiceAccountServiceClient
	PomeriumSessionService        pb.PomeriumSessionServiceClient
	RouteService                  pb.RouteServiceClient
	SettingsService               pb.SettingsServiceClient
	UserService                   pb.UserServiceClient
}

// NewClient returns a Pomerium Enterprise client configured to communicate with a given target API
func NewClient(_ context.Context, target string, authToken string, opts ...Option) (*Client, error) {
	conn, err := dial(target, authToken, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	return &Client{
		conn: conn,

		ActivityLogService:            pb.NewActivityLogServiceClient(conn),
		ClustersService:               pb.NewClustersServiceClient(conn),
		DeviceService:                 pb.NewDeviceServiceClient(conn),
		ExternalDataSourceService:     pb.NewExternalDataSourceServiceClient(conn),
		KeyChainService:               pb.NewKeyChainServiceClient(conn),
		NamespacePermissionService:    pb.NewNamespacePermissionServiceClient(conn),
		NamespaceService:              pb.NewNamespaceServiceClient(conn),
		PolicyService:                 pb.NewPolicyServiceClient(conn),
		PomeriumServiceAccountService: pb.NewPomeriumServiceAccountServiceClient(conn),
		PomeriumSessionService:        pb.NewPomeriumSessionServiceClient(conn),
		RouteService:                  pb.NewRouteServiceClient(conn),
		SettingsService:               pb.NewSettingsServiceClient(conn),
		UserService:                   pb.NewUserServiceClient(conn),
	}, nil
}

func dial(target string, authToken string, opts ...Option) (*grpc.ClientConn, error) {
	cfg := &options{}
	for _, o := range opts {
		o(cfg)
	}

	creds := NewPomeriumAuthCredentials(authToken)
	dialOpts := append(cfg.dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(cfg.tlsConfig)), grpc.WithPerRPCCredentials(creds))

	conn, err := grpc.NewClient(target, dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}
	return conn, nil
}

// WithTlsConfig provides a custom tls.Config to the Client
func WithTlsConfig(config *tls.Config) Option {
	return func(cfg *options) {
		cfg.tlsConfig = config
	}
}

// WithDialOption provides a custom grpc.DialOption to the internal grpc client in Client
func WithDialOption(o grpc.DialOption) Option {
	return func(cfg *options) {
		cfg.dialOpts = append(cfg.dialOpts, o)
	}
}
