// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wirex

import (
	"context"
	"github.com/LyricTian/gin-admin/v10/internal/mods"
	"github.com/LyricTian/gin-admin/v10/internal/mods/rbac"
	"github.com/LyricTian/gin-admin/v10/internal/mods/rbac/api"
	"github.com/LyricTian/gin-admin/v10/internal/mods/rbac/biz"
	"github.com/LyricTian/gin-admin/v10/internal/mods/rbac/dal"
	"github.com/LyricTian/gin-admin/v10/internal/utils"
)

// Injectors from wire.go:

func BuildInjector(ctx context.Context) (*Injector, func(), error) {
	db, cleanup, err := InitDB(ctx)
	if err != nil {
		return nil, nil, err
	}
	cacher, cleanup2, err := InitCacher(ctx)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	auther, cleanup3, err := InitAuth(ctx)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	trans := &utils.Trans{
		DB: db,
	}
	menu := &dal.Menu{
		DB: db,
	}
	menuResource := &dal.MenuResource{
		DB: db,
	}
	roleMenu := &dal.RoleMenu{
		DB: db,
	}
	bizMenu := &biz.Menu{
		Trans:           trans,
		MenuDAL:         menu,
		MenuResourceDAL: menuResource,
		RoleMenuDAL:     roleMenu,
	}
	apiMenu := &api.Menu{
		MenuBIZ: bizMenu,
	}
	role := &dal.Role{
		DB: db,
	}
	userRole := &dal.UserRole{
		DB: db,
	}
	bizRole := &biz.Role{
		Cache:       cacher,
		Trans:       trans,
		RoleDAL:     role,
		RoleMenuDAL: roleMenu,
		UserRoleDAL: userRole,
	}
	apiRole := &api.Role{
		RoleBIZ: bizRole,
	}
	user := &dal.User{
		DB: db,
	}
	bizUser := &biz.User{
		Cache:       cacher,
		Trans:       trans,
		UserDAL:     user,
		UserRoleDAL: userRole,
	}
	apiUser := &api.User{
		UserBIZ: bizUser,
	}
	login := &biz.Login{
		Cache:       cacher,
		Auth:        auther,
		UserDAL:     user,
		UserRoleDAL: userRole,
		MenuDAL:     menu,
		UserBIZ:     bizUser,
	}
	apiLogin := &api.Login{
		LoginBIZ: login,
	}
	casbinx := &rbac.Casbinx{
		Cache:           cacher,
		MenuDAL:         menu,
		MenuResourceDAL: menuResource,
		RoleDAL:         role,
	}
	rbacRBAC := &rbac.RBAC{
		DB:       db,
		MenuAPI:  apiMenu,
		RoleAPI:  apiRole,
		UserAPI:  apiUser,
		LoginAPI: apiLogin,
		Casbinx:  casbinx,
	}
	modsMods := &mods.Mods{
		RBAC: rbacRBAC,
	}
	injector := &Injector{
		DB:    db,
		Cache: cacher,
		Auth:  auther,
		M:     modsMods,
	}
	return injector, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
