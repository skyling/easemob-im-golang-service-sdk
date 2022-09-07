package easemob_im

import (
	"github.com/skyling/easemob-im-golang-service-sdk/types"
	"testing"
)

func GetUser(t *testing.T) *User {
	auth, err := NewAuth(DefaultEaseMobConfig.AppKey, DefaultEaseMobConfig.ClientID, DefaultEaseMobConfig.ClientSecret, false)
	if err != nil {
		t.Error(err)
	}
	return NewUser(auth)
}
func TestUser_Create(t *testing.T) {
	users, err := GetUser(t).Create(types.UserCreateReq{
		Username: "test1",
		Password: "123",
		Nickname: "test",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(users)
}
