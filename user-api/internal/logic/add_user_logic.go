package logic

import (
	"context"
	"go-zero-nacos-demo/user-rpc/user"

	"go-zero-nacos-demo/user-api/internal/svc"
	"go-zero-nacos-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.User) (resp *types.AddUserResp, err error) {
	add, err := l.svcCtx.UserRpc.Add(l.ctx, &user.User{Username: req.Username, Age: req.Age})
	return &types.AddUserResp{
		UserId: add.UserId, Ok: add.Ok,
	}, err
}
