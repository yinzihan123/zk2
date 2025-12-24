package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"math/rand"
	"shop-srv/basic/globals"
	__ "shop-srv/basic/proto"
	"shop-srv/handler/dao"
	"shop-srv/handler/models"
	"shop-srv/pkg"
	"strconv"
)

type Server struct {
	__.UnimplementedShopServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Register(_ context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {
	shopDao := dao.NewShopDao()
	mobile, err := shopDao.FindUserByMobile(globals.DB, in.Mobile)
	if err != nil {
		return nil, err
	}
	if mobile.ID != 0 {
		return nil, errors.New("该用户已存在")
	}
	code := rand.Intn(9000) + 1000
	salt := strconv.Itoa(code)
	user := &models.User{
		Model:    gorm.Model{},
		Mobile:   in.Mobile,
		Password: pkg.Md5Str(in.Password + salt),
		Salt:     salt,
	}
	add, err := shopDao.UserAdd(globals.DB, user)
	if err != nil {
		return nil, err
	}
	handler, err := pkg.TokenHandler(strconv.Itoa(int(add.ID)))
	if err != nil {
		return nil, err
	}
	return &__.RegisterResp{
		Token: handler,
	}, nil
}
func (s *Server) Login(_ context.Context, in *__.LoginReq) (*__.LoginResp, error) {
	shopDao := dao.NewShopDao()
	mobile, err := shopDao.FindUserByMobile(globals.DB, in.Mobile)
	if err != nil {
		return nil, err
	}
	if mobile.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	if mobile.Password != pkg.Md5Str(in.Password+mobile.Salt) {
		return nil, errors.New("密码错误")
	}
	handler, err := pkg.TokenHandler(strconv.Itoa(int(mobile.ID)))
	if err != nil {
		return nil, err
	}
	return &__.LoginResp{
		Token:    handler,
		Mobile:   mobile.Mobile,
		Password: mobile.Password,
	}, nil
}
func (s *Server) FindGoodList(_ context.Context, in *__.FindGoodListReq) (*__.FindGoodListResp, error) {
	shopDao := dao.NewShopDao()
	list, err := shopDao.FindGoodList(globals.DB)
	if err != nil {
		return nil, err
	}
	var goodLists []*__.GoodList
	for _, good := range list {
		goodList := &__.GoodList{
			Id:            int64(good.ID),
			Name:          good.Name,
			Price:         float32(good.Price),
			OriginalPrice: float32(good.OriginalPrice),
			SalesVolume:   float32(good.SalesVolume),
			ShopName:      good.ShopName,
			Inventory:     int64(good.Inventory),
			Pic:           good.Pic,
		}
		goodLists = append(goodLists, goodList)
	}
	return &__.FindGoodListResp{
		List: goodLists,
	}, nil
}
func (s *Server) FindBannerList(_ context.Context, in *__.FindBannerListReq) (*__.FindBannerListResp, error) {
	shopDao := dao.NewShopDao()
	list, err := shopDao.FindBannerList(globals.DB)
	if err != nil {
		return nil, err
	}
	var bannerLists []*__.BannerList
	for _, banner := range list {
		bannerList := &__.BannerList{
			Id:    int64(banner.ID),
			Pic:   banner.Pic,
			Title: banner.Title,
		}
		bannerLists = append(bannerLists, bannerList)
	}
	return &__.FindBannerListResp{
		List: bannerLists,
	}, nil
}
