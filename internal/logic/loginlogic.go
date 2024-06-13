package logic

import (
	"context"
	"fmt"

	"blog_rpc/internal/svc"
	"blog_rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &user.LoginResponse{
		Code: 0,
		Msg:  "登陆成功",
		Data: fmt.Sprintf("登录的用户名为%v,密码为%v", in.Username, in.Password),
	}, nil
}
