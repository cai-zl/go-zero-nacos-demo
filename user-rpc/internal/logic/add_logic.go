package logic

import (
	"context"
	"github.com/google/uuid"

	"go-zero-nacos-demo/user-rpc/internal/svc"
	"go-zero-nacos-demo/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *user.User) (*user.AddResp, error) {
	logx.Infof("新增用户: %s", in.Username)
	return &user.AddResp{Ok: true, UserId: uuid.New().String()}, nil
}
