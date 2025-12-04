package errorx

import (
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RPC 错误映射配置
type RpcErrorMapper struct {
	// NotFound 错误时返回的业务错误
	NotFoundErr *BusinessError

	// AlreadyExists 错误时返回的业务错误
	AlreadyExistsErr *BusinessError

	// InvalidArgument 错误时返回的业务错误(可选，默认使用 ErrInvalidParam)
	InvalidArgumentErr *BusinessError
}

func MapRpcError(err error, logger logx.Logger, rpcMethod string, mapper RpcErrorMapper) error {
	if err == nil {
		return nil
	}

	// 尝试从错误中提取 gRPC 状态
	st, ok := status.FromError(err)
	if !ok {
		// 不是 gRPC 错误，记录日志并返回内部错误
		logger.Errorf("调用 %s 失败（非 gRPC 错误）：%v", rpcMethod, err)
		return ErrInternalError
	}

	// 根据 gRPC 状态码映射到业务错误
	switch st.Code() {
	case codes.InvalidArgument:
		if mapper.InvalidArgumentErr != nil {
			return mapper.InvalidArgumentErr
		}
		return ErrInvalidParam

	case codes.NotFound:
		if mapper.NotFoundErr != nil {
			return mapper.NotFoundErr
		}
		// 如果没有配置 NotFoundErr，返回通用 NotFound 错误
		logger.Errorf("调用 %s 失败：未配置 NotFoundErr, code=%v, msg=%s", rpcMethod, st.Code(), st.Message())
		return ErrNotFound
	case codes.AlreadyExists:
		if mapper.AlreadyExistsErr != nil {
			return mapper.AlreadyExistsErr
		}
		logger.Errorf("调用 %s 失败：未配置 AlreadyExistsErr, code=%v, msg=%s", rpcMethod, st.Code(), st.Message())
		return ErrInternalError

	case codes.Internal:
		logger.Errorf("调用 %s 失败（RPC 内部错误）, code=%v, msg=%s", rpcMethod, st.Code(), st.Message())
		return ErrInternalError

	default:
		logger.Errorf("调用 %s 失败（未映射的错误）: code=%v, msg=%s", rpcMethod, st.Code(), st.Message())
		return ErrInternalError
	}
}
