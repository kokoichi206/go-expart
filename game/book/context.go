package main

import "context"

type contextKey struct{}

type RequestContext struct {
	RequestID string
}

func SetInContext(ctx context.Context, reqCtx *RequestContext) context.Context {
	return context.WithValue(ctx, contextKey{}, reqCtx)
}

func Extract(ctx context.Context) *RequestContext {
	value := ctx.Value(contextKey{})
	if value == nil {
		return nil
	}

	return value.(*RequestContext)
}
