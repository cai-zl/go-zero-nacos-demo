// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"go-zero-nacos-demo/user-rpc/internal/logic"
	"go-zero-nacos-demo/user-rpc/internal/svc"
	"go-zero-nacos-demo/user-rpc/user"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) Add(ctx context.Context, in *user.User) (*user.AddResp, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}
