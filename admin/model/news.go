package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type News struct {
	Id           int64      `json:"id" gorm:"column:id;primaryKey;not null;autoIncrement;" `
	CreateTime   *time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime   *time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	StockID      *int64     `json:"stock_id" gorm:"column:stock_id;type:bigint(20);index;comment:股票相关ID"`
	Security     *string    `json:"security" gorm:"column:security;type:varchar(40);index;comment:股票代码"`
	NewsId       string     `json:"news_id" gorm:"column:news_id;type:varchar(50);not null;uniqueIndex;comment:新闻唯一ID"`
	Category     string     `json:"category" gorm:"column:category;type:varchar(50);index;comment:类别"`
	NewsCategory *string    `json:"news_category" gorm:"column:news_category;type:varchar(50);comment:新闻类别"`
	Datetime     int64      `json:"datetime" gorm:"column:datetime;type:int(11);not null;comment:发布时间"`
	Headline     string     `json:"headline" gorm:"column:headline;type:varchar(255);not null;comment:标题"`
	Image        *string    `json:"image" gorm:"column:image;type:text;comment:缩略图"`
	Related      *string    `json:"related" gorm:"column:related;type:text;comment:相关股票"`
	Source       *string    `json:"source" gorm:"column:source;type:varchar(255);comment:来源"`
	Summary      *string    `json:"summary" gorm:"column:summary;type:text;comment:摘要"`
	Url          *string    `json:"url" gorm:"column:url;type:text;comment:原始文章地址"`
	Content      *string    `json:"content" gorm:"column:content;type:text;comment:文章内容"`
	Show         bool       `json:"show" gorm:"column:show;default:true;comment:是否展示"`
	Likes        int64      `json:"likes" gorm:"column:likes;default:0;comment:点赞数量"`
	Language     string     `json:"language" gorm:"column:language;default:en;comment:语言"`
}

func (News) TableName() string {
	return "news"
}

type NewsInspiration struct {
	Id         int64           `json:"id" gorm:"column:id;primaryKey;not null;autoIncrement;"`
	CreateTime *time.Time      `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime *time.Time      `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	NewsId     int64           `json:"news_id" gorm:"column:news_id;comment:新闻ID"`
	Date       string          `json:"t_date" gorm:"column:t_date;type:date;comment:日期"`
	Price      decimal.Decimal `json:"price" gorm:"column:price;type:float;comment:预测价格"`
	TopPrice   decimal.Decimal `json:"top_price" gorm:"column:top_price;type:float;comment:价格高位"`
	LowPrice   decimal.Decimal `json:"low_price" gorm:"column:low_price;type:float;comment:低位价格"`
	Security   string          `json:"security" gorm:"column:security;type:varchar(20);comment:股票代码"`
	Handler    bool            `json:"handler" gorm:"column:handler;comment:是否已被处理"`
	Question   string          `json:"question" gorm:"column:question;comment:AI提问"`
	Answer     string          `json:"answer" gorm:"column:answer;comment:AI回答"`
}

func (NewsInspiration) TableName() string {
	return "news_inspiration"
}

type MarketNews struct {
	Id         int64      `json:"id" types:"" text:"" json:"id" form:"id" range:"" edit:"0"`                               //
	Title      string     `json:"title" comment:"标题" types:"" text:"" json:"id" form:"id" range:"" edit:"0"`               //
	Content    string     `json:"content" comment:"内容" types:"" text:"" json:"security" form:"security" range:"" edit:"0"` //
	CreateTime *time.Time `json:"create_time" edit:"0"`                                                                    //
	UpdateTime *time.Time `json:"update_time"`                                                                             //
	Remarks    string     `json:"remarks"`                                                                                 //
	Image      string     `json:"image"`                                                                                   //
	Url        string     `json:"url"`                                                                                     //
	Like       int        `json:"like" json:"like" comment:"点赞数" types:"" text:"" json:"like" form:"like" range:""`
}

func (MarketNews) TableName() string {
	return "market_news"
}
