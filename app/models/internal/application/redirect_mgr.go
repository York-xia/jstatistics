package application

import (
	"js_statistics/app/models/tables"
	"time"
)

type RedirectManage struct {
	ON         time.Time `gorm:"column:on;type:timestamptz;comment:开启时间"`
	OFF        time.Time `gorm:"column:off;type:timestamptz;comment:关闭时间"`
	PC         string    `gorm:"column:pc;type:varchar(100);not null;comment:pc跳转地址"`
	Android    string    `gorm:"column:android;type:varchar(100);not null;comment:android跳转地址"`
	IOS        string    `gorm:"column:ios;type:varchar(100);not null;comment:ios跳转地址"`
	Title      string    `gorm:"column:title;type:varchar(40);not null;comment:标题"`
	ID         int64     `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	CategoryID int64     `gorm:"column:category_id;type:bigint;not null;comment:js分类id"`
	Status     bool      `gorm:"column:status;type:boolean;;not nulldefault:true;comment:状态"`
}

func (RedirectManage) TableName() string {
	return tables.RedirectManage
}
