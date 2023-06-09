package user

import (
	"context"
	"data_source_management_center/apps/user/cmd/api/internal/svc"
	"data_source_management_center/apps/user/cmd/api/internal/types"
	"data_source_management_center/apps/user/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	registerRes, err := l.svcCtx.UserRpc.Register(l.ctx, &usercenter.RegisterReq{
		Username: req.UserName,
		Password: req.Password,
		Email:    req.Email,
		Sex:      req.Sex,
		Info:     req.Info,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		AccessToken:  registerRes.AccessToken,
		AccessExpire: registerRes.AccessExpire,
		RefreshAfter: registerRes.RefreshAfter,
	}, nil
}
