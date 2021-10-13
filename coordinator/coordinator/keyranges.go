package coordinator

import (
	"context"

	shards "github.com/pg-sharding/spqr/router/protos"
)

type KeyRangeService struct {
	shards.UnimplementedKeyRangeServiceServer
}

func (c KeyRangeService) ListKeyRange(ctx context.Context, request *shards.ListKeyRangeRequest) (*shards.KeyRangeReply, error) {
	panic("implement me")
}

func (c KeyRangeService) LockKeyRange(ctx context.Context, request *shards.LockKeyRangeRequest) (*shards.KeyRangeReply, error) {
	panic("implement me")
}

func (c KeyRangeService) UnlockKeyRange(ctx context.Context, request *shards.UnlockKeyRangeRequest) (*shards.KeyRangeReply, error) {
	panic("implement me")
}

func (c KeyRangeService) SplitKeyRange(ctx context.Context, request *shards.SplitKeyRangeRequest) (*shards.KeyRangeReply, error) {
	panic("implement me")
}

var _ shards.KeyRangeServiceServer = KeyRangeService{}