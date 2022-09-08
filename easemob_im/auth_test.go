package easemob_im

import (
	"fmt"
	"testing"
)

func TestNewAuth(t *testing.T) {
	auth, err := NewAuth(DefaultEaseMobConfig.AppKey, DefaultEaseMobConfig.ClientID, DefaultEaseMobConfig.ClientSecret, false)
	if err != nil {
		t.Error(err)
	}
	ret := auth.Headers()
	fmt.Println(ret)
}

func TestAuth_GetAgoraToken(t *testing.T) {
	auth, err := NewAuth(DefaultAgoraConfig.AppKey, DefaultAgoraConfig.AppID, DefaultAgoraConfig.AppCertificate, true)
	if err != nil {
		t.Error(err)
	}
	//auth.SetUUID("test")
	ret := auth.GetAgoraToken()
	t.Log("\nAgoraToken", ret)
	//easeMobToken := auth.GetAgoraToken2easemobToken()
	//t.Log("easeMobToken", easeMobToken)
}
