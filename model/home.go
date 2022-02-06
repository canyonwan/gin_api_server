package model

// CustomSwiper 轮播图
type CustomSwiper struct {
	Id         int64  `gorm:"id primary_key AUTOINCREMENT" json:"id"`
	Title      string `gorm:"title" json:"title,omitempty"`
	ImageUrl   string `gorm:"imageUrl" json:"image_url,omitempty"`
	CreateTime int64  `gorm:"create_time" json:"create_time,omitempty"`
}

// HomeGoods 首页商品列表
type HomeGoods struct {
	Id            int64  `gorm:"id primary_key AUTOINCREMENT" json:"id"`
	Title         string `gorm:"title" json:"title,omitempty"`
	ImageUrl      string `gorm:"imageUrl" json:"imageUrl,omitempty"`
	OriginalPrice int64  `gorm:"originalPrice" json:"price,omitempty"`
	PresentPrice  int64  `gorm:"presentPrice" json:"presentPrice,omitempty"`
	CreateTime    int64  `gorm:"create_time" json:"createTime,omitempty"`
}
