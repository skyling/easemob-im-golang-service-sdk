package easemob_im

import (
	"fmt"
	"github.com/skyling/easemob-im-golang-service-sdk/types"
)

// https://docs-im.easemob.com/push/apppush/integration

// Push 推送服务
type Push struct {
	auth *Auth
}

func NewPush(auth *Auth) *Push {
	return &Push{
		auth: auth,
	}
}

// Single 使用单接口批量发送推送消息
func (s *Push) Single(msg *types.PushSingleReq) (*types.PushSingleResp, error) {
	uri := s.auth.BuildURI("/push/single")
	var res types.PushSingleResp
	err := HttpPost(uri, msg, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Message 创建推送消息
func (s *Push) Message(msg *types.PushDataBaseMessage) (*types.PushMessageResp, error) {
	uri := s.auth.BuildURI("/push/message")
	var res types.PushMessageResp
	err := HttpPost(uri, msg, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetMessage 查询推送消息
func (s *Push) GetMessage(id string) (*types.GetPushMessageResp, error) {
	uri := s.auth.BuildURI(fmt.Sprintf("/push/message/%s", id))
	var res types.GetPushMessageResp
	err := HttpGet(uri, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// TaskBroadcast 创建全局推送任务
func (s *Push) TaskBroadcast(msg *types.PushTaskBroadcastReq) (*types.PushMessageResp, error) {
	uri := s.auth.BuildURI("/push/task/broadcast")
	var res types.PushMessageResp
	err := HttpPost(uri, msg, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Task 创建推送任务
func (s *Push) Task(msg *types.PushTaskReq) (*types.PushMessageResp, error) {
	uri := s.auth.BuildURI("/push/task")
	var res types.PushMessageResp
	err := HttpPost(uri, msg, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// List 获取标签推送列表
func (s *Push) List(msg *types.PushListReq) (*types.PushListResp, error) {
	uri := s.auth.BuildURI("/push/list")
	var res types.PushListResp
	err := HttpPost(uri, msg, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}
