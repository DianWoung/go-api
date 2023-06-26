package logic

import (
	"context"

	"open/internal/svc"
	"open/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloWorldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloWorldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloWorldLogic {
	return &HelloWorldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloWorldLogic) HelloWorld() (resp *types.HelloWorldResp, err error) {
	resp = new(types.HelloWorldResp)
	resp.Msg = "Hello World"
	return
}
