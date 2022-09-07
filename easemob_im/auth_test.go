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
