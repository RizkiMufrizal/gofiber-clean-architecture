package service

import "context"

type HttpBinService interface {
	PostMethod(ctx context.Context)
}
