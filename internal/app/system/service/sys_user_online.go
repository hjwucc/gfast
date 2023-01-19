// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type (
	ISysUserOnline interface {
		Invoke(ctx context.Context, params *model.SysUserOnlineParams)
		SaveOnline(ctx context.Context, params *model.SysUserOnlineParams)
		CheckUserOnline(ctx context.Context)
		GetOnlineListPage(ctx context.Context, req *system.SysUserOnlineSearchReq, hasToken ...bool) (res *system.SysUserOnlineSearchRes, err error)
		UserIsOnline(ctx context.Context, token string) bool
		DeleteOnlineByToken(ctx context.Context, token string) (err error)
		ForceLogout(ctx context.Context, ids []int) (err error)
		GetInfosByIds(ctx context.Context, ids []int) (onlineList []*entity.SysUserOnline, err error)
	}
)

var (
	localSysUserOnline ISysUserOnline
)

func SysUserOnline() ISysUserOnline {
	if localSysUserOnline == nil {
		panic("implement not found for interface ISysUserOnline, forgot register?")
	}
	return localSysUserOnline
}

func RegisterSysUserOnline(i ISysUserOnline) {
	localSysUserOnline = i
}