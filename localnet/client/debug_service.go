package client

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	apitypes "github.com/spacemeshos/api/release/go/spacemesh/v1"
)

func (c *gRPCClient) DebugAllAccounts() ([]*apitypes.Account, error) {
	dbgService := c.getDebugServiceClient()
	resp, err := dbgService.Accounts(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return resp.AccountWrapper, nil
}
