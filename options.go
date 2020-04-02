package xormzap

import (
	"context"

	"go.uber.org/zap/zapcore"
)

type options struct {
	contextFunc ContextToFields
}

type Option func(*options)

type ContextToFields func(ctx context.Context) []zapcore.Field

func WithContextFields(f ContextToFields) Option {
	return func(o *options) {
		o.contextFunc = f
	}
}

func evaluateOpt(opts []Option) *options {
	optCopy := &options{}
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}
