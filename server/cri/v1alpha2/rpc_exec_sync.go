package v1alpha2

import (
	"context"

	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

func (c *service) ExecSync(
	ctx context.Context, req *pb.ExecSyncRequest,
) (*pb.ExecSyncResponse, error) {
	return nil, nil
}
