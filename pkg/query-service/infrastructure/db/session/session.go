package session

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/dao/mongo_dao"
)

type Func func(ctx context.Context) error
type SessionType int

const (
	NoSession SessionType = iota
	ReadSession
	WriteSession
)

type Options struct {
	mongo *SessionType
}

func NewOptions(opts ...*Options) *Options {
	return &Options{}
}

func MergeOptions(opts ...*Options) *Options {
	opt := &Options{}
	w := WriteSession
	opt.mongo = &w

	for _, o := range opts {
		if o.mongo != nil {
			opt.mongo = o.mongo
		}
	}
	return opt
}

func (o *Options) SetMongo(v SessionType) *Options {
	o.mongo = &v
	return o
}

func (o *Options) GetMongo() SessionType {
	if o.mongo == nil {
		return NoSession
	}
	return *o.mongo
}

func StartSession(ctx context.Context, fun Func, opts ...*Options) error {
	opt := MergeOptions(opts...)
	if opt.GetMongo() != NoSession {
		mongoSession := mongo_dao.NewSession(opt.GetMongo() == WriteSession)
		return mongoSession.UseTransaction(ctx, func(ctx context.Context) error {
			return fun(ctx)
		})
	}
	return fun(ctx)
}
