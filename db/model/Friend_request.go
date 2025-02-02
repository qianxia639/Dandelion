package model

import (
	"encoding/json"
	"time"
)

type FriendRequest struct {
	Id          uint32    `db:"id" json:"id,omitempty"`                     // 好友请求Id
	FromUserId  uint32    `db:"from_user_id" json:"from_user_id,omitempty"` // 请求者Id
	ToUserId    uint32    `db:"to_user_id" json:"to_user_id,omitempty"`     // 接收者Id
	Status      int8      `db:"status" json:"status"`                       // 请求状态, 1: 待处理, 2: 已同意, 3:已拒绝, 4: 已过期
	RequestDesc string    `db:"request_desc" json:"request_desc,omitempty"` // 请求信息
	RequestedAt time.Time `db:"requested_at" json:"requested_at,omitempty"` // 请求时间
	ChangedAt   time.Time `db:"changed_at" json:"changed_at,omitempty"`     // 变更时间
}

func (fr *FriendRequest) MarshalBinary() ([]byte, error) {
	return json.Marshal(fr)
}

func (fr *FriendRequest) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, fr)
}
