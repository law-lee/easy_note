package mw

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = ClientMiddleware

func ClientMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		// pre-processing logic can be added here
		ri := rpcinfo.GetRPCInfo(ctx)
		// get server info
		klog.Infof("server address: %v, rpc timeout: %v, readwrite timeout: %v\n", ri.To().Address(),
			ri.Config().RPCTimeout(), ri.Config().ConnectTimeout())

		if err := next(ctx, request, response); err != nil {
			// post-processing logic can be added here
			return err
		}
		return nil
	}
}
