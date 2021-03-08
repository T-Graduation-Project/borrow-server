package service

import (
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gvalid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// 拦截器，类似中间件作用。
var Interceptor = new(serviceInterceptor)

type serviceInterceptor struct{}

// 统一StructTag校验器
func (s *serviceInterceptor) UnaryValidate(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 如果没有StructTag，那么校验器什么都不会做
	if err := gvalid.CheckStruct(req, nil); err != nil {
		return nil, gerror.NewCode(
			int(codes.InvalidArgument),
			gerror.Current(err).Error(),
		)
	}
	return handler(ctx, req)
}
