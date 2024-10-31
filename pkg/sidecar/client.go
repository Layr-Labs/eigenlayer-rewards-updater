package sidecar

import (
	"context"
	"crypto/tls"
	"github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

type IRewardsClient interface {
	GenerateRewards(ctx context.Context, req *v1.GenerateRewardsRequest, opts ...grpc.CallOption) (*v1.GenerateRewardsResponse, error)
	GenerateRewardsRoot(ctx context.Context, req *v1.GenerateRewardsRootRequest, opts ...grpc.CallOption) (*v1.GenerateRewardsRootResponse, error)
}

type SidecarClient struct {
	Rewards IRewardsClient
}

func NewSidecarClient(url string) (*SidecarClient, error) {
	var creds grpc.DialOption
	if strings.Contains(url, "localhost:") || strings.Contains(url, "127.0.0.1:") {
		creds = grpc.WithTransportCredentials(insecure.NewCredentials())
	} else {
		creds = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: false}))
	}

	grpcClient, err := grpc.NewClient(url, creds)
	if err != nil {
		return nil, err
	}

	rewardsClient := v1.NewRewardsClient(grpcClient)

	return &SidecarClient{
		Rewards: rewardsClient,
	}, nil
}
