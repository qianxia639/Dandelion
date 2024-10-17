package model

import "time"

type LoginLog struct {
	LogId      int64     `db:"log_id" json:"log_id,omitempty"`             // 日志Id
	UserId     uint32    `db:"user_id" json:"user_id,omitempty"`           // 用户Id
	Ip         string    `db:"ip" json:"ip,omitempty"`                     // Ip地址
	IpAddr     string    `db:"ip_addr" json:"ip_addr,omitempty"`           // Ip所属地
	IsLoginExp bool      `db:"is_login_exp" json:"is_login_exp,omitempty"` // 登录异常, F: 无异常, T: 异常
	Other      string    `db:"other" json:"other,omitempty"`               // 其他属性
	LoginTime  time.Time `db:"login_time" json:"login_time,omitempty"`     // 登录时间
}
