package easemob_im

import (
	"fmt"
	"github.com/skyling/easemob-im-golang-service-sdk/types"
	"testing"
)

func TestPush_Single(t *testing.T) {
	auth, err := NewAuth(DefaultEaseMobConfig.AppKey, DefaultEaseMobConfig.ClientID, DefaultEaseMobConfig.ClientSecret, false)
	if err != nil {
		t.Error(err)
	}
	sign := NewPush(auth)
	msg := types.PushSingleReq{
		Targets:  []string{"D5FA241D0A890CC12B38495771DC2661"},
		Async:    false,
		Strategy: 3,
		PushMessage: types.PushDataBaseMessage{
			Title:   "后端测试推送消息",
			Content: "请及时查看",
		},
	}
	res, err := sign.Single(&msg)
	fmt.Printf("%+v %+v", res, err)
}
