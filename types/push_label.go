package types

type PushLabelCreateReq struct {
	Name        string `json:"name"`                  // 标签名称，限制 64 字符以内。
	Description string `json:"description,omitempty"` // 标签描述 内容，限制 255 字符以内。
}

type PushLabelData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
}

type PushLabelDataResp struct {
	Timestamp int64          `json:"timestamp"`
	Data      *PushLabelData `json:"data"`
	Duration  int            `json:"duration"`
}

type PushLabelListResp struct {
	Timestamp int64            `json:"timestamp"`
	Data      []*PushLabelData `json:"data"`
	Duration  int              `json:"duration"`
}

type PushLabelDeleteResp struct {
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
	Duration  int    `json:"duration"`
}

type PushLabelAddDeleteUserResp struct {
	Timestamp int64 `json:"timestamp"`
	Data      struct {
		Success []string `json:"success"`
		Fail    []string `json:"fail"`
	} `json:"data"`
	Duration int `json:"duration"`
}

type PushLabelUserData struct {
	Username string `json:"username"`
	Created  int64  `json:"created"`
}
type PushLabelUserResp struct {
	Timestamp int64              `json:"timestamp"`
	Data      *PushLabelUserData `json:"data"`
	Duration  int                `json:"duration"`
}

type PushLabelUserListResp struct {
	Timestamp int64                `json:"timestamp"`
	Data      []*PushLabelUserData `json:"data"`
	Duration  int                  `json:"duration"`
}
