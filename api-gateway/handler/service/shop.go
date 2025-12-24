package service

import (
	"api-gateway/basic/globals"
	__ "api-gateway/basic/proto"
	"api-gateway/handler/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var req request.Register
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	register, err := globals.ShopClient.Register(c, &__.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "注册失败",
			"data": nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "注册成功",
			"data": register,
		})
		return
	}
}
func Login(c *gin.Context) {
	var req request.Login
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)
	register, err := globals.ShopClient.Login(c, &__.LoginReq{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "登录失败",
			"data": nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "登录成功",
			"data": register,
		})
		return
	}
}
func FindGoodList(c *gin.Context) {
	list, err := globals.ShopClient.FindGoodList(c, &__.FindGoodListReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询商品列表失败",
			"data": nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询商品列表成功",
			"data": list,
		})
		return
	}
}
func FindBannerList(c *gin.Context) {
	list, err := globals.ShopClient.FindBannerList(c, &__.FindBannerListReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "查询轮播图列表失败",
			"data": nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询轮播图列表成功",
			"data": list,
		})
		return
	}
}
