package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/grpc_testing"
)

func Test_PomeriumAuthCredentials(t *testing.T) {
	var token = "mysecuretesttoken"

	creds := NewPomeriumAuthCredentials(token)
	meta, err := creds.GetRequestMetadata(context.Background(), "/something")
	require.NoError(t, err)
	authHeader, ok := meta["authorization"]
	assert.True(t, ok)
	assert.Equal(t, authHeader, AuthorizationHeader(token))
}

func Test_dial(t *testing.T) {
	var token = "mysecuretesttoken"

	srv := grpc.NewServer()
	h := httptest.NewUnstartedServer(srv)
	h.EnableHTTP2 = true
	h.StartTLS()
	defer h.Close()

	hURL, err := url.Parse(h.URL)
	require.NoError(t, err)
	hCA := x509.NewCertPool()
	hCA.AddCert(h.Certificate())

	testService := &testServiceServer{}
	srv.RegisterService(&grpc_testing.TestService_ServiceDesc, testService)

	ctx := context.Background()

	conn, err := dial(ctx, hURL.Host, token, WithTlsConfig(&tls.Config{RootCAs: hCA}))

	require.NoError(t, err)
	defer conn.Close()

	client := grpc_testing.NewTestServiceClient(conn)
	_, err = client.EmptyCall(context.Background(), &grpc_testing.Empty{})

	assert.Contains(t, testService.md.Get("authorization"), AuthorizationHeader(token))
	assert.NoError(t, err)
}

type testServiceServer struct {
	grpc_testing.TestServiceServer
	md metadata.MD
}

func (t *testServiceServer) EmptyCall(ctx context.Context, msg *grpc_testing.Empty) (*grpc_testing.Empty, error) {
	t.md, _ = metadata.FromIncomingContext(ctx)
	return &grpc_testing.Empty{}, nil
}
