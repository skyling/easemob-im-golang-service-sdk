package easemob_im

import (
	"fmt"
	"github.com/skyling/easemob-im-golang-service-sdk/types"
)

// https://docs-im.easemob.com/push/apppush/label

// PushLabel 推送标签管理
type PushLabel struct {
	auth *Auth
}

func NewPushLabel(auth *Auth) *PushLabel {
	return &PushLabel{
		auth: auth,
	}
}

// Create 创建标签
func (s *PushLabel) Create(msg *types.PushLabelCreateReq) (*types.PushLabelDataResp, error) {
	uri := s.auth.BuildURI("/push/label")
	var res types.PushLabelDataResp
	err := HttpPost(uri, msg, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Detail 查询推送标签
func (s *PushLabel) Detail(name string) (*types.PushLabelDataResp, error) {
	uri := s.auth.BuildURI(fmt.Sprintf("/push/label/%s", name))
	var res types.PushLabelDataResp
	err := HttpGet(uri, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Lists 查询推送标签列表
func (s *PushLabel) Lists(limit int, cursor string) (*types.PushLabelListResp, error) {
	uri := s.auth.BuildURI(fmt.Sprintf("/push/label?limit=%d&cursor=%s", limit, cursor))
	var res types.PushLabelListResp
	err := HttpGet(uri, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Delete 删除推送标签
func (s *PushLabel) Delete(name string) (*types.PushLabelDeleteResp, error) {
	uri := s.auth.BuildURI(fmt.Sprintf("/push/label/%s", name))
	var res types.PushLabelDeleteResp
	err := HttpDelete(uri, nil, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// AddUsers 在推送标签里添加用户
func (s *PushLabel) AddUsers(name string, usernames []string) (*types.PushLabelAddDeleteUserResp, error) {
	uri := s.auth.BuildURI(fmt.Sprintf("/push/label/%s/user", name))
	var res types.PushLabelAddDeleteUserResp
	req := map[string]interface{}{
		"usernames": usernames,
	}
	err := HttpPost(uri, req, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// LabelUserDetail 根据推送标签和用户 ID 查询用户信息
func (s *PushLabel) LabelUserDetail(name string, username string) (*types.PushLabelUserResp, error) {
	uri := s.auth.BuildURI(fmt.Sprintf("/push/label/%s/user/%s", name, username))
	var res types.PushLabelUserResp
	err := HttpGet(uri, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// LabelUserList 分页查询标签下用户
func (s *PushLabel) LabelUserList(name string, limit int, cursor string) (*types.PushLabelUserListResp, error) {
	uri := s.auth.BuildURI(fmt.Sprintf("/push/label/%s/user?limit=%d&cursor=%s", name, limit, cursor))
	var res types.PushLabelUserListResp
	err := HttpGet(uri, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteUsers 批量移出推送标签下的用户
func (s *PushLabel) DeleteUsers(name string, usernames []string) (*types.PushLabelAddDeleteUserResp, error) {
	uri := s.auth.BuildURI(fmt.Sprintf("/push/label/%s/user", name))
	var res types.PushLabelAddDeleteUserResp
	req := map[string]interface{}{
		"usernames": usernames,
	}
	err := HttpDelete(uri, req, &res, s.auth.Headers())
	if err != nil {
		return nil, err
	}
	return &res, nil
}
