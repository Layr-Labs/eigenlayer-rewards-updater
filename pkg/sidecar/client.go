package sidecar

import (
	"context"
	"crypto/tls"
	"fmt"
	rewardsV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/v1/rewards"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

type IRewardsClient interface {
	GenerateRewards(ctx context.Context, req *rewardsV1.GenerateRewardsRequest, opts ...grpc.CallOption) (*rewardsV1.GenerateRewardsResponse, error)
	GenerateRewardsRoot(ctx context.Context, req *rewardsV1.GenerateRewardsRootRequest, opts ...grpc.CallOption) (*rewardsV1.GenerateRewardsRootResponse, error)
	GenerateClaimProof(ctx context.Context, in *rewardsV1.GenerateClaimProofRequest, opts ...grpc.CallOption) (*rewardsV1.GenerateClaimProofResponse, error)
}

type SidecarClient struct {
	Rewards IRewardsClient
}

func NewSidecarClient(url string, insecureConn bool) (*SidecarClient, error) {
	fmt.Printf("url: %s - insecure %+v\n", url, insecureConn)
	var creds grpc.DialOption
	if strings.Contains(url, "localhost:") || strings.Contains(url, "127.0.0.1:") || insecureConn {
		creds = grpc.WithTransportCredentials(insecure.NewCredentials())
	} else {
		creds = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: false}))
	}

	grpcClient, err := grpc.NewClient(url, creds)
	if err != nil {
		return nil, err
	}

	rewardsClient := rewardsV1.NewRewardsClient(grpcClient)

	return &SidecarClient{
		Rewards: rewardsClient,
	}, nil
}
