package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	pb "github.com/pomerium/enterprise-client-go/pb"

	client "github.com/pomerium/enterprise-client-go"
)

var serviceAccountToken = os.Getenv("SERVICE_ACCOUNT")
var target = "console-api.localhost.pomerium.io:443"

func main() {
	err := run()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func run() error {

	ctx := context.Background()

	tlsConfig := &tls.Config{InsecureSkipVerify: true}

	p, err := client.NewClient(ctx, target, serviceAccountToken, client.WithTlsConfig(tlsConfig))
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	nsResp, err := p.NamespaceService.ListNamespaces(ctx, &pb.ListNamespacesRequest{})
	if err != nil {
		return fmt.Errorf("could not list namespaces: %w", err)
	}

	var productionNamespaceId string
	for _, n := range nsResp.GetNamespaces() {
		if n.GetName() == "Production" {
			productionNamespaceId = n.GetId()
		}
	}

	if productionNamespaceId == "" {
		return fmt.Errorf("could not find production namespace")
	}

	policyName := "my policy"
	var policyId string
	polResp, err := p.PolicyService.ListPolicies(ctx, &pb.ListPoliciesRequest{Namespace: productionNamespaceId, Query: &policyName})
	if err != nil {
		return fmt.Errorf("failed to find policy: %w", err)
	}
	if len(polResp.GetPolicies()) == 0 {
		return fmt.Errorf("no policy named '%s' found", policyName)
	}

	policyId = polResp.GetPolicies()[0].GetId()

	passIdHeaders := true
	newRoute := &pb.Route{
		NamespaceId:         productionNamespaceId,
		Name:                "my route",
		From:                "https://test.localhost.pomerium.io",
		To:                  []string{"https://verify.pomerium.com"},
		PolicyIds:           []string{policyId},
		PassIdentityHeaders: &passIdHeaders,
	}

	routeResp, err := p.RouteService.SetRoute(ctx, &pb.SetRouteRequest{Route: newRoute})
	if err != nil {
		return fmt.Errorf("could not create route: %w", err)
	}

	fmt.Printf("created route id: %s\n", routeResp.Route.GetId())
	return nil
}
