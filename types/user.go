package types

import "fmt"

type UserCreateReq struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Nickname string `yaml:"nickname" json:"nickname"`
}

type BaseResp struct {
	Path      string `json:"path,omitempty"`
	Uri       string `json:"uri,omitempty"`
	Timestamp int64  `json:"timestamp"`
	Action    string `json:"action,omitempty"`
	Duration  int    `json:"duration,omitempty"`

	ApplicationName string `json:"applicationName,omitempty"`
	Organization    string `json:"organization,omitempty"`
	Application     string `json:"application,omitempty"`
}

type UserEntity struct {
	Uuid      string `json:"uuid"`
	Type      string `json:"type"`
	Created   int64  `json:"created"`
	Modified  int64  `json:"modified"`
	Username  string `json:"username"`
	Activated bool   `json:"activated"`
}

type UserResp struct {
	BaseResp
	Entities []*UserEntity `json:"entities"`
}

type UserListResp struct {
	BaseResp
	Entities []*UserEntity `json:"entities"`
	Cursor   string        `json:"cursor"`
	Count    int           `json:"count"`
}

type ListUserReq struct {
	Limit     int64
	Cursor    string
	Activated bool
}

var ListUserReqDefault = ListUserReq{
	Limit:     10,
	Cursor:    "",
	Activated: true,
}

func (s *ListUserReq) BuildQuery() string {
	if s.Limit <= 0 {
		s.Limit = 10
	}
	if s.Limit > 100 {
		s.Limit = 100
	}
	activated := 0
	if s.Activated {
		activated = 1
	}
	return fmt.Sprintf("limit=%d&cursor=%s&activated=%d", s.Limit, s.Cursor, activated)
}

type UserResultResp struct {
	BaseResp
	Data struct {
		Result bool `json:"result"`
	} `json:"data"`
}

type UserStatusResp struct {
	BaseResp
	Data map[string]string `json:"data"`
}

type UsersStatusResp struct {
	BaseResp
	Data []map[string]string `json:"data"`
}

// UserMutesReq 禁言
type UserMutesReq struct {
	Username  string `json:"username"`
	Chat      int    `json:"chat"` // 0单聊消息禁言时长，单位为秒，最大值为 2147483647。 - > 0：该用户 ID 具体的单聊消息禁言时长。 - 0：取消该用户的单聊消息禁言。 - -1：该用户被设置永久单聊消息禁言。
	GroupChat int    `json:"groupchat"`
	Chatroom  int    `json:"chatroom"`
}

type UserMutesData struct {
	Userid    string `json:"userid"`
	Chat      int    `json:"chat"`
	Groupchat int    `json:"groupchat"`
	Chatroom  int    `json:"chatroom"`
	Unixtime  int    `json:"unixtime"`
}
type UserMutesResp struct {
	BaseResp
	Data *UserMutesData `json:"data"`
}

// UserOfflineMsgCountResp 用户离线消息个数
type UserOfflineMsgCountResp struct {
	BaseResp
	Data map[string]int64 `json:"data"`
}
