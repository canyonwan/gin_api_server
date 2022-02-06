package controller

import (
	"gin_api_server/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var imageUrl = "https://cdn9-banquan.ituchong.com/weili/l/1397601058514272265.webp"
var imageUrl2 = "https://cdn3-banquan.ituchong.com/weili/l/1304842941827383333.webp"

//
var imageList = &model.CustomSwiper{
	Title:      "第一张图片",
	ImageUrl:   imageUrl,
	CreateTime: time.Now().Unix(),
}
var imageList2 = &model.CustomSwiper{
	Title:      "第二张图片",
	ImageUrl:   imageUrl2,
	CreateTime: time.Now().Unix(),
}

func GetHomeSwiperList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取首页轮播图列表成功",
		"data":    []*model.CustomSwiper{imageList, imageList2},
	})

}

var homeGoods1 = &model.HomeGoods{
	Id:            0,
	Title:         "科星药业-麻杏石甘散1kg/袋",
	ImageUrl:      imageUrl,
	OriginalPrice: 32,
	PresentPrice:  28,
	CreateTime:    time.Now().Unix(),
}
var homeGoods2 = &model.HomeGoods{
	Id:            1,
	Title:         "迈德生物-单过硫酸氢钾1kg/桶",
	ImageUrl:      imageUrl,
	OriginalPrice: 80,
	PresentPrice:  80,
	CreateTime:    time.Now().Unix(),
}
var homeGoods3 = &model.HomeGoods{
	Id:            2,
	Title:         "迈德生物-单过硫酸氢钾100kg/桶",
	ImageUrl:      imageUrl,
	OriginalPrice: 100,
	PresentPrice:  100,
	CreateTime:    time.Now().Unix(),
}

func GetHomeGoods(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取首页商品列表成功",
		"data":    []*model.HomeGoods{homeGoods1, homeGoods2, homeGoods3},
	})
}
